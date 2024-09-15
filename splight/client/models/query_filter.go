package models

type QueryFilter struct {
	Id   string `json:"id"`
	Name string `json:"name"`
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
