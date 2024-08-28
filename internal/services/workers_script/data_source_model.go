// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_script

import (
	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/workers"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type WorkersScriptResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[WorkersScriptDataSourceModel] `json:"result,computed"`
}

type WorkersScriptDataSourceModel struct {
	AccountID     types.String                                  `tfsdk:"account_id" path:"account_id"`
	ScriptName    types.String                                  `tfsdk:"script_name" path:"script_name"`
	CreatedOn     timetypes.RFC3339                             `tfsdk:"created_on" json:"created_on" format:"date-time"`
	Etag          types.String                                  `tfsdk:"etag" json:"etag"`
	ID            types.String                                  `tfsdk:"id" json:"id"`
	Logpush       types.Bool                                    `tfsdk:"logpush" json:"logpush"`
	ModifiedOn    timetypes.RFC3339                             `tfsdk:"modified_on" json:"modified_on" format:"date-time"`
	PlacementMode types.String                                  `tfsdk:"placement_mode" json:"placement_mode"`
	UsageModel    types.String                                  `tfsdk:"usage_model" json:"usage_model"`
	TailConsumers *[]*WorkersScriptTailConsumersDataSourceModel `tfsdk:"tail_consumers" json:"tail_consumers"`
	Filter        *WorkersScriptFindOneByDataSourceModel        `tfsdk:"filter"`
}

func (m *WorkersScriptDataSourceModel) toReadParams() (params workers.ScriptGetParams, diags diag.Diagnostics) {
	params = workers.ScriptGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

func (m *WorkersScriptDataSourceModel) toListParams() (params workers.ScriptListParams, diags diag.Diagnostics) {
	params = workers.ScriptListParams{
		AccountID: cloudflare.F(m.Filter.AccountID.ValueString()),
	}

	return
}

type WorkersScriptTailConsumersDataSourceModel struct {
	Service     types.String `tfsdk:"service" json:"service,computed"`
	Environment types.String `tfsdk:"environment" json:"environment,computed_optional"`
	Namespace   types.String `tfsdk:"namespace" json:"namespace,computed_optional"`
}

type WorkersScriptFindOneByDataSourceModel struct {
	AccountID types.String `tfsdk:"account_id" path:"account_id"`
}