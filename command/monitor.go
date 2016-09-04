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
	"time"

	"github.com/influxdata/influxdb/client/v2"
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
	agent, err := NewAgent(conf)
	if err != nil {
		c.UI.Error(err.Error())
		return 1
	}
	log.Printf("[DEBUG] Skybox agent: %s", agent)

	action := args[0]
	switch action {
	case "display":
		c.doDisplayBoxMonitoring(agent, conf)
	case "box":
		c.doBoxMonitoring(agent, conf)
	default:
		f.Usage()
	}
	return 0
}

func (c *MonitorCommand) doDisplayBoxMonitoring(agent *Agent, conf *config.Configuration) {
	c.UI.Info(fmt.Sprintf("Display box provider statistics: %s", agent.Provider.Description()))
	log.Printf("[DEBUG] Skybox box provider: %v", agent.Provider)
	agent.Provider.Setup(conf)
	agent.Provider.Authenticate()
	resp, err := agent.Provider.Statistics()
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

func (c *MonitorCommand) doBoxMonitoring(agent *Agent, conf *config.Configuration) {
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
	if err := agent.Output.Setup(conf); err != nil {
		c.UI.Error(err.Error())
		return
	}
	if err := agent.Output.Connect(); err != nil {
		c.UI.Error(err.Error())
		return
	}
	tick := time.Tick(time.Second * time.Duration(conf.Interval))
	for _ = range tick {
		resp, err := agent.Provider.Statistics()
		if err != nil {
			fmt.Printf("Error with box statistics: %s\n", err.Error())
			continue
		}
		fmt.Printf("Rate: [Up/Down]: %d / %d\n",
			resp.RateUp, resp.RateDown)
		fmt.Printf("Bytes: [Up/Down]: %d / %d\n",
			resp.BytesUp, resp.BytesDown)
		fmt.Printf("Bandwidth: [Up/Down]: %d / %d\n",
			resp.BandwidthUp, resp.BandwidthDown)

		var points []*client.Point

		rateTags := map[string]string{"rate": "rate-up-down"}
		rateFields := map[string]interface{}{
			"up":   resp.RateUp,
			"down": resp.RateDown,
		}
		ratePt, err := client.NewPoint("rate", rateTags, rateFields, time.Now())
		if err != nil {
			fmt.Printf("Error creating rate statistics for output: %s\n", err.Error())
			continue
		}
		points = append(points, ratePt)

		bytesTags := map[string]string{"bytes": "bytes-up-down"}
		bytesFields := map[string]interface{}{
			"up":   resp.BytesUp,
			"down": resp.BytesDown,
		}
		bytesPt, err := client.NewPoint("bytes", bytesTags, bytesFields, time.Now())
		if err != nil {
			fmt.Printf("Error creating bytes statistics for output: %s\n", err.Error())
			continue
		}
		points = append(points, bytesPt)

		bandwidthTags := map[string]string{"bandwidth": "bandwidth-up-down"}
		bandwidthFields := map[string]interface{}{
			"up":   resp.BandwidthUp,
			"down": resp.BandwidthDown,
		}
		bandwidthPt, err := client.NewPoint("bandwidth", bandwidthTags, bandwidthFields, time.Now())
		if err != nil {
			fmt.Printf("Error creating bandwidth statistics for output: %s\n", err.Error())
			continue
		}
		points = append(points, bandwidthPt)

		err = agent.Output.Write(points)
		if err != nil {
			fmt.Printf("Error writing statistics : %s\n", err.Error())
			continue
		}
	}
}
