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

// Orange Livebox API

// http://192.168.1.1/sysbus/Wificom:getStatus

// http://192.168.1.1/sysbus/NeMo/Intf/dsl0:getDSLStats
// http://192.168.1.1/sysbus/NeMo/Intf/data:getMIBs
// http://192.168.1.1/sysbus/NeMo/Intf/lan:getMIBs
// http://192.168.1.1/sysbus/NeMo/Intf/lan:luckyAddrAddress
// http://192.168.1.1/sysbus/NeMo/Intf/data:luckyAddrAddress

// http://192.168.1.1/sysbus/UserManagement:getUsers

// http://192.168.1.1/sysbus/NMC:getWANStatus : status wan
// http://192.168.1.1/sysbus/NMC/OrangeTV:getIPTVStatus : status TV

// http://192.168.1.1/sysbus/VoiceService/VoiceApplication:listTrunks

package livebox

import (
	"fmt"
	"log"
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
	Client    *http.Client
	Endpoint  *url.URL
	ContextID string
	Username  string
	Password  string
	Cookies   []*http.Cookie
}

// New returns a Livebox Client
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
	if c.Cookies != nil {
		for _, cookie := range c.Cookies {
			request.AddCookie(cookie)
		}
		request.Header.Add("X-Context", c.ContextID)
		request.Header.Add("X-Sah-Request-Type", "idle")
		request.Header.Add("X-Requested-With", "XMLHttpRequest")
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
	c.Authenticate()
	resp, err := c.getTime()
	if err != nil {
		return err
	}
	log.Printf("[DEBUG] Livebox Time: %s", resp.Result.Data.Time)
	return nil
}

func (c *Client) Authenticate() error {
	if len(c.Cookies) == 0 && len(c.ContextID) == 0 {
		_, err := c.authenticate()
		if err != nil {
			return err
		}
		log.Printf("[DEBUG] Livebox authentication done")
	}
	return nil
}

func (c *Client) Statistics() (*providers.ProviderConnectionStatistics, error) {
	_, err := c.connectionStatus()
	if err != nil {
		return nil, err
	}
	return nil, nil
}
