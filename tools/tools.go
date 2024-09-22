//go:build tools

package tools

import (
	// Documentation generation
	_ "github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs"
	// TODO: make this easy installable
	_ "mvdan.cc/gofumpt"
)
