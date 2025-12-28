package config

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// RunWizard runs the interactive configuration wizard
func RunWizard() error {
	reader := bufio.NewReader(os.Stdin)
	
	fmt.Println()
	fmt.Println("╔═══════════════════════════════════════╗")
	fmt.Println("║   🐟 Redfish Configuration Wizard    ║")
	fmt.Println("╚═══════════════════════════════════════╝")
	fmt.Println()

	// Load current config
	currentCfg, err := LoadConfig()
	if err != nil {
		fmt.Printf("Warning: could not load current config: %v\n", err)
		currentCfg = GetDefaultConfig()
	}

	fmt.Println("Configure your preferences:")
	fmt.Println()

	// Language selection
	newCfg := Config{}
	newCfg.Language = promptLanguage(reader, currentCfg.Language)

	// Validate configuration
	if err := ValidateConfig(newCfg); err != nil {
		return fmt.Errorf("invalid configuration: %w", err)
	}

	// Save configuration
	if err := SaveConfig(newCfg); err != nil {
		return fmt.Errorf("saving configuration: %w", err)
	}

	fmt.Println()
	fmt.Println("✓ Configuration saved successfully!")
	fmt.Println()
	fmt.Println("Current settings:")
	fmt.Printf("  Language: %s\n", newCfg.Language)
	fmt.Println()
	
	configPath, _ := GetConfigPath()
	fmt.Printf("Configuration file: %s\n", configPath)
	
	return nil
}

func promptLanguage(reader *bufio.Reader, current string) string {
	fmt.Println("─────────────────────────────────────────")
	fmt.Println("1. Default Language")
	fmt.Println("─────────────────────────────────────────")
	fmt.Println()
	fmt.Println("Select your preferred language for command searches:")
	fmt.Println("  [en] English")
	fmt.Println("  [it] Italian (Italiano)")
	fmt.Println()
	fmt.Printf("Current: %s\n", current)
	fmt.Println()
	fmt.Print("Enter your choice [en/it]: ")

	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("Error reading input, keeping current value: %s\n", current)
		return current
	}

	input = strings.TrimSpace(strings.ToLower(input))
	
	// If empty, keep current
	if input == "" {
		fmt.Printf("Keeping current value: %s\n", current)
		return current
	}

	// Validate input
	if input != "en" && input != "it" {
		fmt.Printf("Invalid choice '%s', keeping current value: %s\n", input, current)
		return current
	}

	return input
}
