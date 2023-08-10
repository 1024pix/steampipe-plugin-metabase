# Table: metabase_db_table

List all tables of one databases of Metabase.

You must provide `db_id`. `db_id` is the database id from `metabase_db` table.

## Examples

### List all tables of one database

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
  id,
  db_id,
  schema,
  show_in_getting_started,
  name,
  description,
  caveats,
  created_at,
  updated_at,
  visibility_type,
  display_name,
  points_of_interest
FROM
  metabase_db_table, metabase_db
WHERE
  db_id = my_db.id;
```

Return:
```
+-----+-------+--------+-------------------------+----------+-------------+---------+---------------------------+---------------------------+-----------------+-----------------------------+--------------------+--------------------------------+
| id  | db_id | schema | show_in_getting_started | name     | description | caveats | created_at                | updated_at                | visibility_type | display_name                | points_of_interest | _ctx                           |
+-----+-------+--------+-------------------------+----------+-------------+---------+---------------------------+---------------------------+-----------------+-----------------------------+--------------------+--------------------------------+
| 209 | 15    | public | false                   | table1   | <null>      | <null>  | 2023-05-31T11:00:48+02:00 | 2023-05-31T11:00:53+02:00 | <null>          | This is my table            | <null>             | {"connection_name":"metabase"} |
| 223 | 15    | public | false                   | table2   | <null>      | <null>  | 2023-05-31T11:00:49+02:00 | 2023-05-31T11:00:52+02:00 | <null>          | A amazing table             | <null>             | {"connection_name":"metabase"} |
| 201 | 15    | public | false                   | table3   | <null>      | <null>  | 2023-05-31T11:00:48+02:00 | 2023-05-31T11:00:52+02:00 | <null>          | Another one bites the dust  | <null>             | {"connection_name":"metabase"} |
...
```

### Search table with special warning

```sql
WITH my_db AS (
  SELECT
    id as my_db_id,
    name as my_db_name
  FROM
    metabase_db
  WHERE name = 'Test'
)
SELECT
  id,
  db_id,
  name,
  caveats
FROM
  metabase_db_table,
  my_db
WHERE
  db_id = my_db.my_db_id AND
  caveats IS NOT NULL;
```

Return:
```
+----+-------+------+---------+
| id | db_id | name | caveats |
+----+-------+------+---------+
+----+-------+------+---------+
```

### Get one table of database

`id` is table id that you want.

```sql
SELECT
  id,
  db_id,
  name
FROM
  metabase_db_table
WHERE
  db_id = 15 and id = 209;
```

### List all tables show in getting started

```sql
SELECT
  metabase_db_table.db_id as db_id,
  metabase_db_table.id as table_id,
  metabase_db_table.name as name
FROM
  metabase_db_table
INNER JOIN
  metabase_db ON metabase_db_table.db_id = metabase_db.id
WHERE 
  metabase_db_table.show_in_getting_started = true
ORDER BY
  db_id ASC, table_id ASC;
```
