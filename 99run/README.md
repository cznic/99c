# Table of Contents

1. [99run](#99run)
     1. Changelog
     1. Installation
     1. Executing compiled programs

# 99run

Command 99run executes binary programs produced by the 99c compiler.

### Changelog

2017-01-07: Initial public release.

### Installation

To install or update 99run

     $ go get [-u] github.com/cznic/99c/99run

Online documentation: [godoc.org/github.com/cznic/99c/99run](http://godoc.org/github.com/cznic/99c/99run)

### Executing compiled programs

Running a binary on Linux

     $ ./a.out
     hello world
     $

Running a binary on Windows

     C:\> 99run a.out
     hello world
     C:\>
