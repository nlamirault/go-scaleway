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
	//"log"
	//"os"

	"github.com/codegangsta/cli"

	// "github.com/nlamirault/go-onlinelabs/api"
	"github.com/nlamirault/go-onlinelabs/logging"
)

var commandListImages = cli.Command{
	Name:        "listImages",
	Usage:       "List availables images",
	Description: ``,
	Action:      doListImages,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "name",
			Value: "",
			Usage: "name of the component",
		},
		cli.StringFlag{
			Name:  "version",
			Value: "",
			Usage: "version of the application",
		},
		verboseFlag,
	},
}

func doListImages(c *cli.Context) {
	logging.Info("List images")
	// client := api.NewClient()
	// b, err := client.GetRequest("/images")
	// if err != nil {
	// 	logging.Error("[Error] Get Images " + err)
	// 	return
	// }
	// print(string(b))
}
