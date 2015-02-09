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
	//"fmt"
)

// Organization represents an Online Labs entity
type Organization struct {
	ID            string `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	Currency      string `json:"currency,omitempty"`
	Locale        string `json:"locale,omitempty"`
	CustomerClass string `json:"customer_class,omitempty"`
}

// OrganizationsResponse represents JSON response of server
type OrganizationsResponse struct {
	Organizations []Organization
}

// Role represents role of an user into an organization
type Role struct {
	Organization Organization
	Role         string `json:"role,omitempty"`
}

// User represents a user account of the Online Labs cloud
type User struct {
	ID            string         `json:"id,omitempty"`
	Fullname      string         `json:"fullname,omitempty"`
	Firstname     string         `json:"firstname,omitempty"`
	Lastname      string         `json:"lastname,omitempty"`
	Email         string         `json:"email,omitempty"`
	PhoneNumber   string         `json:"phone_number,omitempty"`
	Organizations []Organization `json:"organizations,omitempty"`
	Roles         []Role         `json:"roles,omitempty"`
}

// UserResponse represents JSON response of server
type UserResponse struct {
	User User `json:"user,omitempty"`
}

// GetUserFromJSON load bytes and return a UserResponse
func GetUserFromJSON(b []byte) (*UserResponse, error) {
	var response UserResponse
	// fmt.Printf("Response JSON: %s\n", (string(b)))
	err := json.Unmarshal(b, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

// GetUserFromJSON load bytes and return a UserResponse
func GetOrganizationsFromJSON(b []byte) (*OrganizationsResponse, error) {
	var response OrganizationsResponse
	// fmt.Printf("Response JSON: %s\n", (string(b)))
	err := json.Unmarshal(b, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
