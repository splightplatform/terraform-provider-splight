package models

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func mapLogLevelToNumber(level string) string {
	switch level {
	case "critical":
		return "50"
	case "error":
		return "40"
	case "warning":
		return "30"
	case "info":
		return "20"
	case "debug":
		return "10"
	case "all":
		return "0"
	default:
		return "0" // Default to "all" if the level is not recognized
	}
}

func mapNumberToLogLevel(number string) string {
	switch number {
	case "50":
		return "critical"
	case "40":
		return "error"
	case "30":
		return "warning"
	case "20":
		return "info"
	case "10":
		return "debug"
	case "0":
		return "all"
	default:
		return "all" // Default to "all" if the number is not recognized
	}
}

type ConnectorParams struct {
	Name                string           `json:"name"`
	Description         string           `json:"description"`
	Tags                []QueryFilter    `json:"tags"`
	Version             string           `json:"version"`
	Input               []InputParameter `json:"input"`
	Node                string           `json:"compute_node_id,omitempty"`
	MachineInstanceSize string           `json:"deployment_capacity,omitempty"`
	LogLevel            string           `json:"deployment_log_level,omitempty"`
	RestartPolicy       string           `json:"deployment_restart_policy,omitempty"`
}

// computeNodeWrapper is used to extract the nested "compute_node" field from JSON
type computeNodeWrapper struct {
	ID string `json:"id"`
}

type Connector struct {
	ConnectorParams
	Id string `json:"id"`
}

// UnmarshalJSON is a custom method to handle both the top-level "id" and nested "compute_node.id".
func (c *Connector) UnmarshalJSON(data []byte) error {
	// Define an auxiliary struct that includes both `id` and `compute_node`
	type Alias Connector
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

	// Set the nested compute_node.id value to the Node field in ConnectorParams
	c.Node = aux.ComputeNode.ID
	return nil
}

func (m *Connector) GetId() string {
	return m.Id
}

func (m *Connector) GetParams() Params {
	return &m.ConnectorParams
}

func (m *Connector) ResourcePath() string {
	return "v3/engine/component/connectors/"
}

func (m *Connector) FromSchema(d *schema.ResourceData) error {
	m.Id = d.Id()

	tags := convertQueryFilters(d.Get("tags").(*schema.Set).List())
	input, err := convertInputParameters(d.Get("input").(*schema.Set).List())
	if err != nil {
		return fmt.Errorf("error converting input parameters: %v", err)
	}

	// Convert the log level to a numeric string
	logLevel := d.Get("log_level").(string)
	m.ConnectorParams = ConnectorParams{
		Name:                d.Get("name").(string),
		Description:         d.Get("description").(string),
		Version:             d.Get("version").(string),
		Input:               input,
		Tags:                tags,
		Node:                d.Get("node").(string),
		MachineInstanceSize: d.Get("machine_instance_size").(string),
		LogLevel:            mapLogLevelToNumber(logLevel),
		RestartPolicy:       d.Get("restart_policy").(string),
	}

	return nil
}

func (m *Connector) ToSchema(d *schema.ResourceData) error {
	d.SetId(m.Id)

	d.Set("name", m.Name)
	d.Set("description", m.Description)
	d.Set("version", m.Version)

	var tags []map[string]any
	for _, tag := range m.Tags {
		tags = append(tags, map[string]any{
			"id":   tag.Id,
			"name": tag.Name,
		})
	}
	d.Set("tags", tags)

	inputMap := make([]map[string]any, len(m.Input))
	for i, input := range m.Input {
		inputMap[i] = input.ToMap()
	}
	d.Set("input", inputMap)

	d.Set("node", m.Node)
	d.Set("machine_instance_size", m.MachineInstanceSize)

	// Convert numeric string back to log level name
	d.Set("log_level", mapNumberToLogLevel(m.LogLevel))
	d.Set("restart_policy", m.RestartPolicy)

	return nil
}
