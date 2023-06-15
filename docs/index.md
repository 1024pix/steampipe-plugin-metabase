---
organization: 1024pix
category: ["saas"]
brand_color: "#5294E2"
display_name: "Metabase"
short_name: "metabase"
description: "Steampipe plugin for querying Metabase."
og_description: "Query Metabase with SQL! Open source CLI. No DB required."
og_image: "/images/plugins/francois2metz/metabase-social-graphic.png"
icon_url: "/images/plugins/francois2metz/metabase.svg"
---

# Metabase + Steampipe

[Mettabase](https://www.metabase.com) provides fast analytics with the friendly UX and integrated tooling to let your company explore data on their own.

[Steampipe](https://steampipe.io) is an open source CLI to instantly query cloud APIs using SQL.

For example:

```sql
select
  name,
  description,
  id
from
  metabase_db
```

```
+---------------+-------------+-----+
| name          | Description | id  |
+---------------+-------------+-----+
| my-database-1 | Test 1      | 17  |
| my-database-2 | Test 2      | 136 |
+---------------+-------------+-----+
```

## Documentation

- **[Table definitions & examples →](/plugins/1024pix/metabase/tables)**

## Get started

### Install

Download and install the latest Metabase plugin:

```bash
steampipe plugin install 1024pix/metabase
```

### Credentials

| Item        | Description                                                                                                                                                                   |
|-------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| Credentials | Metabase requires a [token](https://www.metabase.com/learn/administration/metabase-api#authenticate-your-requests-with-a-session-token) or login password.                                                                                                   |
| Permissions | Tokens have the same permissions as the user who creates them.                                                                                                                |

### Configuration

Installing the latest metabase plugin will create a config file (`~/.steampipe/config/metabase.spc`) with a single connection named `metabase`:

```hcl
connection "metabase" {
    plugin = "1024pix/metabase"

    # Your metabase url
    # url = https://localhost

    # Metabase credentials
    # user = my_user
    # password = my_password

    # Or token
    # token = xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx

    # Skip tsl verification
    # tls_skip_verify = true
}
```

### Credentials from Environment Variables

The Metabase plugin will use the following environment variables to obtain credentials **only if other argument (`token` or `user` or `password`) is not specified** in the connection:

```sh
export METABASE_TOKEN=xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
```

or

```sh
export METABASE_USER=my-user
export METABASE_PASSWORD=my-password
```

## Get Involved

* Open source: https://github.com/1024pix/steampipe-plugin-metabase

