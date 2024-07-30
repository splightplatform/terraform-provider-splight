package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func schemaDashboardImageChart() map[string]*schema.Schema {
	outputSchema := schemaDashboardChart()
	outputSchema["image_url"] = &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Description: "image url",
	}
	outputSchema["image_file"] = &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Description: "image file",
	}
	return outputSchema
}
