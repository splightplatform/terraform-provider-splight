package models

type DashboardTabParams struct {
	Name      string `json:"name"`
	Order     int    `json:"order"`
	Dashboard string `json:"dashboard"`
}

type DashboardTab struct {
	DashboardTabParams
	ID string `json:"id"`
}
