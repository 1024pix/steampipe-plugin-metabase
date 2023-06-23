![Steampipe + Metabase](docs/metabase-social-graphic.png)

# Metabase plugin for Steampipe

Use SQL to query [Metabase][].

- **[Get started →](docs/index.md)**
- Documentation: [Table definitions & examples](docs/tables)

Compatible with Metabase v0.46.4

## Quick start

Install the plugin with [Steampipe][]:

    steampipe plugin install 1024pix/metabase

## Development

To build the plugin and install it in your `.steampipe` directory

    make

Copy the default config file:

    cp config/metabase.spc ~/.steampipe/config/metabase.spc

## License

Apache 2

[steampipe]: https://steampipe.io
[metabase]: https://metabase.com
