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
		"description": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "chart description",
		},
		"position_x": {
			Type:        schema.TypeInt,
			Optional:    true,
			Description: "chart x position",
		},
		"position_y": {
			Type:        schema.TypeInt,
			Optional:    true,
			Description: "chart y position",
		},
		"min_height": {
			Type:        schema.TypeInt,
			Optional:    true,
			Default:     1,
			Description: "minimum chart height",
		},
		"min_width": {
			Type:        schema.TypeInt,
			Optional:    true,
			Default:     4,
			Description: "minimum chart width",
		},
		"display_time_range": {
			Type:        schema.TypeBool,
			Optional:    true,
			Default:     true,
			Description: "whether to display the time range or not",
		},
		"labels_display": {
			Type:        schema.TypeBool,
			Optional:    true,
			Default:     true,
			Description: "whether to display the labels or not",
		},
		"labels_aggregation": {
			Type:        schema.TypeString,
			Optional:    true,
			Default:     "last",
			Description: "[last|avg|...] aggregation",
		},
		"labels_placement": {
			Type:        schema.TypeString,
			Optional:    true,
			Default:     "bottom",
			Description: "[right|bottom] placement",
		},
		"refresh_interval": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "refresh interval",
		},
		"relative_window_time": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "relative window time",
		},
		"show_beyond_data": {
			Type:        schema.TypeBool,
			Optional:    true,
			Default:     true,
			Description: "whether to show data which is beyond timestamp_lte or not",
		},
		"timezone": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "chart timezone",
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
		"collection": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "default",
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
					"label": {
						Type:     schema.TypeString,
						Optional: true,
						Default:  nil,
					},
					"hidden": {
						Type:     schema.TypeBool,
						Optional: true,
						Default:  nil,
					},
					"query_group_unit": {
						Type:     schema.TypeString,
						Optional: true,
						Default:  "",
					},
					"query_group_function": {
						Type:     schema.TypeString,
						Optional: true,
						Default:  "",
					},
					"expression_plain": {
						Type:     schema.TypeString,
						Required: true,
					},
					"query_filter_asset": &schema.Schema{
						Type:        schema.TypeSet,
						Required:    true,
						Description: "Asset filter",
						Default:     nil,
						MaxItems:    1,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"id": {
									Type:        schema.TypeString,
									Optional:    true,
									Description: "ID of the resource",
								},
								"name": {
									Type:        schema.TypeString,
									Optional:    true,
									Description: "name of the resource",
								},
							},
						},
					},

					"query_filter_attribute": &schema.Schema{
						Type:        schema.TypeSet,
						Required:    true,
						Description: "Attribute filter",
						Default:     nil,
						MaxItems:    1,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"id": {
									Type:        schema.TypeString,
									Optional:    true,
									Description: "ID of the resource",
								},
								"name": {
									Type:        schema.TypeString,
									Optional:    true,
									Description: "name of the resource",
								},
							},
						},
					},

					"query_plain": {
						Type:     schema.TypeString,
						Required: true,
					},
					"query_sort_direction": {
						Type:     schema.TypeInt,
						Optional: true,
						Default:  -1,
					},
					"query_limit": {
						Type:     schema.TypeInt,
						Optional: true,
						Default:  10000,
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
