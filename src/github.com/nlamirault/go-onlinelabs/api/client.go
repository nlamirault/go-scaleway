// Copyright (C) 2015  Nicolas Lamirault <nicolas.lamirault@gmail.com>

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package api

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	log "github.com/Sirupsen/logrus"
)

const (
	computeURL = "https://api.cloud.online.net"
	accountURL = "https://account.cloud.online.net"
)

// OnlineLabsClient is a client for the Online Labs Cloud API.
// UserID represents your user identifiant
// Token is to authenticate to the API
// Organization is the ID of the user's organization
type OnlineLabsClient struct {
	UserID       string
	Token        string
	Organization string
	client       *http.Client
}

// NewClient creates a new OnlineLabs API client using userId, API token and organization
func NewClient(userid string, token string, organization string) *OnlineLabsClient {

	log.Debugf("Creating client using %s %s %s", userid, token, organization)
	client := &OnlineLabsClient{
		UserID:       userid,
		Token:        token,
		Organization: organization,
		client:       &http.Client{},
	}
	return client
}

func (c OnlineLabsClient) GetUserInformations(userID string) ([]byte, error) {
	body, err := c.getAccountAPIResource(fmt.Sprintf("users/%s", userID))
	if err != nil {
		return nil, err
	}
	log.Debugf("Get user response: %s", string(body))
	return body, nil
}

func (c OnlineLabsClient) GetUserOrganzations() ([]byte, error) {
	body, err := c.getAccountAPIResource(fmt.Sprintf("organizations"))
	if err != nil {
		return nil, err
	}
	log.Debugf("Get user organizations response: %s", string(body))
	return body, nil
}

// CreateServer creates a new server
// name is the server name
// organization is the organization unique identifier
// image is the image unique identifier
func (c OnlineLabsClient) CreateServer(name string, organization string,
	image string) ([]byte, error) {
	json := fmt.Sprintf(`{"name": "%s", "organization": "%s", "image": "%s", "tags": ["docker-machine"]}`,
		name, organization, image)
	body, err := c.postAPIResource("servers", []byte(json))
	if err != nil {
		return nil, err
	}
	log.Debugf("Create server response: %s",
		string(body))
	return body, nil
}

// DeleteServer delete a specific server
// serverID ith the server unique identifier
func (c OnlineLabsClient) DeleteServer(serverID string) ([]byte, error) {
	body, err := c.deleteAPIResource(fmt.Sprintf("servers/%s", serverID))
	if err != nil {
		return nil, err
	}
	log.Debugf("Delete server response: %s",
		string(body))
	return body, nil
}

// GetServer list an individual server
// serverID ith the server unique identifier
func (c OnlineLabsClient) GetServer(serverID string) ([]byte, error) {
	body, err := c.getAPIResource(fmt.Sprintf("servers/%s", serverID))
	if err != nil {
		return nil, err
	}
	log.Debugf("Get server response: %s", string(body))
	return body, nil
}

// GetServers list all servers associate with your account
func (c OnlineLabsClient) GetServers() ([]byte, error) {
	body, err := c.getAPIResource("servers")
	if err != nil {
		return nil, err
	}
	log.Debugf("Retrieve servers response: %s",
		string(body))
	return body, nil
}

// PerformServerAction execute an action on a server
// serverID ith the server unique identifier
// action is the action to execute
func (c OnlineLabsClient) PerformServerAction(serverID string, action string) ([]byte, error) {
	json := fmt.Sprintf(`{"action": "%s"}`, action)
	body, err := c.postAPIResource(
		fmt.Sprintf("servers/%s/action", serverID),
		[]byte(json))
	if err != nil {
		return nil, err
	}
	log.Debugf("[OnlineAPI] Server action response: %s",
		string(body))
	return body, nil
}

// UploadPublicKey update user SSH keys
// userId is the user unique identifier
// keyPath is the complete path of the SSH key
func (c OnlineLabsClient) UploadPublicKey(userid string, keyPath string) ([]byte, error) {
	publicKey, err := ioutil.ReadFile(keyPath)
	if err != nil {
		return nil, err
	}
	json := fmt.Sprintf(`{"ssh_public_keys": [{"key": "%s"}]}`,
		strings.TrimSpace(string(publicKey)))
	body, err := c.patchAPIResource(
		fmt.Sprintf("%s/users/%s", accountURL, userid),
		[]byte(json))
	if err != nil {
		return nil, err
	}
	log.Debugf("[OnlineAPI] Server action response: %s",
		string(body))
	return body, nil
}

func (c OnlineLabsClient) getAccountAPIResource(request string) ([]byte, error) {
	url := fmt.Sprintf("%s/%s", accountURL, request)
	log.Debugf("GET: %q", url)
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("X-Auth-Token", c.Token)
	req.Header.Set("Content-Type", "application/json")
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if resp.StatusCode > 299 {
		return nil, fmt.Errorf("Status code: %d", resp.StatusCode)
	}
	return b, nil
}

func (c OnlineLabsClient) getAPIResource(request string) ([]byte, error) {
	url := fmt.Sprintf("%s/%s", computeURL, request)
	log.Debugf("GET: %q", url)
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("X-Auth-Token", c.Token)
	req.Header.Set("Content-Type", "application/json")
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if resp.StatusCode > 299 {
		return nil, fmt.Errorf("Status code: %d", resp.StatusCode)
	}
	return b, nil
}

func (c OnlineLabsClient) postAPIResource(request string, json []byte) ([]byte, error) {
	url := fmt.Sprintf("%s/%s", computeURL, request)
	log.Debugf("POST: %q %s", url, string(json))
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(json))
	req.Header.Set("X-Auth-Token", c.Token)
	req.Header.Set("Content-Type", "application/json")
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode > 299 {
		return nil, fmt.Errorf("%d %s",
			resp.StatusCode, string(b))
	}
	//return ioutil.ReadAll(resp.Body)
	return b, nil
}

func (c OnlineLabsClient) deleteAPIResource(request string) ([]byte, error) {
	url := fmt.Sprintf("%s/%s", computeURL, request)
	log.Debugf("DELETE: %q", url)
	req, err := http.NewRequest("DELETE", url, nil)
	req.Header.Set("X-Auth-Token", c.Token)
	req.Header.Set("Content-Type", "application/json")
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if resp.StatusCode > 299 {
		return nil, fmt.Errorf("Status code: %d", resp.StatusCode)
	}
	return b, nil
}

func (c OnlineLabsClient) patchAPIResource(url string, json []byte) ([]byte, error) {
	//url := fmt.Sprintf("%s/%s", computeURL, request)
	log.Debugf("PATCH: %q %s", url, string(json))
	req, err := http.NewRequest("PATCH", url, bytes.NewBuffer(json))
	req.Header.Set("X-Auth-Token", c.Token)
	req.Header.Set("Content-Type", "application/json")
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode > 299 {
		return nil, fmt.Errorf("%d %s",
			resp.StatusCode, string(b))
	}
	//return ioutil.ReadAll(resp.Body)
	return b, nil
}
