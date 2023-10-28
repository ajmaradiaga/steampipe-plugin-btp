package btp

import (
	"context"
	"encoding/json"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
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
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Technical name of the data center. Must be unique within the cloud deployment"},
			{Name: "display_name", Type: proto.ColumnType_STRING, Description: "Descriptive name of the data center for customer-facing UIs"},
			{Name: "region", Type: proto.ColumnType_STRING, Description: "The region in which the data center is located"},
			{Name: "environment", Type: proto.ColumnType_STRING, Description: "The environment that the data center supports. For example: Kubernetes, Cloud Foundry"},
			{Name: "iaas_provider", Type: proto.ColumnType_STRING, Description: "The infrastructure provider for the data center. Valid values: AWS, GCP, AZURE, SAP: SAP BTP (Neo), ALI: Alibaba Cloud, IBM: IBM Cloud."},
			{Name: "supports_trial", Type: proto.ColumnType_BOOL, Description: "Whether the specified datacenter supports trial accounts"},
			{Name: "provisioning_service_url", Type: proto.ColumnType_STRING, Description: "Provisioning service URL"},
			{Name: "saas_registry_service_url", Type: proto.ColumnType_STRING, Description: "Saas-Registry service URL"},
			{Name: "domain", Type: proto.ColumnType_STRING, Description: "The domain of the data center"},
			{Name: "is_main_data_center", Type: proto.ColumnType_BOOL, Description: "Whether the specified datacenter is a main datacenter"},
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

	equalQuals := d.EqualsQuals
	logger.Warn(fnName, "d.Quals", d.Quals)

	queryStrings := map[string]string{}

	// Including satellite data centers by default
	queryStrings["includeSatelliteDataCenters"] = "true"

	if d.Quals["region"] != nil {
		queryStrings["region"] = equalQuals["region"].GetStringValue()
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
