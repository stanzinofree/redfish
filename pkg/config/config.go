package config

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

const (
	configFile      = "config.yaml"
	defaultLanguage = "en"
)

// Config represents user preferences
type Config struct {
	Language string `yaml:"language"`
}

// GetDefaultConfig returns the default configuration
func GetDefaultConfig() Config {
	return Config{
		Language: defaultLanguage,
	}
}

// GetConfigPath returns the full config file path
func GetConfigPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, ".redfish", configFile), nil
}

// LoadConfig loads the configuration from file
func LoadConfig() (Config, error) {
	configPath, err := GetConfigPath()
	if err != nil {
		return GetDefaultConfig(), err
	}

	data, err := os.ReadFile(configPath)
	if err != nil {
		if os.IsNotExist(err) {
			// Config doesn't exist yet, return default and create it
			defaultCfg := GetDefaultConfig()
			if saveErr := SaveConfig(defaultCfg); saveErr != nil {
				return defaultCfg, fmt.Errorf("creating default config: %w", saveErr)
			}
			return defaultCfg, nil
		}
		return GetDefaultConfig(), err
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return GetDefaultConfig(), err
	}

	return cfg, nil
}

// SaveConfig saves the configuration to file
func SaveConfig(cfg Config) error {
	configPath, err := GetConfigPath()
	if err != nil {
		return err
	}

	// Ensure the directory exists
	if err := os.MkdirAll(filepath.Dir(configPath), 0755); err != nil {
		return err
	}

	data, err := yaml.Marshal(cfg)
	if err != nil {
		return err
	}

	return os.WriteFile(configPath, data, 0644)
}

// ValidateConfig checks if the configuration is valid
func ValidateConfig(cfg Config) error {
	validLanguages := map[string]bool{
		"en": true,
		"it": true,
	}

	if !validLanguages[cfg.Language] {
		return fmt.Errorf("invalid language '%s', supported: en, it", cfg.Language)
	}

	return nil
}
