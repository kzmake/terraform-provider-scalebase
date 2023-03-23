package main

import (
	"context"

	"github.com/kzmake/terraform-provider-scalebase/scalebase"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
)

func main() {
	_ = providerserver.Serve(context.Background(), scalebase.New, providerserver.ServeOpts{
		Address: "kzmake/scalebase",
	})
}
