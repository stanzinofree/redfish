<p align="center">
  <img src="docs/logo.png" alt="Redfish Logo" width="200">
</p>

<h1 align="center">🐟 Redfish</h1>

<p align="center">
  <i>A smart command cheatsheet tool with fuzzy search and beautiful markdown rendering</i>
</p>

<p align="center">
  <a href="https://github.com/stanzinofree/redfish/releases"><img src="https://img.shields.io/badge/version-0.1.0-blue.svg" alt="Version"></a>
  <a href="https://golang.org/"><img src="https://img.shields.io/badge/go-1.23+-00ADD8.svg" alt="Go Version"></a>
  <a href="LICENSE"><img src="https://img.shields.io/badge/license-MIT-green.svg" alt="License"></a>
  <a href="#installation"><img src="https://img.shields.io/badge/platform-macOS%20%7C%20Linux%20%7C%20Windows-lightgrey.svg" alt="Platform"></a>
  <a href="https://stanzinofree.github.io/redfish/"><img src="https://img.shields.io/badge/docs-GitHub%20Pages-blue.svg" alt="Documentation"></a>
  <a href="#contributing"><img src="https://img.shields.io/badge/PRs-welcome-brightgreen.svg" alt="PRs Welcome"></a>
</p>

<p align="center">
  <a href="https://github.com/stanzinofree/redfish/actions/workflows/ci.yml"><img src="https://github.com/stanzinofree/redfish/workflows/CI/badge.svg" alt="CI"></a>
  <a href="https://sonarcloud.io/summary/new_code?id=stanzinofree_redfish"><img src="https://sonarcloud.io/api/project_badges/measure?project=stanzinofree_redfish&metric=alert_status" alt="Quality Gate Status"></a>
  <a href="https://sonarcloud.io/summary/new_code?id=stanzinofree_redfish"><img src="https://sonarcloud.io/api/project_badges/measure?project=stanzinofree_redfish&metric=security_rating" alt="Security Rating"></a>
  <a href="https://sonarcloud.io/summary/new_code?id=stanzinofree_redfish"><img src="https://sonarcloud.io/api/project_badges/measure?project=stanzinofree_redfish&metric=sqale_rating" alt="Maintainability Rating"></a>
  <a href="https://github.com/stanzinofree/redfish/network/updates"><img src="https://img.shields.io/badge/Dependabot-enabled-brightgreen.svg?logo=dependabot" alt="Dependabot Status"></a>
</p>

<p align="center">
  <a href="#features">Features</a> •
  <a href="#installation">Installation</a> •
  <a href="#quick-start">Quick Start</a> •
  <a href="#usage">Usage</a> •
  <a href="https://stanzinofree.github.io/redfish/cheatsheet.html">📖 Cheatsheet</a> •
  <a href="https://stanzinofree.github.io/redfish/roadmap.html">🗺️ Roadmap</a>
</p>

---

Redfish is a blazing-fast CLI tool that helps you find and recall commands using natural language search. Built in Go with embedded markdown cheatsheets and beautiful terminal rendering powered by Glamour.

## ✨ Features

- 🔍 **Fuzzy Search**: Find commands using natural language queries
- 📦 **Embedded Data**: All cheatsheets embedded in the binary - no external files needed
- 🎨 **Beautiful Output**: Markdown rendering with syntax highlighting via Glamour/Glow
- ⚡ **Fast**: Written in Go, instant results
- 🔧 **Extensible**: Easy to add new commands by editing markdown files
- 🌈 **Smart Scoring**: Weighted matching on titles, tags, keywords, and descriptions

## 🚀 Installation

### Via Go Install

```bash
go install github.com/stanzinofree/redfish/cmd/redfish@latest
```

### From Source

```bash
git clone https://github.com/stanzinofree/redfish.git
cd redfish
go build -o redfish cmd/redfish/main.go
sudo mv redfish /usr/local/bin/
```

## 📖 Usage

Search for commands using natural language:

```bash
# Search for git commands
redfish git

# Search for specific workflows
redfish git merge and delete branch

# Search for docker compose
redfish docker compose up

# Multi-word queries
redfish kubernetes deployment rolling update
```

### Example Output

```
Found 4 command(s)

[1] git pull
Tags: git, version-control, sync, remote, fetch
Fetch and merge changes from remote repository

▸ git pull origin main

────────────────────────────────────────────────────────

[2] git push
Tags: git, version-control, sync, remote, upload
Push local commits to remote repository

▸ git push origin main
▸ git push --force-with-lease
```

## 🗂️ Project Structure

```
redfish/
├── cmd/
│   └── redfish/
│       └── main.go              # CLI entry point
├── internal/
│   ├── parser/
│   │   ├── parser.go            # Markdown parser
│   │   └── data/                # Embedded markdown cheatsheets
│   │       ├── git.md
│   │       └── docker.md
│   ├── search/
│   │   └── search.go            # Fuzzy search engine
│   └── render/
│       └── render.go            # Glamour-based renderer
├── docs/
│   ├── cheatsheet.html          # Static documentation
│   └── roadmap.html             # Development roadmap
├── .github/
│   └── workflows/
│       └── ci.yml               # GitHub Actions CI/CD
├── go.mod
├── README.md
└── CHANGELOG.md
```

## 📝 Adding New Commands

1. Create or edit a markdown file in `internal/parser/data/`
2. Follow this format:

```markdown
## command name
**Tags**: tag1, tag2, tag3
**Keywords**: keyword1 keyword2 keyword3

Command description

\`\`\`sh
command example
another example
\`\`\`
```

3. Rebuild the binary:

```bash
go build -o redfish cmd/redfish/main.go
```

## 🔍 How Search Works

Redfish uses a weighted fuzzy matching algorithm:

- **Title matches** (highest priority): 10 points for fuzzy, 8 for exact substring
- **Tag matches**: 5 points for fuzzy, 4 for exact substring  
- **Keyword matches**: 3 points for fuzzy, 2.5 for exact substring
- **Description matches**: 1 point
- **Code matches**: 0.5 points

Scores are multiplied by coverage bonus (percentage of query terms matched).

## 🛠️ Development

### Prerequisites

- Go 1.23 or higher
- Git

### Setup

```bash
git clone https://github.com/stanzinofree/redfish.git
cd redfish
go mod download
```

### Running Tests

```bash
go test ./...
```

### Code Quality

```bash
# Format code
gofmt -w .

# Vet code
go vet ./...

# Run linter (requires golangci-lint)
golangci-lint run
```

### Building

```bash
go build -o redfish cmd/redfish/main.go
```

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Run tests and quality checks (`go test ./... && gofmt -w . && go vet ./...`)
5. Push to the branch (`git push origin feature/amazing-feature`)
6. Open a Pull Request

## 📋 Roadmap

See [docs/roadmap.html](docs/roadmap.html) for the full development roadmap.

**Current Phase**: Core Implementation
- [x] Migrate from Zig to Go
- [x] Modular architecture
- [x] Fuzzy search engine
- [x] Glamour rendering
- [ ] CI/CD setup
- [ ] Quality gates (SonarCloud, tests)

## 📄 License

MIT License - see [LICENSE](LICENSE) file for details

## 🙏 Acknowledgments

- [Glamour](https://github.com/charmbracelet/glamour) - Beautiful markdown rendering
- [Charm](https://charm.sh/) - Amazing CLI tools ecosystem

## 📚 Documentation

- [Command Cheatsheet](docs/cheatsheet.html)
- [Development Roadmap](docs/roadmap.html)
- [GitHub Repository](https://github.com/stanzinofree/redfish)

---

Built with ❤️ using Go and Glamour
