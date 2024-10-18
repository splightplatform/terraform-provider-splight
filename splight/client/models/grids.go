package models

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

type Grids struct {
	Grids []QueryFilter `json:"results"`
}

func (m *Grids) ResourcePath() string {
	return "v2/engine/asset/grids/"
}

func (m *Grids) ToSchema(d *schema.ResourceData) error {
	var gridsMap []map[string]string

	for _, queryFilter := range m.Grids {
		gridMap := map[string]string{
			"id":   queryFilter.Id,
			"name": queryFilter.Name,
		}
		gridsMap = append(gridsMap, gridMap)
	}

	d.Set("grids", gridsMap)
	d.SetId("grids")

	return nil
}
