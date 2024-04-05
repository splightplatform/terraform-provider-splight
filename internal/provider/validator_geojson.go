package provider

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/santhosh-tekuri/jsonschema/v5"
	_ "github.com/santhosh-tekuri/jsonschema/v5/httploader"
)

type geoJSONGeometryCollectionValidator struct{}

// Ensure we implement the validator.String interface
var _ validator.String = &geoJSONGeometryCollectionValidator{}

func (v geoJSONGeometryCollectionValidator) Description(ctx context.Context) string {
	return "GeoJSON RFC7946 validation"
}

func (v geoJSONGeometryCollectionValidator) MarkdownDescription(ctx context.Context) string {
	return "GeoJSON RFC7946 validation"
}

// Validate runs the main validation logic of the validator, reading configuration data out of `req` and updating `resp` with diagnostics.
func (v geoJSONGeometryCollectionValidator) ValidateString(ctx context.Context, req validator.StringRequest, resp *validator.StringResponse) {

	// If the value is unknown or null, there is nothing to validate.
	if req.ConfigValue.IsUnknown() || req.ConfigValue.IsNull() {
		return
	}

	var raw_json interface{}
	if err := json.Unmarshal([]byte(req.ConfigValue.ValueString()), &raw_json); err != nil {
		resp.Diagnostics.AddAttributeError(
			req.Path,
			"Error decoding Geometry Collection. Is it a valid JSON?",
			fmt.Sprintf("Error: %s", err),
		)
	}

	schema, err := jsonschema.Compile("https://geojson.org/schema/GeoJSON.json")

	if err != nil {
		resp.Diagnostics.AddAttributeError(
			req.Path,
			"Error while retrieving JSON Schema",
			fmt.Sprintf("Error: %s", err),
		)
	}

	if err = schema.Validate(raw_json); err != nil {
		resp.Diagnostics.AddAttributeError(
			req.Path,
			"Error while validating Geometry Collection",
			"Make sure it satifies RFC7946: https://datatracker.ietf.org/doc/html/rfc7946#section-3.1.8",
		)
	}

}
