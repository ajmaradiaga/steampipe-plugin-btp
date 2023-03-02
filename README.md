
Commands
```bash
steampipe plugin install steampipe
steampipe query "select name from steampipe_registry_plugin;"
steampipe plugin install rss
steampipe query
go install github.com/turbot/steampipe-plugin-sdk/plugin@latest
ls ~/.steampipe/plugins/
cp btp/config/btp.spc ~/.steampipe/config/btp.spc

# Set steampipe log level
export STEAMPIPE_LOG_LEVEL=LEVEL=WARN

# Runs the command in Makefile and places the plugin in the steampipe
# plugins folder, e.g. ~/.steampipe/plugins/local/btp
make

# Run make and query
make && steampipe query "select * from btp_accounts_global_account"
```

```SQL
select guid, display_name, description, parent_guid, state_message from btp_subaccounts
select * from btp_accounts_global_account
```