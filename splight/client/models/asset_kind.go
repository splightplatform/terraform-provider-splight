package models

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

type AssetKindParams struct {
	Name string `json:"name"`
}

type AssetKind struct {
	AssetKindParams
	ID string `json:"id"`
}

type AssetKinds struct {
	Kinds []AssetKind `json:"results"`
}

func (m *AssetKinds) ResourcePath() string {
	return "v2/engine/asset/kinds/"
}

func (m *AssetKinds) ToSchema(d *schema.ResourceData) error {
	var kindsMap []map[string]string

	for _, kind := range m.Kinds {
		tagMap := map[string]string{
			"id":   kind.ID,
			"name": kind.Name,
		}
		kindsMap = append(kindsMap, tagMap)
	}

	d.Set("kinds", kindsMap)
	d.SetId("kinds")

	return nil
}
