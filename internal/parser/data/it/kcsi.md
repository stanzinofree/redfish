# KCSI - Command Cheatsheet Interface

## kcsi ricerca
**Tags**: kcsi, ricerca, query, trova, comandi
**Keywords**: ricerca query trova comandi cheatsheet cerca lookup

Cerca comandi nel cheatsheet

```sh
kcsi git commit
kcsi docker
kcsi "come fare un merge git"
```

## kcsi con fzf
**Tags**: kcsi, fzf, interattivo, selezione, fuzzy
**Keywords**: fzf interattivo selezione fuzzy picker selettore

Usa il selettore fuzzy interattivo

```sh
kcsi -f git
kcsi --fzf docker
kcsi -f kubernetes
```

## kcsi configurazione
**Tags**: kcsi, config, configurazione, wizard, impostazioni
**Keywords**: config configurazione wizard impostazioni preferenze

Esegui wizard di configurazione

```sh
kcsi -c
kcsi --config
```

## kcsi lingua
**Tags**: kcsi, lingua, locale, internazionalizzazione, i18n
**Keywords**: lingua locale internazionalizzazione i18n traduzione linguaggio

Cerca in diverse lingue

```sh
kcsi -l en git
kcsi -l it docker
kcsi --lang es kubernetes
```

## kcsi versione
**Tags**: kcsi, versione, info, informazioni, build
**Keywords**: versione info informazioni build dettagli

Mostra informazioni versione

```sh
kcsi -v
kcsi --version
```

## kcsi aiuto
**Tags**: kcsi, aiuto, uso, documentazione, opzioni
**Keywords**: aiuto help uso documentazione opzioni flag

Mostra aiuto e opzioni disponibili

```sh
kcsi -h
kcsi --help
```

## kcsi cheatsheet personalizzati
**Tags**: kcsi, personalizzato, utente, personale, cheatsheet
**Keywords**: personalizzato custom utente personale cheatsheet aggiungi crea proprio

Aggiungi cheatsheet personalizzati

```sh
# I cheatsheet personalizzati vanno in ~/.kcsi/{lang}/
# Ad esempio: ~/.kcsi/it/mieicomandi.md

# Crea cheatsheet personalizzato
mkdir -p ~/.kcsi/it
cat > ~/.kcsi/it/miaapp.md << 'EOF'
# Mia Applicazione

## miaapp start
**Tags**: miaapp, avvia, esegui
**Keywords**: avvia esegui lancia start run

Avvia la mia applicazione

```sh
miaapp start --port 8080
```
EOF
```

## kcsi posizione cache
**Tags**: kcsi, cache, directory, posizione, percorso
**Keywords**: cache directory posizione percorso cartella storage

Posizioni cache e configurazione

```sh
# Directory cache
~/.kcsi/

# File di configurazione
~/.kcsi/config.yaml

# Directory lingue
~/.kcsi/en/
~/.kcsi/it/
~/.kcsi/es/

# Cheatsheet personalizzati (uniti con quelli integrati)
~/.kcsi/{lang}/*.md
```

## kcsi formato markdown
**Tags**: kcsi, markdown, formato, sintassi, struttura
**Keywords**: markdown formato sintassi struttura template modello

Formato markdown per cheatsheet personalizzati

```markdown
# Nome Strumento

## nome comando
**Tags**: tag1, tag2, tag3, categoria
**Keywords**: parola1 parola2 parola3 termini ricercabili

Breve descrizione di cosa fa questo comando

```sh
esempio comando
comando --flag valore
comando con opzioni multiple
```

## altro comando
**Tags**: tag, correlati, qui
**Keywords**: altre parole chiave ricercabili

Descrizione altro comando

```sh
altro esempio
```
```

## kcsi funzionalità ricerca
**Tags**: kcsi, ricerca, nlp, fuzzy, intelligente
**Keywords**: ricerca nlp fuzzy intelligente linguaggio naturale

Funzionalità e capacità di ricerca

```sh
# Query in linguaggio naturale
kcsi "come faccio a listare i container docker"
kcsi "voglio fare un merge git"

# Corrispondenza fuzzy
kcsi gti comit  # trova "git commit"

# Ricerca multi-parola con AND
kcsi git merge  # trova comandi con sia "git" CHE "merge"

# Rimozione stopword
kcsi "come faccio voglio vedere i pod kubernetes"  # filtra stopword

# Espansione sinonimi
kcsi "lista container"  # trova "docker ps" (lista→ps, container→docker)
```

## kcsi sviluppo
**Tags**: kcsi, sviluppo, build, contribuisci, golang
**Keywords**: sviluppo build contribuisci golang codice sorgente

Compila dal sorgente

```sh
# Clona repository
git clone https://github.com/yourusername/kcsi.git
cd kcsi

# Compila
go build -o kcsi .

# Esegui test
go test ./...

# Installa
go install
```

## kcsi file configurazione
**Tags**: kcsi, config, yaml, impostazioni, preferenze
**Keywords**: config yaml impostazioni preferenze configurazione

Formato file di configurazione (~/.kcsi/config.yaml)

```yaml
# Preferenza lingua (codice ISO 639-1)
language: it

# Lingue disponibili: en, it, es, fr, de
```

## kcsi ci/cd
**Tags**: kcsi, ci, cd, github, actions, build
**Keywords**: ci cd github actions build pipeline automazione

CI/CD e release

```sh
# GitHub Actions automaticamente:
# - Esegue test al push
# - Compila binari per multiple piattaforme
# - Crea release sui tag
# - Pubblica su GitHub Pages

# Crea una release
git tag -a v1.0.0 -m "Release v1.0.0"
git push origin v1.0.0
```
