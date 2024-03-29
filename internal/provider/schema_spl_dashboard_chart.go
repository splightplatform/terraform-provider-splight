package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func schemaDashboardChart() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "name of the chart",
		},
		"tab": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "id for the tab where to place the chart",
		},
		"type": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "[timeseries|bargauge|..] chart type",
		},
		"timestamp_lte": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "date in isoformat or shortcut string where to start reading",
		},
		"timestamp_gte": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "date in isoformat or shortcut string where to end reading",
		},
		"height": {
			Type:        schema.TypeInt,
			Optional:    true,
			Default:     10,
			Description: "chart height in px",
		},
		"width": {
			Type:        schema.TypeInt,
			Optional:    true,
			Default:     20,
			Description: "chart width in cols (max 20)",
		},
		"chart_items": {
			Type:        schema.TypeSet,
			Required:    true,
			Description: "chart traces to be included",
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"color": {
						Type:     schema.TypeString,
						Required: true,
					},
					"ref_id": {
						Type:     schema.TypeString,
						Required: true,
					},
					"type": {
						Type:     schema.TypeString,
						Required: true,
					},
					"expression_plain": {
						Type:     schema.TypeString,
						Required: true,
					},
					"query_plain": {
						Type:     schema.TypeString,
						Required: true,
					},
				},
			},
		},
		"thresholds": {
			Type:        schema.TypeSet,
			Optional:    true,
			Description: "optional static lines to be added to the chart as references",
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"value": {
						Type:     schema.TypeFloat,
						Required: true,
					},
					"color": {
						Type:     schema.TypeString,
						Required: true,
					},
					"display_text": {
						Type:     schema.TypeString,
						Required: true,
					},
				},
			},
		},
		"value_mappings": {
			Type:        schema.TypeSet,
			Optional:    true,
			Description: "optional mappings to transform data with rules",
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"type": {
						Type:     schema.TypeString,
						Required: true,
					},
					"order": {
						Type:     schema.TypeInt,
						Required: true,
					},
					"match_value": {
						Type:     schema.TypeString,
						Required: true,
					},
					"display_text": {
						Type:     schema.TypeString,
						Required: true,
					},
				},
			},
		},
	}
}
