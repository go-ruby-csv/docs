<p align="center"><img src="https://raw.githubusercontent.com/go-ruby-csv/brand/main/social/go-ruby-csv.png" alt="go-ruby-csv/docs" width="720"></p>

# go-ruby-csv/docs

Versioned documentation for [go-ruby-csv](https://github.com/go-ruby-csv),
built with [MkDocs Material](https://squidfunk.github.io/mkdocs-material/) and
versioned with [mike](https://github.com/jimporter/mike). Published to the
`gh-pages` branch and served at <https://go-ruby-csv.github.io/docs/>.

The organization landing page ([go-ruby-csv.github.io](https://go-ruby-csv.github.io))
links here.

## Local preview

```bash
python -m venv .venv && . .venv/bin/activate
pip install -r requirements.txt
mkdocs serve                       # http://localhost:8000 (current sources)
mike serve                         # preview the versioned site
```

## Releasing a new docs version

```bash
mike deploy --push --update-aliases <version> latest
mike set-default --push latest
```
