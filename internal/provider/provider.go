package provider

import (
	"context"
	"terraform-provider-google-tag-manager/internal/api"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure the implementation satisfies the expected interfaces
var (
	_ provider.Provider = &gtmProvider{}
)

// New is a helper function to simplify provider server and testing implementation.
func New() provider.Provider {
	return &gtmProvider{}
}

// gtmProvider is the provider implementation.
type gtmProvider struct{}

// Metadata returns the provider type name.
func (p *gtmProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "gtm"
}

// Schema defines the provider-level schema for configuration data.
func (p *gtmProvider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"credential_file": schema.StringAttribute{
				Description: "Path to the credential file.",
				Required:    true},
			"account_id": schema.StringAttribute{
				Description: "GTM Account ID.",
				Required:    true},
			"container_id": schema.StringAttribute{
				Description: "GTM Container ID.",
				Required:    true},
			"workspace_name": schema.StringAttribute{
				Description: "Workspace name.",
				Required:    true},
			"retry_limit": schema.Int64Attribute{
				Description: "Number of times to retry requests when rate-limited before giving up.",
				Optional:    true},
		},
	}
}

type gtmProviderModel struct {
	CredentialFile types.String `tfsdk:"credential_file"`
	AccountId      types.String `tfsdk:"account_id"`
	ContainerId    types.String `tfsdk:"container_id"`
	WorkspaceName  types.String `tfsdk:"workspace_name"`
	RetryLimit     types.Int64  `tfsdk:"retry_limit"`
}

// Configure prepares an API client for data sources and resources.
func (p *gtmProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	tflog.Info(ctx, "Provider Configure starts.")
	defer tflog.Info(ctx, "Provider Configure finished.")

	// Retrieve provider data from configuration
	var config gtmProviderModel
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var retryLimit = 10
	if !config.RetryLimit.IsNull() && !config.RetryLimit.IsUnknown() {
		retryLimit = int(config.RetryLimit.ValueInt64())
	}

	client, err := api.NewClientInWorkspace(&api.ClientInWorkspaceOptions{
		ClientOptions: &api.ClientOptions{
			CredentialFile: config.CredentialFile.ValueString(),
			AccountId:      config.AccountId.ValueString(),
			ContainerId:    config.ContainerId.ValueString(),
			RetryLimit:     retryLimit,
		},
		WorkspaceName: config.WorkspaceName.ValueString(),
	})
	if err != nil {
		resp.Diagnostics.AddError("Unable to Create GTM Client", err.Error())
		return
	}
	resp.DataSourceData = client
	resp.ResourceData = client
}

// DataSources defines the data sources implemented in the provider.
func (p *gtmProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return nil
}

// Resources defines the resources implemented in the provider.
func (p *gtmProvider) Resources(_ context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewWorkspaceResource,
		NewTagResource,
		NewVariableResource,
		NewTriggerResource,
	}
}
