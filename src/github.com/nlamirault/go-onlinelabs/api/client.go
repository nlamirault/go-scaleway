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
	body, err := getAPIResource(
		c.Client,
		c.Token,
		fmt.Sprintf("%s/users/%s", c.AccountURL, userID))
	if err != nil {
		return nil, err
	}
	log.Debugf("Get user response: %s", string(body))
	return body, nil
}

// GetUserOrganizations list all organizations associate with your account
func (c OnlineLabsClient) GetUserOrganizations() ([]byte, error) {
	body, err := getAPIResource(
		c.Client,
		c.Token,
		fmt.Sprintf("%s/organizations", c.AccountURL))
	if err != nil {
		return nil, err
	}
	log.Debugf("Get user organizations response: %s", string(body))
	return body, nil
}

// GetUserTokens list all tokens associate with your account
func (c OnlineLabsClient) GetUserTokens() ([]byte, error) {
	body, err := getAPIResource(
		c.Client,
		c.Token,
		fmt.Sprintf("%s/tokens", c.AccountURL))
	if err != nil {
		return nil, err
	}
	log.Debugf("Get tokens response: %s", string(body))
	return body, nil
}

//GetUserToken lList an individual Token
func (c OnlineLabsClient) GetUserToken(tokenID string) ([]byte, error) {
	body, err := getAPIResource(
		c.Client,
		c.Token,
		fmt.Sprintf("%s/tokens/%s", c.AccountURL, tokenID))
	if err != nil {
		return nil, err
	}
	log.Debugf("Get token response: %s", string(body))
	return body, nil
}

// CreateToken authenticates a user against their email, password,
// and then returns a new Token, which can be used until it expires.
// email is the user email
// password is the user password
// expires is if you want a token wich expires or not
func (c OnlineLabsClient) CreateToken(email string, password string,
	expires bool) ([]byte, error) {
	json := fmt.Sprintf(`{"email": "%s", "password": "%s", "expires": %t}`,
		email, password, expires)
	body, err := postAPIResource(
		c.Client,
		c.Token,
		fmt.Sprintf("%s/tokens", c.AccountURL),
		[]byte(json))
	if err != nil {
		return nil, err
	}
	log.Debugf("Create token response: %s",
		string(body))
	return body, nil
}

// DeleteToken delete a specific token
// tokenID ith the token unique identifier
func (c OnlineLabsClient) DeleteToken(tokenID string) ([]byte, error) {
	body, err := deleteAPIResource(
		c.Client,
		c.Token,
		fmt.Sprintf("%s/tokens/%s", c.AccountURL, tokenID))
	if err != nil {
		return nil, err
	}
	log.Debugf("Delete token response: %s", string(body))
	return body, nil
}

// UpdateToken increase Token expiration time of 30 minutes
// tokenID ith the token unique identifier
func (c OnlineLabsClient) UpdateToken(tokenID string) ([]byte, error) {
	json := `{"expires": true}`
	body, err := patchAPIResource(
		c.Client,
		c.Token,
		fmt.Sprintf("%s/tokens/%s", c.AccountURL, tokenID),
		[]byte(json))
	if err != nil {
		return nil, err
	}
	log.Debugf("Update token response: %s", string(body))
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
	log.Debugf("Delete volume response: %s", string(body))
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
	log.Debugf("Create volume response: %s", string(body))
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

// GetImages list all images associate with your account
func (c OnlineLabsClient) GetImages() ([]byte, error) {
	body, err := getAPIResource(
		c.Client,
		c.Token,
		fmt.Sprintf("%s/images", c.ComputeURL))
	if err != nil {
		return nil, err
	}
	log.Debugf("Retrieve images response: %s", string(body))
	return body, nil
}

// GetImage list an individual image
// volumeID ith the image unique identifier
func (c OnlineLabsClient) GetImage(volumeID string) ([]byte, error) {
	body, err := getAPIResource(
		c.Client,
		c.Token,
		fmt.Sprintf("%s/images/%s", c.ComputeURL, volumeID))
	if err != nil {
		return nil, err
	}
	log.Debugf("Get image response: %s", string(body))
	return body, nil
}

// DeleteImage delete a specific volume
// volumeID ith the volume unique identifier
func (c OnlineLabsClient) DeleteImage(imageID string) ([]byte, error) {
	body, err := deleteAPIResource(
		c.Client,
		c.Token,
		fmt.Sprintf("%s/images/%s", c.ComputeURL, imageID))
	if err != nil {
		return nil, err
	}
	log.Debugf("Delete image response: %s", string(body))
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
