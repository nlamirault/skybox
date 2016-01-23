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
	"testing"

	"github.com/nlamirault/skybox/providers"
)

func TestMakeFreeboxAPIErrorResponse(t *testing.T) {
	apiError := &providers.APIError{
		StatusCode: 403,
		Message:    "{\"uid\":\"xxxxxxxxxxxx1cc1191ef\",\"success\":false,\"msg\":\"Vous devez vous connecter pour accéder à cette fonction\",\"result\":{\"password_salt\":\"xxxxxxxxxxxxx0K0Jd\",\"challenge\":\"xxxxxxxxxxxxxBHI9mT\"},\"error_code\":\"auth_required\"}",
	}
	freeboxAPIError, err := makeAPIErrorResponse(apiError)
	if err != nil {
		t.Fatal(err)
	}
	if freeboxAPIError.ErrorCode != AuthRequiredError {
		t.Fatal("Invalid Freebox API error parsing")
	}
}
