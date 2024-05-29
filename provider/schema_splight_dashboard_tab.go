package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func schemaDashboarTab() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "name for the tab",
		},
		"order": {
			Type:        schema.TypeInt,
			Optional:    true,
			Description: "order within the dashboard",
		},
		"dashboard": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "dashboard id where to place it",
		},
	}
}
