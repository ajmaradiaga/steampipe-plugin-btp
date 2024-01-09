![image](https://hub.steampipe.io/images/plugins/ajmaradiaga/btp-social-graphic.png)

# SAP BTP plugin for Steampipe

Use SQL to query your SAP BTP account details.

- **[Get started →](https://hub.steampipe.io/plugins/ajmaradiaga/btp)**
- Documentation: [Table definitions & examples](https://hub.steampipe.io/plugins/ajmaradiaga/btp/tables)
- Community: [Join #steampipe on Slack →](https://turbot.com/community/join)
- Get involved: [Issues](https://github.com/ajmaradiaga/steampipe-plugin-btp/issues)

## Quick start

### Install

Download and install the latest SAP BTP plugin:

```bash
steampipe plugin install ajmaradiaga/btp
```

Configure your [credentials](https://hub.steampipe.io/plugins/ajmaradiaga/btp#credentials) and [config file](https://hub.steampipe.io/plugins/ajmaradiaga/btp#configuration).

Configure your account details in `~/.steampipe/config/btp.spc`:

```hcl
connection "btp" {
  plugin = "ajmaradiaga/btp"

  # You will need to create a service key for the Cloud Management Service. 
  # You can get the instructions on how to get an access token for the SAP Cloud Management Service APIs here: https://help.sap.com/docs/btp/sap-business-technology-platform/getting-access-token-for-sap-cloud-management-service-apis.

  # URL of the Accounts Service. Required.
  # This can also be set via the `BTP_CIS_ACCOUNTS_SERVICE_URL` environment variable.
  # cis_accounts_service_url = "https://accounts-service.cfapps.[region].hana.ondemand.com"
  
  # URL of the Entitlements Service. Required.
  # This can also be set via the `BTP_CIS_ENTITLEMENTS_SERVICE_URL` environment variable.
  # cis_entitlements_service_url = "https://entitlements-service.cfapps.[region].hana.ondemand.com"

  # Access token to communicate with the Cloud Management Service APIs. Optional.
  # You can get the instructions on how to get an access token for the SAP Cloud Management Service APIs here: https://help.sap.com/docs/btp/sap-business-technology-platform/getting-access-token-for-sap-cloud-management-service-apis. 
  # If no access token is provided, the plugin will try getting an access token using the details provided in cis_client_id, cis_client_secret, cis_token_url, username, password.
  # This can also be set via the `BTP_CIS_ACCESS_TOKEN` environment variable.
  # cis_access_token = "eyJhbGciOiDBNsO0JxFoAaodkDJ3Pmk7cFEsEr5ml5BwNWEafrEjy8Hsxt2mVACpD8B4AIPpRuMoGE71qXGoPcW0vCugceTwN4C3xM8qYmH7DLQrdVIlSX6kydYxnRNjSO8je56ckA4oTC8wm2E2clClPhinDBN6DxHhXlB0eVJnerl4ONpxaH43PYXmHjIsArTuBGK6nCFtApIGN1OvMDPmjHFtOjNgcPCPC5GDXTt5oaB6M2gUrfD5QQVGA7L6yFQlXYPvF6BSyMxpoQXywMUwYA6oqV"

  # A service key (https://help.sap.com/docs/btp/sap-business-technology-platform/creating-service-keys) for the Central Management Service is required to get an access token.  
  # The values for cis_client_id, cis_client_secret, cis_token_url can be retrieved from it. 

  # OAuth 2.0 Client ID (field in the service key: uua.clientid). Optional.
  # This can also be set via the `BTP_CIS_CLIENT_ID` environment variable.
  # cis_client_id = "sb-ut-6y0m0wr1-ai10-1a56-4bv1-52m478h011g6-clone!b017880|cis-central!b14"

  # OAuth 2.0 Client Secret (field in the service key: uua.clientsecret). Optional.
  # This can also be set via the `BTP_CIS_CLIENT_SECRET` environment variable. Optional.
  # cis_client_secret = "44s32264-285r-7101-g2n0-g036p1m214us$vm2_UCHAN_mX1FOW0lklBj2-igBsOUG77G-nE1TWsEu="

  # OAuth 2.0 Token URL (field in the service key: uua.url). Optional.
  # The value in the service key doesn't contain the path, /oauth/token, if not specified the plugin will append it automatically.
  # This can also be set via the `BTP_CIS_TOKEN_URL` environment variable. Optional.
  # cis_token_url = "https://[global-account-subdomain].authentication.[region].hana.ondemand.com"

  # User Email used to log in to SAP BTP.
  # This can also be set via the `BTP_USERNAME` environment variable. Optional.
  # username = "user@domain.com"

  # User Password used to log in to SAP BTP.
  # This can also be set via the `BTP_PASSWORD` environment variable. Optional.
  # password = "My-BTP-Passw0rd"
}
```

Or through environment variables:

```sh
export BTP_CIS_ACCOUNTS_SERVICE_URL=https://accounts-service.cfapps.eu10.hana.ondemand.com
export BTP_CIS_ENTITLEMENTS_SERVICE_URL=https://entitlements-service.cfapps.eu10.hana.ondemand.com
export BTP_CIS_ACCESS_TOKEN=eyJhbGciOiDBNsO0JxFoAaodkDJ3Pmk7cFEsEr5ml5BwNWEafrEjy8Hsxt2mVACpD8B4AIPpRuMoGE71qXGoPcW0vCugceTwN4C3xM8qYmH7DLQ
```

Run steampipe:

```shell
steampipe query
```

List your SAP BTP Global account details:

```sql
select
   guid,
   display_name,
   created_date,
   modified_date 
from
   btp_accounts_global_account;
```

```
+--------------------------------------+-----------------------+---------------+---------------+
| guid                                 | display_name          | created_date  | modified_date |
+--------------------------------------+-----------------------+---------------+---------------+
| 010788v8-7s64-1801-6680-l6g2253646b2 | My BTP global account | 1638221010619 | 1693587625761 |
+--------------------------------------+-----------------------+---------------+---------------+
```

## Engines

This plugin is available for the following engines:

| Engine        | Description
|---------------|------------------------------------------
| [Steampipe](https://steampipe.io/docs) | The Steampipe CLI exposes APIs and services as a high-performance relational database, giving you the ability to write SQL-based queries to explore dynamic data. Mods extend Steampipe's capabilities with dashboards, reports, and controls built with simple HCL. The Steampipe CLI is a turnkey solution that includes its own Postgres database, plugin management, and mod support.
| [Postgres FDW](https://steampipe.io/docs/steampipe_postgres/overview) | Steampipe Postgres FDWs are native Postgres Foreign Data Wrappers that translate APIs to foreign tables. Unlike Steampipe CLI, which ships with its own Postgres server instance, the Steampipe Postgres FDWs can be installed in any supported Postgres database version.
| [SQLite Extension](https://steampipe.io/docs//steampipe_sqlite/overview) | Steampipe SQLite Extensions provide SQLite virtual tables that translate your queries into API calls, transparently fetching information from your API or service as you request it.
| [Export](https://steampipe.io/docs/steampipe_export/overview) | Steampipe Plugin Exporters provide a flexible mechanism for exporting information from cloud services and APIs. Each exporter is a stand-alone binary that allows you to extract data using Steampipe plugins without a database.
| [Turbot Pipes](https://turbot.com/pipes/docs) | Turbot Pipes is the only intelligence, automation & security platform built specifically for DevOps. Pipes provide hosted Steampipe database instances, shared dashboards, snapshots, and more.

## Developing

Prerequisites:

- [Steampipe](https://steampipe.io/downloads)
- [Golang](https://golang.org/doc/install)

Clone:

```sh
git clone https://github.com/ajmaradiaga/steampipe-plugin-btp.git
cd steampipe-plugin-btp
```

Build, which automatically installs the new version to your `~/.steampipe/plugins` directory:

```
make
```

Configure the plugin:

```
cp config/* ~/.steampipe/config
vi ~/.steampipe/config/btp.spc
```

Try it!

```
steampipe query
> .inspect btp
```

Further reading:

- [Writing plugins](https://steampipe.io/docs/develop/writing-plugins)
- [Writing your first table](https://steampipe.io/docs/develop/writing-your-first-table)

## Open Source & Contributing

This repository is published under the [Apache 2.0](https://www.apache.org/licenses/LICENSE-2.0) (source code) and [CC BY-NC-ND](https://creativecommons.org/licenses/by-nc-nd/2.0/) (docs) licenses. Please see our [code of conduct](https://github.com/turbot/.github/blob/main/CODE_OF_CONDUCT.md). We look forward to collaborating with you!

[Steampipe](https://steampipe.io) is a product produced from this open source software, exclusively by [Turbot HQ, Inc](https://turbot.com). It is distributed under our commercial terms. Others are allowed to make their own distribution of the software, but cannot use any of the Turbot trademarks, cloud services, etc. You can learn more in our [Open Source FAQ](https://turbot.com/open-source).

## Get Involved

**[Join #steampipe on Slack →](https://turbot.com/community/join)**

Want to help but don't know where to start? Pick up one of the `help wanted` issues:

- [Steampipe](https://github.com/turbot/steampipe/labels/help%20wanted)
- [SAP BTP Plugin](https://github.com/ajmaradiaga/steampipe-plugin-btp/labels/help%20wanted)
