package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func SchemaAssetMetadata() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "name of the resource",
		},
		"type": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "[String|Boolean|Number] type of the data to be ingested in this attribute",
			ValidateFunc: validation.StringInSlice([]string{
				"String",
				"Boolean",
				"Number",
			}, false),
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
