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
			Schema:      ConfigSchema,
		},
		DefaultTransform: transform.FromGo().NullIfZero(),
		TableMap: map[string]*plugin.Table{
			"btp_accounts_global_account":         tableBTPGlobalAccount(),
			"btp_accounts_subaccounts":            tableBTPSubaccounts(),
			"btp_accounts_directories":            tableBTPDirectories(),
			"btp_entitlements_alloweddatacenters": tableBTPDatacenters(),
			"btp_entitlements_assignments":        tableBTPGlobalAccountAssignments(),
		},
	}
	return p
}
