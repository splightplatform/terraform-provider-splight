package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func SchemaComponentRoutine() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "name of the routine",
		},
		"description": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "optional complementary information about the routine",
		},
		"type": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "[IncomingRoutine|OutgoingRoutine] direction of the data flow (from device to system or from system to device)",
		},
		"component_id": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "reference to component to be attached",
		},
		"config": inputParameter(),
		"output": inputDataAddress(),
		"input":  inputDataAddress(),
	}
}
