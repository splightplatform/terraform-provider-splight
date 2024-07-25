package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func schemaDashboardAlertListChart() map[string]*schema.Schema {
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
	outputSchema["alert_list_type"] = &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Description: "alert list type",
	}
	return outputSchema
}
