// Copyright 2017 The 99c Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Command 99prof profiles programs produced by the 99c compiler.
//
// The profile is written to stderr.
//
// Usage
//
// Profile a program by issuing
//
//     99prof [-functions] [-lines] [-instructions] [-rate] a.out [arguments]
//
//     -functions
//       	profile functions
//     -instructions
//       	profile instructions
//     -lines
//       	profile lines
//     -rate int
//       	profile rate (default 1000)
//
// Installation
//
// To install or update 99prof
//
//      $ go get [-u] -tags virtual.profile github.com/cznic/99c/99prof
//
// Online documentation: [godoc.org/github.com/cznic/99c/99prof](http://godoc.org/github.com/cznic/99c/99prof)
//
// Changelog
//
// 2017-10-09: Initial public release.
//
// Sample
//
// Profile a fibonacci program.
//
//     $ cd ../examples/prof/
//     $ ls *
//     fib.c
//     $ cat fib.c
//     #include <stdlib.h>
//     #include <stdio.h>
//
//     int fib(int n) {
//     	switch (n) {
//     	case 0:
//     		return 0;
//     	case 1:
//     		return 1;
//     	default:
//     		return fib(n-1)+fib(n-2);
//     	}
//     }
//
//     int main(int argc, char **argv) {
//     	if (argc != 2) {
//     		return 2;
//     	}
//
//     	int n = atoi(argv[1]);
//     	if (n<=0 || n>40) {
//     		return 1;
//     	}
//
//     	printf("%i\n", fib(n));
//     }
//     $ 99c fib.c && 99prof -functions -lines -instructions a.out 31 2>log
//     1346269
//     $ cat log
//     # [99prof -functions -lines -instructions a.out 31] 781.384628ms, 72.483 MIPS
//     # functions
//     fib   	     56636    100.00%    100.00%
//     _start	         1      0.00%    100.00%
//     # lines
//     fib.c:11:	     32707     57.75%     57.75%
//     fib.c:5:	      8738     15.43%     73.18%
//     fib.c:4:	      4350      7.68%     80.86%
//     fib.c:9:	      4002      7.07%     87.92%
//     fib.c:7:	      2476      4.37%     92.29%
//     fib.c:10:	      2184      3.86%     96.15%
//     fib.c:8:	      1357      2.40%     98.55%
//     fib.c:6:	       822      1.45%    100.00%
//     /home/jnml/src/github.com/cznic/ccir/libc/crt0.c:13:	         1      0.00%    100.00%
//     # instructions
//     Argument32      8796     15.53%	     15.53%
//     Push32          5601      9.89%	     25.42%
//     AddSP           4426      7.81%	     33.23%
//     Return          4413      7.79%	     41.03%
//     SubI32          4375      7.72%	     48.75%
//     AP              4363      7.70%	     56.45%
//     Arguments       4359      7.70%	     64.15%
//     Func            4351      7.68%	     71.83%
//     Call            4346      7.67%	     79.51%
//     SwitchI32       4331      7.65%	     87.15%
//     Store32         4283      7.56%	     94.72%
//     AddI32          2187      3.86%	     98.58%
//     Zero32           806      1.42%	    100.00%
//     $
package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"sort"
	"time"

	"github.com/cznic/virtual"
)

func exit(code int, msg string, arg ...interface{}) {
	if msg != "" {
		fmt.Fprintf(os.Stderr, os.Args[0]+": "+msg, arg...)
	}
	os.Exit(code)
}

func main() {
	if !profile {
		exit(1, `This tool must be built using '-tags virtual.profile', please rebuild it.
Use
	$ go install -tags virtual.profile github.com/cznic/99c/99prof
or
	$ go get [-u] -tags virtual.profile github.com/cznic/99c/99prof
`)
	}

	functions := flag.Bool("functions", false, "profile functions")
	instructions := flag.Bool("instructions", false, "profile instructions")
	lines := flag.Bool("lines", false, "profile lines")
	rate := flag.Int("rate", 1000, "profile rate")
	flag.Parse()

	if flag.NArg() == 0 {
		exit(2, "missing program name %v\n", os.Args)
	}

	nm := flag.Arg(0)
	bin, err := os.Open(nm)
	if err != nil {
		exit(1, "%v\n", err)
	}

	var b virtual.Binary
	if _, err := b.ReadFrom(bufio.NewReader(bin)); err != nil {
		exit(1, "%v\n", err)
	}

	args := os.Args[1:]
	for i, v := range args {
		if v == nm {
			args = args[i:]
			break
		}
	}

	var opts []virtual.Option
	if *functions {
		opts = append(opts, virtual.ProfileFunctions())
	}
	if *lines {
		opts = append(opts, virtual.ProfileLines())
	}
	if *instructions {
		opts = append(opts, virtual.ProfileInstructions())
	}
	if n := *rate; n != 0 {
		opts = append(opts, virtual.ProfileRate(n))
	}

	t0 := time.Now()
	vm, code, err := virtual.New(&b, args, os.Stdin, os.Stdout, os.Stderr, 0, 8<<20, "", opts...)
	d := time.Since(t0)
	if err != nil {
		if code == 0 {
			code = 1
		}
		exit(code, "%v\n", err)
	}

	var s int64
	switch {
	case len(vm.ProfileFunctions) != 0:
		for _, v := range vm.ProfileFunctions {
			s += int64(v)
		}
	case len(vm.ProfileLines) != 0:
		for _, v := range vm.ProfileLines {
			s += int64(v)
		}
	case len(vm.ProfileInstructions) != 0:
		for _, v := range vm.ProfileInstructions {
			s += int64(v)
		}
	}
	if s != 0 {
		w := bufio.NewWriter(os.Stderr)
		fmt.Fprintf(w, "# %v %v, %.3f MIPS\n", os.Args, d, float64(s)/1e6*float64(*rate)*float64(time.Second)/float64(d))
		out(w, vm)
		w.Flush()
		os.Stderr.Sync()
	}
	exit(code, "")
}

func out(w io.Writer, vm *virtual.Machine) {
	rate := vm.ProfileRate
	if rate == 0 {
		rate = 1
	}
	if len(vm.ProfileFunctions) != 0 {
		type u struct {
			virtual.PCInfo
			n int
		}
		var s int64
		var a []u
		var wi int
		for k, v := range vm.ProfileFunctions {
			a = append(a, u{k, v})
			s += int64(v)
			if n := len(k.Name.String()); n > wi {
				wi = n
			}
		}
		sort.Slice(a, func(i, j int) bool {
			if a[i].n < a[j].n {
				return true
			}

			if a[i].n > a[j].n {
				return false
			}

			return a[i].Name < a[j].Name
		})
		fmt.Fprintln(w, "# functions")
		var c int64
		for i := len(a) - 1; i >= 0; i-- {
			c += int64(a[i].n)
			fmt.Fprintf(
				w,
				"%*v\t%10v%10.2f%%%10.2f%%\n",
				-wi, a[i].Name, a[i].n,
				100*float64(a[i].n)/float64(s),
				100*float64(c)/float64(s),
			)
		}
	}
	if len(vm.ProfileLines) != 0 {
		type u struct {
			virtual.PCInfo
			n int
		}
		var s int64
		var a []u
		for k, v := range vm.ProfileLines {
			a = append(a, u{k, v})
			s += int64(v)
		}
		sort.Slice(a, func(i, j int) bool {
			if a[i].n < a[j].n {
				return true
			}

			if a[i].n > a[j].n {
				return false
			}

			if a[i].Name < a[j].Name {
				return true
			}

			if a[i].Name > a[j].Name {
				return false
			}

			return a[i].Line < a[j].Line
		})
		fmt.Fprintln(w, "# lines")
		var c int64
		for i := len(a) - 1; i >= 0; i-- {
			c += int64(a[i].n)
			fmt.Fprintf(
				w,
				"%v:%v:\t%10v%10.2f%%%10.2f%%\n",
				a[i].Name, a[i].Line, a[i].n,
				100*float64(a[i].n)/float64(s),
				100*float64(c)/float64(s),
			)
		}
	}
	if len(vm.ProfileInstructions) != 0 {
		type u struct {
			virtual.Opcode
			n int
		}
		var s int64
		var a []u
		var wi int
		for k, v := range vm.ProfileInstructions {
			a = append(a, u{k, v})
			s += int64(v)
			if n := len(k.String()); n > wi {
				wi = n
			}
		}
		sort.Slice(a, func(i, j int) bool {
			if a[i].n < a[j].n {
				return true
			}

			if a[i].n > a[j].n {
				return false
			}

			return a[i].Opcode < a[j].Opcode
		})
		fmt.Fprintln(w, "# instructions")
		var c int64
		for i := len(a) - 1; i >= 0; i-- {
			c += int64(a[i].n)
			fmt.Fprintf(
				w,
				"%*s%10v%10.2f%%\t%10.2f%%\n",
				-wi, a[i].Opcode, a[i].n,
				100*float64(a[i].n)/float64(s),
				100*float64(c)/float64(s),
			)
		}
	}
}
