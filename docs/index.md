---
organization: ajmaradiaga
category: ["paas"]
icon_url: "/images/plugins/ajmaradiaga/btp.svg"
brand_color: "#002A86"
display_name: "SAP BTP"
short_name: "btp"
description: "Steampipe plugin to query the account details of your SAP Business Technology Platform account."
og_description: "Query SAP BTP with SQL! Open source CLI. No DB required."
og_image: "/images/plugins/ajmaradiaga/sap-btp-graphic.png"
engines: ["steampipe", "sqlite", "postgres", "export"]
---

# SAP BTP + Steampipe

[SAP BTP](https://www.sap.com/products/technology-platform.html) brings together data and analytics, artificial intelligence, application development, automation, and integration in one, unified environment.

[Steampipe](https://steampipe.io) is an open-source zero-ETL engine to instantly query cloud APIs using SQL.

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

## Documentation

- **[Table definitions & examples →](/plugins/ajmaradiaga/btp/tables)**

## Quick start

### Install

Download and install the latest SAP BTP plugin:

```sh
steampipe plugin install ajmaradiaga/btp
```

### Credentials

| Item        | Description                                                                                                                                                                                           |
| ----------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Credentials | You will need to create a service key for the Cloud Management Service. You can get the instructions on how to get an access token for the SAP Cloud Management Service APIs here: https://help.sap.com/docs/btp/sap-business-technology-platform/getting-access-token-for-sap-cloud-management-service-apis.                                                                |                                                               |
| Permissions | Create a [Cloud Management Service with a `Central` service plan](https://discovery-center.cloud.sap/serviceCatalog/cloud-management-service?region=all&tab=service_plan) to manage your global account, subaccounts, directories, and entitlements.  |
| Radius | Each connection represents a single SAP BTP account. |
<<<<<<< HEAD
| Resolution  | You can authenticate by providing an access token or the details of a service key in the connection config file. The plugin prioritises an access token if one is provided.<br/><ol><li>The access token can be provided via environment variables or in the connection config file. <br/><ol><li>Access token specified in environment variables, e.g., `BTP_CIS_ACCESS_TOKEN`.</li><li>Access token explicitly set in a steampipe config file (`~/.steampipe/config/btp.spc`)</li></ol><li>Service key details can be provided by specifying a file path to the service key via environment variable (`BTP_CIS_SERVICE_KEY_PATH`) or in the connection config file (`cis_service_key_path`).<ol><li>Service key details provided in environment variables, e.g., `BTP_CIS_CLIENT_ID`, `BTP_CIS_CLIENT_SECRET`, `BTP_CIS_TOKEN_URL` will be prioritised over the values in the config file.</li><li>Service key details explicitly set in a steampipe config file (`~/.steampipe/config/btp.spc`)</li></ol></ol><br/>Generally the plugin prioritises environment variables then values in the config file and lastly values in the service key file. |
=======
| Resolution  | You can authenticate by providing an access token or the details of a service key in the connection config file. The plugin prioritises an access token if one is provided.<br/><ol><li>The access token can be provided via environment variables or in the connection config file. <br/><ol><li>Access token specified in environment variables, e.g., `BTP_CIS_ACCESS_TOKEN`.</li><li>Access token explicitly set in a steampipe config file (`~/.steampipe/config/btp.spc`)</li></ol><li>Service key details can be provided via environment variables or in the connection config file.<ol><li>Service key details provided in environment variables, e.g., `BTP_CIS_CLIENT_ID`, `BTP_CIS_CLIENT_SECRET`, `BTP_CIS_TOKEN_URL`.</li><li>Service key details explicitly set in a steampipe config file (`~/.steampipe/config/btp.spc`)</li></ol></ol> |
>>>>>>> 2f53a7e (Including additional config supported)

### Configuration

Installing the latest SAP BTP plugin will create a config file (`~/.steampipe/config/btp.spc`) with a single connection named `btp`:

Configure your account details in `~/.steampipe/config/btp.spc`:

```hcl
connection "btp" {
  plugin = "ajmaradiaga/btp"

  # User email used to log in to SAP BTP. Required.
  # This can also be set via the `BTP_USERNAME` environment variable.
  # username = "user@domain.com"

  # User password used to log in to SAP BTP. Required.
  # This can also be set via the `BTP_PASSWORD` environment variable.
  # password = "My-BTP-Passw0rd"

  # You will need to create a service key for the Cloud Management Service. Required.
  # You can get the instructions on how to get an access token for the SAP Cloud Management Service APIs here: https://help.sap.com/docs/btp/sap-business-technology-platform/getting-access-token-for-sap-cloud-management-service-apis.
  # This can also be set via the `BTP_CIS_SERVICE_KEY_PATH` environment variable. Required.
  # cis_service_key_path = "~/service_keys/cis_global.json"

  # URL of the Accounts Service. Optional.
  # This can also be set via the `BTP_CIS_ACCOUNTS_SERVICE_URL` environment variable.
  # cis_accounts_service_url = "https://accounts-service.cfapps.[region].hana.ondemand.com"
  
  # URL of the Entitlements Service. Optional.
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
<<<<<<< HEAD
=======

  # User Email used to log in to SAP BTP.
  # This can also be set via the `BTP_USERNAME` environment variable. Optional.
  # username = "user@domain.com"

  # User Password used to log in to SAP BTP.
  # This can also be set via the `BTP_PASSWORD` environment variable. Optional.
  # password = "My-BTP-Passw0rd"
>>>>>>> 2f53a7e (Including additional config supported)
}
```

Alternatively, you can also use the environment variables specified below:

```sh
export BTP_USERNAME=user@domain.com
export BTP_PASSWORD=My-BTP-Passw0rd
export BTP_CIS_SERVICE_KEY_PATH=~/service_keys/cis_global.json
export BTP_CIS_ACCOUNTS_SERVICE_URL=https://accounts-service.cfapps.eu10.hana.ondemand.com
export BTP_CIS_ENTITLEMENTS_SERVICE_URL=https://entitlements-service.cfapps.eu10.hana.ondemand.com
export BTP_CIS_ACCESS_TOKEN=eyJhbGciOiDBNsO0JxFoAaodkDJ3Pmk7cFEsEr5ml5BwNWEafrEjy8Hsxt2mVACpD8B4AIPpRuMoGE71qXGoPcW0vCugceTwN4C3xM8qYmH7DLQ
```

## Get involved

- Open source: https://github.com/ajmaradiaga/steampipe-plugin-btp
- Community: [Join #steampipe on Slack →](https://turbot.com/community/join)
