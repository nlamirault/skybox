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

package freebox

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/json"
	"fmt"
	// "log"

	"github.com/sirupsen/logrus"

	"github.com/nlamirault/skybox/providers"
)

const (
	defaultURL = "http://mafreebox.free.fr"

	apiVersion = "3"

	// API Errors code

	// Invalid session token, or not session token sent
	authRequiredError string = "auth_required"
	// The app token you are trying to use is invalid or has been revoked
	invalidToken string = "invalid_token"
	// The app token you are trying to use has not been validated by user yet
	pendingToken string = "pending_token"
	// Your app permissions does not allow accessing this API
	insufficientRights string = "insufficient_rights"
	// You are trying to get an app_token from a remote IP
	deniedFromExternalIP string = "denied_from_external_ip"
	// Your request is invalid
	invalidRequest string = "invalid_request"
	// Too many auth error have been made from your IP
	rateLimited string = "ratelimited"
	// New application token request has been disabled
	newAppsDenied string = "new_apps_denied"
	// API access from apps has been disabled
	appsDenied string = "apps_denied"
	// Internal error
	internalError string = "internal_error"
)

type apiErrorResponse struct {
	UID       string `json:"uid"`
	Message   string `json:"msg"`
	Success   bool   `json:"success"`
	ErrorCode string `json:"error_code"`
}

// apiVersionResponse is returned by requesting `GET /api_version`
type apiVersionResponse struct {
	FreeboxID  string `json:"uid"`
	DeviceName string `json:"device_name"`
	Version    string `json:"api_version"`
	BaseURL    string `json:"api_base_url"`
	DeviceType string `json:"device_type"`
}

func (c *Client) version() (*apiVersionResponse, error) {
	var resp *apiVersionResponse
	_, err := providers.Do(
		c,
		"GET",
		fmt.Sprintf("%s/api_version", c.EndPoint()),
		nil,
		&resp)
	logrus.Debugf("FreeboxAPI version response: %v", resp)
	return resp, err
}

// apiAuthorizeRequest is sent by requesting `POST /api/v3/login/authorize/`
type apiAuthorizeRequest struct {
	AppID      string `json:"app_id"`
	AppName    string `json:"app_name"`
	AppVersion string `json:"app_version"`
	DeviceName string `json:"device_name"`
}

// apiAuthorizeResponse is returned by requesting `POST /api/v3/login/authorize/`
type apiAuthorizeResponse struct {
	Success bool `json:"success"`
	Result  struct {
		AppToken string `json:"app_token"`
		TrackID  int    `json:"track_id"`
	}
}

func (c *Client) authorize() (*apiAuthorizeResponse, error) {
	logrus.Debugf("FreeboxAPI retrieve authorization\n")
	var resp *apiAuthorizeResponse
	_, err := providers.Do(
		c,
		"POST",
		c.getFreeboxAPIRequest("login/authorize"),
		apiAuthorizeRequest{
			AppID:      c.Identifier,
			AppName:    c.Name,
			AppVersion: c.Version,
			DeviceName: c.DeviceName,
		},
		&resp)
	if err != nil {
		return nil, err
	}
	logrus.Debugf("FreeboxAPI Authorize response: %v", resp)
	return resp, nil
}

// apiConnectionStatusResponse is returned by requesting `GET /api/v3/connection/`
type apiConnectionStatusResponse struct {
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

func (c *Client) connectionStatus() (*apiConnectionStatusResponse, error) {
	logrus.Debugf("FreeboxAPI connection status\n")
	var resp *apiConnectionStatusResponse
	_, err := providers.Do(
		c,
		"GET",
		c.getFreeboxAPIRequest("connection"),
		nil,
		&resp)
	if err != nil {
		return nil, err
	}
	logrus.Debugf("FreeboxAPI connection status response: %v", resp)
	return resp, nil
}

type apiLoginResponse struct {
	Success bool `json:"success"`
	Result  struct {
		Challenge string `json:"challenge"`
		LoggedIn  bool   `json:"logged_in"`
	} `json:"result"`
}

func (c *Client) login() (*apiLoginResponse, error) {
	logrus.Debugf("FreeboxAPI login\n")
	var resp *apiLoginResponse
	_, err := providers.Do(
		c,
		"GET",
		c.getFreeboxAPIRequest("login"),
		nil,
		&resp)
	if err != nil {
		return nil, err
	}
	c.Challenge = resp.Result.Challenge
	logrus.Debugf("FreeboxAPI login response: %v", resp)
	return resp, nil
}

type apiLoginSessionRequest struct {
	AppID      string `json:"app_id"`
	AppVersion string `json:"app_version"`
	Password   string `json:"password"`
}

type apiLoginSessionResponse struct {
	Success bool `json:"success"`
	Result  struct {
		SessionToken string `json:"session_token"`
		Challenge    string `json:"challenge"`
		PasswordSalt string `json:""`
		Permissions  struct {
			Settings   bool `json:"settings"`
			Contacts   bool `json:"contacts"`
			Calls      bool `json:"calls"`
			Explorer   bool `json:"explorer"`
			Downloader bool `json:"downloader"`
			Parental   bool `json:"parental"`
			Pvr        bool `json:"pvr"`
		}
	} `json:"result"`
}

func (c *Client) openSession() (*apiLoginSessionResponse, error) {
	logrus.Debugf("FreeboxAPI open session\n")
	hash := hmac.New(sha1.New, []byte(c.Token))
	hash.Write([]byte(c.Challenge))
	c.Password = fmt.Sprintf("%x", hash.Sum(nil))
	var resp *apiLoginSessionResponse
	_, err := providers.Do(
		c,
		"POST",
		c.getFreeboxAPIRequest("login/session"),
		apiLoginSessionRequest{
			AppID:      c.Identifier,
			AppVersion: c.Version,
			Password:   c.Password,
		},
		&resp)
	if err != nil {
		return nil, err
	}
	logrus.Debugf("FreeboxAPI open session response: %v", resp)
	c.SessionToken = resp.Result.SessionToken
	return resp, nil
}

type apiLogoutSessionResponse struct {
	Success bool `json:"success"`
}

func (c *Client) closeSession() (*apiLogoutSessionResponse, error) {
	var resp *apiLogoutSessionResponse
	_, err := providers.Do(
		c,
		"POST",
		fmt.Sprintf("%s/login/logout", c.EndPoint()),
		nil,
		&resp)
	if err != nil {
		return nil, err
	}
	logrus.Debugf("FreeboxAPI close session response: %v", resp)
	c.SessionToken = ""
	return resp, err
}

func (c *Client) getFreeboxAPIRequest(request string) string {
	return fmt.Sprintf("%s/api/v%s/%s", c.Endpoint, apiVersion, request)
}

func makeAPIErrorResponse(e error) (*apiErrorResponse, error) {
	var resp *apiErrorResponse
	if err := json.Unmarshal([]byte(e.(*providers.APIError).Message), &resp); err != nil {
		return nil, err
	}
	return resp, nil
}
