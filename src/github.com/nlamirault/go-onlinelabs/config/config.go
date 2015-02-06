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

package config

import (
	"encoding/json"
	// "fmt"
	"io/ioutil"
	//"os"

	"github.com/nlamirault/go-onlinelabs/logging"
	//"inky/utils"
)

// Config represents the Inky configuration
type Config struct {
	Username string `json:"user"`
	Apikey   string `json:"apikey"`
}

// LoadConfiguration read Inky configuration and return a Config struct
func LoadConfiguration(path string) (config *Config, err error) {
	//fmt.Println(chalk.Green, "Loading configuration", chalk.Reset)
	//filename := utils.UserHomeDir() + "/" + path
	logging.Debug("Loading configuration")
	content, err := ioutil.ReadFile(path)
	if err != nil {
		logging.Error("File error: " + path)
		return nil, err
	}
	//fmt.Printf("%s\n", string(content))

	var conf Config
	err = json.Unmarshal(content, &conf)
	if err != nil {
		logging.Error("Load configuration failed " + err.Error())
		return nil, err
	}
	//fmt.Println(conf)

	logging.Debug("User " + conf.Username + " ApiKey : " + conf.Apikey)
	return &conf, nil
}
