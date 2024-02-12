package provider

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/splightplatform/terraform-provider-splight/api/client"
)

func resourceSecret() *schema.Resource {
	return &schema.Resource{
		Schema: schemaSecret(),
		Create: resourceCreateSecret,
		Read:   resourceReadSecret,
		Update: resourceUpdateSecret,
		Delete: resourceDeleteSecret,
		Exists: resourceExistsSecret,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Description: "Provides a Cloudflare Observatory Scheduled Test resource.",
	}
}

func resourceCreateSecret(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	rawValue := d.Get("raw_value").(string)
	item := client.SecretParams{
		Name:  d.Get("name").(string),
		Value: rawValue,
	}
	createdSecret, err := apiClient.CreateSecret(&item)
	if err != nil {
		return err
	}

	d.SetId(createdSecret.ID)
	d.Set("name", createdSecret.Name)
	d.Set("value", createdSecret.Value)
	return nil
}

func resourceUpdateSecret(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	itemId := d.Id()
	rawValue := d.Get("raw_value").(string)
	item := client.SecretParams{
		Name:  d.Get("name").(string),
		Value: rawValue,
	}

	updatedSecret, err := apiClient.UpdateSecret(itemId, &item)
	if err != nil {
		return err
	}

	d.Set("name", updatedSecret.Name)
	d.Set("value", updatedSecret.Value)
	return nil
}

func resourceReadSecret(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	itemId := d.Id()
	retrievedSecret, err := apiClient.RetrieveSecret(itemId)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			d.SetId("")
		} else {
			return fmt.Errorf("error finding Secret with ID %s", itemId)
		}
	}
	storedValue := d.Get("value")
	if retrievedSecret.Value != storedValue.(string) {
		d.Set("raw_value", nil)
		return nil
	}
	d.SetId(retrievedSecret.ID)
	d.Set("name", retrievedSecret.Name)
	d.Set("value", retrievedSecret.Value)
	return nil
}

func resourceDeleteSecret(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	itemId := d.Id()

	err := apiClient.DeleteSecret(itemId)
	if err != nil {
		return err
	}
	d.SetId("")
	return nil
}

func resourceExistsSecret(d *schema.ResourceData, m interface{}) (bool, error) {
	apiClient := m.(*client.Client)

	itemId := d.Id()
	_, err := apiClient.RetrieveSecret(itemId)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return false, nil
		} else {
			return false, err
		}
	}
	return true, nil
}
