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
	language    = flag.String("l", "en", "Language for search (supported: en, it)")
	showHelp    = flag.Bool("h", false, "Show help message")
)

func main() {
	// Parse flags
	flag.Parse()

	// Initialize cache directory on first run
	if err := cache.EnsureCacheDir(); err != nil {
		fmt.Fprintf(os.Stderr, "Warning: could not initialize cache directory: %v\n", err)
	}

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
		if err := rebuildCache(*language); err != nil {
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
	commands, err := loadAllCommands(*language)
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

func loadAllCommands(lang string) ([]parser.Command, error) {
	// Load embedded commands for the specified language
	embedded, err := parser.LoadCommandsForLanguage(lang)
	if err != nil {
		return nil, err
	}

	// Try to load custom commands from cache for the specified language
	custom, err := cache.LoadCustomCheatsheetsForLanguage(lang)
	if err != nil {
		// Non-fatal: just use embedded commands
		return embedded, nil
	}

	// Merge embedded and custom
	return append(embedded, custom...), nil
}

func rebuildCache(lang string) error {
	// Clear existing cache
	if err := cache.ClearCache(); err != nil {
		return fmt.Errorf("clearing cache: %w", err)
	}

	// Load custom cheatsheets for the specified language
	custom, err := cache.LoadCustomCheatsheetsForLanguage(lang)
	if err != nil {
		return fmt.Errorf("loading custom cheatsheets: %w", err)
	}

	if len(custom) == 0 {
		cacheDir, _ := cache.GetCacheDir()
		fmt.Printf("No custom cheatsheets found in %s/%s/\n", cacheDir, lang)
		fmt.Printf("Add .md files to %s/%s/ to include them in your searches\n", cacheDir, lang)
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
  redfish -l it git

Flags:
  -h    Show this help message
  -v    Show version information
  -r    Reload cache from ~/.redfish custom cheatsheets
  -l    Language (default: en, supported: en, it)

Custom Cheatsheets:
  Add your own .md files to ~/.redfish/<lang>/ and run 'redfish -r -l <lang>' to include them
  Example: Add Italian cheatsheets to ~/.redfish/it/`)
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
	fmt.Println("  Create .md files in ~/.redfish/<lang>/ with this format:")
	fmt.Println("  (e.g., ~/.redfish/en/ for English, ~/.redfish/it/ for Italian)")
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
	fmt.Printf("Supported Languages: %v\n", cache.GetSupportedLanguages())
	fmt.Println()
	fmt.Printf("Repository: %s\n", version.GetManifest().Repository)
	fmt.Printf("License: %s\n", version.GetManifest().License)
}
