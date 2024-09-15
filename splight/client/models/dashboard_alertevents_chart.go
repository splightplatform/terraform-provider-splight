package models

type DashboardAlertEventsChartParams struct {
	DashboardChartParams
	Type            string   `json:"type"`
	FilterName      string   `json:"filter_name"`
	FilterOldStatus []string `json:"filter_old_status"`
	FilterNewStatus []string `json:"filter_new_status"`
}

type DashboardAlertEventsChart struct {
	DashboardAlertEventsChartParams
	ID string `json:"id"`
}
