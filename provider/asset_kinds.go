package provider

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func schemaAssetKinds() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"kinds": {
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"id": {
						Type:        schema.TypeString,
						Computed:    true,
						Description: "Id of the resource",
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
