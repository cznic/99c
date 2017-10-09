# Table of Contents

1. Usage
1. Installation
1. Changelog
1. Sample

# 99prof

Command 99prof profiles programs produced by the 99c compiler.

The profile is written to stderr.

### Usage

Profile a program by issuing

    99prof [-functions] [-lines] [-instructions] [-rate] a.out [arguments]

    -functions
      	profile functions
    -instructions
      	profile instructions
    -lines
      	profile lines
    -rate int
      	profile rate (default 1000)

### Installation

To install or update 99prof

     $ go get [-u] -tags virtual.profile github.com/cznic/99c/99prof

Online documentation: [godoc.org/github.com/cznic/99c/99prof](http://godoc.org/github.com/cznic/99c/99prof)

### Changelog

2017-10-09: Initial public release.

### Sample

    $ cd ../examples/prof/
    $ ls *
    fib.c
    $ cat fib.c 
    #include <stdlib.h>
    #include <stdio.h>
    
    int fib(int n) {
    	switch (n) {
    	case 0:
    		return 0;
    	case 1:
    		return 1;
    	default:
    		return fib(n-1)+fib(n-2);
    	}
    }
    
    int main(int argc, char **argv) {
    	if (argc != 2) {
    		return 2;
    	}
    
    	int n = atoi(argv[1]);
    	if (n<=0 || n>40) {
    		return 1;
    	}
    
    	printf("%i\n", fib(n));
    }
    $ 99c fib.c && 99prof -functions -lines -instructions a.out 31 2>log
    1346269
    $ cat log
    # [99prof -functions -lines -instructions a.out 31] 781.384628ms, 72.483 MIPS
    # functions
    fib   	     56636    100.00%    100.00%
    _start	         1      0.00%    100.00%
    # lines
    fib.c:11:	     32707     57.75%     57.75%
    fib.c:5:	      8738     15.43%     73.18%
    fib.c:4:	      4350      7.68%     80.86%
    fib.c:9:	      4002      7.07%     87.92%
    fib.c:7:	      2476      4.37%     92.29%
    fib.c:10:	      2184      3.86%     96.15%
    fib.c:8:	      1357      2.40%     98.55%
    fib.c:6:	       822      1.45%    100.00%
    /home/jnml/src/github.com/cznic/ccir/libc/crt0.c:13:	         1      0.00%    100.00%
    # instructions
    Argument32      8796     15.53%	     15.53%
    Push32          5601      9.89%	     25.42%
    AddSP           4426      7.81%	     33.23%
    Return          4413      7.79%	     41.03%
    SubI32          4375      7.72%	     48.75%
    AP              4363      7.70%	     56.45%
    Arguments       4359      7.70%	     64.15%
    Func            4351      7.68%	     71.83%
    Call            4346      7.67%	     79.51%
    SwitchI32       4331      7.65%	     87.15%
    Store32         4283      7.56%	     94.72%
    AddI32          2187      3.86%	     98.58%
    Zero32           806      1.42%	    100.00%
    $ 
