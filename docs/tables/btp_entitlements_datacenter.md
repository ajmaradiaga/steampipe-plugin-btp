# Table: btp_entitlements_datacenter

Get the details of all the data centers where a subaccount can be created in the SAP BTP global account.

## Examples

### List all the allowed data centers

```sql
select
  display_name,
  region,
  environment,
  iaas_provider,
  supports_trial,
  domain,
  is_main_data_center
from
  btp_entitlements_datacenter;
```

### List all the Cloud Foundry data centers

```sql
select
  display_name,
  region,
  environment,
  iaas_provider,
  supports_trial,
  domain,
  is_main_data_center
from
  btp_entitlements_datacenter 
where
  environment = 'cloudfoundry';
```

### List all the satellite data centers

```sql
select
  display_name,
  region,
  environment,
  iaas_provider,
  supports_trial,
  domain
from
  btp_entitlements_datacenter 
where
  is_main_data_center = False;
```

### Subaccount details with datacenter information

```sql
select
  sa.guid subaccount_guid,
  sa.display_name subaccount_name,
  sa.subdomain subaccount_subdomain,
  dc.name dc_name,
  dc.display_name as dc_location,
  sa.region,
  environment,
  dc.iaas_provider,
  supports_trial,
  saas_registry_service_url,
  domain,
  geo_access 
from
  btp_accounts_subaccount sa 
  join
    btp_entitlements_datacenter dc 
    on sa.region = dc.region 
order by
  region,
  subaccount_name;
```
