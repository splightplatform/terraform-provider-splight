package models

type DashboardParams struct {
	Name          string         `json:"name"`
	Description   string         `json:"description"`
	RelatedAssets []RelatedAsset `json:"assets"`
}

type Dashboard struct {
	DashboardParams
	ID string `json:"id"`
}
