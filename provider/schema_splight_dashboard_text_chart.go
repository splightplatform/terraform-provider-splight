package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func schemaDashboardTextChart() map[string]*schema.Schema {
	outputSchema := schemaDashboardChart()
	outputSchema["text"] = &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Description: "text to display",
	}
	return outputSchema
}
