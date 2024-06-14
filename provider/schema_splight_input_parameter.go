package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func InputParameter() *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeSet,
		Optional:    true,
		Description: "static config parameters of the routine",
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"name": {
					Type:     schema.TypeString,
					Required: true,
				},
				"description": {
					Type:     schema.TypeString,
					Required: true,
				},
				"multiple": {
					Type:     schema.TypeBool,
					Required: true,
				},
				"required": {
					Type:     schema.TypeBool,
					Required: true,
				},
				"sensitive": {
					Type:     schema.TypeBool,
					Required: true,
				},
				"type": {
					Type:     schema.TypeString,
					Required: true,
				},
				"value": {
					Type:     schema.TypeString,
					Required: true,
				},
			},
		},
	}
}