# Table: metabase_db_feature

List all features of one database of Metabase.

You must provide `db_id`. `db_id` is the database id from `metabase_db` table.

## Examples

### Get all enabled features of database

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
  metabase_db_feature,
  my_db
WHERE
  db_id = my_db.id;
```

Return:
```
+----------------------------------------+--------------------------------+
| feature                                | _ctx                           |
+----------------------------------------+--------------------------------+
| actions                                | {"connection_name":"metabase"} |
| nested-queries                         | {"connection_name":"metabase"} |
...
| case-sensitivity-string-filter-options | {"connection_name":"metabase"} |
| inner-join                             | {"connection_name":"metabase"} |
+-------+----------------------------------------+--------------------------------+
```

### Check if regex is enabled on PostgreSQL database

```sql
SELECT
  count(*)
FROM
  metabase_db_feature
WHERE
  db_id = 15 and
  feature = 'regex';
```

Return:
```
+-------+
| count |
+-------+
| 1     |
+-------+
```
