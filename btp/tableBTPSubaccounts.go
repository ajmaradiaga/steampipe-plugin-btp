package btp

import (
	"context"
	"encoding/json"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
)

func tableBTPSubaccounts() *plugin.Table {
	return &plugin.Table{
		Name:        "btp_subaccounts",
		Description: "BTP Subaccounts",
		List: &plugin.ListConfig{
			Hydrate: listSubaccounts,
		},
		Columns: []*plugin.Column{
			{Name: "guid", Type: proto.ColumnType_STRING, Description: "The account GUID"},
			{Name: "technical_name", Type: proto.ColumnType_STRING, Description: "Subaccount technical name"},
			{Name: "display_name", Type: proto.ColumnType_STRING, Description: "Subaccount display name"},
			{Name: "global_account_guid", Type: proto.ColumnType_STRING},
			{Name: "parent_guid", Type: proto.ColumnType_STRING},
			{Name: "parent_type", Type: proto.ColumnType_STRING},
			{Name: "region", Type: proto.ColumnType_STRING},
			{Name: "subdomain", Type: proto.ColumnType_STRING, Description: "Account subdomain"},
			{Name: "beta_enabled", Type: proto.ColumnType_BOOL, Description: "True if the account is beta enabled"},
			{Name: "used_for_production", Type: proto.ColumnType_STRING},
			{Name: "description", Type: proto.ColumnType_STRING},
			{Name: "state", Type: proto.ColumnType_STRING},
			{Name: "state_message", Type: proto.ColumnType_STRING},
			{Name: "created_date", Type: proto.ColumnType_INT, Description: "Date account was last modified"},
			{Name: "created_by", Type: proto.ColumnType_STRING},
			{Name: "modified_date", Type: proto.ColumnType_INT, Description: "Date account was last modified"},
		},
	}
}

func listSubaccounts(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {

	logger := plugin.Logger(ctx)
	logger.Trace("Hydrating list subaccounts")

	conn, err := connect(ctx, d)
	if err != nil {
		return nil, err
	}
	quals := d.KeyColumnQuals
	plugin.Logger(ctx).Warn("listAccount", "quals", quals)
	id := quals["id"].GetInt64Value()
	plugin.Logger(ctx).Warn("listAccount", "id", id)

	btpConf := GetConfig(d.Connection)

	url := *btpConf.EndpointsAccountServiceUrl + "/accounts/v1/subaccounts?derivedAuthorizations=any"
	err = conn.SetEndpointURL(url)

	body, err := conn.Get(ctx, "")

	var data struct {
		Subaccounts []Subaccount `json:"value"`
	}

	err = json.Unmarshal(body, &data)

	plugin.Logger(ctx).Warn("listAccount", "url", url)
	plugin.Logger(ctx).Warn("listAccount", "body", string(body[:]))
	plugin.Logger(ctx).Warn("listAccount", "btpConf", *btpConf.EndpointsAccountServiceUrl)

	plugin.Logger(ctx).Warn("listAccount", "data", data)
	plugin.Logger(ctx).Warn("listAccount", "err", err)

	if err != nil {
		return nil, err
	}

	for _, s := range data.Subaccounts {
		d.StreamListItem(ctx, s)
	}

	return nil, nil
}
