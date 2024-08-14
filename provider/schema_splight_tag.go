package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func schemaTag() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "ID of the resource",
		},
		"name": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "name of the resource",
		},
	}
}
