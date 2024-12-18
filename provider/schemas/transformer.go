package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func SchemaTransformer() map[string]*schema.Schema {
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
			Description: "timezone that overrides location-based timezone of the resource",
		},
		"active_power_hv": {
			Type:        schema.TypeSet,
			Computed:    true,
			Description: "attribute of the resource",
			Elem: &schema.Resource{
				Schema: schemaConstrainedAttribute(false, nil),
			},
		},
		"active_power_lv": {
			Type:        schema.TypeSet,
			Computed:    true,
			Description: "attribute of the resource",
			Elem: &schema.Resource{
				Schema: schemaConstrainedAttribute(false, nil),
			},
		},
		"reactive_power_hv": {
			Type:        schema.TypeSet,
			Computed:    true,
			Description: "attribute of the resource",
			Elem: &schema.Resource{
				Schema: schemaConstrainedAttribute(false, nil),
			},
		},
		"reactive_power_lv": {
			Type:        schema.TypeSet,
			Computed:    true,
			Description: "attribute of the resource",
			Elem: &schema.Resource{
				Schema: schemaConstrainedAttribute(false, nil),
			},
		},
		"active_power_loss": {
			Type:        schema.TypeSet,
			Computed:    true,
			Description: "attribute of the resource",
			Elem: &schema.Resource{
				Schema: schemaConstrainedAttribute(false, nil),
			},
		},
		"reactive_power_loss": {
			Type:        schema.TypeSet,
			Computed:    true,
			Description: "attribute of the resource",
			Elem: &schema.Resource{
				Schema: schemaConstrainedAttribute(false, nil),
			},
		},
		"current_hv": {
			Type:        schema.TypeSet,
			Computed:    true,
			Description: "attribute of the resource",
			Elem: &schema.Resource{
				Schema: schemaConstrainedAttribute(false, nil),
			},
		},
		"current_lv": {
			Type:        schema.TypeSet,
			Computed:    true,
			Description: "attribute of the resource",
			Elem: &schema.Resource{
				Schema: schemaConstrainedAttribute(false, nil),
			},
		},
		"voltage_hv": {
			Type:        schema.TypeSet,
			Computed:    true,
			Description: "attribute of the resource",
			Elem: &schema.Resource{
				Schema: schemaConstrainedAttribute(false, nil),
			},
		},
		"voltage_lv": {
			Type:        schema.TypeSet,
			Computed:    true,
			Description: "attribute of the resource",
			Elem: &schema.Resource{
				Schema: schemaConstrainedAttribute(false, nil),
			},
		},
		"contingency": {
			Type:        schema.TypeSet,
			Computed:    true,
			Description: "attribute of the resource",
			Elem: &schema.Resource{
				Schema: schemaConstrainedAttribute(false, nil),
			},
		},
		"switch_status_lv": {
			Type:        schema.TypeSet,
			Computed:    true,
			Description: "attribute of the resource",
			Elem: &schema.Resource{
				Schema: schemaConstrainedAttribute(false, nil),
			},
		},
		"switch_status_hv": {
			Type:        schema.TypeSet,
			Computed:    true,
			Description: "attribute of the resource",
			Elem: &schema.Resource{
				Schema: schemaConstrainedAttribute(false, nil),
			},
		},
		"tap_pos": {
			Type:        schema.TypeSet,
			Optional:    true,
			MaxItems:    1,
			Description: "attribute of the resource",
			Elem: &schema.Resource{
				Schema: schemaConstrainedAttribute(true, "0"),
			},
		},
		"xn_ohm": {
			Type:        schema.TypeSet,
			Optional:    true,
			MaxItems:    1,
			Description: "attribute of the resource",
			Elem: &schema.Resource{
				Schema: schemaConstrainedAttribute(true, "0"),
			},
		},
		"standard_type": {
			Type:        schema.TypeSet,
			Optional:    true,
			MaxItems:    1,
			Description: "attribute of the resource",
			Elem: &schema.Resource{
				Schema: schemaConstrainedAttribute(true, ""),
			},
		},
		"capacitance": {
			Type:        schema.TypeSet,
			Optional:    true,
			MaxItems:    1,
			Description: "attribute of the resource",
			Elem: &schema.Resource{
				Schema: schemaConstrainedAttribute(true, "10.7"),
			},
		},
		"conductance": {
			Type:        schema.TypeSet,
			Optional:    true,
			MaxItems:    1,
			Description: "attribute of the resource",
			Elem: &schema.Resource{
				Schema: schemaConstrainedAttribute(true, "0.001"),
			},
		},
		"maximum_allowed_current": {
			Type:        schema.TypeSet,
			Optional:    true,
			MaxItems:    1,
			Description: "attribute of the resource",
			Elem: &schema.Resource{
				Schema: schemaConstrainedAttribute(true, "1.18"),
			},
		},
		"maximum_allowed_power": {
			Type:        schema.TypeSet,
			Optional:    true,
			MaxItems:    1,
			Description: "attribute of the resource",
			Elem: &schema.Resource{
				Schema: schemaConstrainedAttribute(true, "450"),
			},
		},
		"reactance": {
			Type:        schema.TypeSet,
			Optional:    true,
			MaxItems:    1,
			Description: "attribute of the resource",
			Elem: &schema.Resource{
				Schema: schemaConstrainedAttribute(true, "21.80042593"),
			},
		},
		"resistance": {
			Type:        schema.TypeSet,
			Optional:    true,
			MaxItems:    1,
			Description: "attribute of the resource",
			Elem: &schema.Resource{
				Schema: schemaConstrainedAttribute(true, "0.217500888"),
			},
		},
		"safety_margin_for_power": {
			Type:        schema.TypeSet,
			Optional:    true,
			MaxItems:    1,
			Description: "attribute of the resource",
			Elem: &schema.Resource{
				Schema: schemaConstrainedAttribute(true, "5"),
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
