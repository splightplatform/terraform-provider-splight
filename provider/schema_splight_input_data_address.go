package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func InputDataAddress() *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeSet,
		Optional:    true,
		Description: "asset attribute where to ingest data. Only valid for IncomingRoutine",
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"name": {
					Type:     schema.TypeString,
					Required: true,
				},
				"description": {
					Type:     schema.TypeString,
					Optional: true,
					Default:  "",
				},
				"multiple": {
					Type:     schema.TypeBool,
					Optional: true,
					Default:  false,
				},
				"required": {
					Type:     schema.TypeBool,
					Optional: true,
					Default:  true,
				},
				"value_type": {
					Type:     schema.TypeString,
					Required: true,
				},
				"type": {
					Type:     schema.TypeString,
					Computed: true,
					Default:  "DataAddress",
				},
				"value": {
					Type:     schema.TypeSet,
					Optional: true,
					Default:  nil,
					MaxItems: 1,
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							"asset": {
								Required: true,
								Type:     schema.TypeString,
							},
							"attribute": {
								Required: true,
								Type:     schema.TypeString,
							},
						},
					},
				},
			},
		},
	}
}
