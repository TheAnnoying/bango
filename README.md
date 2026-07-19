<h1 align="center">bango</h1>

<p align="center">Search with bangs</p>

<p align="center">
    [<a href="https://github.com/TheAnnoying/bango/blob/main/README.he-IL.md">עברית</a>]
</p>

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

### Environment Variables
| Variable | Usage | Default Value |
| -------- | ----- | ------------- |
| PORT | The port that bango will run on | 8479 |
| BANG_PREFIX | Prefix used for bangs in search queries | ! |
| DEFAULT_SEARCH_PREFIX | The URL that comes before the search term if no bang is present | `https://www.google.com/search?q=` |
| DEFAULT_SEARCH_SUFFIX | The rest of the URL after the search term if no bang is present | None by default. Examples: `&submit_search=`, `&searchon=all&suite=all&section=all` |

### Docker
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