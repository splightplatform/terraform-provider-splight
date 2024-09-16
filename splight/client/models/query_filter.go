package models

import (
	"encoding/json"
	"reflect"
)

type QueryFilter struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type TypedQueryFilter struct {
	Type string `json:"type"`
	Id   string `json:"id"`
	Name string `json:"name"`
}

// Implement custom JSON marshalling to omit the struct if all fields are empty
func excludeZeroedStruct(v any, omitEmptyFields bool) ([]byte, error) {
	vValue := reflect.ValueOf(v)
	vType := reflect.TypeOf(v)

	empty := true
	for i := 0; i < vType.NumField(); i++ {
		if fieldValue := vValue.Field(i).Interface(); fieldValue != "" {
			empty = false
			break
		}
	}
	if empty && omitEmptyFields {
		return []byte("null"), nil
	}
	return json.Marshal(v)
}

// Implement custom JSON unmarshalling to initialize the struct if the field is null
func initNilStruct(data []byte, v any) error {
	if string(data) == "null" {
		reflect.ValueOf(v).Elem().Set(reflect.Zero(reflect.TypeOf(v).Elem()))
		return nil
	}
	return json.Unmarshal(data, v)
}

func (qf QueryFilter) MarshalJSON() ([]byte, error) {
	return excludeZeroedStruct(qf, true)
}

func (qf *QueryFilter) UnmarshalJSON(data []byte) error {
	return initNilStruct(data, qf)
}

func (tqf TypedQueryFilter) MarshalJSON() ([]byte, error) {
	return excludeZeroedStruct(tqf, true)
}

func (tqf *TypedQueryFilter) UnmarshalJSON(data []byte) error {
	return initNilStruct(data, tqf)
}

func convertQueryFilters(data []any) []QueryFilter {
	queryFilters := make([]QueryFilter, len(data))
	for i, queryFilterData := range data {
		queryFilterMap := queryFilterData.(map[string]any)
		queryFilters[i] = QueryFilter{
			Id:   queryFilterMap["id"].(string),
			Name: queryFilterMap["name"].(string),
		}
	}

	return queryFilters
}

func convertSingleQueryFilter(data []any) QueryFilter {
	var queryFilter QueryFilter
	if len(data) > 0 {
		queryFilterMap := data[0].(map[string]any)
		queryFilter = QueryFilter{
			Id:   queryFilterMap["id"].(string),
			Name: queryFilterMap["name"].(string),
		}
	}
	return queryFilter
}

func convertSingleTypedQueryFilter(data []any) TypedQueryFilter {
	var queryFilter TypedQueryFilter
	if len(data) > 0 {
		queryFilterMap := data[0].(map[string]any)
		queryFilter = TypedQueryFilter{
			Id:   queryFilterMap["id"].(string),
			Name: queryFilterMap["name"].(string),
			Type: queryFilterMap["type"].(string),
		}
	}
	return queryFilter
}
