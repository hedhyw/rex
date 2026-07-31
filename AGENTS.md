# AGENTS.md

Guide for AI coding agents working on this repository.

## What is this project?

Rex (`github.com/hedhyw/rex`) is a zero-dependency Go library for building
regular expressions from human-friendly, composable tokens. Instead of writing
raw regex syntax, users combine typed helpers (`rex.Chars`, `rex.Common`,
`rex.Group`, `rex.Helper`) and compile the result into a standard
`*regexp.Regexp`. The repository also ships a small CLI
(`cmd/generator`) that converts an existing regular expression into
equivalent rex builder code.

## Public API

The user-facing import path is:

```go
import "github.com/hedhyw/rex/pkg/rex"
```

- `rex.New(tokens ...dialect.Token) *RegExp` — builds an expression; then
  `.Compile()`, `.MustCompile()` (returning `*regexp.Regexp`) or `.String()`.
- `rex.Chars` — character classes: `Begin()`, `End()`, `Any()`,
  `Digits()`, `Range('a','z')`, `Runes("abc")`, `Single('r')`,
  `Unicode(unicode.Greek)`, etc.
- `rex.Common` — core operators: `Raw(s)`, `RawVerbose(s)`, `Text(s)`
  (escaped literal), `Class(...)`, `NotClass(...)`.
- `rex.Group` — grouping: `Define(...)`, `NonCaptured(...)`,
  `Composite(...)` (OR), plus `.WithName(...)` and `.Repeat()`.
- `rex.Helper` — ready-made patterns: `NumberRange`, `Phone`/`PhoneE164`/
  `PhoneE123`, `Email`, `IP`/`IPv4`/`IPv6`, `HostnameRFC952`/`HostnameRFC1123`,
  `MD5Hex`/`SHA1Hex`/`SHA256Hex`.
- Repetitions are chained via `.Repeat()`: `OneOrMore()`, `ZeroOrMore()`,
  `ZeroOrOne()`, `Exactly(n)`, `Between(from, to)`, and `...PreferFewer`
  (lazy) variants.

Supporting packages (also public, but most users only need `pkg/rex`):

- `github.com/hedhyw/rex/pkg/dialect` — `Token`, `ClassToken`,
  `StringByteWriter` interfaces that all tokens implement.
- `github.com/hedhyw/rex/pkg/dialect/base` — the base dialect that
  implements `Chars`, `Common`, `Group`, `Helper`; `pkg/rex` re-exports
  these as aliases.

## Repository layout

- `pkg/rex/` — public facade: `RegExp` builder (`rex.go`) and namespace
  aliases (`base.go`). Usage examples live in `examples_test.go`.
- `pkg/dialect/` — token interfaces shared by dialects.
- `pkg/dialect/base/` — token implementations: `chars.go`, `class.go`,
  `common.go`, `group.go`, `composite.go`, `repetitions.go`, `raw.go`,
  and `helper_*.go` (number, phone, web, hash patterns) with tests.
- `internal/helper/` — token stream processing used by `rex.New`.
- `internal/generator/` — regex-to-rex-code generation logic.
- `cmd/generator/` — CLI entry point: `rex '<regex>'` prints rex builder
  code for a given regular expression.
- `_docs/` — `library.md` (API documentation), `examples.md`, images.
- `bin/` — git-ignored; holds the downloaded golangci-lint and built CLI.

## Build, test, lint

- `make lint` — runs golangci-lint (version pinned via
  `GOLANG_CI_LINT_VER` in the Makefile; downloaded automatically into
  `bin/`). Config: `.golangci.json`.
- `make test` — runs `go test` for all packages with coverage
  (`coverage.out`) and prints a per-function coverage summary.
- `make all` (default) — lint + test.
- `make test.fuzz NAME=FuzzIPv4` — runs a fuzz test from
  `pkg/dialect/base`.
- `make build` — builds the generator CLI to `./bin/rex`.
- `make tidy` — `go mod tidy`.

CI (`.github/workflows/check.yml`) runs `make lint` and `make test` and
uploads coverage to Coveralls. PR titles must follow Conventional Commits
(enforced by `.github/workflows/semantic.yaml`).

## Code generation

There is no generated code in this repository; all source files are
hand-written and may be edited. (`cmd/generator` / `internal/generator`
are a user-facing tool for converting regexes to rex code, not part of
the build.)

## Conventions

- The module has zero external dependencies — do not add any without a
  very strong reason. Test assertions use the local
  `internal/test` helpers instead of a third-party assertion library.
- Public API is stable (v1); avoid breaking changes to exported
  identifiers in `pkg/`.
- Every token change needs tests; token behavior is asserted by
  comparing the produced regex string and match results (see
  `*_test.go` in `pkg/dialect/base`).
- Keep `_docs/library.md` and the README in sync with API changes.
- Commit messages and PR titles follow Conventional Commits
  (`feat:`, `fix:`, `docs:`, `chore:`, ...), terse and imperative.
