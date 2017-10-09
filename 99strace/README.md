# Table of Contents

1. Usage
1. Installation
1. Changelog
1. Sample

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

    $ cd ../example/strace/
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
