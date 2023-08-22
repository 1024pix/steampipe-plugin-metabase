# Table: metabase_db_feature

List all features of one database of Metabase.

You must provide `db_id`. `db_id` is the database id from `metabase_db` table.

## Examples

### Get all enabled features of database

In case of we use PostgreSQL database, we want kown all list of features are activated on this PostgreSQL database.

`feature` column contain the list of activated features.

```sql
with my_db as (
  select
    id,
    name
  from
    metabase_db
  where name = 'Test'
)
select
  feature
from
  metabase_db_feature,
  my_db
where
  db_id = my_db.id;
```

### Check if regex is enabled on PostgreSQL database

```sql
select
  count(*)
from
  metabase_db_feature
where
  db_id = 15 and
  feature = 'regex';
```

