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

package livebox

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/nlamirault/skybox/providers"
)

func newLivebox(handler http.HandlerFunc) (*Client, *httptest.Server, error) {
	server := httptest.NewServer(http.HandlerFunc(handler))
	box := New()
	fakeURL, err := url.Parse(server.URL)
	if err != nil {
		return nil, nil, err
	}
	box.Endpoint = fakeURL
	return box, server, nil
}

func TestLiveboxAuthenticate(t *testing.T) {
	box, server, err := newLivebox(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", providers.AcceptHeader)
		w.Header().Set("Set-Cookie", "25200fcf/sessid=Cusei7vG93RWDrabChZ9SlNJ; Path=/")
		fmt.Fprintln(w, `{
     "status":0,
     "data": {
         "contextID":"RmjJzr2UIXk2zFteSiU0i1bK8wUuS8QyhZ6GeWoLKyC82T0K2TH9HGIF1sXJnD6s"
     }
}`)
	})
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()

	resp, err := box.authenticate()
	if err != nil {
		t.Fatalf("Error API call authenticate : %v", err)
	}
	if resp.Data.ContextID != "RmjJzr2UIXk2zFteSiU0i1bK8wUuS8QyhZ6GeWoLKyC82T0K2TH9HGIF1sXJnD6s" {
		t.Fatalf("Freebox API version response: %s", resp)
	}
	if box.ContextID != "RmjJzr2UIXk2zFteSiU0i1bK8wUuS8QyhZ6GeWoLKyC82T0K2TH9HGIF1sXJnD6s" {
		t.Fatalf("Livebox contextID not set: %v", box)
	}
	fmt.Printf("Cookies: %v\n", box.Cookies)
	if len(box.Cookies) != 1 {
		t.Fatalf("Livebox invalid cookies %d", len(box.Cookies))
	}
	if box.Cookies[0].Name != "25200fcf/sessid" || box.Cookies[0].Value != "Cusei7vG93RWDrabChZ9SlNJ" {
		t.Fatalf("Livebox invalid cookie %#v", box.Cookies[0])
	}
}

func TestLiveboxStatistics(t *testing.T) {
	box, server, err := newLivebox(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", providers.AcceptHeader)
		w.Header().Set("Set-Cookie", "25200fcf/sessid=Cusei7vG93RWDrabChZ9SlNJ; Path=/")
		fmt.Fprintln(w, `{
   "result":{
      "status":{
         "base":{
            "data":{
               "Name":"data",
               "Enable":true,
               "Status":true,
               "Flags":"up nat-config enabled",
               "ULIntf":{

               },
               "LLIntf":{
                  "primdata":{
                     "Name":"primdata"
                  }
               }
            },
            "primdata":{
               "ULIntf":{
                  "data":{

                  }
               },
               "LLIntf":{
                  "ppp_data":{

                  }
               }
            },
            "ppp_data":{
               "Name":"ppp_data",
               "Enable":true,
               "Status":true,
               "Flags":"ppp netdev logical ipv4 nat-enabled enabled netdev-bound up netdev-up ipv4-up",
               "ULIntf":{
                  "primdata":{
                     "Name":"primdata"
                  }
               },
               "LLIntf":{
                  "vlan_data":{
                     "Name":"vlan_data"
                  }
               }
            },
            "vlan_data":{
               "Name":"vlan_data",
               "Enable":true,
               "Status":true,
               "Flags":"vlan netdev ipv4 enabled up netdev-bound netdev-up",
               "ULIntf":{
                  "ppp_data":{
                     "Name":"ppp_data"
                  }
               },
               "LLIntf":{
                  "eth1":{
                     "Name":"eth1"
                  }
               }
            },
            "eth1":{
               "Name":"eth1",
               "Enable":true,
               "Status":true,
               "Flags":"enabled netdev eth atheth physical wan netdev-bound netdev-up up",
               "ULIntf":{
                  "vlan_data":{
                     "Name":"vlan_data"
                  },
                  "vlan_multi":{
                     "Name":"vlan_multi"
                  },
                  "vlan_voip":{
                     "Name":"vlan_voip"
                  },
                  "vlan_iptv1":{
                     "Name":"vlan_iptv1"
                  },
                  "vlan_iptv2":{
                     "Name":"vlan_iptv2"
                  }
               },
               "LLIntf":{

               }
            }
         },
         "6rd":{

         },
         "alias":{
            "data":{
               "Alias":"cpe-data"
            },
            "primdata":{

            },
            "ppp_data":{
               "Alias":"cpe-ppp_data"
            },
            "vlan_data":{
               "Alias":"cpe-vlan_data"
            },
            "eth1":{
               "Alias":"cpe-eth1"
            }
         },
         "atm":{

         },
         "bridge":{

         },
         "copy":{
            "data":{

            },
            "primdata":{

            },
            "ppp_data":{

            },
            "vlan_data":{

            },
            "eth1":{

            }
         },
         "dhcp-api":{
            "data":{

            },
            "primdata":{

            },
            "ppp_data":{

            },
            "vlan_data":{

            },
            "eth1":{

            }
         },
         "dhcp":{

         },
         "dhcpv6":{

         },
         "dhcpv6impl":{

         },
         "dop-slave":{

         },
         "dsl":{

         },
         "dslite":{

         },
         "eth":{
            "eth1":{
               "LastChangeTime":50,
               "LastChange":682671,
               "CurrentBitRate":1000,
               "MaxBitRateSupported":1000,
               "MaxBitRateEnabled":-1,
               "CurrentDuplexMode":"Full",
               "DuplexModeEnabled":"Auto",
               "PowerSavingSupported":false,
               "PowerSavingEnabled":false,
               "PhyDevice":"eth0",
               "PhyId":4,
               "ExternalPhy":false
            }
         },
         "gre":{

         },
         "nat":{
            "ppp_data":{
               "NATEnabled":true
            },
            "vlan_data":{
               "NATEnabled":false
            },
            "eth1":{
               "NATEnabled":false
            }
         },
         "netdev-api":{
            "data":{

            },
            "primdata":{

            },
            "ppp_data":{

            },
            "vlan_data":{

            },
            "eth1":{

            }
         },
         "netdev":{
            "ppp_data":{
               "NetDevIndex":29,
               "NetDevType":"ppp",
               "NetDevFlags":"up pointopoint noarp multicast",
               "NetDevName":"ppp_data",
               "LLAddress":"",
               "TxQueue[52/128$MTU":1492,
               "NetDevState":"unknown",
               "IPv4Forwarding":true,
               "IPv4ForceIGMPVersion":0,
               "IPv4AcceptSourceRoute":true,
               "IPv4AcceptRedirects":false,
               "IPv6AcceptRA":true,
               "IPv6ActAsRouter":true,
               "IPv6AutoConf":true,
               "IPv6MaxRtrSolicitations":3,
               "IPv6RtrSolicitationInterval":4000,
               "IPv6AcceptSourceRoute":false,
               "IPv6AcceptRedirects":true,
               "IPv6OptimisticDAD":false,
               "IPv6Disable":true,
               "IPv6AddrDelegate":"",
               "IPv4Addr":{
                  "dyn8":{
                     "Enable":true,
                     "Status":"dynamic",
                     "Address":"109.214.93.108",
                     "Peer":"193.253.160.3",
                     "PrefixLen":32,
                     "Flags":"permanent",
                     "Scope":"global"
                  }
               },
               "IPv6Addr":{

               },
               "IPv4Route":{
                  "dyn31":{
                     "Enable":true,
                     "Status":"dynamic",
                     "DstLen":32,
                     "Table":"main",
                     "Scope":"link",
                     "Protocol":"kernel",
                     "Type":"unicast",
                     "Dst":"193.253.160.3",
                     "Priority":0,
                     "Gateway":""
                  },
                  "route":{
                     "Enable":true,
                     "Status":"bound",
                     "DstLen":0,
                     "Table":"main",
                     "Scope":"global",
                     "Protocol":"boot",
                     "Type":"unicast",
                     "Dst":"0.0.0.0",
                     "Priority":0,
                     "Gateway":""
                  }
               },
               "IPv6Route":{

               }
            },
            "vlan_data":{
               "NetDevIndex":26,
               "NetDevType":"ether",
               "NetDevFlags":"up broadcast multicast",
               "NetDevName":"vlan_data",
               "LLAddress":"8C:F8:13:02:DD:9C",
               "TxQueueLen":0,
               "MTU":1500,
               "NetDevState":"up",
               "IPv4Forwarding":true,
               "IPv4ForceIGMPVersion":0,
               "IPv4AcceptSourceRoute":true,
               "IPv4AcceptRedirects":false,
               "IPv6AcceptRA":true,
               "IPv6ActAsRouter":true,
               "IPv6AutoConf":true,
               "IPv6MaxRtrSolicitations":3,
               "IPv6RtrSolicitationInterval":4000,
               "IPv6AcceptSourceRoute":false,
               "IPv6AcceptRedirects":true,
               "IPv6OptimisticDAD":false,
               "IPv6Disable":true,
               "IPv6AddrDelegate":"",
               "IPv4Addr":{

               },
               "IPv6Addr":{

               },
               "IPv4Route":{

               },
               "IPv6Route":{

               }
            },
            "eth1":{
               "NetDevIndex":6,
               "NetDevType":"ether",
               "NetDevFlags":"up broadcast multicast",
               "NetDevName":"eth1",
               "LLAddress":"8C:F8:13:02:DD:9C",
               "TxQueueLen":100,
               "MTU":1500,
               "NetDevState":"unknown",
               "IPv4Forwarding":true,
               "IPv4ForceIGM$Version":0,
               "IPv4AcceptSourceRoute":true,
               "IPv4AcceptRedirects":false,
               "IPv6AcceptRA":true,
               "IPv6ActAsRouter":true,
               "IPv6AutoConf":true,
               "IPv6MaxRtrSolicitations":3,
               "IPv6RtrSolicitationInterval":4000,
               "IPv6AcceptSourceRoute":false,
               "IPv6AcceptRedirects":true,
               "IPv6OptimisticDAD":false,
               "IPv6Disable":true,
               "IPv6AddrDelegate":"",
               "IPv4Addr":{

               },
               "IPv6Addr":{

               },
               "IPv4Route":{

               },
               "IPv6Route":{

               }
            }
         },
         "penable":{

         },
         "ppp":{
            "ppp_data":{
               "Username":"fti/2x2zcfy",
               "ConnectionStatus":"Connected",
               "LastConne$tionError":"ERROR_NONE",
               "MaxMRUSize":1492,
               "PPPoESessionID":16668,
               "PPPoEACName":"BSBOR654-H101L1112L02R5",
               "PPPoEServiceName":"",
               "RemoteIPAddress":"193.253.160.$",
               "LocalIPAddress":"109.214.93.108",
               "LastChangeTime":297461,
               "LastChange":385260,
               "DNSServers":"81.253.149.2,80.10.246.132",
               "TransportType":"PPPoE",
               "LCPEcho":30,
               "LCPEchoRetry":3,
               "IPCPEnable":true,
               "IPv6CPEnable":false,
               "IPv6CPLocalInterfaceIdentifier":"0000:0000:0000:0000",
               "IPv6CPRemoteInterfaceIdentifier":"0000:0000:00$0:0000",
               "ConnectionTrigger":"AlwaysOn",
               "IdleDisconnectTime":0
            }
         },
         "ptm":{

         },
         "ra-api":{
            "data":{

            },
            "primdata":{

            },
            "ppp_data":{

            },
            "vlan_data":{

            },
            "eth1":{

            }
         },
         "ra":{

         },
         "sw$tch":{

         },
         "vlan":{
            "vlan_data":{
               "LastChangeTime":131,
               "LastChange":682590,
               "VLANID":835,
               "VLANPriority":0
            }
         },
         "wlanconfig":{

         },
         "wlanradio":{

         },
         "wlanvap":{

         }
      }
   }
}
`)
	})
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()

	resp, err := box.connectionStatus()
	if err != nil {
		t.Fatalf("Error API call statistics : %v", err)
	}
	fmt.Printf("Stats: %s", resp)
	//t.Fatalf("FOO")
}

func TestLiveboxWanStatus(t *testing.T) {
	box, server, err := newLivebox(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", providers.AcceptHeader)
		w.Header().Set("Set-Cookie", "25200fcf/sessid=Cusei7vG93RWDrabChZ9SlNJ; Path=/")
		fmt.Fprintln(w, `{
   "result":{
      "status":true,
      "data":{
         "LinkType":"ethernet",
         "LinkState":"up",
         "MACAddress":"86:ED:13:02:DD:9C",
         "Protocol":"dhcp",
         "ConnectionState":"Bound",
         "LastConnectionError":"None",
         "IPAddress":"193.230.234.111",
         "RemoteGateway":"193.230.234.1",
         "DNSServers":"70.22.246.132,70.14.149.2",
         "IPv6Address":"2a01:cb19:1aa:aabb:efef:13ff:fe02:dd9c",
         "IPv6DelegatedPrefix":"2a01:cb19:1bb:dcdc::/56"
      }
   }
}
`)
	})
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()

	resp, err := box.wanStatus()
	if err != nil {
		t.Fatalf("Error API call wan status : %v", err)
	}
	if resp.Result.Data.IPAddress != "193.230.234.111" ||
		resp.Result.Data.DNSServers != "70.22.246.132,70.14.149.2" {
		t.Fatalf("Error wan status response: %s", resp)
	}
}
func TestLiveboxTVStatus(t *testing.T) {
	box, server, err := newLivebox(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", providers.AcceptHeader)
		w.Header().Set("Set-Cookie", "25200fcf/sessid=Cusei7vG93RWDrabChZ9SlNJ; Path=/")
		fmt.Fprintln(w, `{
   "result":{
      "status":[
         {
            "ChannelStatus":true,
            "ChannelType":"VLAN",
            "ChannelNumber":"838",
            "ChannelFlags":"VOD"
         },
         {
            "ChannelStatus":true,
            "ChannelType":"VLAN",
            "ChannelNumber":"840",
            "ChannelFlags":"Multicast Zapping"
         }
      ]
   }
}
`)
	})
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()

	resp, err := box.tvStatus()
	if err != nil {
		t.Fatalf("Error API call wan status : %v", err)
	}
	if len(resp.Result.Status) != 2 {
		t.Fatalf("Error TV status response: %s", resp)
	}
	if resp.Result.Status[0].ChannelNumber != "838" ||
		resp.Result.Status[0].ChannelFlags != "VOD" ||
		resp.Result.Status[0].ChannelStatus == false ||
		resp.Result.Status[1].ChannelNumber != "840" ||
		resp.Result.Status[1].ChannelFlags != "Multicast Zapping" ||
		resp.Result.Status[1].ChannelStatus == false {
		t.Fatalf("Error TV status channels response: %s", resp)
	}
}
func TestLiveboxDevices(t *testing.T) {
	box, server, err := newLivebox(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", providers.AcceptHeader)
		w.Header().Set("Set-Cookie", "25200fcf/sessid=Cusei7vG93RWDrabChZ9SlNJ; Path=/")
		fmt.Fprintln(w, `{
   "result":{
      "status":{
         "usbM2M":[

         ],
         "usb":[

         ],
         "usblogical":[

         ],
         "wifi":[
            {
               "Key":"B8:27:EB:AE:AD:52",
               "DiscoverySource":"import",
               "Name":"jarvis",
               "DeviceType":"Computer",
               "Active":true,
               "Tags":"lan mac physical edev ipv4 ipv6 dhcp wifi mdns",
               "LastConnection":"2017-02-02T19:55:02Z",
               "LastChanged":"2017-02-02T19:55:01Z",
               "Master":"",
               "PhysAddress":"B8:27:EB:AE:AD:52",
               "Ageing":false,
               "Layer2Interface":"wl0",
               "IPAddress":"192.168.1.12",
               "IPAddressSource":"DHCP",
               "VendorClassID":"dhcpcd-6.7.1:Linux-4.4.39-hypriotos-v7+:armv7l:BCM2709",
               "UserClassID":"",
               "ClientID":"01:B8:27:EB:AE:AD:52",
               "SerialNumber":"",
               "ProductClass":"",
               "OUI":"",
               "SignalStrength":-76,
               "SignalNoiseRatio":0,
               "Index":"7",
               "Actions":[
                  {
                     "Function":"setName",
                     "Name":"Edit Name",
                     "Arguments":[
                        {
                           "Name":"name",
                           "Type":"string",
                           "Mandatory":true
                        },
                        {
                           "Name":"source",
                           "Type":"string",
                           "Mandatory":false
                        }
                     ]
                  }
               ],
               "Alternative":[

               ],
               "Names":[
                  {
                     "Name":"PC-7",
                     "Source":"default"
                  },
                  {
                     "Name":"jarvis",
                     "Source":"dhcp"
                  },
                  {
                     "Name":"jarvis",
                     "Source":"mdns"
                  }
               ],
               "DeviceTypes":[
                  {
                     "Type":"Computer",
                     "Source":"webui"
                  },
                  {
                     "Type":"Computer",
                     "Source":"mdns"
                  }
               ],
               "IPv4Address":[
                  {
                     "Address":"192.168.1.12",
                     "Status":"reachable",
                     "Scope":"global",
                     "AddressSource":"DHCP",
                     "Reserved":false
                  },
                  {
                     "Address":"192.168.1.23",
                     "Status":"error",
                     "Scope":"global",
                     "AddressSource":"Static",
                     "Reserved":true
                  }
               ],
               "IPv6Address":[

               ],
               "mDNSService":[
                  {
                     "Name":"jarvis [b8:27:eb:ae:ad:52]",
                     "ServiceName":"_workstation._tcp",
                     "Domain":"local",
                     "Port":"9",
                     "Text":""
                  }
               ]
            },
            {
               "Key":"8C:EB:C6:09:5E:12",
               "DiscoverySource":"neighborhood",
               "Name":"Android",
               "DeviceType":"Mobile",
               "Active":true,
               "Tags":"lan mac physical ipv4 ipv6 edev dhcp wifi",
               "LastConnection":"2017-02-02T19:55:02Z",
               "LastChanged":"2017-02-02T19:54:33Z",
               "Master":"",
               "IPAddress":"192.168.1.14",
               "IPAddressSource":"DHCP",
               "PhysAddress":"8C:EB:C6:09:5E:12",
               "Ageing":false,
               "Layer2Interface":"wl0",
               "VendorClassID":"dhcpcd-5.5.6",
               "UserClassID":"",
               "ClientID":"01:8C:EB:C6:09:5E:12",
               "SerialNumber":"",
               "ProductClass":"",
               "OUI":"",
               "SignalStrength":-82,
               "SignalNoiseRatio":0,
               "Index":"32",
               "Actions":[
                  {
                     "Function":"setName",
                     "Name":"Edit Name",
                     "Arguments":[
                        {
                           "Name":"name",
                           "Type":"string",
                           "Mandatory":true
                        },
                        {
                           "Name":"source",
                           "Type":"string",
                           "Mandatory":false
                        }
                     ]
                  }
               ],
               "Alternative":[

               ],
               "Names":[
                  {
                     "Name":"PC-32",
                     "Source":"default"
                  },
                  {
                     "Name":"HUAWEI_P9_lite",
                     "Source":"dhcp"
                  },
                  {
                     "Name":"Android",
                     "Source":"mdns"
                  }
               ],
               "DeviceTypes":[
                  {
                     "Type":"Computer",
                     "Source":"default"
                  },
                  {
                     "Type":"Mobile",
                     "Source":"dhcp"
                  }
               ],
               "IPv4Address":[
                  {
                     "Address":"192.168.1.14",
                     "Status":"reachable",
                     "Scope":"global",
                     "AddressSource":"DHCP",
                     "Reserved":false
                  }
               ],
               "IPv6Address":[
                  {
                     "Address":"fe80::8eeb:c6ff:fe09:5e12",
                     "Status":"reachable",
                     "Scope":"link",
                     "AddressSource":"Static"
                  },
                  {
                     "Address":"2a01:cb19:1aa:7700:9081:1da8:a2fb:fe55",
                     "Status":"reachable",
                     "Scope":"global",
                     "AddressSource":"Static"
                  },
                  {
                     "Address":"2a01:cb19:1aa:7700:8eeb:c6ff:fe09:5e12",
                     "Status":"reachable",
                     "Scope":"global",
                     "AddressSource":"Static"
                  }
               ]
            },
            {
               "Key":"18:5E:0F:8C:30:25",
               "DiscoverySource":"neighborhood",
               "Name":"nlamirault",
               "DeviceType":"Computer",
               "Active":true,
               "Tags":"lan mac physical ipv4 ipv6 edev dhcp wifi",
               "LastConnection":"2017-02-02T19:55:02Z",
               "LastChanged":"2017-02-02T19:46:43Z",
               "Master":"",
               "IPAddress":"192.168.1.19",
               "IPAddressSource":"DHCP",
               "PhysAddress":"18:5E:0F:8C:30:25",
               "Ageing":false,
               "Layer2Interface":"wl0",
               "VendorClassID":"dhcpcd-6.11.5:Linux-4.9.6-1-ARCH:x86_64:GenuineIntel",
               "UserClassID":"",
               "ClientID":"FF:0F:8C:30:25:00:01:00:01:1E:42:27:26:18:5E:0F:8C:30:25",
               "SerialNumber":"",
               "ProductClass":"",
               "OUI":"",
               "SignalStrength":-85,
               "SignalNoiseRatio":0,
               "Index":"33",
               "Actions":[
                  {
                     "Function":"setName",
                     "Name":"Edit Name",
                     "Arguments":[
                        {
                           "Name":"name",
                           "Type":"string",
                           "Mandatory":true
                        },
                        {
                           "Name":"source",
                           "Type":"string",
                           "Mandatory":false
                        }
                     ]
                  }
               ],
               "Alternative":[

               ],
               "Names":[
                  {
                     "Name":"PC-33",
                     "Source":"default"
                  },
                  {
                     "Name":"nlamirault",
                     "Source":"dhcp"
                  }
               ],
               "DeviceTypes":[
                  {
                     "Type":"Computer",
                     "Source":"default"
                  }
               ],
               "IPv4Address":[
                  {
                     "Address":"192.168.1.19",
                     "Status":"reachable",
                     "Scope":"global",
                     "AddressSource":"DHCP",
                     "Reserved":false
                  }
               ],
               "IPv6Address":[
                  {
                     "Address":"2a01:cb19:1aa:7700:1a5e:fff:fe8c:3025",
                     "Status":"reachable",
                     "Scope":"global",
                     "AddressSource":"Static"
                  },
                  {
                     "Address":"fe80::1a5e:fff:fe8c:3025",
                     "Status":"reachable",
                     "Scope":"link",
                     "AddressSource":"Static"
                  }
               ]
            }
         ],
         "eth":[
            {
               "Key":"B8:27:EB:34:1B:D5",
               "DiscoverySource":"import",
               "Name":"OSMC",
               "DeviceType":"squeezebox",
               "Active":true,
               "Tags":"lan mac physical edev ipv4 ipv6 dhcp eth mdns",
               "LastConnection":"2017-02-02T19:55:02Z",
               "LastChanged":"2017-01-31T15:29:14Z",
               "Master":"",
               "PhysAddress":"B8:27:EB:34:1B:D5",
               "Ageing":false,
               "Layer2Interface":"eth0",
               "IPAddress":"192.168.1.10",
               "IPAddressSource":"DHCP",
               "VendorClassID":"",
               "UserClassID":"",
               "ClientID":"01:B8:27:EB:34:1B:D5",
               "SerialNumber":"",
               "ProductClass":"",
               "OUI":"",
               "Index":"2",
               "Actions":[
                  {
                     "Function":"setName",
                     "Name":"Edit Name",
                     "Arguments":[
                        {
                           "Name":"name",
                           "Type":"string",
                           "Mandatory":true
                        },
                        {
                           "Name":"source",
                           "Type":"string",
                           "Mandatory":false
                        }
                     ]
                  }
               ],
               "Alternative":[

               ],
               "Names":[
                  {
                     "Name":"OSMC",
                     "Source":"default"
                  },
                  {
                     "Name":"OSMC",
                     "Source":"webui"
                  },
                  {
                     "Name":"osmc",
                     "Source":"dns"
                  },
                  {
                     "Name":"osmc",
                     "Source":"mdns"
                  },
                  {
                     "Name":"osmc",
                     "Source":"dhcp"
                  }
               ],
               "DeviceTypes":[
                  {
                     "Type":"squeezebox",
                     "Source":"webui"
                  },
                  {
                     "Type":"Computer",
                     "Source":"mdns"
                  }
               ],
               "IPv4Address":[
                  {
                     "Address":"192.168.1.10",
                     "Status":"reachable",
                     "Scope":"global",
                     "AddressSource":"DHCP",
                     "Reserved":true
                  }
               ],
               "IPv6Address":[

               ],
               "mDNSService":[
                  {
                     "Name":"osmc [b8:27:eb:34:1b:d5]",
                     "ServiceName":"_workstation._tcp",
                     "Domain":"local",
                     "Port":"9",
                     "Text":""
                  },
                  {
                     "Name":"osmc",
                     "ServiceName":"_http._tcp",
                     "Domain":"local",
                     "Port":"8080",
                     "Text":""
                  },
                  {
                     "Name":"osmc",
                     "ServiceName":"_airplay._tcp",
                     "Domain":"local",
                     "Port":"36667",
                     "Text":"\"features=0x20F7\" \"srcvers=101.28\" \"model=Xbmc,1\" \"deviceid=B8:27:EB:34:1B:D5\""
                  },
                  {
                     "Name":"B827EB341BD5@osmc",
                     "ServiceName":"_raop._tcp",
                     "Domain":"local",
                     "Port":"36666",
                     "Text":"\"txtvers=1\" \"cn=0,1\" \"ch=2\" \"ek=1\" \"et=0,1\" \"sv=false\" \"tp=UDP\" \"sm=false\" \"ss=16\" \"sr=44100\" \"pw=false\" \"vn=3\" \"da=true\" \"md=0,1,2\" \"am=Kodi,1\" \"vs=130.14\""
                  },
                  {
                     "Name":"osmc",
                     "ServiceName":"_xbmc-events._udp",
                     "Domain":"local",
                     "Port":"9777",
                     "Text":""
                  },
                  {
                     "Name":"osmc",
                     "ServiceName":"_xbmc-jsonrpc._tcp",
                     "Domain":"local",
                     "Port":"9090",
                     "Text":""
                  },
                  {
                     "Name":"osmc",
                     "ServiceName":"_xbmc-jsonrpc-h._tcp",
                     "Domain":"local",
                     "Port":"8080",
                     "Text":""
                  },
                  {
                     "Name":"osmc",
                     "ServiceName":"_sftp-ssh._tcp",
                     "Domain":"local",
                     "Port":"22",
                     "Text":"\"path=/home/osmc/\" \"u=osmc\""
                  },
                  {
                     "Name":"osmc",
                     "ServiceName":"_ssh._tcp",
                     "Domain":"local",
                     "Port":"22",
                     "Text":""
                  },
                  {
                     "Name":"osmc",
                     "ServiceName":"_udisks-ssh._tcp",
                     "Domain":"local",
                     "Port":"22",
                     "Text":""
                  }
               ]
            },
            {
               "Key":"00:11:32:13:1F:61",
               "DiscoverySource":"import",
               "Name":"Synology",
               "DeviceType":"computer",
               "Active":true,
               "Tags":"lan mac physical edev ipv4 ipv6 dhcp eth upnp mdns",
               "LastConnection":"2017-02-02T19:55:02Z",
               "LastChanged":"2017-01-31T09:42:11Z",
               "Master":"",
               "PhysAddress":"00:11:32:13:1F:61",
               "Ageing":false,
               "Layer2Interface":"eth0",
               "IPAddress":"192.168.1.13",
               "IPAddressSource":"DHCP",
               "VendorClassID":"",
               "UserClassID":"",
               "ClientID":"01:00:11:32:13:1F:61",
               "SerialNumber":"",
               "ProductClass":"",
               "OUI":"",
               "Index":"3",
               "Actions":[
                  {
                     "Function":"setName",
                     "Name":"Edit Name",
                     "Arguments":[
                        {
                           "Name":"name",
                           "Type":"string",
                           "Mandatory":true
                        },
                        {
                           "Name":"source",
                           "Type":"string",
                           "Mandatory":false
                        }
                     ]
                  }
               ],
               "Alternative":[

               ],
               "Names":[
                  {
                     "Name":"Synology",
                     "Source":"default"
                  },
                  {
                     "Name":"Synology",
                     "Source":"webui"
                  },
                  {
                     "Name":"syno",
                     "Source":"dns"
                  },
                  {
                     "Name":"DiskStation",
                     "Source":"dhcp"
                  }
               ],
               "DeviceTypes":[
                  {
                     "Type":"computer",
                     "Source":"webui"
                  },
                  {
                     "Type":"Basic",
                     "Source":"upnp-uuid:73796E6F-6473-6D00-0000-001132131f61"
                  },
                  {
                     "Type":"Computer",
                     "Source":"mdns"
                  }
               ],
               "IPv4Address":[
                  {
                     "Address":"192.168.1.13",
                     "Status":"reachable",
                     "Scope":"global",
                     "AddressSource":"DHCP",
                     "Reserved":true
                  }
               ],
               "IPv6Address":[
                  {
                     "Address":"2a01:cb19:1aa:7700:211:32ff:fe13:1f61",
                     "Status":"reachable",
                     "Scope":"global",
                     "AddressSource":"Static"
                  },
                  {
                     "Address":"fe80::211:32ff:fe13:1f61",
                     "Status":"reachable",
                     "Scope":"link",
                     "AddressSource":"Static"
                  }
               ],
               "mDNSService":[
                  {
                     "Name":"DiskStation",
                     "ServiceName":"_http._tcp",
                     "Domain":"local",
                     "Port":"5000",
                     "Text":"\"vendor=Synology\" \"model=DS212j\" \"serial=C5KON14633\" \"version_major=6\" \"version_minor=0\" \"version_build=8451\" \"admin_port=5000\" \"secure_admin_port=5001\" \"mac_address=00:11:32:13:1f:61\""
                  },
                  {
                     "Name":"DiskStation",
                     "ServiceName":"_afpovertcp._tcp",
                     "Domain":"local",
                     "Port":"548",
                     "Text":""
                  },
                  {
                     "Name":"DiskStation",
                     "ServiceName":"_sftp._tcp",
                     "Domain":"local",
                     "Port":"22",
                     "Text":""
                  },
                  {
                     "Name":"DiskStation",
                     "ServiceName":"_smb._tcp",
                     "Domain":"local",
                     "Port":"445",
                     "Text":""
                  },
                  {
                     "Name":"DiskStation",
                     "ServiceName":"_device-info._tcp",
                     "Domain":"local",
                     "Port":"0",
                     "Text":"\"model=Xserve\""
                  },
                  {
                     "Name":"DiskStation [00:11:32:13:1f:61]",
                     "ServiceName":"_workstation._tcp",
                     "Domain":"local",
                     "Port":"9",
                     "Text":""
                  }
               ]
            },
            {
               "Key":"18:1E:78:82:C2:25",
               "DiscoverySource":"import",
               "Name":"décodeur TV d'Orange",
               "DeviceType":"SetTopBox",
               "Active":true,
               "Tags":"lan mac physical edev ipv4 ipv6 dhcp stb orange eth upnp",
               "LastConnection":"2017-02-02T19:55:02Z",
               "LastChanged":"2017-02-02T19:39:56Z",
               "Master":"",
               "PhysAddress":"18:1E:78:82:C2:25",
               "Ageing":false,
               "Layer2Interface":"eth0",
               "IPAddress":"192.168.1.11",
               "IPAddressSource":"DHCP",
               "VendorClassID":"sagem",
               "UserClassID":"PC_MLTV_IHD92",
               "ClientID":"01:18:1E:78:82:C2:25",
               "SerialNumber":"",
               "ProductClass":"",
               "OUI":"",
               "Index":"5",
               "Actions":[
                  {
                     "Function":"setName",
                     "Name":"Edit Name",
                     "Arguments":[
                        {
                           "Name":"name",
                           "Type":"string",
                           "Mandatory":true
                        },
                        {
                           "Name":"source",
                           "Type":"string",
                           "Mandatory":false
                        }
                     ]
                  }
               ],
               "Alternative":[

               ],
               "Names":[
                  {
                     "Name":"PC-5",
                     "Source":"default"
                  },
                  {
                     "Name":"PC_MLTV_IHD92",
                     "Source":"UserClassID"
                  },
                  {
                     "Name":"sagem PC_MLTV_IHD92",
                     "Source":"orange"
                  },
                  {
                     "Name":"décodeur TV d'Orange",
                     "Source":"upnp"
                  }
               ],
               "DeviceTypes":[
                  {
                     "Type":"SetTopBox",
                     "Source":"webui"
                  },
                  {
                     "Type":"Set-top Box",
                     "Source":"dhcp"
                  },
                  {
                     "Type":"MediaRenderer",
                     "Source":"upnp-uuid:fd44014a-87d8-11de-a572-181e7882c225"
                  },
                  {
                     "Type":"Basic",
                     "Source":"upnp-uuid:1b13f9b4-23b7-11e2-9360-181e7882c225"
                  }
               ],
               "IPv4Address":[
                  {
                     "Address":"192.168.1.11",
                     "Status":"reachable",
                     "Scope":"global",
                     "AddressSource":"DHCP",
                     "Reserved":false
                  }
               ],
               "IPv6Address":[
                  {
                     "Address":"2a01:cb19:1aa:7700:1a1e:78ff:fe82:c225",
                     "Status":"reachable",
                     "Scope":"global",
                     "AddressSource":"Static"
                  },
                  {
                     "Address":"fe80::1a1e:78ff:fe82:c225",
                     "Status":"reachable",
                     "Scope":"link",
                     "AddressSource":"Static"
                  }
               ]
            },
            {
               "Key":"B8:27:EB:FB:F8:07",
               "DiscoverySource":"import",
               "Name":"wifi bridge",
               "DeviceType":"Computer",
               "Active":true,
               "Tags":"lan mac physical ipv4 ipv6 dhcp eth wifi_bridge hnid",
               "LastConnection":"2017-02-02T19:55:02Z",
               "LastChanged":"2017-02-02T10:33:05Z",
               "Master":"",
               "PhysAddress":"B8:27:EB:FB:F8:07",
               "Ageing":false,
               "Layer2Interface":"eth0",
               "IPAddress":"192.168.1.23",
               "IPAddressSource":"DHCP",
               "VendorClassID":"",
               "UserClassID":"",
               "ClientID":"01:B8:27:EB:FB:F8:07",
               "SerialNumber":"",
               "ProductClass":"",
               "OUI":"",
               "Index":"6",
               "Actions":[
                  {
                     "Function":"setName",
                     "Name":"Edit Name",
                     "Arguments":[
                        {
                           "Name":"name",
                           "Type":"string",
                           "Mandatory":true
                        },
                        {
                           "Name":"source",
                           "Type":"string",
                           "Mandatory":false
                        }
                     ]
                  }
               ],
               "Alternative":[

               ],
               "Names":[
                  {
                     "Name":"wifi bridge",
                     "Source":"default"
                  }
               ],
               "DeviceTypes":[
                  {
                     "Type":"Computer",
                     "Source":"webui"
                  },
                  {
                     "Type":"Computer",
                     "Source":"mdns"
                  },
                  {
                     "Type":"WiFi Bridge",
                     "Source":"default"
                  }
               ],
               "IPv4Address":[
                  {
                     "Address":"192.168.1.23",
                     "Status":"error",
                     "Scope":"global",
                     "AddressSource":"DHCP",
                     "Reserved":true
                  },
                  {
                     "Address":"192.168.1.12",
                     "Status":"error",
                     "Scope":"global",
                     "AddressSource":"DHCP",
                     "Reserved":false
                  }
               ],
               "IPv6Address":[
                  {
                     "Address":"fe80::a383:ed85:ed01:9ebd",
                     "Status":"reachable",
                     "Scope":"link",
                     "AddressSource":"Static"
                  }
               ]
            },
            {
               "Key":"3C:07:54:3C:F9:5A",
               "DiscoverySource":"import",
               "Name":"iMac",
               "DeviceType":"Computer",
               "Active":true,
               "Tags":"lan mac physical edev ipv4 ipv6 dhcp eth mdns",
               "LastConnection":"2017-02-02T19:55:02Z",
               "LastChanged":"2017-02-02T08:44:20Z",
               "Master":"",
               "PhysAddress":"3C:07:54:3C:F9:5A",
               "Ageing":false,
               "Layer2Interface":"eth0",
               "IPAddress":"192.168.1.18",
               "IPAddressSource":"DHCP",
               "VendorClassID":"",
               "UserClassID":"",
               "ClientID":"01:3C:07:54:3C:F9:5A",
               "SerialNumber":"",
               "ProductClass":"",
               "OUI":"",
               "Index":"8",
               "Actions":[
                  {
                     "Function":"setName",
                     "Name":"Edit Name",
                     "Arguments":[
                        {
                           "Name":"name",
                           "Type":"string",
                           "Mandatory":true
                        },
                        {
                           "Name":"source",
                           "Type":"string",
                           "Mandatory":false
                        }
                     ]
                  }
               ],
               "Alternative":[

               ],
               "Names":[
                  {
                     "Name":"PC-8",
                     "Source":"default"
                  },
                  {
                     "Name":"iMac",
                     "Source":"dhcp"
                  },
                  {
                     "Name":"iMac",
                     "Source":"mdns"
                  }
               ],
               "DeviceTypes":[
                  {
                     "Type":"Computer",
                     "Source":"webui"
                  }
               ],
               "IPv4Address":[
                  {
                     "Address":"192.168.1.18",
                     "Status":"reachable",
                     "Scope":"global",
                     "AddressSource":"DHCP",
                     "Reserved":true
                  }
               ],
               "IPv6Address":[
                  {
                     "Address":"fe80::3e07:54ff:fe3c:f95a",
                     "Status":"reachable",
                     "Scope":"link",
                     "AddressSource":"Static"
                  },
                  {
                     "Address":"2a01:cb19:1aa:7700:a594:55d2:b051:f899",
                     "Status":"reachable",
                     "Scope":"global",
                     "AddressSource":"Static"
                  },
                  {
                     "Address":"2a01:cb19:1aa:7700:3e07:54ff:fe3c:f95a",
                     "Status":"reachable",
                     "Scope":"global",
                     "AddressSource":"Static"
                  },
                  {
                     "Address":"2a01:cb19:1aa:7700:e446:e95a:b8f8:68bb",
                     "Status":"reachable",
                     "Scope":"global",
                     "AddressSource":"Static"
                  }
               ],
               "mDNSService":[
                  {
                     "Name":"iMac",
                     "ServiceName":"_sftp-ssh._tcp",
                     "Domain":"local",
                     "Port":"22",
                     "Text":""
                  },
                  {
                     "Name":"iMac",
                     "ServiceName":"_ssh._tcp",
                     "Domain":"local",
                     "Port":"22",
                     "Text":""
                  },
                  {
                     "Name":"iMac",
                     "ServiceName":"_afpovertcp._tcp",
                     "Domain":"local",
                     "Port":"548",
                     "Text":""
                  }
               ]
            }
         ],
         "dect":[

         ]
      }
   }
}
`)
	})
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()

	resp, err := box.devices()
	if err != nil {
		t.Fatalf("Error API call devices : %v", err)
	}
	if len(resp.Result.Status.Wifi) != 3 {
		t.Fatalf("Error devices response wifi: %s", resp.Result.Status.Wifi)
	}
	if len(resp.Result.Status.Eth) != 5 {
		t.Fatalf("Error devices response eth: %s", resp.Result.Status.Eth)
	}
}
