// Copyright 2017 The 99c Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Command 99run executes binary programs produced by the 99c compiler.
//
// Usage
//
//	99run a.out [arguments]
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

	code, err := virtual.Exec(&b, os.Args[1:], os.Stdin, os.Stdout, os.Stderr, 0, 8<<20, "")
	if err != nil {
		if code == 0 {
			code = 1
		}
		exit(code, "%v\n", err)
	}

	exit(code, "")
}
