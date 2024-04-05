package provider

import (
	"context"
	"fmt"
	"runtime"

	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/xeipuuv/gojsonschema"
)

type geoJSONValidator struct{}

// Ensure we implement the validator.String interface
var _ validator.String = &geoJSONValidator{}

func (v geoJSONValidator) Description(ctx context.Context) string {
	return "GeoJSON RFC7946 validation"
}

func (v geoJSONValidator) MarkdownDescription(ctx context.Context) string {
	return "GeoJSON RFC7946 validation"
}

// Validate runs the main validation logic of the validator, reading configuration data out of `req` and updating `resp` with diagnostics.
func (v geoJSONValidator) ValidateString(ctx context.Context, req validator.StringRequest, resp *validator.StringResponse) {

	// If the value is unknown or null, there is nothing to validate.
	if req.ConfigValue.IsUnknown() || req.ConfigValue.IsNull() {
		return
	}

	runtime.Breakpoint()
	data := req.ConfigValue.ValueString()
	schemaLoader := gojsonschema.NewReferenceLoader("https://geojson.org/schema/GeoJSON.json")
	documentLoader := gojsonschema.NewReferenceLoader(data)

	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	if err != nil {
		panic(err.Error())
	}

	if result.Valid() {
		fmt.Printf("The document is valid\n")
	} else {
		fmt.Printf("The document is not valid. see errors :\n")
		for _, desc := range result.Errors() {
			fmt.Printf("- %s\n", desc)
		}
	}
}
