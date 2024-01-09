package btp

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

const (
	globalAccountAssignmentsPath = "/entitlements/v1/globalAccountAssignments"
)

func tableBTPEntitlementsAssignment() *plugin.Table {
	return &plugin.Table{
		Name:        "btp_entitlements_assignment",
		Description: "BTP Global account assignments",
		List: &plugin.ListConfig{
			Hydrate: listGlobalAccountAssignments,
		},
		Columns: []*plugin.Column{
			{Name: "name", Type: proto.ColumnType_STRING, Description: "The unique registration name of the deployed service as defined by the service provider."},
			{Name: "display_name", Type: proto.ColumnType_STRING, Description: "Display name of the service for customer-facing UIs."},
			{Name: "description", Type: proto.ColumnType_STRING, Description: "Description of the service for customer-facing UIs."},
			{Name: "business_category", Type: proto.ColumnType_JSON, Description: "Services grouped according to capabilities and customer’s business needs, for example, Extension Suite - Digital Experience, Extension Suite - Development Efficiency, and Extension Suite - Digital Process Automation. Possible values: OTHER, APPLICATION_DEVELOPMENT_AND_AUTOMATION, INTEGRATION, FOUNDATION_CROSS_SERVICES, AI, DATA_AND_ANALYTICS, EXTENDED_PLANNING_AND_ANALYSIS."},
			{Name: "owner_type", Type: proto.ColumnType_STRING, Description: "The owner type of the service. Possible values:\nVENDOR: The owner is a service owner, who is affiliated with the cloud operator, that added the service to the product catalog for general consumption.\nCUSTOMER: The owner is an SAP customer that added a custom service to the product catalog, and it is available only for consumption within the customer's global account.\nPARTNER: The owner is an SAP partner that added the service to the product catalog, and it is available only to their customers for consumption."},
			{Name: "terms_of_use_url", Type: proto.ColumnType_STRING, Description: "Terms of use."},
			{Name: "service_plans", Type: proto.ColumnType_JSON, Description: "List of service plans associated with the entitled service."},
			{Name: "icon_base64", Type: proto.ColumnType_STRING, Description: "The icon of the service in Base64 format."},
		},
	}
}

func listGlobalAccountAssignments(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {

	fnName := "listGlobalAccountAssignments"
	logger := plugin.Logger(ctx)

	logger.Trace("Hydrating list of assignments")

	btpClient, err := NewBTPClient(nil, ctx, d)
	if err != nil {
		logger.Error(fmt.Sprintf("%s.%s", d.Table.Name, fnName), "connection_error", err)
		return nil, err
	}

	logger.Warn(fnName, "d.Quals", d.Quals)

	queryStrings := map[string]string{}

	if d.Quals["subaccount_guid"] != nil {
		queryStrings["subaccountGUID"] = d.EqualsQualString("subaccount_guid")
	}

	path := globalAccountAssignmentsPath

	logger.Warn(fnName, "path", path)

	// Call the API
	body, err := btpClient.Get(ctx, EntitlementService, path, nil, queryStrings)

	if err != nil {
		logger.Error(fmt.Sprintf("%s.%s", d.Table.Name, fnName), "api_error", err)
		return nil, err
	}

	var data struct {
		EntitledServices []EntitledService `json:"entitledServices"`
	}

	// Convert JSON response to structure
	err = json.Unmarshal(body, &data)

	logger.Warn(fnName, "data", data)
	logger.Warn(fnName, "err", err)

	if err != nil {
		logger.Error(fmt.Sprintf("%s.%s", d.Table.Name, fnName), "api_error", err)
		return nil, err
	}

	for _, s := range data.EntitledServices {
		d.StreamListItem(ctx, s)
	}

	return nil, nil
}
