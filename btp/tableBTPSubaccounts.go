package btp

import (
	"context"
	"encoding/json"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
)

const (
	subAccountsPath = "/accounts/v1/subaccounts"
)

func tableBTPSubaccounts() *plugin.Table {
	return &plugin.Table{
		Name:        "btp_accounts_subaccounts",
		Description: "BTP Subaccounts",
		List: &plugin.ListConfig{
			Hydrate: listSubaccounts,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("guid"),
			Hydrate:    getSubaccount,
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

	// conn, err := connect(ctx, d)
	btpClient, err := NewBTPClient(nil, d.Connection)
	if err != nil {
		return nil, err
	}

	// Call the API
	body, err := btpClient.Get(ctx, AccountsService, subAccountsPath, nil, nil)

	if err != nil {
		return nil, err
	}

	var data struct {
		Subaccounts []Subaccount `json:"value"`
	}

	// Convert JSON response to Subaccounts structure
	err = json.Unmarshal(body, &data)

	logger.Debug("listAccount", "data", data)
	logger.Debug("listAccount", "err", err)

	if err != nil {
		return nil, err
	}

	for _, s := range data.Subaccounts {
		d.StreamListItem(ctx, s)
	}

	return nil, nil
}

func getSubaccount(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {

	fnName := "getSubaccount"

	logger := plugin.Logger(ctx)
	logger.Trace("Hydrating get subaccount")

	btpClient, err := NewBTPClient(nil, d.Connection)
	if err != nil {
		return nil, err
	}

	quals := d.KeyColumnQuals
	logger.Warn(fnName, "quals", quals)
	subaccountGuid := quals["guid"].GetStringValue()
	logger.Warn(fnName, "guid", subaccountGuid)

	path := subAccountsPath + "/" + subaccountGuid

	logger.Warn(fnName, "path", path)

	// Call the API
	body, err := btpClient.Get(ctx, AccountsService, globalAccountsPath, nil, nil)

	if err != nil {
		return nil, err
	}

	var data Directory

	// Convert JSON response to Directory structure
	err = json.Unmarshal(body, &data)

	logger.Warn(fnName, "data", data)
	logger.Warn(fnName, "err", err)

	if err != nil {
		return nil, err
	}

	return data, nil
}
