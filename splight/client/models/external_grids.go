package models

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

type ExternalGrids struct {
	ExternalGrids []QueryFilter `json:"results"`
}

func (m *ExternalGrids) ResourcePath() string {
	return "v2/engine/asset/external-grids/"
}

func (m *ExternalGrids) ToSchema(d *schema.ResourceData) error {
	var linesMap []map[string]string

	for _, queryFilter := range m.ExternalGrids {
		lineMap := map[string]string{
			"id":   queryFilter.Id,
			"name": queryFilter.Name,
		}
		linesMap = append(linesMap, lineMap)
	}

	d.Set("external_grids", linesMap)
	d.SetId("external_grids")

	return nil
}
