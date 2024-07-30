package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func schemaDashboardTimeseriesChart() map[string]*schema.Schema {
	outputSchema := schemaDashboardChart()
	outputSchema["y_axis_max_limit"] = &schema.Schema{
		Type:        schema.TypeInt,
		Optional:    true,
		Description: "y axis max limit",
	}
	outputSchema["y_axis_min_limit"] = &schema.Schema{
		Type:        schema.TypeInt,
		Optional:    true,
		Description: "y axis min limit",
	}
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
	outputSchema["x_axis_format"] = &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Default:     "MM-dd HH:mm",
		Description: "x axis time format",
	}
	outputSchema["x_axis_auto_skip"] = &schema.Schema{
		Type:        schema.TypeBool,
		Optional:    true,
		Default:     true,
		Description: "x axis auto skip",
	}
	outputSchema["x_axis_max_ticks_limit"] = &schema.Schema{
		Type:        schema.TypeInt,
		Optional:    true,
		Description: "x axis max ticks limit",
	}
	outputSchema["line_interpolation_style"] = &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Description: "line interpolation style",
	}
	outputSchema["timeseries_type"] = &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Default:     "line",
		Description: "[line|bar] timeseries type",
	}
	outputSchema["fill"] = &schema.Schema{
		Type:        schema.TypeBool,
		Optional:    true,
		Default:     false,
		Description: "whether to fill the area under the curve or not",
	}
	outputSchema["show_line"] = &schema.Schema{
		Type:        schema.TypeBool,
		Optional:    true,
		Default:     true,
		Description: "whether to show the line or not",
	}
	return outputSchema
}
