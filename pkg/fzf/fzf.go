package fzf

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/stanzinofree/redfish/internal/search"
)

// IsAvailable checks if fzf is installed and available in PATH
func IsAvailable() bool {
	_, err := exec.LookPath("fzf")
	return err == nil
}

// SelectCommand uses fzf to let user interactively select a command
func SelectCommand(results []search.Result) (*search.Result, error) {
	if len(results) == 0 {
		return nil, fmt.Errorf("no results to select from")
	}

	if !IsAvailable() {
		return nil, fmt.Errorf("fzf is not installed. Install it with: brew install fzf")
	}

	// Prepare input for fzf: each line is "index: title - description"
	var lines []string
	for i, r := range results {
		line := fmt.Sprintf("%d: %s - %s", i+1, r.Command.Title, truncate(r.Command.Description, 80))
		lines = append(lines, line)
	}

	// Run fzf
	cmd := exec.Command("fzf",
		"--height=40%",
		"--layout=reverse",
		"--border",
		"--prompt=Select command: ",
		"--preview=echo {}",
		"--preview-window=hidden",
	)

	cmd.Stdin = strings.NewReader(strings.Join(lines, "\n"))
	cmd.Stderr = os.Stderr

	output, err := cmd.Output()
	if err != nil {
		// User cancelled or fzf error
		return nil, err
	}

	// Parse selection
	selection := strings.TrimSpace(string(output))
	if selection == "" {
		return nil, fmt.Errorf("no selection made")
	}

	// Extract index from "index: title - description"
	parts := strings.SplitN(selection, ":", 2)
	if len(parts) < 1 {
		return nil, fmt.Errorf("invalid selection format")
	}

	var selectedIndex int
	if _, err := fmt.Sscanf(parts[0], "%d", &selectedIndex); err != nil {
		return nil, fmt.Errorf("invalid selection index: %w", err)
	}

	// Validate index
	if selectedIndex < 1 || selectedIndex > len(results) {
		return nil, fmt.Errorf("invalid selection index: %d", selectedIndex)
	}

	return &results[selectedIndex-1], nil
}

// truncate truncates a string to maxLen characters
func truncate(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen-3] + "..."
}
