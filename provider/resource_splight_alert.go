package provider

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/splightplatform/terraform-provider-splight/api/client"
)

func resourceAlert() *schema.Resource {
	return &schema.Resource{
		Schema: schemaAlert(),
		Create: resourceCreateAlert,
		Read:   resourceReadAlert,
		Update: resourceUpdateAlert,
		Delete: resourceDeleteAlert,
		Exists: resourceExistsAlert,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceCreateAlert(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	alertItemInterface := d.Get("alert_items").([]interface{})
	alertItemInterfaceList := make([]map[string]interface{}, len(alertItemInterface))
	for i, alertItemInterfaceItem := range alertItemInterface {
		alertItemInterfaceList[i] = alertItemInterfaceItem.(map[string]interface{})
	}
	alertItems := make([]client.AlertItem, len(alertItemInterfaceList))
	for i, alertItem := range alertItemInterfaceList {
		queryFilterAsset := alertItem["query_filter_asset"].(*schema.Set).List()[0].(map[string]interface{})
		queryFilterAttribute := alertItem["query_filter_attribute"].(*schema.Set).List()[0].(map[string]interface{})
		alertItems[i] = client.AlertItem{
			RefID:           alertItem["ref_id"].(string),
			Type:            alertItem["type"].(string),
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
		}
	}

	alertThresholdInterface := d.Get("thresholds").([]interface{})
	alertThresholdInterfaceList := make([]map[string]interface{}, len(alertThresholdInterface))
	for i, alertThresholdInterfaceItem := range alertThresholdInterface {
		alertThresholdInterfaceList[i] = alertThresholdInterfaceItem.(map[string]interface{})
	}
	alertThresholds := make([]client.AlertThreshold, len(alertThresholdInterfaceList))
	for i, alertThresholdItem := range alertThresholdInterfaceList {
		alertThresholds[i] = client.AlertThreshold{
			Value:      alertThresholdItem["value"].(float64),
			Status:     alertThresholdItem["status"].(string),
			StatusText: alertThresholdItem["status_text"].(string),
		}
	}

	alertRelatedAssetsSet := d.Get("related_assets").(*schema.Set).List()
	alertRelatedAssets := make([]client.RelatedAsset, len(alertRelatedAssetsSet))
	for i, relatedAsset := range alertRelatedAssetsSet {
		alertRelatedAssets[i] = client.RelatedAsset{
			Id: relatedAsset.(string),
		}
	}

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

	createdAlert, err := apiClient.CreateAlert(&item)
	if err != nil {
		return err
	}

	d.SetId(createdAlert.ID)
	d.Set("name", createdAlert.Name)
	d.Set("description", createdAlert.Description)
	d.Set("type", createdAlert.Type)
	d.Set("time_window", createdAlert.TimeWindow)
	d.Set("operator", createdAlert.Operator)
	d.Set("aggregation", createdAlert.Aggregation)
	d.Set("rate_unit", createdAlert.RateUnit)
	d.Set("rate_value", createdAlert.RateValue)
	d.Set("cron_minutes", createdAlert.CronMinutes)
	d.Set("cron_hours", createdAlert.CronHours)
	d.Set("cron_dom", createdAlert.CronDOM)
	d.Set("cron_month", createdAlert.CronMonth)
	d.Set("cron_dow", createdAlert.CronDOW)
	d.Set("cron_year", createdAlert.CronYear)
	d.Set("severity", createdAlert.Severity)
	d.Set("related_assets", createdAlert.RelatedAssets)
	alertThresholdsOutput := make([]map[string]interface{}, len(createdAlert.Thresholds))
	for i, alertThresholdItem := range createdAlert.Thresholds {
		alertThresholdsOutput[i] = map[string]interface{}{
			"value":       alertThresholdItem.Value,
			"status":      alertThresholdItem.Status,
			"status_text": alertThresholdItem.StatusText,
		}
	}
	d.Set("thresholds", alertThresholdsOutput)
	alertItemsOutput := make([]map[string]interface{}, len(createdAlert.AlertItems))
	for i, alertItem := range createdAlert.AlertItems {
		alertItemsOutput[i] = map[string]interface{}{
			"id":               alertItem.ID,
			"ref_id":           alertItem.RefID,
			"type":             alertItem.Type,
			"expression_plain": alertItem.ExpressionPlain,
			"query_plain":      alertItem.QueryPlain,
			"query_filter_asset": map[string]interface{}{
				"name": alertItem.QueryFilterAsset.Name,
				"id":   alertItem.QueryFilterAsset.ID,
			},
			"query_filter_attribute": map[string]interface{}{
				"name": alertItem.QueryFilterAttribute.Name,
				"id":   alertItem.QueryFilterAttribute.ID,
			},
		}
	}
	d.Set("alert_items", alertItemsOutput)

	return nil
}

func resourceUpdateAlert(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	itemId := d.Id()

	alertItemInterface := d.Get("alert_items").([]interface{})
	alertItemInterfaceList := make([]map[string]interface{}, len(alertItemInterface))
	for i, alertItemInterfaceItem := range alertItemInterface {
		alertItemInterfaceList[i] = alertItemInterfaceItem.(map[string]interface{})
	}
	alertItems := make([]client.AlertItem, len(alertItemInterfaceList))
	for i, alertItem := range alertItemInterfaceList {
		queryFilterAsset := alertItem["query_filter_asset"].(*schema.Set).List()[0].(map[string]interface{})
		queryFilterAttribute := alertItem["query_filter_attribute"].(*schema.Set).List()[0].(map[string]interface{})
		alertItems[i] = client.AlertItem{
			ID:              alertItem["id"].(string),
			RefID:           alertItem["ref_id"].(string),
			Type:            alertItem["type"].(string),
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
		}
	}

	alertThresholdInterface := d.Get("thresholds").([]interface{})
	alertThresholdInterfaceList := make([]map[string]interface{}, len(alertThresholdInterface))
	for i, alertThresholdInterfaceItem := range alertThresholdInterface {
		alertThresholdInterfaceList[i] = alertThresholdInterfaceItem.(map[string]interface{})
	}
	alertThresholds := make([]client.AlertThreshold, len(alertThresholdInterfaceList))
	for i, alertThresholdItem := range alertThresholdInterfaceList {
		alertThresholds[i] = client.AlertThreshold{
			Value:      alertThresholdItem["value"].(float64),
			Status:     alertThresholdItem["status"].(string),
			StatusText: alertThresholdItem["status_text"].(string),
		}
	}

	alertRelatedAssetsSet := d.Get("related_assets").(*schema.Set).List()
	alertRelatedAssets := make([]client.RelatedAsset, len(alertRelatedAssetsSet))
	for i, relatedAsset := range alertRelatedAssetsSet {
		alertRelatedAssets[i] = client.RelatedAsset{
			Id: relatedAsset.(string),
		}
	}

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

	updateAlert, err := apiClient.UpdateAlert(itemId, &item)
	if err != nil {
		return err
	}

	d.SetId(updateAlert.ID)
	d.Set("name", updateAlert.Name)
	d.Set("description", updateAlert.Description)
	d.Set("type", updateAlert.Type)
	d.Set("time_window", updateAlert.TimeWindow)
	d.Set("operator", updateAlert.Operator)
	d.Set("aggregation", updateAlert.Aggregation)
	d.Set("rate_unit", updateAlert.RateUnit)
	d.Set("rate_value", updateAlert.RateValue)
	d.Set("cron_minutes", updateAlert.CronMinutes)
	d.Set("cron_hours", updateAlert.CronHours)
	d.Set("cron_dom", updateAlert.CronDOM)
	d.Set("cron_month", updateAlert.CronMonth)
	d.Set("cron_dow", updateAlert.CronDOW)
	d.Set("cron_year", updateAlert.CronYear)
	d.Set("severity", updateAlert.Severity)
	d.Set("related_assets", updateAlert.RelatedAssets)
	alertThresholdsOutput := make([]map[string]interface{}, len(updateAlert.Thresholds))
	for i, alertThresholdItem := range updateAlert.Thresholds {
		alertThresholdsOutput[i] = map[string]interface{}{
			"value":       alertThresholdItem.Value,
			"status":      alertThresholdItem.Status,
			"status_text": alertThresholdItem.StatusText,
		}
	}
	d.Set("thresholds", alertThresholdsOutput)
	alertItemsOutput := make([]map[string]interface{}, len(updateAlert.AlertItems))
	for i, alertItem := range updateAlert.AlertItems {
		alertItemsOutput[i] = map[string]interface{}{
			"id":               alertItem.ID,
			"ref_id":           alertItem.RefID,
			"type":             alertItem.Type,
			"expression_plain": alertItem.ExpressionPlain,
			"query_plain":      alertItem.QueryPlain,
			"query_filter_asset": map[string]interface{}{
				"name": alertItem.QueryFilterAsset.Name,
				"id":   alertItem.QueryFilterAsset.ID,
			},
			"query_filter_attribute": map[string]interface{}{
				"name": alertItem.QueryFilterAttribute.Name,
				"id":   alertItem.QueryFilterAttribute.ID,
			},
		}
	}
	d.Set("alert_items", alertItemsOutput)

	return nil
}

func resourceReadAlert(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	itemId := d.Id()
	retrievedAlert, err := apiClient.RetrieveAlert(itemId)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			d.SetId("")
		} else {
			return fmt.Errorf("error finding Alert with ID %s", itemId)
		}
	}

	alertItemsDict := make([]map[interface{}]interface{}, len(retrievedAlert.AlertItems))
	for i, alertItem := range retrievedAlert.AlertItems {
		alertItemsDict[i] = map[interface{}]interface{}{
			"id":               alertItem.ID,
			"ref_id":           alertItem.RefID,
			"type":             alertItem.Type,
			"expression_plain": alertItem.ExpressionPlain,
			"query_plain":      alertItem.QueryPlain,
			"query_filter_asset": map[string]interface{}{
				"name": alertItem.QueryFilterAsset.Name,
				"id":   alertItem.QueryFilterAsset.ID,
			},
			"query_filter_attribute": map[string]interface{}{
				"name": alertItem.QueryFilterAttribute.Name,
				"id":   alertItem.QueryFilterAttribute.ID,
			},
		}
	}

	thresholdsDict := make([]map[interface{}]interface{}, len(retrievedAlert.Thresholds))
	for i, thresholdItem := range retrievedAlert.Thresholds {
		thresholdsDict[i] = map[interface{}]interface{}{
			"value":       thresholdItem.Value,
			"status":      thresholdItem.Status,
			"status_text": thresholdItem.StatusText,
		}
	}

	d.SetId(retrievedAlert.ID)
	d.Set("name", retrievedAlert.Name)
	d.Set("description", retrievedAlert.Description)
	d.Set("type", retrievedAlert.Type)
	d.Set("time_window", retrievedAlert.TimeWindow)
	d.Set("operator", retrievedAlert.Operator)
	d.Set("aggregation", retrievedAlert.Aggregation)
	d.Set("thresholds", thresholdsDict)
	d.Set("rate_unit", retrievedAlert.RateUnit)
	d.Set("rate_value", retrievedAlert.RateValue)
	d.Set("cron_minutes", retrievedAlert.CronMinutes)
	d.Set("cron_hours", retrievedAlert.CronHours)
	d.Set("cron_dom", retrievedAlert.CronDOM)
	d.Set("cron_month", retrievedAlert.CronMonth)
	d.Set("cron_dow", retrievedAlert.CronDOW)
	d.Set("cron_year", retrievedAlert.CronYear)
	d.Set("alert_items", alertItemsDict)
	d.Set("severity", retrievedAlert.Severity)
	d.Set("related_assets", retrievedAlert.RelatedAssets)
	return nil
}

func resourceDeleteAlert(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	itemId := d.Id()

	err := apiClient.DeleteAlert(itemId)
	if err != nil {
		return err
	}
	d.SetId("")
	return nil
}

func resourceExistsAlert(d *schema.ResourceData, m interface{}) (bool, error) {
	apiClient := m.(*client.Client)

	itemId := d.Id()
	_, err := apiClient.RetrieveAlert(itemId)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return false, nil
		} else {
			return false, err
		}
	}
	return true, nil
}
