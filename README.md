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

### Environment Variables
| Variable | Usage | Default Value |
| -------- | ----- | ------------- |
| PORT | The port that bango will run on | 8479 |
| DEFAULT_SEARCH_PREFIX | The URL that comes before the search term if no bang is present | `https://www.google.com/search?q=` |
| DEFAULT_SEARCH_SUFFIX | The rest of the URL after the search term if no bang is present | None by default. Examples: `&submit_search=`, `&searchon=all&suite=all&section=all` |