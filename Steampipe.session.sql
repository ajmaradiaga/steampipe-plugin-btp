---------------------------
-- Global account details
---------------------------

SELECT GUID,
	DISPLAY_NAME,
	CREATED_DATE,
	MODIFIED_DATE,
	ENTITY_STATE,
	STATE_MESSAGE,
	SUBDOMAIN,
	CONTRACT_STATUS,
	COMMERCIAL_MODEL,
	CONSUMPTION_BASED
FROM BTP.BTP_ACCOUNTS_GLOBAL_ACCOUNT;

---------------------------------
-- List all subaccounts in root
---------------------------------

SELECT GUID,
	DISPLAY_NAME,
	PARENT_GUID,
	PARENT_TYPE,
	SUBDOMAIN,
	CUSTOM_PROPERTIES
FROM BTP_ACCOUNTS_SUBACCOUNT 

---------------------------------------
-- List all subaccounts in a directory
---------------------------------------

SELECT DISPLAY_NAME,
	REGION,
	SUBDOMAIN,
	BETA_ENABLED
FROM BTP_ACCOUNTS_SUBACCOUNT
WHERE PARENT_GUID = '00643708-5865-4e15-a0b4-d276c3877502'
ORDER BY REGION,
	DISPLAY_NAME;

---------------------------------
-- List all directories
---------------------------------

SELECT DISTINCT PARENT_GUID,
	PARENT_TYPE
FROM BTP_ACCOUNTS_SUBACCOUNT
WHERE PARENT_TYPE = 'PROJECT';

---------------------------------
-- Count subaccounts by region
---------------------------------

SELECT REGION,
	COUNT(1)
FROM BTP_ACCOUNTS_SUBACCOUNT
GROUP BY REGION
ORDER BY COUNT DESC;

---------------------------------------------------
-- Subaccount details with datacenter information
---------------------------------------------------

SELECT SA.GUID SUBACCOUNT_GUID,
	SA.DISPLAY_NAME SUBACCOUNT_NAME,
	SA.SUBDOMAIN SUBACCOUNT_SUBDOMAIN,
	DC.NAME DC_NAME,
	DC.DISPLAY_NAME AS DC_LOCATION,
	SA.REGION,
	ENVIRONMENT,
	DC.IAAS_PROVIDER,
	SUPPORTS_TRIAL,
	SAAS_REGISTRY_SERVICE_URL,
	DOMAIN,
	GEO_ACCESS
FROM BTP_ACCOUNTS_SUBACCOUNT SA
JOIN BTP.BTP_ENTITLEMENTS_DATACENTER DC ON SA.REGION = DC.REGION
ORDER BY REGION,
	SUBACCOUNT_NAME;

----------------------------------------------
-- Get the business category of all services
----------------------------------------------

SELECT DISTINCT BUSINESS_CATEGORY ->> 'id' BC_ID
FROM BTP_ENTITLEMENTS_ASSIGNMENT BES;

---------------------------------------------------
-- Nested JSON structures in the Entitlements API
---------------------------------------------------

SELECT BES.DISPLAY_NAME,
	SERVICE_PLANS
FROM BTP_ENTITLEMENTS_ASSIGNMENT BES 

-------------------------------------------------------------
-- Assignments and quota for a particular business category
-------------------------------------------------------------

SELECT BES.DISPLAY_NAME,
	SERVICE_PLAN ->> 'name' SP_DISPLAYNAME,
	SERVICE_PLAN ->> 'amount' SP_AMOUNT,
	SERVICE_PLAN ->> 'remainingAmount' SP_REMAINING_AMOUNT
FROM BTP_ENTITLEMENTS_ASSIGNMENT BES
CROSS JOIN JSONB_ARRAY_ELEMENTS(SERVICE_PLANS) SERVICE_PLAN
WHERE BUSINESS_CATEGORY ->> 'id' = 'AI'
ORDER BY BES.DISPLAY_NAME ASC;

------------------------------------------------------------------
-- Assignments and the data centers where they are available
------------------------------------------------------------------

SELECT BES.DISPLAY_NAME,
	SERVICE_PLAN ->> 'name' SP_DISPLAYNAME,
	DATA_CENTERS ->> 'name' DC_NAME
FROM BTP_ENTITLEMENTS_ASSIGNMENT BES
CROSS JOIN JSONB_ARRAY_ELEMENTS(SERVICE_PLANS) SERVICE_PLAN
CROSS JOIN JSONB_ARRAY_ELEMENTS(SERVICE_PLAN -> 'dataCenters') DATA_CENTERS
WHERE BUSINESS_CATEGORY ->> 'id' = 'INTEGRATION'
	AND DATA_CENTERS ->> 'name' = 'cf-eu10'
ORDER BY BES.DISPLAY_NAME ASC


------------------------------------------------------------------
-- Have multiple BTP Global accounts?
------------------------------------------------------------------

SELECT glob.display_name "Global Account", sub.REGION,
	COUNT(1)
FROM BTP.BTP_ACCOUNTS_SUBACCOUNT sub
JOIN BTP.BTP_ACCOUNTS_GLOBAL_ACCOUNT glob ON sub.global_account_guid = glob.guid
GROUP BY glob.display_name, sub.REGION
UNION
SELECT glob.display_name "Global Account", sub.REGION,
	COUNT(1)
FROM BTP_TRIAL.BTP_ACCOUNTS_SUBACCOUNT sub
JOIN BTP_TRIAL.BTP_ACCOUNTS_GLOBAL_ACCOUNT glob ON sub.global_account_guid = glob.guid
GROUP BY glob.display_name, sub.REGION
ORDER BY COUNT DESC, region asc;

