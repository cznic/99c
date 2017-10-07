// Copyright 2017 The 99c Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"bufio"
	"flag"
	"fmt"
	"go/token"
	"os"
	"path/filepath"
	"runtime"
	"runtime/debug"

	"github.com/cznic/cc"
	"github.com/cznic/ccir"
	"github.com/cznic/ir"
	"github.com/cznic/virtual"
	"github.com/cznic/xc"
)

var exitStatus = 1

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Fprintf(os.Stderr, "PANIC: %v\n%s\n", err, debug.Stack())
			os.Exit(exitStatus)
		}
	}()

	t := newTask()
	flag.BoolVar(&t.flags.E, "E", false, "Copy C-language source files to standard output, executing all preprocessor directives; no compilation shall be performed. If any operand is not a text file, the effects are unspecified.")
	flag.BoolVar(&t.flags.c, "c", false, "Suppress the link-edit phase of the compilation, and do not remove any object files that are produced.")
	flag.BoolVar(&t.flags.lib, "99lib", false, "Library link mode.")
	flag.StringVar(&t.flags.o, "o", "", "Use the pathname outfile, instead of the default a.out, for the executable file produced. If the -o option is present with -c or -E, the result is unspecified.")
	flag.Parse()
	if err := t.main(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
	}
}

type flags struct {
	E   bool   // -E
	c   bool   // -c
	lib bool   // -99lib
	o   string // -o
}

type task struct {
	flags  flags
	args   []string
	cfiles []string
	ofiles []string
}

func fatalError(msg string, arg ...interface{}) error {
	return fmt.Errorf("%s: fatal error: %s", os.Args[0], fmt.Sprintf(msg, arg...))
}

func newTask() *task { return &task{} }

func (t *task) main() error {
	t.args = flag.Args()
	if len(t.args) == 0 {
		return fatalError("no input files")
	}

	if t.flags.o != "" && (t.flags.c || t.flags.E) && len(t.args) > 1 {
		exitStatus = 2
		return fatalError("cannot specify -o with -c or -E with multiple files")
	}

	for _, arg := range t.args {
		switch filepath.Ext(arg) {
		case ".c":
			t.cfiles = append(t.cfiles, arg)
		case ".o":
			t.ofiles = append(t.ofiles, arg)
		default:
			return fatalError("unrecognized file type: %v", arg)
		}
	}

	switch {
	case t.flags.E:
		model, err := ccir.NewModel()
		if err != nil {
			fatalError("%v", err)
		}

		o := os.Stdout
		if fn := t.flags.o; fn != "" {
			if o, err = os.Create(fn); err != nil {
				fatalError("%v\n", err)
			}
		}
		out := bufio.NewWriter(o)

		defer out.Flush()

		var lpos token.Position
		_, err = cc.Parse(
			fmt.Sprintf(`
#define __arch__ %s
#define __os__ %s
#include <builtin.h>
`, runtime.GOARCH, runtime.GOOS),
			append(t.cfiles, ccir.CRT0Path),
			model,
			cc.SysIncludePaths([]string{ccir.LibcIncludePath}),
			cc.AllowCompatibleTypedefRedefinitions(),
			cc.Cpp(func(toks []xc.Token) {
				if len(toks) != 0 {
					p := toks[0].Position()
					if p.Filename != lpos.Filename {
						fmt.Fprintf(out, "# %d %q\n", p.Line, p.Filename)
					}
					lpos = p
				}
				for _, v := range toks {
					fmt.Fprintf(out, "%v ", cc.TokSrc(v))
				}
				fmt.Fprintln(out)
			}),
		)
		return err
	}

	var obj [][]ir.Object
	for _, fn := range t.ofiles {
		f, err := os.Open(fn)
		if err != nil {
			return fatalError("%v", err)
		}

		r := bufio.NewReader(f)
		var o ir.Objects
		if _, err := o.ReadFrom(r); err != nil {
			return fatalError("%v", err)
		}

		obj = append(obj, o)
	}

	switch {
	case t.flags.c:
		for _, arg := range t.cfiles {
			model, err := ccir.NewModel()
			if err != nil {
				fatalError("%v", err)
			}

			tu, err := cc.Parse(
				fmt.Sprintf(`
#define __arch__ %s
#define __os__ %s
#include <builtin.h>
`, runtime.GOARCH, runtime.GOOS),
				[]string{arg},
				model,
				cc.SysIncludePaths([]string{ccir.LibcIncludePath}),
				cc.AllowCompatibleTypedefRedefinitions(),
			)
			if err != nil {
				return err
			}

			o, err := ccir.New(tu)
			if err != nil {
				return err
			}

			fn := arg[:len(arg)-len(filepath.Ext(arg))] + ".o"
			f, err := os.Create(fn)
			if err != nil {
				return err
			}

			w := bufio.NewWriter(f)
			if _, err := ir.Objects(o).WriteTo(w); err != nil {
				return err
			}

			if err := w.Flush(); err != nil {
				return err
			}

			if err := f.Close(); err != nil {
				return err
			}
		}
		return nil
	default:
		model, err := ccir.NewModel()
		if err != nil {
			fatalError("%v", err)
		}

		tu, err := cc.Parse(
			fmt.Sprintf(`
#define __arch__ %s
#define __os__ %s
#include <builtin.h>
`, runtime.GOARCH, runtime.GOOS),
			append(t.cfiles, ccir.CRT0Path),
			model,
			cc.SysIncludePaths([]string{ccir.LibcIncludePath}),
			cc.AllowCompatibleTypedefRedefinitions(),
		)
		if err != nil {
			return err
		}

		o, err := ccir.New(tu)
		if err != nil {
			return err
		}

		var out []ir.Object
		switch {
		case t.flags.lib:
			out, err = ir.LinkLib(append(obj, o)...)
		default:
			out, err = ir.LinkMain(append(obj, o)...)
		}
		if err != nil {
			return err
		}

		bin, err := virtual.LoadMain(out)
		if err != nil {
			return err
		}

		fn := t.flags.o
		if fn == "" {
			fn = "a.out"
		}

		f, err := os.Create(fn)
		if err != nil {
			return err
		}

		if runtime.GOOS == "linux" {
			f.WriteString("#!/usr/bin/env 99run\n")
		}

		if _, err := bin.WriteTo(f); err != nil {
			return err
		}

		if err := f.Close(); err != nil {
			return err
		}

		fi, err := os.Stat(fn)
		if err != nil {
			return err
		}

		m := fi.Mode()
		for k, b := os.FileMode(0400), os.FileMode(0100); k != 0; k, b = k>>3, b>>3 {
			if m&k != 0 {
				m |= b
			}
		}

		return os.Chmod(fn, m)
	}
}
