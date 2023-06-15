# Table: metabase_db

List all databases of Metabase.

## Examples

### List all databases

```sql
select
  name,
  description
from
  metabase_db;
```

### Get one database

```sql
select
  name,
  description
from
  metabase_db
where
  id=15;
```
