# Usage & API

The public API lives at the module root (`github.com/go-ruby-csv/csv`). It is **Ruby-shaped but Go-idiomatic**: `Parse` / `Generate` mirror Ruby's `CSV.parse` / `CSV.generate`, while the surface follows Go conventions — an explicit `error`, value types, no global state.

!!! success "Status: implemented"
    The library is built and importable as `github.com/go-ruby-csv/csv`, bound into
    `rbgo` as a native module; see [Roadmap](roadmap.md).

## Install

```sh
go get github.com/go-ruby-csv/csv
```

## Worked example

```go
rows, _ := csv.Parse("a,b\n1,2\n")          // [][]string
t, _ := csv.ParseTable("name,n\nx,1\n")     // CSV::Table with headers
s, _ := csv.Generate([][]string{{"a", "b"}, {"1", "2"}})
```

## Shape

```go
// Parse reads CSV text into rows (CSV.parse).
func Parse(s string, opts ...Option) ([]Row, error)

// ParseTable reads CSV text with a header row into a CSV::Table.
func ParseTable(s string, opts ...Option) (*Table, error)

// Generate emits CSV text for the given rows (CSV.generate).
func Generate(rows [][]string, opts ...Option) (string, error)
```

## MRI conformance

Correctness is defined by reference Ruby. A **differential oracle** runs a wide
corpus through both the system `ruby` and this library and compares the results
**byte-for-byte** — not approximated from memory. The oracle tests skip
themselves where `ruby` is not on `PATH` (e.g. the qemu arch lanes), so the
cross-arch builds still validate the library.

## Relationship to Ruby

`go-ruby-csv/csv` is **standalone and reusable**, and is the backend bound into
[go-embedded-ruby](https://github.com/go-embedded-ruby/ruby) by `rbgo` as a
native module — the same way [go-ruby-regexp](https://github.com/go-ruby-regexp)
and [go-ruby-erb](https://github.com/go-ruby-erb) are bound. The dependency runs
the other way: this library has no dependency on the Ruby runtime.
