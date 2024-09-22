package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func SchemaDashboardAssetListChart() map[string]*schema.Schema {
	outputSchema := schemaDashboardChart()
	outputSchema["filter_name"] = &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Description: "filter name",
	}
	outputSchema["filter_status"] = &schema.Schema{
		Type:        schema.TypeList,
		Optional:    true,
		Elem:        &schema.Schema{Type: schema.TypeString},
		Description: "filter status list",
	}
	outputSchema["asset_list_type"] = &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Description: "asset list type",
	}
	return outputSchema
}
