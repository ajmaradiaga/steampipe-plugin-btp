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

func tableBTPAccountsGlobalAccount() *plugin.Table {
	return &plugin.Table{
		Name:        "btp_accounts_global_account",
		Description: "BTP Global Account details",
		List: &plugin.ListConfig{
			Hydrate: getGlobalAccount,
		},
		Columns: []*plugin.Column{
			{Name: "guid", Type: proto.ColumnType_STRING, Description: "The unique ID of the global account"},
			{Name: "display_name", Type: proto.ColumnType_STRING, Description: "The display name of the global account"},
			{Name: "created_date", Type: proto.ColumnType_INT, Description: "The date the global account was created. Dates and times are in UTC format"},
			{Name: "modified_date", Type: proto.ColumnType_INT, Description: "The date the global account was last modified. Dates and times are in UTC format"},
			{Name: "entity_state", Type: proto.ColumnType_STRING, Description: "The current state of the global account.\n* STARTED: CRUD operation on an entity has started.\n* CREATING: Creating entity operation is in progress.\n* UPDATING: Updating entity operation is in progress.\n* MOVING: Moving entity operation is in progress.\n* PROCESSING: A series of operations related to the entity is in progress.\n* DELETING: Deleting entity operation is in progress.\n* OK: The CRUD operation or series of operations completed successfully.\n* PENDING REVIEW: The processing operation has been stopped for reviewing and can be restarted by the operator.\n* CANCELLED: The operation or processing was canceled by the operator.\n* CREATION_FAILED: The creation operation failed, and the entity was not created or was created but cannot be used.\n* UPDATE_FAILED: The update operation failed, and the entity was not updated.\n* PROCESSING_FAILED: The processing operations failed.\n* DELETION_FAILED: The delete operation failed, and the entity was not deleted.\n* MOVE_FAILED: Entity could not be moved to a different location.\n* MIGRATING: Migrating entity from NEO to CF"},
			{Name: "state_message", Type: proto.ColumnType_STRING, Description: "Information about the state"},
			{Name: "subdomain", Type: proto.ColumnType_STRING, Description: "Relevant only for entities that require authorization (e.g. global account). The subdomain that becomes part of the path used to access the authorization tenant of the global account. Unique within the defined region"},
			{Name: "contract_status", Type: proto.ColumnType_STRING, Description: "The status of the customer contract and its associated root global account.\n* ACTIVE: The customer contract and its associated global account is currently active.\n* PENDING_TERMINATION: A termination process has been triggered for a customer contract (the customer contract has expired, or a customer has given notification that they wish to terminate their contract), and the global account is currently in the validation period. The customer can still access their global account until the end of the validation period.\n* SUSPENDED: For enterprise accounts, specifies that the customer's global account is currently in the grace period of the termination process. Access to the global account by the customer is blocked. No data is deleted until the deletion date is reached at the end of the grace period. For trial accounts, specifies that the account is suspended, and the account owner has not yet extended the trial period"},
			{Name: "commercial_model", Type: proto.ColumnType_STRING, Description: "The type of the commercial contract that was signed"},
			{Name: "consumption_based", Type: proto.ColumnType_BOOL, Description: "Whether the customer of the global account pays only for services that they actually use (consumption-based) or pay for subscribed services at a fixed cost irrespective of consumption (subscription-based).\n* TRUE: Consumption-based commercial model.\n* FALSE: Subscription-based commercial model"},
			{Name: "license_type", Type: proto.ColumnType_STRING, Description: "The type of license for the global account. The license type affects the scope of functions of the account.\n* DEVELOPER: For internal developer global accounts on Staging or Canary landscapes.\n* CUSTOMER: For customer global accounts.\n* PARTNER: For partner global accounts.\n* INTERNAL_DEV: For internal global accounts on the Dev landscape.\n* INTERNAL_PROD: For internal global accounts on the Live landscape.\n* TRIAL: For customer trial accounts"},
			{Name: "geo_access", Type: proto.ColumnType_STRING, Description: "The geographic locations from where the global account can be accessed.\n* STANDARD: The global account can be accessed from any geographic location.\n* EU_ACCESS: The global account can be accessed only within locations in the EU"},
			{Name: "renewal_date", Type: proto.ColumnType_INT, Description: "The date that an expired contract was renewed. Dates and times are in UTC format"},
			{Name: "subaccounts", Type: proto.ColumnType_JSON, Description: "The subaccounts in the global account"},
		},
	}
}

func getGlobalAccount(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {

	logger := plugin.Logger(ctx)
	logger.Trace("Hydrating Global Account")

	btpClient, err := NewBTPClient(nil, d.Connection)

	if err != nil {
		return nil, err
	}

	queryStrings := map[string]string{
		"derivedAuthorizations": "any",
		"expand":                "true",
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

	logger.Debug("getGlobalAccount", "data", data)
	logger.Debug("getGlobalAccount", "err", err)

	if err != nil {
		return nil, err
	}

	d.StreamListItem(ctx, data)

	return nil, nil
}
