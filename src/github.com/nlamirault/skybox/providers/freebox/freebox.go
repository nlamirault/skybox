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
	"crypto/hmac"
	"crypto/sha1"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/nlamirault/skybox/config"
	"github.com/nlamirault/skybox/providers"
	"github.com/nlamirault/skybox/version"
)

const (
	DefaultURL = "http://mafreebox.free.fr/"
	APIVersion = "3"
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
	Identifier   string `json:"app_id",omitempty`
	Name         string `json:"app_name",omitempty`
	Version      string `json:"app_version",omitempty`
	DeviceName   string `json:"device_name",omitempty`
}

// New returns a Freebox Client
func New() *Client {
	baseURL, _ := url.Parse(DefaultURL)
	client := Client{
		Client:     &http.Client{},
		Endpoint:   baseURL,
		Identifier: "skybox",
		Name:       "Skybox",
		Version:    fmt.Sprintf("%s", version.Version),
		DeviceName: "Skybox",
	}
	if os.Getenv("SKYBOX_FREEBOX_TOKEN") != "" {
		client.Token = os.Getenv("SKYBOX_FREEBOX_TOKEN")
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
	var resp *APIVersionResponse
	err := providers.Do(
		c,
		"GET",
		fmt.Sprintf("%s/api_version", c.EndPoint()),
		nil,
		&resp)
	if err != nil {
		return err
	}
	log.Printf("[DEBUG] Freebox Ping response: %s", resp)
	return nil
}
func (c *Client) Authenticate() error {
	if c.Token == "" {
		err := c.authorize()
		if err != nil {
			return err
		}
		return nil
	}
	err := c.login()
	if err != nil {
		return err
	}
	return err

}
func (c *Client) authorize() error {
	log.Printf("[DEBUG] Freebox retrieve authorization\n")
	var resp *APIAuthorizeResponse
	err := providers.Do(
		c,
		"POST",
		c.getFreeboxAPIRequest("login/authorize"),
		APIAuthorizeRequest{
			AppID:      c.Identifier,
			AppName:    c.Name,
			AppVersion: c.Version,
			DeviceName: c.DeviceName,
		},
		&resp)
	if err != nil {
		return err
	}
	log.Printf("[DEBUG] Freebox Authorize response: %s", resp)
	return nil
}

func (c *Client) Statistics() (*providers.ProviderConnectionStatistics, error) {
	log.Printf("[DEBUG] Freebox retrieve statistics\n")
	var resp *APIConnectionStatusResponse
	err := providers.Do(
		c,
		"GET",
		c.getFreeboxAPIRequest("connection"),
		nil,
		&resp)
	if err != nil {
		return nil, err
	}
	log.Printf("[DEBUG] Freebox connection status response: %s", resp)
	return &providers.ProviderConnectionStatistics{
		RateDown:      resp.Result.RateDown,
		RateUp:        resp.Result.RateUp,
		BytesDown:     resp.Result.BytesDown,
		BytesUp:       resp.Result.BytesUp,
		BandwidthDown: resp.Result.BandwidthDown,
		BandwidthUp:   resp.Result.BandwidthUp,
	}, nil
}

func (c *Client) login() error {
	log.Printf("[DEBUG] Freebox login\n")
	if c.Token == "" {
		return fmt.Errorf("No Freebox token found.")
	}
	var resp *APILoginResponse
	err := providers.Do(
		c,
		"GET",
		c.getFreeboxAPIRequest("login"),
		nil,
		&resp)
	if err != nil {
		return err
	}
	log.Printf("[DEBUG] Freebox login response: %s", resp)
	c.Challenge = resp.Result.Challenge
	if c.SessionToken == "" {
		c.openSession()
	}
	return nil
}

func (c *Client) openSession() error {
	log.Printf("[DEBUG] Freebox open session\n")
	hash := hmac.New(sha1.New, []byte(c.Token))
	hash.Write([]byte(c.Challenge))
	c.Password = fmt.Sprintf("%x", hash.Sum(nil))
	var resp *APILoginResponse
	err := providers.Do(
		c,
		"POST",
		c.getFreeboxAPIRequest("login/session"),
		APILoginRequest{
			AppID:      c.Identifier,
			AppVersion: c.Version,
			Password:   c.Password,
		},
		&resp)
	if err != nil {
		return err
	}
	c.SessionToken = resp.Result.SessionToken
	return nil
}

func (c *Client) getFreeboxAPIRequest(request string) string {
	return fmt.Sprintf("%s/api/v%s/%s", c.Endpoint, APIVersion, request)
}
