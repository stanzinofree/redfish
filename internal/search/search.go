package search

import (
	"sort"
	"strings"

	"github.com/stanzinofree/redfish/internal/parser"
)

// Result rappresenta un risultato di ricerca con score
type Result struct {
	Command parser.Command
	Score   float64
}

// Search cerca comandi usando fuzzy matching
func Search(commands []parser.Command, query string) []Result {
	tokens := tokenize(query)
	if len(tokens) == 0 {
		return nil
	}

	var results []Result
	for _, cmd := range commands {
		score := scoreCommand(cmd, tokens)
		if score > 0 {
			results = append(results, Result{
				Command: cmd,
				Score:   score,
			})
		}
	}

	// Ordina per score decrescente
	sort.Slice(results, func(i, j int) bool {
		return results[i].Score > results[j].Score
	})

	return results
}

// tokenize divide la query in token
func tokenize(query string) []string {
	var tokens []string
	for _, token := range strings.Fields(query) {
		token = strings.ToLower(strings.TrimSpace(token))
		if token != "" {
			tokens = append(tokens, token)
		}
	}
	return tokens
}

// scoreCommand calculates the score of a command against search tokens
func scoreCommand(cmd parser.Command, tokens []string) float64 {
	score := 0.0

	for _, token := range tokens {
		// Title match (peso massimo)
		if fuzzyMatch(strings.ToLower(cmd.Title), token) {
			score += 10.0
		}
		if strings.Contains(strings.ToLower(cmd.Title), token) {
			score += 8.0
		}

		// Tags match (peso alto)
		for _, tag := range cmd.Tags {
			tagLower := strings.ToLower(tag)
			if fuzzyMatch(tagLower, token) {
				score += 5.0
			}
			if strings.Contains(tagLower, token) {
				score += 4.0
			}
		}

		// Keywords match (peso medio)
		keywordsLower := strings.ToLower(cmd.Keywords)
		if fuzzyMatch(keywordsLower, token) {
			score += 3.0
		}
		if strings.Contains(keywordsLower, token) {
			score += 2.5
		}

		// Description match (peso basso)
		if strings.Contains(strings.ToLower(cmd.Description), token) {
			score += 1.0
		}

		// Code match (peso minimo)
		if strings.Contains(strings.ToLower(cmd.Code), token) {
			score += 0.5
		}
	}

	// Bonus: coverage (quanti token hanno match)
	matched := countMatches(cmd, tokens)
	coverage := float64(matched) / float64(len(tokens))
	score *= (1.0 + coverage)

	return score
}

// countMatches counts how many tokens match in the command
func countMatches(cmd parser.Command, tokens []string) int {
	matched := 0

	for _, token := range tokens {
		if strings.Contains(strings.ToLower(cmd.Title), token) {
			matched++
			continue
		}

		for _, tag := range cmd.Tags {
			if strings.Contains(strings.ToLower(tag), token) {
				matched++
				break
			}
		}

		if strings.Contains(strings.ToLower(cmd.Keywords), token) {
			matched++
		}
	}

	return matched
}

// fuzzyMatch controlla se tutti i caratteri del pattern appaiono nel testo in ordine
func fuzzyMatch(text, pattern string) bool {
	if len(pattern) == 0 {
		return true
	}
	if len(text) < len(pattern) {
		return false
	}

	textIdx := 0
	patternIdx := 0

	for textIdx < len(text) && patternIdx < len(pattern) {
		if text[textIdx] == pattern[patternIdx] {
			patternIdx++
		}
		textIdx++
	}

	return patternIdx == len(pattern)
}
