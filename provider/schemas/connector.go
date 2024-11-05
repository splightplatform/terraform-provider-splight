package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func SchemaConnector() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "the name of the connector to be created",
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
			Description: "[NAME-VERSION] the version of the hub connector",
		},
		"input": schemaInputParameter(),
		"node": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "id of the compute node where the connector runs",
		},
		"machine_instance_size": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "instance size",
			ValidateFunc: validation.StringInSlice([]string{
				"small",
				"medium",
				"large",
				"very_large",
			}, false),
			Default: "large",
		},
		"log_level": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "log level of the connector",
			ValidateFunc: validation.StringInSlice([]string{
				"critical",
				"error",
				"warning",
				"info",
				"debug",
				"all",
			}, false),
			Default: "info",
		},
		"restart_policy": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "restart policy of the connector",
			ValidateFunc: validation.StringInSlice([]string{
				"Always",
				"Never",
				"OnFailure",
			}, false),
			Default: "OnFailure",
		},
	}
}
