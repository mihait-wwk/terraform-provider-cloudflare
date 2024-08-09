// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package magic_wan_static_route

import (
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type MagicWANStaticRouteResultDataSourceEnvelope struct {
	Result MagicWANStaticRouteDataSourceModel `json:"result,computed"`
}

type MagicWANStaticRouteDataSourceModel struct {
	AccountID types.String         `tfsdk:"account_id" path:"account_id"`
	RouteID   types.String         `tfsdk:"route_id" path:"route_id"`
	Route     jsontypes.Normalized `tfsdk:"route" json:"route"`
}