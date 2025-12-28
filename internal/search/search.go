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
	baseScore := calculateBaseScore(cmd, tokens)
	coverage := calculateCoverage(cmd, tokens)
	return baseScore * (1.0 + coverage)
}

func calculateBaseScore(cmd parser.Command, tokens []string) float64 {
	score := 0.0
	for _, token := range tokens {
		score += scoreTitleMatch(cmd.Title, token)
		score += scoreTagsMatch(cmd.Tags, token)
		score += scoreKeywordsMatch(cmd.Keywords, token)
		score += scoreDescriptionMatch(cmd.Description, token)
		score += scoreCodeMatch(cmd.Code, token)
	}
	return score
}

func scoreTitleMatch(title, token string) float64 {
	titleLower := strings.ToLower(title)
	score := 0.0
	if fuzzyMatch(titleLower, token) {
		score += 10.0
	}
	if strings.Contains(titleLower, token) {
		score += 8.0
	}
	return score
}

func scoreTagsMatch(tags []string, token string) float64 {
	score := 0.0
	for _, tag := range tags {
		tagLower := strings.ToLower(tag)
		if fuzzyMatch(tagLower, token) {
			score += 5.0
		}
		if strings.Contains(tagLower, token) {
			score += 4.0
		}
	}
	return score
}

func scoreKeywordsMatch(keywords, token string) float64 {
	keywordsLower := strings.ToLower(keywords)
	score := 0.0
	if fuzzyMatch(keywordsLower, token) {
		score += 3.0
	}
	if strings.Contains(keywordsLower, token) {
		score += 2.5
	}
	return score
}

func scoreDescriptionMatch(description, token string) float64 {
	if strings.Contains(strings.ToLower(description), token) {
		return 1.0
	}
	return 0.0
}

func scoreCodeMatch(code, token string) float64 {
	if strings.Contains(strings.ToLower(code), token) {
		return 0.5
	}
	return 0.0
}

func calculateCoverage(cmd parser.Command, tokens []string) float64 {
	matched := countMatches(cmd, tokens)
	return float64(matched) / float64(len(tokens))
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
