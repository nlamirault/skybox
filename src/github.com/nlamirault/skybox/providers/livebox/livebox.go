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
	//"log"
	"net/http"
	"net/url"

	"github.com/nlamirault/skybox/config"
	"github.com/nlamirault/skybox/providers"
)

func init() {
	providers.Add("livebox", func() providers.Provider {
		return New()
	})
}

// Client is the Livebox API client
type Client struct {
	// The HTTP client to use when sending requests.
	Client *http.Client
	// Endpoint is the base URL for API requests.
	Endpoint  *url.URL
	ContextID string
	Username  string
	Password  string
	Cookies   []*http.Cookie
}

// New returns a Freebox Client
func New() *Client {
	baseURL, _ := url.Parse(defaultURL)
	client := Client{
		Client:   &http.Client{},
		Endpoint: baseURL,
		Username: "admin",
		Password: "admin",
	}
	return &client
}

func (c *Client) Description() string {
	return "livebox"
}

func (c *Client) EndPoint() *url.URL {
	return c.Endpoint
}

func (c *Client) GetHTTPClient() *http.Client {
	return c.Client
}

func (c *Client) SetupHeaders(request *http.Request) {
	request.Header.Add("Content-Type", providers.MediaType)
	request.Header.Add("Accept", providers.AcceptHeader)
	request.Header.Add("User-Agent", providers.UserAgent)
	request.Header.Add("X-Context", c.ContextID)
	request.Header.Add("X-Sah-Request-Type", "idle")
	request.Header.Add("X-Requested-With", "XMLHttpRequest")
	if c.Cookies != nil {
		for _, cookie := range c.Cookies {
			request.AddCookie(cookie)
		}
	}
}

func (c *Client) Setup(config *config.Configuration) error {
	if config.Livebox == nil {
		return fmt.Errorf("Livebox configuration not found: %v", config)
	}
	url, err := url.Parse(config.Livebox.URL)
	if err != nil {
		return fmt.Errorf("Livebox configuration invalid: %s", err.Error())
	}
	c.Endpoint = url
	c.Username = config.Livebox.Username
	c.Password = config.Livebox.Password
	return nil
}

func (c *Client) Ping() error {
	return fmt.Errorf("Not implemented")
}

func (c *Client) Authenticate() error {
	return fmt.Errorf("Not implemented")
}

func (c *Client) Statistics() (*providers.ProviderConnectionStatistics, error) {
	return nil, fmt.Errorf("Not implemented")
}
