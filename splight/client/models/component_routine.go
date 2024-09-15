package models

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type ComponentRoutineConfigParam struct {
	Name        string           `json:"name"`
	Description string           `json:"description"`
	Multiple    bool             `json:"multiple"`
	Required    bool             `json:"required"`
	Sensitive   bool             `json:"sensitive"`
	Type        string           `json:"type"`
	Value       *json.RawMessage `json:"value"`
}

type ComponentRoutineDataAddress struct {
	Asset     string `json:"asset"`
	Attribute string `json:"attribute"`
}

type ComponentRoutineDataAddresses []ComponentRoutineDataAddress

func (c *ComponentRoutineDataAddresses) UnmarshalJSON(data []byte) error {
	// Attempt to unmarshal data into a single ComponentRoutineDataAddress
	var single ComponentRoutineDataAddress
	if err := json.Unmarshal(data, &single); err == nil {
		*c = ComponentRoutineDataAddresses{single}
		return nil
	}

	// Attempt to unmarshal data into a slice of ComponentRoutineDataAddress
	var slice []ComponentRoutineDataAddress
	if err := json.Unmarshal(data, &slice); err == nil {
		*c = slice
		return nil
	}

	return fmt.Errorf("failed to unmarshal ComponentRoutineDataAddresses")
}

type ComponentRoutineIOParam struct {
	Name        string                        `json:"name"`
	Description string                        `json:"description"`
	Type        string                        `json:"type"`
	ValueType   string                        `json:"value_type"`
	Multiple    bool                          `json:"multiple"`
	Required    bool                          `json:"required"`
	Value       ComponentRoutineDataAddresses `json:"value"`
}

type ComponentRoutineParams struct {
	Name        string                        `json:"name"`
	Description string                        `json:"description"`
	Type        string                        `json:"type"`
	ComponentId string                        `json:"component_id"`
	Input       []ComponentRoutineIOParam     `json:"input"`
	Output      []ComponentRoutineIOParam     `json:"output"`
	Config      []ComponentRoutineConfigParam `json:"config"`
}

type ComponentRoutine struct {
	ComponentRoutineParams
	ID string `json:"id"`
}

func (m *ComponentRoutine) GetID() string {
	return m.ID
}

func (m *ComponentRoutine) GetParams() Params {
	return &m.ComponentRoutineParams
}

func (m *ComponentRoutine) ResourcePath() string {
	return "v2/engine/component/routines/"
}

func (m *ComponentRoutine) FromSchema(d *schema.ResourceData) error {
	// Handling config
	componentConfigInterface := d.Get("config").(*schema.Set).List()
	componentConfig := make([]ComponentRoutineConfigParam, len(componentConfigInterface))
	for i, item := range componentConfigInterface {
		data := item.(map[string]interface{})
		componentConfig[i] = ComponentRoutineConfigParam{
			Name:        data["name"].(string),
			Description: data["description"].(string),
			Multiple:    data["multiple"].(bool),
			Required:    data["required"].(bool),
			Sensitive:   data["sensitive"].(bool),
			Type:        data["type"].(string),
		}
		if data["value"] != nil && data["value"] != "" {
			value := json.RawMessage(data["value"].(string))
			componentConfig[i].Value = &value
		}
	}

	// Handling output
	componentOutputInterface := d.Get("output").(*schema.Set).List()
	componentOutput := make([]ComponentRoutineIOParam, len(componentOutputInterface))
	for i, item := range componentOutputInterface {
		data := item.(map[string]interface{})
		componentOutput[i] = ComponentRoutineIOParam{
			Name:        data["name"].(string),
			Description: data["description"].(string),
			Type:        data["type"].(string),
			Multiple:    data["multiple"].(bool),
			Required:    data["required"].(bool),
			ValueType:   data["value_type"].(string),
		}
		if data["value"] != nil {
			valueList := data["value"].(*schema.Set).List()
			for _, value := range valueList {
				valueData := value.(map[string]interface{})
				componentOutput[i].Value = append(componentOutput[i].Value, ComponentRoutineDataAddress{
					Asset:     valueData["asset"].(string),
					Attribute: valueData["attribute"].(string),
				})
			}
		}
	}

	// Handling input
	componentInputInterface := d.Get("input").(*schema.Set).List()
	componentInput := make([]ComponentRoutineIOParam, len(componentInputInterface))
	for i, item := range componentInputInterface {
		data := item.(map[string]interface{})
		componentInput[i] = ComponentRoutineIOParam{
			Name:        data["name"].(string),
			Description: data["description"].(string),
			Type:        data["type"].(string),
			Multiple:    data["multiple"].(bool),
			Required:    data["required"].(bool),
			ValueType:   data["value_type"].(string),
		}
		if data["value"] != nil {
			valueList := data["value"].(*schema.Set).List()
			for _, value := range valueList {
				valueData := value.(map[string]interface{})
				componentInput[i].Value = append(componentInput[i].Value, ComponentRoutineDataAddress{
					Asset:     valueData["asset"].(string),
					Attribute: valueData["attribute"].(string),
				})
			}
		}
	}

	m.ComponentRoutineParams = ComponentRoutineParams{
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
		Type:        d.Get("type").(string),
		ComponentId: d.Get("component_id").(string),
		Config:      componentConfig,
		Input:       componentInput,
		Output:      componentOutput,
	}

	return nil
}

func (m *ComponentRoutine) ToSchema(d *schema.ResourceData) error {
	d.SetId(m.ID)
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

	d.Set("input", m.Input)
	d.Set("output", m.Output)

	return nil
}
