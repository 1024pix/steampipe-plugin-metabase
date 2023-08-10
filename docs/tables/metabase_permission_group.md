# Table: metabase_permission_group

List all groups in Metabase.

Please read [Permissions introduction](https://www.metabase.com/docs/latest/permissions/introduction) to understand how it works.

## Examples

### List all groups

```sql
SELECT
  id,
  name,
  member_count
FROM
  metabase_permission_group;
```

Return:
```
+----+----------------+--------------+--------------------------------+
| id | name           | member_count | _ctx                           |
+----+----------------+--------------+--------------------------------+
| 1  | All Users      | 35           | {"connection_name":"metabase"} |
| 8  | Administrators | 1            | {"connection_name":"metabase"} |
| 4  | Test           | 1            | {"connection_name":"metabase"} |
+----+----------------+--------------+--------------------------------+
```

### Get one group

```sql
SELECT
  name
FROM
  metabase_permission_group
WHERE
  id = 15;
```

### Search group that can download all data and display name of database and group name

```sql
WITH group_with_download_full AS (
  SELECT
    id,
    name
  FROM
    metabase_permission
  WHERE
    download_native = 'full'
)
SELECT
  metabase_permission_group.id as group_id,
  metabase_permission_group.name as group_name,
  metabase_db.id as db_id,
  metabase_db.name as db_name
FROM 
  metabase_permission_group
INNER JOIN group_with_download_full ON group_with_download_full.group_id = metabase_permission_group.id
INNER JOIN metabase_db ON metabase_db.id = group_with_download_full.db_id
ORDER BY
  group_id
```
