package main

import (
	"context"
	"flag"
	"log"

	"github.com/kzmake/terraform-provider-scalebase/scalebase"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
)

func main() {
	var debug bool

	flag.BoolVar(&debug, "debug", false, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()

	err := providerserver.Serve(context.Background(), scalebase.New, providerserver.ServeOpts{
		Address: "registry.terraform.io/kzmake/scalebase",
	})
	if err != nil {
		log.Fatal(err)
	}
}
