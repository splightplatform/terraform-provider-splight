package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/splightplatform/terraform-provider-splight/utils"
)

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
			StateFunc: func(val interface{}) string {
				return utils.HashStringMD5(val.(string))
			},
		},
		"value": {
			Type:     schema.TypeString,
			Computed: true,
		},
	}
}
