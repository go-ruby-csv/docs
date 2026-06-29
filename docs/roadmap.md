# Roadmap

`go-ruby-csv/csv` is grown **test-first**, each capability differential-tested against MRI
rather than built in isolation. Ruby's csv — the
deterministic, interpreter-independent slice extracted from rbgo's internals — is
**complete**.

| Stage | What | Status |
| --- | --- | --- |
| Parse & generate | `CSV.parse` and `CSV.generate` over the full dialect surface: column/row separators, quote character, and the quoting rules Ruby's CSV uses. | **Done** |
| Dialects & quoting | Configurable `col_sep` / `row_sep` / `quote_char`, with minimal vs. forced quoting and embedded-newline handling matching MRI byte-for-byte. | **Done** |
| Headers | `Row` and `Table` with header access, duplicate-header handling, and header/field converters, the way reference Ruby's `CSV::Row` / `CSV::Table` work. | **Done** |
| Converters | The built-in `:integer` / `:float` / `:date` field converters and header converters, applied with Ruby's ordering and fall-through semantics. | **Done** |
| MalformedCSVError | Raised on unclosed quotes and stray quote characters at the same positions, with the same line/column information as reference Ruby. | **Done** |
| Differential oracle & coverage | A wide corpus parsed and generated both here and by the system `ruby`/`csv`, compared byte-for-byte; 100% coverage, gofmt + go vet clean, green across all six 64-bit Go arches and three OSes. | **Done** |

## Documented out-of-scope boundaries

These are **deliberate**, recorded so the module's surface is unambiguous:

- **No interpreter.** The library implements the deterministic algorithm; it
  never runs arbitrary Ruby. Anything that needs a live binding or evaluation is
  the consumer's job — that is why `rbgo` binds this module rather than the
  reverse.
- **Reference is reference Ruby (MRI).** Byte-for-byte conformance targets MRI's
  behaviour; differences across MRI releases are matched to the reference used by
  the differential oracle.
- **Standalone & reusable.** The module has no dependency on the Ruby runtime;
  the dependency runs the other way.

See [Usage & API](api.md) for the surface and [Why pure Go](why.md) for the
deterministic/interpreter split.
