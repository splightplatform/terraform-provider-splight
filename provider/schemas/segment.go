package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func SchemaSegment() map[string]*schema.Schema {
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
		"temperature": {
			Type:        schema.TypeSet,
			Computed:    true,
			Description: "attribute of the resource",
			Elem: &schema.Resource{
				Schema: schemaConstrainedAttribute(false),
			},
		},
		"wind_speed": {
			Type:        schema.TypeSet,
			Computed:    true,
			Description: "attribute of the resource",
			Elem: &schema.Resource{
				Schema: schemaConstrainedAttribute(false),
			},
		},
		"wind_direction": {
			Type:        schema.TypeSet,
			Computed:    true,
			Description: "attribute of the resource",
			Elem: &schema.Resource{
				Schema: schemaConstrainedAttribute(false),
			},
		},
		"altitude": {
			Type:        schema.TypeSet,
			Optional:    true,
			MaxItems:    1,
			Description: "attribute of the resource",
			Elem: &schema.Resource{
				Schema: schemaConstrainedAttribute(true),
			},
		},
		"azimuth": {
			Type:        schema.TypeSet,
			Optional:    true,
			MaxItems:    1,
			Description: "attribute of the resource",
			Elem: &schema.Resource{
				Schema: schemaConstrainedAttribute(true),
			},
		},
		"cumulative_distance": {
			Type:        schema.TypeSet,
			Optional:    true,
			MaxItems:    1,
			Description: "attribute of the resource",
			Elem: &schema.Resource{
				Schema: schemaConstrainedAttribute(true),
			},
		},
		"reference_sag": {
			Type:        schema.TypeSet,
			Optional:    true,
			MaxItems:    1,
			Description: "attribute of the resource",
			Elem: &schema.Resource{
				Schema: schemaConstrainedAttribute(true),
			},
		},
		"reference_temperature": {
			Type:        schema.TypeSet,
			Optional:    true,
			MaxItems:    1,
			Description: "attribute of the resource",
			Elem: &schema.Resource{
				Schema: schemaConstrainedAttribute(true),
			},
		},
		"span_length": {
			Type:        schema.TypeSet,
			Optional:    true,
			MaxItems:    1,
			Description: "attribute of the resource",
			Elem: &schema.Resource{
				Schema: schemaConstrainedAttribute(true),
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
		"line": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "id of the related Line object",
		},
		"grid": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "id of the related Grid object",
		},
	}
}
