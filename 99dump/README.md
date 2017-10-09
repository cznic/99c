# Table of Contents

1. Usage
1. Installation
1. Changelog
1. Sample

# 99dump

Command 99dump lists object and executable files produced by the 99c compiler.

### Usage

    99dump [files...]

### Installation

To install or update 99dump

     $ go get [-u] github.com/cznic/99c/99dump

Online documentation: [godoc.org/github.com/cznic/99c/99dump](http://godoc.org/github.com/cznic/99c/99dump)

### Changelog

2017-10-09: Initial public release.

### Sample

    $ cd ../examples/hello/
    $ ls *
    hello.c  log
    $ 99c -c hello.c && 99dump hello.o
    ir.Objects hello.o:
    # [0]: *ir.FunctionDefinition { ExternalLinkage __builtin_fopen  func(*int8,*int8)*struct{} X__FILE_TYPE__ /home/jnml/src/github.com/cznic/ccir/libc/builtin.h:45:15} [__filename __modes]
    0x00000		panic           		; /home/jnml/src/github.com/cznic/ccir/libc/builtin.h:45:15
    # [1]: *ir.FunctionDefinition { ExternalLinkage __builtin_strlen  func(*int8)uint64  /home/jnml/src/github.com/cznic/ccir/libc/builtin.h:46:15} [__s]
    ...
    # [62]: *ir.FunctionDefinition { ExternalLinkage main  func()int32  hello.c:3:1} []
    0x00000		result          	&#0, *int32			; hello.c:3:12
    0x00001		const           	0x0, int32			; hello.c:3:12
    0x00002		store           	int32				; hello.c:3:12
    0x00003		drop            	int32				; hello.c:3:12
    0x00004		beginScope      					; hello.c:3:12
    0x00005		allocResult     	int32				;  hello.c:4:2
    0x00006		global          	&printf, *func(*int8...)int32	;  hello.c:4:2
    0x00007		arguments       					; hello.c:4:9
    0x00008		const           	"hello world\n", *int8		; hello.c:4:9
    0x00009		callfp          	1, *func(*int8...)int32		; hello.c:4:2
    0x0000a		drop            	int32				; hello.c:4:2
    0x0000b		return          					; hello.c:5:1
    0x0000c		endScope        					; hello.c:5:1
    # [63]: *ir.DataDefinition { InternalLinkage main__func__0  [5]int8  -} "main"+0
    jnml@r550:~/src/github.com/cznic/99c/examples/hello$ 99c hello.o && 99dump a.out
    virtual.Binary a.out: code 0x00021, text 0x00010, data 0x00030, bss 0x00020, pc2func 2, pc2line 10
    0x00000		call           0x2	; -
    0x00001		ffireturn      		; -
    
    # _start
    0x00002	func	; /home/jnml/src/github.com/cznic/ccir/libc/crt0.c:13:1
    0x00003		arguments      			; /home/jnml/src/github.com/cznic/ccir/libc/crt0.c:14:1
    0x00004		push64         (ds)		; /home/jnml/src/github.com/cznic/ccir/libc/crt0.c:14:1
    0x00005		push64         (ds+0x10)	; /home/jnml/src/github.com/cznic/ccir/libc/crt0.c:14:1
    0x00006		push64         (ds+0x20)	; /home/jnml/src/github.com/cznic/ccir/libc/crt0.c:14:1
    0x00007		#register_stdfiles		; /home/jnml/src/github.com/cznic/ccir/libc/crt0.c:14:1
    0x00008		arguments      			; /home/jnml/src/github.com/cznic/ccir/libc/crt0.c:15:1
    0x00009		sub            sp, 0x8		; /home/jnml/src/github.com/cznic/ccir/libc/crt0.c:15:1
    0x0000a		arguments      			; /home/jnml/src/github.com/cznic/ccir/libc/crt0.c:15:1
    0x0000b		push32         (ap-0x8)		; /home/jnml/src/github.com/cznic/ccir/libc/crt0.c:15:1
    0x0000c		push64         (ap-0x10)	; /home/jnml/src/github.com/cznic/ccir/libc/crt0.c:15:1
    0x0000d		call           0x16		; /home/jnml/src/github.com/cznic/ccir/libc/crt0.c:15:1
    0x0000e		#exit          			; /home/jnml/src/github.com/cznic/ccir/libc/crt0.c:15:1
    
    0x0000f		builtin        		; /home/jnml/src/github.com/cznic/ccir/libc/crt0.c:16:1
    0x00010		#register_stdfiles	; __register_stdfiles:89:1
    0x00011		ffireturn      		; __register_stdfiles:89:1
    
    0x00012		add            sp, 0x8	; __builtin_exit:86:1
    0x00013		#exit          		; __builtin_exit:86:1
    
    0x00014		call           0x16	; __builtin_exit:86:1
    0x00015		ffireturn      		; __builtin_exit:86:1
    
    # main
    0x00016	func	; hello.c:3:1
    0x00017		push           ap	; hello.c:3:1
    0x00018		zero32         		; hello.c:3:1
    0x00019		store32        		; hello.c:3:1
    0x0001a		arguments      		; hello.c:3:1
    0x0001b		push           ts+0x0	; hello.c:4:1
    0x0001c		#printf        		; hello.c:4:1
    0x0001d		return         		; hello.c:4:1
    
    0x0001e		builtin        	; hello.c:5:1
    0x0001f		#printf        	; printf:253:1
    0x00020		ffireturn      	; printf:253:1
    
    Text segment
    00000000  68 65 6c 6c 6f 20 77 6f  72 6c 64 0a 00 00 00 00  |hello world.....|
    
    Data segment
    00000000  30 00 00 00 00 00 00 00  00 00 00 00 00 00 00 00  |0...............|
    00000010  38 00 00 00 00 00 00 00  00 00 00 00 00 00 00 00  |8...............|
    00000020  40 00 00 00 00 00 00 00  00 00 00 00 00 00 00 00  |@...............|
    
    DS relative bitvector
    00000000  01 00 01 00 01                                    |.....|
    
    Symbol table
    0x00012	function	__builtin_exit
    0x0000f	function	__register_stdfiles
    0x00000	function	_start
    0x00014	function	main
    0x0001e	function	printf
    $ 
