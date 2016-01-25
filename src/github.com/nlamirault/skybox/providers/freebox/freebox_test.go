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

package freebox

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/nlamirault/skybox/providers"
)

func newFreebox(handler http.HandlerFunc) (*Client, *httptest.Server, error) {
	server := httptest.NewServer(http.HandlerFunc(handler))
	fbx := New()
	fakeURL, err := url.Parse(server.URL)
	if err != nil {
		return nil, nil, err
	}
	fbx.Endpoint = fakeURL
	return fbx, server, nil
}

func TestMakeFreeboxAPIErrorResponse(t *testing.T) {
	apiError := &providers.APIError{
		StatusCode: 403,
		Message:    "{\"uid\":\"xxxxxxxxxxxx1cc1191ef\",\"success\":false,\"msg\":\"Vous devez vous connecter pour accéder à cette fonction\",\"result\":{\"password_salt\":\"xxxxxxxxxxxxx0K0Jd\",\"challenge\":\"xxxxxxxxxxxxxBHI9mT\"},\"error_code\":\"auth_required\"}",
	}
	freeboxAPIError, err := makeAPIErrorResponse(apiError)
	if err != nil {
		t.Fatal(err)
	}
	if freeboxAPIError.ErrorCode != authRequiredError {
		t.Fatal("Invalid Freebox API error parsing")
	}
}

func TestFreeboxAPIVersion(t *testing.T) {
	fbx, server, err := newFreebox(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", providers.AcceptHeader)
		fmt.Fprintln(w, `{
   "uid": "23b86ec8091013d668829fe12791fdab",
   "device_name": "Freebox Server",
   "api_version": "3.0",
   "api_base_url": "/api/",
   "device_type": "FreeboxServer1,1"
}`)
	})
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()

	resp, err := fbx.version()
	if err != nil {
		t.Fatalf("Error API call version : %v", err)
	}
	if resp.Version != "3.0" ||
		resp.FreeboxID != "23b86ec8091013d668829fe12791fdab" {
		t.Fatalf("Freebox API version response: %s", resp)
	}
}

func TestFreeboxAPIAuthorize(t *testing.T) {
	fbx, server, err := newFreebox(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", providers.AcceptHeader)
		fmt.Fprintln(w, `{
  "success": true,
  "result": {
      "app_token": "dyNYgfK0Ya6FWGqq83sBHa7TwzWo+pg4fDFUJHShcjVYzTfaRrZzm93p7OTAfH/0",
      "track_id": 42
   }
}`)
	})
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()

	resp, err := fbx.authorize()
	if err != nil {
		t.Fatalf("Error API call authorize: %v", err)
	}
	if resp.Result.AppToken != "dyNYgfK0Ya6FWGqq83sBHa7TwzWo+pg4fDFUJHShcjVYzTfaRrZzm93p7OTAfH/0" ||
		resp.Result.TrackID != 42 {
		t.Fatalf("Freebox API authorize response: %s", resp)
	}

}

func TestFreeboxAPILogin(t *testing.T) {
	fbx, server, err := newFreebox(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", providers.AcceptHeader)
		fmt.Fprintln(w, `{
  "success": true,
    "result": {
        "logged_in": false,
        "challenge": "VzhbtpR4r8CLaJle2QgJBEkyd8JPb0zL"
    }
}`)
	})
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()

	resp, err := fbx.login()
	if err != nil {
		t.Fatalf("Error API call login: %v", err)
	}
	if resp.Result.Challenge != "VzhbtpR4r8CLaJle2QgJBEkyd8JPb0zL" {
		t.Fatalf("Freebox API login response: %s", resp)
	}
	if fbx.Challenge != "VzhbtpR4r8CLaJle2QgJBEkyd8JPb0zL" {
		t.Fatalf("Freebox login challenge not set: %v", fbx)
	}

}

func TestFreeboxAPIOpenSession(t *testing.T) {
	fbx, server, err := newFreebox(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{
  "success": true,
  "result" : {
         "session_token": "35JYdQSvkcBYK84IFMU7H86clfhS75OzwlQrKlQN1gBchDd62RGzDpgC7YB9jB2",
         "challenge": "jdGL6CtuJ3Dm7p9nkcIQ8pjB+eLwr4Ya",
         "permissions": {
               "downloader": true
         }
    }
}`)
	})
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()

	resp, err := fbx.openSession()
	if err != nil {
		t.Fatalf("Error API call open session: %v", err)
	}
	if resp.Result.SessionToken != "35JYdQSvkcBYK84IFMU7H86clfhS75OzwlQrKlQN1gBchDd62RGzDpgC7YB9jB2" {
		t.Fatalf("Freebox API open session response: %s", resp)
	}
	if fbx.SessionToken != "35JYdQSvkcBYK84IFMU7H86clfhS75OzwlQrKlQN1gBchDd62RGzDpgC7YB9jB2" {
		t.Fatalf("Freebox session token not set: %v", fbx)
	}
}

func TestFreeboxAPICloseSession(t *testing.T) {
	fbx, server, err := newFreebox(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{
  "success": true
}`)
	})
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()

	resp, err := fbx.closeSession()
	if err != nil {
		t.Fatalf("Error API call close session: %v", err)
	}
	if !resp.Success {
		t.Fatalf("Freebox API close session response: %s", resp)
	}
	if fbx.SessionToken != "" {
		t.Fatalf("Freebox session token set: %v", fbx)
	}
}
