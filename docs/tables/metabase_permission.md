# Table: metabase_permission

List all permissions of Metabase.

## Examples

### List all permissions

```sql
select
  *
from
  metabase_permission;
```

### Get permission of one group

```sql
select
  *
from
  metabase_permission_group
where
  group_id=15;
```
