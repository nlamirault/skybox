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

package main

// import (
// 	"fmt"
// 	"os"

// 	"github.com/mitchellh/cli"

// 	_ "github.com/nlamirault/skybox/providers/freebox"
// 	_ "github.com/nlamirault/skybox/providers/livebox"
// 	"github.com/nlamirault/skybox/version"
// )

// func main() {
// 	os.Exit(realMain())
// }

// func realMain() int {
// 	cli := &cli.CLI{
// 		Args:       os.Args[1:],
// 		Commands:   Commands,
// 		HelpFunc:   cli.BasicHelpFunc("skybox"),
// 		HelpWriter: os.Stdout,
// 		Version:    version.Version,
// 	}

// 	exitCode, err := cli.Run()
// 	if err != nil {
// 		UI.Error(fmt.Sprintf("Error executing CLI: %s", err.Error()))
// 		return 1
// 	}

// 	return exitCode
// }

import (
	// "fmt"
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/urfave/cli"

	"github.com/nlamirault/skybox/cmd"
	_ "github.com/nlamirault/skybox/providers/freebox"
	_ "github.com/nlamirault/skybox/providers/livebox"
	"github.com/nlamirault/skybox/version"
)

func main() {
	app := cli.NewApp()
	app.Name = "skybox"
	app.Usage = "The box provider toolkit"
	app.Version = version.Version

	app.Commands = []cli.Command{
		cmd.VersionCommand,
		cmd.CheckCommand,
		cmd.BoxCommand,
	}

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name: "debug",
			//Value: true,
			Usage: "Enable debug mode",
		},
	}
	app.Action = func(context *cli.Context) error {
		if context.Bool("debug") {
			logrus.SetLevel(logrus.DebugLevel)
		} else {
			logrus.SetLevel(logrus.WarnLevel)
		}
		return nil
	}
	//logrus.SetLevel(logrus.DebugLevel)
	if err := app.Run(os.Args); err != nil {
		logrus.Fatal(err)
	}
}
