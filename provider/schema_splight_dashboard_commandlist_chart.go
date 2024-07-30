package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func schemaDashboardCommandListChart() map[string]*schema.Schema {
	outputSchema := schemaDashboardChart()
	outputSchema["command_list_type"] = &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Default:     "table",
		Description: "[table|button_list]command list type",
	}
	outputSchema["filter_name"] = &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Description: "filter name",
	}
	return outputSchema
}
