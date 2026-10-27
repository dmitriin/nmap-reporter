# nmap-reporter

[English](README.md) · [Русский](README.ru.md)

Turn Nmap XML scans into clean HTML reports in seconds.

`nmap-reporter` parses an Nmap XML file (`nmap -oX`) and generates a self-contained HTML report with a host summary, per-host port tables, and basic security recommendations for common exposed services.

## Why

Pentesters and sysadmins run Nmap all the time, but turning the raw XML into something readable for a client or a ticket takes manual work. `nmap-reporter` does that in one command, with no server, no database, and no dependencies beyond a single binary.

## Install

 go install github.com/dmitriin/nmap-reporter@latest

Or build from source:

```cmd
 git clone https://github.com/dmitriin/nmap-reporter.git
 cd nmap-reporter
 go build -o nmap-reporter .
 ```

## Usage

 nmap -sV -oX scan.xml 127.0.0.1
 nmap-reporter -input scan.xml -output report.html

Options:

| Flag | Default | Description |
| ------ | --------- | ------------- |
| `-input` | (required) | Path to Nmap XML (`-oX`) |
| `-output` | `report.html` | Path to the output file |
| `-lang` | `en` | Report language: `en` or `ru` |

Examples:

# English report (default)

 nmap-reporter -input scan.xml -output report-en.html

# Russian report

 nmap-reporter -input scan.xml -output report-ru.html -lang ru

## What's in the report

- Host summary: IP, hostname, status, number of open ports.
- Per-host table: port, protocol, state, service, version.
- Recommendations for common exposed services (SSH, RDP, SMB, Redis, etc.).

## Extensions (Open Core)

The core is MIT-licensed and stays free. Additional functionality is distributed as separate executables discovered via `PATH`:

- `nmap-reporter-docx` — DOCX reports with corporate templates (proprietary).
- `nmap-reporter-cloud` — report sync and team templates (proprietary, planned).

If `nmap-reporter-<name>` is found in `PATH`, the core dispatches to it:

 nmap-reporter docx -input scan.xml

The core never depends on extensions; they are optional.

## License

MIT. See [LICENSE](LICENSE).
