package schemas

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func SchemaTag() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": {
			Type:        schema.TypeString,
			Description: "name of the resource",
			Required:    true,
		},
	}
}

func SchemaTags() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"tags": {
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
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
				},
			},
		},
	}
}
