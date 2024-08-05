// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package web_analytics_site

import (
	"context"
	"fmt"

	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/rum"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

type WebAnalyticsSitesDataSource struct {
	client *cloudflare.Client
}

var _ datasource.DataSourceWithConfigure = &WebAnalyticsSitesDataSource{}

func NewWebAnalyticsSitesDataSource() datasource.DataSource {
	return &WebAnalyticsSitesDataSource{}
}

func (d *WebAnalyticsSitesDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_web_analytics_sites"
}

func (r *WebAnalyticsSitesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*cloudflare.Client)

	if !ok {
		resp.Diagnostics.AddError(
			"unexpected resource configure type",
			fmt.Sprintf("Expected *cloudflare.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *WebAnalyticsSitesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *WebAnalyticsSitesDataSourceModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	items := &[]*WebAnalyticsSitesResultDataSourceModel{}
	env := WebAnalyticsSitesResultListDataSourceEnvelope{items}
	maxItems := int(data.MaxItems.ValueInt64())
	acc := []*WebAnalyticsSitesResultDataSourceModel{}

	page, err := r.client.RUM.SiteInfo.List(ctx, rum.SiteInfoListParams{
		AccountID: cloudflare.F(data.AccountID.ValueString()),
		OrderBy:   cloudflare.F(rum.SiteInfoListParamsOrderBy(data.OrderBy.ValueString())),
		Page:      cloudflare.F(data.Page.ValueFloat64()),
		PerPage:   cloudflare.F(data.PerPage.ValueFloat64()),
	})
	if err != nil {
		resp.Diagnostics.AddError("failed to make http request", err.Error())
		return
	}

	for page != nil && len(page.Result) > 0 {
		bytes := []byte(page.JSON.RawJSON())
		err = apijson.Unmarshal(bytes, &env)
		if err != nil {
			resp.Diagnostics.AddError("failed to unmarshal http request", err.Error())
			return
		}
		acc = append(acc, *items...)
		if len(acc) >= maxItems {
			break
		}
		page, err = page.GetNextPage()
		if err != nil {
			resp.Diagnostics.AddError("failed to fetch next page", err.Error())
			return
		}
	}

	acc = acc[:maxItems]
	data.Result = &acc

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}