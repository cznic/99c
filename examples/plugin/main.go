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
