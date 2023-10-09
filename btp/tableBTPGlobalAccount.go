package btp

import (
	"context"
	"encoding/json"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
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
			{Name: "guid", Type: proto.ColumnType_STRING, Description: "The unique ID of the global account"},
			{Name: "display_name", Type: proto.ColumnType_STRING, Description: "The display name of the global account"},
			{Name: "created_date", Type: proto.ColumnType_INT, Description: "The date the global account was created. Dates and times are in UTC format"},
			{Name: "modified_date", Type: proto.ColumnType_INT, Description: "The date the global account was last modified. Dates and times are in UTC format"},
			{Name: "entity_state", Type: proto.ColumnType_STRING, Description: "The current state of the global account.\n* <b>STARTED:</b> CRUD operation on an entity has started.\n* <b>CREATING:</b> Creating entity operation is in progress.\n* <b>UPDATING:</b> Updating entity operation is in progress.\n* <b>MOVING:</b> Moving entity operation is in progress.\n* <b>PROCESSING:</b> A series of operations related to the entity is in progress.\n* <b>DELETING:</b> Deleting entity operation is in progress.\n* <b>OK:</b> The CRUD operation or series of operations completed successfully.\n* <b>PENDING REVIEW:</b> The processing operation has been stopped for reviewing and can be restarted by the operator.\n* <b>CANCELLED:</b> The operation or processing was canceled by the operator.\n* <b>CREATION_FAILED:</b> The creation operation failed, and the entity was not created or was created but cannot be used.\n* <b>UPDATE_FAILED:</b> The update operation failed, and the entity was not updated.\n* <b>PROCESSING_FAILED:</b> The processing operations failed.\n* <b>DELETION_FAILED:</b> The delete operation failed, and the entity was not deleted.\n* <b>MOVE_FAILED:</b> Entity could not be moved to a different location.\n* <b>MIGRATING:</b> Migrating entity from NEO to CF"},
			{Name: "state_message", Type: proto.ColumnType_STRING, Description: "Information about the state"},
			{Name: "subdomain", Type: proto.ColumnType_STRING, Description: "Relevant only for entities that require authorization (e.g. global account). The subdomain that becomes part of the path used to access the authorization tenant of the global account. Unique within the defined region"},
			{Name: "contract_status", Type: proto.ColumnType_STRING, Description: "The status of the customer contract and its associated root global account.\n* <b>ACTIVE:</b> The customer contract and its associated global account is currently active.\n* <b>PENDING_TERMINATION:</b> A termination process has been triggered for a customer contract (the customer contract has expired, or a customer has given notification that they wish to terminate their contract), and the global account is currently in the validation period. The customer can still access their global account until the end of the validation period.\n* <b>SUSPENDED:</b> For enterprise accounts, specifies that the customer's global account is currently in the grace period of the termination process. Access to the global account by the customer is blocked. No data is deleted until the deletion date is reached at the end of the grace period. For trial accounts, specifies that the account is suspended, and the account owner has not yet extended the trial period"},
			{Name: "commercial_model", Type: proto.ColumnType_STRING, Description: "The type of the commercial contract that was signed"},
			{Name: "consumption_based", Type: proto.ColumnType_BOOL, Description: "Whether the customer of the global account pays only for services that they actually use (consumption-based) or pay for subscribed services at a fixed cost irrespective of consumption (subscription-based).\n* <b>TRUE:</b> Consumption-based commercial model.\n* <b>FALSE:</b> Subscription-based commercial model"},
			{Name: "license_type", Type: proto.ColumnType_STRING, Description: "The type of license for the global account. The license type affects the scope of functions of the account.\n* <b>DEVELOPER:</b> For internal developer global accounts on Staging or Canary landscapes.\n* <b>CUSTOMER:</b> For customer global accounts.\n* <b>PARTNER:</b> For partner global accounts.\n* <b>INTERNAL_DEV:</b> For internal global accounts on the Dev landscape.\n* <b>INTERNAL_PROD:</b> For internal global accounts on the Live landscape.\n* <b>TRIAL:</b> For customer trial accounts"},
			{Name: "geo_access", Type: proto.ColumnType_STRING, Description: "The geographic locations from where the global account can be accessed.\n* <b>STANDARD:</b> The global account can be accessed from any geographic location.\n* <b>EU_ACCESS:</b> The global account can be accessed only within locations in the EU"},
			{Name: "renewal_date", Type: proto.ColumnType_INT, Description: "The date that an expired contract was renewed. Dates and times are in UTC format"},
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
