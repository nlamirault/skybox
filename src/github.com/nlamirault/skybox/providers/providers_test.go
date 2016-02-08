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
	"testing"
)

// https://www.reddit.com/r/golang/comments/44cjvy/http_response_cant_retrieve_cookies/
func TestLiveboxReadCookie(t *testing.T) {
	header := "Set-Cookie [25200fcf/sessid=Cusei7vG93RWDrabChZ9SlNJ; Path=/]"
	c := readCookie(header)
	str := generateCookie(c)
	if str != header {
		t.Fatalf("Livebox invalid read cookie: %#v %s", c, str)
	}
}
