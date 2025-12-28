# KCSI - Command Cheatsheet Interface

## kcsi search
**Tags**: kcsi, search, query, find, commands
**Keywords**: search query find commands cheatsheet lookup
**Short_Description**: Search for commands in the cheatsheet
**Long_Description**: Performs intelligent search across all cheatsheets using fuzzy matching, NLP processing, and synonym expansion. Searches titles, tags, keywords, descriptions, and code. Returns ranked results with most relevant commands first. Supports natural language queries.

```sh
kcsi git commit
kcsi docker
kcsi "how to merge git branches"
```

## kcsi with fzf
**Tags**: kcsi, fzf, interactive, selection, fuzzy
**Keywords**: fzf interactive selection fuzzy picker
**Short_Description**: Use interactive fuzzy finder
**Long_Description**: Enables interactive command selection using fzf fuzzy finder. Allows filtering and selecting commands interactively with real-time preview. Requires fzf to be installed. Press Ctrl+C to cancel, Enter to select. Improves usability for exploratory searches.

```sh
kcsi -f git
kcsi --fzf docker
kcsi -f kubernetes
```

## kcsi config
**Tags**: kcsi, config, configuration, wizard, settings
**Keywords**: config configuration wizard settings preferences
**Short_Description**: Run configuration wizard
**Long_Description**: Launches interactive configuration wizard to set preferences like default language and description display mode. Settings saved to ~/.kcsi/config.yaml. Prompts for each setting with current values shown. Can be run multiple times to update preferences.

```sh
kcsi -c
kcsi --config
```

## kcsi language
**Tags**: kcsi, language, locale, internationalization, i18n
**Keywords**: language locale internationalization i18n translation
**Short_Description**: Search in different languages
**Long_Description**: Searches cheatsheets in specified language, overriding config file setting. Supports multiple languages with localized commands, keywords, and synonyms. Language-specific NLP processing and verb normalization. Falls back to English if language not available.

```sh
kcsi -l en git
kcsi -l it docker
kcsi --lang es kubernetes
```

## kcsi version
**Tags**: kcsi, version, info, about, build
**Keywords**: version info about build information
**Short_Description**: Show version information
**Long_Description**: Displays current version number, build information, and project metadata. Includes version tag, repository URL, and license information. Useful for troubleshooting and verifying installation.

```sh
kcsi -v
kcsi --version
```

## kcsi help
**Tags**: kcsi, help, usage, documentation, options
**Keywords**: help usage documentation options flags
**Short_Description**: Show help and available options
**Long_Description**: Displays comprehensive usage information including available flags, examples, search features, and custom cheatsheet format. Explains search algorithm, supported languages, and configuration options. Primary reference for tool usage.

```sh
kcsi -h
kcsi --help
```

## kcsi custom cheatsheets
**Tags**: kcsi, custom, user, personal, cheatsheet
**Keywords**: custom user personal cheatsheet add create own
**Short_Description**: Add custom cheatsheets
**Long_Description**: Create personal cheatsheets in ~/.kcsi/{lang}/ directories using markdown format. Custom cheatsheets merged with built-in ones during search. Supports same format as built-in cheatsheets with tags, keywords, and descriptions. Ideal for team-specific commands or personal notes.

```sh
# Custom cheatsheets go in ~/.kcsi/{lang}/
# For example: ~/.kcsi/en/mycommands.md

# Create custom cheatsheet
mkdir -p ~/.kcsi/en
cat > ~/.kcsi/en/myapp.md << 'EOF'
# My Application

## myapp start
**Tags**: myapp, start, run
**Keywords**: start run launch
**Short_Description**: Start my application
**Long_Description**: Launches the application server on specified port with configurable options.

```sh
myapp start --port 8080
```
EOF
```

## kcsi cache location
**Tags**: kcsi, cache, directory, location, path
**Keywords**: cache directory location path folder storage
**Short_Description**: Cache and config locations
**Long_Description**: KCSI stores configuration and custom cheatsheets in ~/.kcsi/ directory. Config file at config.yaml, language-specific cheatsheets in {lang}/ subdirectories. Custom markdown files automatically merged with built-in cheatsheets on search.

```sh
# Cache directory
~/.kcsi/

# Config file
~/.kcsi/config.yaml

# Language directories
~/.kcsi/en/
~/.kcsi/it/
~/.kcsi/es/

# Custom cheatsheets (merged with built-in)
~/.kcsi/{lang}/*.md
```

## kcsi markdown format
**Tags**: kcsi, markdown, format, syntax, structure
**Keywords**: markdown format syntax structure template
**Short_Description**: Custom cheatsheet markdown format
**Long_Description**: Defines standard format for custom cheatsheets using markdown with special headers. Commands start with H2 (##), followed by Tags, Keywords, and optional Short/Long descriptions. Code blocks show examples. Consistent format ensures proper parsing and searching.

```markdown
# Tool Name

## command name
**Tags**: tag1, tag2, tag3, category
**Keywords**: keyword1 keyword2 keyword3 searchable terms
**Short_Description**: Brief one-line summary
**Long_Description**: Detailed explanation with context and usage notes

```sh
command example
command --flag value
command with multiple options
```

## another command
**Tags**: related, tags, here
**Keywords**: more searchable keywords
**Short_Description**: Another brief description
**Long_Description**: More detailed context

```sh
another example
```
```

## kcsi search features
**Tags**: kcsi, search, nlp, fuzzy, intelligent
**Keywords**: search nlp fuzzy intelligent natural language
**Short_Description**: Search features and capabilities
**Long_Description**: KCSI provides advanced search with natural language processing, fuzzy matching for typos, stopword removal, synonym expansion, and multi-word AND logic. Weighted scoring prioritizes title matches over content. Supports informal queries like "how do i..." and handles language-specific verb conjugations.

```sh
# Natural language queries
kcsi "how do i list docker containers"
kcsi "voglio fare un merge git"

# Fuzzy matching
kcsi gti comit  # finds "git commit"

# Multi-word AND search
kcsi git merge  # finds commands with both "git" AND "merge"

# Stopwords removal
kcsi "how do i want to see kubernetes pods"  # filters stopwords

# Synonym expansion
kcsi "list containers"  # matches "docker ps" (list→ps, containers→docker)
```

## kcsi development
**Tags**: kcsi, development, build, contribute, golang
**Keywords**: development build contribute golang source code
**Short_Description**: Build from source
**Long_Description**: Clone repository and build with Go toolchain. Run tests with go test, build binary with go build, or install globally with go install. Requires Go 1.19 or later. Supports cross-compilation for multiple platforms.

```sh
# Clone repository
git clone https://github.com/yourusername/kcsi.git
cd kcsi

# Build
go build -o kcsi .

# Run tests
go test ./...

# Install
go install
```

## kcsi configuration file
**Tags**: kcsi, config, yaml, settings, preferences
**Keywords**: config yaml settings preferences configuration
**Short_Description**: Configuration file format
**Long_Description**: YAML configuration file at ~/.kcsi/config.yaml stores user preferences including default language and description display mode. Language uses ISO 639-1 codes. Description mode can be short, long, or none. Created automatically by configuration wizard or manually editable.

```yaml
# Language preference (ISO 639-1 code)
language: en

# Description display mode
description_mode: short

# Available languages: en, it, es, fr, de
# Available description modes: short, long, none
```

## kcsi ci/cd
**Tags**: kcsi, ci, cd, github, actions, build
**Keywords**: ci cd github actions build pipeline automation
**Short_Description**: CI/CD and releases
**Long_Description**: GitHub Actions workflow automatically runs tests on push, builds multi-platform binaries (Linux, macOS, Windows for amd64/arm64), generates SHA256 checksums, and creates GitHub releases on version tags. Tag with v* pattern triggers release workflow.

```sh
# GitHub Actions automatically:
# - Runs tests on push
# - Builds binaries for multiple platforms
# - Creates releases on tags
# - Publishes to GitHub Pages

# Create a release
git tag -a v1.0.0 -m "Release v1.0.0"
git push origin v1.0.0
```
