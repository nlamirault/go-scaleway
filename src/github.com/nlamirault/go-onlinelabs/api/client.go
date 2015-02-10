// Copyright (C) 2015 Nicolas Lamirault <nicolas.lamirault@gmail.com>

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
	Client       *http.Client
	ComputeURL   string
	AccountURL   string
}

// NewClient creates a new OnlineLabs API client using userId, API token and organization
func NewClient(userid string, token string, organization string) *OnlineLabsClient {
	log.Debugf("Creating client using %s %s %s", userid, token, organization)
	client := &OnlineLabsClient{
		UserID:       userid,
		Token:        token,
		Organization: organization,
		Client:       &http.Client{},
		ComputeURL:   computeURL,
		AccountURL:   accountURL,
	}
	return client
}

// GetUserInformations list informations about your user account
func (c OnlineLabsClient) GetUserInformations(userID string) ([]byte, error) {
	body, err := c.getAccountAPIResource(fmt.Sprintf("users/%s", userID))
	if err != nil {
		return nil, err
	}
	log.Debugf("Get user response: %s", string(body))
	return body, nil
}

// GetUserOrganizations list all organizations associate with your account
func (c OnlineLabsClient) GetUserOrganizations() ([]byte, error) {
	body, err := c.getAccountAPIResource(fmt.Sprintf("organizations"))
	if err != nil {
		return nil, err
	}
	log.Debugf("Get user organizations response: %s", string(body))
	return body, nil
}

// GetUserTokens list all tokens associate with your account
func (c OnlineLabsClient) GetUserTokens() ([]byte, error) {
	body, err := c.getAccountAPIResource(fmt.Sprintf("tokens"))
	if err != nil {
		return nil, err
	}
	log.Debugf("Get tokens response: %s", string(body))
	return body, nil
}

//GetUserToken lList an individual Token
func (c OnlineLabsClient) GetUserToken(tokenID string) ([]byte, error) {
	body, err := c.getAccountAPIResource(fmt.Sprintf("tokens/%s", tokenID))
	if err != nil {
		return nil, err
	}
	log.Debugf("Get token response: %s", string(body))
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
	body, err := postAPIResource(
		c.Client,
		c.Token,
		fmt.Sprintf("%s/servers", c.ComputeURL),
		[]byte(json))
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
	body, err := deleteAPIResource(
		c.Client,
		c.Token,
		fmt.Sprintf("%s/servers/%s", c.ComputeURL, serverID))
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
	body, err := getAPIResource(
		c.Client,
		c.Token,
		fmt.Sprintf("%s/servers/%s", c.ComputeURL, serverID))
	if err != nil {
		return nil, err
	}
	log.Debugf("Get server response: %s", string(body))
	return body, nil
}

// GetServers list all servers associate with your account
func (c OnlineLabsClient) GetServers() ([]byte, error) {
	body, err := getAPIResource(
		c.Client,
		c.Token,
		fmt.Sprintf("%s/servers", c.ComputeURL))
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
	body, err := postAPIResource(
		c.Client,
		c.Token,
		fmt.Sprintf("%s/servers/%s/action", c.ComputeURL, serverID),
		[]byte(json))
	if err != nil {
		return nil, err
	}
	log.Debugf("[OnlineAPI] Server action response: %s",
		string(body))
	return body, nil
}

// GetVolume list an individual volume
// volumeID ith the volume unique identifier
func (c OnlineLabsClient) GetVolume(volumeID string) ([]byte, error) {
	body, err := getAPIResource(
		c.Client,
		c.Token,
		fmt.Sprintf("%s/volumes/%s", c.ComputeURL, volumeID))
	if err != nil {
		return nil, err
	}
	log.Debugf("Get volume response: %s", string(body))
	return body, nil
}

// DeleteVolume delete a specific volume
// volumeID ith the volume unique identifier
func (c OnlineLabsClient) DeleteVolume(volumeID string) ([]byte, error) {
	body, err := deleteAPIResource(
		c.Client,
		c.Token,
		fmt.Sprintf("%s/volumes/%s", c.ComputeURL, volumeID))
	if err != nil {
		return nil, err
	}
	log.Debugf("Delete volume response: %s",
		string(body))
	return body, nil
}

// CreateVolume creates a new volume
// name is the volume name
// organization is the organization unique identifier
// volume_type is the volume type
// size is the volume size
func (c OnlineLabsClient) CreateVolume(name string, organization string, volume_type string, size int) ([]byte, error) {
	json := fmt.Sprintf(`{"name": "%s", "organization": "%s", "volume_type": "%s", "size": %d}`,
		name, organization, volume_type, size)
	body, err := postAPIResource(
		c.Client,
		c.Token,
		fmt.Sprintf("%s/volumes", c.ComputeURL),
		[]byte(json))
	if err != nil {
		return nil, err
	}
	log.Debugf("Create volume response: %s",
		string(body))
	return body, nil
}

// GetVolumes list all volumes associate with your account
func (c OnlineLabsClient) GetVolumes() ([]byte, error) {
	body, err := getAPIResource(
		c.Client,
		c.Token,
		fmt.Sprintf("%s/volumes", c.ComputeURL))
	if err != nil {
		return nil, err
	}
	log.Debugf("Retrieve volumes response: %s",
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
	body, err := patchAPIResource(
		c.Client,
		c.Token,
		fmt.Sprintf("%s/users/%s", c.AccountURL, userid),
		[]byte(json))
	if err != nil {
		return nil, err
	}
	log.Debugf("[OnlineAPI] Server action response: %s",
		string(body))
	return body, nil
}

func (c OnlineLabsClient) getAccountAPIResource(request string) ([]byte, error) {
	url := fmt.Sprintf("%s/%s", c.AccountURL, request)
	log.Debugf("GET: %q", url)
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("X-Auth-Token", c.Token)
	req.Header.Set("Content-Type", "application/json")
	resp, err := c.Client.Do(req)
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
