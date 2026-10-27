# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Planned

- DOCX report generation via the proprietary `nmap-reporter-docx` extension.
- Cloud sync and team templates via `nmap-reporter-cloud` (proprietary).
- Parser tests with a fixture `scan.xml`.
- GitHub Actions release pipeline with GoReleaser.

## [0.1.0] - 2026-10-27

### Added

- **Nmap XML parser** (`internal/parser`): reads `nmap -oX` output and extracts
  hosts, IPs, hostnames, ports, states, services and versions.
- **HTML report generator** (`internal/report`): produces a self-contained
  HTML file with a host summary, per-host port tables and recommendations
  for common exposed services.
- **Bilingual report and CLI** (`internal/i18n`): English (default) and
  Russian via `-lang en|ru`. All labels, messages and recommendations are
  translated through a single template with `{{.T.Field}}` placeholders.
- **PATH-based extension dispatch** (`internal/dispatch`): if an executable
  named `nmap-reporter-<name>` is found in `PATH`, the core dispatches to it
  and passes the remaining arguments. This is the Open Core boundary between
  the MIT-licensed core and proprietary extensions.
- **CLI flags**: `-input` (required), `-output` (default `report.html`),
  `-lang` (default `en`).
- **Recommendations** for common exposed services: FTP, SSH, Telnet, HTTP,
  HTTPS, SMB, MySQL, RDP, PostgreSQL, Redis.

### Security

- The core never executes extension code implicitly: dispatch happens only
  when the first argument is a subcommand (does not start with `-`).
  Built-in commands always take priority over external executables.