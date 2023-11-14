---------------------------
-- Global account details
---------------------------
select
  guid,
  display_name,
  created_date,
  modified_date,
  entity_state,
  state_message,
  subdomain,
  contract_status,
  commercial_model,
  consumption_based 
from
  btp.btp_accounts_global_account;

---------------------------------
-- List all subaccounts in root
---------------------------------
select
  guid,
  display_name,
  parent_guid,
  parent_type,
  subdomain,
  custom_properties 
from
  btp_accounts_subaccount;
  
---------------------------------------
-- List all subaccounts in a directory
---------------------------------------
select
  display_name,
  region,
  subdomain,
  beta_enabled 
from
  btp_accounts_subaccount 
where
  parent_guid = '00643708-5865-4e15-a0b4-d276c3877502' 
order by
  region,
  display_name;

---------------------------------
-- List all directories
---------------------------------
select distinct
  parent_guid,
  parent_type 
from
  btp_accounts_subaccount 
where
  parent_type = 'PROJECT';
  
---------------------------------
-- Count subaccounts by region
---------------------------------
select
  region,
  count(1) 
from
  btp_accounts_subaccount 
group by
  region 
order by
  count desc;

---------------------------------------------------
-- Subaccount details with datacenter information
---------------------------------------------------
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
    btp.btp_entitlements_datacenter dc 
    on sa.region = dc.region 
order by
  region,
  subaccount_name;
  
----------------------------------------------
-- Get the business category of all services
----------------------------------------------
select distinct
  business_category ->> 'id' bc_id 
from
  btp_entitlements_assignment bes;
  
---------------------------------------------------
-- Nested JSON structures in the Entitlements API
---------------------------------------------------
select
  bes.display_name,
  service_plans 
from
  btp_entitlements_assignment bes;
  
-------------------------------------------------------------
-- Assignments and quota for a particular business category
-------------------------------------------------------------
select
  bes.display_name,
  service_plan ->> 'name' sp_displayname,
  service_plan ->> 'amount' sp_amount,
  service_plan ->> 'remainingAmount' sp_remaining_amount 
from
  btp_entitlements_assignment bes 
  join
    jsonb_array_elements(service_plans) service_plan 
    on true 
where
  business_category ->> 'id' = 'INTEGRATION' 
order by
  bes.display_name asc;

------------------------------------------------------------------
-- Assignments and the data centers where they are available
------------------------------------------------------------------
select
  bes.name,
  bes.display_name,
  service_plan ->> 'name' sp_displayname,
  data_centers ->> 'name' dc_name 
from
  btp_entitlements_assignment bes 
  join
    jsonb_array_elements(service_plans) service_plan 
    on true 
  join
    jsonb_array_elements(service_plan -> 'dataCenters') data_centers 
    on true 
where
  business_category ->> 'id' = 'AI' 
  and data_centers ->> 'name' = 'cf-eu10' 
order by
  bes.display_name asc;

select
  bes.name,
  bes.display_name,
  service_plan ->> 'name' sp_displayname,
  data_centers ->> 'name' dc_name 
from
  btp_entitlements_assignment bes 
  cross join
    jsonb_array_elements(service_plans) service_plan 
  cross join
    jsonb_array_elements(service_plan -> 'dataCenters') data_centers 
where
  business_category ->> 'id' = 'AI' 
  and data_centers ->> 'name' = 'cf-eu10' 
order by
  bes.display_name asc; 		
	
------------------------------------------------------------------
-- Have multiple BTP Global accounts?
------------------------------------------------------------------
select
  glob.display_name "global account",
  sub.region,
  count(1) 
from
  btp.btp_accounts_subaccount sub 
  join
	btp.btp_accounts_global_account glob 
	on sub.global_account_guid = glob.guid 
group by
  glob.display_name,
  sub.region 
union
select
  glob.display_name "global account",
  sub.region,
  count(1) 
from
  btp_trial.btp_accounts_subaccount sub 
  join
	btp_trial.btp_accounts_global_account glob 
	on sub.global_account_guid = glob.guid 
group by
  glob.display_name,
  sub.region 
order by
  count desc,
  region asc;
