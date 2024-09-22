package schemas

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// HashStringMD5 computes the MD5 hash of a string. It returns the hash as a hexadecimal string.
// If the input is not a string, it returns an empty string and logs an error.
func HashStringMD5(data interface{}) string {
	str, ok := data.(string)
	if !ok {
		// Handle the case where data is not a string
		fmt.Println("Error: HashStringMD5 requires a string input")
		return ""
	}

	hasher := md5.New()
	hasher.Write([]byte(str))
	return hex.EncodeToString(hasher.Sum(nil))
}

func SchemaSecret() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"raw_value": {
			Type:      schema.TypeString,
			Required:  true,
			Sensitive: true,
			StateFunc: HashStringMD5,
		},
		"value": {
			Type:     schema.TypeString,
			Computed: true,
		},
	}
}
