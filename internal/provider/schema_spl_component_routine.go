package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func schemaComponentRoutine() map[string]*schema.Schema {
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
		"config": {
			Type:        schema.TypeList,
			Optional:    true,
			Description: "static config parameters of the routine",
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"name": {
						Type:     schema.TypeString,
						Required: true,
					},
					"description": {
						Type:     schema.TypeString,
						Required: true,
					},
					"multiple": {
						Type:     schema.TypeBool,
						Required: true,
					},
					"required": {
						Type:     schema.TypeBool,
						Required: true,
					},
					"sensitive": {
						Type:     schema.TypeBool,
						Required: true,
					},
					"type": {
						Type:     schema.TypeString,
						Required: true,
					},
					"value": {
						Type:     schema.TypeString,
						Required: true,
					},
				},
			},
		},
		"output": {
			Type:        schema.TypeList,
			Optional:    true,
			Description: "asset attribute where to ingest data. Only valid for IncomingRoutine",
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"name": {
						Type:     schema.TypeString,
						Required: true,
					},
					"description": {
						Type:     schema.TypeString,
						Required: true,
					},
					"multiple": {
						Type:     schema.TypeBool,
						Required: true,
					},
					"required": {
						Type:     schema.TypeBool,
						Required: true,
					},
					"value_type": {
						Type:     schema.TypeString,
						Required: true,
					},
					"type": {
						Type:     schema.TypeString,
						Required: true,
					},
					"value": {
						Type:     schema.TypeString,
						Required: true,
					},
				},
			},
		},
		"input": {
			Type:        schema.TypeList,
			Optional:    true,
			Description: "asset attribute where to read data. Only valid for OutgoingRoutine",
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"name": {
						Type:     schema.TypeString,
						Required: true,
					},
					"description": {
						Type:     schema.TypeString,
						Required: true,
					},
					"multiple": {
						Type:     schema.TypeBool,
						Required: true,
					},
					"required": {
						Type:     schema.TypeBool,
						Required: true,
					},
					"value_type": {
						Type:     schema.TypeString,
						Required: true,
					},
					"type": {
						Type:     schema.TypeString,
						Required: true,
					},
					"value": {
						Type:     schema.TypeString,
						Required: true,
					},
				},
			},
		},
	}
}
