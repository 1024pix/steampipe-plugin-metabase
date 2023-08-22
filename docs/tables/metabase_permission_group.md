# Table: metabase_permission_group

List all groups in Metabase.

Please read [Permissions introduction](https://www.metabase.com/docs/latest/permissions/introduction) to understand how it works.

## Examples

### List all groups

```sql
select
  id,
  name,
  member_count
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
  id = 15;
```

### List groups that can download all data and display name of database and group name

```sql
with group_with_download_full as (
  select
    id,
    name
  from
    metabase_permission
  where
    download_native = 'full'
)
select
  metabase_permission_group.id as group_id,
  metabase_permission_group.name as group_name,
  metabase_db.id as db_id,
  metabase_db.name as db_name
from 
  metabase_permission_group
inner join group_with_download_full on group_with_download_full.group_id = metabase_permission_group.id
inner join metabase_db on metabase_db.id = group_with_download_full.db_id
order by
  group_id
```
