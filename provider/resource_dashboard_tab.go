package provider

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/splightplatform/splight-terraform-provider/api/client"
)

func resourceDashboardTab() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"order": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"dashboard": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
		Create: resourceCreateDashboardTab,
		Read:   resourceReadDashboardTab,
		Update: resourceUpdateDashboardTab,
		Delete: resourceDeleteDashboardTab,
		Exists: resourceExistsDashboardTab,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
	}
}

func resourceCreateDashboardTab(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)
	item := client.DashboardTabParams{
		Name:      d.Get("name").(string),
		Order:     d.Get("order").(int),
		Dashboard: d.Get("dashboard").(string),
	}
	createdDashboardTab, err := apiClient.CreateDashboardTab(&item)
	if err != nil {
		return err
	}

	d.SetId(createdDashboardTab.ID)
	d.Set("name", createdDashboardTab.Name)
	d.Set("order", createdDashboardTab.Order)
	d.Set("dashboard", createdDashboardTab.Dashboard)
	return nil
}

func resourceUpdateDashboardTab(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	itemId := d.Id()
	item := client.DashboardTabParams{
		Name:      d.Get("name").(string),
		Dashboard: d.Get("dashboard").(string),
		Order:     d.Get("order").(int),
	}

	updatedDashboardTab, err := apiClient.UpdateDashboardTab(itemId, &item)
	if err != nil {
		return err
	}

	d.Set("name", updatedDashboardTab.Name)
	d.Set("order", updatedDashboardTab.Order)
	d.Set("dashboard", updatedDashboardTab.Dashboard)
	return nil
}

func resourceReadDashboardTab(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	itemId := d.Id()
	retrievedDashboardTab, err := apiClient.RetrieveDashboardTab(itemId)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			d.SetId("")
		} else {
			return fmt.Errorf("error finding DashboardTab with ID %s", itemId)
		}
	}

	d.SetId(retrievedDashboardTab.ID)
	d.Set("name", retrievedDashboardTab.Name)
	d.Set("dashboard", retrievedDashboardTab.Dashboard)
	d.Set("order", retrievedDashboardTab.Order)

	return nil
}

func resourceDeleteDashboardTab(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	itemId := d.Id()

	err := apiClient.DeleteDashboardTab(itemId)
	if err != nil {
		return err
	}
	d.SetId("")
	return nil
}

func resourceExistsDashboardTab(d *schema.ResourceData, m interface{}) (bool, error) {
	apiClient := m.(*client.Client)

	itemId := d.Id()
	_, err := apiClient.RetrieveDashboardTab(itemId)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return false, nil
		} else {
			return false, err
		}
	}
	return true, nil
}
