// Copyright (C) 2016, 2017 Nicolas Lamirault <nicolas.lamirault@gmail.com>

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"

	"github.com/urfave/cli"

	"github.com/nlamirault/skybox/config"
)

// CheckCommand is the command which check providers
var CheckCommand = cli.Command{
	Name: "check",
	Subcommands: []cli.Command{
		boxCommand,
	},
}

var boxCommand = cli.Command{
	Name:  "box",
	Usage: "Check box provider",
	Action: func(context *cli.Context) error {
		configFile, err := getConfigurationFile()
		if err != nil {
			return err
		}
		conf, agent, err := setup(configFile)
		if err != nil {
			return err
		}
		return checkBoxProvider(agent, conf)
	},
}

func checkBoxProvider(agent *Agent, conf *config.Configuration) error {
	fmt.Printf("Check box provider: %s\n", agent.Provider.Description())
	if err := agent.Provider.Setup(conf); err != nil {
		return err
	}
	if err := agent.Provider.Ping(); err != nil {
		return err
	}
	if err := agent.Provider.Authenticate(); err != nil {
		return err
	}
	fmt.Printf("Box provider successfully configured\n")
	return nil
}
