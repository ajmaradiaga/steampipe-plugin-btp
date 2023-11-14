# Table: btp_accounts_directory

Get the details of a directory in your global account.

## Examples

### Get directory

```sql
select
  * 
from
  btp.btp_accounts_directory 
where
  guid = '35170704-2231-2v14-q4x1-c017f2525076';
```

### List all directories in a Global account

```sql
select distinct
  parent_guid,
  parent_type 
from
  btp_accounts_subaccount 
where
  parent_type = 'PROJECT';
```
