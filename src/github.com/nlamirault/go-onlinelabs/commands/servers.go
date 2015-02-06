// Copyright (C) 2015  Nicolas Lamirault <nicolas.lamirault@gmail.com>

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package commands

import (
	//"fmt"

	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"

	"github.com/nlamirault/go-onlinelabs/api"
)

var commandListServers = cli.Command{
	Name:        "listServers",
	Usage:       "List all servers associate with your account",
	Description: ``,
	Action:      doListServers,
	Flags: []cli.Flag{
		verboseFlag,
	},
}

func doListServers(c *cli.Context) {
	log.Infof("List servers")
	client := api.NewClient(
		c.GlobalString("onlinelabs-userid"),
		c.GlobalString("onlinelabs-token"),
		c.GlobalString("onlinelabs-organization"))
	b, err := client.GetServers()
	if err != nil {
		log.Errorf("[Error] Retrieving servers %v", err)
		return
	}
	response, err := api.GetServersFromJson(b)
	if err != nil {
		log.Errorf("[Error] Reading servers %v", err)
		return
	}
	log.Infof("Servers: ")
	for _, server := range response.Servers {
		log.Infof("----------------------------------------------")
		log.Infof("Id:    %s", server.Id)
		log.Infof("Name:  %s", server.Name)
		log.Infof("Date:  %s", server.ModificationDate)
		log.Infof("IP:    %s", server.PublicIp.Address)
		log.Infof("Tags:  %s", server.Tags)
	}
	// b, err := client.GetRequest("/images")
	// if err != nil {
	// 	logging.Error("[Error] Get Images " + err)
	// 	return
	// }
	// print(string(b))
}
