package models

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

type InputDataAddress struct {
	Name        string                        `json:"name"`
	Description string                        `json:"description"`
	Type        string                        `json:"type"`
	ValueType   string                        `json:"value_type"`
	Multiple    bool                          `json:"multiple"`
	Required    bool                          `json:"required"`
	Value       ComponentRoutineDataAddresses `json:"value"`
}
type DataAddress struct {
	Asset     string `json:"asset"`
	Attribute string `json:"attribute"`
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
