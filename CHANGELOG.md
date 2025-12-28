# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [0.2.0] - 2025-12-28

### Added
- **Description Display Modes**: Configurable description display with three modes
  - `short`: Brief one-line descriptions (default)
  - `long`: Detailed explanations with context and best practices
  - `none`: Hide descriptions completely
- **Enhanced Markdown Format**: Support for `Short_Description` and `Long_Description` fields in all cheatsheets
- **Configuration Wizard Enhancement**: Interactive prompt for description display preference
- **CLI Flag `-d`**: Override description mode from command line (e.g., `redfish -d long git pull`)
- **Complete Description Coverage**: All embedded cheatsheets now include both short and long descriptions
  - English (EN): docker, git, kubernetes, kubectl, markdown, kcsi
  - Italian (IT): docker, git, kubernetes, kubectl, markdown, kcsi
- **Automatic Release Workflow**: GitHub Actions workflow for automated releases on version tags
  - Multi-platform builds (Linux, macOS, Windows × amd64/arm64/arm)
  - SHA256 checksum generation
  - Automatic GitHub Release creation

### Changed
- **Parser Enhancement**: Updated markdown parser to recognize `Short_Description` and `Long_Description` fields
- **Render Logic**: Modified render engine to display descriptions based on selected mode
- **Config Structure**: Extended configuration with `description_mode` field
- **Backward Compatibility**: Old markdown format (plain `Description`) still fully supported

### Fixed
- **SonarCloud Security Hotspot**: Pinned GitHub Actions to commit SHA instead of mutable tags
- **Code Complexity Issues**: Refactored `synonyms.go` and `main.go` to reduce cognitive complexity
  - Extracted helper functions in synonyms.go: `expandToken()`, `getSynonymMap()`, `containsString()`, `filterOutString()`
  - Split main.go into 13 focused functions with single responsibilities
- **Code Smell**: Grouped consecutive parameters of same type in function signatures
- **CSS Accessibility**: Extracted inline CSS to dedicated files with WCAG 2.1 AA compliant contrast ratios
  - `docs/cheatsheet.css` with 4.5:1+ contrast ratios
  - `docs/roadmap.css` with 7:1+ enhanced contrast

### Technical Details
- Configuration file format updated: `~/.redfish/config.yaml` now includes `description_mode` field
- Render precedence: CLI flag `-d` > config file > default (short)
- All 62 commands across 12 files enhanced with dual descriptions
- Release workflow triggers on `v*` tags with version injection via ldflags

## [0.1.1] - 2024-12-27

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
