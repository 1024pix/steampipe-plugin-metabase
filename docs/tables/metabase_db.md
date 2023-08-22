# Table: metabase_db

List all databases of Metabase.

This table need to be use first cause all tables started with `metabase_db_xxx` need to know the database id.

No field required.

## Examples

### List all databases

```sql
select
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
from
  metabase_db;
```

### Get one database by name

```sql
select
  id
  name,
  description
from
  metabase_db
where
  name = 'Test';
```

### List all database use PostgreSQL

```sql
select
  engine
from
  metabase_db
where
   engine = 'postgres';
```

### List all databases are not fully synchronized after a Metabase crash

```sql
select
  name,
  is_full_sync
from
  metabase_db
where
   is_full_sync = false;
```

### Ensure no database can be use with native SQL

For this, we list all database with `native_permissions` set `write`:
```sql
select
  name
from
  metabase_db
where
   native_permissions = 'write';
```
