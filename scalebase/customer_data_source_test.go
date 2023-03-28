package scalebase

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/jarcoal/httpmock"

	"github.com/kzmake/terraform-provider-scalebase/api/models"
)

func TestCustomerDataSource(t *testing.T) {
	if isUnitTest {
		customer := &models.Publicv1Customer{
			ID:         "00000000-00000000-00000000-00000000",
			OptionalID: "test",
			Name:       "テスト",
		}

		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		json, _ := (&models.V1GetCustomerResponse{Customer: customer}).MarshalBinary()
		httpmock.RegisterResponder("POST", "https://localhost/v1/customer/get", httpmock.NewStringResponder(200, string(json)))
	}

	resource.Test(t, resource.TestCase{
		IsUnitTest:               isUnitTest,
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: providerConfig() + `data "scalebase_customer" "test" { optional_id = "test" }`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.scalebase_customer.test", "id"),
					resource.TestCheckResourceAttr("data.scalebase_customer.test", "optional_id", "test"),
					resource.TestCheckResourceAttr("data.scalebase_customer.test", "name", "テスト"),
				),
			},
		},
	})
}
