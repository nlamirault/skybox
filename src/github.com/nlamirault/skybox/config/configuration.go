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

package config

import (
	"log"

	"github.com/BurntSushi/toml"
)

// Configuration holds configuration for Skybox.
type Configuration struct {
	// BoxProvider is the name of the box provider
	BoxProvider string
	// OutputPlugin is the name of the output plugin to store data
	OutputPlugin string

	// Debug is the option for running in debug mode
	Debug bool

	Freebox *FreeboxConfiguration

	InfluxDB *InfluxdbConfiguration
}

// New returns a Configuration with default values
func New() *Configuration {
	return &Configuration{
		OutputPlugin: "influxdb",
		BoxProvider:  "freebox",
		Freebox: &FreeboxConfiguration{
			URL: "http://mafreebox.freebox.fr",
		},
		InfluxDB: &InfluxdbConfiguration{
			Host:     "localhost:8086",
			Username: "admin",
			Password: "admin"},
	}
}

// LoadFileConfig returns a Configuration from reading the specified file (a toml file).
func LoadFileConfig(file string) (*Configuration, error) {
	log.Printf("[DEBUG] Load configuration file: %s", file)
	configuration := New()
	if _, err := toml.DecodeFile(file, configuration); err != nil {
		return nil, err
	}
	log.Printf("[DEBUG] Configuration : %#v", configuration)
	if configuration.Freebox != nil {
		log.Printf("[DEBUG] Configuration : %#v", configuration.Freebox)
	}
	if configuration.InfluxDB != nil {
		log.Printf("[DEBUG] Configuration : %#v", configuration.InfluxDB)
	}
	return configuration, nil
}

// FreeboxProviderConfiguration defines the configuration for the Freebox provider
type FreeboxConfiguration struct {
	URL   string `toml:"url"`
	Token string `toml:"token"`
}

// InfluxdbConfiguration defines the configuration for AWS KMS provider
type InfluxdbConfiguration struct {
	Host     string `toml:"host"`
	Username string `toml:"username"`
	Password string `toml:"password"`
}
