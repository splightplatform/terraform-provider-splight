package models

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type ComponentInputParam struct {
	Name        string           `json:"name"`
	Description string           `json:"description"`
	Multiple    bool             `json:"multiple"`
	Required    bool             `json:"required"`
	Sensitive   bool             `json:"sensitive"`
	Type        string           `json:"type"`
	Value       *json.RawMessage `json:"value"`
}

type ComponentParams struct {
	Name        string                `json:"name"`
	Description string                `json:"description"`
	Tags        []Tag                 `json:"tags"`
	Version     string                `json:"version"`
	Input       []ComponentInputParam `json:"input"`
}

type Component struct {
	ComponentParams
	ID string `json:"id"`
}

func (m *Component) GetID() string {
	return m.ID
}

func (m *Component) GetParams() Params {
	return &m.ComponentParams
}

func (m *Component) ResourcePath() string {
	return "v2/engine/component/components/"
}

func (m *Component) FromSchema(d *schema.ResourceData) error {
	tagsInterface := d.Get("tags").(*schema.Set).List()
	tags := make([]Tag, len(tagsInterface))
	for i, tagInterface := range tagsInterface {
		tagMap := tagInterface.(map[string]interface{})
		tags[i] = Tag{
			ID: tagMap["id"].(string),
			TagParams: TagParams{
				Name: tagMap["name"].(string),
			},
		}
	}

	componentInputInterface := d.Get("input").(*schema.Set).List()
	componentInputInterfaceList := make([]map[string]interface{}, len(componentInputInterface))
	for i, componentInputInterfaceItem := range componentInputInterface {
		componentInputInterfaceList[i] = componentInputInterfaceItem.(map[string]interface{})
	}
	componentInput := make([]ComponentInputParam, len(componentInputInterfaceList))
	for i, componentInputItem := range componentInputInterfaceList {
		componentInput[i] = ComponentInputParam{
			Name:        componentInputItem["name"].(string),
			Description: componentInputItem["description"].(string),
			Multiple:    componentInputItem["multiple"].(bool),
			Required:    componentInputItem["required"].(bool),
			Sensitive:   componentInputItem["sensitive"].(bool),
			Type:        componentInputItem["type"].(string),
		}
		if componentInputItem["value"] != nil && componentInputItem["value"] != "" {
			value := json.RawMessage(componentInputItem["value"].(string))
			componentInput[i].Value = &value
		}
	}

	m.ComponentParams = ComponentParams{
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
		Version:     d.Get("version").(string),
		Input:       componentInput,
		Tags:        tags,
	}

	return nil
}

func (m *Component) ToSchema(d *schema.ResourceData) error {
	d.SetId(m.ID)
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
