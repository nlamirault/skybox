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
	"log"
	"net/http"
	"net/url"

	"github.com/nlamirault/skybox/config"
	"github.com/nlamirault/skybox/providers"
	"github.com/nlamirault/skybox/version"
)

func init() {
	providers.Add("freebox", func() providers.Provider {
		return New()
	})
}

// Client is the Freebox API client
type Client struct {
	// The client to use when sending requests.
	Client *http.Client
	// Endpoint is the base URL for API requests.
	Endpoint     *url.URL
	ID           string
	Token        string
	SessionToken string
	Challenge    string
	Password     string
	Identifier   string `json:"app_id"`
	Name         string `json:"app_name"`
	Version      string `json:"app_version"`
	DeviceName   string `json:"device_name"`
}

// New returns a Freebox Client
func New() *Client {
	baseURL, _ := url.Parse(defaultURL)
	client := Client{
		Client:     &http.Client{},
		Endpoint:   baseURL,
		Identifier: "skybox",
		Name:       "Skybox",
		Version:    fmt.Sprintf("%s", version.Version),
		DeviceName: "Skybox",
	}
	return &client
}

func (c *Client) Description() string {
	return "freebox"
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
	if c.SessionToken != "" {
		request.Header.Add("X-Fbx-App-Auth", c.SessionToken)
	}
}

func (c *Client) Setup(config *config.Configuration) error {
	if config.Freebox == nil {
		return fmt.Errorf("Freebox configuration not found: %v", config)
	}
	url, err := url.Parse(config.Freebox.URL)
	if err != nil {
		return fmt.Errorf("Freebox configuration invalid: %s", err.Error())
	}
	c.Endpoint = url
	c.Token = config.Freebox.Token
	return nil
}

// Ping contact the Freebox server, and check the API version
func (c *Client) Ping() error {
	_, err := c.version()
	if err != nil {
		return err
	}
	log.Printf("[DEBUG] Freebox Ping received")
	return nil
}

func (c *Client) Authenticate() error {
	if c.Token == "" {
		_, err := c.authorize()
		if err != nil {
			return err
		}
		log.Printf("[DEBUG] Freebox authentication done")
		return nil
	}
	_, err := c.login()
	if err != nil {
		return err
	}
	_, err = c.login()
	if err != nil {
		return err
	}
	log.Printf("[DEBUG] Freebox login done")
	if c.SessionToken == "" {
		_, err = c.openSession()
		if err != nil {
			return err
		}
	}
	log.Printf("[DEBUG] Freebox open session done")
	return err

}

func (c *Client) Statistics() (*providers.ProviderConnectionStatistics, error) {
	log.Printf("[DEBUG] Freebox retrieve statistics\n")
	resp, err := c.connectionStatus()
	if err != nil {
		apiError, err := makeAPIErrorResponse(err)
		if err != nil {
			return nil, err
		}
		if apiError.ErrorCode == authRequiredError {
			c.SessionToken = ""
			_, err := c.openSession()
			if err != nil {
				return nil, err
			}
			c.Statistics()
		}
		return nil, err
	}
	log.Printf("[DEBUG] Freebox connection status received")
	return &providers.ProviderConnectionStatistics{
		RateDown:      resp.Result.RateDown,
		RateUp:        resp.Result.RateUp,
		BytesDown:     resp.Result.BytesDown,
		BytesUp:       resp.Result.BytesUp,
		BandwidthDown: resp.Result.BandwidthDown,
		BandwidthUp:   resp.Result.BandwidthUp,
	}, nil
}
