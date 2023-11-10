# Table: btp_accounts_directory

Get the details of a directory in your global account.

## Examples

### Get directory

```sql
SELECT
    *
FROM
    BTP.BTP_ACCOUNTS_DIRECTORY
WHERE
    GUID = '35170704-2231-2v14-q4x1-c017f2525076'
```

### List all directories in a Global account

```sql
SELECT DISTINCT PARENT_GUID,
	PARENT_TYPE
FROM BTP_ACCOUNTS_SUBACCOUNT
WHERE PARENT_TYPE = 'PROJECT';
```
