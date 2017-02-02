// Copyright (C) 2016, 2017 Nicolas Lamirault <nicolas.lamirault@gmail.com>

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
	"net/http"
	"net/url"

	"github.com/nlamirault/skybox/config"
)

// Creator return a new Provider
type Creator func() Provider

// Providers defines all available box providers
var Providers = map[string]Creator{}

// Add define a new box provider to the available providers
func Add(name string, creator Creator) {
	Providers[name] = creator
}

// Provider represents a client for a Box API
type Provider interface {

	// Description return the provider name
	Description() string

	// EndPoint returns the API base URL
	EndPoint() *url.URL

	// GetHTTPClient returns the HTTP client to use
	GetHTTPClient() *http.Client

	// SetupHeaders add customer headers
	SetupHeaders(request *http.Request)

	// Setup finalize the provider configuration
	Setup(config *config.Configuration) error

	// Ping call the box provider to check connection
	Ping() error

	// Authenticate perform a call to authenticate the application
	Authenticate() error

	// Statistics retrieve box provider statistics
	Statistics() (*ConnectionStatistics, error)

	// Description retrieve some description about the box provider
	Network() (*NetworkInformations, error)

	// Wifi retrieve informations about the wifi
	Wifi() (*WifiStatus, error)

	// TV retrieve informations about the TV
	TV() ([]*TVStatus, error)

	// Devices retrieve connected devices
	Devices() ([]*BoxDevice, error)
}

// ConnectionStatistics represents commons statistics for box provider
type ConnectionStatistics struct {
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
}

// NetworkInformations define some informations about the box provider
type NetworkInformations struct {
	IPV4Address string `json:"ipv4_address"`
	IPV6Address string `json:"ipv6_address"`
	DNS         string `json:"dns"`
	State       string `json:"state"`
}

// WifiStatus define informations related to the wifi
type WifiStatus struct {
	State  bool `json:"state"`
	Enable bool `json:"enable"`
}

// TVStatus define the informations related to the TV
type TVStatus struct {
	State bool   `json:"state"`
	Name  string `json:"name"`
}

// BoxDevice define a device connected to the box provider network
type BoxDevice struct {
	Type      string `json:"type"`
	Name      string `json:"name"`
	IPAddress string `json:"ipaddress"`
}
