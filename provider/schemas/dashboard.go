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
			Description: "dashboard description",
		},
		"tags": {
			Type:        schema.TypeSet,
			Optional:    true,
			Description: "tags of the resource",
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"id": {
						Type:        schema.TypeString,
						Required:    true,
						Description: "tag id",
					},
					"name": {
						Type:        schema.TypeString,
						Required:    true,
						Description: "tag name",
					},
				},
			},
		},
		"assets": {
			Type:        schema.TypeSet,
			Optional:    true,
			Description: "related assets of the resource",
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"id": {
						Type:        schema.TypeString,
						Required:    true,
						Description: "asset id",
					},
					"name": {
						Type:        schema.TypeString,
						Required:    true,
						Description: "asset name",
					},
				},
			},
		},
	}
}
