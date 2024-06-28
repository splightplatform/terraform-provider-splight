package provider

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/splightplatform/terraform-provider-splight/api/client"
)

func resourceNode() *schema.Resource {
	return &schema.Resource{
		Schema: schemaNode(),
		Create: resourceCreateNode,
		Read:   resourceReadNode,
		Delete: resourceDeleteNode,
		Exists: resourceExistsNode,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceCreateNode(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	orgId, err := apiClient.RetrieveOrgId()

	if err != nil {
		return fmt.Errorf("Error getting organization id %s", err)
	}

	item := client.NodeParams{
		Name:           d.Get("name").(string),
		InstanceType:   d.Get("instance_type").(string),
		Region:         d.Get("region").(string),
		OrganizationId: orgId,
	}
	createdNode, err := apiClient.CreateNode(&item)
	if err != nil {
		return fmt.Errorf("Error creating node %s", err)
	}
	d.SetId(createdNode.ID)
	return nil
}

func resourceReadNode(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	itemId := d.Id()
	retrievedNode, err := apiClient.RetrieveNode(itemId)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			d.SetId("")
		} else {
			return fmt.Errorf("error finding Node with ID %s", itemId)
		}
	}

	d.SetId(retrievedNode.ID)
	d.Set("name", retrievedNode.Name)
	d.Set("instance_type", retrievedNode.InstanceType)
	d.Set("region", retrievedNode.Region)
	d.Set("organization_id", retrievedNode.OrganizationId)
	return nil
}

func resourceDeleteNode(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	itemId := d.Id()

	err := apiClient.DeleteNode(itemId)
	if err != nil {
		return err
	}
	d.SetId("")
	return nil
}

func resourceExistsNode(d *schema.ResourceData, m interface{}) (bool, error) {
	apiClient := m.(*client.Client)

	itemId := d.Id()
	_, err := apiClient.RetrieveNode(itemId)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return false, nil
		} else {
			return false, err
		}
	}
	return true, nil
}
