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
	"log"

	"github.com/influxdata/influxdb/client/v2"

	"github.com/nlamirault/skybox/config"
	"github.com/nlamirault/skybox/outputs"
	"github.com/nlamirault/skybox/version"
)

func init() {
	outputs.Add("influxdb", func() outputs.Output {
		return New()
	})

}

type InfluxDB struct {
	URL       string
	URLs      []string `toml:"urls"`
	Username  string
	Password  string
	Database  string
	UserAgent string
	Client    client.Client

	Precision  string
	UDPPayload int `toml:"udp_payload"`
}

// New returns a InfluxDB Client
func New() *InfluxDB {
	return &InfluxDB{
		UserAgent: fmt.Sprintf("skybox-influxdb-%s", version.Version),
	}
}

func (i *InfluxDB) Setup(config *config.Configuration) error {
	if config.InfluxDB == nil {
		return fmt.Errorf("InfluxDB configuration not found: %v", config)
	}
	i.URL = config.InfluxDB.URL
	i.Username = config.InfluxDB.Username
	i.Password = config.InfluxDB.Password
	i.Database = config.InfluxDB.Database
	log.Printf("[DEBUG] InfluxDB output: %v", i)
	return nil
}

func (i *InfluxDB) Ping() error {
	if i.Client == nil {
		return fmt.Errorf("InfluxDB Client not configured")
	}
	resp, err := i.Client.Query(client.Query{
		Command: fmt.Sprintf("SHOW DATABASES"),
	})
	if err != nil {
		return err
	}
	log.Printf("[DEBUG] InfluxDB Check database response: %v", resp)
	return nil
}

func (i *InfluxDB) Connect() error {
	log.Printf("[DEBUG] InfluxDB Connect: %v", i)
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:      i.URL,
		Username:  i.Username,
		Password:  i.Password,
		UserAgent: i.UserAgent,
		//Timeout:   5,
	})
	if err != nil {
		return err
	}
	// Create Database if it doesn't exist
	log.Printf("[DEBUG] InfluxDB Create database if not exists")
	resp, err := c.Query(client.Query{
		Command: fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", i.Database),
	})
	if err != nil {
		return err
	}
	log.Printf("[DEBUG] InfluxDB connect response: %v", resp)
	i.Client = c
	return nil
}

func (i *InfluxDB) Close() error {
	return nil
}

func (i *InfluxDB) Description() string {
	return "Configuration for InfluxDB server to send metrics to"
}

func (i *InfluxDB) Write(points []*client.Point) error {
	log.Printf("[DEBUG] InfluxDB Make points")
	bp, _ := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  i.Database,
		Precision: "s",
	})

	for _, point := range points {
		bp.AddPoint(point)
	}
	log.Printf("[DEBUG] InfluxDB Write points")
	// Write the batch
	i.Client.Write(bp)
	return nil
}
