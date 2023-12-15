package btp

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name: "btp",
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
		},
		DefaultTransform: transform.FromGo().NullIfZero(),
		TableMap: map[string]*plugin.Table{
			"btp_accounts_directory":      tableBTPAccountsDirectory(),
			"btp_accounts_global_account": tableBTPAccountsGlobalAccount(),
			"btp_accounts_subaccount":     tableBTPAccountsSubaccount(),
			"btp_entitlements_assignment": tableBTPEntitlementsAssignment(),
			"btp_entitlements_datacenter": tableBTPEntitlementsDatacenter(),
		},
	}
	return p
}
