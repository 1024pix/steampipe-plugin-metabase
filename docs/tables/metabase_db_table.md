# Table: metabase_db_table

List all tables of one databases of Metabase.

You must provide `db_id`. `db_id` is the database id from `metabase_db` table.

## Examples

### List all tables of one database

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
from
  metabase_db_table, metabase_db
where
  db_id = my_db.id;
```

### Search table with special warning

```sql
with my_db as (
  select
    id as my_db_id,
    name as my_db_name
  from
    metabase_db
  where name = 'Test'
)
select
  id,
  db_id,
  name,
  caveats
from
  metabase_db_table,
  my_db
where
  db_id = my_db.my_db_id AND
  caveats IS NOT NULL;
```

### Get one table of database

`id` is table id that you want.

```sql
select
  id,
  db_id,
  name
from
  metabase_db_table
where
  db_id = 15 and id = 209;
```

### List all tables show in getting started

```sql
select
  metabase_db_table.db_id as db_id,
  metabase_db_table.id as table_id,
  metabase_db_table.name as name
from
  metabase_db_table
inner join
  metabase_db ON metabase_db_table.db_id = metabase_db.id
where 
  metabase_db_table.show_in_getting_started = true
order by
  db_id asC, table_id asC;
```
