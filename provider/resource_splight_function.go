package provider

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/splightplatform/terraform-provider-splight/api/client"
)

func resourceFunction() *schema.Resource {
	return &schema.Resource{
		Schema: schemaFunction(),
		Create: resourceCreateFunction,
		Read:   resourceReadFunction,
		Update: resourceUpdateFunction,
		Delete: resourceDeleteFunction,
		Exists: resourceExistsFunction,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceCreateFunction(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	valueList := d.Get("target_asset").(*schema.Set).List()
	var targetAsset client.FunctionTargetItem
	for _, value := range valueList {
		valueData := value.(map[string]interface{})
		targetAsset = client.FunctionTargetItem{
			Name: valueData["name"].(string),
			ID:   valueData["id"].(string),
		}
	}

	valueList = d.Get("target_attribute").(*schema.Set).List()
	var targetAttribute client.FunctionTargetItem
	for _, value := range valueList {
		valueData := value.(map[string]interface{})
		targetAttribute = client.FunctionTargetItem{
			Name: valueData["name"].(string),
			ID:   valueData["id"].(string),
		}
	}

	functionItemInterface := d.Get("function_items").([]interface{})
	functionItemInterfaceList := make([]map[string]interface{}, len(functionItemInterface))
	for i, functionItemInterfaceItem := range functionItemInterface {
		functionItemInterfaceList[i] = functionItemInterfaceItem.(map[string]interface{})
	}
	functionItems := make([]client.FunctionItem, len(functionItemInterfaceList))
	for i, functionItemItem := range functionItemInterfaceList {
		functionItems[i] = client.FunctionItem{
			RefID:           functionItemItem["ref_id"].(string),
			Type:            functionItemItem["type"].(string),
			ExpressionPlain: functionItemItem["expression_plain"].(string),
			QueryPlain:      functionItemItem["query_plain"].(string),
		}
	}

	item := client.FunctionParams{
		Name:            d.Get("name").(string),
		Description:     d.Get("description").(string),
		Type:            d.Get("type").(string),
		TimeWindow:      d.Get("time_window").(int),
		RateUnit:        d.Get("rate_unit").(string),
		RateValue:       d.Get("rate_value").(int),
		TargetVariable:  d.Get("target_variable").(string),
		TargetAsset:     targetAsset,
		TargetAttribute: targetAttribute,
		FunctionItems:   functionItems,
	}

	createdFunction, err := apiClient.CreateFunction(&item)
	if err != nil {
		return err
	}

	d.SetId(createdFunction.ID)
	d.Set("name", createdFunction.Name)
	d.Set("description", createdFunction.Description)
	d.Set("type", createdFunction.Type)
	d.Set("time_window", createdFunction.TimeWindow)

	targetAssetOutput := make([]map[string]interface{}, 1)
	targetAssetOutput[0] = map[string]interface{}{
		"name": createdFunction.TargetAsset.Name,
		"id":   createdFunction.TargetAsset.ID,
	}
	targetAttributeOutput := make([]map[string]interface{}, 1)
	targetAttributeOutput[0] = map[string]interface{}{
		"name": createdFunction.TargetAttribute.Name,
		"id":   createdFunction.TargetAttribute.ID,
	}
	d.Set("target_asset", targetAssetOutput)
	d.Set("target_attribute", targetAttributeOutput)
	d.Set("rate_unit", createdFunction.RateUnit)
	d.Set("rate_value", createdFunction.RateValue)
	d.Set("cron_minutes", createdFunction.CronMinutes)
	d.Set("cron_hours", createdFunction.CronHours)
	d.Set("cron_dom", createdFunction.CronDOM)
	d.Set("cron_month", createdFunction.CronMonth)
	d.Set("cron_dow", createdFunction.CronDOW)
	d.Set("cron_year", createdFunction.CronYear)
	functionItemsOutput := make([]map[string]interface{}, len(createdFunction.FunctionItems))
	for i, functionItemItem := range createdFunction.FunctionItems {
		functionItemsOutput[i] = map[string]interface{}{
			"id":               functionItemItem.ID,
			"ref_id":           functionItemItem.RefID,
			"type":             functionItemItem.Type,
			"expression_plain": functionItemItem.ExpressionPlain,
			"query_plain":      functionItemItem.QueryPlain,
		}
	}
	d.Set("function_items", functionItemsOutput)

	return nil
}

func resourceUpdateFunction(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	itemId := d.Id()

	valueList := d.Get("target_asset").(*schema.Set).List()
	var targetAsset client.FunctionTargetItem
	for _, value := range valueList {
		valueData := value.(map[string]interface{})
		targetAsset = client.FunctionTargetItem{
			Name: valueData["name"].(string),
			ID:   valueData["id"].(string),
		}
	}

	valueList = d.Get("target_attribute").(*schema.Set).List()
	var targetAttribute client.FunctionTargetItem
	for _, value := range valueList {
		valueData := value.(map[string]interface{})
		targetAttribute = client.FunctionTargetItem{
			Name: valueData["name"].(string),
			ID:   valueData["id"].(string),
		}
	}

	functionItemInterface := d.Get("function_items").([]interface{})
	functionItemInterfaceList := make([]map[string]interface{}, len(functionItemInterface))
	for i, functionItemInterfaceItem := range functionItemInterface {
		functionItemInterfaceList[i] = functionItemInterfaceItem.(map[string]interface{})
	}
	functionItems := make([]client.FunctionItem, len(functionItemInterfaceList))
	for i, functionItemItem := range functionItemInterfaceList {
		functionItems[i] = client.FunctionItem{
			ID:              functionItemItem["id"].(string),
			RefID:           functionItemItem["ref_id"].(string),
			Type:            functionItemItem["type"].(string),
			ExpressionPlain: functionItemItem["expression_plain"].(string),
			QueryPlain:      functionItemItem["query_plain"].(string),
		}
	}

	item := client.FunctionParams{
		Name:            d.Get("name").(string),
		Description:     d.Get("description").(string),
		Type:            d.Get("type").(string),
		TimeWindow:      d.Get("time_window").(int),
		RateUnit:        d.Get("rate_unit").(string),
		RateValue:       d.Get("rate_value").(int),
		TargetVariable:  d.Get("target_variable").(string),
		TargetAsset:     targetAsset,
		TargetAttribute: targetAttribute,
		FunctionItems:   functionItems,
	}
	updatedFunction, err := apiClient.UpdateFunction(itemId, &item)
	if err != nil {
		return err
	}

	d.Set("name", updatedFunction.Name)
	d.Set("description", updatedFunction.Description)
	d.Set("type", updatedFunction.Type)
	d.Set("time_window", updatedFunction.TimeWindow)
	targetAssetOutput := make([]map[string]interface{}, 1)
	targetAssetOutput[0] = map[string]interface{}{
		"name": updatedFunction.TargetAsset.Name,
		"id":   updatedFunction.TargetAsset.ID,
	}
	targetAttributeOutput := make([]map[string]interface{}, 1)
	targetAttributeOutput[0] = map[string]interface{}{
		"name": updatedFunction.TargetAttribute.Name,
		"id":   updatedFunction.TargetAttribute.ID,
	}
	d.Set("target_asset", targetAssetOutput)
	d.Set("target_attribute", targetAttributeOutput)
	d.Set("rate_unit", updatedFunction.RateUnit)
	d.Set("rate_value", updatedFunction.RateValue)
	d.Set("cron_minutes", updatedFunction.CronMinutes)
	d.Set("cron_hours", updatedFunction.CronHours)
	d.Set("cron_dom", updatedFunction.CronDOM)
	d.Set("cron_month", updatedFunction.CronMonth)
	d.Set("cron_dow", updatedFunction.CronDOW)
	d.Set("cron_year", updatedFunction.CronYear)
	functionItemsOutput := make([]map[string]interface{}, len(updatedFunction.FunctionItems))
	for i, functionItemItem := range updatedFunction.FunctionItems {
		functionItemsOutput[i] = map[string]interface{}{
			"id":               functionItemItem.ID,
			"ref_id":           functionItemItem.RefID,
			"type":             functionItemItem.Type,
			"expression_plain": functionItemItem.ExpressionPlain,
			"query_plain":      functionItemItem.QueryPlain,
		}
	}
	d.Set("function_items", functionItemsOutput)

	return nil
}

func resourceReadFunction(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	itemId := d.Id()
	retrievedFunction, err := apiClient.RetrieveFunction(itemId)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			d.SetId("")
		} else {
			return fmt.Errorf("error finding Function with ID %s", itemId)
		}
	}

	functionItemsDict := make([]map[interface{}]interface{}, len(retrievedFunction.FunctionItems))
	for i, functionItemItem := range retrievedFunction.FunctionItems {
		functionItemsDict[i] = map[interface{}]interface{}{
			"id":               functionItemItem.ID,
			"ref_id":           functionItemItem.RefID,
			"type":             functionItemItem.Type,
			"expression_plain": functionItemItem.ExpressionPlain,
			"query_plain":      functionItemItem.QueryPlain,
		}
	}

	targetAsset := make([]map[string]string, 1)
	targetAsset[0] = map[string]string{
		"name": retrievedFunction.TargetAsset.Name,
		"id":   retrievedFunction.TargetAsset.ID,
	}

	targetAttribute := make([]map[string]string, 1)
	targetAttribute[0] = map[string]string{
		"name": retrievedFunction.TargetAttribute.Name,
		"id":   retrievedFunction.TargetAttribute.ID,
	}

	d.SetId(retrievedFunction.ID)
	d.Set("name", retrievedFunction.Name)
	d.Set("description", retrievedFunction.Description)
	d.Set("name", retrievedFunction.Name)
	d.Set("description", retrievedFunction.Description)
	d.Set("type", retrievedFunction.Type)
	d.Set("time_window", retrievedFunction.TimeWindow)
	d.Set("target_asset", targetAsset)
	d.Set("target_attribute", targetAttribute)
	d.Set("rate_unit", retrievedFunction.RateUnit)
	d.Set("rate_value", retrievedFunction.RateValue)
	d.Set("cron_minutes", retrievedFunction.CronMinutes)
	d.Set("cron_hours", retrievedFunction.CronHours)
	d.Set("cron_dom", retrievedFunction.CronDOM)
	d.Set("cron_month", retrievedFunction.CronMonth)
	d.Set("cron_dow", retrievedFunction.CronDOW)
	d.Set("cron_year", retrievedFunction.CronYear)
	d.Set("function_items", functionItemsDict)
	return nil
}

func resourceDeleteFunction(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	itemId := d.Id()

	err := apiClient.DeleteFunction(itemId)
	if err != nil {
		return err
	}
	d.SetId("")
	return nil
}

func resourceExistsFunction(d *schema.ResourceData, m interface{}) (bool, error) {
	apiClient := m.(*client.Client)

	itemId := d.Id()
	_, err := apiClient.RetrieveFunction(itemId)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return false, nil
		} else {
			return false, err
		}
	}
	return true, nil
}
