package settings

import (
	"fmt"
	"os"
	"sync"

	"gopkg.in/yaml.v3"
)

type Workspace struct {
	AccessID  string `yaml:"SPLIGHT_ACCESS_ID"`
	SecretKey string `yaml:"SPLIGHT_SECRET_KEY"`
	Hostname  string `yaml:"SPLIGHT_PLATFORM_API_HOST"`
}

type Config struct {
	Current    string               `yaml:"current_workspace"`
	Workspaces map[string]Workspace `yaml:"workspaces"`
}

type SplightConfigOverrides struct {
	HostnameOverride string
	TokenOverride    string
}

var (
	config     *Config
	configOnce sync.Once
	configErr  error
)

// LoadSplightConfig reads the Splight configuration from the YAML file once and caches it.
// It takes optional configuration options to override specific values.
func LoadSplightConfig(options *SplightConfigOverrides) (*Config, error) {
	configOnce.Do(func() {
		filename := os.Getenv("HOME") + "/.splight/config"
		buf, err := os.ReadFile(filename)
		if err != nil {
			configErr = err
			return
		}

		c := &Config{}
		err = yaml.Unmarshal(buf, c)
		if err != nil {
			configErr = fmt.Errorf("in file %q: %w", filename, err)
			return
		}

		// Apply overrides if provided
		if workspace, ok := c.Workspaces[c.Current]; ok {
			if options.HostnameOverride != "" {
				workspace.Hostname = options.HostnameOverride
			}
			if options.TokenOverride != "" {
				workspace.SecretKey = options.TokenOverride
			}
			c.Workspaces[c.Current] = workspace
		}

		config = c
	})

	return config, configErr
}
