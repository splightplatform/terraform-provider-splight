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
// We could just use a switch too
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

	if err := model.FromSchema(d); err != nil {
		return diag.Errorf("error mapping schema to model: %s", err.Error())
	}

	if err := client.Save(apiClient, model); err != nil {
		return diag.Errorf("error creating resource: %s", err.Error())
	}

	if err := model.ToSchema(d); err != nil {
		return diag.Errorf("error mapping model to schema: %s", err.Error())
	}

	return nil
}

func RetrieveResource[T models.SplightModel](ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	model := InstantiateType[T]()
	apiClient := meta.(*client.Client)

	if err := client.Retrieve(apiClient, model, d.Id()); err != nil {
		if httpErr, ok := err.(*client.HttpError); ok && httpErr.StatusCode == http.StatusNotFound {
			d.SetId("") // Resource not found, clear the Id to remove it from the state
			return nil
		}
		return diag.Errorf("error reading resource with Id '%s': %s", model.GetId(), err.Error())
	}

	if err := model.ToSchema(d); err != nil {
		return diag.Errorf("error mapping model to schema: %s", err.Error())
	}

	return nil
}

func ListDataSource[T models.DataSource](ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	model := InstantiateType[T]()
	apiClient := meta.(*client.Client)

	if err := client.List(apiClient, model); err != nil {
		return diag.Errorf("error listing resource: %s", err.Error())
	}

	if err := model.ToSchema(d); err != nil {
		return diag.Errorf("error mapping model to schema: %s", err.Error())
	}

	return nil
}

func DeleteResource[T models.SplightModel](ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	model := InstantiateType[T]()
	apiClient := meta.(*client.Client)

	if err := client.Delete(apiClient, model, d.Id()); err != nil {
		return diag.Errorf("error deleting resource with Id '%s': %s", d.Id(), err.Error())
	}

	d.SetId("")
	return nil
}
