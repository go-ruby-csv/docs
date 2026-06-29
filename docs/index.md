# go-ruby-csv documentation

**Ruby's CSV parser and generator in pure Go — MRI-compatible, no cgo.**

`go-ruby-csv/csv` is a faithful, pure-Go (zero cgo) reimplementation of Ruby's CSV,
matching reference Ruby (MRI) byte-for-byte. The module path is
`github.com/go-ruby-csv/csv`.

It was **extracted from rbgo's prelude/internals into a reusable standalone
library**: the module is standalone and importable by any Go program, and it is
the backend bound into [go-embedded-ruby](https://github.com/go-embedded-ruby/ruby)
by `rbgo` as a native module — just like
[go-ruby-regexp](https://github.com/go-ruby-regexp) and
[go-ruby-erb](https://github.com/go-ruby-erb). The dependency runs the other
way: this library has **no dependency on the Ruby runtime**.

!!! success "Status: parser + generator complete — MRI byte-exact"
    Faithful port of Ruby's CSV: **`Parse`** and **`Generate`** over the full **dialect** surface (`col_sep` / `row_sep` / `quote_char`), **minimal vs. forced quoting** with embedded newlines, **headers** via **`Row`** / **`Table`**, the **`:integer` / `:float` / `:date`** converters, and **`MalformedCSVError`**. Validated by a **differential oracle** against the system `ruby` / `csv` — parsed and generated output compared byte-for-byte — at 100% coverage, `gofmt` + `go vet` clean, CI green across the six 64-bit Go targets and three OSes.

## Quick taste

```go
rows, _ := csv.Parse("a,b\n1,2\n")          // [][]string
t, _ := csv.ParseTable("name,n\nx,1\n")     // CSV::Table with headers
s, _ := csv.Generate([][]string{{"a", "b"}, {"1", "2"}})
```

## Repositories

| Repo | What it is |
| --- | --- |
| [`csv`](https://github.com/go-ruby-csv/csv) | the library — Ruby's CSV in pure Go |
| [`docs`](https://github.com/go-ruby-csv/docs) | this documentation site (MkDocs Material, versioned with mike) |
| [`go-ruby-csv.github.io`](https://github.com/go-ruby-csv/go-ruby-csv.github.io) | the organization landing page (Hugo) |
| [`brand`](https://github.com/go-ruby-csv/brand) | logo and brand assets |

## Principles

- **Pure Go, `CGO_ENABLED=0`** — trivial cross-compilation, a single static
  binary, no C toolchain.
- **MRI byte-exact.** Output matches reference Ruby exactly, not approximately,
  validated by a differential oracle against the `ruby` binary.
- **Standalone & reusable.** Extracted from rbgo's internals; no dependency on
  the Ruby runtime — the dependency runs the other way.
- **100% test coverage** is the target, enforced as a CI gate, across 6 arches
  and 3 OSes.

## Where to go next

- [Why pure Go](why.md) — why this slice of Ruby is deterministic enough to live
  as a standalone, interpreter-independent Go library.
- [Usage & API](api.md) — the public surface and worked examples.
- [Roadmap](roadmap.md) — what is done and what is downstream by design.

Source lives at [github.com/go-ruby-csv/csv](https://github.com/go-ruby-csv/csv).
