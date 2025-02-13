package models

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

type AssetKinds struct {
	Kinds []QueryFilter `json:"results"`
}

func (m *AssetKinds) ResourcePath() string {
	return "v3/engine/asset/kinds/"
}

func (m *AssetKinds) ToSchema(d *schema.ResourceData) error {
	var kindsMap []map[string]string

	for _, queryFilter := range m.Kinds {
		kindMap := map[string]string{
			"id":   queryFilter.Id,
			"name": queryFilter.Name,
		}
		kindsMap = append(kindsMap, kindMap)
	}

	d.Set("kinds", kindsMap)
	d.SetId("kinds")

	return nil
}
