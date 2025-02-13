package models

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

type ExternalGrids struct {
	ExternalGrids []QueryFilter `json:"results"`
}

func (m *ExternalGrids) ResourcePath() string {
	return "v3/engine/asset/external-grids/"
}

func (m *ExternalGrids) ToSchema(d *schema.ResourceData) error {
	var gridsMap []map[string]string

	for _, queryFilter := range m.ExternalGrids {
		gridMap := map[string]string{
			"id":   queryFilter.Id,
			"name": queryFilter.Name,
		}
		gridsMap = append(gridsMap, gridMap)
	}

	d.Set("external_grids", gridsMap)
	d.SetId("external_grids")

	return nil
}
