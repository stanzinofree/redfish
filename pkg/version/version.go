package version

import (
	_ "embed"
	"fmt"
	"runtime"
	"strings"

	"gopkg.in/yaml.v3"
)

//go:embed version.yaml
var manifestData []byte

// Manifest holds the version and metadata information
type Manifest struct {
	Version      string   `yaml:"version"`
	Name         string   `yaml:"name"`
	FullName     string   `yaml:"fullName"`
	Description  string   `yaml:"description"`
	Author       string   `yaml:"author"`
	Contact      string   `yaml:"contact"`
	License      string   `yaml:"license"`
	Repository   string   `yaml:"repository"`
	BuildDate    string   `yaml:"buildDate"`
	Spirit       string   `yaml:"spirit"`
	Contributors []string `yaml:"contributors"`
	Releases     struct {
		Latest string `yaml:"latest"`
		Stable string `yaml:"stable"`
	} `yaml:"releases"`
}

var manifest Manifest

func init() {
	if yaml.Unmarshal(manifestData, &manifest) != nil {
		// Fallback to defaults if manifest cannot be read
		manifest.Version = "dev"
		manifest.Name = "redfish"
		manifest.FullName = "Redfish Command Cheatsheet"
		manifest.Description = "A smart command cheatsheet tool with fuzzy search"
		manifest.Author = "Alessandro Middei"
	}
}

// GetVersion returns the current version
func GetVersion() string {
	return manifest.Version
}

// GetVersionInfo returns formatted version information
func GetVersionInfo() string {
	return fmt.Sprintf("%s version %s\nAuthor: %s", manifest.Name, manifest.Version, manifest.Author)
}

// GetDetailedVersion returns detailed version information
func GetDetailedVersion() string {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("%s (%s)\n", manifest.FullName, manifest.Name))
	sb.WriteString(fmt.Sprintf("Version: %s\n", manifest.Version))
	sb.WriteString(fmt.Sprintf("Author: %s\n", manifest.Author))

	if manifest.BuildDate != "" {
		sb.WriteString(fmt.Sprintf("Build Date: %s\n", manifest.BuildDate))
	}

	sb.WriteString(fmt.Sprintf("Go Version: %s\n", runtime.Version()))
	sb.WriteString(fmt.Sprintf("OS/Arch: %s/%s\n", runtime.GOOS, runtime.GOARCH))

	return sb.String()
}

// GetManifest returns the full manifest
func GetManifest() Manifest {
	return manifest
}

// GetAuthor returns the author name
func GetAuthor() string {
	return manifest.Author
}
