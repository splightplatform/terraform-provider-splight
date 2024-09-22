package settings

import (
	"fmt"
	"os"
	"sync"

	"gopkg.in/yaml.v3"
)

type Workspace struct {
	AccessId  string `yaml:"SPLIGHT_ACCESS_ID"`
	SecretKey string `yaml:"SPLIGHT_SECRET_KEY"`
	Hostname  string `yaml:"SPLIGHT_PLATFORM_API_HOST"`
}

type ConfigFile struct {
	CurrentWorkspace string               `yaml:"current_workspace"`
	Workspaces       map[string]Workspace `yaml:"workspaces"`
}

type SplightConfig struct {
	Hostname string
	Token    string
}

type SplightConfigOverrides struct {
	HostnameOverride string
	TokenOverride    string
}

var (
	config     *SplightConfig
	configOnce sync.Once
	configErr  error
)

// LoadSplightConfig reads the Splight configuration from the YAML file once and caches it.
// It takes optional configuration options to override specific values.
func LoadSplightConfig(options *SplightConfigOverrides) (*SplightConfig, error) {
	configOnce.Do(func() {
		filename := os.Getenv("HOME") + "/.splight/config"
		buf, err := os.ReadFile(filename)
		if err != nil {
			configErr = err
			return
		}

		c := &ConfigFile{}
		err = yaml.Unmarshal(buf, c)
		if err != nil {
			configErr = fmt.Errorf("in file %q: %w", filename, err)
			return
		}

		// Fetch the current workspace
		workspace, ok := c.Workspaces[c.CurrentWorkspace]
		if !ok {
			configErr = fmt.Errorf("current workspace %q not found", c.CurrentWorkspace)
			return
		}

		// Apply overrides if provided
		if options != nil {
			if options.HostnameOverride != "" {
				workspace.Hostname = options.HostnameOverride
			}
			if options.TokenOverride != "" {
				workspace.SecretKey = options.TokenOverride
			}
		}

		// Build the SplightConfig struct with the token format "Splight <access_id> <secret_key>"
		config = &SplightConfig{
			Hostname: workspace.Hostname,
			Token:    fmt.Sprintf("Splight %s %s", workspace.AccessId, workspace.SecretKey),
		}
	})

	return config, configErr
}
