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
		configFile := os.Getenv("HOME") + "/.splight/config"
		config, err := readConf(configFile)
		if err != nil {
			return nil, err
		}
		return fmt.Sprintf("Splight %s %s", config.Workspaces[config.Current].AccessID, config.Workspaces[config.Current].SecretKey), nil

	}
}

func HostnameDefaultFunc() schema.SchemaDefaultFunc {
	return func() (interface{}, error) {
		configFile := os.Getenv("HOME") + "/.splight/config"
		config, err := readConf(configFile)
		if err != nil {
			return nil, err
		}
		return config.Workspaces[config.Current].Hostname, nil

	}
}
