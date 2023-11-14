# Table: btp_accounts_global_account

Get the details of an SAP BTP global account.

## Examples

### Get global account details

```sql
select guid,
	display_name,
	created_date,
	modified_date,
	entity_state,
	state_message,
	subdomain,
	contract_status,
	commercial_model,
	consumption_based
from btp.btp_accounts_global_account;
```
