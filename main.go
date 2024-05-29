//go:generate terraform fmt -recursive examples/
//go:generate go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs

package main

import (
	"context"
	"flag"

	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6/tf6server"
	"github.com/hashicorp/terraform-plugin-mux/tf5to6server"
	"github.com/splightplatform/terraform-provider-splight/provider"
)

func main() {
	var debug bool
	ctx := context.Background()

	flag.BoolVar(&debug, "debug", false, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()

	upgradedSdkProvider, err := tf5to6server.UpgradeServer(
		ctx,
		provider.Provider().GRPCProvider,
	)

	if err != nil {
		// TODO: error
	}

	if debug {

		err = tf6server.Serve(
			"registry.terraform.io/splightplatform/splight",
			func() tfprotov6.ProviderServer {
				return upgradedSdkProvider
			},
			tf6server.WithManagedDebug(),
		)
	} else {
		err = tf6server.Serve(
			"registry.terraform.io/splightplatform/splight",
			func() tfprotov6.ProviderServer {
				return upgradedSdkProvider
			},
		)
	}
}
