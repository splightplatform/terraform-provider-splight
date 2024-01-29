package provider

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/splightplatform/splight-terraform-provider/api/client"
)

func validateName(v interface{}, k string) (ws []string, es []error) {
	var errs []error
	var warns []string
	_, ok := v.(string)
	if !ok {
		errs = append(errs, fmt.Errorf("Expected name to be string"))
		return warns, errs
	}
	return warns, errs
}

func strPtrOrNil(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}

func metadataHash(v interface{}) int {
	item := v.(map[string]interface{})
	return schema.HashString(item["id"].(string))
}

func resourceAsset() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the resource",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "A description of the resource",
			},
			// "metadata": {
			// 	Type:        schema.TypeSet,
			// 	Optional:    true,
			// 	Description: "Metadata to be added to the resource",
			// 	Set:         metadataHash,
			// 	Elem: &schema.Resource{
			// 		Schema: map[string]*schema.Schema{
			// 			"id": {
			// 				Type:     schema.TypeString,
			// 				Computed: true,
			// 			},
			// 			"name": {
			// 				Type:     schema.TypeString,
			// 				Required: true,
			// 			},
			// 			"type": {
			// 				Type:     schema.TypeString,
			// 				Required: true,
			// 			},
			// 			"value": {
			// 				Type:     schema.TypeString,
			// 				Required: true,
			// 			},
			// 			"unit": {
			// 				Type:     schema.TypeString,
			// 				Required: false,
			// 				Optional: true,
			// 			},
			// 		},
			// 	},
			// },
			// "attribute": {
			// 	Type:        schema.TypeSet,
			// 	Optional:    true,
			// 	Description: "Attribute to be added to the resource",
			// 	Elem: &schema.Resource{
			// 		Schema: map[string]*schema.Schema{
			// 			"name": {
			// 				Type:     schema.TypeString,
			// 				Required: true,
			// 			},
			// 			"type": {
			// 				Type:     schema.TypeString,
			// 				Required: true,
			// 			},
			// 			"unit": {
			// 				Type:     schema.TypeString,
			// 				Required: false,
			// 				Optional: true,
			// 			},
			// 		},
			// 	},
			// },
		},
		Create: resourceCreateAsset,
		Read:   resourceReadAsset,
		Update: resourceUpdateAsset,
		Delete: resourceDeleteAsset,
		Exists: resourceExistsAsset,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
	}
}

func resourceCreateAsset(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)
	item := client.AssetParams{
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
		// Metadata:    make([]client.Metadata, 0),
		// Attributes:  make([]client.Attribute, 0),
	}
	createdAsset, err := apiClient.CreateAsset(&item)
	if err != nil {
		return fmt.Errorf("Error creating asset %s", err)
	}
	d.SetId(createdAsset.ID)
	d.Set("name", createdAsset.Name)
	d.Set("description", createdAsset.Description)
	return nil
}

func resourceUpdateAsset(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	// tfMetadatas := d.Get("metadata").(*schema.Set).List()
	// metadata := make([]client.Metadata, len(tfMetadatas))
	// for i, item := range tfMetadatas {
	// 	metadataItem, _ := item.(map[string]interface{})
	// 	metadata[i] = client.Metadata{
	// 		MetadataParams: client.MetadataParams{
	// 			Name:  metadataItem["name"].(string),
	// 			Value: metadataItem["value"].(string),
	// 			Type:  metadataItem["type"].(string),
	// 			Unit:  strPtrOrNil(metadataItem["unit"].(string)),
	// 		},
	// 		ID: strPtrOrNil(metadataItem["id"].(string)),
	// 	}
	// }

	// tfAttributes := d.Get("attribute").(*schema.Set).List()
	// attributes := make([]client.Attribute, len(tfAttributes))
	// for i, m := range tfAttributes {
	// 	attributeItem := m.(map[string]interface{})
	// 	attributes[i] = client.Attribute{
	// 		AttributeParams: client.AttributeParams{
	// 			Name: attributeItem["name"].(string),
	// 			Type: attributeItem["type"].(string),
	// 			Unit: strPtrOrNil(attributeItem["unit"].(string)),
	// 		},
	// 		ID: strPtrOrNil(attributeItem["id"].(string)),
	// 	}
	// }
	itemId := d.Id()
	item := client.AssetParams{
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
		// Metadata:    metadata,
		// Attributes:  make([]client.Attribute, 0),
	}
	updatedAsset, err := apiClient.UpdateAsset(itemId, &item)

	if err != nil {
		return err
	}
	d.Set("name", updatedAsset.Name)
	d.Set("description", updatedAsset.Description)
	// d.Set("metadata", updatedAsset.Metadata)
	// d.Set("attribute", updatedAsset.Attributes)
	return nil
}

func resourceReadAsset(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	itemId := d.Id()
	retrievedAsset, err := apiClient.RetrieveAsset(itemId)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			d.SetId("")
		} else {
			return fmt.Errorf("error finding Asset with ID %s", itemId)
		}
	}

	d.SetId(retrievedAsset.ID)
	d.Set("name", retrievedAsset.Name)
	d.Set("description", retrievedAsset.Description)
	// d.Set("metadata", retrievedAsset.Metadata)
	// d.Set("attribute", retrievedAsset.Attributes)
	return nil
}

func resourceDeleteAsset(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	itemId := d.Id()

	err := apiClient.DeleteAsset(itemId)
	if err != nil {
		return err
	}
	d.SetId("")
	return nil
}

func resourceExistsAsset(d *schema.ResourceData, m interface{}) (bool, error) {
	apiClient := m.(*client.Client)

	itemId := d.Id()
	_, err := apiClient.RetrieveAsset(itemId)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return false, nil
		} else {
			return false, err
		}
	}
	return true, nil
}
