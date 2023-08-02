connection "metabase" {
    plugin = "metabase"

    # Your metabase url (requiried)
    # url = "https://localhost"

    # Username/password is required for requests. Required except if token (see after) is provided.
    # This can also be set via the `METABASE_USER` and `METABASE_PASSWORD` environment variable.
    # user = "my_user"
    # password = "my_password"

    # Token is required for requests. Required except if user/password (see before) is provided.
    # This can also be set via the `METABASE_TOKEN` environment variable.
    # token = "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"

    # Skip TLS verification, useful in local test. Optionnal.
    # tls_skip_verify = false
}
