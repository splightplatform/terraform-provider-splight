package models

type QueryFilter struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func (m *QueryFilter) isEmpty() bool {
	return m.Id == "" && m.Name == ""
}

func (m *QueryFilter) toMap() []map[string]string {
	return []map[string]string{
		{
			"id":   m.Id,
			"name": m.Name,
		},
	}
}

type TypedQueryFilter struct {
	QueryFilter
	Type string `json:"type"`
}

func (m *TypedQueryFilter) isEmpty() bool {
	return m.Type == "" && m.Id == "" && m.Name == ""
}

func (m *TypedQueryFilter) toMap() []map[string]string {
	return []map[string]string{
		{
			"id":   m.Id,
			"name": m.Name,
			"type": m.Type,
		},
	}
}

// Generic converter for any QueryFilter type
func convertToQueryFilter[T any](data []any) *T {
	if len(data) == 0 {
		return nil
	}
	var qf T
	queryFilterMap := data[0].(map[string]interface{})
	if q, ok := any(&qf).(*QueryFilter); ok {
		q.Id = queryFilterMap["id"].(string)
		q.Name = queryFilterMap["name"].(string)
	}
	if tq, ok := any(&qf).(*TypedQueryFilter); ok {
		tq.Id = queryFilterMap["id"].(string)
		tq.Name = queryFilterMap["name"].(string)
		tq.Type = queryFilterMap["type"].(string)
	}
	return &qf
}

// Conversion functions
func convertSingleQueryFilter(data []any) *QueryFilter {
	return convertToQueryFilter[QueryFilter](data)
}

func convertSingleTypedQueryFilter(data []any) *TypedQueryFilter {
	return convertToQueryFilter[TypedQueryFilter](data)
}

func convertQueryFilters(data []any) []QueryFilter {
	queryFilters := make([]QueryFilter, len(data))
	for i, queryFilterData := range data {
		queryFilters[i] = *convertToQueryFilter[QueryFilter]([]any{queryFilterData})
	}
	return queryFilters
}
