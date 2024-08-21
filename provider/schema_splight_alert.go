package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func schemaAlert() map[string]*schema.Schema {
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
		"alert_items": Trace(),
		"related_assets": {
			Type:        schema.TypeSet,
			Optional:    true,
			Description: "related assets to be linked. In case one of these alerts triggers it will be reflected on each of these assets.",
			Elem:        &schema.Schema{Type: schema.TypeString},
		},
	}
}
