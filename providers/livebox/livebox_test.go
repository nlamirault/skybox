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
	"result": {
		"status": {
			"base": {
				"data": {
					"Name": "data",
					"Enable": true,
					"Status": true,
					"Flags": "up nat-config enabled",
					"ULIntf": {},
					"LLIntf": {
						"primdata": {
							"Name": "primdata"
						}
					}
				},
				"primdata": {
					"ULIntf": {
						"data": {}
					},
					"LLIntf": {
						"ppp_data": {}
					}
				},
				"ppp_data": {
					"Name": "ppp_data",
					"Enable": true,
					"Status": true,
					"Flags": "ppp netdev logical ipv4 nat-enabled enabled netdev-bound up netdev-up ipv4-up",
					"ULIntf": {
						"primdata": {
							"Name": "primdata"
						}
					},
					"LLIntf": {
						"vlan_data": {
							"Name": "vlan_data"
						}
					}
				},
				"vlan_data": {
					"Name": "vlan_data",
					"Enable": true,
					"Status": true,
					"Flags": "vlan netdev ipv4 enabled up netdev-bound netdev-up",
					"ULIntf": {
						"ppp_data": {
							"Name": "ppp_data"
						}
					},
					"LLIntf": {
						"eth1": {
							"Name": "eth1"
						}
					}
				},
				"eth1": {
					"Name": "eth1",
					"Enable": true,
					"Status": true,
					"Flags": "enabled netdev eth atheth physical wan netdev-bound netdev-up up",
					"ULIntf": {
						"vlan_data": {
							"Name": "vlan_data"
						},
						"vlan_multi": {
							"Name": "vlan_multi"
						},
						"vlan_voip": {
							"Name": "vlan_voip"
						},
						"vlan_iptv1": {
							"Name": "vlan_iptv1"
						},
						"vlan_iptv2": {
							"Name": "vlan_iptv2"
						}
					},
					"LLIntf": {}
				}
			},
			"6rd": {},
			"alias": {
				"data": {
					"Alias": "cpe-data"
				},
				"primdata": {},
				"ppp_data": {
					"Alias": "cpe-ppp_data"
				},
				"vlan_data": {
					"Alias": "cpe-vlan_data"
				},
				"eth1": {
					"Alias": "cpe-eth1"
				}
			},
			"atm": {},
			"bridge": {},
			"copy": {
				"data": {},
				"primdata": {},
				"ppp_data": {},
				"vlan_data": {},
				"eth1": {}
			},
			"dhcp-api": {
				"data": {},
				"primdata": {},
				"ppp_data": {},
				"vlan_data": {},
				"eth1": {}
			},
			"dhcp": {},
			"dhcpv6": {},
			"dhcpv6impl": {},
			"dop-slave": {},
			"dsl": {},
			"dslite": {},
			"eth": {
				"eth1": {
					"LastChangeTime": 50,
					"LastChange": 682671,
					"CurrentBitRate": 1000,
					"MaxBitRateSupported": 1000,
					"MaxBitRateEnabled": -1,
					"CurrentDuplexMode": "Full",
					"DuplexModeEnabled": "Auto",
					"PowerSavingSupported": false,
					"PowerSavingEnabled": false,
					"PhyDevice": "eth0",
					"PhyId": 4,
					"ExternalPhy": false
				}
			},
			"gre": {},
			"nat": {
				"ppp_data": {
					"NATEnabled": true
				},
				"vlan_data": {
					"NATEnabled": false
				},
				"eth1": {
					"NATEnabled": false
				}
			},
			"netdev-api": {
				"data": {},
				"primdata": {},
				"ppp_data": {},
				"vlan_data": {},
				"eth1": {}
			},
			"netdev": {
				"ppp_data": {
					"NetDevIndex": 29,
					"NetDevType": "ppp",
					"NetDevFlags": "up pointopoint noarp multicast",
					"NetDevName": "ppp_data",
					"LLAddress": "",
					"TxQueue[52/128$MTU": 1492,
					"NetDevState": "unknown",
					"IPv4Forwarding": true,
					"IPv4ForceIGMPVersion": 0,
					"IPv4AcceptSourceRoute": true,
					"IPv4AcceptRedirects": false,
					"IPv6AcceptRA": true,
					"IPv6ActAsRouter": true,
					"IPv6AutoConf": true,
					"IPv6MaxRtrSolicitations": 3,
					"IPv6RtrSolicitationInterval": 4000,
					"IPv6AcceptSourceRoute": false,
					"IPv6AcceptRedirects": true,
					"IPv6OptimisticDAD": false,
					"IPv6Disable": true,
					"IPv6AddrDelegate": "",
					"IPv4Addr": {
						"dyn8": {
							"Enable": true,
							"Status": "dynamic",
							"Address": "109.214.93.108",
							"Peer": "193.253.160.3",
							"PrefixLen": 32,
							"Flags": "permanent",
							"Scope": "global"
						}
					},
					"IPv6Addr": {},
					"IPv4Route": {
						"dyn31": {
							"Enable": true,
							"Status": "dynamic",
							"DstLen": 32,
							"Table": "main",
							"Scope": "link",
							"Protocol": "kernel",
							"Type": "unicast",
							"Dst": "193.253.160.3",
							"Priority": 0,
							"Gateway": ""
						},
						"route": {
							"Enable": true,
							"Status": "bound",
							"DstLen": 0,
							"Table": "main",
							"Scope": "global",
							"Protocol": "boot",
							"Type": "unicast",
							"Dst": "0.0.0.0",
							"Priority": 0,
							"Gateway": ""
						}
					},
					"IPv6Route": {}
				},
				"vlan_data": {
					"NetDevIndex": 26,
					"NetDevType": "ether",
					"NetDevFlags": "up broadcast multicast",
					"NetDevName": "vlan_data",
					"LLAddress": "8C:F8:13:02:DD:9C",
					"TxQueueLen": 0,
					"MTU": 1500,
					"NetDevState": "up",
					"IPv4Forwarding": true,
					"IPv4ForceIGMPVersion": 0,
					"IPv4AcceptSourceRoute": true,
					"IPv4AcceptRedirects": false,
					"IPv6AcceptRA": true,
					"IPv6ActAsRouter": true,
					"IPv6AutoConf": true,
					"IPv6MaxRtrSolicitations": 3,
					"IPv6RtrSolicitationInterval": 4000,
					"IPv6AcceptSourceRoute": false,
					"IPv6AcceptRedirects": true,
					"IPv6OptimisticDAD": false,
					"IPv6Disable": true,
					"IPv6AddrDelegate": "",
					"IPv4Addr": {},
					"IPv6Addr": {},
					"IPv4Route": {},
					"IPv6Route": {}
				},
				"eth1": {
					"NetDevIndex": 6,
					"NetDevType": "ether",
					"NetDevFlags": "up broadcast multicast",
					"NetDevName": "eth1",
					"LLAddress": "8C:F8:13:02:DD:9C",
					"TxQueueLen": 100,
					"MTU": 1500,
					"NetDevState": "unknown",
					"IPv4Forwarding": true,
					"IPv4ForceIGM$Version": 0,
					"IPv4AcceptSourceRoute": true,
					"IPv4AcceptRedirects": false,
					"IPv6AcceptRA": true,
					"IPv6ActAsRouter": true,
					"IPv6AutoConf": true,
					"IPv6MaxRtrSolicitations": 3,
					"IPv6RtrSolicitationInterval": 4000,
					"IPv6AcceptSourceRoute": false,
					"IPv6AcceptRedirects": true,
					"IPv6OptimisticDAD": false,
					"IPv6Disable": true,
					"IPv6AddrDelegate": "",
					"IPv4Addr": {},
					"IPv6Addr": {},
					"IPv4Route": {},
					"IPv6Route": {}
				}
			},
			"penable": {},
			"ppp": {
				"ppp_data": {
					"Username": "fti/2x2zcfy",
					"ConnectionStatus": "Connected",
					"LastConne$tionError": "ERROR_NONE",
					"MaxMRUSize": 1492,
					"PPPoESessionID": 16668,
					"PPPoEACName": "BSBOR654-H101L1112L02R5",
					"PPPoEServiceName": "",
					"RemoteIPAddress": "193.253.160.$",
					"LocalIPAddress": "109.214.93.108",
					"LastChangeTime": 297461,
					"LastChange": 385260,
					"DNSServers": "81.253.149.2,80.10.246.132",
					"TransportType": "PPPoE",
					"LCPEcho": 30,
					"LCPEchoRetry": 3,
					"IPCPEnable": true,
					"IPv6CPEnable": false,
					"IPv6CPLocalInterfaceIdentifier": "0000:0000:0000:0000",
					"IPv6CPRemoteInterfaceIdentifier": "0000:0000:00$0:0000",
					"ConnectionTrigger": "AlwaysOn",
					"IdleDisconnectTime": 0
				}
			},
			"ptm": {},
			"ra-api": {
				"data": {},
				"primdata": {},
				"ppp_data": {},
				"vlan_data": {},
				"eth1": {}
			},
			"ra": {},
			"sw$tch": {},
			"vlan": {
				"vlan_data": {
					"LastChangeTime": 131,
					"LastChange": 682590,
					"VLANID": 835,
					"VLANPriority": 0
				}
			},
			"wlanconfig": {},
			"wlanradio": {},
			"wlanvap": {}
		}
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
