// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package tunnel

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/cloudflare/cloudflare-go/v2/zero_trust"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/logging"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

type TunnelDataSource struct {
	client *cloudflare.Client
}

var _ datasource.DataSourceWithConfigure = &TunnelDataSource{}

func NewTunnelDataSource() datasource.DataSource {
	return &TunnelDataSource{}
}

func (d *TunnelDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_tunnel"
}

func (r *TunnelDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *TunnelDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *TunnelDataSourceModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	if data.FindOneBy == nil {
		res := new(http.Response)
		env := TunnelResultDataSourceEnvelope{*data}
		_, err := r.client.ZeroTrust.Tunnels.Get(
			ctx,
			data.TunnelID.ValueString(),
			zero_trust.TunnelGetParams{
				AccountID: cloudflare.F(data.AccountID.ValueString()),
			},
			option.WithResponseBodyInto(&res),
			option.WithMiddleware(logging.Middleware(ctx)),
		)
		if err != nil {
			resp.Diagnostics.AddError("failed to make http request", err.Error())
			return
		}
		bytes, _ := io.ReadAll(res.Body)
		err = apijson.Unmarshal(bytes, &env)
		if err != nil {
			resp.Diagnostics.AddError("failed to deserialize http request", err.Error())
			return
		}
		data = &env.Result
	} else {
		dataFindOneByExistedAt, errs := data.FindOneBy.ExistedAt.ValueRFC3339Time()
		resp.Diagnostics.Append(errs...)
		dataFindOneByWasActiveAt, errs := data.FindOneBy.WasActiveAt.ValueRFC3339Time()
		resp.Diagnostics.Append(errs...)
		dataFindOneByWasInactiveAt, errs := data.FindOneBy.WasInactiveAt.ValueRFC3339Time()
		resp.Diagnostics.Append(errs...)
		if resp.Diagnostics.HasError() {
			return
		}

		items := &[]*TunnelDataSourceModel{}
		env := TunnelResultListDataSourceEnvelope{items}

		page, err := r.client.ZeroTrust.Tunnels.List(ctx, zero_trust.TunnelListParams{
			AccountID:     cloudflare.F(data.FindOneBy.AccountID.ValueString()),
			ExcludePrefix: cloudflare.F(data.FindOneBy.ExcludePrefix.ValueString()),
			ExistedAt:     cloudflare.F(dataFindOneByExistedAt),
			IncludePrefix: cloudflare.F(data.FindOneBy.IncludePrefix.ValueString()),
			IsDeleted:     cloudflare.F(data.FindOneBy.IsDeleted.ValueBool()),
			Name:          cloudflare.F(data.FindOneBy.Name.ValueString()),
			Page:          cloudflare.F(data.FindOneBy.Page.ValueFloat64()),
			PerPage:       cloudflare.F(data.FindOneBy.PerPage.ValueFloat64()),
			Status:        cloudflare.F(zero_trust.TunnelListParamsStatus(data.FindOneBy.Status.ValueString())),
			TunTypes:      cloudflare.F(data.FindOneBy.TunTypes.ValueString()),
			UUID:          cloudflare.F(data.FindOneBy.UUID.ValueString()),
			WasActiveAt:   cloudflare.F(dataFindOneByWasActiveAt),
			WasInactiveAt: cloudflare.F(dataFindOneByWasInactiveAt),
		})
		if err != nil {
			resp.Diagnostics.AddError("failed to make http request", err.Error())
			return
		}

		bytes := []byte(page.JSON.RawJSON())
		err = apijson.Unmarshal(bytes, &env)
		if err != nil {
			resp.Diagnostics.AddError("failed to unmarshal http request", err.Error())
			return
		}

		if count := len(*items); count != 1 {
			resp.Diagnostics.AddError("failed to find exactly one result", fmt.Sprint(count)+" found")
			return
		}
		data = (*items)[0]
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}