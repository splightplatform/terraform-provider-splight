package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func schemaFile() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"file": {
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
			Description: "assets to be linked",
			Optional:    true,
			Elem:        &schema.Schema{Type: schema.TypeString},
		},
		"checksum": {
			Type:     schema.TypeString,
			Computed: true,
		},
	}
}
