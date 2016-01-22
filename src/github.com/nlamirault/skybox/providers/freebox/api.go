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
)

// APIVersionResponse is returned by requesting `GET /api_version`
type APIVersionResponse struct {
	FreeboxID  string `json:"uid",omitempty`
	DeviceName string `json:"device_name",omitempty`
	Version    string `json:"api_version",omitempty`
	BaseURL    string `json:"api_base_url",omitempty`
	DeviceType string `json:"device_type",omitempty`
}

// APIAuthorizeRequest is sent by requesting `POST /api/v3/login/authorize/`
type APIAuthorizeRequest struct {
	AppID      string `json:"app_id",omitempty`
	AppName    string `json:"app_name",omitempty`
	AppVersion string `json:"app_version",omitempty`
	DeviceName string `json:"device_name",omitempty`
}

// APIAuthorizeResponse is returned by requesting `POST /api/v3/login/authorize/`
type APIAuthorizeResponse struct {
	Success bool `json:"success",omitempty`
	Result  struct {
		AppToken string `json:"app_token",omitempty`
		TrackID  int    `json:"track_id",omitempty`
	}
}

// APIConnectionStatusResponse is returned by requesting `GET /api/v3/connection/`
type APIConnectionStatusResponse struct {
	Success bool `json:"success",omitempty`
	Result  struct {
		// ehernet FTTH, or rfc2684 xDSL (unbundled), or pppoatm xDSL
		Type string `json:"type",omitempty`
		// current download rate in byte/s
		RateDown int `json:"rate_down",omitempty`
		// current download rate in byte/s
		RateUp int `json:"rate_up",omitempty`
		// total downloaded bytes since last connection
		BytesDown int `json:"bytes_down",omitempty`
		// total uploaded bytes since last connection
		BytesUp int `json:"bytes_up",omitempty`
		// available upload bandwidth in bit/s
		BandwidthUp int `json:"bandwidth_up",omitempty`
		// available download bandwidth in bit/s
		BandwidthDown int `json:"bandwidth_down",omitempty`
		// Freebox IPv4 address
		IPv4 string `json:"ipv4",omitempty`
		// Freebox IPv6 address
		IPv6 string `json:"ipv6",omitempty`
		// State of the connection
		State string `json:"state",omitempty`
		// ftth	FTTH or xdsl xDSL
		Media string `json:"media",omitempty`
	}
}

type APILoginRequest struct {
	AppID      string `json:"app_id",omitempty`
	AppVersion string `json:"app_version",omitempty`
	Password   string `json:"password",omitempty`
}

type APILoginResponse struct {
	Success bool `json:"success"`
	Result  struct {
		SessionToken string          `json:"session_token",omitempty`
		Challenge    string          `json:"challenge",omitempty`
		PasswordSalt string          `json:"",omitempty`
		Permissions  map[string]bool `json:"",omitempty`
	} `json:"result"`
}
