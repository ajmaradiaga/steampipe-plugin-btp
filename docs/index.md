---
organization: ajmaradiaga
category: ["public cloud"]
icon_url: "/images/plugins/ajmaradiaga/btp.svg"
brand_color: "#009FDD"
display_name: "SAP BTP"
short_name: "btp"
description: "Steampipe plugin to query the account details of your SAP Business Technology Platform account."
og_description: "Query SAP BTP with SQL! Open source CLI. No DB required."
og_image: "/images/plugins/ajmaradiaga/sap-btp-graphic.png"
---

# SAP BTP + Steampipe

[SAP BTP](https://www.sap.com/products/technology-platform.html) brings together data and analytics, artificial intelligence, application development, automation, and integration in one, unified environment.

[Steampipe](https://steampipe.io) is an open source CLI to instantly query cloud APIs using SQL.

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

## Documentation

- **[Table definitions & examples →](/plugins/ajmaradiaga/btp/tables)**

## Quick start

### Install

Download and install the latest SAP BTP plugin:

```sh
steampipe plugin install btp
```

### Credentials

| Item        | Description                                                                                                                                                                                           |
| ----------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Credentials | You will need to create a service key for the Cloud Management Service. You can get the instructions on how to get an access token for the SAP Cloud Management Service APIs here: https://help.sap.com/docs/btp/sap-business-technology-platform/getting-access-token-for-sap-cloud-management-service-apis.                                                                |                                                               |
| Resolution  | 1. Credentials specified in environment variables, e.g., `BTP_CIS_ACCESS_TOKEN`.<br />2. Credentials explicitly set in a steampipe config file (`~/.steampipe/config/btp.spc`) |

### Configuration

Installing the latest namecheap plugin will create a config file (`~/.steampipe/config/btp.spc`) with a single connection named `btp`:

Configure your account details in `~/.steampipe/config/btp.spc`:

```hcl
connection "btp" {
  plugin = "btp"

  #########
  #  CIS  #
  #########

  # You will need to create a service key for the Cloud Management Service. You can get the instructions on how to get an access token for the SAP Cloud Management Service APIs here: https://help.sap.com/docs/btp/sap-business-technology-platform/getting-access-token-for-sap-cloud-management-service-apis.

  # URL of the Accounts Service. Required.
  # This can also be set via the `BTP_CIS_ACCOUNTS_SERVICE_URL` environment variable.
  cis_accounts_service_url = "https://accounts-service.cfapps.[region].hana.ondemand.com"
  
  # URL of the Entitlements Service. Required.
  # This can also be set via the `BTP_CIS_ENTITLEMENTS_SERVICE_URL` environment variable.
  cis_entitlements_service_url = "https://entitlements-service.cfapps.[region].hana.ondemand.com"

  # Access token to communicate with the Cloud Management Service APIs. Required.
  # This can also be set via the `BTP_CIS_ACCESS_TOKEN` environment variable.
  cis_access_token = "eyJhbGciOiDBNsO0JxFoAaodkDJ3Pmk7cFEsEr5ml5BwNWEafrEjy8Hsxt2mVACpD8B4AIPpRuMoGE71qXGoPcW0vCugceTwN4C3xM8qYmH7DLQ"
 
}
```

Alternatively, you can also use the environment variables specified below:

```sh
export BTP_CIS_ACCOUNTS_SERVICE_URL=https://accounts-service.cfapps.eu10.hana.ondemand.com
export BTP_CIS_ENTITLEMENTS_SERVICE_URL=https://entitlements-service.cfapps.eu10.hana.ondemand.com
export BTP_CIS_ACCESS_TOKEN=eyJhbGciOiDBNsO0JxFoAaodkDJ3Pmk7cFEsEr5ml5BwNWEafrEjy8Hsxt2mVACpD8B4AIPpRuMoGE71qXGoPcW0vCugceTwN4C3xM8qYmH7DLQ
```

## Get involved

- Open source: https://github.com/ajmaradiaga/steampipe-plugin-btp
