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
	"github.com/stanzinofree/redfish/pkg/config"
	"github.com/stanzinofree/redfish/pkg/fzf"
	"github.com/stanzinofree/redfish/pkg/version"
)

var (
	showVersion = flag.Bool("v", false, "Show version information")
	reloadCache = flag.Bool("r", false, "Reload cache from ~/.redfish custom cheatsheets")
	language    = flag.String("l", "", "Language for search (supported: en, it) - overrides config")
	showHelp    = flag.Bool("h", false, "Show help message")
	showConfig  = flag.Bool("c", false, "Run configuration wizard")
	useFzf      = flag.Bool("fzf", false, "Use fzf for interactive selection")
	useShortFzf = flag.Bool("f", false, "Use fzf for interactive selection (short form)")
)

func main() {
	// Parse flags
	flag.Parse()

	// Initialize cache directory on first run
	if err := cache.EnsureCacheDir(); err != nil {
		fmt.Fprintf(os.Stderr, "Warning: could not initialize cache directory: %v\n", err)
	}

	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Warning: could not load config: %v\n", err)
		cfg = config.GetDefaultConfig()
	}

	// Determine language to use (flag overrides config)
	selectedLang := cfg.Language
	if *language != "" {
		selectedLang = *language
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

	if *showConfig {
		if err := config.RunWizard(); err != nil {
			fmt.Fprintf(os.Stderr, "Error running configuration wizard: %v\n", err)
			os.Exit(1)
		}
		os.Exit(0)
	}

	if *reloadCache {
		if err := rebuildCache(selectedLang); err != nil {
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
	commands, err := loadAllCommands(selectedLang)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading commands: %v\n", err)
		os.Exit(1)
	}

	// Search with language-aware stopword filtering
	results := search.SearchWithLanguage(commands, query, selectedLang)

	if len(results) == 0 {
		fmt.Printf("No commands found for: %s\n", query)
		fmt.Println("\nTip: Try different keywords or use -h for help")
		os.Exit(0)
	}

	// Smart output: limit results for specific searches
	results = limitResultsIfSpecific(results, query)

	// Use fzf for interactive selection if requested
	if *useFzf || *useShortFzf {
		if !fzf.IsAvailable() {
			fmt.Fprintln(os.Stderr, "Error: fzf is not installed")
			fmt.Fprintln(os.Stderr, "Install it with: brew install fzf (macOS) or apt install fzf (Linux)")
			os.Exit(1)
		}

		selected, err := fzf.SelectCommand(results)
		if err != nil {
			// User cancelled or error
			os.Exit(0)
		}

		// Render only the selected command
		if err := render.RenderResults([]search.Result{*selected}); err != nil {
			fmt.Fprintf(os.Stderr, "Error rendering result: %v\n", err)
			os.Exit(1)
		}
	} else {
		// Render all results
		if err := render.RenderResults(results); err != nil {
			fmt.Fprintf(os.Stderr, "Error rendering results: %v\n", err)
			os.Exit(1)
		}
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
  -h      Show this help message
  -v      Show version information
  -c      Run configuration wizard (set default language, etc.)
  -r      Reload cache from ~/.redfish custom cheatsheets
  -l      Language (overrides config, supported: en, it)
  -f/--fzf  Use fzf for interactive command selection (requires fzf installed)

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

// limitResultsIfSpecific limits results for multi-token specific searches
func limitResultsIfSpecific(results []search.Result, query string) []search.Result {
	// Count meaningful words in query (after potential stopword removal)
	words := strings.Fields(strings.ToLower(query))
	meaningfulWords := 0
	for _, word := range words {
		if len(word) > 2 { // Count words longer than 2 chars
			meaningfulWords++
		}
	}

	// If it's a specific search (2+ meaningful words), limit to top 5 results
	// This shows only the most relevant snippets
	if meaningfulWords >= 2 && len(results) > 5 {
		return results[:5]
	}

	// For single-word searches, show more results (up to 10)
	if len(results) > 10 {
		return results[:10]
	}

	return results
}
