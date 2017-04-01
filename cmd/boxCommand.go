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

// BoxCommand is the command which display informations about box provider
var BoxCommand = cli.Command{
	Name: "box",
	Subcommands: []cli.Command{
		infosCommand,
	},
}

var infosCommand = cli.Command{
	Name:  "infos",
	Usage: "Display box provider informations",
	Action: func(context *cli.Context) error {
		configFile, err := getConfigurationFile()
		if err != nil {
			return err
		}
		conf, agent, err := setup(configFile)
		if err != nil {
			return err
		}
		return boxProviderInformations(agent, conf)
	},
}

func boxProviderInformations(agent *Agent, conf *config.Configuration) error {
	fmt.Printf("Box provider: %s\n", agent.Provider.Description())
	if err := agent.Provider.Setup(conf); err != nil {
		return err
	}
	if err := agent.Provider.Authenticate(); err != nil {
		return err
	}

	description, err := agent.Provider.Informations()
	if err != nil {
		return err
	}
	fmt.Printf(yellowOut("== Box ==\n"))
	for k, v := range description.Informations {
		fmt.Printf("%s: %s\n", k, v)
	}

	network, err := agent.Provider.Network()
	if err != nil {
		return err
	}
	fmt.Printf(yellowOut("== Network ==\n"))
	fmt.Printf("IPAddress: %s\n", network.IPV4Address)
	fmt.Printf("DNS: %s\n", network.DNS)
	fmt.Printf("State: %s\n", network.State)

	wifi, err := agent.Provider.Wifi()
	if err != nil {
		return err
	}
	fmt.Printf(yellowOut("== Wifi ==\n"))
	fmt.Printf("Enabled: %t\n", wifi.Enable)
	fmt.Printf("State: %t\n", wifi.State)

	tv, err := agent.Provider.TV()
	if err != nil {
		return err
	}
	fmt.Printf(yellowOut("== TV ==\n"))
	for _, st := range tv {
		fmt.Printf("- %s: %t\n", st.Name, st.State)
	}

	devices, err := agent.Provider.Devices()
	if err != nil {
		return err
	}
	fmt.Printf(yellowOut("== Devices ==\n"))
	for _, dev := range devices {
		fmt.Printf("- %s: %s [%s]\n", dev.Name, dev.IPAddress, dev.Type)
	}
	return nil
}
