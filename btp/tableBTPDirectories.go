package btp

import (
	"context"
	"encoding/json"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
)

const (
	directoriesPath = "/accounts/v1/directories"
)

func tableBTPDirectories() *plugin.Table {
	return &plugin.Table{
		Name:        "btp_accounts_directories",
		Description: "BTP Directories",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("guid"),
			Hydrate:    getDirectory,
		},
		Columns: []*plugin.Column{
			{Name: "guid", Type: proto.ColumnType_STRING, Description: "The directory GUID"},
			{Name: "parent_type", Type: proto.ColumnType_STRING},
			{Name: "global_account_guid", Type: proto.ColumnType_STRING},
			{Name: "display_name", Type: proto.ColumnType_STRING, Description: "Directory display name"},
			{Name: "created_date", Type: proto.ColumnType_INT, Description: "Date account was last modified"},
			{Name: "created_by", Type: proto.ColumnType_STRING},
			{Name: "modified_date", Type: proto.ColumnType_INT, Description: "Date account was last modified"},
			{Name: "entity_state", Type: proto.ColumnType_STRING},
			{Name: "state_message", Type: proto.ColumnType_STRING},
			{Name: "directory_type", Type: proto.ColumnType_STRING},
			{Name: "contract_status", Type: proto.ColumnType_STRING},
			{Name: "consumption_based", Type: proto.ColumnType_BOOL, Description: "True if the account is consumption based"},
			{Name: "parent_guid0", Type: proto.ColumnType_STRING},
			{Name: "parent_guid1", Type: proto.ColumnType_STRING},
		},
	}
}

func getDirectory(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {

	fnName := "getDirectory"
	logger := plugin.Logger(ctx)

	logger.Trace("Hydrating list subaccounts")

	btpClient, err := NewBTPClient(nil, d.Connection)
	if err != nil {
		return nil, err
	}

	quals := d.KeyColumnQuals
	logger.Warn(fnName, "quals", quals)
	directoryGuid := quals["guid"].GetStringValue()
	logger.Warn(fnName, "guid", directoryGuid)

	path := directoriesPath + "/" + directoryGuid

	logger.Warn(fnName, "path", path)

	// Call the API
	body, err := btpClient.Get(ctx, AccountsService, path, nil, nil)

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
