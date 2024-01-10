## v0.1.2 [2024-01-10]

_Fixes:_

- Fixing issues in markdown table HTML tags

## v0.1.1 [2024-01-10]

_Fixes:_

- Minor issue in documentation

## v0.1.0 [2024-01-09]

_What's new?_

- The plugin now supports:
  - Reading service key details from a file: There is no need to specify lots of variables in the connection if a path to the service key is provided in the config cis_service_key_path (BTP_CIS_SERVICE_KEY_PATH).
  - Authenticating on behalf of a user to get an access token. Two new variables are required in the config: username (BTP_USERNAME) and password (BTP_PASSWORD). Note: The access token will be cached by the plugin. 

## v0.0.2 [2024-01-03]

_What's new?_

- Aligning documentation with Turbot release week changes

## v0.0.1 [2023-12-15]

_What's new?_

- New tables added
  
  - [btp_accounts_directory](https://github.com/ajmaradiaga/steampipe-plugin-btp/blob/main/btp/table_btp_accounts_directory.go)
  - [btp_accounts_global_account](https://github.com/ajmaradiaga/steampipe-plugin-btp/blob/main/btp/table_btp_accounts_global_account.go)
  - [btp_accounts_subaccount](https://github.com/ajmaradiaga/steampipe-plugin-btp/blob/main/btp/table_btp_accounts_subaccount.go)
  - [btp_entitlements_assignment](https://github.com/ajmaradiaga/steampipe-plugin-btp/blob/main/btp/table_btp_entitlements_assignment.go)
  - [btp_entitlements_datacenter](https://github.com/ajmaradiaga/steampipe-plugin-btp/blob/main/btp/table_btp_entitlements_datacenter.go)
