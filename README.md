# notion2html

---

* [Command line arguments](#command-line-arguments)
  * [Generate once](#generate-once)
  * [Watch for changes](#watch-for-changes)
* [Run in Docker](#run-in-docker)
* [Example - generate a static website from Notion](#example---generate-a-static-website-from-notion)
* [Example - serve a website from a Notion space (with HTTPS!)](#example---serve-a-website-from-a-notion-space-with-https)
* [How to get an access token](#how-to-get-an-access-token)
* [License](#license)

---

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
* `-p <period>` - watch timer period (by default - `1h`)

## Run in Docker

Use the following docker compose file as a reference:

```yaml
version: '2.4'
services:
    notion2html:
        image: ghcr.io/kapitanov/notion2html:latest
        environment:
            # Access token for Notion (required)
            NOTION_API_TOKEN: '<access_token>'
            # Watch timer period (optional, '1h' by default)
            TIMER_PERIOD: '1h'
            # Path to output directory (optional, '/out' by default)
            OUTPUT_DIR: '/out'
        volumes:
            # Output directory mapping
            - ./out:/out
```

## Example - generate a static website from Notion

First, you would need [an access token for Notion](#how-to-get-an-access-token).

Then, you can run the following command:

```shell
docker run -t --rm -v $(pwd):/mnt ghcr.io/kapitanov/notion2html:latest generate --token "<access_token>" --output /mnt --force
```

> Note `"<1access_token>"` above - you should put your access token there.

This command will export all pages from Notion to the current directory as a single static website.

## Example - serve a website from a Notion space (with HTTPS!)

First, you would need [an access token for Notion](#how-to-get-an-access-token).
Second, you would need a domain name configured for your server.

Then, you can run the following command (on your server):

```bash
$ ACCESS_TOKEN="<access_token>" DOMAIN="<domain>" cat << EOF > docker-compose.yaml
services:
    notion2html:
        image: ghcr.io/kapitanov/notion2html:latest
        restart: always
        logging:
            driver: "json-file"
            options:
                max-size: "10m"
                max-file: "1"
        environment:
            NOTION_API_TOKEN: '$ACCESS_TOKEN'
            TIMER_PERIOD: '1h'
            OUTPUT_DIR: '/out'
        volumes:
            - ./www:/out
    nginx:
        image: steveltn/https-portal:latest
        ports:
            - 80:80
            - 443:443
        restart: always
        logging:
            driver: "json-file"
            options:
                max-size: "10m"
                max-file: "1"
        volumes:
            - ./out:/usr/share/nginx/html
        environment:
            DOMAINS: '$DOMAIN'
        volumes:
            - ./https-portal-data:/var/lib/https-portal
            - ./www:/var/www/vhosts/$DOMAIN
EOF
$ docker-compose up -d
```

> Note `ACCESS_TOKEN="<access_token>"` and `DOMAIN="<domain>"` above - you should put your access token and domain name there.

This command will export all pages from Notion to the `www` directory as a single static website and serve it with HTTPS.
HTTPS will be powered by an awesome [https-portal](https://github.com/SteveLTN/https-portal),
with automatic LetsEncrypt certificates.

## How to get an access token

1. Visit [notion.so/my-integrations](https://www.notion.so/my-integrations) and click "New integration" button.
2. Enter a name for your integration (you will need it to connect this integration to your page).
3. Unckeck everything except "Read content".
4. Click "Submit" button.
5. On the next page, you will see a "Internal Integration Token" field - that is the access token.
6. Now, open a root Notion page and click "..." button in the top right corner.
7. CLick on "Add connection" menu item and add your integration (yoi might need to use search by name here).

## License

[MIT](LICENSE)
