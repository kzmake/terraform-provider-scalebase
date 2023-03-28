package scalebase

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/jarcoal/httpmock"

	"github.com/kzmake/terraform-provider-scalebase/api/models"
)

func TestProductDataSource(t *testing.T) {
	if isUnitTest {
		product := &models.V1Product{
			ID:          "00000000-00000000-00000000-00000000",
			OptionalID:  "test",
			Name:        "テスト",
			Description: "ああああああ",
		}

		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		json, _ := (&models.V1GetProductResponse{Product: product}).MarshalBinary()
		httpmock.RegisterResponder("POST", "https://localhost/v1/product/get", httpmock.NewStringResponder(200, string(json)))
	}

	resource.Test(t, resource.TestCase{
		IsUnitTest:               isUnitTest,
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: providerConfig() + `data "scalebase_product" "test" { optional_id = "test" }`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.scalebase_product.test", "id"),
					resource.TestCheckResourceAttr("data.scalebase_product.test", "optional_id", "test"),
					resource.TestCheckResourceAttr("data.scalebase_product.test", "name", "テスト"),
					resource.TestCheckResourceAttr("data.scalebase_product.test", "description", "ああああああ"),
				),
			},
		},
	})
}
