// Copyright 2017 The 99c Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:generate wget --no-clobber http://xcb.freedesktop.org/dist/libxcb-1.12.tar.gz
//go:generate rm -rf libxcb-1.12/
//go:generate tar xzf libxcb-1.12.tar.gz
//go:generate sh -c "cd libxcb-1.12/ && ./configure CC=99c --prefix=$HOME/.99c"
//go:generate make -C libxcb-1.12/ install

// Package xcb installs a 99c version of libxcb on your system.
//
// Run
//
//     $ go generate
//
// to download and install the package on your system. The package will be
// installed in '$HOME/.99c'. Currently supported only on Linux.
package xcb
