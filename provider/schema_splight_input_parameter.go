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
					Optional: true,
					Default:  "",
				},
				"multiple": {
					Type:     schema.TypeBool,
					Optional: true,
					Default:  false,
				},
				"required": {
					Type:     schema.TypeBool,
					Optional: true,
					Default:  true,
				},
				"sensitive": {
					Type:     schema.TypeBool,
					Optional: true,
					Default:  false,
				},
				"type": {
					Type:         schema.TypeString,
					Required:     true,
					ExactlyOneOf: []string{"str", "float", "int", "bool"},
				},
				"value": {
					Type:     schema.TypeString,
					Optional: true,
					Default:  nil,
				},
			},
		},
	}
}
