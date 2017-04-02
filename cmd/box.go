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
	"github.com/nlamirault/skybox/providers"
)

// BoxCommand is the command which display informations about box provider
var BoxCommand = cli.Command{
	Name: "box",
	Subcommands: []cli.Command{
		infosCommand,
		checkCommand,
	},
}
var checkCommand = cli.Command{
	Name:  "check",
	Usage: "Check box provider",
	Action: func(context *cli.Context) error {
		configFile, err := getConfigurationFile()
		if err != nil {
			return err
		}
		conf, provider, err := setup(configFile)
		if err != nil {
			return err
		}
		return checkBoxProvider(provider, conf)
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
		conf, provider, err := setup(configFile)
		if err != nil {
			return err
		}
		return boxProviderInformations(provider, conf)
	},
}

func boxProviderInformations(provider providers.Provider, conf *config.Configuration) error {
	fmt.Printf("Box provider: %s\n", provider.Description())
	if err := provider.Setup(conf); err != nil {
		return err
	}
	if err := provider.Authenticate(); err != nil {
		return err
	}

	description, err := provider.Informations()
	if err != nil {
		return err
	}
	fmt.Printf(yellowOut("== Box ==\n"))
	for k, v := range description.Informations {
		fmt.Printf("%s: %s\n", k, v)
	}

	network, err := provider.Network()
	if err != nil {
		return err
	}
	fmt.Printf(yellowOut("== Network ==\n"))
	fmt.Printf("IPAddress: %s\n", network.IPV4Address)
	fmt.Printf("DNS: %s\n", network.DNS)
	fmt.Printf("State: %s\n", network.State)

	wifi, err := provider.Wifi()
	if err != nil {
		return err
	}
	fmt.Printf(yellowOut("== Wifi ==\n"))
	fmt.Printf("Enabled: %t\n", wifi.Enable)
	fmt.Printf("State: %t\n", wifi.State)

	tv, err := provider.TV()
	if err != nil {
		return err
	}
	fmt.Printf(yellowOut("== TV ==\n"))
	for _, st := range tv {
		fmt.Printf("- %s: %t\n", st.Name, st.State)
	}

	devices, err := provider.Devices()
	if err != nil {
		return err
	}
	fmt.Printf(yellowOut("== Devices ==\n"))
	for _, dev := range devices {
		fmt.Printf("- %s: %s [%s]\n", dev.Name, dev.IPAddress, dev.Type)
	}
	return nil
}

func checkBoxProvider(provider providers.Provider, conf *config.Configuration) error {
	fmt.Printf("Check box provider: %s\n", provider.Description())
	if err := provider.Setup(conf); err != nil {
		return err
	}
	if err := provider.Ping(); err != nil {
		return err
	}
	if err := provider.Authenticate(); err != nil {
		return err
	}
	fmt.Printf("Box provider successfully configured\n")
	return nil
}
