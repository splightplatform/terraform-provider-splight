package models

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type ConnectorParams struct {
	Name        string           `json:"name"`
	Description string           `json:"description"`
	Tags        []QueryFilter    `json:"tags"`
	Version     string           `json:"version"`
	Input       []InputParameter `json:"input"`
}

type Connector struct {
	ConnectorParams
	Id string `json:"id"`
}

func (m *Connector) GetId() string {
	return m.Id
}

func (m *Connector) GetParams() Params {
	return &m.ConnectorParams
}

func (m *Connector) ResourcePath() string {
	return "v2/engine/component/connectors/"
}

func (m *Connector) FromSchema(d *schema.ResourceData) error {
	m.Id = d.Id()

	tags := convertQueryFilters(d.Get("tags").(*schema.Set).List())
	input := convertInputParameters(d.Get("input").(*schema.Set).List())

	m.ConnectorParams = ConnectorParams{
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
		Version:     d.Get("version").(string),
		Input:       input,
		Tags:        tags,
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

	// We need to initialize the memory for nested elements
	// Needed because d.Set() can not handle properly json.RawMessage
	if _, ok := d.GetOk("input"); !ok {
		d.Set("input", []interface{}{})
	}

	inputsMap := make([]map[string]any, len(m.Input))
	for i, input := range m.Input {
		inputsMap[i] = input.ToMap()
	}

	d.Set("input", inputsMap)

	return nil
}
