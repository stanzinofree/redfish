package parser

import (
	"bufio"
	"embed"
	"strings"
)

//go:embed data/*/*.md
var embeddedData embed.FS

// Command represents a command with its metadata
type Command struct {
	Title            string
	Tags             []string
	Keywords         string
	ShortDescription string
	LongDescription  string
	Description      string // Deprecated: use ShortDescription
	Code             string
}

// LoadCommands loads all commands from embedded markdown files for a specific language
func LoadCommands() ([]Command, error) {
	return LoadCommandsForLanguage("en")
}

// LoadCommandsForLanguage loads commands for a specific language
func LoadCommandsForLanguage(lang string) ([]Command, error) {
	var commands []Command

	langDir := "data/" + lang
	entries, err := embeddedData.ReadDir(langDir)
	if err != nil {
		// Fallback to English if language not found
		if lang != "en" {
			return LoadCommandsForLanguage("en")
		}
		return nil, err
	}

	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".md") {
			continue
		}

		content, err := embeddedData.ReadFile(langDir + "/" + entry.Name())
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
	var codeLines []string
	inCodeBlock := false

	scanner := bufio.NewScanner(strings.NewReader(content))
	for scanner.Scan() {
		line := scanner.Text()
		trimmed := strings.TrimSpace(line)

		if isNewCommandHeader(trimmed) {
			current = saveAndStartNewCommand(current, codeLines, &commands)
			current.Title = strings.TrimPrefix(trimmed, "## ")
			codeLines = []string{}
			continue
		}

		if current == nil {
			continue
		}

		inCodeBlock = processLine(trimmed, current, &codeLines, inCodeBlock)
	}

	saveCommand(current, codeLines, &commands)
	return commands
}

func isNewCommandHeader(line string) bool {
	return strings.HasPrefix(line, "## ")
}

func saveAndStartNewCommand(current *Command, codeLines []string, commands *[]Command) *Command {
	saveCommand(current, codeLines, commands)
	return &Command{}
}

func saveCommand(cmd *Command, codeLines []string, commands *[]Command) {
	if cmd != nil {
		cmd.Code = strings.Join(codeLines, "\n")
		*commands = append(*commands, *cmd)
	}
}

func processLine(line string, cmd *Command, codeLines *[]string, inCodeBlock bool) bool {
	if strings.HasPrefix(line, "**Tags**:") {
		cmd.Tags = parseTags(strings.TrimPrefix(line, "**Tags**:"))
		return inCodeBlock
	}

	if strings.HasPrefix(line, "**Keywords**:") {
		cmd.Keywords = strings.TrimSpace(strings.TrimPrefix(line, "**Keywords**:"))
		return inCodeBlock
	}

	if strings.HasPrefix(line, "**Short_Description**:") {
		cmd.ShortDescription = strings.TrimSpace(strings.TrimPrefix(line, "**Short_Description**:"))
		// Also set Description for backward compatibility
		cmd.Description = cmd.ShortDescription
		return inCodeBlock
	}

	if strings.HasPrefix(line, "**Long_Description**:") {
		cmd.LongDescription = strings.TrimSpace(strings.TrimPrefix(line, "**Long_Description**:"))
		return inCodeBlock
	}

	if strings.HasPrefix(line, "```") {
		return !inCodeBlock
	}

	if inCodeBlock {
		*codeLines = append(*codeLines, line)
	} else if line != "" && !strings.HasPrefix(line, "**") {
		// Fallback: if no Short_Description specified, use first line as description
		if cmd.ShortDescription == "" && cmd.Description == "" {
			cmd.Description = line
			cmd.ShortDescription = line
		}
	}

	return inCodeBlock
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
