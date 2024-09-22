package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func SchemaFunction() map[string]*schema.Schema {
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
		"target_variable": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "variable to be considered to be ingested",
			Elem:        &schema.Schema{Type: schema.TypeString},
		},
		"target_asset": {
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
					},
					"name": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "name of the resource",
					},
				},
			},
		},
		"target_attribute": {
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
					},
					"name": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "name of the resource",
					},
					"type": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "type of the resource",
						ValidateFunc: validation.StringInSlice([]string{
							"String",
							"Boolean",
							"Number",
						}, false),
					},
				},
			},
		},
		"function_items": {
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
					},
					"type": {
						Type:        schema.TypeString,
						Required:    true,
						Description: "either QUERY or EXPRESSION",
						ValidateFunc: validation.StringInSlice([]string{
							"QUERY",
							"EXPRESSION",
						}, false),
					},
					"expression": {
						Type:        schema.TypeString,
						Required:    true,
						Description: "how the expression is shown (i.e 'A * 2')",
					},
					"expression_plain": {
						Type:        schema.TypeString,
						Required:    true,
						Description: "actual mongo query containing the expression",
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
								},
								"name": {
									Type:        schema.TypeString,
									Optional:    true,
									Description: "name of the resource",
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
								},
								"name": {
									Type:        schema.TypeString,
									Optional:    true,
									Description: "name of the resource",
								},
								"type": {
									Type:        schema.TypeString,
									Optional:    true,
									Description: "type of the resource",
									ValidateFunc: validation.StringInSlice([]string{
										"String",
										"Boolean",
										"Number",
									}, false),
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
					},
					"query_plain": {
						Type:        schema.TypeString,
						Required:    true,
						Description: "actual mongo query",
					},
				},
			},
		},
	}
}
