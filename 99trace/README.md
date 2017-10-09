# Table of Contents

1. Usage
1. Installation
1. Changelog
1. Sample

# 99trace

Command 99trace traces execution of binary programs produced by the 99c compiler.

The trace is written to stderr.

### Usage

    99trace a.out [arguments]

### Installation

To install or update 99trace

     $ go get [-u] -tags virtual.trace github.com/cznic/99c/99trace

Online documentation: [godoc.org/github.com/cznic/99c/99trace](http://godoc.org/github.com/cznic/99c/99trace)

### Changelog

2017-10-09: Initial public release.

### Sample

    $ cd ../examples/hello/
    $ 99c hello.c && 99trace a.out 2>log
    hello world
    $ cat log
    # _start
    0x00002	func	; /home/jnml/src/github.com/cznic/ccir/libc/crt0.c:13:1
    0x00003		arguments      		; /home/jnml/src/github.com/cznic/ccir/libc/crt0.c:14:1
    0x00004		push64         (ds)	; /home/jnml/src/github.com/cznic/ccir/libc/crt0.c:14:1
    0x00005		push64         (ds+0x10); /home/jnml/src/github.com/cznic/ccir/libc/crt0.c:14:1
    0x00006		push64         (ds+0x20); /home/jnml/src/github.com/cznic/ccir/libc/crt0.c:14:1
    0x00007		#register_stdfiles	; /home/jnml/src/github.com/cznic/ccir/libc/crt0.c:14:1
    0x00008		arguments      		; /home/jnml/src/github.com/cznic/ccir/libc/crt0.c:15:1
    0x00009		sub            sp, 0x8	; /home/jnml/src/github.com/cznic/ccir/libc/crt0.c:15:1
    0x0000a		arguments      		; /home/jnml/src/github.com/cznic/ccir/libc/crt0.c:15:1
    0x0000b		push32         (ap-0x8)	; /home/jnml/src/github.com/cznic/ccir/libc/crt0.c:15:1
    0x0000c		push64         (ap-0x10); /home/jnml/src/github.com/cznic/ccir/libc/crt0.c:15:1
    0x0000d		call           0x16	; /home/jnml/src/github.com/cznic/ccir/libc/crt0.c:15:1
    # main
    0x00016	func	; hello.c:3:1
    0x00017		push           ap	; hello.c:3:1
    0x00018		zero32         		; hello.c:3:1
    0x00019		store32        		; hello.c:3:1
    0x0001a		arguments      		; hello.c:3:1
    0x0001b		push           ts+0x0	; hello.c:4:1
    0x0001c		#printf        		; hello.c:4:1
    0x0001d		return         		; hello.c:4:1
    
    0x0000e	#exit          	; /home/jnml/src/github.com/cznic/ccir/libc/crt0.c:15:1
    
    $ 
