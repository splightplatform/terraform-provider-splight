package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func schemaAssetMetadata() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "name of the resource",
		},
		"type": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "[string|boolean|number] type of the data to be ingested in this attribute",
		},
		"unit": {
			Type:        schema.TypeString,
			Required:    false,
			Optional:    true,
			Description: "optional reference to the unit of the measure",
		},
		"value": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "metadata value",
		},
		"asset": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "reference to the asset to be linked to",
			ForceNew:    true,
		},
	}
}
