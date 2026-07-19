<h1 align="center">bango</h1>

<p align="center">Search with bangs</p>

> Bangs are sourced from [Helium](https://helium.computer/bangs).

> [!TIP]
> After running, add bango to your browser using a URL like so:
>
> ```
> http://localhost:8479/?q=%s
> ```
> If you don't want to run bango yourself, use my public instance:
>
> ```
> https://bango.theannoying.dev/?q=%s
> ```

## Installation

<details>
<summary>Build from source</summary>

### Prerequisites
Install [Go 1.26+](https://go.dev/doc/install)

### Build from source
```sh
git clone https://github.com/TheAnnoying/bango.git
cd bango
go generate
CGO_ENABLED=0 go build -o bango
```
> If you're on Windows, rename the file to end with `.exe`

Run the binary using `./bango`

</details>

<details>
<summary>Pre-built binary</summary>

Download and run a binary from [the latest release](https://github.com/TheAnnoying/bango/releases/latest)

</details>

<details>
<summary>Docker Compose</summary>

### Docker Compose
Create a file `docker-compose.yml` with the following:
```yml
services:
  bango:
    image: ghcr.io/theannoying/bango:latest
    ports:
      - "8479:8479"
    restart: unless-stopped
```
Run using the command: `docker compose up -d`

</details>

## Environment Variables
| Variable | Usage | Default Value |
| -------- | ----- | ------------- |
| PORT | The port that bango will run on | 8479 |
| BANG_PREFIX | Prefix used for bangs in search queries | ! |
| DEFAULT_SEARCH_PREFIX | The URL that comes before the search term if no bang is present | `https://www.google.com/search?q=` |
| DEFAULT_SEARCH_SUFFIX | The rest of the URL after the search term if no bang is present | None by default. Examples: `&submit_search=`, `&searchon=all&suite=all&section=all` |
