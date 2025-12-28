package search

// Stopwords are common words that should be filtered from search queries
var stopwordsEN = map[string]bool{
	"a": true, "an": true, "and": true, "are": true, "as": true, "at": true,
	"be": true, "by": true, "for": true, "from": true, "has": true, "he": true,
	"in": true, "is": true, "it": true, "its": true, "of": true, "on": true,
	"that": true, "the": true, "to": true, "was": true, "will": true, "with": true,
	"i": true, "you": true, "me": true, "my": true, "we": true, "us": true,
	"do": true, "does": true, "did": true, "can": true, "could": true, "would": true,
	"should": true, "want": true, "need": true, "how": true, "what": true, "where": true,
	"when": true, "why": true, "which": true, "who": true, "whom": true,
	"make": true, "get": true, "have": true, "this": true, "there": true,
}

var stopwordsIT = map[string]bool{
	"un": true, "una": true, "uno": true, "il": true, "lo": true, "la": true,
	"i": true, "gli": true, "le": true, "di": true, "da": true, "a": true,
	"in": true, "con": true, "su": true, "per": true, "tra": true, "fra": true,
	"e": true, "o": true, "ma": true, "come": true, "quando": true, "dove": true,
	"che": true, "chi": true, "cui": true, "quale": true, "quanto": true,
	"sono": true, "sei": true, "è": true, "siamo": true, "siete": true,
	"ho": true, "hai": true, "ha": true, "abbiamo": true, "avete": true, "hanno": true,
	"voglio": true, "vuoi": true, "vuole": true, "vogliamo": true, "volete": true, "vogliono": true,
	"fare": true, "faccio": true, "fai": true, "fa": true, "facciamo": true, "fate": true, "fanno": true,
	"devo": true, "devi": true, "deve": true, "dobbiamo": true, "dovete": true, "devono": true,
	"posso": true, "puoi": true, "può": true, "possiamo": true, "potete": true, "possono": true,
	"mi": true, "ti": true, "ci": true, "vi": true, "si": true,
	"mio": true, "tuo": true, "suo": true, "nostro": true, "vostro": true, "loro": true,
}

// RemoveStopwords removes common stopwords from tokens based on language
func RemoveStopwords(tokens []string, lang string) []string {
	stopwords := stopwordsEN
	if lang == "it" {
		stopwords = stopwordsIT
	}

	filtered := make([]string, 0, len(tokens))
	for _, token := range tokens {
		if !stopwords[token] {
			filtered = append(filtered, token)
		}
	}

	// If all tokens were stopwords, return original to avoid empty search
	if len(filtered) == 0 {
		return tokens
	}

	return filtered
}

// HasStopwords checks if query contains stopwords (indicates NLP-like query)
func HasStopwords(tokens []string, lang string) bool {
	stopwords := stopwordsEN
	if lang == "it" {
		stopwords = stopwordsIT
	}

	for _, token := range tokens {
		if stopwords[token] {
			return true
		}
	}
	return false
}
