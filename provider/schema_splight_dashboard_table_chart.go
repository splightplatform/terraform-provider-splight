package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func schemaDashboardTableChart() map[string]*schema.Schema {
	outputSchema := schemaDashboardChart()
	outputSchema["y_axis_unit"] = &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Description: "y axis unit",
	}
	outputSchema["number_of_decimals"] = &schema.Schema{
		Type:        schema.TypeInt,
		Optional:    true,
		Description: "number of decimals",
	}
	return outputSchema
}
