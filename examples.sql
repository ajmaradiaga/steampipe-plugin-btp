----------------------------------------------
---- Steampipe query examples for SAP BTP ----
----------------------------------------------


-- SQL to get the list of subaccounts and their regions
select guid, display_name, region 
from btp_trial.btp_accounts_subaccounts 
where region = 'us10'

-- SQL to get the regions and total subaccounts in it
select region, count(1) as total
from btp_trial.btp_accounts_subaccounts
group by region 
order by total desc

-- SQL to get the list of subaccounts and information of the Datacenters where they are hosted
select sa.display_name, sa.region, sa.subdomain, dc.iaas_provider, dc.domain, dc.geo_access
from btp_trial.btp_accounts_subaccounts sa
join btp_trial.btp_entitlements_alloweddatacenters dc on sa.region = dc.region