package i18n

// Strings — все переводимые строки приложения.
type Strings struct {
	LangCode  string
	EmptyValue string
	// CLI
	UsageInput       string
	UsageExtensions  string
	ParseError       string
	RenderError      string
	Done             string
	DispatchError    string

	// Report (HTML)
	ReportTitle      string
	ReportSource     string
	ReportHosts      string
	HostSummary      string
	ColIP            string
	ColHostname      string
	ColStatus        string
	ColOpenPorts     string
	HostLabel        string
	ColPort          string
	ColProtocol      string
	ColState         string
	ColService       string
	ColVersion       string
	Recommendations  string
	PortLabel        string

	// Recommendations
	Recs map[int]string
}

// New возвращает словарь для языка. Неизвестный язык → английский.
func New(lang string) *Strings {
	switch lang {
	case "ru":
		return &ru
	default:
		return &en
	}
}

var en = Strings{
	UsageInput:      "usage: nmap-reporter -input scan.xml [-output report.html] [-lang en|ru]",
	UsageExtensions: "external extensions are discovered as nmap-reporter-<name> in PATH",
	ParseError:      "parse error: %v",
	RenderError:     "render error: %v",
	Done:            "Done: %s (hosts: %d)",
	DispatchError:   "failed to run %s: %v",

	ReportTitle:     "Nmap Report",
	ReportSource:    "Source",
	ReportHosts:     "Hosts",
	HostSummary:     "Host Summary",
	ColIP:           "IP",
	ColHostname:     "Hostname",
	ColStatus:       "Status",
	ColOpenPorts:    "Open Ports",
	HostLabel:       "Host",
	ColPort:         "Port",
	ColProtocol:     "Protocol",
	ColState:        "State",
	ColService:      "Service",
	ColVersion:      "Version",
	Recommendations: "Recommendations",
	PortLabel:       "Port",
	EmptyValue:      "—",

	Recs: map[int]string{
		21:   "FTP is open — check for anonymous access, replace with SFTP/FTPS.",
		22:   "SSH is open — disable password auth, use keys.",
		23:   "Telnet is open — insecure protocol, replace with SSH.",
		80:   "HTTP is open — check redirect to HTTPS and security headers.",
		443:  "HTTPS is open — check TLS version and certificate expiry.",
		445:  "SMB is open — critical: check patches (EternalBlue).",
		3306: "MySQL is open — should not be exposed to the internet.",
		3389: "RDP is open — enable NLA, restrict access by IP.",
		5432: "PostgreSQL is open — restrict via firewall.",
		6379: "Redis is open — often unauthenticated; enable auth.",
	},
}

var ru = Strings{
	UsageInput:      "использование: nmap-reporter -input scan.xml [-output report.html] [-lang en|ru]",
	UsageExtensions: "внешние расширения ищутся в PATH как nmap-reporter-<name>",
	ParseError:      "ошибка парсинга: %v",
	RenderError:     "ошибка рендера: %v",
	Done:            "Готово: %s (хостов: %d)",
	DispatchError:   "ошибка запуска %s: %v",

	ReportTitle:     "Отчёт Nmap",
	ReportSource:    "Источник",
	ReportHosts:     "Хостов",
	HostSummary:     "Сводка по хостам",
	ColIP:           "IP",
	ColHostname:     "Hostname",
	ColStatus:       "Статус",
	ColOpenPorts:    "Открытых портов",
	HostLabel:       "Хост",
	ColPort:         "Порт",
	ColProtocol:     "Протокол",
	ColState:        "Состояние",
	ColService:      "Сервис",
	ColVersion:      "Версия",
	Recommendations: "Рекомендации",
	PortLabel:       "Порт",
	EmptyValue:      "—",

	Recs: map[int]string{
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
	},
}