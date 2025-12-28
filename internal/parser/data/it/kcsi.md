# KCSI - Command Cheatsheet Interface

## kcsi ricerca
**Tags**: kcsi, ricerca, query, trova, comandi
**Keywords**: ricerca query trova comandi cheatsheet cerca lookup cercare
**Short_Description**: Cerca comandi nel cheatsheet
**Long_Description**: Esegue ricerca intelligente attraverso tutti i cheatsheet usando corrispondenza fuzzy, elaborazione NLP ed espansione sinonimi. Cerca titoli, tag, parole chiave, descrizioni e codice. Restituisce risultati ordinati con comandi più rilevanti per primi. Supporta query in linguaggio naturale.

```sh
kcsi git commit
kcsi docker
kcsi "come fare un merge git"
```

## kcsi con fzf
**Tags**: kcsi, fzf, interattivo, selezione, fuzzy
**Keywords**: fzf interattivo selezione fuzzy picker selettore interazione
**Short_Description**: Usa il selettore fuzzy interattivo
**Long_Description**: Abilita selezione interattiva comandi usando il fuzzy finder fzf. Permette filtraggio e selezione comandi interattivamente con anteprima in tempo reale. Richiede fzf installato. Premi Ctrl+C per annullare, Invio per selezionare. Migliora usabilità per ricerche esplorative.

```sh
kcsi -f git
kcsi --fzf docker
kcsi -f kubernetes
```

## kcsi configurazione
**Tags**: kcsi, config, configurazione, wizard, impostazioni
**Keywords**: config configurazione wizard impostazioni preferenze configurare
**Short_Description**: Esegui wizard di configurazione
**Long_Description**: Lancia wizard di configurazione interattivo per impostare preferenze come lingua predefinita e modalità visualizzazione descrizioni. Impostazioni salvate in ~/.kcsi/config.yaml. Richiede ogni impostazione con valori correnti mostrati. Può essere eseguito più volte per aggiornare preferenze.

```sh
kcsi -c
kcsi --config
```

## kcsi lingua
**Tags**: kcsi, lingua, locale, internazionalizzazione, i18n
**Keywords**: lingua locale internazionalizzazione i18n traduzione linguaggio language
**Short_Description**: Cerca in diverse lingue
**Long_Description**: Cerca cheatsheet nella lingua specificata, sovrascrivendo impostazione file config. Supporta multiple lingue con comandi localizzati, parole chiave e sinonimi. Elaborazione NLP e normalizzazione verbi specifica per lingua. Fallback a inglese se lingua non disponibile.

```sh
kcsi -l en git
kcsi -l it docker
kcsi --lang es kubernetes
```

## kcsi versione
**Tags**: kcsi, versione, info, informazioni, build
**Keywords**: versione info informazioni build dettagli version
**Short_Description**: Mostra informazioni versione
**Long_Description**: Visualizza numero versione corrente, informazioni build e metadata progetto. Include tag versione, URL repository e informazioni licenza. Utile per troubleshooting e verificare installazione.

```sh
kcsi -v
kcsi --version
```

## kcsi aiuto
**Tags**: kcsi, aiuto, uso, documentazione, opzioni
**Keywords**: aiuto help uso documentazione opzioni flag guida
**Short_Description**: Mostra aiuto e opzioni disponibili
**Long_Description**: Visualizza informazioni uso complete includendo flag disponibili, esempi, funzionalità ricerca e formato cheatsheet personalizzati. Spiega algoritmo ricerca, lingue supportate e opzioni configurazione. Riferimento primario per uso tool.

```sh
kcsi -h
kcsi --help
```

## kcsi cheatsheet personalizzati
**Tags**: kcsi, personalizzato, utente, personale, cheatsheet
**Keywords**: personalizzato custom utente personale cheatsheet aggiungi crea proprio
**Short_Description**: Aggiungi cheatsheet personalizzati
**Long_Description**: Crea cheatsheet personali in directory ~/.kcsi/{lang}/ usando formato markdown. Cheatsheet personalizzati uniti con quelli integrati durante ricerca. Supporta stesso formato dei cheatsheet integrati con tag, parole chiave e descrizioni. Ideale per comandi team-specifici o note personali.

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
**Short_Description**: Avvia la mia applicazione
**Long_Description**: Lancia il server applicazione sulla porta specificata con opzioni configurabili.

```sh
miaapp start --port 8080
```
EOF
```

## kcsi posizione cache
**Tags**: kcsi, cache, directory, posizione, percorso
**Keywords**: cache directory posizione percorso cartella storage archiviazione
**Short_Description**: Posizioni cache e configurazione
**Long_Description**: KCSI memorizza configurazione e cheatsheet personalizzati nella directory ~/.kcsi/. File config in config.yaml, cheatsheet lingua-specifici in sottodirectory {lang}/. File markdown personalizzati uniti automaticamente con cheatsheet integrati durante ricerca.

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
**Keywords**: markdown formato sintassi struttura template modello schema
**Short_Description**: Formato markdown per cheatsheet personalizzati
**Long_Description**: Definisce formato standard per cheatsheet personalizzati usando markdown con header speciali. I comandi iniziano con H2 (##), seguiti da Tag, Parole chiave e descrizioni Short/Long opzionali. Blocchi codice mostrano esempi. Formato consistente assicura parsing e ricerca corretti.

```markdown
# Nome Strumento

## nome comando
**Tags**: tag1, tag2, tag3, categoria
**Keywords**: parola1 parola2 parola3 termini ricercabili
**Short_Description**: Breve riassunto una riga
**Long_Description**: Spiegazione dettagliata con contesto e note uso

```sh
esempio comando
comando --flag valore
comando con opzioni multiple
```

## altro comando
**Tags**: tag, correlati, qui
**Keywords**: altre parole chiave ricercabili
**Short_Description**: Altra descrizione breve
**Long_Description**: Più contesto dettagliato

```sh
altro esempio
```
```

## kcsi funzionalità ricerca
**Tags**: kcsi, ricerca, nlp, fuzzy, intelligente
**Keywords**: ricerca nlp fuzzy intelligente linguaggio naturale query
**Short_Description**: Funzionalità e capacità di ricerca
**Long_Description**: KCSI fornisce ricerca avanzata con elaborazione linguaggio naturale, corrispondenza fuzzy per errori battitura, rimozione stopword, espansione sinonimi e logica AND multi-parola. Punteggio ponderato prioritizza corrispondenze titolo su contenuto. Supporta query informali come "come faccio..." e gestisce coniugazioni verbi lingua-specifica.

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
**Keywords**: sviluppo build contribuisci golang codice sorgente compilare
**Short_Description**: Compila dal sorgente
**Long_Description**: Clona repository e compila con toolchain Go. Esegui test con go test, compila binario con go build, o installa globalmente con go install. Richiede Go 1.19 o successivo. Supporta cross-compilazione per multiple piattaforme.

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
**Keywords**: config yaml impostazioni preferenze configurazione settings
**Short_Description**: Formato file di configurazione
**Long_Description**: File configurazione YAML in ~/.kcsi/config.yaml memorizza preferenze utente includendo lingua predefinita e modalità visualizzazione descrizioni. Lingua usa codici ISO 639-1. Modalità descrizione può essere short, long o none. Creato automaticamente da wizard configurazione o editabile manualmente.

```yaml
# Preferenza lingua (codice ISO 639-1)
language: it

# Modalità visualizzazione descrizioni
description_mode: short

# Lingue disponibili: en, it, es, fr, de
# Modalità descrizione disponibili: short, long, none
```

## kcsi ci/cd
**Tags**: kcsi, ci, cd, github, actions, build
**Keywords**: ci cd github actions build pipeline automazione workflow
**Short_Description**: CI/CD e release
**Long_Description**: Workflow GitHub Actions esegue automaticamente test al push, compila binari multi-piattaforma (Linux, macOS, Windows per amd64/arm64), genera checksum SHA256 e crea GitHub release su tag versione. Tag con pattern v* attiva workflow release.

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
