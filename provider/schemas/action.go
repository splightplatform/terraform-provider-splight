package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func SchemaAction() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "the name of the action to be created",
		},
		"asset": {
			Type:        schema.TypeSet,
			Required:    true,
			Description: "target asset of the setpoint",
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
		"setpoints": {
			Type:        schema.TypeSet,
			Required:    true,
			Description: "action setpoints",
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"id": {
						Type:        schema.TypeString,
						Computed:    true,
						Description: "setpoint ID",
					},
					"name": {
						Type:        schema.TypeString,
						Computed:    true,
						Description: "setpoint name",
					},
					"value": {
						Type:        schema.TypeString,
						Required:    true,
						Description: "JSON encoded scalar value",
					},
					"attribute": {
						Type:        schema.TypeSet,
						Required:    true,
						Description: "the target attribute of the setpoint which should also be an attribute of the specified asset",
						MaxItems:    1,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"id": {
									Type:        schema.TypeString,
									Required:    true,
									Description: "attribute id",
								},
								"name": {
									Type:        schema.TypeString,
									Required:    true,
									Description: "attribute name",
								},
							},
						},
					},
				},
			},
		},
	}
}
