// Copyright 2017 The 99c Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
	"testing"

	"github.com/cznic/ir"
	"github.com/cznic/virtual"
	"github.com/cznic/xc"
)

func caller(s string, va ...interface{}) {
	if s == "" {
		s = strings.Repeat("%v ", len(va))
	}
	_, fn, fl, _ := runtime.Caller(2)
	fmt.Fprintf(os.Stderr, "# caller: %s:%d: ", path.Base(fn), fl)
	fmt.Fprintf(os.Stderr, s, va...)
	fmt.Fprintln(os.Stderr)
	_, fn, fl, _ = runtime.Caller(1)
	fmt.Fprintf(os.Stderr, "# \tcallee: %s:%d: ", path.Base(fn), fl)
	fmt.Fprintln(os.Stderr)
	os.Stderr.Sync()
}

func dbg(s string, va ...interface{}) {
	if s == "" {
		s = strings.Repeat("%v ", len(va))
	}
	_, fn, fl, _ := runtime.Caller(1)
	fmt.Fprintf(os.Stderr, "# dbg %s:%d: ", path.Base(fn), fl)
	fmt.Fprintf(os.Stderr, s, va...)
	fmt.Fprintln(os.Stderr)
	os.Stderr.Sync()
}

func TODO(...interface{}) string { //TODOOK
	_, fn, fl, _ := runtime.Caller(1)
	return fmt.Sprintf("# TODO: %s:%d:\n", path.Base(fn), fl) //TODOOK
}

func use(...interface{}) {}

func init() {
	use(caller, dbg, TODO) //TODOOK
}

// ============================================================================

// https://github.com/cznic/99c/issues/4
func TestIssue4(t *testing.T) {
	src, err := filepath.Abs("testdata/issue4.c")
	if err != nil {
		t.Fatal(err)
	}

	wd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	defer func() {
		if err := os.Chdir(wd); err != nil {
			t.Fatal(err)
		}
	}()

	dir, err := ioutil.TempDir("", "99c-test-")
	if err != nil {
		t.Fatal(err)
	}

	if err := os.Chdir(dir); err != nil {
		t.Fatal(err)
	}

	defer os.RemoveAll(dir)

	var obj, obj2 ir.Objects
	j := newTask()
	j.args.c = true
	j.args.hooks.obj = &obj
	j.args.args = []string{src}
	if err := j.main(); err != nil {
		t.Fatal(err)
	}

	objf := filepath.Join(dir, "issue4.o")
	f, err := os.Open(objf)
	if err != nil {
		t.Fatal(err)
	}

	if _, err := obj2.ReadFrom(f); err != nil {
		t.Fatal(err)
	}

	if err := f.Close(); err != nil {
		t.Fatal(err)
	}

	if g, e := ir.PrettyString(obj2), ir.PrettyString(obj); g != e {
		t.Fatalf("got\n%s\nexp\n%s", g, e)
	}

	var bin *virtual.Binary
	j = newTask()
	j.args.args = []string{objf}
	j.args.g = true
	j.args.hooks.bin = &bin
	if err := j.main(); err != nil {
		t.Fatal(err)
	}

	if _, ok := bin.Sym[ir.NameID(xc.Dict.SID("fib"))]; !ok {
		t.Fatalf("fib symbol missing: %v", bin.Sym)
	}
}

func TestLibToolConfig(t *testing.T) {
	m, err := filepath.Glob(filepath.Join("testdata", "*.la"))
	if err != nil {
		t.Fatal(err)
	}

	for _, v := range m {
		c, err := newLibToolConfigFile(v)
		if err != nil {
			t.Fatal(err)
		}

		t.Logf("%s: %#v", v, c)
	}
}
