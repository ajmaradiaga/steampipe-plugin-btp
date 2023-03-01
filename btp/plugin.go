package btp

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
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
			// "zendesk_brand":        tableZendeskBrand(),
			// "zendesk_group":        tableZendeskGroup(),
			// "zendesk_organization": tableZendeskOrganization(),
			// "zendesk_search":       tableZendeskSearch(),
			// "zendesk_ticket":       tableZendeskTicket(),
			// "zendesk_ticket_audit": tableZendeskTicketAudit(),
			// "zendesk_trigger":      tableZendeskTrigger(),
			"btp_subaccounts":             tableBTPSubaccounts(),
			"btp_accounts_global_account": tableBTPAccountsGlobalAccount(),
		},
	}
	return p
}
