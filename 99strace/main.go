// Copyright 2017 The 99c Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Command 99strace traces system calls of programs produced by the 99c compiler.
//
// The trace is written to stderr.
//
// Usage
//
// To trace a.out
//
//     99strace a.out [arguments]
//
// Installation
//
// To install or update 99strace
//
//      $ go get [-u] -tags virtual.strace github.com/cznic/99c/99strace
//
// Online documentation: [godoc.org/github.com/cznic/99c/99strace](http://godoc.org/github.com/cznic/99c/99strace)
//
// Changelog
//
// 2017-10-09: Initial public release.
//
// Sample
//
// Example session
//
//     $ cd ../examples/strace/
//     $ ls *
//     data.txt  main.c
//     $ cat main.c
//     #include <stdlib.h>
//     #include <fcntl.h>
//     #include <unistd.h>
//
//     #define BUFSIZE 1<<16
//
//     int main(int argc, char **argv) {
//     	char *buf = malloc(BUFSIZE);
//     	if (!buf) {
//     		return 1;
//     	}
//
//     	for (int i = 1; i < argc; i++) {
//     		int fd = open(argv[i], O_RDWR);
//     		if (fd < 0) {
//     			return 1;
//     		}
//
//     		ssize_t n;
//     		while ((n = read(fd, buf, BUFSIZE)) > 0) {
//     			write(0, buf, n);
//     		}
//     	}
//     	free(buf);
//     }
//     $ cat data.txt
//     Lorem ipsum.
//     $ 99c main.c && ./a.out data.txt
//     Lorem ipsum.
//     $ 99strace a.out data.txt
//     malloc(0x10000) 0x7f83a9400020
//     open("data.txt", O_RDWR, 0) 5 errno 0
//     read(5, 0x7f83a9400020, 65536) 13 errno 0
//     Lorem ipsum.
//     write(0, 0x7f83a9400020, 13) 13 errno 0
//     read(5, 0x7f83a9400020, 65536) 0 errno 0
//     freep(0x7f83a9400020)
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
	if !strace {
		exit(1, `This tool must be built using '-tags virtual.strace', please rebuild it.
Use
	$ go install -tags virtual.strace github.com/cznic/99c/99strace
or
	$ go get [-u] -tags virtual.strace github.com/cznic/99c/99strace
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

	code, err := virtual.Exec(&b, os.Args[1:], os.Stdin, os.Stdout, os.Stderr, 0, 8<<20, "")
	if err != nil {
		if code == 0 {
			code = 1
		}
		exit(code, "%v\n", err)
	}

	exit(code, "")
}
