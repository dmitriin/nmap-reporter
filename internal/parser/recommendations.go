package parser

type Recommendation struct {
	Port int
	Text string
}

var portRecommendations = map[int]string{
	21:   "FTP открыт — проверьте анонимный доступ, замените на SFTP/FTPS.",
	22:   "SSH открыт — отключите вход по паролю, используйте ключи.",
	23:   "Telnet открыт — небезопасный протокол, замените на SSH.",
	80:   "HTTP открыт — проверьте редирект на HTTPS и заголовки безопасности.",
	443:  "HTTPS открыт — проверьте версию TLS и срок сертификата.",
	445:  "SMB открыт — критично: проверьте патчи (EternalBlue).",
	3306: "MySQL открыт — не должен быть доступен извне.",
	3389: "RDP открыт — включите NLA, ограничьте доступ по IP.",
	5432: "PostgreSQL открыт — ограничьте firewall.",
	6379: "Redis открыт — часто без пароля; включите auth.",
}

func RecommendationsForHost(h Host) []Recommendation {
	var recs []Recommendation
	for _, p := range h.Ports {
		if p.State != "open" {
			continue
		}
		if text, ok := portRecommendations[p.PortID]; ok {
			recs = append(recs, Recommendation{Port: p.PortID, Text: text})
		}
	}
	return recs
}