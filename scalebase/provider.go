package scalebase

import (
	"context"
	"net/http"
	"os"
	"strings"
	"time"

	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	retryablehttp "github.com/hashicorp/go-retryablehttp"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"

	apiclient "github.com/kzmake/terraform-provider-scalebase/api/client"
)

var _ provider.Provider = &scalebaseProvider{}

func New() provider.Provider {
	return &scalebaseProvider{}
}

type scalebaseProvider struct{}

type scalebaseProviderModel struct {
	Host  types.String `tfsdk:"host"`
	Token types.String `tfsdk:"token"`
}

func (p *scalebaseProvider) Metadata(_ context.Context, _ provider.MetadataRequest, res *provider.MetadataResponse) {
	res.TypeName = "scalebase"
}

func (p *scalebaseProvider) Schema(_ context.Context, _ provider.SchemaRequest, res *provider.SchemaResponse) {
	res.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"host": schema.StringAttribute{
				Optional: true,
			},
			"token": schema.StringAttribute{
				Optional:  true,
				Sensitive: true,
			},
		},
	}
}

func (p *scalebaseProvider) Configure(ctx context.Context, req provider.ConfigureRequest, res *provider.ConfigureResponse) {
	var config scalebaseProviderModel
	diags := req.Config.Get(ctx, &config)
	res.Diagnostics.Append(diags...)
	if res.Diagnostics.HasError() {
		return
	}

	if config.Host.IsUnknown() {
		res.Diagnostics.AddAttributeError(
			path.Root("host"),
			"Unknown Scalebase API Host",
			strings.Join([]string{
				"The provider cannot create the Scalebase API client as there is an unknown configuration value for the Scalebase API host.",
				"Either target apply the source of the value first, set the value statically in the configuration, or use the SCALEBASE_HOST environment variable.",
			}, " "),
		)
	}

	if config.Token.IsUnknown() {
		res.Diagnostics.AddAttributeError(
			path.Root("password"),
			"Unknown Scalebase API Token",
			strings.Join([]string{
				"The provider cannot create the Scalebase API client as there is an unknown configuration value for the Scalebase API token.",
				"Either target apply the source of the value first, set the value statically in the configuration, or use the SCALEBASE_TOKEN environment variable.",
			}, " "),
		)
	}

	if res.Diagnostics.HasError() {
		return
	}

	host := os.Getenv("SCALEBASE_HOST")
	token := os.Getenv("SCALEBASE_TOKEN")

	if !config.Host.IsNull() {
		host = config.Host.ValueString()
	}

	if !config.Token.IsNull() {
		token = config.Token.ValueString()
	}

	if host == "" {
		res.Diagnostics.AddAttributeError(
			path.Root("host"),
			"Missing Scalebase API Host",
			strings.Join([]string{
				"The provider cannot create the Scalebase API client as there is a missing or empty value for the Scalebase API host.",
				"Set the host value in the configuration or use the SCALEBASE_HOST environment variable.",
				"If either is already set, ensure the value is not empty.",
			}, " "),
		)
	}

	if token == "" {
		res.Diagnostics.AddAttributeError(
			path.Root("token"),
			"Missing Scalebase API Token",
			strings.Join([]string{
				"The provider cannot create the Scalebase API client as there is a missing or empty value for the Scalebase API token.",
				"Set the token value in the configuration or use the SCALEBASE_TOKEN environment variable.",
				"If either is already set, ensure the value is not empty.",
			}, " "),
		)
	}

	if res.Diagnostics.HasError() {
		return
	}

	retry := retryablehttp.NewClient()
	retry.HTTPClient = http.DefaultClient // NOTE: for using httpmock
	retry.RetryWaitMin = 0500 * time.Millisecond
	retry.RetryWaitMax = 2000 * time.Second
	retry.RetryMax = 5

	r := httptransport.NewWithClient(host, apiclient.DefaultBasePath, apiclient.DefaultSchemes, retry.StandardClient())
	r.DefaultAuthentication = httptransport.BearerToken(token)
	client := apiclient.New(r, strfmt.Default)

	res.DataSourceData = client
	res.ResourceData = client
}

func (p *scalebaseProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewProductDataSource,
		NewCatalogItemDataSource,
		NewCustomerDataSource,
		NewCustomFieldMasterDataSource,
	}
}

func (p *scalebaseProvider) Resources(_ context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewCustomerResource,
		NewCustomerStaffResource,
		NewResourceResource,
	}
}
