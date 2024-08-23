package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func schemaCommand() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "the name of the command to be created",
		},
		"asset": {
			Type:        schema.TypeSet,
			Optional:    true,
			Description: "associated asset to the command",
			MaxItems:    1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"id": {
						Type:        schema.TypeString,
						Required:    true,
						Description: "asset id",
					},
					"name": {
						Type:        schema.TypeString,
						Required:    true,
						Description: "asset name",
					},
				},
			},
		},
		"actions": {
			Type:        schema.TypeSet,
			Required:    true,
			Description: "command actions",
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"id": {
						Type:        schema.TypeString,
						Computed:    true,
						Description: "action ID",
					},
					"name": {
						Type:        schema.TypeString,
						Computed:    true,
						Description: "setpoint name",
					},
				},
			},
		},
	}
}
