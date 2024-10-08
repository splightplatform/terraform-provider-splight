package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func SchemaDashboardBarChart() map[string]*schema.Schema {
	outputSchema := schemaDashboardChart()
	outputSchema["y_axis_unit"] = &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Description: "y axis units",
	}
	outputSchema["number_of_decimals"] = &schema.Schema{
		Type:        schema.TypeInt,
		Optional:    true,
		Description: "number of decimals",
	}
	outputSchema["orientation"] = &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Description: "orientation",
	}
	return outputSchema
}
