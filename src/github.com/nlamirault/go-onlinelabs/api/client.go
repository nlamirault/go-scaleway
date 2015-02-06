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
	ComputeURL = "https://api.cloud.online.net"
	AccountURL = "https://account.cloud.online.net"
)

type OnlineLabsClient struct {
	UserId       string
	Token        string
	Organization string
	client       *http.Client
}

func NewClient(userid string, token string, organization string) *OnlineLabsClient {

	log.Debugf("Creating client using %s %s %s", userid, token, organization)
	client := &OnlineLabsClient{
		UserId:       userid,
		Token:        token,
		Organization: organization,
		client:       &http.Client{},
	}
	return client
}

func (c OnlineLabsClient) CreateServer(name string, organization string,
	image string) ([]byte, error) {
	json := fmt.Sprintf(`{"name": "%s", "organization": "%s", "image": "%s", "tags": ["docker-machine"]}`,
		name, organization, image)
	body, err := c.postApiResource("servers", []byte(json))
	if err != nil {
		return nil, err
	}
	log.Debugf("Create server response: %s",
		string(body))
	return body, nil
}

func (c OnlineLabsClient) DeleteServer(serverId string) ([]byte, error) {
	body, err := c.deleteApiResource(fmt.Sprintf("servers/%s", serverId))
	if err != nil {
		return nil, err
	}
	log.Debugf("Delete server response: %s",
		string(body))
	return body, nil
}

func (c OnlineLabsClient) GetServer(serverId string) ([]byte, error) {
	body, err := c.getApiResource(fmt.Sprintf("servers/%s", serverId))
	if err != nil {
		return nil, err
	}
	log.Debugf("Retrieve server response: %s",
		string(body))
	return body, nil
}

func (c OnlineLabsClient) GetServers() ([]byte, error) {
	body, err := c.getApiResource("servers")
	if err != nil {
		return nil, err
	}
	log.Debugf("Retrieve servers response: %s",
		string(body))
	return body, nil
}

func (c OnlineLabsClient) PerformServerAction(serverId string, action string) ([]byte, error) {
	json := fmt.Sprintf(`{"action": "%s"}`, action)
	body, err := c.postApiResource(
		fmt.Sprintf("servers/%s/action", serverId),
		[]byte(json))
	if err != nil {
		return nil, err
	}
	log.Debugf("[OnlineAPI] Server action response: %s",
		string(body))
	return body, nil
}

func (c OnlineLabsClient) UploadPublicKey(userid string, keyPath string) ([]byte, error) {
	publicKey, err := ioutil.ReadFile(keyPath)
	if err != nil {
		return nil, err
	}
	json := fmt.Sprintf(`{"ssh_public_keys": [{"key": "%s"}]}`,
		strings.TrimSpace(string(publicKey)))
	body, err := c.patchApiResource(
		fmt.Sprintf("%s/users/%s", AccountURL, userid),
		[]byte(json))
	if err != nil {
		return nil, err
	}
	log.Debugf("[OnlineAPI] Server action response: %s",
		string(body))
	return body, nil
}

func (c OnlineLabsClient) getApiResource(request string) ([]byte, error) {
	url := fmt.Sprintf("%s/%s", ComputeURL, request)
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

func (c OnlineLabsClient) postApiResource(request string, json []byte) ([]byte, error) {
	url := fmt.Sprintf("%s/%s", ComputeURL, request)
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

func (c OnlineLabsClient) deleteApiResource(request string) ([]byte, error) {
	url := fmt.Sprintf("%s/%s", ComputeURL, request)
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

func (c OnlineLabsClient) patchApiResource(url string, json []byte) ([]byte, error) {
	//url := fmt.Sprintf("%s/%s", ComputeURL, request)
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
