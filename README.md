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

### Freebox

Setup configuration :

```toml
[freebox]
url= "http://mafreebox.freebox.fr/"
token = ""
```
    $ skybox check box

*skybox* will ask for an `app_token` using the API. A message will be displayed on
the Freebox LCD asking the user to grant/deny access to the requesting app.

Once the app has obtained a valid `app_token`, edit your configuration file, and setup this token into the
specific entry: `token`.

### Livebox

Setup configuration :

```toml
# Box Provider
provider = "livebox"

[livebox]
url = "http://192.168.1.1"
username = "xxx"
password = "xxx"
```


## Usage

* Help:

```
$ skybox help
NAME:
   skybox - The box provider toolkit

USAGE:
   skybox [global options] command [command options] [arguments...]

VERSION:
   0.1.0

COMMANDS:
     version
     metrics
     box
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --debug        Enable debug mode
   --help, -h     show help
   --version, -v  print the version
```

* Box provider

```
$ skybox box check
Check box provider: livebox
Box provider successfully configured

$ skybox --debug box infos
Box provider: livebox
== Box ==
AdditionalHardwareVersion:
ProvisioningCode: AUTH.1243.4545.65GB.66GF
DeviceLog:
Country: fr
NumberOfReboots: 7
Manufacturer: Sagemcom
SerialNumber: OPS3343453536
HardwareVersion: SG_LB3_1.2.1
SoftwareVersion: SG30_sip-fr-5.21.1.1
ModemFirmwareVersion:
EnabledOptions:
AdditionalSoftwareVersion: g5-r-sip-fr
SpecVersion: 1.1
ManufacturerOUI: 8CF813
ModelName: SagemcomFast3965_LB2.8
ProductClass: Livebox 3
UpTime: 240356
VendorConfigFileNumberOfEntries: 1
ExternalIPAddress: x.x.x.x
DeviceStatus: Up
Description: SagemcomFast3965_LB2.8 Sagemcom fr
RescueVersion: SG30_sip-fr-5.17.5.1
ManufacturerURL: http://www.sagemcom.com/
FirstUseDate: 0001-01-01T00:00:00Z
== Network ==
IPAddress: x.x.x.x
DNS: 80.10.246.132,81.253.149.2
State: up
== Wifi ==
Enabled: true
State: true
== TV ==
- VOD: true
- Multicast Zapping: true
== Devices ==
- OSMC: 192.168.1.10 [ethernet]
- Synology: 192.168.1.13 [ethernet]
- décodeur TV d'Orange: 192.168.1.11 [ethernet]
- jarvis: 192.168.1.12 [wifi]
- Android: 192.168.1.14 [wifi]
- LIVEBOX: 192.168.1.1 [wifi]
```

* Metrics

```
$ skybox metrics export

```


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

[Grafana]: http://grafana.org/

[toml]: https://github.com/toml-lang/toml
