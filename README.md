# Table of Contents

1. [99c](#99c)
     1. Usage
     1. Installation
     1. Changelog
     1. Supported platforms and operating systems
     1. Project status
     1. Executing compiled programs
     1. Compiling a simple program
     1. Setting the output file name
     1. Obtaining the preprocessor output
     1. Multiple C files projects
     1. Using object files
     1. Stack traces
     1. Argument passing
     1. Executing a C program embedded in a Go program
     1. Calling into an embedded C library from Go
     1. Loading C plugins at run-time
     1. Inserting defines
     1. Specifying include paths
1. [99run](#99run)
     1. Usage
     1. Installation
     1. Changelog
1. [99trace](#99trace)
     1. Usage
     1. Installation
     1. Changelog
     1. Sample
1. [99strace](#99strace)
     1. Usage
     1. Installation
     1. Changelog
     1. Sample
1. [99dump](#99dump)
     1. Usage
     1. Installation
     1. Changelog
     1. Sample
1. [99prof](#99prof)
     1. Usage
     1. Installation
     1. Changelog
     1. Sample
     1. Bogomips

# 99c

Command 99c is a c99 compiler targeting a virtual machine.

### Usage

Output of 99c -h

    99c: Flags:
      -99lib
        	Library link mode.
      -Dname
    	Equivalent to inserting '#define name 1' at the start of the
    	translation unit.
      -Dname=definition
    	Equivalent to inserting '#define name definition' at the start of the
    	translation unit.
      -E	Copy C-language source files to standard output, executing all
      	preprocessor directives; no compilation shall be performed. If any
      	operand is not a text file, the effects are unspecified.
      -Ipath
    	Add path to the include files search paths.
      -Olevel
    	Optimization setting, ignored.
      -Wwarn
    	Warning level, ignored.
      -c	Suppress the link-edit phase of the compilation, and do not
      	remove any object files that are produced.
      -g	Produce debugging information, ignored.
      -o pathname
        	Use the specified pathname, instead of the default a.out, for
        	the executable file produced. If the -o option is present with
        	-c or -E, the result is unspecified.
      -99extra flag
         Extra cc flags:
            AlignOf
            AlternateKeywords
            AnonymousStructFields
            Asm
            BuiltinClassifyType
            BuiltinConstantP
            ComputedGotos
            DefineOmitCommaBeforeDDD
            DlrInIdentifiers
            EmptyDeclarations
            EmptyDefine
            EmptyStructs
            ImaginarySuffix
            ImplicitFuncDef
            ImplicitIntType
            IncludeNext
            LegacyDesignators
            NonConstStaticInitExpressions
            Noreturn
            OmitConditionalOperand
            OmitFuncArgTypes
            OmitFuncRetType
            ParenthesizedCompoundStatemen
            StaticAssert
            TypeOf
            UndefExtraTokens
            UnsignedEnums
            WideBitFieldTypes
            WideEnumValues

Rest of the input is a list of file names, either C (.c) files or object (.o) files.

### Installation

To install or update the compiler and the virtual machine

    $ go get [-u] github.com/cznic/99c github.com/cznic/99c/99run

To update the toolchain and rebuild all commands

    $ go generate

Online documentation: [godoc.org/github.com/cznic/99c](http://godoc.org/github.com/cznic/99c)

### Changelog

2017-10-07: Initial public release.

### Supported platforms and operating systems

See: https://godoc.org/github.com/cznic/ccir#hdr-Supported_platforms_and_architectures

At the time of this writing, in GOOS_GOARCH form

    linux_386
    linux_amd64
    windows_386
    windows_amd64

Porting to other platforms/architectures is considered not difficult.

### Project status

Both the compiler and the C runtime library implementation are known to be
incomplete.  Missing pieces are added as needed. Please fill an issue when
you run into problems and be patient. Only limited resources can be
allocated to this project and to the related parts of the tool chain.

Also, contributions are welcome.

### Executing compiled programs

Running a binary on Linux

    $ ./a.out
    hello world
    $

Running a binary on Windows

    C:\> 99run a.out
    hello world
    C:\>

### Compiling a simple program

All in just a single C file.

    $ cd examples/hello/
    $ cat hello.c
    #include <stdio.h>
    
    int main() {
    	printf("hello world\n");
    }
    $ 99c hello.c && ./a.out
    hello world
    $

### Setting the output file name

If the output is a single file, use -o to set its name. (POSIX option)

    $ cd examples/hello/
    $ cat hello.c
    #include <stdio.h>
    
    int main() {
    	printf("hello world\n");
    }
    $ 99c -o hello hello.c && ./hello
    hello world
    $

### Obtaining the preprocessor output

Use -E to produce the cpp results. (POSIX option)

    $ cd examples/hello/
    /home/jnml/src/github.com/cznic/99c/examples/hello
    $ 99c -E hello.c
    # 24 "/home/jnml/src/github.com/cznic/ccir/libc/predefined.h"
    typedef char * __builtin_va_list ; 
    # 41 "/home/jnml/src/github.com/cznic/ccir/libc/builtin.h"
    typedef __builtin_va_list __gnuc_va_list ; 
    typedef void * __FILE_TYPE__ ; 
    typedef void * __jmp_buf [ 7 ] ; 
    __FILE_TYPE__ __builtin_fopen ( char * __filename , char * __modes ) ; 
    long unsigned int __builtin_strlen ( char * __s ) ; 
    long unsigned int __builtin_bswap64 ( long unsigned int x ) ; 
    char * __builtin_strchr ( char * __s , int __c ) ; 
    char * __builtin_strcpy ( char * __dest , char * __src ) ; 
    double __builtin_copysign ( double x , double y ) ; 
    int __builtin_abs ( int j ) ; 
    ...
    extern void flockfile ( FILE * __stream ) ; 
    extern int ftrylockfile ( FILE * __stream ) ; 
    extern void funlockfile ( FILE * __stream ) ; 
    # 3 "hello.c"
    int main ( ) { 
    printf ( "hello world\n" ) ; 
    } 
    $

### Multiple C files projects

A translation unit may consist of multiple source files.

    $ cd examples/multifile/
    /home/jnml/src/github.com/cznic/99c/examples/multifile
    $ cat main.c
    char *hello();
    
    #include <stdio.h>
    int main() {
    	printf("%s\n", hello());
    }
    $ cat hello.c
    char *hello() {
    	return "hello world";
    }
    $ 99c main.c hello.c && ./a.out
    hello world
    $

### Using object files

Use -c to output object files. (POSIX option)

    $ cd examples/multifile/
    /home/jnml/src/github.com/cznic/99c/examples/multifile
    $ 99c -c hello.c main.c
    $ 99c hello.o main.o && ./a.out
    hello world
    $ 99c hello.o main.c && ./a.out
    hello world
    $ 99c hello.c main.o && ./a.out
    hello world
    $

### Stack traces

If the program source(s) are available at the same location(s) as when the
program was compiled, then any stack trace produced is annotated using the
source code lines.

    $ cd examples/stack/
    /home/jnml/src/github.com/cznic/99c/examples/stack
    $ cat stack.c
    void f(int n) {
    	if (n) {
    		f(n-1);
    		return;
    	}
    
    	*(char *)n;
    }
    
    int main() {
    	f(4);
    }
    $ 99c stack.c && ./a.out
    panic: runtime error: invalid memory address or nil pointer dereference [recovered]
    	panic: runtime error: invalid memory address or nil pointer dereference
    stack.c.f(0x0)
    	stack.c:7:1	0x0002c		load8          0x0	; -	// *(char *)n;
    stack.c.f(0x1)
    	stack.c:3:1	0x00028		call           0x21	; -	// f(n-1);
    stack.c.f(0x2)
    	stack.c:3:1	0x00028		call           0x21	; -	// f(n-1);
    stack.c.f(0x3)
    	stack.c:3:1	0x00028		call           0x21	; -	// f(n-1);
    stack.c.f(0x7f3400000004)
    	stack.c:3:1	0x00028		call           0x21	; -	// f(n-1);
    stack.c.main(0x7f3400000001, 0x7f34c9400030)
    	stack.c:11:1	0x0001d		call           0x21	; -	// f(4);
    /home/jnml/src/github.com/cznic/ccir/libc/crt0.c._start(0x1, 0x7f34c9400030)
    	/home/jnml/src/github.com/cznic/ccir/libc/crt0.c:15:1	0x0000d		call           0x16	; -	// __builtin_exit(((int (*)())main) (argc, argv));
    
    [signal SIGSEGV: segmentation violation code=0x1 addr=0x0 pc=0x511e46]
    
    goroutine 1 [running]:
    github.com/cznic/virtual.(*cpu).run.func1(0xc42009e2a0)
    	/home/jnml/src/github.com/cznic/virtual/cpu.go:222 +0x26e
    panic(0x555340, 0x66a270)
    	/home/jnml/go/src/runtime/panic.go:491 +0x283
    github.com/cznic/virtual.readI8(...)
    	/home/jnml/src/github.com/cznic/virtual/cpu.go:74
    github.com/cznic/virtual.(*cpu).run(0xc42009e2a0, 0x2, 0x0, 0x0, 0x0)
    	/home/jnml/src/github.com/cznic/virtual/cpu.go:993 +0x2116
    github.com/cznic/virtual.New(0xc4201d2000, 0xc42000a090, 0x1, 0x1, 0x656540, 0xc42000c010, 0x656580, 0xc42000c018, 0x656580, 0xc42000c020, ...)
    	/home/jnml/src/github.com/cznic/virtual/virtual.go:73 +0x2be
    github.com/cznic/virtual.Exec(0xc4201d2000, 0xc42000a090, 0x1, 0x1, 0x656540, 0xc42000c010, 0x656580, 0xc42000c018, 0x656580, 0xc42000c020, ...)
    	/home/jnml/src/github.com/cznic/virtual/virtual.go:84 +0xe9
    main.main()
    	/home/jnml/src/github.com/cznic/99c/99run/main.go:37 +0x382
    $

### Argument passing

Command line arguments are passed the standard way.

    $ cd examples/args/
    /home/jnml/src/github.com/cznic/99c/examples/args
    $ cat args.c
    #include <stdio.h>
    
    int main(int argc, char **argv) {
    	for (int i = 0; i < argc; i++) {
    		printf("%i: %s\n", i, argv[i]);
    	}
    }
    $ 99c args.c && ./a.out foo bar -x -y - qux
    0: ./a.out
    1: foo
    2: bar
    3: -x
    4: -y
    5: -
    6: qux
    $

### Executing a C program embedded in a Go program

This example requires installation of additional tools

    $ go get -u github.com/cznic/assets github.com/cznic/httpfs
    $ cd examples/embedding/
    /home/jnml/src/github.com/cznic/99c/examples/embedding
    $ ls *
    main.c  main.go
    
    assets:
    keepdir
    $ cat main.c
    // +build ignore
    
    #include <stdio.h>
    
    int main() {
    	int c;
    	while ((c = getc(stdin)) != EOF) {
    		printf("%c", c >= 'a' && c <= 'z' ? c^' ' : c);
    	}
    }
    $ cat main.go
    //go:generate 99c -o assets/a.out main.c
    //go:generate assets
    
    package main
    
    import (
    	"bytes"
    	"fmt"
    	"strings"
    	"time"
    
    	"github.com/cznic/httpfs"
    	"github.com/cznic/virtual"
    )
    
    func main() {
    	fs := httpfs.NewFileSystem(assets, time.Now())
    	f, err := fs.Open("/a.out")
    	if err != nil {
    		panic(err)
    	}
    
    	var bin virtual.Binary
    	if _, err := bin.ReadFrom(f); err != nil {
    		panic(err)
    	}
    
    	var out bytes.Buffer
    	exitCode, err := virtual.Exec(&bin, nil, strings.NewReader("Foo Bar"), &out, &out, 0, 1<<20, "")
    	if err != nil {
    		panic(err)
    	}
    
    	fmt.Printf("%s\n%v\n", out.Bytes(), exitCode)
    }
    $ go generate && go build && ./embedding
    FOO BAR
    0
    $

### Calling into an embedded C library from Go

It's possible to call individual C functions from Go.

This example requires installation of additional tools

    $ go get -u github.com/cznic/assets github.com/cznic/httpfs
    $ cd examples/ffi/
    /home/jnml/src/github.com/cznic/99c/examples/ffi
    $ ls *
    lib42.c  main.go
    
    assets:
    keepdir
    $ cat lib42.c
    // +build ignore
    
    static int answer;
    
    int main() {
    	// Any library initialization comes here.
    	answer = 42;
    }
    
    // Use the -99lib option to prevent the linker from eliminating this function.
    int f42(int arg) {
    	return arg*answer;
    }
    $ cat main.go
    //go:generate 99c -99lib -o assets/a.out lib42.c
    //go:generate assets
    
    package main
    
    import (
    	"fmt"
    	"time"
    
    	"github.com/cznic/httpfs"
    	"github.com/cznic/ir"
    	"github.com/cznic/virtual"
    	"github.com/cznic/xc"
    )
    
    func main() {
    	fs := httpfs.NewFileSystem(assets, time.Now())
    	f, err := fs.Open("/a.out")
    	if err != nil {
    		panic(err)
    	}
    
    	var bin virtual.Binary
    	if _, err := bin.ReadFrom(f); err != nil {
    		panic(err)
    	}
    
    	m, _, err := virtual.New(&bin, nil, nil, nil, nil, 0, 1<<10, "")
    	if err != nil {
    		panic(err)
    	}
    
    	defer m.Close()
    
    	pc, ok := bin.Sym[ir.NameID(xc.Dict.SID("f42"))]
    	if !ok {
    		panic("symbol not found")
    	}
    
    	t, err := m.NewThread(1 << 10)
    	if err != nil {
    		panic(err)
    	}
    
    	for _, v := range []int{-1, 0, 1} {
    		var y int32
    		_, err := t.FFI1(pc, virtual.Int32Result{&y}, virtual.Int32(int32(v)))
    		if err != nil {
    			panic(err)
    		}
    
    		fmt.Println(y)
    	}
    }
    $ go generate && go build && ./ffi
    -42
    0
    42
    $

### Loading C plugins at run-time

It's possible to load C plugins at run-time.

    $ cd examples/plugin/
    /home/jnml/src/github.com/cznic/99c/examples/plugin
    $ ls *
    lib42.c  main.go
    $ cat lib42.c
    // +build ignore
    
    static int answer;
    
    int main() {
    	// Any library initialization comes here.
    	answer = 42;
    }
    
    // Use the -99lib option to prevent the linker from eliminating this function.
    int f42(int arg) {
    	return arg*answer;
    }
    $ cat main.go
    //go:generate 99c -99lib lib42.c
    
    package main
    
    import (
    	"fmt"
    	"os"
    
    	"github.com/cznic/ir"
    	"github.com/cznic/virtual"
    	"github.com/cznic/xc"
    )
    
    func main() {
    	f, err := os.Open("a.out")
    	if err != nil {
    		panic(err)
    	}
    
    	var bin virtual.Binary
    	if _, err := bin.ReadFrom(f); err != nil {
    		panic(err)
    	}
    
    	m, _, err := virtual.New(&bin, nil, nil, nil, nil, 0, 1<<10, "")
    	if err != nil {
    		panic(err)
    	}
    
    	defer m.Close()
    
    	pc, ok := bin.Sym[ir.NameID(xc.Dict.SID("f42"))]
    	if !ok {
    		panic("symbol not found")
    	}
    
    	t, err := m.NewThread(1 << 10)
    	if err != nil {
    		panic(err)
    	}
    
    	for _, v := range []int{1, 2, 3} {
    		var y int32
    		_, err := t.FFI1(pc, virtual.Int32Result{&y}, virtual.Int32(int32(v)))
    		if err != nil {
    			panic(err)
    		}
    
    		fmt.Println(y)
    	}
    }
    $ go generate && go run main.go
    42
    84
    126
    $

### Inserting defines

Use the -D flag to define additional macros on the command line.

    $ cd examples/define/
    /home/jnml/src/github.com/cznic/99c/examples/define
    $ ls *
    main.c
    $ cat main.c
    #include <stdio.h>
    
    int main() {
    #ifdef VERBOSE
    	printf(GREETING);
    #endif
    }
    $ 99c -DVERBOSE -DGREETING=\"hello\\n\" main.c && ./a.out
    hello
    $ 99c -DGREETING=\"hello\\n\" main.c && ./a.out
    $

### Specifying include paths

The -I flag defines additional include files search path(s).

    $ cd examples/include/
    /home/jnml/src/github.com/cznic/99c/examples/include
    $ ls *
    main.c
    
    foo:
    main.h
    $ cat main.c
    #include <stdio.h>
    #include "main.h"
    
    int main() {
    	printf(HELLO);
    }
    $ cat foo/main.h
    #ifndef _MAIN_H_
    #define _MAIN_H_
    
    #define HELLO "hello\n"
    
    #endif
    $ 99c main.c && ./a.out
    99c: main.c:2:10: include file not found: main.h (and 2 more errors)
    $ 99c -Ifoo main.c && ./a.out
    hello
    $

# 99run

Command 99run executes binary programs produced by the 99c compiler.

### Usage

    $ 99run a.out

On Linux a.out can be executed directly.

### Installation

To install or update 99run

    $ go get [-u] github.com/cznic/99c/99run

Online documentation: [godoc.org/github.com/cznic/99c/99run](http://godoc.org/github.com/cznic/99c/99run)

### Changelog

2017-10-07: Initial public release.

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

    $ cd examples/hello/
    /home/jnml/src/github.com/cznic/99c/examples/hello
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

# 99strace

Command 99strace traces system calls of programs produced by the 99c compiler.

The trace is written to stderr.

### Usage

    99strace a.out [arguments]

### Installation

To install or update 99strace

    $ go get [-u] -tags virtual.strace github.com/cznic/99c/99strace

Online documentation: [godoc.org/github.com/cznic/99c/99strace](http://godoc.org/github.com/cznic/99c/99strace)

### Changelog

2017-10-09: Initial public release.

### Sample

    $ cd examples/strace/
    /home/jnml/src/github.com/cznic/99c/examples/strace
    $ ls *
    data.txt  main.c
    $ cat main.c 
    #include <stdlib.h>
    #include <fcntl.h>
    #include <unistd.h>
    
    #define BUFSIZE 1<<16
    
    int main(int argc, char **argv) {
    	char *buf = malloc(BUFSIZE);
    	if (!buf) {
    		return 1;
    	}
    
    	for (int i = 1; i < argc; i++) {
    		int fd = open(argv[i], O_RDWR);
    		if (fd < 0) {
    			return 1;
    		}
    
    		ssize_t n;
    		while ((n = read(fd, buf, BUFSIZE)) > 0) {
    			write(0, buf, n);
    		}
    	}
    	free(buf);
    }
    $ cat data.txt 
    Lorem ipsum.
    $ 99c main.c && ./a.out data.txt
    Lorem ipsum.
    $ 99strace a.out data.txt
    malloc(0x10000) 0x7f83a9400020
    open("data.txt", O_RDWR, 0) 5 errno 0
    read(5, 0x7f83a9400020, 65536) 13 errno 0
    Lorem ipsum.
    write(0, 0x7f83a9400020, 13) 13 errno 0
    read(5, 0x7f83a9400020, 65536) 0 errno 0
    freep(0x7f83a9400020)
    $ 

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
    /home/jnml/src/github.com/cznic/99c/examples/hello
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

    $ cd examples/prof/
    /home/jnml/src/github.com/cznic/99c/examples/prof
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

### Bogomips

Let's try to estimate the VM bogomips value on an older Intel® Core™ i5-4670 CPU @ 3.40GHz × 4 machine.

    $ cd ../examples/prof/
    $ ls *
    bogomips.c  fib.c
    $ cat bogomips.c 
    #include <stdlib.h>
    #include <stdio.h>
    
    // src: https://en.wikipedia.org/wiki/BogoMips#Computation_of_BogoMIPS
    static void delay_loop(long loops) {
    	long d0 = loops;
    	do {
    		--d0;
    	} while (d0 >= 0);
    }
    
    int main(int argc, char **argv) {
    	if (argc != 2) {
    		return 2;
    	}
    
    	int n = atoi(argv[1]);
    	if (n<=0) {
    		return 1;
    	}
    
    	delay_loop(n);
    }
    $ 99c bogomips.c && 99prof -functions a.out 9900000
    # [99prof -functions a.out 9900000] 999.017618ms, 79.279 MIPS
    # functions
    delay_loop	     79200    100.00%    100.00%
    _start    	         1      0.00%    100.00%
    $ time ./a.out 34900000
    
    real	0m1,001s
    user	0m0,996s
    sys	0m0,000s
    $

In both cases the program executes for ~1 second. 34900000/9900000 = 3.525 and that's the slowdown coefficient when running the binary under 99prof. The bogomips value is thus ~279 MIPS on this machine.

    $ 99dump a.out 
    virtual.Binary a.out: code 0x0004d, text 0x00000, data 0x00030, bss 0x00020, pc2func 3, pc2line 23
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
    0x00016	func	variables      [0x8]byte	; bogomips.c:12:1
    0x00017		push           ap		; bogomips.c:12:1
    0x00018		zero32         			; bogomips.c:12:1
    0x00019		store32        			; bogomips.c:12:1
    0x0001a		add            sp, 0x8		; bogomips.c:12:1
    0x0001b		push32         (ap-0x8)		; bogomips.c:13:1
    0x0001c		push32         0x2		; bogomips.c:13:1
    0x0001d		neqi32         			; bogomips.c:13:1
    0x0001e		jz             0x23		; bogomips.c:13:1
    
    0x0001f		push           ap	; bogomips.c:14:1
    0x00020		push32         0x2	; bogomips.c:14:1
    0x00021		store32        		; bogomips.c:14:1
    0x00022		return         		; bogomips.c:14:1
    
    0x00023		push           bp-0x8		; bogomips.c:13:1
    0x00024		sub            sp, 0x8		; bogomips.c:17:1
    0x00025		arguments      			; bogomips.c:17:1
    0x00026		push64         (ap-0x10)	; bogomips.c:17:1
    0x00027		push32         0x1		; bogomips.c:17:1
    0x00028		indexi32       0x8		; bogomips.c:17:1
    0x00029		load64         0x0		; bogomips.c:17:1
    0x0002a		#atoi          			; bogomips.c:17:1
    0x0002b		store32        			; bogomips.c:17:1
    0x0002c		add            sp, 0x8		; bogomips.c:17:1
    0x0002d		push32         (bp-0x8)		; bogomips.c:18:1
    0x0002e		zero32         			; bogomips.c:18:1
    0x0002f		leqi32         			; bogomips.c:18:1
    0x00030		jz             0x35		; bogomips.c:18:1
    
    0x00031		push           ap	; bogomips.c:19:1
    0x00032		push32         0x1	; bogomips.c:19:1
    0x00033		store32        		; bogomips.c:19:1
    0x00034		return         		; bogomips.c:19:1
    
    0x00035		arguments      		; bogomips.c:18:1
    0x00036		push32         (bp-0x8)	; bogomips.c:22:1
    0x00037		convi32i64     		; bogomips.c:22:1
    0x00038		call           0x3f	; bogomips.c:22:1
    0x00039		return         		; bogomips.c:23:1
    
    0x0003a		builtin        	; atoi:69:1
    0x0003b		#atoi          	; atoi:69:1
    0x0003c		ffireturn      	; atoi:69:1
    
    0x0003d		call           0x3f	; atoi:69:1
    0x0003e		ffireturn      		; atoi:69:1
    
    # delay_loop
    0x0003f	func	variables      [0x8]byte		; bogomips.c:5:1
    0x00040		push           bp-0x8			; bogomips.c:6:1
    0x00041		push64         (ap-0x8)			; bogomips.c:6:1
    0x00042		store64        				; bogomips.c:6:1
    0x00043		add            sp, 0x8			; bogomips.c:6:1
    0x00044		push           bp-0x8			; bogomips.c:7:1
    0x00045		preinci64      0xffffffffffffffff	; bogomips.c:8:1
    0x00046		add            sp, 0x8			; bogomips.c:8:1
    0x00047		push64         (bp-0x8)			; bogomips.c:9:1
    0x00048		zero32         				; bogomips.c:9:1
    0x00049		convi32i64     				; bogomips.c:9:1
    0x0004a		geqi64         				; bogomips.c:9:1
    0x0004b		jnz            0x44			; bogomips.c:9:1
    
    0x0004c		return         	; bogomips.c:10:1
    
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
    0x0003a	function	atoi
    0x00014	function	main
    $

Alternatively, using 99dump, we can see that the loop consists of 8 instructions at addresses 0x00044-0x0004b. 34900000*8 = 279200000 confirming the estimated ~279MIPS value.
