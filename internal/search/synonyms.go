package search

import "strings"

// Synonyms and verb forms mapping
var synonymsEN = map[string][]string{
	"list":      {"listing", "show", "showing", "display", "displaying", "view", "viewing", "ls", "ps"},
	"create":    {"creating", "make", "making", "add", "adding", "new"},
	"delete":    {"deleting", "remove", "removing", "rm", "destroy", "destroying"},
	"run":       {"running", "execute", "executing", "start", "starting", "launch", "launching"},
	"stop":      {"stopping", "kill", "killing", "terminate", "terminating"},
	"update":    {"updating", "modify", "modifying", "change", "changing", "edit", "editing"},
	"get":       {"getting", "fetch", "fetching", "retrieve", "retrieving", "obtain", "obtaining"},
	"push":      {"pushing", "upload", "uploading", "send", "sending"},
	"pull":      {"pulling", "download", "downloading", "fetch", "fetching"},
	"merge":     {"merging", "combine", "combining", "join", "joining"},
	"reset":     {"resetting", "revert", "reverting", "undo", "undoing"},
	"commit":    {"committing", "save", "saving", "record", "recording"},
	"build":     {"building", "compile", "compiling", "construct", "constructing"},
	"install":   {"installing", "setup", "setting"},
	"docker":    {"container", "containers"},
	"container": {"docker", "containers"},
}

var synonymsIT = map[string][]string{
	"listare":      {"lista", "list", "elencare", "elenco", "mostrare", "mostra", "visualizzare", "visualizza", "ls", "ps"},
	"vedere":       {"vedi", "vedere", "view", "show", "mostra", "mostrare"},
	"creare":       {"crea", "create", "nuovo", "nuova", "aggiungere", "aggiungi"},
	"eliminare":    {"elimina", "delete", "rimuovere", "rimuovi", "cancellare", "cancella"},
	"eseguire":     {"esegui", "run", "avviare", "avvia", "lanciare", "lancia"},
	"fermare":      {"ferma", "stop", "terminare", "termina"},
	"aggiornare":   {"aggiorna", "update", "modificare", "modifica"},
	"ottenere":     {"ottieni", "get", "prendere", "prendi"},
	"caricare":     {"carica", "push", "upload"},
	"scaricare":    {"scarica", "pull", "download"},
	"unire":        {"unisci", "merge", "combinare", "combina"},
	"ripristinare": {"ripristina", "reset", "annullare", "annulla"},
	"salvare":      {"salva", "commit", "save"},
	"compilare":    {"compila", "build", "costruire", "costruisci"},
	"installare":   {"installa", "install", "setup"},
	"docker":       {"container", "contenitore", "contenitori", "containers"},
	"container":    {"docker", "contenitore", "contenitori"},
	"contenitore":  {"docker", "container", "containers"},
}

// ExpandWithSynonyms expands tokens with their synonyms
func ExpandWithSynonyms(tokens []string, lang string) []string {
	synonyms := synonymsEN
	if lang == "it" {
		synonyms = synonymsIT
	}

	expanded := make([]string, 0, len(tokens)*2)

	for _, token := range tokens {
		// Add original token
		expanded = append(expanded, token)

		// Check if token matches any base word or its synonyms
		for baseWord, variants := range synonyms {
			// Check if token is the base word
			if token == baseWord {
				expanded = append(expanded, variants...)
				break
			}

			// Check if token is one of the variants
			for _, variant := range variants {
				if token == variant {
					// Add base word and all other variants
					expanded = append(expanded, baseWord)
					for _, v := range variants {
						if v != variant {
							expanded = append(expanded, v)
						}
					}
					break
				}
			}
		}
	}

	// Remove duplicates
	return uniqueTokens(expanded)
}

// uniqueTokens removes duplicate tokens
func uniqueTokens(tokens []string) []string {
	seen := make(map[string]bool)
	result := make([]string, 0, len(tokens))

	for _, token := range tokens {
		if !seen[token] {
			seen[token] = true
			result = append(result, token)
		}
	}

	return result
}

// NormalizeVerb tries to normalize Italian/English verb forms to base form
func NormalizeVerb(word string, lang string) string {
	word = strings.ToLower(word)

	if lang == "it" {
		// Italian verb endings to base forms
		// Remove common verb endings
		if strings.HasSuffix(word, "are") || strings.HasSuffix(word, "ire") || strings.HasSuffix(word, "ere") {
			return word // Already in infinitive
		}

		// Simple normalization for common patterns
		replacements := map[string]string{
			"lista":    "listare",
			"elenca":   "listare",
			"mostra":   "mostrare",
			"crea":     "creare",
			"elimina":  "eliminare",
			"esegui":   "eseguire",
			"ferma":    "fermare",
			"aggiorna": "aggiornare",
			"ottieni":  "ottenere",
			"carica":   "caricare",
			"scarica":  "scaricare",
			"unisci":   "unire",
			"salva":    "salvare",
			"compila":  "compilare",
			"installa": "installare",
		}

		if normalized, ok := replacements[word]; ok {
			return normalized
		}
	} else {
		// English verb forms
		replacements := map[string]string{
			"listing":  "list",
			"showing":  "show",
			"creating": "create",
			"making":   "make",
			"deleting": "delete",
			"removing": "remove",
			"running":  "run",
			"stopping": "stop",
			"updating": "update",
			"getting":  "get",
			"pushing":  "push",
			"pulling":  "pull",
			"merging":  "merge",
			"building": "build",
		}

		if normalized, ok := replacements[word]; ok {
			return normalized
		}
	}

	return word
}
