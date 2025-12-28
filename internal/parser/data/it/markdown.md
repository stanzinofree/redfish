# Guida alla Sintassi Markdown

## markdown intestazioni
**Tags**: markdown, intestazioni, titoli, h1, h2, h3
**Keywords**: intestazioni titoli header h1 h2 h3 h4 h5 h6

Crea intestazioni

```markdown
# Intestazione H1
## Intestazione H2
### Intestazione H3
#### Intestazione H4
##### Intestazione H5
###### Intestazione H6
```

## markdown enfasi
**Tags**: markdown, enfasi, grassetto, corsivo, formattazione
**Keywords**: enfasi grassetto corsivo strong formattazione testo stile

Formattazione del testo

```markdown
**testo grassetto**
*testo corsivo*
***grassetto e corsivo***
~~testo barrato~~
`codice inline`
```

## markdown liste
**Tags**: markdown, liste, ordinate, non-ordinate, elenchi
**Keywords**: liste ordinate non-ordinate elenchi puntati numerati elementi

Crea liste

```markdown
# Lista non ordinata
- Elemento 1
- Elemento 2
  - Elemento annidato
  - Altro elemento annidato

# Lista ordinata
1. Primo elemento
2. Secondo elemento
   1. Numerato annidato
   2. Altro numerato annidato

# Lista di task
- [x] Task completato
- [ ] Task incompleto
```

## markdown link
**Tags**: markdown, link, collegamenti, url, riferimenti
**Keywords**: link collegamenti ipertesti url riferimenti href ancora

Crea collegamenti

```markdown
[Testo del link](https://example.com)
[Link con titolo](https://example.com "Titolo del Link")

# Link in stile riferimento
[Testo del link][riferimento]
[riferimento]: https://example.com

# Auto-link
<https://example.com>
<email@example.com>
```

## markdown immagini
**Tags**: markdown, immagini, foto, media, grafica
**Keywords**: immagini foto media grafica picture embed incorpora

Incorpora immagini

```markdown
![Testo alternativo](immagine.jpg)
![Testo alternativo](immagine.jpg "Titolo Immagine")

# Immagini in stile riferimento
![Testo alternativo][riferimento-immagine]
[riferimento-immagine]: immagine.jpg

# Con link
[![Testo alternativo](immagine.jpg)](https://example.com)
```

## markdown blocchi codice
**Tags**: markdown, codice, sintassi, evidenziazione, fenced, blocchi
**Keywords**: codice sintassi evidenziazione fenced blocchi backtick triplo

Blocchi di codice con evidenziazione sintassi

````markdown
```javascript
function ciao() {
  console.log("Ciao, Mondo!");
}
```

```python
def ciao():
    print("Ciao, Mondo!")
```

```bash
echo "Ciao, Mondo!"
```
````

## markdown citazioni
**Tags**: markdown, citazioni, blockquote, quote
**Keywords**: citazioni blockquote quote riferimento indentato

Crea citazioni

```markdown
> Questa è una citazione
> Può estendersi su più righe

> Citazione multi-paragrafo
>
> Secondo paragrafo

> Citazione annidata
>> Annidata più in profondità
```

## markdown tabelle
**Tags**: markdown, tabelle, griglia, dati, tabulare
**Keywords**: tabelle griglia dati tabulare righe colonne allineamento

Crea tabelle

```markdown
| Intestazione 1 | Intestazione 2 | Intestazione 3 |
|----------------|----------------|----------------|
| Riga 1         | Dato           | Altro          |
| Riga 2         | Dato           | Altro          |

# Allineamento
| Sinistra | Centro | Destra |
|:---------|:------:|-------:|
| S        |   C    |      D |
```

## markdown linea orizzontale
**Tags**: markdown, orizzontale, linea, separatore, divisore
**Keywords**: orizzontale linea separatore divisore hr interruzione

Crea linee orizzontali

```markdown
---
***
___
```

## markdown note a piè di pagina
**Tags**: markdown, note, piè di pagina, riferimenti, citazioni
**Keywords**: note piè pagina riferimenti citazioni apice

Aggiungi note a piè di pagina

```markdown
Ecco una frase con una nota[^1].

[^1]: Questo è il contenuto della nota.

Più riferimenti[^nota] alla stessa nota[^nota].

[^nota]: Nota condivisa.
```

## markdown liste definizioni
**Tags**: markdown, definizione, liste, termini, glossario
**Keywords**: definizione liste termini glossario dizionario

Crea liste di definizioni

```markdown
Termine 1
: Definizione per il termine 1

Termine 2
: Prima definizione per il termine 2
: Seconda definizione per il termine 2
```

## markdown html
**Tags**: markdown, html, raw, inline, incorporato
**Keywords**: html raw inline incorporato tag elementi

Incorpora HTML grezzo

```markdown
<div style="color: red;">
  Questo è testo rosso
</div>

<details>
  <summary>Clicca per espandere</summary>
  Contenuto nascosto qui
</details>

<kbd>Ctrl</kbd> + <kbd>C</kbd>
```

## markdown escape
**Tags**: markdown, escape, speciali, caratteri, backslash
**Keywords**: escape speciali caratteri backslash letterale simboli

Escape di caratteri speciali

```markdown
\* Non un punto elenco
\# Non un'intestazione
\[Non un link\]
\`Non codice\`

Usa backslash \ prima di:
* \* \_ \{ \} \[ \] \( \)
* \# \+ \- \. \!
* \| \~
```

## markdown estensioni github
**Tags**: markdown, github, gfm, estensioni, flavored
**Keywords**: github gfm estensioni flavored sintassi speciale

Funzionalità GitHub Flavored Markdown

```markdown
# Liste di task
- [x] Completato
- [ ] Incompleto

# Menzioni username
@username

# Riferimenti issue
#123
owner/repo#123

# Emoji
:smile: :rocket: :heart:

# Tabelle (semplificate)
Intestazione | Intestazione
------------ | ------------
Cella        | Cella

# Alert (GitHub)
> [!NOTE]
> Informazioni utili

> [!WARNING]
> Contenuto critico

> [!IMPORTANT]
> Informazioni chiave
```
