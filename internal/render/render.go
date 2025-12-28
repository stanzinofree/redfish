package render

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/glamour"
	"github.com/stanzinofree/redfish/internal/parser"
	"github.com/stanzinofree/redfish/internal/search"
)

// RenderResults renderizza i risultati usando Glow/Glamour
func RenderResults(results []search.Result, descMode string) error {
	// Crea renderer Glamour con tema dark
	r, err := glamour.NewTermRenderer(
		glamour.WithAutoStyle(),
		glamour.WithWordWrap(100),
	)
	if err != nil {
		return err
	}

	// Header
	header := fmt.Sprintf("# Found %d command(s)\n\n", len(results))

	var output strings.Builder
	output.WriteString(header)

	for i, result := range results {
		// Format each command as markdown
		md := formatCommandAsMarkdown(result.Command, i+1, descMode)
		output.WriteString(md)

		// Separatore tra comandi
		if i < len(results)-1 {
			output.WriteString("\n---\n\n")
		}
	}

	// Renderizza tutto il markdown
	rendered, err := r.Render(output.String())
	if err != nil {
		return err
	}

	fmt.Print(rendered)
	return nil
}

// formatCommandAsMarkdown converts a command to formatted markdown
func formatCommandAsMarkdown(cmd parser.Command, index int, descMode string) string {
	var md strings.Builder

	// Titolo con numero
	md.WriteString(fmt.Sprintf("## [%d] %s\n\n", index, cmd.Title))

	// Tags
	if len(cmd.Tags) > 0 {
		md.WriteString("**Tags:** ")
		md.WriteString(strings.Join(cmd.Tags, ", "))
		md.WriteString("\n\n")
	}

	// Description based on mode
	switch descMode {
	case "short":
		if cmd.ShortDescription != "" {
			md.WriteString(cmd.ShortDescription)
			md.WriteString("\n\n")
		} else if cmd.Description != "" {
			md.WriteString(cmd.Description)
			md.WriteString("\n\n")
		}
	case "long":
		if cmd.LongDescription != "" {
			md.WriteString(cmd.LongDescription)
			md.WriteString("\n\n")
		} else if cmd.ShortDescription != "" {
			md.WriteString(cmd.ShortDescription)
			md.WriteString("\n\n")
		} else if cmd.Description != "" {
			md.WriteString(cmd.Description)
			md.WriteString("\n\n")
		}
	case "none":
		// Don't show any description
	default:
		// Default to short mode
		if cmd.ShortDescription != "" {
			md.WriteString(cmd.ShortDescription)
			md.WriteString("\n\n")
		} else if cmd.Description != "" {
			md.WriteString(cmd.Description)
			md.WriteString("\n\n")
		}
	}

	// Code
	if cmd.Code != "" {
		md.WriteString("```sh\n")
		md.WriteString(cmd.Code)
		md.WriteString("\n```\n\n")
	}

	return md.String()
}
