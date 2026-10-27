package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/dmitriin/nmap-reporter/internal/dispatch"
	"github.com/dmitriin/nmap-reporter/internal/parser"
	"github.com/dmitriin/nmap-reporter/internal/report"
)

func main() {
	// Если первый аргумент — не флаг и не встроенная команда,
	// пробуем диспетчеризовать во внешний бинарник.
	if len(os.Args) > 1 && os.Args[1][0] != '-' {
		if dispatched, err := dispatch.TryDispatch(os.Args[1], os.Args[2:]); dispatched {
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
			return
		}
	}

	input := flag.String("input", "", "путь к Nmap XML (-oX)")
	output := flag.String("output", "report.html", "путь к выходному файлу")
	flag.Parse()

	if *input == "" {
		fmt.Fprintln(os.Stderr, "укажите -input scan.xml")
		fmt.Fprintln(os.Stderr, "внешние расширения: nmap-reporter-<name> в PATH")
		os.Exit(1)
	}

	result, err := parser.ParseNmapXML(*input)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ошибка парсинга: %v\n", err)
		os.Exit(1)
	}

	if err := report.RenderHTML(result, *output); err != nil {
		fmt.Fprintf(os.Stderr, "ошибка рендера: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Готово: %s (хостов: %d)\n", *output, len(result.Hosts))
}