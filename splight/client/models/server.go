package models

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Port struct {
	Name         string           `json:"name"`
	Protocol     string           `json:"protocol"`
	InternalPort int              `json:"internal_port"`
	ExposedPort  int              `json:"exposed_port"`
	Value        *json.RawMessage `json:"value"`
}

func (p Port) ToMap() map[string]interface{} {
	var valueStr string
	if p.Value != nil {
		valueStr = string(*p.Value)
	}
	return map[string]interface{}{
		"name":          p.Name,
		"protocol":      p.Protocol,
		"internal_port": p.InternalPort,
		"exposed_port":  p.ExposedPort,
		"value":         valueStr,
	}
}

func convertPorts(data []any) []Port {
	ports := make([]Port, len(data))
	for i, port := range data {
		portMap := port.(map[string]interface{})
		ports[i] = Port{
			Name:         portMap["name"].(string),
			Protocol:     portMap["protocol"].(string),
			InternalPort: portMap["internal_port"].(int),
			ExposedPort:  portMap["exposed_port"].(int),
		}
		if value, exists := portMap["value"]; exists && value != "" {
			rawValue := json.RawMessage(value.(string))
			ports[i].Value = &rawValue
		}
	}
	return ports
}

type EnvVar struct {
	Name  string           `json:"name"`
	Value *json.RawMessage `json:"value"`
}

func (e EnvVar) ToMap() map[string]interface{} {
	var valueStr string
	if e.Value != nil {
		valueStr = string(*e.Value)
	}
	return map[string]interface{}{
		"name":  e.Name,
		"value": valueStr,
	}
}

func convertEnvVars(data []any) []EnvVar {
	envVars := make([]EnvVar, len(data))
	for i, env := range data {
		envMap := env.(map[string]interface{})
		envVars[i] = EnvVar{
			Name: envMap["name"].(string),
		}
		if value, exists := envMap["value"]; exists && value != "" {
			rawValue := json.RawMessage(value.(string))
			envVars[i].Value = &rawValue
		}
	}
	return envVars
}

type ServerParams struct {
	Name        string           `json:"name"`
	Description string           `json:"description"`
	Tags        []QueryFilter    `json:"tags"`
	Version     string           `json:"version"`
	Config      []InputParameter `json:"config"`
	Ports       []Port           `json:"ports"`
	EnvVars     []EnvVar         `json:"env_vars"`
}

type Server struct {
	ServerParams
	Id string `json:"id"`
}

func (m *Server) GetId() string {
	return m.Id
}

func (m *Server) GetParams() Params {
	return &m.ServerParams
}

func (m *Server) ResourcePath() string {
	return "v2/engine/server/servers/"
}

func (m *Server) FromSchema(d *schema.ResourceData) error {
	m.Id = d.Id()
	tags := convertQueryFilters(d.Get("tags").(*schema.Set).List())
	config := convertInputParameters(d.Get("config").(*schema.Set).List())
	ports := convertPorts(d.Get("ports").(*schema.Set).List())
	envVars := convertEnvVars(d.Get("env_vars").(*schema.Set).List())

	m.ServerParams = ServerParams{
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
		Version:     d.Get("version").(string),
		Config:      config,
		Tags:        tags,
		Ports:       ports,
		EnvVars:     envVars,
	}
	return nil
}

func (m *Server) ToSchema(d *schema.ResourceData) error {
	d.SetId(m.Id)
	d.Set("name", m.Name)
	d.Set("description", m.Description)
	d.Set("version", m.Version)

	// Handle Tags
	var tags []map[string]any
	for _, tag := range m.Tags {
		tags = append(tags, map[string]any{
			"id":   tag.Id,
			"name": tag.Name,
		})
	}
	d.Set("tags", tags)

	// Handle Config
	if _, ok := d.GetOk("config"); !ok {
		d.Set("config", []interface{}{})
	}
	configMap := make([]map[string]any, len(m.Config))
	for i, config := range m.Config {
		configMap[i] = config.ToMap()
	}
	d.Set("config", configMap)

	// Handle Ports
	if _, ok := d.GetOk("ports"); !ok {
		d.Set("ports", []interface{}{})
	}
	portsMap := make([]map[string]any, len(m.Ports))
	for i, port := range m.Ports {
		portsMap[i] = port.ToMap()
	}
	d.Set("ports", portsMap)

	// Handle EnvVars
	if _, ok := d.GetOk("env_vars"); !ok {
		d.Set("env_vars", []interface{}{})
	}
	envVarsMap := make([]map[string]any, len(m.EnvVars))
	for i, envVar := range m.EnvVars {
		envVarsMap[i] = envVar.ToMap()
	}
	d.Set("env_vars", envVarsMap)

	return nil
}
