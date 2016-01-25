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
	"log"

	"github.com/nlamirault/skybox/providers"
)

const (
	defaultURL = "http://192.168.1.1"
)

type apiAuthenticateResponse struct {
}

func (c *Client) authenticate() (*apiAuthenticateResponse, error) {
	log.Printf("[DEBUG] LiveboxAPI authenticate\n")
	var resp *apiAuthenticateResponse
	err := providers.Do(
		c,
		"POST",
		fmt.Sprintf("/authenticate?username=%s&password=%s",
			c.Username, c.Password),
		nil,
		&resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

type apiDisconnectResponse struct {
}

func (c *Client) disconnect() (*apiDisconnectResponse, error) {
	log.Printf("[DEBUG] LiveboxAPI Disconnect\n")
	var resp *apiDisconnectResponse
	err := providers.Do(
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

type apiMibDslResponse struct {
	LastChangeTime        int    `json:"LastChangeTime"`
	LastChange            int    `json:"LastChange"`
	LinkStatus            string `json:"LinkStatus"`
	UpstreamCurrRate      int    `json:"UpstreamCurrRate"`
	DownstreamCurrRate    int    `json:"DownstreamCurrRate"`
	UpstreamMaxRate       int    `json:"UpstreamMaxRate"`
	DownstreamMaxRate     int    `json:"DownstreamMaxRate"`
	UpstreamNoiseMargin   int    `json:"UpstreamNoiseMargin"`
	DownstreamNoiseMargin int    `json:"DownstreamNoiseMargin"`
	UpstreamAttenuation   int    `json:"UpstreamAttenuation"`
	DownstreamAttenuation int    `json:"DownstreamAttenuation"`
	UpstreamPower         int    `json:"UpstreamPower"`
	DownstreamPower       int    `json:"DownstreamPower"`
	DataPath              string `json:"DataPath"`
	InterleaveDepth       int    `json:"InterleaveDepth"`
	ModulationType        string `json:"ModulationType"`
	ModulationHint        string `json:"ModulationHint"`
	FirmwareVersion       string `json:"FirmwareVersion"`
	StandardsSupported    string `json:"StandardsSupported"`
	StandardUsed          string `json:"StandardUsed"`
	CurrentProfile        string `json:"CurrentProfile"`
	Upbokle               int    `json:"UPBOKLE"`
}

type apiConnectionStatusResponse struct {
	Result struct {
		Status struct {
			Dsl struct {
				Dls0 apiMibDslResponse
			}
		}
	}
}

func (c *Client) connectionStatus() (*apiConnectionStatusResponse, error) {
	log.Printf("[DEBUG] LiveboxAPI retrieve MIBs\n")
	var resp *apiConnectionStatusResponse
	err := providers.Do(
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

func (c *Client) getLiveboxAPIRequest(request string) string {
	return fmt.Sprintf("%s/sysbus/%s", c.Endpoint, request)
}
