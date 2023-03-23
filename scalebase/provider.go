package scalebase

import (
	"context"
	"os"
	"strings"

	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
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

type scalebaseProvider struct {
	Host  types.String `tfsdk:"host"`
	Token types.String `tfsdk:"token"`
}

func (p *scalebaseProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "scalebase"
}

func (p *scalebaseProvider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
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

func (p *scalebaseProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var config scalebaseProvider
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if config.Host.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("host"),
			"Unknown Scalebase API Host",
			strings.Join([]string{
				"The provider cannot create the Scalebase API client as there is an unknown configuration value for the Scalebase API host.",
				"Either target apply the source of the value first, set the value statically in the configuration, or use the SCALEBASE_HOST environment variable.",
			}, " "),
		)
	}

	if config.Token.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("password"),
			"Unknown Scalebase API Token",
			strings.Join([]string{
				"The provider cannot create the Scalebase API client as there is an unknown configuration value for the Scalebase API token.",
				"Either target apply the source of the value first, set the value statically in the configuration, or use the SCALEBASE_TOKEN environment variable.",
			}, " "),
		)
	}

	if resp.Diagnostics.HasError() {
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
		resp.Diagnostics.AddAttributeError(
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
		resp.Diagnostics.AddAttributeError(
			path.Root("token"),
			"Missing Scalebase API Token",
			strings.Join([]string{
				"The provider cannot create the Scalebase API client as there is a missing or empty value for the Scalebase API token.",
				"Set the token value in the configuration or use the SCALEBASE_TOKEN environment variable.",
				"If either is already set, ensure the value is not empty.",
			}, " "),
		)
	}

	if resp.Diagnostics.HasError() {
		return
	}

	r := httptransport.New(host, apiclient.DefaultBasePath, apiclient.DefaultSchemes)
	r.DefaultAuthentication = httptransport.BearerToken(token)
	client := apiclient.New(r, strfmt.Default)

	resp.DataSourceData = client
	resp.ResourceData = client
}

func (p *scalebaseProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return nil
}

func (p *scalebaseProvider) Resources(_ context.Context) []func() resource.Resource {
	return nil
}
