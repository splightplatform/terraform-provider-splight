//go:generate terraform fmt -recursive examples/
//go:generate go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs

package main

import (
	"context"
	"flag"
	"log"

	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6/tf6server"
	"github.com/hashicorp/terraform-plugin-mux/tf5to6server"
	"github.com/splightplatform/terraform-provider-splight/provider"
)

// Version is initialized by the Go linker to contain the semver of this build.
var Version string = "dev"

func main() {
	ctx := context.Background()

	// Define and parse command-line flags.
	var debug bool
	flag.BoolVar(&debug, "debug", false, "Enable debugging support for tools like delve")
	flag.Parse()

	// Upgrade the provider to the TF6 protocol version.
	upgradedSdkProvider, err := tf5to6server.UpgradeServer(
		ctx,
		provider.Provider().GRPCProvider,
	)
	if err != nil {
		log.Fatalf("Failed to upgrade provider server: %v", err)
	}

	// Serve the provider with or without debugging based on the flag.
	if err := serveProvider(ctx, upgradedSdkProvider, debug); err != nil {
		log.Fatalf("Failed to serve provider: %v", err)
	}
}

// serveProvider sets up the provider server with optional debugging.
func serveProvider(ctx context.Context, providerServer tfprotov6.ProviderServer, debug bool) error {
	if debug {
		return tf6server.Serve(
			"registry.terraform.io/splightplatform/splight",
			func() tfprotov6.ProviderServer {
				return providerServer
			},
			tf6server.WithManagedDebug(),
		)
	}
	return tf6server.Serve(
		"registry.terraform.io/splightplatform/splight",
		func() tfprotov6.ProviderServer {
			return providerServer
		},
	)
}
