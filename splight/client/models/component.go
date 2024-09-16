package models

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type ComponentParams struct {
	Name        string           `json:"name"`
	Description string           `json:"description"`
	Tags        []QueryFilter    `json:"tags"`
	Version     string           `json:"version"`
	Input       []InputParameter `json:"input"`
}

type Component struct {
	ComponentParams
	Id string `json:"id"`
}

func (m *Component) GetId() string {
	return m.Id
}

func (m *Component) GetParams() Params {
	return &m.ComponentParams
}

func (m *Component) ResourcePath() string {
	return "v2/engine/component/components/"
}

func (m *Component) FromSchema(d *schema.ResourceData) error {
	tags := convertQueryFilters(d.Get("tags").(*schema.Set).List())
	input := convertInputParameters(d.Get("input").(*schema.Set).List())

	m.ComponentParams = ComponentParams{
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
		Version:     d.Get("version").(string),
		Input:       input,
		Tags:        tags,
	}

	return nil
}

func (m *Component) ToSchema(d *schema.ResourceData) error {
	d.SetId(m.Id)
	d.Set("name", m.Name)
	d.Set("description", m.Description)
	d.Set("version", m.Version)
	d.Set("tags", m.Tags)

	// We need to initialize the memory for nested elements
	// Needed because d.Set() can not handle properly json.RawMessage
	if _, ok := d.GetOk("input"); !ok {
		d.Set("input", []interface{}{})
	}

	inputInterface := make([]map[string]interface{}, len(m.Input))
	for i, inputItem := range m.Input {
		inputInterface[i] = map[string]interface{}{
			"name":        inputItem.Name,
			"description": inputItem.Description,
			"multiple":    inputItem.Multiple,
			"required":    inputItem.Required,
			"sensitive":   inputItem.Sensitive,
			"type":        inputItem.Type,
			"value":       inputItem.Value,
		}
	}
	d.Set("input", inputInterface)

	return nil
}
