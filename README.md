# Skybox

[![License Apache 2][badge-license]](LICENSE)

Master :
* [![Circle CI](https://circleci.com/gh/nlamirault/skybox/tree/master.svg?style=svg)](https://circleci.com/gh/nlamirault/skybox/tree/master)

Develop :
* [![Circle CI](https://circleci.com/gh/nlamirault/skybox/tree/develop.svg?style=svg)](https://circleci.com/gh/nlamirault/skybox/tree/develop)

*Skybox* is an agent collecting metrics from a box provider, and writing them into outputs.
You could use [Grafana][] to display nice dashboards.

Supported box providers :

* [Freebox][]

Supported outputs :

* [InfluxDB][]

## Installation

You can download the binaries :

* Architecture i386 [ [linux](https://bintray.com/artifact/download/nlamirault/oss/skybox_linux_386) / [darwin](https://bintray.com/artifact/download/nlamirault/oss/skybox_darwin_386) / [freebsd](https://bintray.com/artifact/download/nlamirault/oss/skybox_freebsd_386) / [netbsd](https://bintray.com/artifact/download/nlamirault/oss/skybox_netbsd_386) / [openbsd](https://bintray.com/artifact/download/nlamirault/oss/skybox_openbsd_386) / [windows](https://bintray.com/artifact/download/nlamirault/oss/skybox_windows_386.exe) ]
* Architecture amd64 [ [linux](https://bintray.com/artifact/download/nlamirault/oss/skybox_linux_amd64) / [darwin](https://bintray.com/artifact/download/nlamirault/oss/skybox_darwin_amd64) / [freebsd](https://bintray.com/artifact/download/nlamirault/oss/skybox_freebsd_amd64) / [netbsd](https://bintray.com/artifact/download/nlamirault/oss/skybox_netbsd_amd64) / [openbsd](https://bintray.com/artifact/download/nlamirault/oss/skybox_openbsd_amd64) / [windows](https://bintray.com/artifact/download/nlamirault/oss/skybox_windows_amd64.exe) ]
* Architecture arm [ [linux](https://bintray.com/artifact/download/nlamirault/oss/skybox_linux_arm) / [freebsd](https://bintray.com/artifact/download/nlamirault/oss/skybox_freebsd_arm) / [netbsd](https://bintray.com/artifact/download/nlamirault/oss/skybox_netbsd_arm) ]



## Configuration

*Skybox* configuration use [toml][] format. File is located into `$HOME/.config/skybox/skybox.toml`.

## Usage

### Freebox

    $ skybox check box

*skybox* will ask for an `app_token` using the API. A message will be displayed on
the Freebox LCD asking the user to grant/deny access to the requesting app.

Once the app has obtained a valid `app_token`, edit your configuration file, and setup this token into the
specific entry:

```toml
[freebox]
url= "http://mafreebox.freebox.fr/"
token = "...."
```

### InfluxDB

Setup configuration :

```toml
[influxdb]
url = "http://localhost:8086/"
username = "root"
password = "root"
database = "skybox"
retentionPolicy = "default"
```

And check connection :

    $ skybox check output

## Development

* Initialize environment

        $ make init

* Build tool :

        $ make build

* Start InfluxDB output (port `8083`) and Grafana (port `3000)
using [Docker Compose][] :

        $ docker-compose up

* Launch unit tests :

        $ make test

## Contributing

See [CONTRIBUTING](CONTRIBUTING.md).


## License

See [LICENSE](LICENSE) for the complete license.


## Changelog

A [changelog](ChangeLog.md) is available


## Contact

Nicolas Lamirault <nicolas.lamirault@gmail.com>


[badge-license]: https://img.shields.io/badge/license-Apache2-green.svg?style=flat

[Freebox]: http://www.free.fr/adsl/freebox-revolution.html

[InfluxDB]: https://influxdata.com/time-series-platform/influxdb/

[Grafana]: http://grafana.org/

[toml]: https://github.com/toml-lang/toml
