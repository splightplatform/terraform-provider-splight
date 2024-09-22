package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func SchemaDashboardActionListChart() map[string]*schema.Schema {
	outputSchema := schemaDashboardChart()
	outputSchema["action_list_type"] = &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Default:     "table",
		Description: "action list type",
	}
	outputSchema["filter_name"] = &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Description: "filter name",
	}
	outputSchema["filter_asset_name"] = &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Description: "filter asset name",
	}
	return outputSchema
}
