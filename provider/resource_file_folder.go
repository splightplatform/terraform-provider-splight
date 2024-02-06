package provider

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/splightplatform/splight-terraform-provider/api/client"
)

func resourceFileFolder() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
		Create: resourceCreateFileFolder,
		Read:   resourceReadFileFolder,
		Delete: resourceDeleteFileFolder,
		Exists: resourceExistsFileFolder,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
	}
}

func resourceCreateFileFolder(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)
	item := client.FileFolderParams{
		Name: d.Get("name").(string),
	}
	createdFileFolder, err := apiClient.CreateFileFolder(&item)
	if err != nil {
		return err
	}

	d.SetId(createdFileFolder.ID)
	d.Set("name", createdFileFolder.Name)
	return nil
}

func resourceReadFileFolder(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	itemId := d.Id()
	retrievedFileFolder, err := apiClient.RetrieveFileFolder(itemId)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			d.SetId("")
		} else {
			return fmt.Errorf("error finding FileFolder with ID %s", itemId)
		}
	}

	d.SetId(retrievedFileFolder.ID)
	d.Set("name", retrievedFileFolder.Name)
	return nil
}

func resourceDeleteFileFolder(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	itemId := d.Id()

	err := apiClient.DeleteFileFolder(itemId)
	if err != nil {
		return err
	}
	d.SetId("")
	return nil
}

func resourceExistsFileFolder(d *schema.ResourceData, m interface{}) (bool, error) {
	apiClient := m.(*client.Client)

	itemId := d.Id()
	_, err := apiClient.RetrieveFileFolder(itemId)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return false, nil
		} else {
			return false, err
		}
	}
	return true, nil
}
