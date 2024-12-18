package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func schemaConstrainedAttribute(isMetadata bool, defaultValue interface{}) map[string]*schema.Schema {
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
			Optional:    true,
			Description: "metadata value",
		}

		// Set the default value if provided
		if defaultValue != nil {
			valueSchema.Default = defaultValue
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
			Optional:    true,
			Description: "timezone that overrides location-based timezone of the resource",
		},
		"active_power": {
			Type:        schema.TypeSet,
			Computed:    true,
			Description: "attribute of the resource",
			Elem: &schema.Resource{
				Schema: schemaConstrainedAttribute(false, nil),
			},
		},
		"active_power_end": {
			Type:        schema.TypeSet,
			Computed:    true,
			Description: "attribute of the resource",
			Elem: &schema.Resource{
				Schema: schemaConstrainedAttribute(false, nil),
			},
		},
		"ampacity": {
			Type:        schema.TypeSet,
			Computed:    true,
			Description: "attribute of the resource",
			Elem: &schema.Resource{
				Schema: schemaConstrainedAttribute(false, nil),
			},
		},
		"current": {
			Type:        schema.TypeSet,
			Computed:    true,
			Description: "attribute of the resource",
			Elem: &schema.Resource{
				Schema: schemaConstrainedAttribute(false, nil),
			},
		},
		"current_r": {
			Type:        schema.TypeSet,
			Computed:    true,
			Description: "attribute of the resource",
			Elem: &schema.Resource{
				Schema: schemaConstrainedAttribute(false, nil),
			},
		},
		"current_s": {
			Type:        schema.TypeSet,
			Computed:    true,
			Description: "attribute of the resource",
			Elem: &schema.Resource{
				Schema: schemaConstrainedAttribute(false, nil),
			},
		},
		"current_t": {
			Type:        schema.TypeSet,
			Computed:    true,
			Description: "attribute of the resource",
			Elem: &schema.Resource{
				Schema: schemaConstrainedAttribute(false, nil),
			},
		},
		"energy": {
			Type:        schema.TypeSet,
			Computed:    true,
			Description: "attribute of the resource",
			Elem: &schema.Resource{
				Schema: schemaConstrainedAttribute(false, nil),
			},
		},
		"max_temperature": {
			Type:        schema.TypeSet,
			Computed:    true,
			Description: "attribute of the resource",
			Elem: &schema.Resource{
				Schema: schemaConstrainedAttribute(false, nil),
			},
		},
		"reactive_power": {
			Type:        schema.TypeSet,
			Computed:    true,
			Description: "attribute of the resource",
			Elem: &schema.Resource{
				Schema: schemaConstrainedAttribute(false, nil),
			},
		},
		"voltage_rs": {
			Type:        schema.TypeSet,
			Computed:    true,
			Description: "attribute of the resource",
			Elem: &schema.Resource{
				Schema: schemaConstrainedAttribute(false, nil),
			},
		},
		"voltage_st": {
			Type:        schema.TypeSet,
			Computed:    true,
			Description: "attribute of the resource",
			Elem: &schema.Resource{
				Schema: schemaConstrainedAttribute(false, nil),
			},
		},
		"voltage_tr": {
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
		"switch_status_start": {
			Type:        schema.TypeSet,
			Computed:    true,
			Description: "attribute of the resource",
			Elem: &schema.Resource{
				Schema: schemaConstrainedAttribute(false, nil),
			},
		},
		"switch_status_end": {
			Type:        schema.TypeSet,
			Computed:    true,
			Description: "attribute of the resource",
			Elem: &schema.Resource{
				Schema: schemaConstrainedAttribute(false, nil),
			},
		},
		"diameter": {
			Type:        schema.TypeSet,
			Optional:    true,
			MaxItems:    1,
			Description: "attribute of the resource",
			Elem: &schema.Resource{
				Schema: schemaConstrainedAttribute(true, "30,37"),
			},
		},
		"absorptivity": {
			Type:        schema.TypeSet,
			Optional:    true,
			MaxItems:    1,
			Description: "attribute of the resource",
			Elem: &schema.Resource{
				Schema: schemaConstrainedAttribute(true, "0.6"),
			},
		},
		"atmosphere": {
			Type:        schema.TypeSet,
			Optional:    true,
			MaxItems:    1,
			Description: "attribute of the resource",
			Elem: &schema.Resource{
				Schema: schemaConstrainedAttribute(true, "clean"),
			},
		},
		"capacitance": {
			Type:        schema.TypeSet,
			Optional:    true,
			MaxItems:    1,
			Description: "attribute of the resource",
			Elem: &schema.Resource{
				Schema: schemaConstrainedAttribute(true, "0"),
			},
		},
		"conductance": {
			Type:        schema.TypeSet,
			Optional:    true,
			MaxItems:    1,
			Description: "attribute of the resource",
			Elem: &schema.Resource{
				Schema: schemaConstrainedAttribute(true, "0"),
			},
		},
		"emissivity": {
			Type:        schema.TypeSet,
			Optional:    true,
			MaxItems:    1,
			Description: "attribute of the resource",
			Elem: &schema.Resource{
				Schema: schemaConstrainedAttribute(true, "0.8"),
			},
		},
		"length": {
			Type:        schema.TypeSet,
			Optional:    true,
			MaxItems:    1,
			Description: "attribute of the resource",
			Elem: &schema.Resource{
				Schema: schemaConstrainedAttribute(true, "0"),
			},
		},
		"maximum_allowed_current": {
			Type:        schema.TypeSet,
			Optional:    true,
			MaxItems:    1,
			Description: "attribute of the resource",
			Elem: &schema.Resource{
				Schema: schemaConstrainedAttribute(true, "0"),
			},
		},
		"maximum_allowed_power": {
			Type:        schema.TypeSet,
			Optional:    true,
			MaxItems:    1,
			Description: "attribute of the resource",
			Elem: &schema.Resource{
				Schema: schemaConstrainedAttribute(true, "0"),
			},
		},
		"maximum_allowed_temperature": {
			Type:        schema.TypeSet,
			Optional:    true,
			MaxItems:    1,
			Description: "attribute of the resource",
			Elem: &schema.Resource{
				Schema: schemaConstrainedAttribute(true, "160"),
			},
		},
		"maximum_allowed_temperature_lte": {
			Type:        schema.TypeSet,
			Optional:    true,
			MaxItems:    1,
			Description: "attribute of the resource",
			Elem: &schema.Resource{
				Schema: schemaConstrainedAttribute(true, "160"),
			},
		},
		"maximum_allowed_temperature_ste": {
			Type:        schema.TypeSet,
			Optional:    true,
			MaxItems:    1,
			Description: "attribute of the resource",
			Elem: &schema.Resource{
				Schema: schemaConstrainedAttribute(true, "160"),
			},
		},
		"number_of_conductors": {
			Type:        schema.TypeSet,
			Optional:    true,
			MaxItems:    1,
			Description: "attribute of the resource",
			Elem: &schema.Resource{
				Schema: schemaConstrainedAttribute(true, "3"),
			},
		},
		"reactance": {
			Type:        schema.TypeSet,
			Optional:    true,
			MaxItems:    1,
			Description: "attribute of the resource",
			Elem: &schema.Resource{
				Schema: schemaConstrainedAttribute(true, "0.0293784"),
			},
		},
		"reference_resistance": {
			Type:        schema.TypeSet,
			Optional:    true,
			MaxItems:    1,
			Description: "attribute of the resource",
			Elem: &schema.Resource{
				Schema: schemaConstrainedAttribute(true, "0.0557992717"),
			},
		},
		"resistance": {
			Type:        schema.TypeSet,
			Optional:    true,
			MaxItems:    1,
			Description: "attribute of the resource",
			Elem: &schema.Resource{
				Schema: schemaConstrainedAttribute(true, "0.0293784"),
			},
		},
		"safety_margin_for_power": {
			Type:        schema.TypeSet,
			Optional:    true,
			MaxItems:    1,
			Description: "attribute of the resource",
			Elem: &schema.Resource{
				Schema: schemaConstrainedAttribute(true, "0"),
			},
		},
		"susceptance": {
			Type:        schema.TypeSet,
			Optional:    true,
			MaxItems:    1,
			Description: "attribute of the resource",
			Elem: &schema.Resource{
				Schema: schemaConstrainedAttribute(true, "0.0293784"),
			},
		},
		"temperature_coeff_resistance": {
			Type:        schema.TypeSet,
			Optional:    true,
			MaxItems:    1,
			Description: "attribute of the resource",
			Elem: &schema.Resource{
				Schema: schemaConstrainedAttribute(true, "0.006164923"),
			},
		},
		"specific_heat": {
			Type:        schema.TypeSet,
			Optional:    true,
			MaxItems:    1,
			Description: "attribute of the resource",
			Elem: &schema.Resource{
				Schema: schemaConstrainedAttribute(true, "880"),
			},
		},
		"conductor_mass": {
			Type:        schema.TypeSet,
			Optional:    true,
			MaxItems:    1,
			Description: "attribute of the resource",
			Elem: &schema.Resource{
				Schema: schemaConstrainedAttribute(true, "0.7815"),
			},
		},
		"thermal_elongation_coef": {
			Type:        schema.TypeSet,
			Optional:    true,
			MaxItems:    1,
			Description: "attribute of the resource",
			Elem: &schema.Resource{
				Schema: schemaConstrainedAttribute(true, "0.006164923"),
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
