package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

type sodiumProvider struct {
}

// Enforce interfaces we want provider to implement.
var _ provider.Provider = (*sodiumProvider)(nil)

func New() provider.Provider {
	return &sodiumProvider{}
}

func (p *sodiumProvider) resetConfig() {

}

func (p *sodiumProvider) Metadata(_ context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "sodium"
}

func (p *sodiumProvider) Schema(_ context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Provider configuration",
	}
}

func (p *sodiumProvider) Configure(ctx context.Context, req provider.ConfigureRequest, res *provider.ConfigureResponse) {
	tflog.Debug(ctx, "Configuring provider")
	p.resetConfig()

	// Since the provider instance is being passed, ensure these response
	// values are always set before early returns, etc.
	res.DataSourceData = p
	res.ResourceData = p
}

func (p *sodiumProvider) Resources(_ context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewKeyPairResource,
	}
}

func (p *sodiumProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
	}
}
