# Table: metabase_permission

List all permissions from Metabase.

To understand how permision works, you can read official documentation of Metabase [](https://www.metabase.com/docs/latest/permissions/data).

`db_id` is the database id from `metabase_db` table.

`group_id` is Metabase group that you can find in `metabase_permission_group` table.

## Examples

### List all permissions

```sql
SELECT
  group_id,
  db_id,
  download_native,
  download_schema,
  data_native,
  data_schema
FROM
  metabase_permission;
```

Return:
```
+----------+-------+-----------------+-----------------+-------------+-------------+--------------------------------+
| group_id | db_id | download_native | download_schema | data_native | data_schema | _ctx                           |
+----------+-------+-----------------+-----------------+-------------+-------------+--------------------------------+
| 4        | 1     | <null>          | <null>          | <null>      | limited     | {"connection_name":"metabase"} |
| 2        | 5     | full            | full            | write       | all         | {"connection_name":"metabase"} |
...
```

### Display group name with permission

```sql
SELECT
  group_id,
  name,
  db_id,
  download_native,
  download_schema,
  data_native,
  data_schema
FROM
  metabase_permission
INNER JOIN metabase_permission_group ON metabase_permission_group.id = metabase_permission.group_id
ORDER BY
  group_id;
```
