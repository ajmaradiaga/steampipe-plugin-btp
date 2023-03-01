
Commands
```bash
steampipe plugin install steampipe
steampipe query "select name from steampipe_registry_plugin;"
steampipe plugin install rss
steampipe query
go install github.com/turbot/steampipe-plugin-sdk/plugin@latest
ls ~/.steampipe/plugins/
cp btp/config/btp.spc ~/.steampipe/config/btp.spc
make
```

```SQL
select guid, display_name, description, parent_guid, state_message from btp_subaccounts
select * from btp_accounts_global_account
```