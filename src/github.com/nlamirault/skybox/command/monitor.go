// Copyright (C) 2016 Nicolas Lamirault <nicolas.lamirault@gmail.com>

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

	"github.com/mitchellh/cli"

	"github.com/nlamirault/skybox/config"
)

// MonitorCommand defines the CLI command to manage buckets
type MonitorCommand struct {
	UI cli.Ui
}

// Help display help message about the command
func (c *MonitorCommand) Help() string {
	helpText := `
Usage: skybox monitor [options] action
	Monitor box provider and send data using the output plugin
Options:
	` + generalOptionsUsage() + `
Action :
        display                      Retrieve the box provider statistics
        box                          Monitor the box provider
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
	conf, err := getConfiguration(configFile)
	if err != nil {
		c.UI.Error(err.Error())
		return 1
	}
	client, err := NewClient(conf)
	if err != nil {
		c.UI.Error(err.Error())
		return 1
	}
	log.Printf("[DEBUG] Skybox Client: %s", client)

	action := args[0]
	switch action {
	case "display":
		c.doDisplayBoxMonitoring(client, conf)
	case "box":
		c.doBoxMonitoring(client, conf)
	default:
		f.Usage()
	}
	return 0
}

func (c *MonitorCommand) doDisplayBoxMonitoring(client *Client, conf *config.Configuration) {
	c.UI.Info(fmt.Sprintf("Display box provider statistics: %s", client.Provider.Description()))
	log.Printf("[DEBUG] Skybox box provider: %v", client.Provider)
	client.Provider.Setup(conf)
	client.Provider.Authenticate()
	resp, err := client.Provider.Statistics()
	if err != nil {
		c.UI.Error(err.Error())
		return
	}
	c.UI.Output(fmt.Sprintf("Rate: [Up/Down]: %d / %d",
		resp.RateUp, resp.RateDown))
	c.UI.Output(fmt.Sprintf("Bytes: [Up/Down]: %d / %d",
		resp.BytesUp, resp.BytesDown))
	c.UI.Output(fmt.Sprintf("Bandwidth: [Up/Down]: %d / %d",
		resp.BandwidthUp, resp.BandwidthDown))
	c.UI.Output(fmt.Sprintf("Box provider statistics successfully retrieve"))
}

func (c *MonitorCommand) doBoxMonitoring(client *Client, conf *config.Configuration) {
	c.UI.Info(fmt.Sprintf("Display box provider statistics: %s", client.Provider.Description()))
	log.Printf("[DEBUG] Skybox box provider: %v", client.Provider)
	client.Provider.Setup(conf)
	client.Provider.Authenticate()
	client.Provider.Statistics()
	c.UI.Output(fmt.Sprintf("Box provider statistics successfully retrieve"))
}
