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
		"tags": {
			Type:        schema.TypeSet,
			Optional:    true,
			Description: "tags for the component",
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"id": {
						Type:        schema.TypeString,
						Required:    true,
						Description: "ID of the resource",
					},
					"name": {
						Type:        schema.TypeString,
						Description: "name of the resource",
						Required:    true,
					},
				},
			},
		},
		"version": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "[NAME-VERSION] the version of the hub component",
		},
		"input": InputParameter(),
	}
}
