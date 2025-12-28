# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added
- Initial project setup with Go
- Modular architecture with clean separation of concerns
  - `parser` package for markdown parsing with embed support
  - `search` package for fuzzy search engine
  - `render` package for Glamour-based beautiful rendering
- Embedded markdown cheatsheets system using `//go:embed`
- Fuzzy search algorithm with weighted scoring
  - Title matching (highest priority)
  - Tag matching
  - Keyword matching
  - Description and code matching
- Beautiful terminal output using Glamour/Glow
- Initial command examples:
  - Git commands (pull, push, stash, commit, branch operations, rebase, reset)
  - Docker commands (ps, build, run, compose, logs, exec, prune)
- Project documentation:
  - README.md with comprehensive usage guide
  - HTML templates for cheatsheet and roadmap
  - CHANGELOG.md for tracking changes
- Development infrastructure:
  - `.gitignore` for Go projects
  - Modular project structure following Go best practices

### Changed
- Migrated from Zig to Go for better ecosystem support and faster development
- Repository URL set to `github.com/stanzinofree/redfish`

### Technical Details
- Go module initialized with `github.com/stanzinofree/redfish`
- Dependencies:
  - `github.com/charmbracelet/glamour` for markdown rendering
- Search algorithm uses weighted fuzzy matching with coverage bonus
- All command metadata embedded at compile time (no runtime file dependencies)

## [0.1.0] - TBD

### Planned for First Release
- [ ] GitHub Actions CI/CD workflow
- [ ] Code quality checks (gofmt, go vet, golangci-lint)
- [ ] Unit tests for all packages
- [ ] Pre-commit hooks configuration
- [ ] SonarCloud integration
- [ ] Binary releases via goreleaser
- [ ] Homebrew tap for easy installation

---

## Version History Notes

### Migration from Zig (2024-12-27)
Project was initially prototyped in Zig but migrated to Go for:
- More mature ecosystem for CLI tools
- Better library support (Glamour for rendering)
- Faster development iteration
- Easier community contributions
- Superior cross-platform support
