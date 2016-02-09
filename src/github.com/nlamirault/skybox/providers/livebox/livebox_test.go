// Copyright (C) 2016 Nicolas Lamirault <nicolas.lamirault@gmail.com>

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package livebox

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/nlamirault/skybox/providers"
)

func newLivebox(handler http.HandlerFunc) (*Client, *httptest.Server, error) {
	server := httptest.NewServer(http.HandlerFunc(handler))
	box := New()
	fakeURL, err := url.Parse(server.URL)
	if err != nil {
		return nil, nil, err
	}
	box.Endpoint = fakeURL
	return box, server, nil
}

func TestLiveboxAuthenticate(t *testing.T) {
	box, server, err := newLivebox(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", providers.AcceptHeader)
		w.Header().Set("Set-Cookie", "25200fcf/sessid=Cusei7vG93RWDrabChZ9SlNJ; Path=/")
		fmt.Fprintln(w, `{
     "status":0,
     "data": {
         "contextID":"RmjJzr2UIXk2zFteSiU0i1bK8wUuS8QyhZ6GeWoLKyC82T0K2TH9HGIF1sXJnD6s"
     }
}`)
	})
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()

	resp, err := box.authenticate()
	if err != nil {
		t.Fatalf("Error API call authenticate : %v", err)
	}
	if resp.Data.ContextID != "RmjJzr2UIXk2zFteSiU0i1bK8wUuS8QyhZ6GeWoLKyC82T0K2TH9HGIF1sXJnD6s" {
		t.Fatalf("Freebox API version response: %s", resp)
	}
	if box.ContextID != "RmjJzr2UIXk2zFteSiU0i1bK8wUuS8QyhZ6GeWoLKyC82T0K2TH9HGIF1sXJnD6s" {
		t.Fatalf("Livebox contextID not set: %v", box)
	}
	fmt.Printf("Cookies: %v\n", box.Cookies)
	if len(box.Cookies) != 1 {
		t.Fatalf("Livebox invalid cookies %d", len(box.Cookies))
	}
	if box.Cookies[0].Name != "25200fcf/sessid" || box.Cookies[0].Value != "Cusei7vG93RWDrabChZ9SlNJ" {
		t.Fatalf("Livebox invalid cookie %#v", box.Cookies[0])
	}
}
