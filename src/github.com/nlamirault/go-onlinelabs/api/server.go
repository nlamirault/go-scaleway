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
	"encoding/json"
)

type PublicIp struct {
	Id      string `json:"id,omitempty"`
	Dynamic bool   `json:"dynamic,omitempty"`
	Address string `json:"address,omitempty"`
}

type Server struct {
	Id               string   `json:"id,omitempty"`
	Name             string   `json:"name,omitempty"`
	Organization     string   `json:"organization,omitempty"`
	CreationDate     string   `json:"creation_date,omitempty"`
	ModificationDate string   `json:"modification_date,omitempty"`
	Arch             string   `json:"arch,omitempty"`
	PublicIp         PublicIp `json:"public_ip,omitempty"`
	State            string   `json:"state,omitempty"`
	Tags             []string `json:"tags,omitempty"`
}

type ServerResponse struct {
	Server Server `json:"server,omitempty"`
}

type ServersResponse struct {
	Servers []Server
}

func GetServerFromJson(b []byte) (*ServerResponse, error) {
	var response ServerResponse
	err := json.Unmarshal(b, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func GetServersFromJson(b []byte) (*ServersResponse, error) {
	var response ServersResponse
	err := json.Unmarshal(b, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
