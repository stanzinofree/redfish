# Guida alla Sintassi Markdown

## markdown intestazioni
**Tags**: markdown, intestazioni, titoli, h1, h2, h3
**Keywords**: intestazioni titoli header h1 h2 h3 h4 h5 h6
**Short_Description**: Crea intestazioni
**Long_Description**: Definisce la struttura del documento con sei livelli di intestazione usando simboli cancelletto. H1 è il più grande e importante, H6 il più piccolo. Il numero di cancelletti (1-6) determina il livello intestazione. Essenziale per organizzare la gerarchia dei contenuti e migliorare la leggibilità del documento.

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
**Short_Description**: Formattazione del testo
**Long_Description**: Applica stili al testo usando asterischi o underscore. Singolo asterisco/underscore per corsivo, doppio per grassetto, triplo per entrambi. Tilde creano barrato, backtick creano codice inline. Si combinano per formattazione ricca del testo senza HTML.

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
**Short_Description**: Crea liste
**Long_Description**: Crea liste strutturate usando trattini/asterischi per elenchi puntati o numeri per liste ordinate. Supporta annidamento con indentazione. Le liste di task usano parentesi quadre con x per elementi completati. Essenziali per organizzare informazioni e creare checklist.

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
**Short_Description**: Crea collegamenti
**Long_Description**: Crea link cliccabili con testo e URL. Supporta link inline, link in stile riferimento per riuso, e auto-link per URL semplici. Attributo titolo opzionale appare al passaggio mouse. Lo stile riferimento mantiene pulito il documento quando stesso link appare più volte.

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
**Short_Description**: Incorpora immagini
**Long_Description**: Incorpora immagini usando sintassi simile ai link con punto esclamativo iniziale. Il testo alternativo descrive l'immagine per accessibilità. Supporta sintassi inline e stile riferimento. Può essere avvolto in link per rendere l'immagine cliccabile. Attributo titolo opzionale.

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
**Short_Description**: Blocchi di codice con evidenziazione sintassi
**Long_Description**: Crea blocchi di codice delimitati usando tripli backtick con identificatore linguaggio opzionale per evidenziazione sintassi. Supporta centinaia di linguaggi inclusi JavaScript, Python, Go, Bash, ecc. Essenziale per documentazione tecnica e condivisione snippet di codice.

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
**Short_Description**: Crea citazioni
**Long_Description**: Crea blocchi di testo citato usando simbolo maggiore-di. Supporta citazioni multi-riga, paragrafi multipli e annidamento con simboli aggiuntivi. Comunemente usato per citazioni, callout o enfatizzare informazioni importanti. Può contenere altra formattazione markdown.

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
**Short_Description**: Crea tabelle
**Long_Description**: Crea tabelle dati usando pipe e trattini. Riga intestazione separata da riga allineamento con trattini. Due punti controllano allineamento: sinistra (default), centro (:---:), o destra (---:). Essenziale per presentare dati strutturati in formato leggibile.

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
**Short_Description**: Crea linee orizzontali
**Long_Description**: Crea separatori visivi usando tre o più trattini, asterischi o underscore. Crea linea orizzontale che attraversa la larghezza del documento. Utile per separare sezioni, dividere contenuti o creare gerarchia visiva in documenti lunghi.

```markdown
---
***
___
```

## markdown note a piè di pagina
**Tags**: markdown, note, piè di pagina, riferimenti, citazioni
**Keywords**: note piè pagina riferimenti citazioni apice
**Short_Description**: Aggiungi note a piè di pagina
**Long_Description**: Crea note a piè di pagina con riferimenti numerati nel testo e definizioni in fondo. Usa sintassi parentesi-accento circonflesso-numero per riferimento, stessa sintassi con due punti per definizione. Supporta riferimenti multipli a stessa nota. Utile per citazioni e contesto aggiuntivo.

```markdown
Ecco una frase con una nota[^1].

[^1]: Questo è il contenuto della nota.

Più riferimenti[^nota] alla stessa nota[^nota].

[^nota]: Nota condivisa.
```

## markdown liste definizioni
**Tags**: markdown, definizione, liste, termini, glossario
**Keywords**: definizione liste termini glossario dizionario
**Short_Description**: Crea liste di definizioni
**Long_Description**: Crea liste stile glossario con termini e definizioni. Termine su una riga, definizione su riga successiva iniziando con due punti e spazio. Supporta definizioni multiple per termine. Ideale per glossari, documentazione e FAQ.

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
**Short_Description**: Incorpora HTML grezzo
**Long_Description**: Permette incorporazione HTML grezzo per funzionalità non disponibili in Markdown. Supporta elementi HTML inline e blocco. Utile per stile personalizzato, elementi interattivi come details/summary, o scorciatoie tastiera. Molti processori Markdown permettono passthrough HTML.

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
**Short_Description**: Escape di caratteri speciali
**Long_Description**: Usa backslash per mostrare caratteri speciali Markdown letteralmente senza attivare formattazione. Essenziale quando serve mostrare asterischi, parentesi reali o altri caratteri sintassi. Permette di scrivere sulla sintassi Markdown in documenti Markdown.

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
**Short_Description**: Funzionalità GitHub Flavored Markdown
**Long_Description**: Estensioni Markdown specifiche GitHub incluse liste task, menzioni @, riferimenti issue (#123), shortcode emoji e callout alert. GFM include anche supporto tabelle e linking automatico URL. Queste funzionalità funzionano su GitHub ma potrebbero non funzionare in tutti i processori Markdown.

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
