# Table: btp_entitlements_datacenter

Get the details of all the data centers where a subaccount can be created in the SAP BTP global account.

## Examples

### List all the allowed data centers

```sql
SELECT * 
FROM BTP.BTP_ENTITLEMENTS_DATACENTER;
```

### List all the Cloud Foundry data centers

```sql
SELECT * 
FROM BTP.BTP_ENTITLEMENTS_DATACENTER
WHERE environment = 'cloudfoundry';
```

### List all the satellite data centers

```sql
SELECT * 
FROM BTP.BTP_ENTITLEMENTS_DATACENTER
WHERE is_main_data_center = False;
```

### Subaccount details with datacenter information

```sql
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
```
