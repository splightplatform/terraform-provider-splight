package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func schemaEnvVars() *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeSet,
		Optional:    true,
		Description: "ports of the server",
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"name": {
					Type:     schema.TypeString,
					Required: true,
				},
				"value": {
					Type:     schema.TypeString,
					Required: true,
				},
			},
		},
	}
}

func schemaPorts() *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeSet,
		Optional:    true,
		Description: "ports of the server",
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"name": {
					Type:     schema.TypeString,
					Required: true,
				},
				"protocol": {
					Type:     schema.TypeString,
					Required: true,
				},
				"internal_port": {
					Type:     schema.TypeInt,
					Required: true,
				},
				"exposed_port": {
					Type:     schema.TypeInt,
					Required: true,
				},
			},
		},
	}
}

func SchemaServer() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "the name of the component to be created",
		},
		"description": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "optional description to add details of the resource",
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
		"version": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "[NAME-VERSION] the version of the hub component",
		},
		"config":   schemaInputParameter(),
		"ports":    schemaPorts(),
		"env_vars": schemaEnvVars(),
	}
}
