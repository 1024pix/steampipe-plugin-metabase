# Table: metabase_permission_data

List data permission of Metabase for getting data from database.

See [Guide to data permissions](https://www.metabase.com/learn/permissions/data-permissions) of Metabase documentation.

No field required.

`db_id` is the database id from `metabase_db` table.

`group_id` is Metabase group that you can find in `metabase_permission_group` table.

`table_id` is the table id from `metabase_db_table` table.

## Examples

### List all permissions

```sql
select
  group_id,
  db_id,
  schema_name,
  table_id,
  level_access
from
  metabase_permission_data;
```

### Seach all group in Metabase that have all level access on database that have 'Granular access'

```sql
with group_with_all_level_access as (
  select
    group_id,
    db_id,
    table_id
  from
    metabase_permission_data
  where
    level_access = 'all'
)
select
  metabase_permission_group.name as group_name,
  metabase_permission_group.id as group_id,
  metabase_db_table.id as table_id,
  metabase_db_table.name as table_name,
  metabase_db.id as db_id,
  metabase_db.name as db_name
from
  metabase_db_table
inner join
  group_with_all_level_access ON metabase_db_table.db_id = group_with_all_level_access.db_id AND metabase_db_table.id = group_with_all_level_access.table_id
inner join
  metabase_permission_group ON metabase_permission_group.id = group_with_all_level_access.group_id
inner join
  metabase_db ON metabase_db.id = group_with_all_level_access.db_id
```

