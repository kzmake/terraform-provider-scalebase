package scalebase

import (
	"os"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
)

const (
	defaultConfigForUnitTest = `
provider "scalebase" {
  host = "localhost"
  token = "test_token"
}
`

	defaultConfigForAcc = `
provider "scalebase" {
}
`
)

var (
	testAccProtoV6ProviderFactories = map[string]func() (tfprotov6.ProviderServer, error){"scalebase": providerserver.NewProtocol6WithError(New())}

	isUnitTest = os.Getenv("SCALEBASE_HOST") == "" || os.Getenv("SCALEBASE_HOST") == "localhost"
)

func providerConfig() string {
	if isUnitTest {
		return defaultConfigForUnitTest
	}

	return defaultConfigForAcc
}
