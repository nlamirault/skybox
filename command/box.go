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

package command

import (
	"flag"
	"fmt"
	"log"
	"strings"
	// "time"

	"github.com/mitchellh/cli"

	"github.com/nlamirault/skybox/config"
)

// BoxCommand defines the CLI command to display informations about a box provider
type BoxCommand struct {
	UI cli.Ui
}

// Help display help message about the command
func (c *BoxCommand) Help() string {
	helpText := `
Usage: skybox box [options] action
	Box display informations about the box provider
Options:
	` + generalOptionsUsage() + `
`
	return strings.TrimSpace(helpText)
}

// Synopsis return the command message
func (c *BoxCommand) Synopsis() string {
	return "Display box provider informations"
}

// Run launch the command
func (c *BoxCommand) Run(args []string) int {
	var debug bool
	var configFile string
	f := flag.NewFlagSet("monitor", flag.ContinueOnError)
	f.Usage = func() { c.UI.Error(c.Help()) }

	defaultConfigFile, err := getConfigurationFile()
	if err != nil {
		return 1
	}

	f.BoolVar(&debug, "debug", false, "Debug mode enabled")
	f.StringVar(&configFile, "configFile", defaultConfigFile, "Configuration filename")

	if err := f.Parse(args); err != nil {
		return 1
	}
	args = f.Args()
	if len(args) != 0 {
		f.Usage()
		return 1
	}
	setLogging(debug)

	conf, agent, err := setup(configFile)
	if err != nil {
		c.UI.Error(err.Error())
		return 1
	}

	log.Printf("[DEBUG] Skybox agent: %s", agent)
	c.doDisplayBoxInformations(agent, conf)
	return 0
}

func (c *BoxCommand) doDisplayBoxInformations(agent *Agent, conf *config.Configuration) {
	c.UI.Info(fmt.Sprintf("Display box provider statistics: %s", agent.Provider.Description()))
	log.Printf("[DEBUG] Skybox box provider: %v", agent.Provider)

	if err := agent.Provider.Setup(conf); err != nil {
		c.UI.Error(err.Error())
		return
	}
	if err := agent.Provider.Authenticate(); err != nil {
		c.UI.Error(err.Error())
		return
	}

	network, err := agent.Provider.Network()
	if err != nil {
		c.UI.Error(err.Error())
		return
	}
	c.UI.Info("== Network ==")
	c.UI.Output(fmt.Sprintf("IPAddress: %s", network.IPV4Address))
	c.UI.Output(fmt.Sprintf("DNS: %s", network.DNS))
	c.UI.Output(fmt.Sprintf("State: %s", network.State))

	wifi, err := agent.Provider.Wifi()
	if err != nil {
		c.UI.Error(err.Error())
		return
	}
	c.UI.Info("== Wifi ==")
	c.UI.Output(fmt.Sprintf("Enabled: %t", wifi.Enable))
	c.UI.Output(fmt.Sprintf("State: %t", wifi.State))

	devices, err := agent.Provider.Devices()
	if err != nil {
		c.UI.Error(err.Error())
		return
	}
	c.UI.Info("== Devices ==")
	for _, dev := range devices {
		fmt.Printf("- %s: %s [%s]\n", dev.Name, dev.IPAddress, dev.Type)
	}
}
