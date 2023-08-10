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
