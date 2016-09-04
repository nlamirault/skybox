// Copyright (C) 2015, 2016 Nicolas Lamirault <nicolas.lamirault@gmail.com>

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
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestGetConfiguration(t *testing.T) {
	templateFile, err := ioutil.TempFile("", "configuration")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(templateFile.Name())
	data := []byte(`# Skybox configuration file

# Box Provider
provider = "freebox"

# Output plugin
output = "influxdb"

[freebox]
url= "http://mafreebox.freebox.fr"
token = "xxxxxxxx"

[influxdb]
url = "http://localhost:8086"
username = "admin"
password = "foo"
database = "skybox"
`)
	err = ioutil.WriteFile(templateFile.Name(), data, 0700)
	if err != nil {
		t.Fatal(err)
	}
	configuration, err := LoadFileConfig(templateFile.Name())
	if err != nil {
		t.Fatalf("Error with configuration: %v", err)
	}
	fmt.Printf("Configuration : %#v\n", configuration)
	if configuration.BoxProvider != "freebox" {
		t.Fatalf("Configuration box provider failed")
	}
	if configuration.OutputPlugin != "influxdb" {
		t.Fatalf("Configuration output plugin failed")
	}

	// Box Provider
	if configuration.Freebox.URL != "http://mafreebox.freebox.fr" ||
		configuration.Freebox.Token != "xxxxxxxx" {
		t.Fatalf("Configuration Freeboxfailed")
	}

	// Output plugin
	if configuration.InfluxDB.URL != "http://localhost:8086" ||
		configuration.InfluxDB.Username != "admin" ||
		configuration.InfluxDB.Password != "foo" ||
		configuration.InfluxDB.Database != "skybox" {
		t.Fatalf("Configuration InfluxDB failed")
	}

}
