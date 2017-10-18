// Copyright 2017 The 99c Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Command 99trace traces execution of binary programs produced by the 99c
// compiler.
//
// The trace is output to stderr.
//
// Usage
//
// To trace a compiled program named a.out
//
//     99trace a.out [arguments]
//
// Installation
//
// To install or update 99trace
//
//      $ go get [-u] -tags virtual.trace github.com/cznic/99c/99trace
//
// Online documentation: [godoc.org/github.com/cznic/99c/99trace](http://godoc.org/github.com/cznic/99c/99trace)
//
// Changelog
//
// 2017-10-09: Initial public release.
//
// Sample
//
// Example session
//
//     $ cd ../examples/hello/
//     $ 99c hello.c && 99trace a.out 2>log
//     hello world
//     $ cat log
//     # _start
//     0x00002	func	; /home/jnml/src/github.com/cznic/ccir/libc/crt0.c:13:1
//     0x00003		arguments      		; /home/jnml/src/github.com/cznic/ccir/libc/crt0.c:14:1
//     0x00004		push64         (ds)	; /home/jnml/src/github.com/cznic/ccir/libc/crt0.c:14:1
//     0x00005		push64         (ds+0x10); /home/jnml/src/github.com/cznic/ccir/libc/crt0.c:14:1
//     0x00006		push64         (ds+0x20); /home/jnml/src/github.com/cznic/ccir/libc/crt0.c:14:1
//     0x00007		#register_stdfiles	; /home/jnml/src/github.com/cznic/ccir/libc/crt0.c:14:1
//     0x00008		arguments      		; /home/jnml/src/github.com/cznic/ccir/libc/crt0.c:15:1
//     0x00009		sub            sp, 0x8	; /home/jnml/src/github.com/cznic/ccir/libc/crt0.c:15:1
//     0x0000a		arguments      		; /home/jnml/src/github.com/cznic/ccir/libc/crt0.c:15:1
//     0x0000b		push32         (ap-0x8)	; /home/jnml/src/github.com/cznic/ccir/libc/crt0.c:15:1
//     0x0000c		push64         (ap-0x10); /home/jnml/src/github.com/cznic/ccir/libc/crt0.c:15:1
//     0x0000d		call           0x16	; /home/jnml/src/github.com/cznic/ccir/libc/crt0.c:15:1
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
//     0x0000e	#exit          	; /home/jnml/src/github.com/cznic/ccir/libc/crt0.c:15:1
//
//     $
package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/cznic/virtual"
)

func exit(code int, msg string, arg ...interface{}) {
	if msg != "" {
		fmt.Fprintf(os.Stderr, os.Args[0]+": "+msg, arg...)
	}
	os.Exit(code)
}

func main() {
	if !trace {
		exit(1, `This tool must be built using '-tags virtual.trace', please rebuild it.
Use
	$ go install -tags virtual.trace github.com/cznic/99c/99trace
or
	$ go get [-u] -tags virtual.trace github.com/cznic/99c/99trace
`)
	}

	if len(os.Args) < 2 {
		exit(2, "invalid arguments %v\n", os.Args)
	}

	bin, err := os.Open(os.Args[1])
	if err != nil {
		exit(1, "%v\n", err)
	}

	var b virtual.Binary
	if _, err := b.ReadFrom(bufio.NewReader(bin)); err != nil {
		exit(1, "%v\n", err)
	}

	code, err := virtual.Exec(&b, os.Args[1:], os.Stdin, os.Stdout, os.Stderr, 0, 8<<20, "", virtual.AttachProcessSignals())
	if err != nil {
		if code == 0 {
			code = 1
		}
		exit(code, "%v\n", err)
	}

	exit(code, "")
}
