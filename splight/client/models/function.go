package models

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type FunctionItem struct {
	ID                   string                  `json:"id,omitempty"`
	RefID                string                  `json:"ref_id"`
	Type                 string                  `json:"type"`
	Expression           string                  `json:"expression"`
	ExpressionPlain      string                  `json:"expression_plain"`
	QueryPlain           string                  `json:"query_plain"`
	QueryFilterAsset     FunctionTargetItem      `json:"query_filter_asset"`
	QueryFilterAttribute TypedFunctionTargetItem `json:"query_filter_attribute"`
	QueryGroupFunction   string                  `json:"query_group_function"`
	QueryGroupUnit       string                  `json:"query_group_unit"`
}

type TypedFunctionTargetItem struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type FunctionTargetItem struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// Implement custom JSON marshalling to omit the struct if both fields are empty
func (fti TypedFunctionTargetItem) MarshalJSON() ([]byte, error) {
	if fti.ID == "" && fti.Name == "" && fti.Type == "" {
		return []byte("null"), nil
	}
	type Alias TypedFunctionTargetItem
	return json.Marshal((Alias)(fti))
}

// Implement custom JSON unmarshalling to initialize the struct if the field is null
func (fti *TypedFunctionTargetItem) UnmarshalJSON(data []byte) error {
	type Alias TypedFunctionTargetItem
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(fti),
	}

	if string(data) == "null" {
		*fti = TypedFunctionTargetItem{}
		return nil
	}

	return json.Unmarshal(data, &aux)
}

// Implement custom JSON marshalling to omit the struct if both fields are empty
func (fti FunctionTargetItem) MarshalJSON() ([]byte, error) {
	if fti.ID == "" && fti.Name == "" {
		return []byte("null"), nil
	}
	type Alias FunctionTargetItem
	return json.Marshal((Alias)(fti))
}

// Implement custom JSON unmarshalling to initialize the struct if the field is null
func (fti *FunctionTargetItem) UnmarshalJSON(data []byte) error {
	type Alias FunctionTargetItem
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(fti),
	}

	if string(data) == "null" {
		*fti = FunctionTargetItem{}
		return nil
	}

	return json.Unmarshal(data, &aux)
}

type FunctionParams struct {
	Name            string                  `json:"name"`
	Description     string                  `json:"description"`
	Type            string                  `json:"type"`
	TimeWindow      int                     `json:"time_window"`
	TargetAsset     FunctionTargetItem      `json:"target_asset"`
	TargetAttribute TypedFunctionTargetItem `json:"target_attribute"`
	TargetVariable  string                  `json:"target_variable"`
	RateUnit        string                  `json:"rate_unit"`
	RateValue       int                     `json:"rate_value"`
	CronMinutes     int                     `json:"cron_minutes"`
	CronHours       int                     `json:"cron_hours"`
	CronDOM         int                     `json:"cron_dom"`
	CronMonth       int                     `json:"cron_month"`
	CronDOW         int                     `json:"cron_dow"`
	CronYear        int                     `json:"cron_year"`
	FunctionItems   []FunctionItem          `json:"function_items"`
}

type Function struct {
	FunctionParams
	ID string `json:"id"`
}

func (m *Function) GetID() string {
	return m.ID
}

func (m *Function) GetParams() Params {
	return &m.FunctionParams
}

func (m *Function) ResourcePath() string {
	return "v2/engine/function/functions/"
}

func (m *Function) FromSchema(d *schema.ResourceData) error {
	// Convert target asset
	targetAsset := convertFunctionTargetItem(
		d.Get("target_asset").(*schema.Set).List(),
		false,
	)

	// Convert target attribute
	targetAttribute := convertFunctionTargetItem(
		d.Get("target_attribute").(*schema.Set).List(),
		true,
	)

	// Convert function items
	functionItems := convertFunctionItems(d.Get("function_items").([]interface{}))

	// Create the FunctionParams object
	m.FunctionParams = FunctionParams{
		Name:            d.Get("name").(string),
		Description:     d.Get("description").(string),
		Type:            d.Get("type").(string),
		TimeWindow:      d.Get("time_window").(int),
		RateUnit:        d.Get("rate_unit").(string),
		RateValue:       d.Get("rate_value").(int),
		TargetVariable:  d.Get("target_variable").(string),
		TargetAsset:     targetAsset.(FunctionTargetItem),
		TargetAttribute: targetAttribute.(TypedFunctionTargetItem),
		FunctionItems:   functionItems,
	}

	return nil
}

func convertFunctionTargetItem(targetItemList []interface{}, typed bool) interface{} {
	if len(targetItemList) == 0 {
		if typed {
			return TypedFunctionTargetItem{}
		}
		return FunctionTargetItem{}
	}

	valueData := targetItemList[0].(map[string]interface{})

	if typed {
		return TypedFunctionTargetItem{
			Name: valueData["name"].(string),
			ID:   valueData["id"].(string),
			Type: valueData["type"].(string),
		}
	}

	return FunctionTargetItem{
		Name: valueData["name"].(string),
		ID:   valueData["id"].(string),
	}
}

func convertFunctionItems(functionItemsInterface []interface{}) []FunctionItem {
	functionItems := make([]FunctionItem, len(functionItemsInterface))
	for i, item := range functionItemsInterface {
		functionItem := item.(map[string]interface{})
		queryFilterAsset := functionItem["query_filter_asset"].(*schema.Set).List()[0].(map[string]interface{})
		queryFilterAttribute := functionItem["query_filter_attribute"].(*schema.Set).List()[0].(map[string]interface{})
		queryGroupFunction := functionItem["query_group_function"].(string)
		queryGroupUnit := functionItem["query_group_unit"].(string)
		functionItems[i] = FunctionItem{
			RefID:           functionItem["ref_id"].(string),
			Type:            functionItem["type"].(string),
			Expression:      functionItem["expression"].(string),
			ExpressionPlain: functionItem["expression_plain"].(string),
			QueryPlain:      functionItem["query_plain"].(string),
			QueryFilterAsset: FunctionTargetItem{
				Name: queryFilterAsset["name"].(string),
				ID:   queryFilterAsset["id"].(string),
			},
			QueryFilterAttribute: TypedFunctionTargetItem{
				Name: queryFilterAttribute["name"].(string),
				ID:   queryFilterAttribute["id"].(string),
				Type: queryFilterAttribute["type"].(string),
			},
			QueryGroupFunction: queryGroupFunction,
			QueryGroupUnit:     queryGroupUnit,
		}
	}
	return functionItems
}

func (m *Function) ToSchema(d *schema.ResourceData) error {
	d.SetId(m.ID)

	d.Set("name", m.Name)
	d.Set("description", m.Description)
	d.Set("type", m.Type)
	d.Set("time_window", m.TimeWindow)

	// Since the schemas for these params are 'TypeSet' we must convert
	// our structs to a slice of 'FunctionTargetItem'.
	// Otherwise the SDK will raise a type error.
	d.Set("target_asset", []map[string]interface{}{
		{
			"id":   m.TargetAsset.ID,
			"name": m.TargetAsset.Name,
		},
	})
	d.Set("target_attribute", []map[string]interface{}{
		{
			"id":   m.TargetAttribute.ID,
			"name": m.TargetAttribute.Name,
			"type": m.TargetAttribute.Type,
		},
	})

	d.Set("rate_unit", m.RateUnit)
	d.Set("rate_value", m.RateValue)
	d.Set("cron_minutes", m.CronMinutes)
	d.Set("cron_hours", m.CronHours)
	d.Set("cron_dom", m.CronDOM)
	d.Set("cron_month", m.CronMonth)
	d.Set("cron_dow", m.CronDOW)
	d.Set("cron_year", m.CronYear)

	functionItems := make([]map[string]interface{}, len(m.FunctionItems))
	for i, function := range m.FunctionItems {
		functionItems[i] = map[string]interface{}{
			"id":               function.ID,
			"ref_id":           function.RefID,
			"type":             function.Type,
			"expression":       function.Expression,
			"expression_plain": function.ExpressionPlain,
			"query_plain":      function.QueryPlain,
			"query_filter_asset": []map[string]interface{}{
				{
					"id":   function.QueryFilterAsset.ID,
					"name": function.QueryFilterAsset.Name,
				},
			},
			"query_filter_attribute": []map[string]interface{}{
				{
					"id":   function.QueryFilterAttribute.ID,
					"name": function.QueryFilterAttribute.Name,
					"type": function.QueryFilterAttribute.Type,
				},
			},
			"query_group_function": function.QueryGroupFunction,
			"query_group_unit":     function.QueryGroupUnit,
		}
	}

	d.Set("function_items", functionItems)

	return nil
}
