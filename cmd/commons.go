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
	// "log"
	"path/filepath"

	"github.com/Sirupsen/logrus"
	"github.com/fatih/color"
	"github.com/mitchellh/go-homedir"

	"github.com/nlamirault/skybox/config"
	"github.com/nlamirault/skybox/providers"
)

var (
	greenOut  = color.New(color.FgGreen).SprintFunc()
	yellowOut = color.New(color.FgYellow).SprintFunc()
	redOut    = color.New(color.FgRed).SprintFunc()
)

const (
	defaultConfigurationFile = ".config/skybox/skybox.toml"
)

func getConfigurationFile() (string, error) {
	home, err := homedir.Dir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, defaultConfigurationFile), nil
}

// // Agent provides a box provider client
// type Agent struct {
// 	Provider providers.Provider
// }

func getConfiguration(filename string) (*config.Configuration, error) {
	conf, err := config.LoadFileConfig(filename)
	if err != nil {
		return nil, err
	}
	return conf, nil
}

func newProvider(conf *config.Configuration) (providers.Provider, error) {
	logrus.Debugf("Box Providers: %v\n", providers.Providers)
	providerCreator := providers.Providers[conf.BoxProvider]
	if providerCreator == nil {
		return nil, fmt.Errorf("No box provider found for %s", conf.BoxProvider)
	}
	provider := providerCreator()
	logrus.Debugf("Box Provider: %s\n", provider)
	// return &Agent{
	// 	Provider: provider,
	// }, nil
	return provider, nil
}

func setup(filename string) (*config.Configuration, providers.Provider, error) {
	conf, err := getConfiguration(filename)
	if err != nil {
		return nil, nil, err
	}
	// agent, err := NewAgent(conf)
	provider, err := newProvider(conf)
	if err != nil {
		return nil, nil, err
	}
	return conf, provider, nil
}
