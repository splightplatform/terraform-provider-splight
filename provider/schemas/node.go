package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func SchemaNode() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": {
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
			Description: "name of the resource",
		},
		"type": {
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
			Description: "either splight_hosted or self_hosted",
			ValidateFunc: validation.StringInSlice([]string{
				"self_hosted",
				"splight_hosted",
			}, false),
		},
	}
}
