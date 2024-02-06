package provider

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/splightplatform/splight-terraform-provider/api/client"
)

func resourceDashboard() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
		Create: resourceCreateDashboard,
		Read:   resourceReadDashboard,
		Update: resourceUpdateDashboard,
		Delete: resourceDeleteDashboard,
		Exists: resourceExistsDashboard,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceCreateDashboard(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)
	item := client.DashboardParams{
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
	}
	createdDashboard, err := apiClient.CreateDashboard(&item)
	if err != nil {
		return err
	}

	d.SetId(createdDashboard.ID)
	d.Set("name", createdDashboard.Name)
	d.Set("description", createdDashboard.Description)
	return nil
}

func resourceUpdateDashboard(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	itemId := d.Id()
	item := client.DashboardParams{
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
	}

	updatedDashboard, err := apiClient.UpdateDashboard(itemId, &item)
	if err != nil {
		return err
	}

	d.Set("name", updatedDashboard.Name)
	d.Set("description", updatedDashboard.Description)
	return nil
}

func resourceReadDashboard(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	itemId := d.Id()
	retrievedDashboard, err := apiClient.RetrieveDashboard(itemId)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			d.SetId("")
		} else {
			return fmt.Errorf("error finding Dashboard with ID %s", itemId)
		}
	}

	d.SetId(retrievedDashboard.ID)
	d.Set("name", retrievedDashboard.Name)
	d.Set("description", retrievedDashboard.Description)
	return nil
}

func resourceDeleteDashboard(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	itemId := d.Id()

	err := apiClient.DeleteDashboard(itemId)
	if err != nil {
		return err
	}
	d.SetId("")
	return nil
}

func resourceExistsDashboard(d *schema.ResourceData, m interface{}) (bool, error) {
	apiClient := m.(*client.Client)

	itemId := d.Id()
	_, err := apiClient.RetrieveDashboard(itemId)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return false, nil
		} else {
			return false, err
		}
	}
	return true, nil
}
