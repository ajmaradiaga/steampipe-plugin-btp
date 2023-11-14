# Table: btp_accounts_subaccount

Get the details of subaccounts in the SAP BTP global account.

## Examples

### List all subaccounts in root

```sql
select guid,
	display_name,
	parent_guid,
	parent_type,
	subdomain,
	custom_properties
from btp_accounts_subaccount;
```

### List all subaccounts in a directory

```sql
select display_name,
	region,
	subdomain,
	beta_enabled
from btp_accounts_subaccount
where parent_guid = '00643708-5865-4e15-a0b4-d276c3877502'
order by region,
	display_name;
```

### List all directories

```sql
select distinct parent_guid,
	parent_type
from btp_accounts_subaccount
where parent_type = 'PROJECT';
```

### Count subaccounts by region

```sql
select region,
	count(1)
from btp_accounts_subaccount
group by region
order by count desc;
```

### Subaccount details with datacenter information

```sql
select sa.guid subaccount_guid,
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
from btp_accounts_subaccount sa
join btp.btp_entitlements_datacenter dc on sa.region = dc.region
order by region,
	subaccount_name;
```
