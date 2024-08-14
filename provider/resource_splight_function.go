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

func toFunction(d *schema.ResourceData) *client.FunctionParams {
	// Convert target asset
	targetAsset := convertFunctionTargetItem(d.Get("target_asset").(*schema.Set).List())

	// Convert target attribute
	targetAttribute := convertFunctionTargetItem(d.Get("target_attribute").(*schema.Set).List())

	// Convert function items
	functionItems := convertFunctionItems(d.Get("function_items").([]interface{}))

	// Create the FunctionParams object
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

	return &item
}

func convertFunctionTargetItem(targetItemList []interface{}) client.FunctionTargetItem {
	if len(targetItemList) == 0 {
		return client.FunctionTargetItem{}
	}
	valueData := targetItemList[0].(map[string]interface{})
	return client.FunctionTargetItem{
		Name: valueData["name"].(string),
		ID:   valueData["id"].(string),
	}
}

func convertFunctionItems(functionItemsInterface []interface{}) []client.FunctionItem {
	functionItems := make([]client.FunctionItem, len(functionItemsInterface))
	for i, item := range functionItemsInterface {
		functionItem := item.(map[string]interface{})
		queryFilterAsset := functionItem["query_filter_asset"].(*schema.Set).List()[0].(map[string]interface{})
		queryFilterAttribute := functionItem["query_filter_attribute"].(*schema.Set).List()[0].(map[string]interface{})
		queryGroupFunction := functionItem["query_group_function"].(string)
		queryGroupUnit := functionItem["query_group_unit"].(string)
		functionItems[i] = client.FunctionItem{
			RefID:           functionItem["ref_id"].(string),
			Type:            functionItem["type"].(string),
			Expression:      functionItem["expression"].(string),
			ExpressionPlain: functionItem["expression_plain"].(string),
			QueryPlain:      functionItem["query_plain"].(string),
			QueryFilterAsset: client.FunctionTargetItem{
				Name: queryFilterAsset["name"].(string),
				ID:   queryFilterAsset["id"].(string),
			},
			QueryFilterAttribute: client.FunctionTargetItem{
				Name: queryFilterAttribute["name"].(string),
				ID:   queryFilterAttribute["id"].(string),
			},
			QueryGroupFunction: queryGroupFunction,
			QueryGroupUnit:     queryGroupUnit,
		}
	}
	return functionItems
}

func saveFunctionToState(d *schema.ResourceData, function *client.Function) {
	d.SetId(function.ID)

	d.Set("name", function.Name)
	d.Set("description", function.Description)
	d.Set("type", function.Type)
	d.Set("time_window", function.TimeWindow)

	// Since the schemas for these params are 'TypeSet' we must convert
	// our structs to a slice of 'FunctionTargetItem'.
	// Otherwise the SDK will raise a type error.
	d.Set("target_asset", []client.FunctionTargetItem{function.TargetAsset})
	d.Set("target_attribute", []client.FunctionTargetItem{function.TargetAttribute})

	d.Set("rate_unit", function.RateUnit)
	d.Set("rate_value", function.RateValue)
	d.Set("cron_minutes", function.CronMinutes)
	d.Set("cron_hours", function.CronHours)
	d.Set("cron_dom", function.CronDOM)
	d.Set("cron_month", function.CronMonth)
	d.Set("cron_dow", function.CronDOW)
	d.Set("cron_year", function.CronYear)
	d.Set("function_items", function.FunctionItems)
}

func resourceCreateFunction(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	item := toFunction(d)

	createdFunction, err := apiClient.CreateFunction(item)

	if err != nil {
		return fmt.Errorf("error creating Function %s", err.Error())
	}

	saveFunctionToState(d, createdFunction)

	return nil
}

func resourceUpdateFunction(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	itemId := d.Id()

	item := toFunction(d)

	updatedFunction, err := apiClient.UpdateFunction(itemId, item)

	if err != nil {
		return fmt.Errorf("error updating Function with ID '%s': %s", itemId, err.Error())
	}

	saveFunctionToState(d, updatedFunction)

	return nil
}

func resourceReadFunction(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	itemId := d.Id()

	retrievedFunction, err := apiClient.RetrieveFunction(itemId)

	if err != nil {
		return fmt.Errorf("error reading Function with ID '%s': %s", itemId, err.Error())
	}

	saveFunctionToState(d, retrievedFunction)

	return nil
}

func resourceDeleteFunction(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	itemId := d.Id()

	err := apiClient.DeleteFunction(itemId)

	if err != nil {
		return fmt.Errorf("error deleting Function with ID '%s': %s", itemId, err.Error())
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
			return false, fmt.Errorf("error finding Function with ID '%s': %s", itemId, err.Error())
		}
	}

	return true, nil
}
