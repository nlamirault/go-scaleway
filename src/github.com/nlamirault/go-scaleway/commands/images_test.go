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

package commands

// import (
// 	"flag"
// 	"fmt"
// 	"os"
// 	"testing"

// 	"github.com/codegangsta/cli"
// )

// func getTestContext() *cli.Context {
// 	app := cli.NewApp()
// 	set := flag.NewFlagSet("images", 0)
// 	set.String("scaleway-userid", os.Getenv("SCALEWAY_USERID"), "")
// 	set.String("scaleway-organization", os.Getenv("SCALEWAY_ORGANIZATION"), "")
// 	c := cli.NewContext(app, set, set)
// 	// fmt.Printf("Args: %v -- %s-- \n",
// 	// 	c.Args(), c.GlobalString("scaleway-userid"))
// 	return c
// }

// func TestCommandListImages(t *testing.T) {
// 	c := getTestContext()

// 	err := commandListImages.Run(c)
// 	if err != nil {
// 		fmt.Printf("List Image: %v", err.Error())
// 	}
// }
