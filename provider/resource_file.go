package provider

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/splightplatform/splight-terraform-provider/api/client"
)

func resourceFile() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"file": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The path for the file resource",
				ForceNew:    true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
		Create: resourceCreateFile,
		Read:   resourceReadFile,
		Update: resourceUpdateFile,
		Delete: resourceDeleteFile,
		Exists: resourceExistsFile,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
	}
}

func resourceCreateFile(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)
	item := client.FileParams{
		Description: d.Get("description").(string),
	}
	filepath := d.Get("file").(string)
	createdFile, err := apiClient.CreateFile(&item, filepath)
	if err != nil {
		return err
	}

	d.SetId(createdFile.ID)
	d.Set("checksum", createdFile.Checksum)
	// Update json fields
	updatedFile, err := apiClient.UpdateFile(createdFile.ID, &item)
	d.Set("description", updatedFile.Description)
	return nil
}

func resourceUpdateFile(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	itemId := d.Id()
	item := client.FileParams{
		Description: d.Get("description").(string),
	}

	updatedFile, err := apiClient.UpdateFile(itemId, &item)
	if err != nil {
		return err
	}

	d.Set("checksum", updatedFile.Checksum)
	d.Set("description", updatedFile.Description)
	return nil
}

func resourceReadFile(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	itemId := d.Id()
	retrievedFile, err := apiClient.RetrieveFile(itemId)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			d.SetId("")
		} else {
			return fmt.Errorf("error finding File with ID %s", itemId)
		}
	}

	d.SetId(retrievedFile.ID)
	d.Set("checksum", retrievedFile.Checksum)
	d.Set("description", retrievedFile.Description)
	return nil
}

func resourceDeleteFile(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	itemId := d.Id()

	err := apiClient.DeleteFile(itemId)
	if err != nil {
		return err
	}
	d.SetId("")
	return nil
}

func resourceExistsFile(d *schema.ResourceData, m interface{}) (bool, error) {
	apiClient := m.(*client.Client)

	itemId := d.Id()
	_, err := apiClient.RetrieveFile(itemId)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return false, nil
		} else {
			return false, err
		}
	}
	return true, nil
}
