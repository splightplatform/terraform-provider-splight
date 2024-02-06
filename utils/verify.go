package utils

import (
	"bytes"
	"encoding/json"
	"reflect"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ValidateNullableString(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}

func JSONStringEqualSupressFunc(k, old, new string, d *schema.ResourceData) bool {
	return JSONStringEqual(old, new)
}

func JSONStringEqual(s1, s2 string) bool {
	b1 := bytes.NewBufferString("")
	if err := json.Compact(b1, []byte(s1)); err != nil {
		return false
	}

	b2 := bytes.NewBufferString("")
	if err := json.Compact(b2, []byte(s2)); err != nil {
		return false
	}

	return JSONBytesEqual(b1.Bytes(), b2.Bytes())
}
func JSONBytesEqual(b1, b2 []byte) bool {
	var o1 interface{}
	if err := json.Unmarshal(b1, &o1); err != nil {
		return false
	}

	var o2 interface{}
	if err := json.Unmarshal(b2, &o2); err != nil {
		return false
	}

	return reflect.DeepEqual(o1, o2)
}
