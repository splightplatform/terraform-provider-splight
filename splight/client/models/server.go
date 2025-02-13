package models

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Port struct {
	Name         string `json:"name"`
	Protocol     string `json:"protocol"`
	InternalPort int    `json:"internal_port"`
	ExposedPort  int    `json:"exposed_port"`
}

func (p Port) ToMap() map[string]any {
	return map[string]interface{}{
		"name":          p.Name,
		"protocol":      p.Protocol,
		"internal_port": p.InternalPort,
		"exposed_port":  p.ExposedPort,
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
	}
	return ports
}

type EnvVar struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

func (e EnvVar) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"name":  e.Name,
		"value": e.Value,
	}
}

func convertEnvVars(data []any) []EnvVar {
	if len(data) == 0 {
		return nil
	}
	envVars := make([]EnvVar, len(data))
	for i, envVarData := range data {
		envVarMap := envVarData.(map[string]any)
		envVars[i] = EnvVar{
			Name:  envVarMap["name"].(string),
			Value: envVarMap["value"].(string),
		}
	}
	return envVars
}

type ServerParams struct {
	Name                string           `json:"name"`
	Description         string           `json:"description"`
	Tags                []QueryFilter    `json:"tags"`
	Version             string           `json:"version"`
	Config              []InputParameter `json:"config"`
	Ports               []Port           `json:"ports"`
	EnvVars             []EnvVar         `json:"env_vars"`
	Node                string           `json:"compute_node_id,omitempty"`
	MachineInstanceSize string           `json:"deployment_capacity,omitempty"`
	LogLevel            string           `json:"deployment_log_level,omitempty"`
	RestartPolicy       string           `json:"deployment_restart_policy,omitempty"`
}

type Server struct {
	ServerParams
	Id string `json:"id"`
}

// UnmarshalJSON is a custom method to handle both the top-level "id" and nested "compute_node.id".
func (c *Server) UnmarshalJSON(data []byte) error {
	// Define an auxiliary struct that includes both `id` and `compute_node`
	type Alias Server
	aux := &struct {
		ComputeNode computeNodeWrapper `json:"compute_node"`
		*Alias
	}{
		Alias: (*Alias)(c),
	}

	// Unmarshal JSON into the auxiliary struct
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	// Set the nested compute_node.id value to the Node field in ServerParams
	c.Node = aux.ComputeNode.ID
	return nil
}

func (m *Server) GetId() string {
	return m.Id
}

func (m *Server) GetParams() Params {
	return &m.ServerParams
}

func (m *Server) ResourcePath() string {
	return "v3/engine/server/servers/"
}

func (m *Server) FromSchema(d *schema.ResourceData) error {
	m.Id = d.Id()

	tags := convertQueryFilters(d.Get("tags").(*schema.Set).List())

	config, err := convertInputParameters(d.Get("config").(*schema.Set).List())
	if err != nil {
		return fmt.Errorf("error converting input parameters: %v", err)
	}

	ports := convertPorts(d.Get("ports").(*schema.Set).List())
	envVars := convertEnvVars(d.Get("env_vars").(*schema.Set).List())

	logLevel := d.Get("log_level").(string)

	m.ServerParams = ServerParams{
		Name:                d.Get("name").(string),
		Description:         d.Get("description").(string),
		Version:             d.Get("version").(string),
		Config:              config,
		Tags:                tags,
		Ports:               ports,
		EnvVars:             envVars,
		Node:                d.Get("node").(string),
		MachineInstanceSize: d.Get("machine_instance_size").(string),
		LogLevel:            mapLogLevelToNumber(logLevel),
		RestartPolicy:       d.Get("restart_policy").(string),
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

	configMap := make([]map[string]any, len(m.Config))
	for i, config := range m.Config {
		configMap[i] = config.ToMap()
	}
	d.Set("config", configMap)

	portsMap := make([]map[string]any, len(m.Ports))
	for i, port := range m.Ports {
		portsMap[i] = port.ToMap()
	}
	d.Set("ports", portsMap)

	envVarsMap := make([]map[string]any, len(m.EnvVars))
	for i, envVar := range m.EnvVars {
		envVarsMap[i] = envVar.ToMap()
	}
	d.Set("env_vars", envVarsMap)

	d.Set("node", m.Node)
	d.Set("machine_instance_size", m.MachineInstanceSize)

	// Convert numeric string back to log level name
	d.Set("log_level", mapNumberToLogLevel(m.LogLevel))
	d.Set("restart_policy", m.RestartPolicy)

	return nil
}
