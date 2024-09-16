package models

import "encoding/json"

type QueryFilter struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

// Implement custom JSON marshalling to omit the struct if both fields are empty
func (qf QueryFilter) MarshalJSON() ([]byte, error) {
	if qf.Id == "" && qf.Name == "" {
		return []byte("null"), nil
	}
	type Alias QueryFilter
	return json.Marshal((Alias)(qf))
}

// Implement custom JSON unmarshalling to initialize the struct if the field is null
func (qf *QueryFilter) UnmarshalJSON(data []byte) error {
	type Alias QueryFilter
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(qf),
	}

	if string(data) == "null" {
		*qf = QueryFilter{}
		return nil
	}

	return json.Unmarshal(data, &aux)
}

func convertQueryFilters(data []any) []QueryFilter {
	queryFilters := make([]QueryFilter, len(data))
	for i, queryFilterData := range data {
		queryFilterMap := queryFilterData.(map[string]interface{})
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
		queryFilterMap := data[0].(map[string]interface{})
		queryFilter = QueryFilter{
			Id:   queryFilterMap["id"].(string),
			Name: queryFilterMap["name"].(string),
		}
	}
	return queryFilter
}
