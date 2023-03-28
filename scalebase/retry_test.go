package scalebase

import (
	"net/http"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/jarcoal/httpmock"
	"github.com/kzmake/terraform-provider-scalebase/api/models"
)

func TestRetryWhen429(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	n := 0
	httpmock.RegisterResponder("POST", "https://localhost/v1/customer/get", func(req *http.Request) (*http.Response, error) {
		n = n + 1
		// ~ 3
		if n <= 3 {
			return httpmock.NewStringResponse(429, `{"error": "too many requests"}`), nil
		}
		// 3 ~
		customer := &models.Publicv1Customer{
			ID:         "00000000-00000000-00000000-00000000",
			OptionalID: "test",
			Name:       "テスト",
		}
		json, _ := (&models.V1GetCustomerResponse{Customer: customer}).MarshalBinary()
		return httpmock.NewStringResponse(200, string(json)), nil
	})

	resource.Test(t, resource.TestCase{
		IsUnitTest:               true,
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: defaultConfigForUnitTest + `data "scalebase_customer" "test" { optional_id = "test" }`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.scalebase_customer.test", "id"),
					resource.TestCheckResourceAttr("data.scalebase_customer.test", "optional_id", "test"),
					resource.TestCheckResourceAttr("data.scalebase_customer.test", "name", "テスト"),
				),
			},
		},
	})
}

func TestRetryWhen5XX(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	n := 0
	httpmock.RegisterResponder("POST", "https://localhost/v1/customer/get", func(req *http.Request) (*http.Response, error) {
		n = n + 1
		// ~ 3
		if n <= 3 {
			return httpmock.NewStringResponse(500, `{"error": "internal server error"}`), nil
		}
		// 3 ~
		customer := &models.Publicv1Customer{
			ID:         "00000000-00000000-00000000-00000000",
			OptionalID: "test",
			Name:       "テスト",
		}
		json, _ := (&models.V1GetCustomerResponse{Customer: customer}).MarshalBinary()
		return httpmock.NewStringResponse(200, string(json)), nil
	})

	resource.Test(t, resource.TestCase{
		IsUnitTest:               true,
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: defaultConfigForUnitTest + `data "scalebase_customer" "test" { optional_id = "test" }`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.scalebase_customer.test", "id"),
					resource.TestCheckResourceAttr("data.scalebase_customer.test", "optional_id", "test"),
					resource.TestCheckResourceAttr("data.scalebase_customer.test", "name", "テスト"),
				),
			},
		},
	})
}
