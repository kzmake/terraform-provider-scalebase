package scalebase

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/jarcoal/httpmock"

	"github.com/kzmake/terraform-provider-scalebase/api/models"
)

func TestCatalogItemDataSource(t *testing.T) {
	if true {
		catalogItem := &models.V1CatalogItem{
			ID:         "00000000-00000000-00000000-00000000",
			OptionalID: "test",
			Name:       "テストカタログアイテム",
			ProductID:  "00000000-00000000-00000000-00000000",
			Status:     models.V1CatalogItemStatusSTATUSONSALE,
		}

		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		json, _ := (&models.V1GetCatalogItemResponse{CatalogItem: catalogItem}).MarshalBinary()
		httpmock.RegisterResponder("POST", "https://localhost/v1/catalogitem/get", httpmock.NewStringResponder(200, string(json)))
	}

	resource.Test(t, resource.TestCase{
		IsUnitTest:               true,
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: defaultConfigForUnitTest + `data "scalebase_catalog_item" "test" { id = "00000000-00000000-00000000-00000000" }`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.scalebase_catalog_item.test", "id"),
					resource.TestCheckResourceAttr("data.scalebase_catalog_item.test", "optional_id", "test"),
					resource.TestCheckResourceAttr("data.scalebase_catalog_item.test", "name", "テストカタログアイテム"),
					resource.TestCheckResourceAttrSet("data.scalebase_catalog_item.test", "product_id"),
					resource.TestCheckResourceAttr("data.scalebase_catalog_item.test", "status", "STATUS_ON_SALE"),
				),
			},
		},
	})
}
