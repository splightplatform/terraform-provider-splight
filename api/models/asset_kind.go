package models

type AssetKind struct {
	SplightDatabaseBaseModel
	Name string `json:"name"`
}

type AssetKinds struct {
	Kinds []AssetKind `json:"results"`
}
