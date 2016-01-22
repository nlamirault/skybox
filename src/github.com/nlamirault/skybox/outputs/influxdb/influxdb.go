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

package influxdb

import (
	"fmt"

	"github.com/influxdata/influxdb/client/v2"

	"github.com/nlamirault/skybox/outputs"
	"github.com/nlamirault/skybox/version"
)

func init() {
	outputs.Add("influxdb", func() outputs.Output {
		return New()
	})

}

type InfluxDB struct {
	URL        string
	URLs       []string `toml:"urls"`
	Username   string
	Password   string
	Database   string
	UserAgent  string
	Precision  string
	UDPPayload int `toml:"udp_payload"`
	conns      []client.Client
}

// New returns a InfluxDB Client
func New() *InfluxDB {
	client := InfluxDB{
		URL:       "http://localhost:8086",
		Username:  "admin",
		Password:  "admin",
		Database:  "skybox",
		UserAgent: fmt.Sprintf("skybox-influxdb-%s", version.Version),
	}
	return &client
}

func (i *InfluxDB) Connect() error {
	return nil
}

func (i *InfluxDB) Close() error {
	return nil
}

func (i *InfluxDB) Description() string {
	return "Configuration for InfluxDB server to send metrics to"
}

func (i *InfluxDB) Write(points []*client.Point) error {
	return nil
}
