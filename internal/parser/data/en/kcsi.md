# KCSI - Command Cheatsheet Interface

## kcsi search
**Tags**: kcsi, search, query, find, commands
**Keywords**: search query find commands cheatsheet lookup

Search for commands in the cheatsheet

```sh
kcsi git commit
kcsi docker
kcsi "how to merge git branches"
```

## kcsi with fzf
**Tags**: kcsi, fzf, interactive, selection, fuzzy
**Keywords**: fzf interactive selection fuzzy picker

Use interactive fuzzy finder

```sh
kcsi -f git
kcsi --fzf docker
kcsi -f kubernetes
```

## kcsi config
**Tags**: kcsi, config, configuration, wizard, settings
**Keywords**: config configuration wizard settings preferences

Run configuration wizard

```sh
kcsi -c
kcsi --config
```

## kcsi language
**Tags**: kcsi, language, locale, internationalization, i18n
**Keywords**: language locale internationalization i18n translation

Search in different languages

```sh
kcsi -l en git
kcsi -l it docker
kcsi --lang es kubernetes
```

## kcsi version
**Tags**: kcsi, version, info, about, build
**Keywords**: version info about build information

Show version information

```sh
kcsi -v
kcsi --version
```

## kcsi help
**Tags**: kcsi, help, usage, documentation, options
**Keywords**: help usage documentation options flags

Show help and available options

```sh
kcsi -h
kcsi --help
```

## kcsi custom cheatsheets
**Tags**: kcsi, custom, user, personal, cheatsheet
**Keywords**: custom user personal cheatsheet add create own

Add custom cheatsheets

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

Start my application

```sh
myapp start --port 8080
```
EOF
```

## kcsi cache location
**Tags**: kcsi, cache, directory, location, path
**Keywords**: cache directory location path folder storage

Cache and config locations

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

Custom cheatsheet markdown format

```markdown
# Tool Name

## command name
**Tags**: tag1, tag2, tag3, category
**Keywords**: keyword1 keyword2 keyword3 searchable terms

Brief description of what this command does

```sh
command example
command --flag value
command with multiple options
```

## another command
**Tags**: related, tags, here
**Keywords**: more searchable keywords

Another command description

```sh
another example
```
```

## kcsi search features
**Tags**: kcsi, search, nlp, fuzzy, intelligent
**Keywords**: search nlp fuzzy intelligent natural language

Search features and capabilities

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

Build from source

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

Configuration file format (~/.kcsi/config.yaml)

```yaml
# Language preference (ISO 639-1 code)
language: en

# Available languages: en, it, es, fr, de
```

## kcsi ci/cd
**Tags**: kcsi, ci, cd, github, actions, build
**Keywords**: ci cd github actions build pipeline automation

CI/CD and releases

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
