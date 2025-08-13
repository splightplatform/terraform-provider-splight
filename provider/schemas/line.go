package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func schemaConstrainedAttribute(isMetadata bool) map[string]*schema.Schema {
	schemaMap := map[string]*schema.Schema{
		"id": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "id of the resource",
		},
		"name": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "name of the resource",
		},
		"type": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "[String|Boolean|Number] type of the data to be ingested in this attribute",
		},
		"unit": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "unit of measure",
		},
		"asset": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "reference to the asset to be linked to",
		},
	}

	// Add "value" field only if it's metadata
	if isMetadata {
		valueSchema := &schema.Schema{
			Type:        schema.TypeString,
			Required:    true,
			Description: "metadata value",
		}

		schemaMap["value"] = valueSchema
	}

	return schemaMap
}

func SchemaLine() map[string]*schema.Schema {
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
			Computed:    true,
			Description: "timezone of the resource (set by the geo-location)",
		},
		"custom_timezone": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "custom timezone to use instead of the one computed from the geo-location",
		},
		"active_power": {
			Type:        schema.TypeSet,
			Computed:    true,
			Description: "attribute of the resource",
			Elem: &schema.Resource{
				Schema: schemaConstrainedAttribute(false),
			},
		},
		"active_power_end": {
			Type:        schema.TypeSet,
			Computed:    true,
			Description: "attribute of the resource",
			Elem: &schema.Resource{
				Schema: schemaConstrainedAttribute(false),
			},
		},
		"ampacity": {
			Type:        schema.TypeSet,
			Computed:    true,
			Description: "attribute of the resource",
			Elem: &schema.Resource{
				Schema: schemaConstrainedAttribute(false),
			},
		},
		"current": {
			Type:        schema.TypeSet,
			Computed:    true,
			Description: "attribute of the resource",
			Elem: &schema.Resource{
				Schema: schemaConstrainedAttribute(false),
			},
		},
		"current_r": {
			Type:        schema.TypeSet,
			Computed:    true,
			Description: "attribute of the resource",
			Elem: &schema.Resource{
				Schema: schemaConstrainedAttribute(false),
			},
		},
		"current_s": {
			Type:        schema.TypeSet,
			Computed:    true,
			Description: "attribute of the resource",
			Elem: &schema.Resource{
				Schema: schemaConstrainedAttribute(false),
			},
		},
		"current_t": {
			Type:        schema.TypeSet,
			Computed:    true,
			Description: "attribute of the resource",
			Elem: &schema.Resource{
				Schema: schemaConstrainedAttribute(false),
			},
		},
		"energy": {
			Type:        schema.TypeSet,
			Computed:    true,
			Description: "attribute of the resource",
			Elem: &schema.Resource{
				Schema: schemaConstrainedAttribute(false),
			},
		},
		"max_temperature": {
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
		"voltage_rs": {
			Type:        schema.TypeSet,
			Computed:    true,
			Description: "attribute of the resource",
			Elem: &schema.Resource{
				Schema: schemaConstrainedAttribute(false),
			},
		},
		"voltage_st": {
			Type:        schema.TypeSet,
			Computed:    true,
			Description: "attribute of the resource",
			Elem: &schema.Resource{
				Schema: schemaConstrainedAttribute(false),
			},
		},
		"voltage_tr": {
			Type:        schema.TypeSet,
			Computed:    true,
			Description: "attribute of the resource",
			Elem: &schema.Resource{
				Schema: schemaConstrainedAttribute(false),
			},
		},
		"contingency": {
			Type:        schema.TypeSet,
			Computed:    true,
			Description: "attribute of the resource",
			Elem: &schema.Resource{
				Schema: schemaConstrainedAttribute(false),
			},
		},
		"switch_status_start": {
			Type:        schema.TypeSet,
			Computed:    true,
			Description: "attribute of the resource",
			Elem: &schema.Resource{
				Schema: schemaConstrainedAttribute(false),
			},
		},
		"switch_status_end": {
			Type:        schema.TypeSet,
			Computed:    true,
			Description: "attribute of the resource",
			Elem: &schema.Resource{
				Schema: schemaConstrainedAttribute(false),
			},
		},
		"diameter": {
			Type:        schema.TypeSet,
			Optional:    true,
			Computed:    true,
			MaxItems:    1,
			Description: "attribute of the resource",
			Elem: &schema.Resource{
				Schema: schemaConstrainedAttribute(true),
			},
		},
		"absorptivity": {
			Type:        schema.TypeSet,
			Optional:    true,
			Computed:    true,
			MaxItems:    1,
			Description: "attribute of the resource",
			Elem: &schema.Resource{
				Schema: schemaConstrainedAttribute(true),
			},
		},
		"atmosphere": {
			Type:        schema.TypeSet,
			Optional:    true,
			Computed:    true,
			MaxItems:    1,
			Description: "attribute of the resource",
			Elem: &schema.Resource{
				Schema: schemaConstrainedAttribute(true),
			},
		},
		"capacitance": {
			Type:        schema.TypeSet,
			Optional:    true,
			Computed:    true,
			MaxItems:    1,
			Description: "attribute of the resource",
			Elem: &schema.Resource{
				Schema: schemaConstrainedAttribute(true),
			},
		},
		"conductance": {
			Type:        schema.TypeSet,
			Optional:    true,
			Computed:    true,
			MaxItems:    1,
			Description: "attribute of the resource",
			Elem: &schema.Resource{
				Schema: schemaConstrainedAttribute(true),
			},
		},
		"emissivity": {
			Type:        schema.TypeSet,
			Optional:    true,
			Computed:    true,
			MaxItems:    1,
			Description: "attribute of the resource",
			Elem: &schema.Resource{
				Schema: schemaConstrainedAttribute(true),
			},
		},
		"length": {
			Type:        schema.TypeSet,
			Optional:    true,
			Computed:    true,
			MaxItems:    1,
			Description: "attribute of the resource",
			Elem: &schema.Resource{
				Schema: schemaConstrainedAttribute(true),
			},
		},
		"maximum_allowed_current": {
			Type:        schema.TypeSet,
			Optional:    true,
			Computed:    true,
			MaxItems:    1,
			Description: "attribute of the resource",
			Elem: &schema.Resource{
				Schema: schemaConstrainedAttribute(true),
			},
		},
		"maximum_allowed_power": {
			Type:        schema.TypeSet,
			Optional:    true,
			Computed:    true,
			MaxItems:    1,
			Description: "attribute of the resource",
			Elem: &schema.Resource{
				Schema: schemaConstrainedAttribute(true),
			},
		},
		"maximum_allowed_temperature": {
			Type:        schema.TypeSet,
			Optional:    true,
			Computed:    true,
			MaxItems:    1,
			Description: "attribute of the resource",
			Elem: &schema.Resource{
				Schema: schemaConstrainedAttribute(true),
			},
		},
		"maximum_allowed_temperature_lte": {
			Type:        schema.TypeSet,
			Optional:    true,
			Computed:    true,
			MaxItems:    1,
			Description: "attribute of the resource",
			Elem: &schema.Resource{
				Schema: schemaConstrainedAttribute(true),
			},
		},
		"maximum_allowed_temperature_ste": {
			Type:        schema.TypeSet,
			Optional:    true,
			Computed:    true,
			MaxItems:    1,
			Description: "attribute of the resource",
			Elem: &schema.Resource{
				Schema: schemaConstrainedAttribute(true),
			},
		},
		"number_of_conductors": {
			Type:        schema.TypeSet,
			Optional:    true,
			Computed:    true,
			MaxItems:    1,
			Description: "attribute of the resource",
			Elem: &schema.Resource{
				Schema: schemaConstrainedAttribute(true),
			},
		},
		"reactance": {
			Type:        schema.TypeSet,
			Optional:    true,
			Computed:    true,
			MaxItems:    1,
			Description: "attribute of the resource",
			Elem: &schema.Resource{
				Schema: schemaConstrainedAttribute(true),
			},
		},
		"reference_resistance": {
			Type:        schema.TypeSet,
			Optional:    true,
			Computed:    true,
			MaxItems:    1,
			Description: "attribute of the resource",
			Elem: &schema.Resource{
				Schema: schemaConstrainedAttribute(true),
			},
		},
		"resistance": {
			Type:        schema.TypeSet,
			Optional:    true,
			Computed:    true,
			MaxItems:    1,
			Description: "attribute of the resource",
			Elem: &schema.Resource{
				Schema: schemaConstrainedAttribute(true),
			},
		},
		"safety_margin_for_power": {
			Type:        schema.TypeSet,
			Optional:    true,
			Computed:    true,
			MaxItems:    1,
			Description: "attribute of the resource",
			Elem: &schema.Resource{
				Schema: schemaConstrainedAttribute(true),
			},
		},
		"susceptance": {
			Type:        schema.TypeSet,
			Optional:    true,
			Computed:    true,
			MaxItems:    1,
			Description: "attribute of the resource",
			Elem: &schema.Resource{
				Schema: schemaConstrainedAttribute(true),
			},
		},
		"temperature_coeff_resistance": {
			Type:        schema.TypeSet,
			Optional:    true,
			Computed:    true,
			MaxItems:    1,
			Description: "attribute of the resource",
			Elem: &schema.Resource{
				Schema: schemaConstrainedAttribute(true),
			},
		},
		"specific_heat": {
			Type:        schema.TypeSet,
			Optional:    true,
			Computed:    true,
			MaxItems:    1,
			Description: "attribute of the resource",
			Elem: &schema.Resource{
				Schema: schemaConstrainedAttribute(true),
			},
		},
		"conductor_mass": {
			Type:        schema.TypeSet,
			Optional:    true,
			Computed:    true,
			MaxItems:    1,
			Description: "attribute of the resource",
			Elem: &schema.Resource{
				Schema: schemaConstrainedAttribute(true),
			},
		},
		"thermal_elongation_coef": {
			Type:        schema.TypeSet,
			Optional:    true,
			Computed:    true,
			MaxItems:    1,
			Description: "attribute of the resource",
			Elem: &schema.Resource{
				Schema: schemaConstrainedAttribute(true),
			},
		},
		"tags": {
			Type:        schema.TypeSet,
			Optional:    true,
			Computed:    true,
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
