// Copyright (C) 2017 Nicolas Lamirault <nicolas.lamirault@gmail.com>

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
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/urfave/cli"

	"github.com/nlamirault/skybox/config"
	"github.com/nlamirault/skybox/metrics"
	"github.com/nlamirault/skybox/providers"
)

// MetricsCommand is the command which manage metrics for ski resorts
var MetricsCommand = cli.Command{
	Name: "metrics",
	Subcommands: []cli.Command{
		metricsExportCommand,
		metricsDryRunCommand,
	},
}

var metricsExportCommand = cli.Command{
	Name:  "export",
	Usage: "Export metrics for the box provider",
	Flags: []cli.Flag{
		cli.IntFlag{
			Name:  "port",
			Usage: "Port to listen on for web interface and telemetry.",
			Value: 9114,
		},
		cli.StringFlag{
			Name:  "metricsPath",
			Usage: "Path under which to expose metrics",
			Value: "/metrics",
		},
	},
	Action: func(context *cli.Context) error {
		configFile, err := getConfigurationFile()
		if err != nil {
			return err
		}
		conf, provider, err := setup(configFile)
		if err != nil {
			return err
		}
		if err := exportMetrics(provider, conf, context.Int("port"), context.String("metricsPath")); err != nil {
			fmt.Println(redOut(err))
		}
		return nil
	},
}
var metricsDryRunCommand = cli.Command{
	Name:  "dryrun",
	Usage: "Display metrics for the box provider",
	Action: func(context *cli.Context) error {
		configFile, err := getConfigurationFile()
		if err != nil {
			return err
		}
		conf, provider, err := setup(configFile)
		if err != nil {
			return err
		}
		if err := displayMetrics(provider, conf); err != nil {
			fmt.Println(redOut(err))
		}
		return nil
	},
}

func exportMetrics(provider providers.Provider, conf *config.Configuration, port int, metricsPath string) error {
	exporter, err := metrics.NewExporter(provider, conf)
	if err != nil {
		return err
	}
	prometheus.MustRegister(exporter)

	http.Handle(metricsPath, prometheus.Handler())
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`<html>
             <head><title>Skybox Exporter</title></head>
             <body>
             <h1>Skybox Exporter</h1>
             <p><a href='` + metricsPath + `'>Metrics</a></p>
             </body>
             </html>`))
	})

	fmt.Println("Listening on", port)
	return http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}

func displayMetrics(provider providers.Provider, conf *config.Configuration) error {
	// exporter, err := metrics.NewExporter(provider, conf)
	// if err != nil {
	// 	return err
	// }
	return nil
}
