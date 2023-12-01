# SAP BTP plugin for Steampipe

Use SQL to query your SAP BTP account details.

- **[Get started →](https://hub.steampipe.io/plugins/ajmaradiaga/btp)**
- Documentation: [Table definitions & examples](https://hub.steampipe.io/plugins/ajmaradiaga/btp/tables)
- Get involved: [Issues](https://github.com/ajmaradiaga/steampipe-plugin-btp/issues)

## Quick start

### Install

Download and install the latest SAP BTP plugin:

```bash
steampipe plugin install btp
```

Configure your [credentials](https://hub.steampipe.io/plugins/ajmaradiaga/btp#credentials) and [config file](https://hub.steampipe.io/plugins/ajmaradiaga/btp#configuration).

Configure your account details in `~/.steampipe/config/btp.spc`:

```hcl
connection "btp" {
  plugin = "ajmaradiaga/btp"

  # You will need to create a service key for the Cloud Management Service. You can get the instructions on how to get an access token for the SAP Cloud Management Service APIs here: https://help.sap.com/docs/btp/sap-business-technology-platform/getting-access-token-for-sap-cloud-management-service-apis.

  # URL of the Accounts Service. Required.
  # This can also be set via the `BTP_CIS_ACCOUNTS_SERVICE_URL` environment variable.
  # cis_accounts_service_url = "https://accounts-service.cfapps.[region].hana.ondemand.com"
  
  # URL of the Entitlements Service. Required.
  # This can also be set via the `BTP_CIS_ENTITLEMENTS_SERVICE_URL` environment variable.
  # cis_entitlements_service_url = "https://entitlements-service.cfapps.[region].hana.ondemand.com"

  # Access token to communicate with the Cloud Management Service APIs. At the moment only access token is supported. You can get the instructions on how to get an access token for the SAP Cloud Management Service APIs here: https://help.sap.com/docs/btp/sap-business-technology-platform/getting-access-token-for-sap-cloud-management-service-apis. Required.
  # This can also be set via the `BTP_CIS_ACCESS_TOKEN` environment variable.
  # cis_access_token = "eyJhbGciOiDBNsO0JxFoAaodkDJ3Pmk7cFEsEr5ml5BwNWEafrEjy8Hsxt2mVACpD8B4AIPpRuMoGE71qXGoPcW0vCugceTwN4C3xM8qYmH7DLQrdVIlSX6kydYxnRNjSO8je56ckA4oTC8wm2E2clClPhinDBN6DxHhXlB0eVJnerl4ONpxaH43PYXmHjIsArTuBGK6nCFtApIGN1OvMDPmjHFtOjNgcPCPC5GDXTt5oaB6M2gUrfD5QQVGA7L6yFQlXYPvF6BSyMxpoQXywMUwYA6oqV"
 
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
   btp.btp_accounts_global_account;
```

```
+--------------------------------------+-----------------------+---------------+---------------+
| guid                                 | display_name          | created_date  | modified_date |
+--------------------------------------+-----------------------+---------------+---------------+
| 010788v8-7s64-1801-6680-l6g2253646b2 | My BTP global account | 1638221010619 | 1693587625761 |
+--------------------------------------+-----------------------+---------------+---------------+
```

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

## Contributing

Please see the [contribution guidelines](https://github.com/ajmaradiaga/steampipe-plugin-btp/blob/main/CONTRIBUTING.md) and our [code of conduct](https://github.com/ajmaradiaga/steampipe-plugin-btp/blob/main/CODE_OF_CONDUCT.md). All contributions are subject to the [Apache 2.0 open source license](https://github.com/ajmaradiaga/steampipe-plugin-btp/blob/main/LICENSE).

`help wanted` issues:

- [Steampipe](https://github.com/turbot/steampipe/labels/help%20wanted)
- [SAP BTP Plugin](https://github.com/ajmaradiaga/steampipe-plugin-btp/labels/help%20wanted)
