# Table: metabase_permission

List all permissions from Metabase.

To understand how permision works, you can read official documentation of [Metabase](https://www.metabase.com/docs/latest/permissions/data).

`db_id` is the database id from `metabase_db` table.

`group_id` is Metabase group that you can find in `metabase_permission_group` table.

## Examples

### List permissions

```sql
select
  group_id,
  db_id,
  download_native,
  download_schema,
  data_native,
  data_schema
from
  metabase_permission;
```

### List groups with permission

```sql
select
  group_id,
  name,
  db_id,
  download_native,
  download_schema,
  data_native,
  data_schema
from
  metabase_permission
inner join metabase_permission_group on metabase_permission_group.id = metabase_permission.group_id
order by
  group_id;
```
