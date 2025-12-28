package cache

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/stanzinofree/redfish/internal/parser"
)

const (
	cacheDir          = ".redfish"
	cacheFile         = "commands.cache"
	defaultLanguage   = "en"
	fallbackLanguage  = "en"
)

var supportedLanguages = []string{"en", "it"}

// GetCacheDir returns the cache directory path
func GetCacheDir() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, cacheDir), nil
}

// GetCachePath returns the full cache file path
func GetCachePath() (string, error) {
	dir, err := GetCacheDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(dir, cacheFile), nil
}

// EnsureCacheDir creates the cache directory with language subdirectories if they don't exist
func EnsureCacheDir() error {
	dir, err := GetCacheDir()
	if err != nil {
		return err
	}
	
	// Create main cache directory
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}
	
	// Create language subdirectories
	for _, lang := range supportedLanguages {
		langDir := filepath.Join(dir, lang)
		if err := os.MkdirAll(langDir, 0755); err != nil {
			return err
		}
	}
	
	return nil
}

// SaveCache saves commands to cache file
func SaveCache(commands []parser.Command) error {
	if err := EnsureCacheDir(); err != nil {
		return err
	}

	cachePath, err := GetCachePath()
	if err != nil {
		return err
	}

	data, err := json.MarshalIndent(commands, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(cachePath, data, 0644)
}

// LoadCache loads commands from cache file
func LoadCache() ([]parser.Command, error) {
	cachePath, err := GetCachePath()
	if err != nil {
		return nil, err
	}

	data, err := os.ReadFile(cachePath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil // Cache doesn't exist yet
		}
		return nil, err
	}

	var commands []parser.Command
	if err := json.Unmarshal(data, &commands); err != nil {
		return nil, err
	}

	return commands, nil
}

// ClearCache removes the cache file
func ClearCache() error {
	cachePath, err := GetCachePath()
	if err != nil {
		return err
	}

	if err := os.Remove(cachePath); err != nil && !os.IsNotExist(err) {
		return err
	}

	return nil
}

// LoadCustomCheatsheets loads markdown files from ~/.redfish/<lang>/ directory for a specific language
func LoadCustomCheatsheets() ([]parser.Command, error) {
	return LoadCustomCheatsheetsForLanguage(defaultLanguage)
}

// LoadCustomCheatsheetsForLanguage loads markdown files for a specific language
func LoadCustomCheatsheetsForLanguage(lang string) ([]parser.Command, error) {
	baseDir, err := GetCacheDir()
	if err != nil {
		return nil, err
	}

	dir := filepath.Join(baseDir, lang)
	var commands []parser.Command

	// Read all .md files in the language directory
	entries, err := os.ReadDir(dir)
	if err != nil {
		if os.IsNotExist(err) {
			return commands, nil // Directory doesn't exist yet
		}
		return nil, err
	}

	for _, entry := range entries {
		if entry.IsDir() || filepath.Ext(entry.Name()) != ".md" {
			continue
		}

		filePath := filepath.Join(dir, entry.Name())
		content, err := os.ReadFile(filePath)
		if err != nil {
			continue // Skip files that can't be read
		}

		// Parse the markdown content
		cmds := parser.ParseMarkdownContent(string(content))
		commands = append(commands, cmds...)
	}

	return commands, nil
}

// GetSupportedLanguages returns list of supported language codes
func GetSupportedLanguages() []string {
	return supportedLanguages
}
