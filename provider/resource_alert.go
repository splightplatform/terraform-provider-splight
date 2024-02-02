package provider

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/splightplatform/splight-terraform-provider/api/client"
)

func resourceAlert() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"description": {
				Type:     schema.TypeString,
				Required: true,
			},
			"type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"rate_unit": {
				Type:     schema.TypeString,
				Optional: true, // Optional for CronAlert
				Computed: true, // Computed for RateAlert
			},
			"rate_value": {
				Type:     schema.TypeInt,
				Optional: true, // Optional for CronAlert
				Computed: true, // Computed for RateAlert
			},
			"cron_minutes": {
				Type:     schema.TypeInt,
				Optional: true, // Optional for RateAlert
				Computed: true, // Computed for CronAlert
			},
			"cron_hours": {
				Type:     schema.TypeInt,
				Optional: true, // Optional for RateAlert
				Computed: true, // Computed for CronAlert
			},
			"cron_dom": {
				Type:     schema.TypeInt,
				Optional: true, // Optional for RateAlert
				Computed: true, // Computed for CronAlert
			},
			"cron_month": {
				Type:     schema.TypeInt,
				Optional: true, // Optional for RateAlert
				Computed: true, // Computed for CronAlert
			},
			"cron_dow": {
				Type:     schema.TypeInt,
				Optional: true, // Optional for RateAlert
				Computed: true, // Computed for CronAlert
			},
			"cron_year": {
				Type:     schema.TypeInt,
				Optional: true, // Optional for RateAlert
				Computed: true, // Computed for CronAlert
			},
			"time_window": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"thresholds": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"value": {
							Type:     schema.TypeFloat,
							Required: true,
						},
						"status": {
							Type:     schema.TypeString,
							Required: true,
						},
						"status_text": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"severity": {
				Type:     schema.TypeString,
				Required: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"operator": {
				Type:     schema.TypeString,
				Required: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"aggregation": {
				Type:     schema.TypeString,
				Required: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"target_variable": {
				Type:     schema.TypeString,
				Required: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"alert_items": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
		},
		Create: resourceCreateAlert,
		Read:   resourceReadAlert,
		Update: resourceUpdateAlert,
		Delete: resourceDeleteAlert,
		Exists: resourceExistsAlert,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
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
	for i, alertItemItem := range alertItemInterfaceList {
		alertItems[i] = client.AlertItem{
			RefID:           alertItemItem["ref_id"].(string),
			Type:            alertItemItem["type"].(string),
			ExpressionPlain: alertItemItem["expression_plain"].(string),
			QueryPlain:      alertItemItem["query_plain"].(string),
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
	d.Set("thresholds", createdAlert.Thresholds)
	d.Set("rate_unit", createdAlert.RateUnit)
	d.Set("rate_value", createdAlert.RateValue)
	d.Set("cron_minutes", createdAlert.CronMinutes)
	d.Set("cron_hours", createdAlert.CronHours)
	d.Set("cron_dom", createdAlert.CronDOM)
	d.Set("cron_month", createdAlert.CronMonth)
	d.Set("cron_dow", createdAlert.CronDOW)
	d.Set("cron_year", createdAlert.CronYear)
	d.Set("alert_items", createdAlert.AlertItems)
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
	for i, alertItemItem := range alertItemInterfaceList {
		alertItems[i] = client.AlertItem{
			RefID:           alertItemItem["ref_id"].(string),
			Type:            alertItemItem["type"].(string),
			ExpressionPlain: alertItemItem["expression_plain"].(string),
			QueryPlain:      alertItemItem["query_plain"].(string),
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
	d.Set("thresholds", updateAlert.Thresholds)
	d.Set("rate_unit", updateAlert.RateUnit)
	d.Set("rate_value", updateAlert.RateValue)
	d.Set("cron_minutes", updateAlert.CronMinutes)
	d.Set("cron_hours", updateAlert.CronHours)
	d.Set("cron_dom", updateAlert.CronDOM)
	d.Set("cron_month", updateAlert.CronMonth)
	d.Set("cron_dow", updateAlert.CronDOW)
	d.Set("cron_year", updateAlert.CronYear)
	d.Set("alert_items", updateAlert.AlertItems)
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
	for i, alertItemItem := range retrievedAlert.AlertItems {
		alertItemsDict[i] = map[interface{}]interface{}{
			"ref_id":           alertItemItem.RefID,
			"type":             alertItemItem.Type,
			"expression_plain": alertItemItem.ExpressionPlain,
			"query_plain":      alertItemItem.QueryPlain,
		}
	}

	thresholdsDict := make([]map[interface{}]interface{}, len(retrievedAlert.Thresholds))
	for i, thresholdItem := range retrievedAlert.Thresholds {
		alertItemsDict[i] = map[interface{}]interface{}{
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
