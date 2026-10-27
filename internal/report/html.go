package report

import (
	_ "embed"
	"html/template"
	"os"

	"github.com/dmitriin/nmap-reporter/internal/i18n"
	"github.com/dmitriin/nmap-reporter/internal/parser"
)

//go:embed report.html
var reportTemplate string

// ReportData — то, что получает шаблон: результат скана + словарь строк.
type ReportData struct {
	*parser.ScanResult
	T *i18n.Strings
}

func RenderHTML(result *parser.ScanResult, strings *i18n.Strings, outPath string) error {
	funcs := template.FuncMap{
		"recommendations": func(h parser.Host) []parser.Recommendation {
			return parser.RecommendationsForHost(h, strings.Recs)
		},
	}
	tmpl, err := template.New("report").Funcs(funcs).Parse(reportTemplate)
	if err != nil {
		return err
	}
	f, err := os.Create(outPath)
	if err != nil {
		return err
	}
	defer f.Close()
	return tmpl.Execute(f, ReportData{ScanResult: result, T: strings})
}