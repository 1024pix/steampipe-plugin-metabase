# Table: metabase_db_table

List all tables of one databases of Metabase.

## Examples

### List all tables of one database

```sql
select
  *
from
  metabase_db_table
where
  db_id=15;
```

### Get one database

```sql
select
  *
from
  metabase_db_table
where
  db_id=15 and id=7;
```
