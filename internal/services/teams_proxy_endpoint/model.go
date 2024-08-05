// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package teams_proxy_endpoint

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type TeamsProxyEndpointResultEnvelope struct {
	Result TeamsProxyEndpointModel `json:"result,computed"`
}

type TeamsProxyEndpointModel struct {
	ID        types.String      `tfsdk:"id" json:"id,computed"`
	AccountID types.String      `tfsdk:"account_id" path:"account_id"`
	Name      types.String      `tfsdk:"name" json:"name"`
	IPs       *[]types.String   `tfsdk:"ips" json:"ips"`
	CreatedAt timetypes.RFC3339 `tfsdk:"created_at" json:"created_at,computed"`
	Subdomain types.String      `tfsdk:"subdomain" json:"subdomain,computed"`
	UpdatedAt timetypes.RFC3339 `tfsdk:"updated_at" json:"updated_at,computed"`
}