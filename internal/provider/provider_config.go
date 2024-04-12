package provider

import (
	"fmt"
	"os"

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

func loadYaml(filename string) (*Config, error) {
	c := &Config{}

	buf, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("Error while opening file '%s': %s", filename, err)
	}

	err = yaml.Unmarshal(buf, c)
	if err != nil {
		return nil, fmt.Errorf("Error parsing file '%s': %s", filename, err)
	}

	return c, err
}

func LoadSplightHostname() (string, error) {
	var configFile string = os.Getenv("HOME") + "/.splight/config"
	var hostname string

	hostname = os.Getenv("SPLIGHT_PLATFORM_API_HOST")

	if hostname == "" {
		config, err := loadYaml(configFile)
		if err != nil {
			return "", fmt.Errorf("Missing Splight configuration file at '%s'", configFile)
		}

		hostname = config.Workspaces[config.Current].Hostname
	}

	return hostname, nil

}
func LoadSplightToken() (string, error) {
	var configFile string = os.Getenv("HOME") + "/.splight/config"
	var token string

	accessID := os.Getenv("SPLIGHT_ACCESS_ID")
	secretKey := os.Getenv("SPLIGHT_SECRET_KEY")

	if accessID != "" || secretKey != "" {
		token = fmt.Sprintf("Splight %s %s", accessID, secretKey)
	} else {
		config, err := loadYaml(configFile)
		if err != nil {
			return "", fmt.Errorf("Missing Splight configuration file at '%s'", configFile)
		}

		token = fmt.Sprintf(
			"Splight %s %s",
			config.Workspaces[config.Current].AccessID,
			config.Workspaces[config.Current].SecretKey,
		)
	}

	return token, nil
}
