package provider

import (
	"context"
	"net/http"
	"reflect"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/splightplatform/terraform-provider-splight/splight/client"
	"github.com/splightplatform/terraform-provider-splight/splight/client/models"
)

// InstantiateType creates a new instance of type T, ensuring that T is a pointer type
func InstantiateType[T models.SplightObject]() T {
	var model T
	modelType := reflect.TypeOf(model)
	return reflect.New(modelType.Elem()).Interface().(T)
}

func resourceForType[T models.SplightModel](schemaFunc func() map[string]*schema.Schema) *schema.Resource {
	return &schema.Resource{
		Schema:        schemaFunc(),
		CreateContext: SaveResource[T],
		UpdateContext: SaveResource[T],
		ReadContext:   RetrieveResource[T],
		DeleteContext: DeleteResource[T],
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func dataSourceForType[T models.DataSource](schemaFunc func() map[string]*schema.Schema) *schema.Resource {
	return &schema.Resource{
		Schema:      schemaFunc(),
		ReadContext: ListDataSource[T],
	}
}

func SaveResource[T models.SplightModel](ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	model := InstantiateType[T]()
	apiClient := meta.(*client.Client)

	err := model.FromSchema(d)
	if err != nil {
		return diag.Errorf("error mapping schema to model: %s", err.Error())
	}

	err = client.Save(apiClient, model)
	if err != nil {
		return diag.Errorf("error creating resource: %s", err.Error())
	}

	model.ToSchema(d)

	return nil
}

func RetrieveResource[T models.SplightModel](ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	model := InstantiateType[T]()
	apiClient := meta.(*client.Client)

	err := client.Retrieve(apiClient, model, d.Id())
	if err != nil {
		if httpErr, ok := err.(*client.HttpError); ok && httpErr.StatusCode == http.StatusNotFound {
			d.SetId("") // Resource not found, clear the ID to remove it from the state
			return nil
		}
		return diag.Errorf("error reading resource with ID '%s': %s", model.GetID(), err.Error())
	}

	model.ToSchema(d)

	return nil
}

func ListDataSource[T models.DataSource](ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	model := InstantiateType[T]()
	apiClient := meta.(*client.Client)

	err := client.List(apiClient, model)
	if err != nil {
		return diag.Errorf("error listing resource: %s", err.Error())
	}

	model.ToSchema(d)

	return nil
}

func DeleteResource[T models.SplightModel](ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	model := InstantiateType[T]()
	apiClient := meta.(*client.Client)

	err := client.Delete(apiClient, model, d.Id())
	if err != nil {
		return diag.Errorf("error deleting resource with ID '%s': %s", d.Id(), err.Error())
	}

	d.SetId("") // Clear the resource ID to remove it from the state
	return nil
}
