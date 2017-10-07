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
