---
organization: 1024pix
category: ["saas"]
brand_color: "#5294E2"
display_name: "Metabase"
short_name: "metabase"
description: "Steampipe plugin for querying Metabase."
og_description: "Query Metabase with SQL! Open source CLI. No DB required."
og_image: "/images/plugins/1024pix/metabase-social-graphic.png"
icon_url: "/images/plugins/1024pix/metabase.svg"
---

# Metabase + Steampipe

[Metabase](https://www.metabase.com) provides fast analytics with the friendly UX and integrated tooling to let your company explore data on their own.

[Steampipe](https://steampipe.io) is an open source CLI to instantly query cloud APIs using SQL.

List database in your Metabase instance:

```sql
select
  name,
  description,
  id
from
  metabase_db;
```

```
+---------------+-------------+-----+
| name          | description | id  |
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

    # Your metabase url (required)
    # url = "https://localhost"

    # Username/password is required for requests. Required except if token (see after) is provided.
    # This can also be set via the `METABASE_USER` and `METABASE_PASSWORD` environment variable.
    # user = "my_user"
    # password = "my_password"

    # Token is required for requests. Required except if user/password (see before) is provided.
    # This can also be set via the `METABASE_TOKEN` environment variable.
    # token = "33d0d62a-6a16-3083-ba7b-3bab31bd6612"

    # Skip TLS verification, useful in local test. Optionnal.
    # tls_skip_verify = false
}
```

### Credentials from Environment Variables

Alternatively, you can also use the standard Metabase environment variables to obtain credentials **only if other arguments (`token` or `user`/`password`) are not specified** in the connection:

```sh
export METABASE_TOKEN=33d0d62a-6a16-3083-ba7b-3bab31bd6612
```

or

```sh
export METABASE_USER=my-user
export METABASE_PASSWORD=my-password
```

## Get involved

- Open source: https://github.com/1024pix/steampipe-plugin-metabase
- Community: [Join #steampipe on Slack →](https://turbot.com/community/join)
