# Table: btp_accounts_global_account

Get the details of an SAP BTP global account.

## Examples

### Get global account details

```sql
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
```
