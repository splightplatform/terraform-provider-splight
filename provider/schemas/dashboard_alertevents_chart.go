package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func SchemaDashboardAlertEventsChart() map[string]*schema.Schema {
	outputSchema := schemaDashboardChart()
	outputSchema["filter_name"] = &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Description: "filter name",
	}
	outputSchema["filter_old_status"] = &schema.Schema{
		Type:        schema.TypeList,
		Optional:    true,
		Elem:        &schema.Schema{Type: schema.TypeString},
		Description: "filter old status",
	}
	outputSchema["filter_new_status"] = &schema.Schema{
		Type:        schema.TypeList,
		Optional:    true,
		Elem:        &schema.Schema{Type: schema.TypeString},
		Description: "filter new status",
	}
	return outputSchema
}
