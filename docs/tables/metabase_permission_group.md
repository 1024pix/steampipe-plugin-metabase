# Table: metabase_permission_group

List all groups of Metabase.

## Examples

### List all groups

```sql
select
  name,
from
  metabase_permission_group;
```

### Get one group

```sql
select
  name
from
  metabase_permission_group
where
  id=15;
```
