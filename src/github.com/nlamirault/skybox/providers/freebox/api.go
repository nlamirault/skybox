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

const (
	FreeboxAPIVersion string = "v3"

	// API Errors code

	// Invalid session token, or not session token sent
	AuthRequiredError string = "auth_required"
	// The app token you are trying to use is invalid or has been revoked
	InvalidToken string = "invalid_token"
	// The app token you are trying to use has not been validated by user yet
	PendingToken string = "pending_token"
	// Your app permissions does not allow accessing this API
	InsufficientRights string = "insufficient_rights"

// denied_from_external_ip	You are trying to get an app_token from a remote IP
// invalid_request	Your request is invalid
// ratelimited	Too many auth error have been made from your IP
// new_apps_denied	New application token request has been disabled
// apps_denied	API access from apps has been disabled
// internal_error	Internal error
)

// APIVersionResponse is returned by requesting `GET /api_version`
type APIVersionResponse struct {
	FreeboxID  string `json:"uid"`
	DeviceName string `json:"device_name"`
	Version    string `json:"api_version"`
	BaseURL    string `json:"api_base_url"`
	DeviceType string `json:"device_type"`
}

type APIErrorResponse struct {
	UID       string `json:"uid"`
	Message   string `json:"msg"`
	Success   bool   `json:"success"`
	ErrorCode string `json:"error_code"`
}

// APIAuthorizeRequest is sent by requesting `POST /api/v3/login/authorize/`
type APIAuthorizeRequest struct {
	AppID      string `json:"app_id"`
	AppName    string `json:"app_name"`
	AppVersion string `json:"app_version"`
	DeviceName string `json:"device_name"`
}

// APIAuthorizeResponse is returned by requesting `POST /api/v3/login/authorize/`
type APIAuthorizeResponse struct {
	Success bool `json:"success"`
	Result  struct {
		AppToken string `json:"app_token"`
		TrackID  int    `json:"track_id"`
	}
}

// APIConnectionStatusResponse is returned by requesting `GET /api/v3/connection/`
type APIConnectionStatusResponse struct {
	Success bool `json:"success"`
	Result  struct {
		// ehernet FTTH, or rfc2684 xDSL (unbundled), or pppoatm xDSL
		Type string `json:"type"`
		// current download rate in byte/s
		RateDown int `json:"rate_down"`
		// current download rate in byte/s
		RateUp int `json:"rate_up"`
		// total downloaded bytes since last connection
		BytesDown int `json:"bytes_down"`
		// total uploaded bytes since last connection
		BytesUp int `json:"bytes_up"`
		// available upload bandwidth in bit/s
		BandwidthUp int `json:"bandwidth_up"`
		// available download bandwidth in bit/s
		BandwidthDown int `json:"bandwidth_down"`
		// Freebox IPv4 address
		IPv4 string `json:"ipv4"`
		// Freebox IPv6 address
		IPv6 string `json:"ipv6"`
		// State of the connection
		State string `json:"state"`
		// ftth	FTTH or xdsl xDSL
		Media string `json:"media"`
	}
}

type APILoginRequest struct {
	AppID      string `json:"app_id"`
	AppVersion string `json:"app_version"`
	Password   string `json:"password"`
}

type APILoginResponse struct {
	Success bool `json:"success"`
	Result  struct {
		SessionToken string          `json:"session_token"`
		Challenge    string          `json:"challenge"`
		PasswordSalt string          `json:""`
		Permissions  map[string]bool `json:""`
	} `json:"result"`
}
