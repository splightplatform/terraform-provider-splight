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
			},
			"parent": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
		Create: resourceCreateFileFolder,
		Update: resourceUpdateFileFolder,
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
		Name:   d.Get("name").(string),
		Parent: d.Get("parent").(string),
	}
	createdFileFolder, err := apiClient.CreateFileFolder(&item)
	if err != nil {
		return err
	}

	d.SetId(createdFileFolder.ID)
	d.Set("name", createdFileFolder.Name)
	d.Set("parent", createdFileFolder.Parent)
	return nil
}

func resourceUpdateFileFolder(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	itemId := d.Id()
	item := client.FileFolderParams{
		Parent: d.Get("parent").(string),
	}

	updatedFile, err := apiClient.UpdateFileFolder(itemId, &item)
	if err != nil {
		return err
	}
	d.Set("name", updatedFile.Name)
	d.Set("parent", updatedFile.Parent)
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
	d.Set("parent", retrievedFileFolder.Parent)
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
