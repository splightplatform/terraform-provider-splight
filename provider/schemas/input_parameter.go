package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func schemaInputParameter() *schema.Schema {
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
					Type:     schema.TypeString,
					Required: true,
					ValidateFunc: validation.StringInSlice([]string{
						"str",
						"float",
						"int",
						"bool",
						"datetime",
						"url",
						"Asset",
						"Attribute",
						"DataAddress",
						"Component",
						"File",
						"Crontab",
					}, false),
				},
				"value": {
					Type:     schema.TypeString,
					Optional: true,
					Default:  "",
				},
			},
		},
	}
}
