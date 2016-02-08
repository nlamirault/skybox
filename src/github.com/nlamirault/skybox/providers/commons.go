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

package providers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	//"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/nlamirault/skybox/version"
)

const (
	// AcceptHeader is the default Accept Header : application/json
	AcceptHeader = "application/json"

	// MediaType is the default mediaType header : application/json
	MediaType = "application/json"
)

var (
	// UserAgent represents the user agent used
	UserAgent = fmt.Sprintf("skybox/%s", version.Version)
)

func createRequest(method, uri string, body interface{}) (*http.Request, error) {
	buf := new(bytes.Buffer)
	if body != nil {
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, uri, buf)
	if err != nil {
		return nil, err
	}
	log.Printf("[DEBUG] HTTP Request : %v", req)
	return req, nil
}

func getURL(base *url.URL, urlStr string) (*url.URL, error) {
	rel, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}
	return base.ResolveReference(rel), nil
}

func performRequest(provider Provider, method, urlStr string, body interface{}) (*http.Response, error) {
	u, err := getURL(provider.EndPoint(), urlStr)
	if err != nil {
		return nil, err
	}
	req, err := createRequest(method, u.String(), body)
	if err != nil {
		return nil, err
	}
	provider.SetupHeaders(req)
	return provider.GetHTTPClient().Do(req)
}

// APIError represents an error from REST API
type APIError struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}

func (a *APIError) Error() string {
	return fmt.Sprintf("%d / %s", a.StatusCode, a.Message)
}

// Do perform a HTTP request using the REST API Client.
// body is used for the content of the request
// result contains the JSON decoded response
// apiError contains the JSON error response if HTTP status code isn't OK.
func Do(provider Provider, method, urlStr string, body interface{}, result interface{}) ([]*http.Cookie, error) {
	resp, err := performRequest(provider, method, urlStr, body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		return decodeResponse(resp, result)
	}
	content, err := getResponseBody(resp)
	if err != nil {
		return nil, fmt.Errorf("Can't read HTTP Error : %s", err.Error())
	}
	apiError := APIError{
		StatusCode: resp.StatusCode,
		Message:    content,
	}
	log.Printf("[DEBUG] HTTP Error: %v\n", apiError)
	return nil, &apiError
}

func decodeResponse(resp *http.Response, v interface{}) ([]*http.Cookie, error) {
	log.Printf("[DEBUG] Decode response")
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	cookies := resp.Cookies()
	if len(resp.Cookies()) == 0 { // Check invalid cookies
		// line := resp.Header.Get("Set-Cookie")
		// log.Printf("[DEBUG] Cookie: %s", line)
		// if len(line) > 0 {
		// 	c := readCookie(line)
		// 	if c != nil {
		// 		cookies = append(cookies, c)
		// 	}
		// }
		for k, v := range resp.Header {
			fmt.Printf("H: %s ** %s\n", k, v)
			if k == "Set-Cookie" {
				c := readCookie(fmt.Sprintf("%s %s", k, v))
				if c != nil {
					cookies = append(cookies, c)
				}
			}
		}

	}
	log.Printf("[DEBUG] HTTP Response: %d / %s / %v",
		resp.StatusCode, string(body), cookies)
	err = json.Unmarshal(body, v)
	if err != nil {
		return nil, err
	}
	return cookies, nil
}

func getResponseBody(resp *http.Response) (string, error) {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
