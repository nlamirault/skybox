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

package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/common/log"

	"github.com/nlamirault/skybox/config"
	"github.com/nlamirault/skybox/providers"
)

const (
	namespace = "skybox"
)

var (
	devices = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "devices"),
		"Connected devices.",
		[]string{"name", "network_type"}, nil,
	)

	states = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "states"),
		"States of provider services.",
		[]string{"name"}, nil,
	)
)

// Exporter collects metrics from the given ski resort and exports them using
// the prometheus metrics package.
type Exporter struct {
	Provider providers.Provider
	Conf     *config.Configuration
}

// NewExporter returns an initialized Exporter.
func NewExporter(provider providers.Provider, conf *config.Configuration) (*Exporter, error) {
	log.Debugln("Init exporter")
	return &Exporter{
		Provider: provider,
		Conf:     conf,
	}, nil
}

// Describe describes all the metrics ever exported by the exporter.
// It implements prometheus.Collector.
func (e *Exporter) Describe(ch chan<- *prometheus.Desc) {
	ch <- devices
	ch <- states
}

// Collect fetches the stats from configured box and delivers them
// as Prometheus metrics.
// It implements prometheus.Collector.
func (e *Exporter) Collect(ch chan<- prometheus.Metric) {
	log.Infof("Exporter starting")
	if err := e.Provider.Setup(e.Conf); err != nil {
		log.Errorf("Can't setup provider: %s", err.Error())
		return
	}
	if err := e.Provider.Authenticate(); err != nil {
		log.Errorf("Authenticate error: %s", err.Error())
		return
	}

	description, err := e.Provider.Informations()
	if err != nil {
		log.Errorf("Can't retrieve informations: %s", err.Error())
		return
	}

	// for k, v := range description.Informations {
	// 	if k == "DeviceStatus" {
	// 		if v == "Up" {
	// 			ch <- prometheus.MustNewConstMetric(states, prometheus.GaugeValue, 1, k)
	// 		} else {
	// 			ch <- prometheus.MustNewConstMetric(states, prometheus.GaugeValue, 0, k)
	// 		}
	// 	}
	// }
	deviceState := 0
	val, ok := description.Informations["DeviceStatus"]
	if ok {
		if val == "Up" {
			deviceState = 1
		}
	}
	ch <- prometheus.MustNewConstMetric(states, prometheus.GaugeValue, float64(deviceState), "DeviceStatus")

	network, err := e.Provider.Network()
	if err != nil {
		log.Errorf("Can't retrieve network: %s", err.Error())
		return
	}
	networkState := 0
	if network.State == "up" {
		networkState = 1
	}
	log.Infof("Add network metric: %s %d", "network", networkState)
	ch <- prometheus.MustNewConstMetric(states, prometheus.GaugeValue, float64(networkState), "network")

	wifi, err := e.Provider.Wifi()
	if err != nil {
		log.Errorf("Can't retrieve wifi: %s", err.Error())
		return
	}
	wifiState := 0
	if wifi.State {
		wifiState = 1
	}
	log.Infof("Add wifi metric: %s %d", "wifi", wifiState)
	ch <- prometheus.MustNewConstMetric(states, prometheus.GaugeValue, float64(wifiState), "wifi")

	tv, err := e.Provider.TV()
	if err != nil {
		log.Errorf("Can't retrieve TV: %s", err.Error())
		return
	}
	for _, st := range tv {
		val := 0
		if st.State {
			val = 1
		}
		log.Infof("Add metric: %s %d", st.Name, val)
		ch <- prometheus.MustNewConstMetric(states, prometheus.GaugeValue, float64(val), st.Name)

	}

	boxDevices, err := e.Provider.Devices()
	if err != nil {
		log.Errorf("Can't retrieve devices: %s", err.Error())
		return
	}

	for _, dev := range boxDevices {
		ch <- prometheus.MustNewConstMetric(devices, prometheus.GaugeValue, 1, dev.Name, dev.Type)
	}
}
