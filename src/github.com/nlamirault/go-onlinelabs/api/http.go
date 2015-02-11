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
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	log "github.com/Sirupsen/logrus"
)

func performAPIRequest(client *http.Client, req *http.Request, token string) ([]byte, error) {
	req.Header.Set("X-Auth-Token", token)
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	msg := string(b)
	log.Debugf("HTTP Response: [%d] %s", resp.StatusCode, msg)
	if resp.StatusCode > 299 {
		return nil, fmt.Errorf("[%d] %s",
			resp.StatusCode, msg)
	}
	return b, nil
}

func getAPIResource(client *http.Client, token string, url string) ([]byte, error) {
	log.Debugf("GET: %q", url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	return performAPIRequest(client, req, token)
}

func postAPIResource(client *http.Client, token string, url string, json []byte) ([]byte, error) {
	log.Debugf("POST: %q %s", url, string(json))
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(json))
	if err != nil {
		return nil, err
	}
	return performAPIRequest(client, req, token)
}

func deleteAPIResource(client *http.Client, token string, url string) ([]byte, error) {
	log.Debugf("DELETE: %q", url)
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}
	return performAPIRequest(client, req, token)
}

func patchAPIResource(client *http.Client, token string, url string, json []byte) ([]byte, error) {
	log.Debugf("PATCH: %q %s", url, string(json))
	req, err := http.NewRequest("PATCH", url, bytes.NewBuffer(json))
	if err != nil {
		return nil, err
	}
	return performAPIRequest(client, req, token)
}
