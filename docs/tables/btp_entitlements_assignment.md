# Table: btp_entitlements_assignment

Get the details of all the service assignments available to the SAP BTP global account.

## Examples

### Get the business category of all services

```sql
select 
  distinct
  business_category ->> 'id' bc_id 
from
  btp_entitlements_assignment;
```

### Nested JSON structures in the Entitlements API

```sql
select
  bes.display_name,
  service_plans 
from
  btp_entitlements_assignment;
```

### Assignments and quota for a particular business category

```sql
select
  bes.display_name,
  service_plan ->> 'name' sp_displayname,
  service_plan ->> 'amount' sp_amount,
  service_plan ->> 'remainingamount' sp_remaining_amount 
from
  btp_entitlements_assignment bes 
  cross join
    jsonb_array_elements(service_plans) service_plan 
where
  business_category ->> 'id' = 'AI' 
order by
  bes.display_name asc;
```

### Assignments and the data centers where they are available

```sql
select
  bes.display_name,
  service_plan ->> 'name' sp_displayname,
  data_centers ->> 'name' dc_name 
from
  btp_entitlements_assignment bes 
  cross join
    jsonb_array_elements(service_plans) service_plan 
  cross join
    jsonb_array_elements(service_plan -> 'datacenters') data_centers 
where
  business_category ->> 'id' = 'INTEGRATION' 
  and data_centers ->> 'name' = 'cf-eu10' 
order by
  bes.display_name asc;
```
