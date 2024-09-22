package models

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type ComponentRoutineParams struct {
	Name        string             `json:"name"`
	Description string             `json:"description"`
	Type        string             `json:"type"`
	ComponentId string             `json:"component_id"`
	Input       []InputDataAddress `json:"input"`
	Output      []InputDataAddress `json:"output"`
	Config      []InputParameter   `json:"config"`
}

type ComponentRoutine struct {
	ComponentRoutineParams
	Id string `json:"id"`
}

func (m *ComponentRoutine) GetId() string {
	return m.Id
}

func (m *ComponentRoutine) GetParams() Params {
	return &m.ComponentRoutineParams
}

func (m *ComponentRoutine) ResourcePath() string {
	return "v2/engine/component/routines/"
}

func (m *ComponentRoutine) FromSchema(d *schema.ResourceData) error {
	m.Id = d.Id()

	config := convertInputParameters(d.Get("config").(*schema.Set).List())
	outputs := convertInputDataAddresses(d.Get("output").(*schema.Set).List())
	inputs := convertInputDataAddresses(d.Get("input").(*schema.Set).List())

	m.ComponentRoutineParams = ComponentRoutineParams{
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
		Type:        d.Get("type").(string),
		ComponentId: d.Get("component_id").(string),
		Config:      config,
		Input:       inputs,
		Output:      outputs,
	}

	return nil
}

func (m *ComponentRoutine) ToSchema(d *schema.ResourceData) error {
	d.SetId(m.Id)

	d.Set("name", m.Name)
	d.Set("description", m.Description)
	d.Set("type", m.Type)
	d.Set("component_id", m.ComponentId)

	// We need to initialize the memory for nested elements
	// Needed because d.Set() can not handle properly json.RawMessage
	if _, ok := d.GetOk("config"); !ok {
		d.Set("config", []interface{}{})
	}

	routineConfigInterface := make([]map[string]interface{}, len(m.Config))
	for i, configItem := range m.Config {
		routineConfigInterface[i] = map[string]interface{}{
			"name":        configItem.Name,
			"description": configItem.Description,
			"multiple":    configItem.Multiple,
			"required":    configItem.Required,
			"sensitive":   configItem.Sensitive,
			"type":        configItem.Type,
			"value":       configItem.Value,
		}
	}
	d.Set("config", routineConfigInterface)

	// TODO: i don't believe this works
	d.Set("input", m.Input)
	d.Set("output", m.Output)

	return nil
}
