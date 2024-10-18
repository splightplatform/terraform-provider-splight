package models

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

type Buses struct {
	Buses []QueryFilter `json:"results"`
}

func (m *Buses) ResourcePath() string {
	return "v2/engine/asset/buses/"
}

func (m *Buses) ToSchema(d *schema.ResourceData) error {
	var busesMap []map[string]string

	for _, queryFilter := range m.Buses {
		busMap := map[string]string{
			"id":   queryFilter.Id,
			"name": queryFilter.Name,
		}
		busesMap = append(busesMap, busMap)
	}

	d.Set("buses", busesMap)
	d.SetId("buses")

	return nil
}
