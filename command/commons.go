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
	"fmt"
	"log"
	"path/filepath"
	"strings"

	"github.com/mitchellh/go-homedir"

	"github.com/nlamirault/skybox/config"
	"github.com/nlamirault/skybox/logging"
	"github.com/nlamirault/skybox/providers"
)

const (
	defaultConfigurationFile = ".config/skybox/skybox.toml"
)

// generalOptionsUsage returns the usage documenation for commonly
// available options
func generalOptionsUsage() string {
	general := `
        --debug                       Debug mode enabled
`
	return strings.TrimSpace(general)
}

func checkArguments(args ...string) bool {
	for _, arg := range args {
		if len(arg) == 0 {
			return false
		}
	}
	return true
}

func setLogging(debug bool) {
	if debug {
		logging.SetLogging("DEBUG")
	} else {
		logging.SetLogging("INFO")
	}
}

func getConfigurationFile() (string, error) {
	home, err := homedir.Dir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, defaultConfigurationFile), nil
}

// Agent provides a box provider client
type Agent struct {
	Provider providers.Provider
}

func getConfiguration(filename string) (*config.Configuration, error) {
	conf, err := config.LoadFileConfig(filename)
	if err != nil {
		return nil, err
	}
	return conf, nil
}

// NewAgent creates a new instance of Agent.
func NewAgent(conf *config.Configuration) (*Agent, error) {
	log.Printf("[DEBUG] Box Providers: %v\n", providers.Providers)
	providerCreator := providers.Providers[conf.BoxProvider]
	if providerCreator == nil {
		return nil, fmt.Errorf("No box provider found for %s", conf.BoxProvider)
	}
	provider := providerCreator()
	log.Printf("[DEBUG] Box Provider: %s\n", provider)
	return &Agent{
		Provider: provider,
		// Output:   output,
	}, nil
}

func setup(filename string) (*config.Configuration, *Agent, error) {
	conf, err := getConfiguration(filename)
	if err != nil {
		return nil, nil, err
	}
	agent, err := NewAgent(conf)
	if err != nil {
		return nil, nil, err
	}
	return conf, agent, nil
}
