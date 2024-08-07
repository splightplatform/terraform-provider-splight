package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/splightplatform/terraform-provider-splight/utils"
)

func schemaAsset() map[string]*schema.Schema {
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
			DiffSuppressFunc: utils.JSONStringEqualSupressFunc,
		},
		"kind": {
			Type:        schema.TypeSet,
			Optional:    true,
			Description: "kind of the resource",
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"id": {
						Type:        schema.TypeString,
						Required:    true,
						Description: "kind id",
					},
					"name": {
						Type:        schema.TypeString,
						Required:    true,
						Description: "kind name",
					},
				},
			},
		},
		"related_assets": {
			Type:        schema.TypeSet,
			Optional:    true,
			Description: "linked assets",
			Elem:        &schema.Schema{Type: schema.TypeString},
		},
	}
}
