package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func QueryFilter() *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeSet,
		Required:    true,
		Description: "Asset/Attribute filter",
		Default:     nil,
		MaxItems:    1,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"id": {
					Type:        schema.TypeString,
					Required:    true,
					Description: "ID of the resource",
				},
				"name": {
					Type:        schema.TypeString,
					Required:    true,
					Description: "name of the resource",
				},
			},
		},
	}
}
