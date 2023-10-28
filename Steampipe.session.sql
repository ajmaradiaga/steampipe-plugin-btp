-- Global account details
SELECT guid, display_name, created_date, modified_date, entity_state, state_message, subdomain, contract_status, commercial_model, consumption_based, license_type, geo_access, renewal_date, _ctx
	FROM btp.btp_accounts_global_account;

-- List all subaccounts in root
SELECT guid, display_name, parent_guid, parent_type, subdomain
FROM btp_accounts_subaccounts

-- List all subaccounts in a directory
SELECT display_name,
    region,
    subdomain,
    beta_enabled
FROM btp_accounts_subaccounts
WHERE parent_guid = '00643708-5865-4e15-a0b4-d276c3877502'
ORDER BY region,
    display_name;

-- List all directories
SELECT DISTINCT parent_guid,
    parent_type
FROM btp_accounts_subaccounts
WHERE parent_type = 'PROJECT';

-- Count subaccounts by region
SELECT region,
    count(1)
FROM btp_accounts_subaccounts
GROUP BY region
ORDER BY count DESC;

-- Subaccount details with datacenter information
SELECT sa.guid subaccount_guid,
    sa.display_name subaccount_name,
    sa.subdomain subaccount_subdomain,
    dc.name dc_name,
    dc.display_name as dc_location,
    sa.region,
    environment,
    iaas_provider,
    supports_trial,
    saas_registry_service_url,
    domain,
    geo_access
FROM btp_accounts_subaccounts sa
    JOIN btp.btp_entitlements_alloweddatacenters dc ON sa.region = dc.region
ORDER BY region,
    subaccount_name;

SELECT DISTINCT business_category ->> 'id' bc_id
FROM       btp_entitlements_assignments bes;

SELECT DISTINCT business_category ->> 'id' bc_id
FROM       btp_entitlements_assignments bes
WHERE subaccount_guid = '700941ea-995f-4faa-95a4-10dcd448ad5c';

-- Assignments and service plans
SELECT     bes.display_name,
           service_plan ->> 'name' sp_displayName,
		   service_plan ->> 'amount' sp_amount,
		   service_plan ->> 'remainingAmount' sp_remaining_amount
FROM       btp_entitlements_assignments bes
CROSS JOIN jsonb_array_elements(service_plans) service_plan
WHERE business_category ->> 'id' = 'APPLICATION_DEVELOPMENT_AND_AUTOMATION'  
ORDER BY bes.display_name asc

-- Assignments and the data centers where they are available 
SELECT     bes.display_name,
           service_plan ->> 'name' sp_displayName,
		   data_centers ->> 'name' dc_name
FROM       btp_entitlements_assignments bes
CROSS JOIN jsonb_array_elements(service_plans) service_plan
CROSS JOIN jsonb_array_elements(service_plan -> 'dataCenters') data_centers
WHERE business_category ->> 'id' = 'INTEGRATION'  
ORDER BY bes.display_name asc

SELECT * 
FROM btp_entitlements_assignments_subaccount


SELECT title, link, author_name FROM rss_item where feed_link = 'https://blogs.sap.com/feed/';