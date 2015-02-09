// Copyright (C) 2015 Nicolas Lamirault <nicolas.lamirault@gmail.com>

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

var commandGetUser = cli.Command{
	Name:        "getUser",
	Usage:       "List informations about your user account",
	Description: ``,
	Action:      doListUserInformations,
	Flags: []cli.Flag{
		verboseFlag,
		cli.StringFlag{
			Name:  "userid",
			Usage: "User unique identifier",
			Value: "",
		},
	},
}

var commandGetOrganizations = cli.Command{
	Name:        "getOrganizations",
	Usage:       "List all Organizations associate with your account",
	Description: ``,
	Action:      doListUserOrganizations,
	Flags: []cli.Flag{
		verboseFlag,
	},
}

func doListUserInformations(c *cli.Context) {
	log.Infof("List user informations")
	client := getClient(c)
	b, err := client.GetUserInformations(c.String("userid"))
	if err != nil {
		log.Errorf("Failed user response %v", err)
		return
	}
	response, err := api.GetUserFromJSON(b)
	if err != nil {
		log.Errorf("Failed user informations %v", err)
		return
	}
	log.Infof("User: ")
	log.Infof("Id            : %s", response.User.ID)
	log.Infof("Fullname      : %s", response.User.Fullname)
	log.Infof("Firstname     : %s", response.User.Firstname)
	log.Infof("Lastname      : %s", response.User.Lastname)
	log.Infof("Email         : %s", response.User.Email)
	log.Infof("Phone         : %s", response.User.PhoneNumber)
	log.Infof("Roles         : %s", response.User.Roles)
	log.Infof("Organizations : %s", response.User.Organizations)
}

func doListUserOrganizations(c *cli.Context) {
	log.Infof("List user organizations")
	client := getClient(c)
	b, err := client.GetUserOrganizations()
	if err != nil {
		log.Errorf("Failed user organizations response %v", err)
		return
	}
	response, err := api.GetOrganizationsFromJSON(b)
	if err != nil {
		log.Errorf("Failed user organizations %v", err)
		return
	}
	log.Infof("User organizations:")
	for _, org := range response.Organizations {
		log.Infof("----------------------------------------------")
		log.Infof("Id              : %s", org.ID)
		log.Infof("Name            : %s", org.Name)
		log.Infof("Currency        : %s", org.Currency)
		log.Infof("Locale          : %s", org.Locale)
		log.Infof("Customer class  : %s", org.CustomerClass)
	}
}
