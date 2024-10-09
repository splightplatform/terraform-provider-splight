package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func SchemaFile() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"path": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "the path for the file resource in your system",
			ForceNew:    true,
		},
		"description": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "complementary information to describe the file",
		},
		"parent": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "the id reference for a folder to be placed in",
		},
		"related_assets": {
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
		"checksum": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"uploaded": {
			Type:     schema.TypeBool,
			Computed: true,
		},
	}
}
