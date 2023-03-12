package btp

import (
	"context"
	"encoding/json"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
)

const (
	datacentersPath = "/entitlements/v1/globalAccountAllowedDataCenters"
)

func tableBTPDatacenters() *plugin.Table {
	return &plugin.Table{
		Name:        "btp_entitlements_alloweddatacenters",
		Description: "BTP Allowed Data Centers",
		List: &plugin.ListConfig{
			KeyColumns: plugin.OptionalColumns([]string{"region"}),
			Hydrate:    listDataCenters,
		},
		Columns: []*plugin.Column{
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Data center name"},
			{Name: "display_name", Type: proto.ColumnType_STRING, Description: "Data center display name"},
			{Name: "region", Type: proto.ColumnType_STRING},
			{Name: "environment", Type: proto.ColumnType_STRING},
			{Name: "iaas_provider", Type: proto.ColumnType_STRING, Description: "Infrastructure-as-a-service provider, e.g. AWS, Azure, GCP"},
			{Name: "supports_trial", Type: proto.ColumnType_BOOL, Description: "True if the data center supports trial"},
			{Name: "provisioning_service_url", Type: proto.ColumnType_STRING},
			{Name: "saas_registry_service_url", Type: proto.ColumnType_STRING},
			{Name: "domain", Type: proto.ColumnType_STRING},
			{Name: "geo_access", Type: proto.ColumnType_STRING},
		},
	}
}

func listDataCenters(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {

	fnName := "listAllowedDataCenters"
	logger := plugin.Logger(ctx)

	logger.Trace("Hydrating list allowed data centers")

	btpClient, err := NewBTPClient(nil, d.Connection)
	if err != nil {
		return nil, err
	}

	quals := d.KeyColumnQuals
	logger.Warn(fnName, "quals", quals)
	region, ok := quals["region"]

	queryStrings := map[string]string{}

	if ok {
		queryStrings["region"] = region.GetStringValue()
	}

	path := datacentersPath

	logger.Warn(fnName, "path", path)

	// Call the API
	body, err := btpClient.Get(ctx, EntitlementService, path, nil, queryStrings)

	if err != nil {
		return nil, err
	}

	var data struct {
		Datacenters []DataCenter `json:"datacenters"`
	}

	// Convert JSON response to Subaccounts structure
	err = json.Unmarshal(body, &data)

	logger.Warn(fnName, "data", data)
	logger.Warn(fnName, "err", err)

	if err != nil {
		return nil, err
	}

	for _, s := range data.Datacenters {
		d.StreamListItem(ctx, s)
	}

	return nil, nil
}
