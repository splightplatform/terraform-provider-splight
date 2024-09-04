package models

import (
	"encoding/json"
)

type RelatedAsset struct {
	Id string `json:"id"`
}

type AssetParams struct {
	Geometry      json.RawMessage `json:"geometry"`
	Name          string          `json:"name"`
	Description   string          `json:"description"`
	RelatedAssets []RelatedAsset  `json:"assets"`
	Tags          []Tag           `json:"tags"`
	Kind          *AssetKind      `json:"kind"`
}

type Asset struct {
	SplightDatabaseBaseModel
	AssetParams
}

func (a Asset) ResourcePath() string {
	return "v2/engine/asset/assets/"
}
