package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func schemaDashboardHistogramChart() map[string]*schema.Schema {
	outputSchema := schemaDashboardChart()
	outputSchema["number_of_decimals"] = &schema.Schema{
		Type:        schema.TypeInt,
		Optional:    true,
		Description: "number of decimals",
	}
	outputSchema["bucket_count"] = &schema.Schema{
		Type:        schema.TypeInt,
		Optional:    true,
		Default:     20,
		Description: "bucket count",
	}
	outputSchema["bucket_size"] = &schema.Schema{
		Type:        schema.TypeInt,
		Optional:    true,
		Description: "bucket size",
	}
	outputSchema["histogram_type"] = &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Default:     "numerical",
		Description: "histogram type",
	}
	outputSchema["sorting"] = &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Default:     "count",
		Description: "sorting type",
	}
	outputSchema["stacked"] = &schema.Schema{
		Type:        schema.TypeBool,
		Optional:    true,
		Default:     false,
		Description: "whether to stack or not the histogram",
	}
	outputSchema["categories_top_max_limit"] = &schema.Schema{
		Type:        schema.TypeInt,
		Optional:    true,
		Description: "categories top max limit",
	}
	return outputSchema
}
