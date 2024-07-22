// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package web_analytics_site

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSourceWithConfigValidators = &WebAnalyticsSitesDataSource{}
var _ datasource.DataSourceWithValidateConfig = &WebAnalyticsSitesDataSource{}

func (r WebAnalyticsSitesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description: "Identifier",
				Required:    true,
			},
			"order_by": schema.StringAttribute{
				Description: "The property used to sort the list of results.",
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("host", "created"),
				},
			},
			"page": schema.Float64Attribute{
				Description: "Current page within the paginated list of results.",
				Optional:    true,
			},
			"per_page": schema.Float64Attribute{
				Description: "Number of items to return per page of results.",
				Optional:    true,
			},
			"max_items": schema.Int64Attribute{
				Description: "Max items to fetch, default: 1000",
				Optional:    true,
			},
			"items": schema.ListNestedAttribute{
				Description: "The items returned by the data source",
				Computed:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"auto_install": schema.BoolAttribute{
							Description: "If enabled, the JavaScript snippet is automatically injected for orange-clouded sites.",
							Computed:    true,
							Optional:    true,
						},
						"created": schema.StringAttribute{
							Computed: true,
						},
						"rules": schema.ListNestedAttribute{
							Description: "A list of rules.",
							Computed:    true,
							Optional:    true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										Description: "The Web Analytics rule identifier.",
										Computed:    true,
										Optional:    true,
									},
									"created": schema.StringAttribute{
										Computed: true,
									},
									"host": schema.StringAttribute{
										Description: "The hostname the rule will be applied to.",
										Computed:    true,
										Optional:    true,
									},
									"inclusive": schema.BoolAttribute{
										Description: "Whether the rule includes or excludes traffic from being measured.",
										Computed:    true,
										Optional:    true,
									},
									"is_paused": schema.BoolAttribute{
										Description: "Whether the rule is paused or not.",
										Computed:    true,
										Optional:    true,
									},
									"paths": schema.ListAttribute{
										Description: "The paths the rule will be applied to.",
										Computed:    true,
										Optional:    true,
										ElementType: types.StringType,
									},
									"priority": schema.Float64Attribute{
										Computed: true,
										Optional: true,
									},
								},
							},
						},
						"ruleset": schema.SingleNestedAttribute{
							Computed: true,
							Optional: true,
							Attributes: map[string]schema.Attribute{
								"id": schema.StringAttribute{
									Description: "The Web Analytics ruleset identifier.",
									Computed:    true,
									Optional:    true,
								},
								"enabled": schema.BoolAttribute{
									Description: "Whether the ruleset is enabled.",
									Computed:    true,
									Optional:    true,
								},
								"zone_name": schema.StringAttribute{
									Computed: true,
									Optional: true,
								},
								"zone_tag": schema.StringAttribute{
									Description: "The zone identifier.",
									Computed:    true,
									Optional:    true,
								},
							},
						},
						"site_tag": schema.StringAttribute{
							Description: "The Web Analytics site identifier.",
							Computed:    true,
							Optional:    true,
						},
						"site_token": schema.StringAttribute{
							Description: "The Web Analytics site token.",
							Computed:    true,
							Optional:    true,
						},
						"snippet": schema.StringAttribute{
							Description: "Encoded JavaScript snippet.",
							Computed:    true,
							Optional:    true,
						},
					},
				},
			},
		},
	}
}

func (r *WebAnalyticsSitesDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}

func (r *WebAnalyticsSitesDataSource) ValidateConfig(ctx context.Context, req datasource.ValidateConfigRequest, resp *datasource.ValidateConfigResponse) {
}