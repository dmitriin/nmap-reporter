package parser

type Host struct {
	IP       string
	Hostname string
	Status   string
	Ports    []Port
}

type Port struct {
	PortID   int
	Protocol string
	State    string
	Service  string
	Product  string
	Version  string
}

type ScanResult struct {
	Target string
	Hosts  []Host
}

func (h Host) OpenPorts() int {
	n := 0
	for _, p := range h.Ports {
		if p.State == "open" {
			n++
		}
	}
	return n
}