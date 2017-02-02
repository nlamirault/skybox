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
	"log"

	"github.com/nlamirault/skybox/providers"
)

const (
	defaultURL = "http://192.168.1.1"
)

func (c *Client) authenticate() (*apiAuthenticateResponse, error) {
	log.Printf("[DEBUG] LiveboxAPI authenticate\n")
	var resp *apiAuthenticateResponse
	cookies, err := providers.Do(
		c,
		"POST",
		fmt.Sprintf("/authenticate?username=%s&password=%s",
			c.Username, c.Password),
		nil,
		&resp)
	if err != nil {
		return nil, err
	}
	if len(cookies) == 0 {
		return nil, fmt.Errorf("Can't read cookie from Livebox")
	}
	c.Cookies = cookies
	c.ContextID = resp.Data.ContextID
	log.Printf("[DEBUG] ContextID: %s", c.ContextID)
	return resp, nil
}

func (c *Client) disconnect() (*apiDisconnectResponse, error) {
	log.Printf("[DEBUG] LiveboxAPI Disconnect\n")
	var resp *apiDisconnectResponse
	_, err := providers.Do(
		c,
		"POST",
		fmt.Sprintf("/logout"),
		nil,
		&resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *Client) connectionStatus() (*apiConnectionStatusResponse, error) {
	log.Printf("[DEBUG] LiveboxAPI retrieve MIBs\n")
	var resp *apiConnectionStatusResponse
	_, err := providers.Do(
		c,
		"POST",
		c.getLiveboxAPIRequest("NeMo/Intf/data:getMIBs"),
		nil,
		&resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *Client) wanStatus() (*apiWanStatusResponse, error) {
	log.Printf("[DEBUG] LiveboxAPI retrieve wan informations")
	var resp *apiWanStatusResponse
	_, err := providers.Do(
		c,
		"POST",
		c.getLiveboxAPIRequest("NMC:getWANStatus"),
		nil,
		&resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *Client) wifiStatus() (*apiWifiStatusResponse, error) {
	log.Printf("[DEBUG] LiveboxAPI retrieve wifi informations")
	var resp *apiWifiStatusResponse
	_, err := providers.Do(
		c,
		"POST",
		c.getLiveboxAPIRequest("NMC/Wifi:get"),
		nil,
		&resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *Client) tvStatus() (*apiTVStatusResponse, error) {
	log.Printf("[DEBUG] LiveboxAPI retrieve device informations")
	var resp *apiTVStatusResponse
	_, err := providers.Do(
		c,
		"POST",
		c.getLiveboxAPIRequest("NMC/OrangeTV:getIPTVConfig"),
		nil,
		&resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *Client) devices() (*apiDevicesResponse, error) {
	log.Printf("[DEBUG] LiveboxAPI retrieve connected devices")
	var resp *apiDevicesResponse
	_, err := providers.Do(
		c,
		"POST",
		c.getLiveboxAPIRequest("Devices:get"),
		nil,
		&resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *Client) getTime() (*apiTimeResponse, error) {
	log.Printf("[DEBUG] LiveboxAPI Get time\n")
	var resp *apiTimeResponse
	_, err := providers.Do(
		c,
		"POST",
		c.getLiveboxAPIRequest("Time:getTime"),
		nil,
		&resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *Client) getLiveboxAPIRequest(request string) string {
	return fmt.Sprintf("%s/sysbus/%s", c.Endpoint, request)
}
