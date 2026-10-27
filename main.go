package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/dmitriin/nmap-reporter/internal/dispatch"
	"github.com/dmitriin/nmap-reporter/internal/i18n"
	"github.com/dmitriin/nmap-reporter/internal/parser"
	"github.com/dmitriin/nmap-reporter/internal/report"
)

func main() {
	// dispatch до flag.Parse: подкоманда — первый аргумент без "-"
	if len(os.Args) > 1 && os.Args[1][0] != '-' {
		if dispatched, err := dispatch.TryDispatch(os.Args[1], os.Args[2:]); dispatched {
			if err != nil {
				fmt.Fprintf(os.Stderr, i18n.New("en").DispatchError+"\n", err)
				os.Exit(1)
			}
			return
		}
	}

	lang := flag.String("lang", "en", "report language: en or ru")
	input := flag.String("input", "", "path to Nmap XML (-oX)")
	output := flag.String("output", "report.html", "path to output file")
	flag.Parse()

	T := i18n.New(*lang)

	if *input == "" {
		fmt.Fprintln(os.Stderr, T.UsageInput)
		fmt.Fprintln(os.Stderr, T.UsageExtensions)
		os.Exit(1)
	}

	result, err := parser.ParseNmapXML(*input)
	if err != nil {
		fmt.Fprintf(os.Stderr, T.ParseError+"\n", err)
		os.Exit(1)
	}

	if err := report.RenderHTML(result, T, *output); err != nil {
		fmt.Fprintf(os.Stderr, T.RenderError+"\n", err)
		os.Exit(1)
	}

	fmt.Printf(T.Done+"\n", *output, len(result.Hosts))
}