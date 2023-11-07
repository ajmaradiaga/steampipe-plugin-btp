SELECT 'Show me the money!' as kick_off;

---------------------------------
-- List all subaccounts in root
---------------------------------
SELECT GUID,
	DISPLAY_NAME,
	PARENT_GUID,
	PARENT_TYPE,
	SUBDOMAIN,
	CUSTOM_PROPERTIES
FROM BTP_ACCOUNTS_SUBACCOUNT;

---------------------------------
-- Filter subaccounts by label
---------------------------------

SELECT GUID,
	DISPLAY_NAME,
	PARENT_GUID,
	PARENT_TYPE,
	SUBDOMAIN,
	CUSTOM_PROPERTIES
FROM BTP_ACCOUNTS_SUBACCOUNT
CROSS JOIN JSONB_ARRAY_ELEMENTS(CUSTOM_PROPERTIES) CP
WHERE CP ->> 'key' = 'viewer' and CP ->> 'value' = 'DJ Adams'

------------------------------------------------------
-- What about if you have more than one account?
------------------------------------------------------

SELECT GUID,
	DISPLAY_NAME,
	PARENT_GUID,
	PARENT_TYPE,
	SUBDOMAIN,
	CUSTOM_PROPERTIES
FROM BTP_TRIAL.BTP_ACCOUNTS_SUBACCOUNT
