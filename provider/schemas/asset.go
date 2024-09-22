package schemas

import (
	"encoding/json"
	"reflect"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// JSONStringEqualSupressFunc is a function for comparing JSON strings.
func JSONStringEqualSupressFunc(k, old, new string, d *schema.ResourceData) bool {
	return JSONStringEqual(old, new)
}

// JSONStringEqual compares two JSON strings for equality,
// ignoring differences in whitespace and order.
func JSONStringEqual(s1, s2 string) bool {
	return JSONBytesEqual([]byte(s1), []byte(s2))
}

// JSONBytesEqual compares two JSON byte slices for equality,
// ignoring differences in whitespace and order.
func JSONBytesEqual(b1, b2 []byte) bool {
	var o1, o2 interface{}

	// Unmarshal JSON bytes into empty interfaces
	if err := json.Unmarshal(b1, &o1); err != nil {
		return false
	}
	if err := json.Unmarshal(b2, &o2); err != nil {
		return false
	}

	// Compare the unmarshaled objects
	return reflect.DeepEqual(o1, o2)
}

func SchemaAsset() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "name of the resource",
		},
		"description": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "description of the resource",
		},
		"geometry": {
			Type:             schema.TypeString,
			Required:         true,
			Description:      "GeoJSON GeomtryCollection",
			DiffSuppressFunc: JSONStringEqualSupressFunc,
		},
		"tags": {
			Type:        schema.TypeSet,
			Optional:    true,
			Description: "tags of the resource",
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"id": {
						Type:        schema.TypeString,
						Required:    true,
						Description: "tag id",
					},
					"name": {
						Type:        schema.TypeString,
						Required:    true,
						Description: "tag name",
					},
				},
			},
		},
		"kind": {
			Type:        schema.TypeSet,
			Optional:    true,
			MaxItems:    1,
			ForceNew:    true,
			Description: "kind of the resource",
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"id": {
						Type:        schema.TypeString,
						Required:    true,
						ForceNew:    true,
						Description: "kind id",
					},
					"name": {
						Type:        schema.TypeString,
						Required:    true,
						ForceNew:    true,
						Description: "kind name",
					},
				},
			},
		},
	}
}
