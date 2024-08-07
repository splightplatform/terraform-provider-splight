package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func schemaFunction() map[string]*schema.Schema {
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
		},
		"rate_unit": {
			Type:        schema.TypeString,
			Optional:    true, // Optional for CronAlert
			Computed:    true, // Computed for RateAlert
			Description: "[day|hour|minute] schedule unit",
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
		"target_asset":     QueryFilter(),
		"target_attribute": QueryFilter(),
		"function_items":   Trace(),
	}
}
