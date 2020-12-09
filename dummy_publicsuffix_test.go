// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pcookiejar_test

import (
	"github.com/carted/pcookiejar"
)

type dummypsl struct {
	List pcookiejar.PublicSuffixList
}

func (dummypsl) PublicSuffix(domain string) string {
	return domain
}

func (dummypsl) String() string {
	return "dummy"
}

var publicsuffix = dummypsl{}
