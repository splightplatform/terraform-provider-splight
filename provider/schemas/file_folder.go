package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func SchemaFileFolder() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": {
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
			Description: "folder name",
		},
		"parent": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "optional folder id where to place this folder",
		},
	}
}
