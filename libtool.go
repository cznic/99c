// Copyright 2017 The 99c Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

var (
	libToolConfigs = map[string]libToolConfig{}
)

type libToolConfig map[string]interface{}

func newLibToolConfig(r io.Reader) (libToolConfig, error) {
	s := bufio.NewScanner(r)
	c := libToolConfig{}
	for s.Scan() {
		l := strings.TrimSpace(s.Text())
		if l == "" || strings.HasPrefix(l, "#") {
			continue
		}

		a := strings.SplitN(l, "=", 2)
		if len(a) != 2 {
			return nil, fmt.Errorf("invalid libtool config line: %s", l)
		}

		nm := strings.TrimSpace(a[0])
		val := strings.TrimSpace(a[1])
		if _, ok := c[nm]; ok {
			return nil, fmt.Errorf("duplicate libtool config item: %s", l)
		}

		if strings.HasPrefix(val, "'") {
			if !strings.HasSuffix(val, "'") {
				return nil, fmt.Errorf("invalid libtool config value: %s", l)
			}

			c[nm] = val[1 : len(val)-1]
			continue
		}

		if val == "yes" {
			c[nm] = true
			continue
		}

		if val == "no" {
			c[nm] = false
			continue
		}

		n, err := strconv.ParseUint(val, 10, 31)
		if err == nil {
			c[nm] = int(n)
			continue
		}

		return nil, fmt.Errorf("invalid libtool config line: %s", l)
	}
	if err := s.Err(); err != nil {
		return nil, err
	}

	return c, nil
}

func newLibToolConfigFile(fn string) (libToolConfig, error) {
	if x, ok := libToolConfigs[fn]; ok {
		return x, nil
	}

	f, err := os.Open(fn)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	c, err := newLibToolConfig(f)
	if err != nil {
		return nil, err
	}

	libToolConfigs[fn] = c
	return c, nil
}

func (c libToolConfig) dependencyLibs() ([]string, error) {
	var r []string
	rq, ok := c["dependency_libs"]
	if !ok {
		return nil, nil
	}

	s, ok := rq.(string)
	if !ok {
		return nil, fmt.Errorf("invalid dependency_libs value: %T(%#v)", rq, rq)
	}

	s = strings.TrimSpace(s)
	a := strings.Split(s, " ")
	for _, v := range a {
		v = strings.TrimSpace(v)
		if v == "" {
			continue
		}

		if v == "-l" || !strings.HasPrefix(v, "-l") {
			return nil, fmt.Errorf("invalid dependency_libs value: %v", s)
		}

		r = append(r, v[2:])
	}
	return r, nil
}
