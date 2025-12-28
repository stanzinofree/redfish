package parser

import (
	"bufio"
	"embed"
	"strings"
)

//go:embed data/*.md
var embeddedData embed.FS

// Command represents a command with its metadata
type Command struct {
	Title       string
	Tags        []string
	Keywords    string
	Description string
	Code        string
}

// LoadCommands carica tutti i comandi dai file markdown embedded
func LoadCommands() ([]Command, error) {
	var commands []Command

	entries, err := embeddedData.ReadDir("data")
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".md") {
			continue
		}

		content, err := embeddedData.ReadFile("data/" + entry.Name())
		if err != nil {
			return nil, err
		}

		cmds := parseMarkdown(string(content))
		commands = append(commands, cmds...)
	}

	return commands, nil
}

// parseMarkdown parsea un file markdown e estrae i comandi
func parseMarkdown(content string) []Command {
	var commands []Command
	var current *Command

	scanner := bufio.NewScanner(strings.NewReader(content))
	inCodeBlock := false
	var codeLines []string

	for scanner.Scan() {
		line := scanner.Text()
		trimmed := strings.TrimSpace(line)

		// Header H2 = new command
		if strings.HasPrefix(trimmed, "## ") {
			// Save previous command if exists
			if current != nil {
				current.Code = strings.Join(codeLines, "\n")
				commands = append(commands, *current)
			}

			// Start new command
			current = &Command{
				Title: strings.TrimPrefix(trimmed, "## "),
			}
			codeLines = []string{}
			continue
		}

		if current == nil {
			continue
		}

		// Parse metadata
		if strings.HasPrefix(trimmed, "**Tags**:") {
			tagsStr := strings.TrimSpace(strings.TrimPrefix(trimmed, "**Tags**:"))
			current.Tags = parseTags(tagsStr)
			continue
		}

		if strings.HasPrefix(trimmed, "**Keywords**:") {
			current.Keywords = strings.TrimSpace(strings.TrimPrefix(trimmed, "**Keywords**:"))
			continue
		}

		// Code blocks
		if strings.HasPrefix(trimmed, "```") {
			inCodeBlock = !inCodeBlock
			continue
		}

		if inCodeBlock {
			codeLines = append(codeLines, trimmed)
		} else if trimmed != "" && !strings.HasPrefix(trimmed, "**") {
			// Description
			if current.Description != "" {
				current.Description += " "
			}
			current.Description += trimmed
		}
	}

	// Save last command
	if current != nil {
		current.Code = strings.Join(codeLines, "\n")
		commands = append(commands, *current)
	}

	return commands
}

// parseTags divide la stringa di tag in slice
func parseTags(tagsStr string) []string {
	var tags []string
	for _, tag := range strings.Split(tagsStr, ",") {
		tag = strings.TrimSpace(tag)
		if tag != "" {
			tags = append(tags, tag)
		}
	}
	return tags
}

// ParseMarkdownContent espone parseMarkdown per uso esterno (cache)
func ParseMarkdownContent(content string) []Command {
	return parseMarkdown(content)
}
