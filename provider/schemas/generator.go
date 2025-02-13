package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func SchemaGenerator() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "name of the resource",
		},
		"description": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "description of the resource",
		},
		"geometry": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "geo position and shape of the resource",
			DiffSuppressFunc: JSONStringEqualSupressFunc,
		},
		"timezone": {
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "timezone that overrides location-based timezone of the resource",
		},
		"active_power": {
			Type:        schema.TypeSet,
			Computed:    true,
			Description: "attribute of the resource",
			Elem: &schema.Resource{
				Schema: schemaConstrainedAttribute(false),
			},
		},
		"reactive_power": {
			Type:        schema.TypeSet,
			Computed:    true,
			Description: "attribute of the resource",
			Elem: &schema.Resource{
				Schema: schemaConstrainedAttribute(false),
			},
		},
		"daily_energy": {
			Type:        schema.TypeSet,
			Computed:    true,
			Description: "attribute of the resource",
			Elem: &schema.Resource{
				Schema: schemaConstrainedAttribute(false),
			},
		},
		"daily_emission_avoided": {
			Type:        schema.TypeSet,
			Computed:    true,
			Description: "attribute of the resource",
			Elem: &schema.Resource{
				Schema: schemaConstrainedAttribute(false),
			},
		},
		"monthly_energy": {
			Type:        schema.TypeSet,
			Computed:    true,
			Description: "attribute of the resource",
			Elem: &schema.Resource{
				Schema: schemaConstrainedAttribute(false),
			},
		},
		"switch_status": {
			Type:        schema.TypeSet,
			Computed:    true,
			Description: "attribute of the resource",
			Elem: &schema.Resource{
				Schema: schemaConstrainedAttribute(false),
			},
		},
		"tags": {
			Type:        schema.TypeSet,
			Optional:    true,
			Description: "tags of the resource",
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"id": {
						Type:        schema.TypeString,
						Required:    true,
						Description: "tag id",
					},
					"name": {
						Type:        schema.TypeString,
						Required:    true,
						Description: "tag name",
					},
				},
			},
		},
		"kind": {
			Type:        schema.TypeSet,
			Computed:    true,
			Description: "kind of the resource",
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"id": {
						Type:        schema.TypeString,
						Required:    true,
						ForceNew:    true,
						Description: "kind id",
					},
					"name": {
						Type:        schema.TypeString,
						Required:    true,
						ForceNew:    true,
						Description: "kind name",
					},
				},
			},
		},
	}
}
