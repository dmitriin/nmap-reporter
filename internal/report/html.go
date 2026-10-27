package report

import (
	_ "embed"
	"html/template"
	"os"

	"github.com/dmitriin/nmap-reporter/internal/parser"
)

//go:embed ../../templates/report.html
var reportTemplate string

func RenderHTML(result *parser.ScanResult, outPath string) error {
	funcs := template.FuncMap{
		"recommendations": parser.RecommendationsForHost,
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
	return tmpl.Execute(f, result)
}