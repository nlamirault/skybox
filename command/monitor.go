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

// MonitorCommand defines the CLI command to monitor
type MonitorCommand struct {
	UI cli.Ui
}

// Help display help message about the command
func (c *MonitorCommand) Help() string {
	helpText := `
Usage: skybox monitor [options] action
	Monitor box provider and export metrics
Options:
	` + generalOptionsUsage() + `

        export                     Export the box provider metrics
        dryrun                     Display live metrics
`
	return strings.TrimSpace(helpText)
}

// Synopsis return the command message
func (c *MonitorCommand) Synopsis() string {
	return "Check box provider and output plugins"
}

// Run launch the command
func (c *MonitorCommand) Run(args []string) int {
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
	if len(args) != 1 {
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

	action := args[0]
	switch action {
	case "dryrun":
		c.doDisplayBoxMonitoring(agent, conf)
	case "export":
		c.doExportBoxMonitoring(agent, conf)
	default:
		f.Usage()
	}
	return 0
}

func (c *MonitorCommand) doDisplayBoxMonitoring(agent *Agent, conf *config.Configuration) {
	c.UI.Info(fmt.Sprintf("Display box provider metrics: %s", agent.Provider.Description()))
	log.Printf("[DEBUG] Skybox box provider: %v", agent.Provider)
	if err := agent.Provider.Setup(conf); err != nil {
		c.UI.Error(err.Error())
		return
	}
	if err := agent.Provider.Authenticate(); err != nil {
		c.UI.Error(err.Error())
		return
	}

	statistics, err := agent.Provider.Statistics()
	if err != nil {
		c.UI.Error(err.Error())
		return
	}
	log.Printf("[DEBUG] Skybox response: %s", statistics)
	c.UI.Info("Connection status:")
	c.UI.Output(fmt.Sprintf("Rate: [Up/Down]: %d / %d",
		statistics.RateUp, statistics.RateDown))
	c.UI.Output(fmt.Sprintf("Bytes: [Up/Down]: %d / %d",
		statistics.BytesUp, statistics.BytesDown))
	c.UI.Output(fmt.Sprintf("Bandwidth: [Up/Down]: %d / %d",
		statistics.BandwidthUp, statistics.BandwidthDown))

}

func (c *MonitorCommand) doExportBoxMonitoring(agent *Agent, conf *config.Configuration) {
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
}
