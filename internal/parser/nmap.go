package parser

import (
	"encoding/xml"
	"os"
	"strings"
)

type nmapRun struct {
	XMLName xml.Name   `xml:"nmaprun"`
	Hosts   []nmapHost `xml:"host"`
}

type nmapHost struct {
	Status    nmapStatus    `xml:"status"`
	Addresses []nmapAddress `xml:"address"`
	Hostnames nmapHostnames `xml:"hostnames"`
	Ports     []nmapPort    `xml:"ports>port"`
}

type nmapStatus struct {
	State string `xml:"state,attr"`
}

type nmapAddress struct {
	Addr     string `xml:"addr,attr"`
	AddrType string `xml:"addrtype,attr"`
}

type nmapHostnames struct {
	Hostnames []nmapHostname `xml:"hostname"`
}

type nmapHostname struct {
	Name string `xml:"name,attr"`
}

type nmapPort struct {
	PortID   int         `xml:"portid,attr"`
	Protocol string      `xml:"protocol,attr"`
	State    nmapState   `xml:"state"`
	Service  nmapService `xml:"service"`
}

type nmapState struct {
	State string `xml:"state,attr"`
}

type nmapService struct {
	Name    string `xml:"name,attr"`
	Product string `xml:"product,attr"`
	Version string `xml:"version,attr"`
}

func ParseNmapXML(path string) (*ScanResult, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var run nmapRun
	if err := xml.Unmarshal(data, &run); err != nil {
		return nil, err
	}

	result := &ScanResult{
		Target: strings.TrimSuffix(path, ".xml"),
		Hosts:  make([]Host, 0, len(run.Hosts)),
	}

	for _, nh := range run.Hosts {
		h := Host{Status: nh.Status.State}
		for _, a := range nh.Addresses {
			if a.AddrType == "ipv4" || h.IP == "" {
				h.IP = a.Addr
			}
		}
		if len(nh.Hostnames.Hostnames) > 0 {
			h.Hostname = nh.Hostnames.Hostnames[0].Name
		}
		for _, np := range nh.Ports {
			h.Ports = append(h.Ports, Port{
				PortID: np.PortID, Protocol: np.Protocol, State: np.State.State,
				Service: np.Service.Name, Product: np.Service.Product, Version: np.Service.Version,
			})
		}
		result.Hosts = append(result.Hosts, h)
	}
	return result, nil
}