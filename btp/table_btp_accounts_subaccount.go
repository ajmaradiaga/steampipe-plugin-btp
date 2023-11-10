package btp

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

const (
	subAccountsPath = "/accounts/v1/subaccounts"
)

func tableBTPAccountsSubaccount() *plugin.Table {
	return &plugin.Table{
		Name:        "btp_accounts_subaccount",
		Description: "BTP Subaccounts",
		List: &plugin.ListConfig{
			Hydrate: listSubaccounts,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("guid"),
			Hydrate:    getSubaccount,
		},
		Columns: []*plugin.Column{
			{Name: "guid", Type: proto.ColumnType_STRING, Description: "Unique ID of the subaccount."},
			{Name: "technical_name", Type: proto.ColumnType_STRING, Description: "The technical name of the subaccount. Refers to: (1) the platform-based account name for Neo subaccounts, or (2) the account identifier (tenant ID) in XSUAA for multi-environment subaccounts."},
			{Name: "display_name", Type: proto.ColumnType_STRING, Description: "A descriptive name of the subaccount for customer-facing UIs."},
			{Name: "global_account_guid", Type: proto.ColumnType_STRING, Description: "The unique ID of the subaccount's global account."},
			{Name: "parent_guid", Type: proto.ColumnType_STRING, Description: "The GUID of the subaccount's parent entity. If the subaccount is located directly in the global account (not in a directory), then this is the GUID of the global account."},
			{Name: "parent_type", Type: proto.ColumnType_STRING},
			{Name: "parent_features", Type: proto.ColumnType_JSON, Description: "The parent features of the subaccount."},
			{Name: "region", Type: proto.ColumnType_STRING, Description: "The region in which the subaccount was created."},
			{Name: "subdomain", Type: proto.ColumnType_STRING, Description: "The subdomain that becomes part of the path used to access the authorization tenant of the subaccount. Must be unique within the defined region. Use only letters (a-z), digits (0-9), and hyphens (not at the start or end). Maximum length is 63 characters. Cannot be changed after the subaccount has been created."},
			{Name: "beta_enabled", Type: proto.ColumnType_BOOL, Description: "Whether the subaccount can use beta services and applications."},
			{Name: "used_for_production", Type: proto.ColumnType_STRING, Description: "Whether the subaccount is used for production purposes. This flag can help your cloud operator to take appropriate action when handling incidents that are related to mission-critical accounts in production systems. Do not apply for subaccounts that are used for non-production purposes, such as development, testing, and demos. Applying this setting this does not modify the subaccount.\n* <b>UNSET:</b> Global account or subaccount admin has not set the production-relevancy flag. Default value.\n* <b>NOT_USED_FOR_PRODUCTION:</b> Subaccount is not used for production purposes.\n* <b>USED_FOR_PRODUCTION:</b> Subaccount is used for production purposes."},
			{Name: "description", Type: proto.ColumnType_STRING, Description: "A description of the subaccount for customer-facing UIs."},
			{Name: "state", Type: proto.ColumnType_STRING, Description: "The current state of the subaccount."},
			{Name: "state_message", Type: proto.ColumnType_STRING, Description: "Information about the state of the subaccount."},
			{Name: "created_date", Type: proto.ColumnType_INT, Description: "The date the subaccount was created. Dates and times are in UTC format."},
			{Name: "created_by", Type: proto.ColumnType_STRING, Description: "Details of the user that created the subaccount."},
			{Name: "modified_date", Type: proto.ColumnType_INT, Description: "The date the subaccount was last modified. Dates and times are in UTC format."},
			{Name: "custom_properties", Type: proto.ColumnType_JSON, Description: "The custom properties of the subaccount."},
		},
	}
}

func listSubaccounts(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {

	fnName := "listSubaccounts"

	logger := plugin.Logger(ctx)
	logger.Trace("Hydrating list subaccounts")

	// conn, err := connect(ctx, d)
	btpClient, err := NewBTPClient(nil, d.Connection)
	if err != nil {
		plugin.Logger(ctx).Error(fmt.Sprintf("%s.%s", d.Table.Name, fnName), "connection_error", err)
		return nil, err
	}

	// Call the API
	body, err := btpClient.Get(ctx, AccountsService, subAccountsPath, nil, nil)

	if err != nil {
		plugin.Logger(ctx).Error(fmt.Sprintf("%s.%s", d.Table.Name, fnName), "api_error", err)
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
		plugin.Logger(ctx).Error(fmt.Sprintf("%s.%s", d.Table.Name, fnName), "api_error", err)
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

	logger.Warn(fnName, "d.Quals", d.Quals)

	subaccountGuid := d.EqualsQualString("guid")
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
