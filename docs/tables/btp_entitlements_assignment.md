# Table: btp_entitlements_assignment

Get the details of all the service assignments available to the SAP BTP global account.

## Examples

### Get the business category of all services

```sql
SELECT DISTINCT BUSINESS_CATEGORY ->> 'id' BC_ID
FROM BTP_ENTITLEMENTS_ASSIGNMENT BES;
```

### Nested JSON structures in the Entitlements API

```sql
SELECT BES.DISPLAY_NAME,
	SERVICE_PLANS
FROM BTP_ENTITLEMENTS_ASSIGNMENT BES 
```

### Assignments and quota for a particular business category

```sql
SELECT BES.DISPLAY_NAME,
	SERVICE_PLAN ->> 'name' SP_DISPLAYNAME,
	SERVICE_PLAN ->> 'amount' SP_AMOUNT,
	SERVICE_PLAN ->> 'remainingAmount' SP_REMAINING_AMOUNT
FROM BTP_ENTITLEMENTS_ASSIGNMENT BES
CROSS JOIN JSONB_ARRAY_ELEMENTS(SERVICE_PLANS) SERVICE_PLAN
WHERE BUSINESS_CATEGORY ->> 'id' = 'AI'
ORDER BY BES.DISPLAY_NAME ASC;
```

### Assignments and the data centers where they are available

```sql
SELECT BES.DISPLAY_NAME,
	SERVICE_PLAN ->> 'name' SP_DISPLAYNAME,
	DATA_CENTERS ->> 'name' DC_NAME
FROM BTP_ENTITLEMENTS_ASSIGNMENT BES
CROSS JOIN JSONB_ARRAY_ELEMENTS(SERVICE_PLANS) SERVICE_PLAN
CROSS JOIN JSONB_ARRAY_ELEMENTS(SERVICE_PLAN -> 'dataCenters') DATA_CENTERS
WHERE BUSINESS_CATEGORY ->> 'id' = 'INTEGRATION'
	AND DATA_CENTERS ->> 'name' = 'cf-eu10'
ORDER BY BES.DISPLAY_NAME ASC
```
