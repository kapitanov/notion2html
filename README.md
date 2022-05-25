# notion2html

Generates a static website from [Notion](https://www.notion.so/) and watches for changes.

## Command line arguments

### Generate once

```shell
notion2html generate -o <output_dir> -t <access_token> [-f|--force]
```

Here:

* `-o <output_dir>` - path to output directory
* `-t <access_token>` - access token for Notion
* `-f`, `--force` - disable incremental update

### Watch for changes

```shell
notion2html watch -o <output_dir> -t <access_token> [-p <period>]
```

Here:

* `-o <output_dir>` - path to output directory
* `-t <access_token>` - access token for Notion
* `-p <period>` - watch timer period (by default - `10m`)

## Run in Docker

Use the following docker compose file as a reference:

```yaml
version: '2.4'
services:
    notion2html:
        build: .
        image: notion2html:latest
        container_name: notion2html
        restart: always
        environment:
            # Access token for Notion (required)
            NOTION_API_TOKEN: '<access_token>'
            # Watch timer period (optional, '10m' by default)
            TIMER_PERIOD: '10m'
            # Path to output directory (optional, '/out' by default)
            OUTPUT_DIR: '/out'
        volumes:
            # Output directory mapping
            - ./out:/out
```

## License

[MIT](LICENSE)
