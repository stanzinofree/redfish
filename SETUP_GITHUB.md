# GitHub Repository Setup Guide

Questa guida ti aiuterà a configurare il repository GitHub per Redfish con tutte le integrazioni necessarie.

## 📋 Checklist Setup Repository

### 1. Creare e Inizializzare il Repository

```bash
# Inizializza git se non già fatto
cd /Users/alessandro/Documents/PLAY/REPOS/CSI/redifsh
git init
git add .
git commit -m "Initial commit: Redfish v0.1.0"

# Aggiungi remote
git remote add origin https://github.com/stanzinofree/redfish.git
git branch -M main
git push -u origin main
```

### 2. Configurare Branch Protection su GitHub

1. Vai su: **Settings** → **Branches** → **Branch protection rules**
2. Clicca **Add rule**
3. Branch name pattern: `main`
4. Abilita:
   - ✅ **Require a pull request before merging**
   - ✅ **Require approvals** (almeno 1)
   - ✅ **Dismiss stale pull request approvals when new commits are pushed**
   - ✅ **Require status checks to pass before merging**
   - ✅ **Require branches to be up to date before merging**
   - ✅ **Require conversation resolution before merging**
   - ✅ **Do not allow bypassing the above settings**
5. Salva le modifiche

### 3. Abilitare GitHub Actions

Le GitHub Actions sono già configurate in `.github/workflows/ci.yml`.

**Verifica che siano abilitate:**
1. Vai su: **Settings** → **Actions** → **General**
2. In "Actions permissions", seleziona: **Allow all actions and reusable workflows**
3. In "Workflow permissions", seleziona: **Read and write permissions**
4. Abilita: **Allow GitHub Actions to create and approve pull requests**

### 4. Configurare SonarCloud

1. **Vai su [SonarCloud](https://sonarcloud.io/)**
2. **Login con GitHub**
3. **Import organization**: Seleziona `stanzinofree`
4. **Analyze new project**: Seleziona `redfish`
5. **Choose analysis method**: Seleziona **GitHub Actions**
6. **Copy** il token fornito

**Aggiungi il token ai GitHub Secrets:**
1. Vai su: **Settings** → **Secrets and variables** → **Actions**
2. Clicca **New repository secret**
3. Name: `SONAR_TOKEN`
4. Value: [incolla il token da SonarCloud]
5. Clicca **Add secret**

**Crea il file di configurazione SonarCloud:**

```bash
# Crea sonar-project.properties
cat > sonar-project.properties << 'EOF'
sonar.projectKey=stanzinofree_redfish
sonar.organization=stanzinofree
sonar.sources=.
sonar.exclusions=**/*_test.go,**/vendor/**,**/testdata/**
sonar.tests=.
sonar.test.inclusions=**/*_test.go
sonar.go.coverage.reportPaths=coverage.out
EOF

git add sonar-project.properties
git commit -m "Add SonarCloud configuration"
git push
```

**Aggiorna `.github/workflows/ci.yml` per includere SonarCloud:**

```yaml
# Aggiungi questo job dopo gli altri
  sonarcloud:
    name: SonarCloud Analysis
    runs-on: ubuntu-latest
    needs: test
    steps:
      - name: Checkout code
        uses: actions/checkout@v4
        with:
          fetch-depth: 0

      - name: Download coverage
        uses: actions/download-artifact@v3
        with:
          name: coverage
          path: .

      - name: SonarCloud Scan
        uses: SonarSource/sonarcloud-github-action@master
        env:
          GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}
          SONAR_TOKEN: ${{ secrets.SONAR_TOKEN }}
```

### 5. Configurare Dependabot

GitHub Dependabot è già incluso nel repository.

**Crea il file di configurazione:**

```bash
# Crea .github/dependabot.yml
mkdir -p .github
cat > .github/dependabot.yml << 'EOF'
version: 2
updates:
  - package-ecosystem: "gomod"
    directory: "/"
    schedule:
      interval: "weekly"
    open-pull-requests-limit: 10
    reviewers:
      - "stanzinofree"
    labels:
      - "dependencies"
      - "go"

  - package-ecosystem: "github-actions"
    directory: "/"
    schedule:
      interval: "weekly"
    open-pull-requests-limit: 5
    reviewers:
      - "stanzinofree"
    labels:
      - "dependencies"
      - "ci"
EOF

git add .github/dependabot.yml
git commit -m "Add Dependabot configuration"
git push
```

### 6. Configurare GitHub Pages

1. Vai su: **Settings** → **Pages**
2. Source: **Deploy from a branch**
3. Branch: `main` / `docs`
4. Salva

I file `docs/cheatsheet.html` e `docs/roadmap.html` saranno accessibili a:
- https://stanzinofree.github.io/redfish/cheatsheet.html
- https://stanzinofree.github.io/redfish/roadmap.html

### 7. Configurare CodeQL (Advanced Security)

Se hai GitHub Advanced Security, aggiungi CodeQL:

```bash
# Crea .github/workflows/codeql.yml
cat > .github/workflows/codeql.yml << 'EOF'
name: "CodeQL"

on:
  push:
    branches: [ main, develop ]
  pull_request:
    branches: [ main ]
  schedule:
    - cron: '0 6 * * 1'

jobs:
  analyze:
    name: Analyze
    runs-on: ubuntu-latest
    permissions:
      actions: read
      contents: read
      security-events: write

    strategy:
      fail-fast: false
      matrix:
        language: [ 'go' ]

    steps:
    - name: Checkout repository
      uses: actions/checkout@v4

    - name: Initialize CodeQL
      uses: github/codeql-action/init@v3
      with:
        languages: ${{ matrix.language }}

    - name: Autobuild
      uses: github/codeql-action/autobuild@v3

    - name: Perform CodeQL Analysis
      uses: github/codeql-action/analyze@v3
EOF

git add .github/workflows/codeql.yml
git commit -m "Add CodeQL security scanning"
git push
```

### 8. Aggiungere LICENSE

```bash
# Crea MIT License
cat > LICENSE << 'EOF'
MIT License

Copyright (c) 2024 Alessandro Middei

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
EOF

git add LICENSE
git commit -m "Add MIT License"
git push
```

### 9. Configurare Repository Settings

**General Settings:**
1. **Settings** → **General**
2. Features:
   - ✅ **Issues**
   - ✅ **Projects**
   - ✅ **Discussions** (opzionale)
   - ✅ **Wiki** (opzionale)
3. Pull Requests:
   - ✅ **Allow squash merging**
   - ✅ **Allow merge commits**
   - ✅ **Allow rebase merging**
   - ✅ **Automatically delete head branches**

**Topics:**
1. Vai su repository home
2. Clicca su ⚙️ accanto a "About"
3. Aggiungi topics: `cli`, `go`, `golang`, `cheatsheet`, `fuzzy-search`, `command-line`, `markdown`, `glamour`

### 10. Creare il Logo

```bash
# Placeholder per il logo
# Crea un'immagine 200x200 px e salvala come docs/logo.png
# Puoi usare questo placeholder temporaneo:
echo "TODO: Aggiungi logo.png in docs/logo.png" > docs/logo.md
```

### 11. Configurare Issue Templates

```bash
# Crea template per bug report
mkdir -p .github/ISSUE_TEMPLATE
cat > .github/ISSUE_TEMPLATE/bug_report.md << 'EOF'
---
name: Bug Report
about: Create a report to help us improve
title: '[BUG] '
labels: bug
assignees: ''
---

**Describe the bug**
A clear and concise description of what the bug is.

**To Reproduce**
Steps to reproduce the behavior:
1. Run command '...'
2. See error

**Expected behavior**
A clear and concise description of what you expected to happen.

**Screenshots**
If applicable, add screenshots to help explain your problem.

**Environment:**
 - OS: [e.g. macOS, Linux, Windows]
 - Redfish Version: [e.g. 0.1.0]
 - Go Version: [e.g. 1.21.0]

**Additional context**
Add any other context about the problem here.
EOF

# Crea template per feature request
cat > .github/ISSUE_TEMPLATE/feature_request.md << 'EOF'
---
name: Feature Request
about: Suggest an idea for this project
title: '[FEATURE] '
labels: enhancement
assignees: ''
---

**Is your feature request related to a problem? Please describe.**
A clear and concise description of what the problem is.

**Describe the solution you'd like**
A clear and concise description of what you want to happen.

**Describe alternatives you've considered**
A clear and concise description of any alternative solutions or features you've considered.

**Additional context**
Add any other context or screenshots about the feature request here.
EOF

git add .github/ISSUE_TEMPLATE/
git commit -m "Add issue templates"
git push
```

### 12. Prima PR usando Task

```bash
# Crea una feature branch
task branch NAME=feature/initial-setup

# Fai le tue modifiche
git add .
git commit -m "feat: initial repository setup"

# Crea PR
task pr BRANCH=feature/initial-setup TITLE="Initial repository setup" DESC="Complete initial setup with CI/CD, documentation, and quality gates"
```

## ✅ Verifica Finale

Dopo aver completato tutti i passaggi, verifica:

- [ ] Il repository è su https://github.com/stanzinofree/redfish
- [ ] CI passa su ogni push
- [ ] SonarCloud mostra quality gate
- [ ] Dependabot è attivo
- [ ] GitHub Pages funziona
- [ ] Branch protection su `main` è attiva
- [ ] Badges nel README sono funzionanti
- [ ] I template per issues funzionano

## 🚀 Workflow Quotidiano

```bash
# Creare nuova feature
task branch NAME=feature/new-command

# Fare modifiche e testare
task check

# Creare PR
task pr BRANCH=feature/new-command TITLE="Add new command" DESC="Description"

# Aggiornare PR
task pr:update MESSAGE="Address review comments"
```

## 📚 Risorse

- [SonarCloud Documentation](https://docs.sonarcloud.io/)
- [GitHub Actions Documentation](https://docs.github.com/en/actions)
- [Dependabot Documentation](https://docs.github.com/en/code-security/dependabot)
- [Task Documentation](https://taskfile.dev/)
