// Copyright 2017 The 99c Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Command 99dump lists object and executable files produced by the 99c compiler.
//
// Usage
//
// To dump object or binary files produced by the 99c compiler
//
//     99dump [files...]
//
// Installation
//
// To install or update 99dump
//
//      $ go get [-u] github.com/cznic/99c/99dump
//
// Online documentation: [godoc.org/github.com/cznic/99c/99dump](http://godoc.org/github.com/cznic/99c/99dump)
//
// Changelog
//
// 2017-10-09: Initial public release.
//
// Sample
//
// Dump hello.c
//
//     $ cd ../examples/hello/
//     $ ls *
//     hello.c  log
//     $ 99c -c hello.c && 99dump hello.o
//     ir.Objects hello.o:
//     # [0]: *ir.FunctionDefinition { ExternalLinkage __builtin_fopen  func(*int8,*int8)*struct{} X__FILE_TYPE__ /home/jnml/src/github.com/cznic/ccir/libc/builtin.h:45:15} [__filename __modes]
//     0x00000		panic           		; /home/jnml/src/github.com/cznic/ccir/libc/builtin.h:45:15
//     # [1]: *ir.FunctionDefinition { ExternalLinkage __builtin_strlen  func(*int8)uint64  /home/jnml/src/github.com/cznic/ccir/libc/builtin.h:46:15} [__s]
//     ...
//     # [62]: *ir.FunctionDefinition { ExternalLinkage main  func()int32  hello.c:3:1} []
//     0x00000		result          	&#0, *int32			; hello.c:3:12
//     0x00001		const           	0x0, int32			; hello.c:3:12
//     0x00002		store           	int32				; hello.c:3:12
//     0x00003		drop            	int32				; hello.c:3:12
//     0x00004		beginScope      					; hello.c:3:12
//     0x00005		allocResult     	int32				;  hello.c:4:2
//     0x00006		global          	&printf, *func(*int8...)int32	;  hello.c:4:2
//     0x00007		arguments       					; hello.c:4:9
//     0x00008		const           	"hello world\n", *int8		; hello.c:4:9
//     0x00009		callfp          	1, *func(*int8...)int32		; hello.c:4:2
//     0x0000a		drop            	int32				; hello.c:4:2
//     0x0000b		return          					; hello.c:5:1
//     0x0000c		endScope        					; hello.c:5:1
//     # [63]: *ir.DataDefinition { InternalLinkage main__func__0  [5]int8  -} "main"+0
//     jnml@r550:~/src/github.com/cznic/99c/examples/hello$ 99c hello.o && 99dump a.out
//     virtual.Binary a.out: code 0x00021, text 0x00010, data 0x00030, bss 0x00020, pc2func 2, pc2line 10
//     0x00000		call           0x2	; -
//     0x00001		ffireturn      		; -
//
//     # _start
//     0x00002	func	; /home/jnml/src/github.com/cznic/ccir/libc/crt0.c:13:1
//     0x00003		arguments      			; /home/jnml/src/github.com/cznic/ccir/libc/crt0.c:14:1
//     0x00004		push64         (ds)		; /home/jnml/src/github.com/cznic/ccir/libc/crt0.c:14:1
//     0x00005		push64         (ds+0x10)	; /home/jnml/src/github.com/cznic/ccir/libc/crt0.c:14:1
//     0x00006		push64         (ds+0x20)	; /home/jnml/src/github.com/cznic/ccir/libc/crt0.c:14:1
//     0x00007		#register_stdfiles		; /home/jnml/src/github.com/cznic/ccir/libc/crt0.c:14:1
//     0x00008		arguments      			; /home/jnml/src/github.com/cznic/ccir/libc/crt0.c:15:1
//     0x00009		sub            sp, 0x8		; /home/jnml/src/github.com/cznic/ccir/libc/crt0.c:15:1
//     0x0000a		arguments      			; /home/jnml/src/github.com/cznic/ccir/libc/crt0.c:15:1
//     0x0000b		push32         (ap-0x8)		; /home/jnml/src/github.com/cznic/ccir/libc/crt0.c:15:1
//     0x0000c		push64         (ap-0x10)	; /home/jnml/src/github.com/cznic/ccir/libc/crt0.c:15:1
//     0x0000d		call           0x16		; /home/jnml/src/github.com/cznic/ccir/libc/crt0.c:15:1
//     0x0000e		#exit          			; /home/jnml/src/github.com/cznic/ccir/libc/crt0.c:15:1
//
//     0x0000f		builtin        		; /home/jnml/src/github.com/cznic/ccir/libc/crt0.c:16:1
//     0x00010		#register_stdfiles	; __register_stdfiles:89:1
//     0x00011		ffireturn      		; __register_stdfiles:89:1
//
//     0x00012		add            sp, 0x8	; __builtin_exit:86:1
//     0x00013		#exit          		; __builtin_exit:86:1
//
//     0x00014		call           0x16	; __builtin_exit:86:1
//     0x00015		ffireturn      		; __builtin_exit:86:1
//
//     # main
//     0x00016	func	; hello.c:3:1
//     0x00017		push           ap	; hello.c:3:1
//     0x00018		zero32         		; hello.c:3:1
//     0x00019		store32        		; hello.c:3:1
//     0x0001a		arguments      		; hello.c:3:1
//     0x0001b		push           ts+0x0	; hello.c:4:1
//     0x0001c		#printf        		; hello.c:4:1
//     0x0001d		return         		; hello.c:4:1
//
//     0x0001e		builtin        	; hello.c:5:1
//     0x0001f		#printf        	; printf:253:1
//     0x00020		ffireturn      	; printf:253:1
//
//     Text segment
//     00000000  68 65 6c 6c 6f 20 77 6f  72 6c 64 0a 00 00 00 00  |hello world.....|
//
//     Data segment
//     00000000  30 00 00 00 00 00 00 00  00 00 00 00 00 00 00 00  |0...............|
//     00000010  38 00 00 00 00 00 00 00  00 00 00 00 00 00 00 00  |8...............|
//     00000020  40 00 00 00 00 00 00 00  00 00 00 00 00 00 00 00  |@...............|
//
//     DS relative bitvector
//     00000000  01 00 01 00 01                                    |.....|
//
//     Symbol table
//     0x00012	function	__builtin_exit
//     0x0000f	function	__register_stdfiles
//     0x00000	function	_start
//     0x00014	function	main
//     0x0001e	function	printf
//     $
package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"text/tabwriter"

	"github.com/cznic/ir"
	"github.com/cznic/virtual"
	"github.com/cznic/xc"
)

func exit(code int, msg string, arg ...interface{}) {
	if msg != "" {
		fmt.Fprintf(os.Stderr, os.Args[0]+": "+msg, arg...)
	}
	os.Exit(code)
}

const (
	bin = true
	obj = false
)

func use(...interface{}) {}

func main() {
	w := bufio.NewWriter(os.Stdout)

	defer w.Flush()

	for _, arg := range os.Args[1:] {
		switch {
		case filepath.Ext(arg) == ".o":
			use(try(w, arg, obj) || try(w, arg, bin) || unknown(arg))
		default:
			use(try(w, arg, bin) || try(w, arg, obj) || unknown(arg))
		}
	}
}

func unknown(fn string) bool {
	exit(1, "unrecognized file format: %s\n", fn)
	panic("unreachable")
}

func try(w io.Writer, fn string, bin bool) bool {
	tw := new(tabwriter.Writer)
	tw.Init(w, 0, 8, 1, '\t', 0)

	defer tw.Flush()

	w = tw
	f, err := os.Open(fn)
	if err != nil {
		exit(1, "%v\n", err)
	}

	r := bufio.NewReader(f)
	switch {
	case bin:
		var b virtual.Binary
		if _, err := b.ReadFrom(r); err != nil {
			return false
		}

		fmt.Fprintf(w, "%T %s: code %#05x, text %#05x, data %#05x, bss %#05x, pc2func %v, pc2line %v\n",
			b, fn, len(b.Code), len(b.Text), len(b.Data), b.BSS, len(b.Functions), len(b.Lines),
		)
		virtual.DumpCode(w, b.Code, 0, b.Functions, b.Lines)
		if len(b.Text) != 0 {
			fmt.Fprintf(w, "Text segment\n%s\n", hex.Dump(b.Text))
		}
		if len(b.Data) != 0 {
			fmt.Fprintf(w, "Data segment\n%s\n", hex.Dump(b.Data))
		}
		if len(b.TSRelative) != 0 {
			fmt.Fprintf(w, "TS relative bitvector\n%s\n", hex.Dump(b.TSRelative))
		}
		if len(b.DSRelative) != 0 {
			fmt.Fprintf(w, "DS relative bitvector\n%s\n", hex.Dump(b.DSRelative))
		}
		var a []string
		for k := range b.Sym {
			a = append(a, string(xc.Dict.S(int(k))))
		}
		sort.Strings(a)
		fmt.Fprintln(w, "Symbol table")
		for _, k := range a {
			fmt.Fprintf(w, "%#05x\tfunction\t%s\n", b.Sym[ir.NameID(xc.Dict.SID(k))], k)
		}
	default:
		var o ir.Objects
		if _, err := o.ReadFrom(r); err != nil {
			return false
		}

		fmt.Fprintf(w, "%T %s:\n", o, fn)
		for i, v := range o {
			for j, v := range v {
				switch x := v.(type) {
				case *ir.DataDefinition:
					fmt.Fprintf(w, "# [%v, %v]: %T %v %v\n", i, j, x, x.ObjectBase, x.Value)
				case *ir.FunctionDefinition:
					fmt.Fprintf(w, "# [%v, %v]: %T %v %v\n", i, j, x, x.ObjectBase, x.Arguments)
					for i, v := range x.Body {
						fmt.Fprintf(w, "%#05x\t%v\n", i, v)
					}
				}
			}
		}

	}
	return true
}
