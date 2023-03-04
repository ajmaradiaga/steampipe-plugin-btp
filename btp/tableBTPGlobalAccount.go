package btp

import (
	"context"
	"encoding/json"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
)

const (
	globalAccountsPath = "/accounts/v1/globalAccount"
)

func tableBTPGlobalAccount() *plugin.Table {
	return &plugin.Table{
		Name:        "btp_accounts_global_account",
		Description: "BTP Global Account details",
		List: &plugin.ListConfig{
			Hydrate: listAccount,
		},
		Columns: []*plugin.Column{
			{Name: "guid", Type: proto.ColumnType_STRING, Description: "The account GUID"},
			{Name: "display_name", Type: proto.ColumnType_STRING, Description: "Account display name"},
			{Name: "created_date", Type: proto.ColumnType_INT, Description: "Date account was created"},
			{Name: "modified_date", Type: proto.ColumnType_INT, Description: "Date account was last modified"},
			{Name: "entity_state", Type: proto.ColumnType_STRING, Description: "Entity state"},
			{Name: "state_message", Type: proto.ColumnType_STRING},
			{Name: "subdomain", Type: proto.ColumnType_STRING, Description: "Account subdomain"},
			{Name: "contract_status", Type: proto.ColumnType_STRING},
			{Name: "commercial_model", Type: proto.ColumnType_STRING},
			{Name: "consumption_based", Type: proto.ColumnType_BOOL, Description: "True if the account is consumption based"},
			{Name: "license_type", Type: proto.ColumnType_STRING},
			{Name: "geo_access", Type: proto.ColumnType_STRING},
			{Name: "renewal_date", Type: proto.ColumnType_INT},
		},
	}
}

func listAccount(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {

	logger := plugin.Logger(ctx)
	logger.Trace("Hydrating Global Account")

	btpClient, err := NewBTPClient(nil, d.Connection)

	if err != nil {
		return nil, err
	}

	queryStrings := map[string]string{
		"derivedAuthorizations": "any",
	}

	if err != nil {
		return nil, err
	}

	// Call the API
	body, err := btpClient.Get(ctx, AccountsService, globalAccountsPath, nil, queryStrings)

	if err != nil {
		return nil, err
	}

	var data GlobalAccount

	// Convert JSON response to structure
	err = json.Unmarshal(body, &data)

	logger.Debug("listAccount", "data", data)
	logger.Debug("listAccount", "err", err)

	if err != nil {
		return nil, err
	}

	d.StreamListItem(ctx, data)

	return nil, nil
}
