package utils

import (
	"fmt"
	"os"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
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

func readConf(filename string) (*Config, error) {
	buf, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	c := &Config{}
	err = yaml.Unmarshal(buf, c)
	if err != nil {
		return nil, fmt.Errorf("in file %q: %w", filename, err)
	}

	return c, err
}

func TokenDefaultFunc() schema.SchemaDefaultFunc {
	return func() (interface{}, error) {
		access_id := os.Getenv("SPLIGHT_ACCESS_ID")
		secret_key := os.Getenv("SPLIGHT_SECRET_KEY")
		if access_id != "" && secret_key != "" {
			return fmt.Sprintf("Splight %s %s", access_id, secret_key), nil
		}

		configFile := os.Getenv("HOME") + "/.splight/config"
		if _, err := os.Stat(configFile); os.IsNotExist(err) {
			return nil, fmt.Errorf("config file does not exist and no env vars found: %s", configFile)
		}
		config, err := readConf(configFile)
		if err != nil {
			return nil, err
		}
		return fmt.Sprintf("Splight %s %s", config.Workspaces[config.Current].AccessID, config.Workspaces[config.Current].SecretKey), nil

	}
}

func HostnameDefaultFunc() schema.SchemaDefaultFunc {
	return func() (interface{}, error) {
		hostname := os.Getenv("SPLIGHT_PLATFORM_API_HOST")
		if hostname != "" {
			return hostname, nil
		}

		configFile := os.Getenv("HOME") + "/.splight/config"
		if _, err := os.Stat(configFile); os.IsNotExist(err) {
			return nil, fmt.Errorf("config file does not exist and no env vars found: %s", configFile)
		}
		config, err := readConf(configFile)
		if err != nil {
			return nil, err
		}
		return config.Workspaces[config.Current].Hostname, nil

	}
}
