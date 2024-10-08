package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func SchemaAlert() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The name of the resource",
		},
		"description": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The description of the resource",
		},
		"type": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "[cron|rate] type for the cron",
			ValidateFunc: validation.StringInSlice([]string{
				"cron",
				"rate",
			}, false),
		},
		"rate_unit": {
			Type:        schema.TypeString,
			Optional:    true, // Optional for CronAlert
			Computed:    true, // Computed for RateAlert
			Description: "[day|hour|minute] schedule unit",
			ValidateFunc: validation.StringInSlice([]string{
				"day",
				"hour",
				"minute",
			}, false),
		},
		"rate_value": {
			Type:        schema.TypeInt,
			Optional:    true, // Optional for CronAlert
			Computed:    true, // Computed for RateAlert
			Description: "schedule value",
		},
		"cron_minutes": {
			Type:        schema.TypeInt,
			Optional:    true, // Optional for RateAlert
			Computed:    true, // Computed for CronAlert
			Description: "schedule value for cron",
		},
		"cron_hours": {
			Type:        schema.TypeInt,
			Optional:    true, // Optional for RateAlert
			Computed:    true, // Computed for CronAlert
			Description: "schedule value for cron",
		},
		"cron_dom": {
			Type:        schema.TypeInt,
			Optional:    true, // Optional for RateAlert
			Computed:    true, // Computed for CronAlert
			Description: "schedule value for cron",
		},
		"cron_month": {
			Type:        schema.TypeInt,
			Optional:    true, // Optional for RateAlert
			Computed:    true, // Computed for CronAlert
			Description: "schedule value for cron",
		},
		"cron_dow": {
			Type:        schema.TypeInt,
			Optional:    true, // Optional for RateAlert
			Computed:    true, // Computed for CronAlert
			Description: "schedule value for cron",
		},
		"cron_year": {
			Type:        schema.TypeInt,
			Optional:    true, // Optional for RateAlert
			Computed:    true, // Computed for CronAlert
			Description: "schedule value for cron",
		},
		"time_window": {
			Type:        schema.TypeInt,
			Required:    true,
			Description: "window to fetch data from. Data out of that window will not be considered for evaluation ",
		},
		"thresholds": {
			Type:     schema.TypeList,
			Required: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"value": {
						Type:        schema.TypeFloat,
						Required:    true,
						Description: "value to be considered to compare",
					},
					"status": {
						Type:        schema.TypeString,
						Required:    true,
						Description: "[alert|warning|no_alert] status value for the threshold",
						ValidateFunc: validation.StringInSlice([]string{
							"alert",
							"warning",
							"no_alert",
						}, false),
					},
					"status_text": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "optional custom value to be displayed in the platform.",
					},
				},
			},
		},
		"severity": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "[sev1,...,sev8] severity for the alert",
			Elem:        &schema.Schema{Type: schema.TypeString},
			ValidateFunc: validation.StringInSlice([]string{
				"sev1",
				"sev2",
				"sev3",
				"sev4",
				"sev5",
				"sev6",
				"sev7",
				"sev8",
			}, false),
		},
		"operator": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "operator to be used to compare the read value with the threshold value",
			Elem:        &schema.Schema{Type: schema.TypeString},
			ValidateFunc: validation.StringInSlice([]string{
				"gt",
				"ge",
				"lt",
				"le",
				"eq",
			}, false),
		},
		"aggregation": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "aggregation to be applied to reads before comparisson",
			Elem:        &schema.Schema{Type: schema.TypeString},
			ValidateFunc: validation.StringInSlice([]string{
				"max",
				"min",
				"avg",
				"sum",
				"last",
			}, false),
		},
		"target_variable": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "variable to be used to compare with thresholds",
			Elem:        &schema.Schema{Type: schema.TypeString},
		},
		"tags": {
			Type:        schema.TypeSet,
			Optional:    true,
			Description: "tags of the resource",
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"id": {
						Type:        schema.TypeString,
						Required:    true,
						Description: "tag id",
					},
					"name": {
						Type:        schema.TypeString,
						Required:    true,
						Description: "tag name",
					},
				},
			},
		},
		"alert_items": {
			Type:        schema.TypeList,
			Required:    true,
			Description: "traces to be used to compute the results",
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"id": {
						Type:        schema.TypeString,
						Computed:    true,
						Description: "Id of the function item",
					},
					"ref_id": {
						Type:        schema.TypeString,
						Required:    true,
						Description: "identifier of the variable (i.e 'A')",
						ForceNew:    true,
					},
					"type": {
						Type:        schema.TypeString,
						Required:    true,
						Description: "either QUERY or EXPRESSION",
						ValidateFunc: validation.StringInSlice([]string{
							"QUERY",
							"EXPRESSION",
						}, false),
						ForceNew: true,
					},
					"expression": {
						Type:        schema.TypeString,
						Required:    true,
						Description: "how the expression is shown (i.e 'A * 2')",
						ForceNew:    true,
					},
					"expression_plain": {
						Type:        schema.TypeString,
						Required:    true,
						Description: "actual mongo query containing the expression",
						ForceNew:    true,
					},
					"query_filter_asset": {
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
									Description: "Id of the resource",
									ForceNew:    true,
								},
								"name": {
									Type:        schema.TypeString,
									Optional:    true,
									Description: "name of the resource",
									ForceNew:    true,
								},
							},
						},
					},
					"query_filter_attribute": {
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
									Description: "Id of the resource",
									ForceNew:    true,
								},
								"name": {
									Type:        schema.TypeString,
									Optional:    true,
									Description: "name of the resource",
									ForceNew:    true,
								},
							},
						},
					},
					"query_group_function": {
						Type:        schema.TypeString,
						Required:    true,
						Description: "function used to aggregate data",
						ValidateFunc: validation.StringInSlice([]string{
							"",
							"max",
							"min",
							"avg",
							"sum",
							"last",
						}, false),
						ForceNew: true,
					},
					"query_group_unit": {
						Type:        schema.TypeString,
						Required:    true,
						Description: "time window to apply the aggregation",
						ValidateFunc: validation.StringInSlice([]string{
							"",
							"second",
							"minute",
							"hour",
							"day",
							"month",
						}, false),
						ForceNew: true,
					},
					"query_plain": {
						Type:        schema.TypeString,
						Required:    true,
						Description: "actual mongo query",
						ForceNew:    true,
					},
				},
			},
		},
		"related_assets": {
			Type:        schema.TypeSet,
			Optional:    true,
			Description: "related assets of the resource",
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"id": {
						Type:        schema.TypeString,
						Required:    true,
						Description: "asset id",
					},
					"name": {
						Type:        schema.TypeString,
						Required:    true,
						Description: "asset name",
					},
				},
			},
		},
	}
}
