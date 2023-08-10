# Table: metabase_db_detail

List all details of one database of Metabase. This is a map of string.

Informations depend of database vendor.

You must provide `db_id`. `db_id` is the database id from `metabase_db` table.

## Examples

### Get all details about one database

In case of we use PostgreSQL database, we want kown all list of features are activated on this PostgreSQL database.

`feature` column contain the list of activated features.

```sql
WITH my_db AS (
  SELECT
    id,
    name
  FROM
    metabase_db
  WHERE name = 'Test'
)
SELECT
  feature
FROM
  metabase_db_detail,
  my_db
WHERE
  db_id = my_db.id;
```

Return:
```
+-------+-----------------------------+----------------------------------------------+--------------------------------+
| db_id | key                         | value                                        | _ctx                           |
+-------+-----------------------------+----------------------------------------------+--------------------------------+
| 7     | ssl-client-cert-source      | <nil>                                        | {"connection_name":"metabase"} |
| 7     | schema-filters-type         | all                                          | {"connection_name":"metabase"} |
| 7     | ssl-mode                    | require                                      | {"connection_name":"metabase"} |
| 7     | ssl-key-creator-id          | <nil>                                        | {"connection_name":"metabase"} |
| 7     | ssl-key-password-creator-id | <nil>                                        | {"connection_name":"metabase"} |
| 7     | ssl                         | true                                         | {"connection_name":"metabase"} |
| 7     | ssl-key-source              | <nil>                                        | {"connection_name":"metabase"} |
| 7     | port                        | 99999                                        | {"connection_name":"metabase"} |
| 7     | ssl-use-client-auth         | false                                        | {"connection_name":"metabase"} |
| 7     | ssl-root-cert-creator-id    | <nil>                                        | {"connection_name":"metabase"} |
| 7     | ssl-client-cert-creator-id  | <nil>                                        | {"connection_name":"metabase"} |
| 7     | ssl-root-cert-source        | <nil>                                        | {"connection_name":"metabase"} |
| 7     | user                        | this_is_user                                 | {"connection_name":"metabase"} |
| 7     | password                    | **MetabasePass**                             | {"connection_name":"metabase"} |
| 7     | host                        | my_dp.postgresql.dbs.test.com                | {"connection_name":"metabase"} |
| 7     | ssl-key-password-source     | <nil>                                        | {"connection_name":"metabase"} |
| 7     | dbname                      | this_is_my_db                                | {"connection_name":"metabase"} |
| 7     | tunnel-enabled              | false                                        | {"connection_name":"metabase"} |
| 7     | advanced-options            | false                                        | {"connection_name":"metabase"} |
+-------+-----------------------------+----------------------------------------------+--------------------------------+
```
