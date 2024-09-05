package provider

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/splightplatform/terraform-provider-splight/api/client"
)

func resourceAlert() *schema.Resource {
	return &schema.Resource{
		Schema:        schemaAlert(),
		CreateContext: resourceCreateAlert,
		ReadContext:   resourceReadAlert,
		UpdateContext: resourceUpdateAlert,
		DeleteContext: resourceDeleteAlert,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func toAlert(d *schema.ResourceData) *client.AlertParams {
	// Convert alert items
	alertItems := convertAlertItems(d.Get("alert_items").([]interface{}))

	// Convert alert thresholds
	alertThresholds := convertAlertThresholds(d.Get("thresholds").([]interface{}))

	// Convert related assets
	alertRelatedAssets := convertRelatedAssets(d.Get("related_assets").(*schema.Set).List())

	// Create the AlertParams object
	item := client.AlertParams{
		Name:           d.Get("name").(string),
		Description:    d.Get("description").(string),
		Type:           d.Get("type").(string),
		TimeWindow:     d.Get("time_window").(int),
		RateUnit:       d.Get("rate_unit").(string),
		RateValue:      d.Get("rate_value").(int),
		Severity:       d.Get("severity").(string),
		Operator:       d.Get("operator").(string),
		Aggregation:    d.Get("aggregation").(string),
		Thresholds:     alertThresholds,
		TargetVariable: d.Get("target_variable").(string),
		AlertItems:     alertItems,
		RelatedAssets:  alertRelatedAssets,
	}

	return &item
}

func convertAlertItems(alertItemsInterface []interface{}) []client.AlertItem {
	alertItems := make([]client.AlertItem, len(alertItemsInterface))
	for i, item := range alertItemsInterface {
		alertItem := item.(map[string]interface{})
		queryFilterAsset := alertItem["query_filter_asset"].(*schema.Set).List()[0].(map[string]interface{})
		queryFilterAttribute := alertItem["query_filter_attribute"].(*schema.Set).List()[0].(map[string]interface{})
		queryGroupFunction := alertItem["query_group_function"].(string)
		queryGroupUnit := alertItem["query_group_unit"].(string)
		alertItems[i] = client.AlertItem{
			RefID:           alertItem["ref_id"].(string),
			Type:            alertItem["type"].(string),
			Expression:      alertItem["expression"].(string),
			ExpressionPlain: alertItem["expression_plain"].(string),
			QueryPlain:      alertItem["query_plain"].(string),
			QueryFilterAsset: client.AlertTargetItem{
				Name: queryFilterAsset["name"].(string),
				ID:   queryFilterAsset["id"].(string),
			},
			QueryFilterAttribute: client.AlertTargetItem{
				Name: queryFilterAttribute["name"].(string),
				ID:   queryFilterAttribute["id"].(string),
			},
			QueryGroupFunction: queryGroupFunction,
			QueryGroupUnit:     queryGroupUnit,
		}
	}
	return alertItems
}

func convertAlertThresholds(thresholdsInterface []interface{}) []client.AlertThreshold {
	alertThresholds := make([]client.AlertThreshold, len(thresholdsInterface))
	for i, item := range thresholdsInterface {
		threshold := item.(map[string]interface{})
		alertThresholds[i] = client.AlertThreshold{
			Value:      threshold["value"].(float64),
			Status:     threshold["status"].(string),
			StatusText: threshold["status_text"].(string),
		}
	}
	return alertThresholds
}

func convertRelatedAssets(relatedAssetsInterface []interface{}) []client.RelatedAsset {
	alertRelatedAssets := make([]client.RelatedAsset, len(relatedAssetsInterface))
	for i, item := range relatedAssetsInterface {
		alertRelatedAssets[i] = client.RelatedAsset{
			Id: item.(string),
		}
	}
	return alertRelatedAssets
}

func saveAlertToState(d *schema.ResourceData, alert *client.Alert) {
	d.SetId(alert.ID)

	d.Set("name", alert.Name)
	d.Set("description", alert.Description)
	d.Set("type", alert.Type)
	d.Set("time_window", alert.TimeWindow)
	d.Set("operator", alert.Operator)
	d.Set("aggregation", alert.Aggregation)
	d.Set("rate_unit", alert.RateUnit)
	d.Set("rate_value", alert.RateValue)
	d.Set("cron_minutes", alert.CronMinutes)
	d.Set("cron_hours", alert.CronHours)
	d.Set("cron_dom", alert.CronDOM)
	d.Set("cron_month", alert.CronMonth)
	d.Set("cron_dow", alert.CronDOW)
	d.Set("cron_year", alert.CronYear)
	d.Set("severity", alert.Severity)
	d.Set("related_assets", alert.RelatedAssets)

	thresholds := make([]map[string]interface{}, len(alert.Thresholds))
	for i, alert := range alert.Thresholds {
		thresholds[i] = map[string]interface{}{
			"value":       alert.Value,
			"status":      alert.Status,
			"status_text": alert.StatusText,
		}
	}
	d.Set("thresholds", thresholds)

	alertItems := make([]map[string]interface{}, len(alert.AlertItems))
	for i, alert := range alert.AlertItems {
		alertItems[i] = map[string]interface{}{
			"id":               alert.ID,
			"ref_id":           alert.RefID,
			"type":             alert.Type,
			"expression":       alert.Expression,
			"expression_plain": alert.ExpressionPlain,
			"query_plain":      alert.QueryPlain,
			"query_filter_asset": []map[string]interface{}{
				{
					"id":   alert.QueryFilterAsset.ID,
					"name": alert.QueryFilterAsset.Name,
				},
			},
			"query_filter_attribute": []map[string]interface{}{
				{
					"id":   alert.QueryFilterAttribute.ID,
					"name": alert.QueryFilterAttribute.Name,
				},
			},
			"query_group_function": alert.QueryGroupFunction,
			"query_group_unit":     alert.QueryGroupUnit,
		}
	}
	d.Set("alert_items", alertItems)
}

func resourceCreateAlert(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	apiClient := m.(*client.Client)

	item := toAlert(d)

	createdAlert, err := apiClient.CreateAlert(item)

	if err != nil {
		return diag.Errorf("error creating Alert: %s", err.Error())
	}

	saveAlertToState(d, createdAlert)

	return nil
}

func resourceUpdateAlert(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	apiClient := m.(*client.Client)

	itemId := d.Id()

	item := toAlert(d)

	updateAlert, err := apiClient.UpdateAlert(itemId, item)

	if err != nil {
		return diag.Errorf("error updating Alert with ID '%s': %s", itemId, err.Error())
	}

	saveAlertToState(d, updateAlert)

	return nil
}

func resourceReadAlert(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	apiClient := m.(*client.Client)

	itemId := d.Id()

	retrievedAlert, err := apiClient.RetrieveAlert(itemId)

	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			d.SetId("")
			return nil
		} else {
			return diag.Errorf("error reading Alert with ID '%s': %s", itemId, err.Error())
		}
	}

	saveAlertToState(d, retrievedAlert)

	return nil
}

func resourceDeleteAlert(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	apiClient := m.(*client.Client)

	itemId := d.Id()

	err := apiClient.DeleteAlert(itemId)

	if err != nil {
		return diag.Errorf("error deleting Alert with ID '%s': %s", itemId, err.Error())
	}

	d.SetId("")

	return nil
}
