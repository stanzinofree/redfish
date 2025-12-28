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
	synonyms := getSynonymMap(lang)
	expanded := make([]string, 0, len(tokens)*2)

	for _, token := range tokens {
		expanded = append(expanded, token)
		expanded = append(expanded, expandToken(token, synonyms)...)
	}

	return uniqueTokens(expanded)
}

// getSynonymMap returns the appropriate synonym map for the language
func getSynonymMap(lang string) map[string][]string {
	if lang == "it" {
		return synonymsIT
	}
	return synonymsEN
}

// expandToken expands a single token with its synonyms
func expandToken(token string, synonyms map[string][]string) []string {
	var result []string

	// Check if token is a base word
	if variants, found := synonyms[token]; found {
		result = append(result, variants...)
		return result
	}

	// Check if token is a variant
	for baseWord, variants := range synonyms {
		if containsString(variants, token) {
			result = append(result, baseWord)
			result = append(result, filterOutString(variants, token)...)
			return result
		}
	}

	return result
}

// containsString checks if a slice contains a string
func containsString(slice []string, str string) bool {
	for _, s := range slice {
		if s == str {
			return true
		}
	}
	return false
}

// filterOutString returns slice without the specified string
func filterOutString(slice []string, exclude string) []string {
	result := make([]string, 0, len(slice))
	for _, s := range slice {
		if s != exclude {
			result = append(result, s)
		}
	}
	return result
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
func NormalizeVerb(word, lang string) string {
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
