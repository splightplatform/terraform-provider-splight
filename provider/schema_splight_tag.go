package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func schemaTag() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		// NOTE: I included this field to able to reuse the schema in the data source
		// but its not necessary
		"id": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "ID of the resource",
		},
		"name": {
			Type:        schema.TypeString,
			Description: "name of the resource",
			Required:    true,
		},
	}
}
