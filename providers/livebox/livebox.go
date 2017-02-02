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

package livebox

import (
	"fmt"
	//"log"
	"net/http"
	"net/url"

	"github.com/Sirupsen/logrus"

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
	_, err := c.authenticate()
	if err != nil {
		return err
	}
	resp, err := c.getTime()
	if err != nil {
		return err
	}
	logrus.Debugf("Livebox Time: %s", resp.Result.Data.Time)
	return nil
}

func (c *Client) Authenticate() error {
	if len(c.Cookies) == 0 && len(c.ContextID) == 0 {
		_, err := c.authenticate()
		if err != nil {
			return err
		}
		logrus.Debugf("Livebox authentication done")
	}
	return nil
}

func (c *Client) Statistics() (*providers.ConnectionStatistics, error) {
	if err := c.Authenticate(); err != nil {
		return nil, err
	}
	conStatus, err := c.connectionStatus()
	if err != nil {
		return nil, err
	}
	logrus.Debugf("Statistics: %s", conStatus)
	return &providers.ConnectionStatistics{}, nil
}

func (c *Client) Network() (*providers.NetworkInformations, error) {
	if err := c.Authenticate(); err != nil {
		return nil, err
	}
	wanStatus, err := c.wanStatus()
	if err != nil {
		return nil, err
	}
	logrus.Debugf("Wan: %s", wanStatus)
	return &providers.NetworkInformations{
		IPV4Address: wanStatus.Result.Data.IPAddress,
		IPV6Address: wanStatus.Result.Data.IPv6Address,
		DNS:         wanStatus.Result.Data.DNSServers,
		State:       wanStatus.Result.Data.LinkState,
	}, nil
}

func (c *Client) Wifi() (*providers.WifiStatus, error) {
	if err := c.Authenticate(); err != nil {
		return nil, err
	}
	wifiStatus, err := c.wifiStatus()
	if err != nil {
		return nil, err
	}
	logrus.Debugf("Wifi: %s", wifiStatus)
	return &providers.WifiStatus{
		State:  wifiStatus.Result.Status.Status,
		Enable: wifiStatus.Result.Status.Enable,
	}, nil
}

func (c *Client) Devices() ([]*providers.BoxDevice, error) {
	if err := c.Authenticate(); err != nil {
		return nil, err
	}
	devices, err := c.devices()
	if err != nil {
		return nil, err
	}
	logrus.Debugf("Devices: %s", devices)
	thedevices := []*providers.BoxDevice{}
	for _, dev := range devices.Result.Status {
		if dev.Active && len(dev.IPAddress) > 0 {
			devType := "wifi"
			if dev.Layer2Interface == "eth0" {
				devType = "ethernet"
			}
			thedevices = append(thedevices, &providers.BoxDevice{
				Name:      dev.Name,
				Type:      devType,
				IPAddress: dev.IPAddress,
			})
		}
	}
	return thedevices, nil
}
