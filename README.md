<p align="center">
  <img src="docs/logo.png?v=2" alt="Redfish Logo" width="200">
</p>

<h1 align="center">🐟 Redfish</h1>

<p align="center">
  <i>A smart command cheatsheet tool with fuzzy search and beautiful markdown rendering</i>
</p>

<p align="center">
  <a href="https://github.com/stanzinofree/redfish/releases"><img src="https://img.shields.io/badge/version-0.1.1-blue.svg" alt="Version"></a>
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

- 🔍 **Intelligent Search**: NLP-like queries with stopwords filtering and synonym expansion
- 🌍 **Multi-Language**: Full support for English and Italian (with ES, FR, DE planned)
- 📦 **Rich Embedded Cheatsheets**: Git, Docker, Kubernetes, Kubectl, Markdown, and more
- 🎨 **Beautiful Output**: Markdown rendering with syntax highlighting via Glamour
- ⚡ **Lightning Fast**: Written in Go, instant results even with large cheatsheet libraries
- 🛡️ **Secure Architecture**: Protected embedded cheatsheets + user custom commands
- 🎯 **Smart Scoring**: Weighted fuzzy matching with coverage-based filtering
- 🔧 **Extensible**: Add your own cheatsheets in `~/.redfish/{lang}/`
- 🎨 **Interactive Mode**: Optional fzf integration for visual command selection
- ⚙️ **User-Friendly**: Configuration wizard, language preferences, help system

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

## 📖 Quick Start

### Basic Search

```bash
# Simple search
redfish git commit

# Multi-word search
redfish docker compose up

# Natural language queries (NLP-like)
redfish "how do i list docker containers"
redfish "voglio fare un merge git"  # Italian
```

### Interactive Mode

```bash
# Use fzf for interactive selection
redfish -f kubernetes
redfish --fzf docker
```

### Configuration

```bash
# Run configuration wizard
redfish -c

# Search in specific language
redfish -l en "kubernetes pods"
redfish -l it "docker container"

# Show version
redfish -v

# Show help
redfish -h
```

### Custom Cheatsheets

```bash
# Add your own commands
mkdir -p ~/.redfish/en
cat > ~/.redfish/en/myapp.md << 'EOF'
## myapp start
**Tags**: myapp, start, run
**Keywords**: start run launch

Start my application

```sh
myapp start --port 8080
```
EOF

# Search your custom commands
redfish myapp
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
│       └── main.go              # CLI entry point (deprecated, use root main.go)
├── main.go                      # Main CLI entry point
├── internal/
│   ├── parser/
│   │   ├── parser.go            # Markdown parser with multi-language support
│   │   └── data/                # Embedded markdown cheatsheets
│   │       ├── en/              # English cheatsheets
│   │       │   ├── git.md
│   │       │   ├── docker.md
│   │       │   ├── kubernetes.md
│   │       │   ├── kubectl.md
│   │       │   ├── markdown.md
│   │       │   └── kcsi.md
│   │       └── it/              # Italian cheatsheets
│   │           ├── git.md
│   │           ├── docker.md
│   │           ├── kubernetes.md
│   │           ├── kubectl.md
│   │           ├── markdown.md
│   │           └── kcsi.md
│   ├── search/
│   │   ├── search.go            # Intelligent fuzzy search engine
│   │   ├── stopwords.go         # Stopwords filtering (EN/IT)
│   │   └── synonyms.go          # Synonym expansion & verb forms
│   └── render/
│       └── render.go            # Glamour-based markdown renderer
├── pkg/
│   ├── cache/
│   │   └── cache.go             # User cheatsheet cache & merge logic
│   ├── config/
│   │   ├── config.go            # YAML configuration management
│   │   └── wizard.go            # Interactive configuration wizard
│   ├── fzf/
│   │   └── fzf.go               # fzf integration for interactive mode
│   └── version/
│       ├── version.go           # Version information
│       └── version.yaml         # Version manifest
├── docs/
│   ├── cheatsheet.html          # Interactive documentation
│   ├── roadmap.html             # Development roadmap
│   └── logo.png                 # Project logo
├── .github/
│   └── workflows/
│       ├── build.yml            # Build & release workflow
│       └── pages.yml            # GitHub Pages deployment
├── go.mod
├── README.md
├── CHANGELOG.md
└── Taskfile.yml                 # Task automation
```

## 📝 Adding Custom Commands

### User Custom Cheatsheets

Create your own cheatsheets in `~/.redfish/{lang}/`:

1. Create a markdown file:

```bash
mkdir -p ~/.redfish/en
nano ~/.redfish/en/mycommands.md
```

2. Follow this format:

```markdown
## command name
**Tags**: tag1, tag2, tag3
**Keywords**: keyword1 keyword2 keyword3
**Short_Description**: Brief one-line description
**Long_Description**: Detailed explanation of what the command does, when to use it, and any important details.

\`\`\`sh
command example
another example
\`\`\`
```

**Note**: `Short_Description` and `Long_Description` are optional but recommended. If omitted, the first text line after keywords will be used as the description.

3. Your commands are immediately available:

```bash
redfish mycommand
```

### Contributing Embedded Cheatsheets

To add commands to the embedded library:

1. Edit files in `internal/parser/data/{lang}/`
2. Follow the enhanced markdown format (see below)
3. Add both English (`en/`) and Italian (`it/`) versions
4. Submit a Pull Request

**Enhanced Format with Descriptions:**

```markdown
## command name
**Tags**: category, tool, action, context
**Keywords**: searchable terms synonyms alternatives
**Short_Description**: One-line summary of what the command does
**Long_Description**: Detailed explanation including when to use it, what it does, important notes, and best practices.

\`\`\`sh
command example
command --with-options
\`\`\`
```

**Guidelines:**
- **Short_Description**: Keep it brief (one line, <80 chars), focus on the main action
- **Long_Description**: Provide context, explain when to use, mention important flags or behaviors
- **Both fields are optional** but highly recommended for better user experience
- **Backward compatible**: Old format (plain text description) still works

**Example:**

```markdown
## docker ps
**Tags**: docker, container, list, running, status
**Keywords**: ps list containers running active status show
**Short_Description**: List running containers
**Long_Description**: Displays all currently running Docker containers with their status, names, and IDs. Use -a flag to show all containers including stopped ones. Essential for monitoring container health and status.

\`\`\`sh
docker ps
docker ps -a
docker ps --format "table {{.Names}}\t{{.Status}}"
\`\`\`
```

## 🔍 How Search Works

### Intelligent Search Pipeline

1. **Tokenization**: Query split into individual terms
2. **Stopwords Filtering**: Common words removed (how, do, i, want, voglio, fare, etc.)
3. **Synonym Expansion**: Terms expanded with related words
   - Verb forms: `listare` → `list`, `show`, `display`, `ls`, `ps`
   - Concepts: `docker` → `container`, `containers`
   - Language-aware (EN/IT)
4. **Weighted Fuzzy Matching**:
   - **Title matches** (highest): 10pt fuzzy, 8pt exact
   - **Tag matches**: 5pt fuzzy, 4pt exact
   - **Keyword matches**: 3pt fuzzy, 2.5pt exact
   - **Description matches**: 1pt
   - **Code matches**: 0.5pt
5. **Coverage Filtering**: Multi-token searches require 70% match threshold
6. **Score Calculation**: Base score × (1 + coverage percentage)
7. **Smart Limiting**: Top 5 results for specific queries, top 10 for general

### Example Query Flow

```
Input:  "voglio listare i docker"
↓
Tokens: ["voglio", "listare", "i", "docker"]
↓
Filtered: ["listare", "docker"]  (stopwords removed)
↓
Expanded: ["listare", "lista", "list", "show", "ls", "ps", "docker", "container", "containers"]
↓
Matched: "docker ps" (score: 148.06, coverage: 100%)
```

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

See [docs/roadmap.html](https://stanzinofree.github.io/redfish/roadmap.html) for the full development roadmap.

**Completed Phases**:
- ✅ **Phase 1**: Core Implementation (Go migration, modular architecture, fuzzy search, Glamour rendering)
- ✅ **Phase 2**: Intelligent Search (NLP queries, stopwords, synonyms, multi-language)
- ✅ **Phase 3**: User Experience (config wizard, custom cheatsheets, secure architecture)

**Current Phase**: Distribution & Quality
- [x] SonarCloud integration
- [x] GitHub Actions CI/CD
- [x] Multi-platform builds
- [x] GitHub Pages documentation
- [ ] Comprehensive unit tests
- [ ] Release automation
- [ ] Homebrew/package managers

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
