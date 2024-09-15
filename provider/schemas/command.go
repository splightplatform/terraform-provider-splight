package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func SchemaCommand() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "the name of the command to be created",
		},
		"description": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "the description of the command to be created",
		},
		"actions": {
			Type:        schema.TypeSet,
			Required:    true,
			Description: "command actions",
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"asset": {
						Type:        schema.TypeSet,
						Required:    true,
						Description: "asset associated with the action (to be deprecated)",
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
					"id": {
						Type:        schema.TypeString,
						Required:    true,
						Description: "action ID",
					},
					"name": {
						Type:        schema.TypeString,
						Required:    true,
						Description: "setpoint name",
					},
				},
			},
		},
	}
}
