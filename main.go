package main

import (
	"context"
	"flag"
	"log"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/splightplatform/terraform-provider-splight/internal/provider"
)

var version string = "dev"

func main() {
	var debug bool

	flag.BoolVar(&debug, "debug", false, "Set to true to run the provider with support for debuggers")
	flag.Parse()

	opts := providerserver.ServeOpts{
		// Also update the tfplugindocs generate command to either remove the
		// -provider-name flag or set its value to the updated provider name.
		Address: "registry.terraform.io/hashicorp/splight",
		Debug:   debug,
	}

	err := providerserver.Serve(context.Background(), provider.NewSplightProvider(version), opts)

	if err != nil {
		log.Fatal(err.Error())
	}
}
