# Table: metabase_db

List all databases of Metabase.

This table need to be use first cause all tables started with `metabase_db_xxx` need to know the database id.

No field required.

## Examples

### List all databases

```sql
SELECT
  id,
  name,
  description,
  is_full_sync,
  is_sample,
  cache_field_values_schedule,
  metadata_sync_schedule,
  caveats,
  engine,
  created_at,
  updated_at,
  native_permissions,
  points_of_interest 
FROM
  metabase_db;
```

### Get one database by name

```sql
SELECT
  id
  name,
  description
FROM
  metabase_db
WHERE
  name = 'Test';
```

Return:
```
+------+------+-------------+
| id   | name | description |
+------+------+-------------+
| 7    | Test | <null>      |
+------+------+-------------+
```

### List all database use PostgreSQL

```sql
SELECT
  engine
FROM
  metabase_db
WHERE
   engine = 'postgres';
```

### List all databases are not fully syncrhonized after a Meatabase crash

```sql
SELECT
  name,
  is_full_sync
FROM
  metabase_db
WHERE
   is_full_sync = false;
```

Return:
```
+-----------------+--------------+
| name            | is_full_sync |
+-----------------+--------------+
| Test            | false        |
+-----------------+--------------+
```

### Ensure no database can be use with native SQL

For this, we list all database with `native_permissions` set `write`:
```sql
SELECT
  name
FROM
  metabase_db
WHERE
   native_permissions = 'write';
```
