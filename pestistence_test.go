// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pcookiejar

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"strings"
	"testing"
	"time"
)

// run runs the jarTest.
func (test jarTest) persist(t *testing.T, jar *Jar) {
	now := tNow

	// Populate jar with cookies.
	setCookies := make([]*http.Cookie, len(test.setCookies))
	for i, cs := range test.setCookies {
		cookies := (&http.Response{Header: http.Header{"Set-Cookie": {cs}}}).Cookies()
		if len(cookies) != 1 {
			panic(fmt.Sprintf("Wrong cookie line %q: %#v", cs, cookies))
		}
		setCookies[i] = cookies[0]
	}
	jar.setCookies(mustParseURL(test.fromURL), setCookies, now)
	now = now.Add(1001 * time.Millisecond)

	// Serialize non-expired entries in the form "name1=val1 name2=val2".
	var cs []string
	for _, submap := range jar.entries {
		for _, cookie := range submap {
			if !cookie.Expires.After(now) {
				continue
			}
			cs = append(cs, cookie.Name+"="+cookie.Value)
		}
	}
	sort.Strings(cs)
	got := strings.Join(cs, " ")

	// Make sure jar content matches our expectations.
	if got != test.content {
		t.Errorf("Test %q Content\ngot  %q\nwant %q",
			test.description, got, test.content)
	}

	// Test different calls to Cookies.
	for i, query := range test.queries {
		now = now.Add(1001 * time.Millisecond)
		var s []string
		jsonJar, err := json.Marshal(jar)
		if err != nil {
			t.Error(err)
		}
		newJar := newTestJar()
		err = json.Unmarshal(jsonJar, newJar)
		if err != nil {
			t.Error(err)
		}
		for _, c := range newJar.cookies(mustParseURL(query.toURL), now) {
			s = append(s, c.Name+"="+c.Value)
		}
		if got := strings.Join(s, " "); got != query.want {
			t.Errorf("Test %q #%d\ngot  %q\nwant %q", test.description, i, got, query.want)
		}
	}
}

func TestBasicsPersist(t *testing.T) {
	for _, test := range basicsTests {
		jar := newTestJar()
		test.persist(t, jar)
	}
}