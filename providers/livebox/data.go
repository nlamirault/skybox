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
	"time"
)

type apiErrorEntryResponse struct {
	ErrorCode   int    `json:"error"`
	Description string `json:"description"`
	Info        string `json:"info"`
}

type apiErrorResponse struct {
	Result struct {
		Status struct {
			Errors []apiErrorEntryResponse
		}
	}
}

type apiAuthenticateResponse struct {
	Status int `json:"status"`
	Data   struct {
		ContextID string `json:"contextID"`
	}
}

type apiDisconnectResponse struct {
}

type apiTimeResponse struct {
	Result struct {
		Status bool `json:"status"`
		Data   struct {
			Time string `json:"time"`
		}
	}
}

type apiConnectionStatusResponse struct {
	Result struct {
		Status struct {
			Base struct {
				Data struct {
					Name   string `json:"Name"`
					Enable bool   `json:"Enable"`
					Status bool   `json:"Status"`
					Flags  string `json:"Flags"`
					ULIntf struct {
					} `json:"ULIntf"`
					LLIntf struct {
						Primdata struct {
							Name string `json:"Name"`
						} `json:"primdata"`
					} `json:"LLIntf"`
				} `json:"data"`
				Primdata struct {
					ULIntf struct {
						Data struct {
						} `json:"data"`
					} `json:"ULIntf"`
					LLIntf struct {
						PppData struct {
						} `json:"ppp_data"`
					} `json:"LLIntf"`
				} `json:"primdata"`
				PppData struct {
					Name   string `json:"Name"`
					Enable bool   `json:"Enable"`
					Status bool   `json:"Status"`
					Flags  string `json:"Flags"`
					ULIntf struct {
						Primdata struct {
							Name string `json:"Name"`
						} `json:"primdata"`
					} `json:"ULIntf"`
					LLIntf struct {
						VlanData struct {
							Name string `json:"Name"`
						} `json:"vlan_data"`
					} `json:"LLIntf"`
				} `json:"ppp_data"`
				VlanData struct {
					Name   string `json:"Name"`
					Enable bool   `json:"Enable"`
					Status bool   `json:"Status"`
					Flags  string `json:"Flags"`
					ULIntf struct {
						PppData struct {
							Name string `json:"Name"`
						} `json:"ppp_data"`
					} `json:"ULIntf"`
					LLIntf struct {
						Eth1 struct {
							Name string `json:"Name"`
						} `json:"eth1"`
					} `json:"LLIntf"`
				} `json:"vlan_data"`
				Eth1 struct {
					Name   string `json:"Name"`
					Enable bool   `json:"Enable"`
					Status bool   `json:"Status"`
					Flags  string `json:"Flags"`
					ULIntf struct {
						VlanData struct {
							Name string `json:"Name"`
						} `json:"vlan_data"`
						VlanMulti struct {
							Name string `json:"Name"`
						} `json:"vlan_multi"`
						VlanVoip struct {
							Name string `json:"Name"`
						} `json:"vlan_voip"`
						VlanIptv1 struct {
							Name string `json:"Name"`
						} `json:"vlan_iptv1"`
						VlanIptv2 struct {
							Name string `json:"Name"`
						} `json:"vlan_iptv2"`
					} `json:"ULIntf"`
					LLIntf struct {
					} `json:"LLIntf"`
				} `json:"eth1"`
			} `json:"base"`
			SixRd struct {
			} `json:"6rd"`
			Alias struct {
				Data struct {
					Alias string `json:"Alias"`
				} `json:"data"`
				Primdata struct {
				} `json:"primdata"`
				PppData struct {
					Alias string `json:"Alias"`
				} `json:"ppp_data"`
				VlanData struct {
					Alias string `json:"Alias"`
				} `json:"vlan_data"`
				Eth1 struct {
					Alias string `json:"Alias"`
				} `json:"eth1"`
			} `json:"alias"`
			Atm struct {
			} `json:"atm"`
			Bridge struct {
			} `json:"bridge"`
			Copy struct {
				Data struct {
				} `json:"data"`
				Primdata struct {
				} `json:"primdata"`
				PppData struct {
				} `json:"ppp_data"`
				VlanData struct {
				} `json:"vlan_data"`
				Eth1 struct {
				} `json:"eth1"`
			} `json:"copy"`
			DhcpAPI struct {
				Data struct {
				} `json:"data"`
				Primdata struct {
				} `json:"primdata"`
				PppData struct {
				} `json:"ppp_data"`
				VlanData struct {
				} `json:"vlan_data"`
				Eth1 struct {
				} `json:"eth1"`
			} `json:"dhcp-api"`
			Dhcp struct {
			} `json:"dhcp"`
			Dhcpv6 struct {
			} `json:"dhcpv6"`
			Dhcpv6Impl struct {
			} `json:"dhcpv6impl"`
			DopSlave struct {
			} `json:"dop-slave"`
			Dsl struct {
			} `json:"dsl"`
			Dslite struct {
			} `json:"dslite"`
			Eth struct {
				Eth1 struct {
					LastChangeTime       int    `json:"LastChangeTime"`
					LastChange           int    `json:"LastChange"`
					CurrentBitRate       int    `json:"CurrentBitRate"`
					MaxBitRateSupported  int    `json:"MaxBitRateSupported"`
					MaxBitRateEnabled    int    `json:"MaxBitRateEnabled"`
					CurrentDuplexMode    string `json:"CurrentDuplexMode"`
					DuplexModeEnabled    string `json:"DuplexModeEnabled"`
					PowerSavingSupported bool   `json:"PowerSavingSupported"`
					PowerSavingEnabled   bool   `json:"PowerSavingEnabled"`
					PhyDevice            string `json:"PhyDevice"`
					PhyID                int    `json:"PhyId"`
					ExternalPhy          bool   `json:"ExternalPhy"`
				} `json:"eth1"`
			} `json:"eth"`
			Gre struct {
			} `json:"gre"`
			Nat struct {
				PppData struct {
					NATEnabled bool `json:"NATEnabled"`
				} `json:"ppp_data"`
				VlanData struct {
					NATEnabled bool `json:"NATEnabled"`
				} `json:"vlan_data"`
				Eth1 struct {
					NATEnabled bool `json:"NATEnabled"`
				} `json:"eth1"`
			} `json:"nat"`
			NetdevAPI struct {
				Data struct {
				} `json:"data"`
				Primdata struct {
				} `json:"primdata"`
				PppData struct {
				} `json:"ppp_data"`
				VlanData struct {
				} `json:"vlan_data"`
				Eth1 struct {
				} `json:"eth1"`
			} `json:"netdev-api"`
			Netdev struct {
				PppData struct {
					NetDevIndex                 int    `json:"NetDevIndex"`
					NetDevType                  string `json:"NetDevType"`
					NetDevFlags                 string `json:"NetDevFlags"`
					NetDevName                  string `json:"NetDevName"`
					LLAddress                   string `json:"LLAddress"`
					TxQueue52128MTU             int    `json:"TxQueue[52/128$MTU"`
					NetDevState                 string `json:"NetDevState"`
					IPv4Forwarding              bool   `json:"IPv4Forwarding"`
					IPv4ForceIGMPVersion        int    `json:"IPv4ForceIGMPVersion"`
					IPv4AcceptSourceRoute       bool   `json:"IPv4AcceptSourceRoute"`
					IPv4AcceptRedirects         bool   `json:"IPv4AcceptRedirects"`
					IPv6AcceptRA                bool   `json:"IPv6AcceptRA"`
					IPv6ActAsRouter             bool   `json:"IPv6ActAsRouter"`
					IPv6AutoConf                bool   `json:"IPv6AutoConf"`
					IPv6MaxRtrSolicitations     int    `json:"IPv6MaxRtrSolicitations"`
					IPv6RtrSolicitationInterval int    `json:"IPv6RtrSolicitationInterval"`
					IPv6AcceptSourceRoute       bool   `json:"IPv6AcceptSourceRoute"`
					IPv6AcceptRedirects         bool   `json:"IPv6AcceptRedirects"`
					IPv6OptimisticDAD           bool   `json:"IPv6OptimisticDAD"`
					IPv6Disable                 bool   `json:"IPv6Disable"`
					IPv6AddrDelegate            string `json:"IPv6AddrDelegate"`
					IPv4Addr                    struct {
						Dyn8 struct {
							Enable    bool   `json:"Enable"`
							Status    string `json:"Status"`
							Address   string `json:"Address"`
							Peer      string `json:"Peer"`
							PrefixLen int    `json:"PrefixLen"`
							Flags     string `json:"Flags"`
							Scope     string `json:"Scope"`
						} `json:"dyn8"`
					} `json:"IPv4Addr"`
					IPv6Addr struct {
					} `json:"IPv6Addr"`
					IPv4Route struct {
						Dyn31 struct {
							Enable   bool   `json:"Enable"`
							Status   string `json:"Status"`
							DstLen   int    `json:"DstLen"`
							Table    string `json:"Table"`
							Scope    string `json:"Scope"`
							Protocol string `json:"Protocol"`
							Type     string `json:"Type"`
							Dst      string `json:"Dst"`
							Priority int    `json:"Priority"`
							Gateway  string `json:"Gateway"`
						} `json:"dyn31"`
						Route struct {
							Enable   bool   `json:"Enable"`
							Status   string `json:"Status"`
							DstLen   int    `json:"DstLen"`
							Table    string `json:"Table"`
							Scope    string `json:"Scope"`
							Protocol string `json:"Protocol"`
							Type     string `json:"Type"`
							Dst      string `json:"Dst"`
							Priority int    `json:"Priority"`
							Gateway  string `json:"Gateway"`
						} `json:"route"`
					} `json:"IPv4Route"`
					IPv6Route struct {
					} `json:"IPv6Route"`
				} `json:"ppp_data"`
				VlanData struct {
					NetDevIndex                 int    `json:"NetDevIndex"`
					NetDevType                  string `json:"NetDevType"`
					NetDevFlags                 string `json:"NetDevFlags"`
					NetDevName                  string `json:"NetDevName"`
					LLAddress                   string `json:"LLAddress"`
					TxQueueLen                  int    `json:"TxQueueLen"`
					MTU                         int    `json:"MTU"`
					NetDevState                 string `json:"NetDevState"`
					IPv4Forwarding              bool   `json:"IPv4Forwarding"`
					IPv4ForceIGMPVersion        int    `json:"IPv4ForceIGMPVersion"`
					IPv4AcceptSourceRoute       bool   `json:"IPv4AcceptSourceRoute"`
					IPv4AcceptRedirects         bool   `json:"IPv4AcceptRedirects"`
					IPv6AcceptRA                bool   `json:"IPv6AcceptRA"`
					IPv6ActAsRouter             bool   `json:"IPv6ActAsRouter"`
					IPv6AutoConf                bool   `json:"IPv6AutoConf"`
					IPv6MaxRtrSolicitations     int    `json:"IPv6MaxRtrSolicitations"`
					IPv6RtrSolicitationInterval int    `json:"IPv6RtrSolicitationInterval"`
					IPv6AcceptSourceRoute       bool   `json:"IPv6AcceptSourceRoute"`
					IPv6AcceptRedirects         bool   `json:"IPv6AcceptRedirects"`
					IPv6OptimisticDAD           bool   `json:"IPv6OptimisticDAD"`
					IPv6Disable                 bool   `json:"IPv6Disable"`
					IPv6AddrDelegate            string `json:"IPv6AddrDelegate"`
					IPv4Addr                    struct {
					} `json:"IPv4Addr"`
					IPv6Addr struct {
					} `json:"IPv6Addr"`
					IPv4Route struct {
					} `json:"IPv4Route"`
					IPv6Route struct {
					} `json:"IPv6Route"`
				} `json:"vlan_data"`
				Eth1 struct {
					NetDevIndex                 int    `json:"NetDevIndex"`
					NetDevType                  string `json:"NetDevType"`
					NetDevFlags                 string `json:"NetDevFlags"`
					NetDevName                  string `json:"NetDevName"`
					LLAddress                   string `json:"LLAddress"`
					TxQueueLen                  int    `json:"TxQueueLen"`
					MTU                         int    `json:"MTU"`
					NetDevState                 string `json:"NetDevState"`
					IPv4Forwarding              bool   `json:"IPv4Forwarding"`
					IPv4ForceIGMVersion         int    `json:"IPv4ForceIGM$Version"`
					IPv4AcceptSourceRoute       bool   `json:"IPv4AcceptSourceRoute"`
					IPv4AcceptRedirects         bool   `json:"IPv4AcceptRedirects"`
					IPv6AcceptRA                bool   `json:"IPv6AcceptRA"`
					IPv6ActAsRouter             bool   `json:"IPv6ActAsRouter"`
					IPv6AutoConf                bool   `json:"IPv6AutoConf"`
					IPv6MaxRtrSolicitations     int    `json:"IPv6MaxRtrSolicitations"`
					IPv6RtrSolicitationInterval int    `json:"IPv6RtrSolicitationInterval"`
					IPv6AcceptSourceRoute       bool   `json:"IPv6AcceptSourceRoute"`
					IPv6AcceptRedirects         bool   `json:"IPv6AcceptRedirects"`
					IPv6OptimisticDAD           bool   `json:"IPv6OptimisticDAD"`
					IPv6Disable                 bool   `json:"IPv6Disable"`
					IPv6AddrDelegate            string `json:"IPv6AddrDelegate"`
					IPv4Addr                    struct {
					} `json:"IPv4Addr"`
					IPv6Addr struct {
					} `json:"IPv6Addr"`
					IPv4Route struct {
					} `json:"IPv4Route"`
					IPv6Route struct {
					} `json:"IPv6Route"`
				} `json:"eth1"`
			} `json:"netdev"`
			Penable struct {
			} `json:"penable"`
			Ppp struct {
				PppData struct {
					Username                        string `json:"Username"`
					ConnectionStatus                string `json:"ConnectionStatus"`
					LastConneTionError              string `json:"LastConne$tionError"`
					MaxMRUSize                      int    `json:"MaxMRUSize"`
					PPPoESessionID                  int    `json:"PPPoESessionID"`
					PPPoEACName                     string `json:"PPPoEACName"`
					PPPoEServiceName                string `json:"PPPoEServiceName"`
					RemoteIPAddress                 string `json:"RemoteIPAddress"`
					LocalIPAddress                  string `json:"LocalIPAddress"`
					LastChangeTime                  int    `json:"LastChangeTime"`
					LastChange                      int    `json:"LastChange"`
					DNSServers                      string `json:"DNSServers"`
					TransportType                   string `json:"TransportType"`
					LCPEcho                         int    `json:"LCPEcho"`
					LCPEchoRetry                    int    `json:"LCPEchoRetry"`
					IPCPEnable                      bool   `json:"IPCPEnable"`
					IPv6CPEnable                    bool   `json:"IPv6CPEnable"`
					IPv6CPLocalInterfaceIdentifier  string `json:"IPv6CPLocalInterfaceIdentifier"`
					IPv6CPRemoteInterfaceIdentifier string `json:"IPv6CPRemoteInterfaceIdentifier"`
					ConnectionTrigger               string `json:"ConnectionTrigger"`
					IdleDisconnectTime              int    `json:"IdleDisconnectTime"`
				} `json:"ppp_data"`
			} `json:"ppp"`
			Ptm struct {
			} `json:"ptm"`
			RaAPI struct {
				Data struct {
				} `json:"data"`
				Primdata struct {
				} `json:"primdata"`
				PppData struct {
				} `json:"ppp_data"`
				VlanData struct {
				} `json:"vlan_data"`
				Eth1 struct {
				} `json:"eth1"`
			} `json:"ra-api"`
			Ra struct {
			} `json:"ra"`
			SwTch struct {
			} `json:"sw$tch"`
			Vlan struct {
				VlanData struct {
					LastChangeTime int `json:"LastChangeTime"`
					LastChange     int `json:"LastChange"`
					VLANID         int `json:"VLANID"`
					VLANPriority   int `json:"VLANPriority"`
				} `json:"vlan_data"`
			} `json:"vlan"`
			Wlanconfig struct {
			} `json:"wlanconfig"`
			Wlanradio struct {
			} `json:"wlanradio"`
			Wlanvap struct {
			} `json:"wlanvap"`
		} `json:"status"`
	} `json:"result"`
}

type apiWanStatusResponse struct {
	Result struct {
		Status bool `json:"status"`
		Data   struct {
			LinkType            string `json:"LinkType"`
			LinkState           string `json:"LinkState"`
			MACAddress          string `json:"MACAddress"`
			Protocol            string `json:"Protocol"`
			ConnectionState     string `json:"ConnectionState"`
			LastConnectionError string `json:"LastConnectionError"`
			IPAddress           string `json:"IPAddress"`
			RemoteGateway       string `json:"RemoteGateway"`
			DNSServers          string `json:"DNSServers"`
			IPv6Address         string `json:"IPv6Address"`
			IPv6DelegatedPrefix string `json:"IPv6DelegatedPrefix"`
		} `json:"data"`
	} `json:"result"`
}

type apiTVStatusResponse struct {
	Result struct {
		Status []struct {
			ChannelStatus bool   `json:"ChannelStatus"`
			ChannelType   string `json:"ChannelType"`
			ChannelNumber string `json:"ChannelNumber"`
			ChannelFlags  string `json:"ChannelFlags"`
		} `json:"status"`
	} `json:"result"`
}

type apiWifiStatusResponse struct {
	Result struct {
		Status struct {
			Enable            bool `json:"Enable"`
			Status            bool `json:"Status"`
			ConfigurationMode bool `json:"ConfigurationMode"`
		} `json:"status"`
	} `json:"result"`
}

type apiDevicesResponse struct {
	Result struct {
		Status []struct {
			Key              string    `json:"Key"`
			DiscoverySource  string    `json:"DiscoverySource"`
			Name             string    `json:"Name"`
			DeviceType       string    `json:"DeviceType"`
			Active           bool      `json:"Active"`
			Tags             string    `json:"Tags"`
			LastConnection   time.Time `json:"LastConnection"`
			LastChanged      time.Time `json:"LastChanged"`
			Master           string    `json:"Master"`
			PhysAddress      string    `json:"PhysAddress"`
			Ageing           bool      `json:"Ageing"`
			Layer2Interface  string    `json:"Layer2Interface"`
			IPAddress        string    `json:"IPAddress"`
			IPAddressSource  string    `json:"IPAddressSource"`
			VendorClassID    string    `json:"VendorClassID"`
			UserClassID      string    `json:"UserClassID"`
			ClientID         string    `json:"ClientID"`
			SerialNumber     string    `json:"SerialNumber"`
			ProductClass     string    `json:"ProductClass"`
			OUI              string    `json:"OUI"`
			SignalStrength   int       `json:"SignalStrength"`
			SignalNoiseRatio int       `json:"SignalNoiseRatio"`
			Index            string    `json:"Index"`
			Actions          []struct {
				Function  string `json:"Function"`
				Name      string `json:"Name"`
				Arguments []struct {
					Name      string `json:"Name"`
					Type      string `json:"Type"`
					Mandatory bool   `json:"Mandatory"`
				} `json:"Arguments"`
			} `json:"Actions"`
			Alternative []interface{} `json:"Alternative"`
			Names       []struct {
				Name   string `json:"Name"`
				Source string `json:"Source"`
			} `json:"Names"`
			DeviceTypes []struct {
				Type   string `json:"Type"`
				Source string `json:"Source"`
			} `json:"DeviceTypes"`
			IPv4Address []struct {
				Address       string `json:"Address"`
				Status        string `json:"Status"`
				Scope         string `json:"Scope"`
				AddressSource string `json:"AddressSource"`
				Reserved      bool   `json:"Reserved"`
			} `json:"IPv4Address"`
			IPv6Address []interface{} `json:"IPv6Address"`
			MDNSService []struct {
				Name        string `json:"Name"`
				ServiceName string `json:"ServiceName"`
				Domain      string `json:"Domain"`
				Port        string `json:"Port"`
				Text        string `json:"Text"`
			} `json:"mDNSService,omitempty"`

			// UsbM2M     []interface{} `json:"usbM2M"`
			// Usb        []interface{} `json:"usb"`
			// Usblogical []interface{} `json:"usblogical"`
			// Wifi       []struct {
			// 	Key              string    `json:"Key"`
			// 	DiscoverySource  string    `json:"DiscoverySource"`
			// 	Name             string    `json:"Name"`
			// 	DeviceType       string    `json:"DeviceType"`
			// 	Active           bool      `json:"Active"`
			// 	Tags             string    `json:"Tags"`
			// 	LastConnection   time.Time `json:"LastConnection"`
			// 	LastChanged      time.Time `json:"LastChanged"`
			// 	Master           string    `json:"Master"`
			// 	PhysAddress      string    `json:"PhysAddress"`
			// 	Ageing           bool      `json:"Ageing"`
			// 	Layer2Interface  string    `json:"Layer2Interface"`
			// 	IPAddress        string    `json:"IPAddress"`
			// 	IPAddressSource  string    `json:"IPAddressSource"`
			// 	VendorClassID    string    `json:"VendorClassID"`
			// 	UserClassID      string    `json:"UserClassID"`
			// 	ClientID         string    `json:"ClientID"`
			// 	SerialNumber     string    `json:"SerialNumber"`
			// 	ProductClass     string    `json:"ProductClass"`
			// 	OUI              string    `json:"OUI"`
			// 	SignalStrength   int       `json:"SignalStrength"`
			// 	SignalNoiseRatio int       `json:"SignalNoiseRatio"`
			// 	Index            string    `json:"Index"`
			// 	Actions          []struct {
			// 		Function  string `json:"Function"`
			// 		Name      string `json:"Name"`
			// 		Arguments []struct {
			// 			Name      string `json:"Name"`
			// 			Type      string `json:"Type"`
			// 			Mandatory bool   `json:"Mandatory"`
			// 		} `json:"Arguments"`
			// 	} `json:"Actions"`
			// 	Alternative []interface{} `json:"Alternative"`
			// 	Names       []struct {
			// 		Name   string `json:"Name"`
			// 		Source string `json:"Source"`
			// 	} `json:"Names"`
			// 	DeviceTypes []struct {
			// 		Type   string `json:"Type"`
			// 		Source string `json:"Source"`
			// 	} `json:"DeviceTypes"`
			// 	IPv4Address []struct {
			// 		Address       string `json:"Address"`
			// 		Status        string `json:"Status"`
			// 		Scope         string `json:"Scope"`
			// 		AddressSource string `json:"AddressSource"`
			// 		Reserved      bool   `json:"Reserved"`
			// 	} `json:"IPv4Address"`
			// 	IPv6Address []interface{} `json:"IPv6Address"`
			// 	MDNSService []struct {
			// 		Name        string `json:"Name"`
			// 		ServiceName string `json:"ServiceName"`
			// 		Domain      string `json:"Domain"`
			// 		Port        string `json:"Port"`
			// 		Text        string `json:"Text"`
			// 	} `json:"mDNSService,omitempty"`
			// } `json:"wifi"`
			// Eth []struct {
			// 	Key             string    `json:"Key"`
			// 	DiscoverySource string    `json:"DiscoverySource"`
			// 	Name            string    `json:"Name"`
			// 	DeviceType      string    `json:"DeviceType"`
			// 	Active          bool      `json:"Active"`
			// 	Tags            string    `json:"Tags"`
			// 	LastConnection  time.Time `json:"LastConnection"`
			// 	LastChanged     time.Time `json:"LastChanged"`
			// 	Master          string    `json:"Master"`
			// 	PhysAddress     string    `json:"PhysAddress"`
			// 	Ageing          bool      `json:"Ageing"`
			// 	Layer2Interface string    `json:"Layer2Interface"`
			// 	IPAddress       string    `json:"IPAddress"`
			// 	IPAddressSource string    `json:"IPAddressSource"`
			// 	VendorClassID   string    `json:"VendorClassID"`
			// 	UserClassID     string    `json:"UserClassID"`
			// 	ClientID        string    `json:"ClientID"`
			// 	SerialNumber    string    `json:"SerialNumber"`
			// 	ProductClass    string    `json:"ProductClass"`
			// 	OUI             string    `json:"OUI"`
			// 	Index           string    `json:"Index"`
			// 	Actions         []struct {
			// 		Function  string `json:"Function"`
			// 		Name      string `json:"Name"`
			// 		Arguments []struct {
			// 			Name      string `json:"Name"`
			// 			Type      string `json:"Type"`
			// 			Mandatory bool   `json:"Mandatory"`
			// 		} `json:"Arguments"`
			// 	} `json:"Actions"`
			// 	Alternative []interface{} `json:"Alternative"`
			// 	Names       []struct {
			// 		Name   string `json:"Name"`
			// 		Source string `json:"Source"`
			// 	} `json:"Names"`
			// 	DeviceTypes []struct {
			// 		Type   string `json:"Type"`
			// 		Source string `json:"Source"`
			// 	} `json:"DeviceTypes"`
			// 	IPv4Address []struct {
			// 		Address       string `json:"Address"`
			// 		Status        string `json:"Status"`
			// 		Scope         string `json:"Scope"`
			// 		AddressSource string `json:"AddressSource"`
			// 		Reserved      bool   `json:"Reserved"`
			// 	} `json:"IPv4Address"`
			// 	IPv6Address []interface{} `json:"IPv6Address"`
			// 	MDNSService []struct {
			// 		Name        string `json:"Name"`
			// 		ServiceName string `json:"ServiceName"`
			// 		Domain      string `json:"Domain"`
			// 		Port        string `json:"Port"`
			// 		Text        string `json:"Text"`
			// 	} `json:"mDNSService,omitempty"`
			// } `json:"eth"`
			// Dect []interface{} `json:"dect"`
		} `json:"status"`
	} `json:"result"`
}
