package scalebase

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var _ provider.Provider = &scalebaseProvider{}

func New() provider.Provider {
	return &scalebaseProvider{}
}

type scalebaseProvider struct{}

func (p *scalebaseProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "scalebase"
}

func (p *scalebaseProvider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{}
}

func (p *scalebaseProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
}

func (p *scalebaseProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return nil
}

func (p *scalebaseProvider) Resources(_ context.Context) []func() resource.Resource {
	return nil
}
