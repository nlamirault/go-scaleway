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

package main

import (
	"fmt"
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"

	"github.com/nlamirault/go-onlinelabs/commands"
	//"github.com/nlamirault/go-onlinelabs/logging"
	"github.com/nlamirault/go-onlinelabs/version"
)

func makeApp() *cli.App {
	app := cli.NewApp()
	app.Name = "onlinelabs"
	app.Version = version.Version
	app.Usage = "A CLI for Online Labs"
	app.Author = "Nicolas Lamirault"
	app.Email = "nicolas.lamirault@gmail.com"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "log-level, l",
			Value: "info",
			Usage: fmt.Sprintf("Log level (options: debug, info, warn, error, fatal, panic)"),
		},
		cli.StringFlag{
			Name:   "onlinelabs-userid",
			Usage:  "Onlinelabs UserID",
			Value:  "",
			EnvVar: "ONLINELABS_USERID",
		},
		cli.StringFlag{
			Name:   "onlinelabs-token",
			Usage:  "Onlinelabs Token",
			Value:  "",
			EnvVar: "ONLINELABS_TOKEN",
		},
		cli.StringFlag{
			Name:   "onlinelabs-organization",
			Usage:  "Organization identifier",
			Value:  "",
			EnvVar: "ONLINELABS_ORGANIZATION",
		},
	}
	app.Before = func(c *cli.Context) error {
		//log.SetFormatter(&logging.CustomFormatter{})
		log.SetOutput(os.Stderr)
		level, err := log.ParseLevel(c.String("log-level"))
		if err != nil {
			log.Fatalf(err.Error())
		}
		log.SetLevel(level)
		return nil
	}

	app.Commands = commands.Commands
	//app.Flags = commands.Flags
	return app
}

func main() {
	app := makeApp()
	app.Run(os.Args)
}
