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
     1. Installing C packages
     1. Talking to X server
     1. Creating a X window
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
1. [99nm](#99nm)
     1. Usage
     1. Installation
     1. Changelog
     1. Sample

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
      -E    Copy C-language source files to standard output, executing all
            preprocessor directives; no compilation shall be performed. If any
            operand is not a text file, the effects are unspecified.
      -Ipath
            Add path to the include files search paths.
      -Lpath
            Add path to search paths for -l.
      -Olevel
            Optimization setting, ignored.
      -Wwarn
            Warning level, ignored.
      -ansi
            Ignored.
      -c    Suppress the link-edit phase of the compilation, and do not
            remove any object files that are produced.
      -g    Produce debugging information.
      -l<name>
            Link with lib<name>.
      -o pathname
            Use the specified pathname, instead of the default a.out, for
            the executable file produced. If the -o option is present with
            -c or -E, the result is unspecified.
      -pedantic
            Ignored.
      -pthread
            Ignored. (TODO)
      -rdynamic
            Ignored. (TODO)
      -rpath pathname
            Ignored. (TODO)
      -shared
            Link mode shared library.
      -soname arg
            Ignored. (TODO)
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


Rest of the input is a list of file names, either C (.c) files or object (.o, .a) files.

### Installation

To install or update the compiler and the virtual machine

    $ go get [-u] github.com/cznic/99c github.com/cznic/99c/99run

To update the toolchain and rebuild all commands

    $ go generate

Use the -x flag to view the commands executed.

Online documentation: [godoc.org/github.com/cznic/99c](http://godoc.org/github.com/cznic/99c)

### Changelog

2017-10-19: Handle ar files (.a).

2017-10-18: Executables should be from now on no more tied to a single compatibility number but to a minimal compatibility number. No more permanent recompiling of everything.

2017-10-18: Initial support for using C packages.

2017-10-18: The -g flag is no more ignored. Add the -g flag to have the symbol and line information included in the executable. Without using -g some tools may not work and stack traces will not be really useful. The advantage of not including the additional info by default are substantially smaller executables.

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

### Installing C packages

To use a C package with programs compiled by 99c it's necessary to install a 99c version of the package. The lib directory contains some such installers. For example

    $ cd lib/xcb
    $ go generate

or equivalently

    $ go generate github.com/cznic/99c/lib/xcb

will install the 99c version of libxcb on your system in '$HOME/.99c'. Currently supported only on Linux.

### Talking to X server

A bare bones example, currently supported only on Linux.

    $ go generate ./lib/xcb ./lib/xau
    ... lot of output
    $ cd examples/xcb/
    /home/jnml/src/github.com/cznic/99c/examples/xcb
    $ ls
    screen.c
    $ cat screen.c 
    // +build ignore
    
    // src: https://xcb.freedesktop.org/tutorial/
    
    #include <stdio.h>
    #include <xcb/xcb.h>
    #include <inttypes.h>
    
    int main()
    {
    	/* Open the connection to the X server. Use the DISPLAY environment variable */
    
    	int i, screenNum;
    	xcb_connection_t *connection = xcb_connect(NULL, &screenNum);
    
    	/* Get the screen whose number is screenNum */
    
    	const xcb_setup_t *setup = xcb_get_setup(connection);
    	xcb_screen_iterator_t iter = xcb_setup_roots_iterator(setup);
    
    	// we want the screen at index screenNum of the iterator
    	for (i = 0; i < screenNum; ++i) {
    		xcb_screen_next(&iter);
    	}
    
    	xcb_screen_t *screen = iter.data;
    
    	/* report */
    
    	printf("\n");
    	printf("Informations of screen %" PRIu32 ":\n", screen->root);
    	printf("  width.........: %" PRIu16 "\n", screen->width_in_pixels);
    	printf("  height........: %" PRIu16 "\n", screen->height_in_pixels);
    	printf("  white pixel...: %" PRIu32 "\n", screen->white_pixel);
    	printf("  black pixel...: %" PRIu32 "\n", screen->black_pixel);
    	printf("\n");
    
    	return 0;
    }
    $ 99c screen.c -lxcb && ./a.out
    
    Informations of screen 927:
      width.........: 1920
      height........: 1200
      white pixel...: 16777215
      black pixel...: 0
    
    $ 

### Creating a X window

This example will show a small 150x150 pixel window in the top left corner of the screen. The window content is not handled by this example, but it can be moved, resized and closed.

    $ go generate ./lib/xcb ./lib/xau
    ... lot of output
    $ cd examples/xcb/
    $ cat helloworld.c 
    // +build ignore
    
    // src: https://www.x.org/releases/current/doc/libxcb/tutorial/index.html#helloworld
    
    #include <stdio.h>
    #include <unistd.h>		/* pause() */
    
    #include <xcb/xcb.h>
    
    int main()
    {
    	xcb_connection_t *c;
    	xcb_screen_t *screen;
    	xcb_window_t win;
    
    	/* Open the connection to the X server */
    	c = xcb_connect(NULL, NULL);
    
    	/* Get the first screen */
    	screen = xcb_setup_roots_iterator(xcb_get_setup(c)).data;
    
    	/* Ask for our window's Id */
    	win = xcb_generate_id(c);
    
    	/* Create the window */
    	xcb_create_window(c,	/* Connection          */
    			  XCB_COPY_FROM_PARENT,	/* depth (same as root) */
    			  win,	/* window Id           */
    			  screen->root,	/* parent window       */
    			  0, 0,	/* x, y                */
    			  150, 150,	/* width, height       */
    			  10,	/* border_width        */
    			  XCB_WINDOW_CLASS_INPUT_OUTPUT,	/* class               */
    			  screen->root_visual,	/* visual              */
    			  0, NULL);	/* masks, not used yet */
    
    	/* Map the window on the screen */
    	xcb_map_window(c, win);
    
    	/* Make sure commands are sent before we pause, so window is shown */
    	xcb_flush(c);
    
    	printf("Close the demo window and/or press ctrl-c while the terminal is focused to exit.\n");
    	int i = pause();		/* hold client until Ctrl-C */
    	printf("pause() returned %i\n", i);
    
    	return 0;
    }
    $ 99c helloworld.c -lxcb && ./a.out
    Close the demo window and/or press ctrl-c while the terminal is focused to exit.
    ... close the demo window (optional)
    ... focus the terminal and press ctrl-c
    ^Cpause() returned -1
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
    bogomips.c  fib.c
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
    $ 99c fib.c -g && 99prof -functions -lines -instructions a.out 32 2>log
    2178309
    $ cat log
    # [99prof -functions -lines -instructions a.out 32] 1.109159422s, 82.621 MIPS
    # functions
    fib   	     91639    100.00%    100.00%
    _start	         1      0.00%    100.00%
    # lines
    fib.c:12:	     52722     57.53%     57.53%
    fib.c:6:	     14133     15.42%     72.95%
    fib.c:5:	      7125      7.77%     80.73%
    fib.c:10:	      6611      7.21%     87.94%
    fib.c:8:	      4076      4.45%     92.39%
    fib.c:11:	      3528      3.85%     96.24%
    fib.c:9:	      2127      2.32%     98.56%
    fib.c:7:	      1317      1.44%    100.00%
    /home/jnml/src/github.com/cznic/ccir/libc/crt0.c:13:	         1      0.00%    100.00%
    # instructions
    Argument32     14175     15.47%	     15.47%
    Push32          9240     10.08%	     25.55%
    Func            7126      7.78%	     33.33%
    Call            7100      7.75%	     41.07%
    Return          7084      7.73%	     48.81%
    AddSP           7040      7.68%	     56.49%
    SwitchI32       7019      7.66%	     64.15%
    SubI32          7003      7.64%	     71.79%
    Store32         7002      7.64%	     79.43%
    AP              6972      7.61%	     87.04%
    Arguments       6968      7.60%	     94.64%
    AddI32          3528      3.85%	     98.49%
    Zero32          1383      1.51%	    100.00%
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
    $ 99c bogomips.c -g && 99prof -functions a.out 11370000
    # [99prof -functions a.out 11370000] 996.425292ms, 91.287 MIPS
    # functions
    delay_loop	     90960    100.00%    100.00%
    _start    	         1      0.00%    100.00%
    $ time ./a.out 36600000
    
    real	0m1,007s
    user	0m1,004s
    sys	0m0,004s
    $

In both cases the program executes for ~1 second. 36600000/113700000 = 3.219 and that's the slowdown coefficient when running the binary under 99prof. The bogomips value is thus ~293 MIPS on this machine.

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

Alternatively, using 99dump, we can see that the loop consists of 8 instructions at addresses 0x00044-0x0004b. 36600000*8 = 292800000 confirming the estimated ~293 MIPS value.

# 99nm

Command 99nm lists names in object and executable files produced by the 99c compiler.

### Usage

    99nm [files...]

### Installation

To install or update 99nm

     $ go get [-u] github.com/cznic/99c/99nm

Online documentation: [godoc.org/github.com/cznic/99c/99nm](http://godoc.org/github.com/cznic/99c/99nm)

### Changelog

2017-10-14: Initial public release.

### Sample

    $ cd ../examples/nm/
    $ ls *
    foo.c
    $ cat foo.c
    int i;
    
    static int j;
    
    int foo()
    {
    }
    
    static int bar()
    {
    }
    
    int main()
    {
    }
    $ 99c foo.c && 99nm a.out
    0x00012	__builtin_exit
    0x0000f	__register_stdfiles
    0x00000	_start
    0x00014	main
    $ 99c -99lib foo.c && 99nm a.out
    0x00077	__builtin_abort
    0x00079	__builtin_abs
    0x0001b	__builtin_alloca
    0x0001e	__builtin_bswap64
    0x00021	__builtin_clrsb
    0x00024	__builtin_clrsbl
    0x00027	__builtin_clrsbll
    0x0002a	__builtin_clz
    0x0002d	__builtin_clzl
    0x00030	__builtin_clzll
    0x0005f	__builtin_copysign
    0x00033	__builtin_ctz
    0x00036	__builtin_ctzl
    0x00039	__builtin_ctzll
    0x00019	__builtin_exit
    0x00091	__builtin_ffs
    0x00094	__builtin_ffsl
    0x00097	__builtin_ffsll
    0x00068	__builtin_fopen
    0x0006b	__builtin_fprintf
    0x0003c	__builtin_frame_address
    0x00056	__builtin_isprint
    0x00062	__builtin_longjmp
    0x00074	__builtin_malloc
    0x00085	__builtin_memcmp
    0x0008b	__builtin_memcpy
    0x0007c	__builtin_memset
    0x0003f	__builtin_parity
    0x00042	__builtin_parityl
    0x00045	__builtin_parityll
    0x00048	__builtin_popcount
    0x0004b	__builtin_popcountl
    0x0004e	__builtin_popcountll
    0x0006e	__builtin_printf
    0x00051	__builtin_return_address
    0x00065	__builtin_setjmp
    0x00071	__builtin_sprintf
    0x0008e	__builtin_strchr
    0x0007f	__builtin_strcmp
    0x00082	__builtin_strcpy
    0x00088	__builtin_strlen
    0x00054	__builtin_trap
    0x00016	__register_stdfiles
    0x00059	__signbit
    0x0005c	__signbitf
    0x00007	_start
    0x0009a	foo
    0x00000	main
    $ 99c -c foo.c && 99nm foo.o
    __builtin_abort			func()
    __builtin_abs			func(int32)int32
    __builtin_alloca		func(uint64)*struct{}
    __builtin_bswap64		func(uint64)uint64
    __builtin_clrsb			func(int32)int32
    __builtin_clrsbl		func(int64)int32
    __builtin_clrsbll		func(int64)int32
    __builtin_clz			func(uint32)int32
    __builtin_clzl			func(uint64)int32
    __builtin_clzll			func(uint64)int32
    __builtin_copysign		func(float64,float64)float64
    __builtin_ctz			func(uint32)int32
    __builtin_ctzl			func(uint64)int32
    __builtin_ctzll			func(uint64)int32
    __builtin_exit			func(int32)
    __builtin_ffs			func(int32)int32
    __builtin_ffsl			func(int64)int32
    __builtin_ffsll			func(int64)int32
    __builtin_fopen			func(*int8,*int8)*struct{}
    __builtin_fprintf		func(*struct{},*int8...)int32
    __builtin_frame_address		func(uint32)*struct{}
    __builtin_isprint		func(int32)int32
    __builtin_longjmp		func(*struct{},int32)
    __builtin_malloc		func(uint64)*struct{}
    __builtin_memcmp		func(*struct{},*struct{},uint64)int32
    __builtin_memcpy		func(*struct{},*struct{},uint64)*struct{}
    __builtin_memset		func(*struct{},int32,uint64)*struct{}
    __builtin_parity		func(uint32)int32
    __builtin_parityl		func(uint64)int32
    __builtin_parityll		func(uint64)int32
    __builtin_popcount		func(uint32)int32
    __builtin_popcountl		func(uint64)int32
    __builtin_popcountll		func(uint64)int32
    __builtin_printf		func(*int8...)int32
    __builtin_return_address	func(uint32)*struct{}
    __builtin_setjmp		func(*struct{})int32
    __builtin_sprintf		func(*int8,*int8...)int32
    __builtin_strchr		func(*int8,int32)*int8
    __builtin_strcmp		func(*int8,*int8)int32
    __builtin_strcpy		func(*int8,*int8)*int8
    __builtin_strlen		func(*int8)uint64
    __builtin_trap			func()
    __register_stdfiles		func(*struct{},*struct{},*struct{})
    __signbit			func(float64)int32
    __signbitf			func(float32)int32
    foo				func()int32
    i				int32
    main				func()int32
    $ 
