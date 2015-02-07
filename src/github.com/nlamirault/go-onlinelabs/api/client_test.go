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
	//"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

const (
	serverResponse = `{"server": {"tags": ["docker-machine"], "state_detail": "provisioning node", "image": {"default_bootscript": {"kernel": {"dtb": "dtb/pimouss-computing.dtb.3.17", "path": "kernel/pimouss-uImage-3.17-119-std", "id": "efff7963-2c2f-4467-837a-d14391218e36", "title": "Pimouss 3.17-119-std-with-aufs"}, "title": "NBD Boot - Linux 3.17 119-std", "public": true, "initrd": {"path": "initrd/pimouss-uInitrd", "id": "fe70e4dc-fb87-47e8-bf61-4c75c6f5a61e", "title": "pimouss-uInitrd"}, "bootcmdargs": {"id": "d22c4dde-e5a4-47ad-abb9-d23b54d542ff", "value": "ip=dhcp boot=local root=/dev/nbd0 USE_XNBD=1 nbd.max_parts=8"}, "organization": "11111111-1111-4111-8111-111111111111", "id": "d28611ff-08bd-4bdd-9f73-084a0e1ec9dc"}, "creation_date": "2015-01-19T18:08:41.906454+00:00", "name": "Debian Wheezy (7.8)", "modification_date": "2015-01-19T18:31:30.354525+00:00", "organization": "a283af0b-d13e-42e1-a43f-855ffbf281ab", "extra_volumes": "[]", "arch": "arm", "id": "cd66fa55-684a-4dd4-b809-956440b7a57f", "root_volume": {"size": 20000000000, "id": "7f98d217-e7ec-4ce0-87ab-a76e9615400c", "volume_type": "l_ssd", "name": "distrib-debian-wheezy-2015-01-19_19:01-snapshot"}, "public": true}, "creation_date": "2015-01-23T11:57:43.458120+00:00", "public_ip": {"dynamic": true, "id": "65c88668-577d-4fd5-aede-aa4bb2bb6427", "address": "212.47.230.241"}, "private_ip": "10.1.12.192", "id": "56e98092-6e05-4c89-9e76-b3610d38478c", "modification_date": "2015-01-23T11:57:43.458120+00:00", "name": "docker-lam", "dynamic_public_ip": true, "hostname": "docker-lam", "state": "starting", "bootscript": null, "volumes": {"0": {"size": 20000000000, "name": "distrib-debian-wheezy-2015-01-19_19:01-snapshot", "modification_date": "2015-01-23T11:54:02.800507+00:00", "organization": "19446e97-4a3b-4ccc-88f3-b65e3f31fb75", "export_uri": "nbd://10.1.12.199:4544", "creation_date": "2015-01-23T11:57:43.458120+00:00", "id": "aae55698-d4be-49e6-929b-bea9f6d7b09a", "volume_type": "l_ssd", "server": {"id": "56e98092-6e05-4c89-9e76-b3610d38478c", "name": "docker-lam"}}}, "organization": "19446e97-4a3b-4ccc-88f3-b65e3f31fb75"}} `
)

func httpClientWithProxy(server *httptest.Server) *http.Client {
	tr := &http.Transport{
		Proxy: func(req *http.Request) (*url.URL, error) {
			//fmt.Printf("Url: %v\n", server.URL)
			return url.Parse(server.URL)
		},
	}
	return &http.Client{Transport: tr}
}

func ServerHandler(res http.ResponseWriter, req *http.Request) {
	//fmt.Printf("Response to send : %s\n", serverResponse)
	//data, _ := json.Marshal(serverResponse)
	res.Header().Set("Content-Type", "application/json; charset=utf-8")
	// res.Write(data)
	fmt.Println(res, serverResponse)
}

func TestGettingServer(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(ServerHandler))
	defer server.Close()
	c := NewClient("02468", "02468", "13579")
	c.client = httpClientWithProxy(server)
	b, err := c.GetServer("56e98092-6e05-4c89-9e76-b3610d38478c")
	response, err := GetServerFromJson(b)
	fmt.Printf("Response: %v %v\n", response, err)
	// if err != nil {
	// 	fmt.Printf("Error : %v", err)
	// 	t.Errorf("Can't decode json: %v", err)
	// }
	// if response.Server.Id != "d28611ff-08bd-4bdd-9f73-084a0e1ec9dc" {
	// 	t.Errorf("Invalid server id")
	// }
}
