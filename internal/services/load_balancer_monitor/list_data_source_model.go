// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package load_balancer_monitor

import (
	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/load_balancers"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type LoadBalancerMonitorsResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[LoadBalancerMonitorsResultDataSourceModel] `json:"result,computed"`
}

type LoadBalancerMonitorsDataSourceModel struct {
	AccountID types.String                                                            `tfsdk:"account_id" path:"account_id"`
	MaxItems  types.Int64                                                             `tfsdk:"max_items"`
	Result    customfield.NestedObjectList[LoadBalancerMonitorsResultDataSourceModel] `tfsdk:"result"`
}

func (m *LoadBalancerMonitorsDataSourceModel) toListParams() (params load_balancers.MonitorListParams, diags diag.Diagnostics) {
	params = load_balancers.MonitorListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

type LoadBalancerMonitorsResultDataSourceModel struct {
	ID              types.String               `tfsdk:"id" json:"id,computed_optional"`
	AllowInsecure   types.Bool                 `tfsdk:"allow_insecure" json:"allow_insecure,computed"`
	ConsecutiveDown types.Int64                `tfsdk:"consecutive_down" json:"consecutive_down,computed"`
	ConsecutiveUp   types.Int64                `tfsdk:"consecutive_up" json:"consecutive_up,computed"`
	CreatedOn       timetypes.RFC3339          `tfsdk:"created_on" json:"created_on,computed"`
	Description     types.String               `tfsdk:"description" json:"description,computed_optional"`
	ExpectedBody    types.String               `tfsdk:"expected_body" json:"expected_body,computed_optional"`
	ExpectedCodes   types.String               `tfsdk:"expected_codes" json:"expected_codes,computed"`
	FollowRedirects types.Bool                 `tfsdk:"follow_redirects" json:"follow_redirects,computed"`
	Header          map[string]*[]types.String `tfsdk:"header" json:"header,computed_optional"`
	Interval        types.Int64                `tfsdk:"interval" json:"interval,computed"`
	Method          types.String               `tfsdk:"method" json:"method,computed"`
	ModifiedOn      timetypes.RFC3339          `tfsdk:"modified_on" json:"modified_on,computed"`
	Path            types.String               `tfsdk:"path" json:"path,computed"`
	Port            types.Int64                `tfsdk:"port" json:"port,computed"`
	ProbeZone       types.String               `tfsdk:"probe_zone" json:"probe_zone,computed_optional"`
	Retries         types.Int64                `tfsdk:"retries" json:"retries,computed"`
	Timeout         types.Int64                `tfsdk:"timeout" json:"timeout,computed"`
	Type            types.String               `tfsdk:"type" json:"type,computed"`
}
