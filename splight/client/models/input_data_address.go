package models

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type DataAddress struct {
	Asset     string `json:"asset"`
	Attribute string `json:"attribute"`
}

type InputDataAddress struct {
	Name        string                        `json:"name"`
	Description string                        `json:"description"`
	Type        string                        `json:"type"`
	ValueType   string                        `json:"value_type"`
	Multiple    bool                          `json:"multiple"`
	Required    bool                          `json:"required"`
	Value       ComponentRoutineDataAddresses `json:"value"`
}

type ComponentRoutineDataAddresses []DataAddress

// We do this since API might return a list or a single data address
func (c *ComponentRoutineDataAddresses) UnmarshalJSON(data []byte) error {
	// Attempt to unmarshal data into a single ComponentRoutineDataAddress
	var single DataAddress
	if err := json.Unmarshal(data, &single); err == nil {
		*c = ComponentRoutineDataAddresses{single}
		return nil
	}

	// Attempt to unmarshal data into a slice of ComponentRoutineDataAddress
	var slice []DataAddress
	if err := json.Unmarshal(data, &slice); err == nil {
		*c = slice
		return nil
	}

	return fmt.Errorf("failed to unmarshal ComponentRoutineDataAddresses")
}

// Method to convert InputDataAddress to a map format
func (m InputDataAddress) ToMap() map[string]interface{} {
	valueList := make([]map[string]interface{}, len(m.Value))
	for i, dataAddr := range m.Value {
		valueList[i] = map[string]interface{}{
			"asset":     dataAddr.Asset,
			"attribute": dataAddr.Attribute,
		}
	}

	return map[string]interface{}{
		"name":        m.Name,
		"description": m.Description,
		"type":        m.Type,
		"value_type":  m.ValueType,
		"multiple":    m.Multiple,
		"required":    m.Required,
		"value":       valueList,
	}
}

func convertInputDataAddresses(data []any) []InputDataAddress {
	inputs := make([]InputDataAddress, len(data))
	for i, item := range data {
		data := item.(map[string]interface{})
		inputs[i] = InputDataAddress{
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
				inputs[i].Value = append(inputs[i].Value, DataAddress{
					Asset:     valueData["asset"].(string),
					Attribute: valueData["attribute"].(string),
				})
			}
		}
	}
	return inputs
}
