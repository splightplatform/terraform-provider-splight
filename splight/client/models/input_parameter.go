package models

import (
	"encoding/json"
	"fmt"
)

type InputParameter struct {
	Name        string           `json:"name"`
	Description string           `json:"description"`
	Multiple    bool             `json:"multiple"`
	Required    bool             `json:"required"`
	Sensitive   bool             `json:"sensitive"`
	Type        string           `json:"type"`
	Value       *json.RawMessage `json:"value"`
}

func (m InputParameter) ToMap() map[string]interface{} {
	var valueStr string
	if m.Value != nil {
		valueStr = string(*m.Value)
	} else {
		valueStr = ""
	}

	return map[string]interface{}{
		"name":        m.Name,
		"description": m.Description,
		"multiple":    m.Multiple,
		"required":    m.Required,
		"sensitive":   m.Sensitive,
		"type":        m.Type,
		"value":       valueStr,
	}
}

func convertInputParameters(data []any) ([]InputParameter, error) {
	inputs := make([]InputParameter, len(data))
	for i, input := range data {
		inputMap := input.(map[string]interface{})
		inputs[i] = InputParameter{
			Name:        inputMap["name"].(string),
			Description: inputMap["description"].(string),
			Multiple:    inputMap["multiple"].(bool),
			Required:    inputMap["required"].(bool),
			Sensitive:   inputMap["sensitive"].(bool),
			Type:        inputMap["type"].(string),
		}
		if value, exists := inputMap["value"]; exists && value != "" {

			// Validate geometry JSON
			valueStr := value.(string)
			if err := validateJSONString(valueStr); err != nil {
				return nil, fmt.Errorf("Value for input parameter %q must be json encoded", inputs[i].Name)
			}

			rawValue := json.RawMessage(valueStr)
			inputs[i].Value = &rawValue
		}
	}
	return inputs, nil
}
