package models

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

type Generators struct {
	Generators []QueryFilter `json:"results"`
}

func (m *Generators) ResourcePath() string {
	return "v2/engine/asset/generators/"
}

func (m *Generators) ToSchema(d *schema.ResourceData) error {
	var generatorsMap []map[string]string

	for _, queryFilter := range m.Generators {
		generatorMap := map[string]string{
			"id":   queryFilter.Id,
			"name": queryFilter.Name,
		}
		generatorsMap = append(generatorsMap, generatorMap)
	}

	d.Set("generators", generatorsMap)
	d.SetId("generators")

	return nil
}
