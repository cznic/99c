// Copyright 2017 The 99c Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:generate wget --no-clobber https://www.x.org/releases/individual/lib/libXau-1.0.8.tar.gz
//go:generate rm -rf libXau-1.0.8/
//go:generate tar xzf libXau-1.0.8.tar.gz
//go:generate sh -c "cd libXau-1.0.8/ && ./configure CC=99c --prefix=$HOME/.99c"
//go:generate make -C libXau-1.0.8/ install

// Package xau installs a 99c version of libXau on your system.
//
// Run
//
//     $ go generate
//
// to download and install the package on your system. The package will be
// installed in '$HOME/.99c'. Currently supported only on Linux.
package xau
