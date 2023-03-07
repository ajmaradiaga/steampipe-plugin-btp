# Steampipe plugin for SAP BTP

> ⚠️ This plugin is a proof of concept and is in no way supported by SAP.

Using this plugin will allow you to use SQL to query your SAP BTP account. As this is a PoC, the plugin only supports a subset of the APIs available in [SAP BTP Core Services](https://api.sap.com/package/SAPCloudPlatformCoreServices/rest), e.g. querying Global Accounts, Subaccounts and Directories. In the future, I plan to expand its capabilities and include the Entitlements, Events, Provisioning, and Resource consumption services that are part of the [SAP BTP Core Services](https://api.sap.com/package/SAPCloudPlatformCoreServices/rest).

![Steampipe plugin for SAP BTP](./assets/demo.gif)

```sql
-- Apply filter in your queries, e.g. retrieve all subaccounts in a specific region.
select guid, display_name, region 
from btp_trial.btp_accounts_subaccounts 
where region = 'ap21'
```
```txt
+--------------------------------------+------------------+--------+
| guid                                 | display_name     | region |
+--------------------------------------+------------------+--------+
| 4e923803-c32b-4cb6-b008-738892b7cd8f | subaccount_under | ap21   |
| a6285f10-2dae-4318-b2c6-2075f2fcc3f3 | subaccount_2     | ap21   |
+--------------------------------------+------------------+--------+
```
```sql
-- Count all subaccounts per region
select region, count(1) as total
from btp_trial.btp_accounts_subaccounts
group by region 
order by total desc
```
```txt
+--------+-------+
| region | total |
+--------+-------+
| ap21   | 2     |
| us10   | 1     |
+--------+-------+
```

To query our SAP BTP account we need to install Steampipe and the SAP BTP plugin locally. Also, the plugin requires an access token to consume the SAP BTP Core Services API.

## Install Steampipe

```bash
$ brew install turbot/tap/steampipe
$ steampipe plugin install steampipe

# Validate installation is working fine
$ steampipe query "select name from steampipe_registry_plugin;"

```

## SAP BTP Access Token

You can set the access token in the plugin config file, e.g. `~/.steampipe/config/btp.spc` or set it as an environment variable. Below I'm using the great [`generate-password-grant-type` script](https://github.com/SAP-samples/cloud-btp-cli-api-codejam/blob/main/scripts/generate-password-grant-type) (created by my friend and colleague [@qmacro](https://github.com/qmacro)) to retrieve an access token. The script expects a Cloud Management Service instance service key. Meaning, you will need to create an instance of this service and then a service key in the SAP BTP Cockpit. 

```bash
export BTP_ACCESS_TOKEN=$(./generate-password-grant-type cis-central-trial-sk | jq -r .access_token)
```

## Install the SAP BTP plugin

```bash
# Go to the directory where you've cloned this repo
$ cd repos/steampipe-plugin-btp

# Copy the sample configuration to the steampipe config folder
$ cp btp/config/btp.spc.sample ~/.steampipe/config/btp.spc

# Runs the command in Makefile and places the plugin in the 
# steampipe plugins folder, e.g. ~/.steampipe/plugins/local/btp
$ make

# Query your SAP BTP account
$ steampipe query "select * from btp_accounts_global_account"

```

Below, the tree structure of .steampipe after you've installed the plugin and copied the sample configuration.

```
~/.steampipe/
├── config
│   ├── ...
│   ├── btp.spc
├── db
│   ......
└── plugins
    ├── hub.steampipe.io
    │   └── plugins
    │       └── turbot
    ├── local
    │   └── btp
    │       └── btp.plugin
```

## Troubleshooting

If facing any issues when running the project, we can set the steampipe log level and check the logs located in `~/.steampipe/logs`.

```bash
# Set steampipe log level
$ export STEAMPIPE_LOG_LEVEL=LEVEL=WARN

$ tail -f ~/.steampipe/logs/plugin-2023-03-07.log
```

