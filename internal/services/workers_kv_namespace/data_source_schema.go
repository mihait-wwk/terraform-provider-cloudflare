// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_kv_namespace

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = &WorkersKVNamespaceDataSource{}
var _ datasource.DataSourceWithValidateConfig = &WorkersKVNamespaceDataSource{}

func (r WorkersKVNamespaceDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description: "Identifier",
				Optional:    true,
			},
			"namespace_id": schema.StringAttribute{
				Description: "Namespace identifier tag.",
				Optional:    true,
			},
			"id": schema.StringAttribute{
				Description: "Namespace identifier tag.",
				Computed:    true,
			},
			"title": schema.StringAttribute{
				Description: "A human-readable string name for a Namespace.",
				Computed:    true,
			},
			"supports_url_encoding": schema.BoolAttribute{
				Description: "True if keys written on the URL will be URL-decoded before storing. For example, if set to \"true\", a key written on the URL as \"%3F\" will be stored as \"?\".",
				Computed:    true,
			},
			"find_one_by": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"account_id": schema.StringAttribute{
						Description: "Identifier",
						Required:    true,
					},
					"direction": schema.StringAttribute{
						Description: "Direction to order namespaces.",
						Optional:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("asc", "desc"),
						},
					},
					"order": schema.StringAttribute{
						Description: "Field to order results by.",
						Optional:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("id", "title"),
						},
					},
					"page": schema.Float64Attribute{
						Description: "Page number of paginated results.",
						Computed:    true,
						Optional:    true,
					},
					"per_page": schema.Float64Attribute{
						Description: "Maximum number of results per page.",
						Computed:    true,
						Optional:    true,
					},
				},
			},
		},
	}
}

func (r *WorkersKVNamespaceDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}

func (r *WorkersKVNamespaceDataSource) ValidateConfig(ctx context.Context, req datasource.ValidateConfigRequest, resp *datasource.ValidateConfigResponse) {
}