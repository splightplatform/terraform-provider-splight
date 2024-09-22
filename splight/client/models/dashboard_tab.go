package models

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

type DashboardTabParams struct {
	Name      string `json:"name"`
	Order     int    `json:"order"`
	Dashboard string `json:"dashboard"`
}

type DashboardTab struct {
	DashboardTabParams
	Id string `json:"id"`
}

func (m *DashboardTab) GetId() string {
	return m.Id
}

func (m *DashboardTab) GetParams() Params {
	return &m.DashboardTabParams
}

func (m *DashboardTab) ResourcePath() string {
	return "v2/engine/dashboard/charts/"
}

func (m *DashboardTab) FromSchema(d *schema.ResourceData) error {
	m.Id = d.Id()

	m.DashboardTabParams = DashboardTabParams{
		Name:      d.Get("name").(string),
		Order:     d.Get("order").(int),
		Dashboard: d.Get("dashboard").(string),
	}

	return nil
}

func (m *DashboardTab) ToSchema(d *schema.ResourceData) error {
	d.SetId(m.Id)

	d.Set("name", m.Name)
	d.Set("order", m.Order)
	d.Set("dashboard", m.Dashboard)

	return nil
}
