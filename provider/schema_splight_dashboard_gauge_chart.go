package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func schemaDashboardGaugeChart() map[string]*schema.Schema {
	outputSchema := schemaDashboardChart()
	outputSchema["max_limit"] = &schema.Schema{
		Type:        schema.TypeInt,
		Optional:    true,
		Description: "bar gauge max limit",
	}
	outputSchema["number_of_decimals"] = &schema.Schema{
		Type:        schema.TypeInt,
		Optional:    true,
		Description: "number of decimals",
	}
	return outputSchema
}
