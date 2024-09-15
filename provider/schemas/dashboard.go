package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func SchemaDashboard() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "dashboard name",
		},
		"description": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "complementary information for the dashboard",
		},
		"related_assets": {
			Type:        schema.TypeSet,
			Optional:    true,
			Description: "assets linked",
			Elem:        &schema.Schema{Type: schema.TypeString},
		},
	}
}
