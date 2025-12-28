package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/stanzinofree/redfish/internal/parser"
	"github.com/stanzinofree/redfish/internal/render"
	"github.com/stanzinofree/redfish/internal/search"
	"github.com/stanzinofree/redfish/pkg/cache"
	"github.com/stanzinofree/redfish/pkg/version"
)

var (
	showVersion = flag.Bool("v", false, "Show version information")
	reloadCache = flag.Bool("r", false, "Reload cache from ~/.redfish custom cheatsheets")
	_           = flag.String("l", "en", "Language for search (currently only 'en' supported)")
	showHelp    = flag.Bool("h", false, "Show help message")
)

func main() {
	// Parse flags
	flag.Parse()

	// Handle flags
	if *showHelp {
		printHelp()
		os.Exit(0)
	}

	if *showVersion {
		fmt.Println(version.GetVersionInfo())
		os.Exit(0)
	}

	if *reloadCache {
		if err := rebuildCache(); err != nil {
			fmt.Fprintf(os.Stderr, "Error rebuilding cache: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("✓ Cache rebuilt successfully")
		os.Exit(0)
	}

	// Get query from remaining arguments
	args := flag.Args()
	if len(args) == 0 {
		printUsage()
		os.Exit(1)
	}

	query := strings.Join(args, " ")

	// Load commands (embedded + cached custom)
	commands, err := loadAllCommands()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading commands: %v\n", err)
		os.Exit(1)
	}

	// Search
	results := search.Search(commands, query)

	if len(results) == 0 {
		fmt.Printf("No commands found for: %s\n", query)
		fmt.Println("\nTip: Try different keywords or use -h for help")
		os.Exit(0)
	}

	// Render results
	if err := render.RenderResults(results); err != nil {
		fmt.Fprintf(os.Stderr, "Error rendering results: %v\n", err)
		os.Exit(1)
	}
}

func loadAllCommands() ([]parser.Command, error) {
	// Load embedded commands
	embedded, err := parser.LoadCommands()
	if err != nil {
		return nil, err
	}

	// Try to load custom commands from cache
	custom, err := cache.LoadCustomCheatsheets()
	if err != nil {
		// Non-fatal: just use embedded commands
		return embedded, nil
	}

	// Merge embedded and custom
	return append(embedded, custom...), nil
}

func rebuildCache() error {
	// Clear existing cache
	if err := cache.ClearCache(); err != nil {
		return fmt.Errorf("clearing cache: %w", err)
	}

	// Load custom cheatsheets
	custom, err := cache.LoadCustomCheatsheets()
	if err != nil {
		return fmt.Errorf("loading custom cheatsheets: %w", err)
	}

	if len(custom) == 0 {
		cacheDir, _ := cache.GetCacheDir()
		fmt.Printf("No custom cheatsheets found in %s\n", cacheDir)
		fmt.Println("Add .md files there to include them in your searches")
		return nil
	}

	// Save to cache
	if err := cache.SaveCache(custom); err != nil {
		return fmt.Errorf("saving cache: %w", err)
	}

	fmt.Printf("Loaded %d custom command(s)\n", len(custom))
	return nil
}

func printUsage() {
	fmt.Println(`Usage: redfish [flags] <search terms...>

Examples:
  redfish git
  redfish git merge
  redfish git merge and delete branch
  redfish docker compose up

Flags:
  -h    Show this help message
  -v    Show version information
  -r    Reload cache from ~/.redfish custom cheatsheets
  -l    Language (default: en, currently only 'en' supported)

Custom Cheatsheets:
  Add your own .md files to ~/.redfish/ and run 'redfish -r' to include them`)
}

func printHelp() {
	fmt.Printf("%s\n\n", version.GetDetailedVersion())
	fmt.Println("A smart command cheatsheet tool with fuzzy search")
	fmt.Println()
	printUsage()
	fmt.Println()
	fmt.Println("Search Algorithm:")
	fmt.Println("  Redfish uses weighted fuzzy matching:")
	fmt.Println("  - Title matches (highest priority)")
	fmt.Println("  - Tag matches")
	fmt.Println("  - Keyword matches")
	fmt.Println("  - Description matches")
	fmt.Println("  - Code matches (lowest priority)")
	fmt.Println()
	fmt.Println("Custom Cheatsheets Format:")
	fmt.Println("  Create .md files in ~/.redfish/ with this format:")
	fmt.Println()
	fmt.Println("  ## command name")
	fmt.Println("  **Tags**: tag1, tag2, tag3")
	fmt.Println("  **Keywords**: keyword1 keyword2 keyword3")
	fmt.Println()
	fmt.Println("  Description of the command")
	fmt.Println()
	fmt.Println("  ```sh")
	fmt.Println("  command example")
	fmt.Println("  ```")
	fmt.Println()
	fmt.Printf("Repository: %s\n", version.GetManifest().Repository)
	fmt.Printf("License: %s\n", version.GetManifest().License)
}
