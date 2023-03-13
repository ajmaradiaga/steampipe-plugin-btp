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
			{Name: "guid", Type: proto.ColumnType_STRING, Description: "The unique ID of the directory"},
			{Name: "parent_type", Type: proto.ColumnType_STRING, Description: "The Type of the directory parent entity"},
			{Name: "global_account_guid", Type: proto.ColumnType_STRING, Description: "The GUID of the directory's global account entity"},
			{Name: "display_name", Type: proto.ColumnType_STRING, Description: "The display name of the directory"},
			{Name: "created_date", Type: proto.ColumnType_INT, Description: "The date the directory was created. Dates and times are in UTC format"},
			{Name: "created_by", Type: proto.ColumnType_STRING, Description: "Details of the user that created the directory"},
			{Name: "modified_date", Type: proto.ColumnType_INT, Description: "The date the directory was last modified. Dates and times are in UTC format"},
			{Name: "entity_state", Type: proto.ColumnType_STRING, Description: "The current state of the directory.\n* <b>STARTED:</b> CRUD operation on an entity has started.\n* <b>CREATING:</b> Creating entity operation is in progress.\n* <b>UPDATING:</b> Updating entity operation is in progress.\n* <b>MOVING:</b> Moving entity operation is in progress.\n* <b>PROCESSING:</b> A series of operations related to the entity is in progress.\n* <b>DELETING:</b> Deleting entity operation is in progress.\n* <b>OK:</b> The CRUD operation or series of operations completed successfully.\n* <b>PENDING REVIEW:</b> The processing operation has been stopped for reviewing and can be restarted by the operator.\n* <b>CANCELLED:</b> The operation or processing was canceled by the operator.\n* <b>CREATION_FAILED:</b> The creation operation failed, and the entity was not created or was created but cannot be used.\n* <b>UPDATE_FAILED:</b> The update operation failed, and the entity was not updated.\n* <b>PROCESSING_FAILED:</b> The processing operations failed.\n* <b>DELETION_FAILED:</b> The delete operation failed, and the entity was not deleted.\n* <b>MOVE_FAILED:</b> Entity could not be moved to a different location.\n* <b>MIGRATING:</b> Migrating entity from NEO to CF"},
			{Name: "state_message", Type: proto.ColumnType_STRING, Description: "Information about the state"},
			{Name: "directory_type", Type: proto.ColumnType_STRING},
			{Name: "contract_status", Type: proto.ColumnType_STRING, Description: "The status of the customer contract and its associated root global account.\n* <b>ACTIVE:</b> The customer contract and its associated global account is currently active.\n* <b>PENDING_TERMINATION:</b> A termination process has been triggered for a customer contract (the customer contract has expired, or a customer has given notification that they wish to terminate their contract), and the global account is currently in the validation period. The customer can still access their global account until the end of the validation period.\n* <b>SUSPENDED:</b> For enterprise accounts, specifies that the customer's global account is currently in the grace period of the termination process. Access to the global account by the customer is blocked. No data is deleted until the deletion date is reached at the end of the grace period. For trial accounts, specifies that the account is suspended, and the account owner has not yet extended the trial period"},
			{Name: "consumption_based", Type: proto.ColumnType_BOOL, Description: "True if the account is consumption based"},
			{Name: "parent_guid0", Type: proto.ColumnType_STRING, Description: "The GUID of the directory's parent entity. Typically this is the global account"},
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
