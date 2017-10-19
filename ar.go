// Copyright 2017 The 99c Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"bufio"
	"io"
	"os"

	"github.com/blakesmith/ar"
	"github.com/cznic/ir"
)

func archive(fn string) (ir.Objects, error) {
	f, err := os.Open(fn)
	if err != nil {
		return nil, err
	}

	var a ir.Objects
	r := ar.NewReader(bufio.NewReader(f))
	for {
		if _, err := r.Next(); err != nil {
			if err == io.EOF {
				return a, nil
			}

			return nil, err
		}

		var o ir.Objects
		if _, err := o.ReadFrom(r); err != nil {
			return nil, err
		}

		if len(o) != 1 {
			panic("TODO")
		}

		a = append(a, o[0])
	}
}
