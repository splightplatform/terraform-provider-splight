package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func schemaNode() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "name of the resource",
			ForceNew:    true,
		},
		"instance_type": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "intance type",
			ForceNew:    true,
			ValidateFunc: validation.StringInSlice([]string{
				"t2.nano",
				"t2.micro",
				"t2.medium",
				"t2.large",
				"t2.xlarge",
				"t2.2xlarge",
			}, false),
		},
		"organization_id": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "organization id",
			ForceNew:    true,
		},
		"region": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "node region",
			ForceNew:    true,
			ValidateFunc: validation.StringInSlice([]string{
				"us-east-1",
				"us-east-2",
			}, false),
		},
	}
}
