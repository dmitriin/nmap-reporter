package parser

type Recommendation struct {
	Port int
	Text string
}

func RecommendationsForHost(h Host, recs map[int]string) []Recommendation {
	var out []Recommendation
	for _, p := range h.Ports {
		if p.State != "open" {
			continue
		}
		if text, ok := recs[p.PortID]; ok {
			out = append(out, Recommendation{Port: p.PortID, Text: text})
		}
	}
	return out
}