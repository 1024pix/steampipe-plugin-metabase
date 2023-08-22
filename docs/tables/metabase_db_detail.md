# Table: metabase_db_detail

List all details of one database of Metabase. This is a map of string.

Informations depend of database vendor.

You must provide `db_id`. `db_id` is the database id from `metabase_db` table.

## Examples

### Get all details about one database

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
  metabase_db_detail,
  my_db
where
  db_id = my_db.id;
```
