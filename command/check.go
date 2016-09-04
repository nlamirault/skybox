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

// CheckCommand defines the CLI command to manage buckets
type CheckCommand struct {
	UI cli.Ui
}

// Help display help message about the command
func (c *CheckCommand) Help() string {
	helpText := `
Usage: skybox check [options] action
	Check box provider and output plugins
Options:
	` + generalOptionsUsage() + `
Action :
        box                           Check box provider
        output                        Check output plugin
`
	return strings.TrimSpace(helpText)
}

// Synopsis return the command message
func (c *CheckCommand) Synopsis() string {
	return "Check box provider and output plugins"
}

// Run launch the command
func (c *CheckCommand) Run(args []string) int {
	var debug bool
	var configFile string
	f := flag.NewFlagSet("check", flag.ContinueOnError)
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
	agent, err := NewAgent(conf)
	if err != nil {
		c.UI.Error(err.Error())
		return 1
	}
	log.Printf("[DEBUG] Skybox Client: %s", agent)

	action := args[0]
	switch action {
	case "box":
		c.doCheckBoxProvider(agent, conf)
	case "output":
		c.doCheckOutputPlugin(agent, conf)
	default:
		f.Usage()
	}
	return 0
}

func (c *CheckCommand) doCheckBoxProvider(agent *Agent, conf *config.Configuration) {
	c.UI.Info(fmt.Sprintf("Check box provider: %s", agent.Provider.Description()))
	log.Printf("[DEBUG] Skybox box provider: %v", agent.Provider)
	if err := agent.Provider.Setup(conf); err != nil {
		c.UI.Error(err.Error())
		return
	}
	if err := agent.Provider.Ping(); err != nil {
		c.UI.Error(err.Error())
		return
	}
	if err := agent.Provider.Authenticate(); err != nil {
		c.UI.Error(err.Error())
		return
	}
	c.UI.Output(fmt.Sprintf("Box provider successfully configured"))
}

func (c *CheckCommand) doCheckOutputPlugin(agent *Agent, conf *config.Configuration) {
	c.UI.Info(fmt.Sprintf("Check output plugin: %s", agent.Output.Description()))
	log.Printf("[DEBUG] Skybox output plugin: %v", agent.Output)
	if err := agent.Output.Setup(conf); err != nil {
		c.UI.Error(err.Error())
		return
	}
	if err := agent.Output.Connect(); err != nil {
		c.UI.Error(err.Error())
		return
	}
	if err := agent.Output.Ping(); err != nil {
		c.UI.Error(err.Error())
		return
	}
	c.UI.Output(fmt.Sprintf("Output plugin successfully configured"))
}
