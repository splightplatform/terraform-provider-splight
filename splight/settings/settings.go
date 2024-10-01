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

// LoadSplightConfig reads the Splight configuration based on the priority:
// overrides -> env vars -> YAML file. The YAML file is only read if not all values are
// provided by overrides or environment variables.
func LoadSplightConfig(options *SplightConfigOverrides) (*SplightConfig, error) {
	configOnce.Do(func() {
		// Load environment variables
		envAccessId := os.Getenv("SPLIGHT_ACCESS_ID")
		envSecretKey := os.Getenv("SPLIGHT_SECRET_KEY")
		envHostname := os.Getenv("SPLIGHT_PLATFORM_API_HOST")

		// Check if all necessary env vars are present
		envVarsPresent := envAccessId != "" && envSecretKey != "" && envHostname != ""

		// Apply overrides if provided
		overrideHostname := ""
		overrideToken := ""
		if options != nil {
			overrideHostname = options.HostnameOverride
			overrideToken = options.TokenOverride
		}

		// If both overrides and environment variables are fully set, use them and skip the file
		if overrideHostname != "" && overrideToken != "" {
			config = &SplightConfig{
				Hostname: overrideHostname,
				Token:    overrideToken,
			}
			return
		}

		if envVarsPresent {
			// If environment variables are fully set, use them and skip the file
			config = &SplightConfig{
				Hostname: envHostname,
				Token:    fmt.Sprintf("Splight %s %s", envAccessId, envSecretKey),
			}
			return
		}

		// Otherwise, read from the YAML file
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

		// Combine the settings from the file, env vars, and overrides in the correct priority
		if overrideHostname != "" {
			workspace.Hostname = overrideHostname
		} else if envHostname != "" {
			workspace.Hostname = envHostname
		}

		if overrideToken != "" {
			workspace.SecretKey = overrideToken
		} else if envAccessId != "" && envSecretKey != "" {
			workspace.AccessId = envAccessId
			workspace.SecretKey = envSecretKey
		}

		// Build the SplightConfig struct with the token format "Splight <access_id> <secret_key>"
		config = &SplightConfig{
			Hostname: workspace.Hostname,
			Token:    fmt.Sprintf("Splight %s %s", workspace.AccessId, workspace.SecretKey),
		}
	})

	return config, configErr
}
