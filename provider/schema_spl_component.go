package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func schemaComponent() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "the name of the component to be created",
		},
		"description": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "optinal description to add details of the resource",
		},
		"version": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "[NAME-VERSION] the version of the hub component",
		},
		"input": {
			Type:        schema.TypeList,
			Required:    true,
			Description: "The input parameters based on hubcomponent spec",
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
		},
	}
}
