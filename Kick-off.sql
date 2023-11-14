select
  'Show me the money!' as kick_off;
  
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

---------------------------------
-- Filter subaccounts by label
---------------------------------
select
  guid,
  display_name,
  parent_guid,
  parent_type,
  subdomain,
  custom_properties 
from
  btp_accounts_subaccount
  cross join
    jsonb_array_elements(custom_properties) cp 
where
  cp ->> 'key' = 'label1' 
  and cp ->> 'value' = 'value1';
  
------------------------------------------------------
-- What about if you have more than one account?
------------------------------------------------------
select
  guid,
  display_name,
  parent_guid,
  parent_type,
  subdomain,
  custom_properties 
from
  btp_trial.btp_accounts_subaccount;