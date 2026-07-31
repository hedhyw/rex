---
name: regex-to-rex
description: Convert a plain regular expression into equivalent rex builder-style Go code. Use when the user pastes a regex and wants the rex DSL equivalent, asks to "rexify" a pattern, or migrates hand-written regexes to rex.
---

# regex-to-rex

Translate a plain Go (RE2) regular expression into `rex` builder code by hand,
construct by construct, then verify the result.

Do NOT use `cmd/generator` for this: it is an unfinished experiment that wraps
almost everything in `rex.Common.Raw(...)` instead of translating it.

## Process

1. Parse the input pattern mentally into a sequence of constructs
   (anchors, literals, classes, groups, quantifiers, alternations).
2. Map each construct using the table below, top to bottom, preferring the
   most specific helper.
3. Assemble tokens inside `rex.New(...)`, one token per line, with a short
   trailing comment showing the regex fragment it produces.
4. Verify (see "Verify" below). Always.

The only import users need is `github.com/hedhyw/rex/pkg/rex`
(plus `unicode` when using `rex.Chars.Unicode`).

## Mapping table

All helper names below are verified against `pkg/dialect/base`. Do not invent
others; if a construct has no row here, check the source first, then fall back
to `rex.Common.Raw`.

### Literals and dot

| Regex | rex |
|---|---|
| single char `a` | `rex.Chars.Single('a')` |
| literal text `abc.de` (metachars auto-escaped) | `rex.Common.Text("abc.de")` |
| escaped metachar `\.` | `rex.Chars.Single('.')` (escapes automatically) |
| `.` (any char) | `rex.Chars.Any()` |

### Character classes

| Regex | rex |
|---|---|
| `[a-z]` | `rex.Chars.Range('a', 'z')` |
| `[abc]` | `rex.Chars.Runes("abc")` |
| `\d` / `[0-9]` | `rex.Chars.Digits()` |
| `\w` | `rex.Chars.WordCharacter()` |
| `\s` | `rex.Chars.Whitespace()` |
| `\D`, `\W`, `\S` | `rex.Common.NotClass(rex.Chars.Digits())` etc. |
| combined class `[0-9a]` | `rex.Common.Class(rex.Chars.Digits(), rex.Chars.Single('a'))` |
| negated class `[^0-9a]` | `rex.Common.NotClass(rex.Chars.Digits(), rex.Chars.Single('a'))` |
| `[[:alnum:]]` | `rex.Chars.Alphanumeric()` |
| `[[:alpha:]]` | `rex.Chars.Alphabetic()` |
| `[[:ascii:]]` / `[\x00-\x7F]` | `rex.Chars.ASCII()` |
| `[[:blank:]]` / `[\t ]` | `rex.Chars.Blank()` |
| `[[:cntrl:]]` | `rex.Chars.Control()` |
| `[[:graph:]]` | `rex.Chars.Graphical()` |
| `[[:lower:]]` | `rex.Chars.Lower()` |
| `[[:print:]]` | `rex.Chars.Printable()` |
| `[[:punct:]]` | `rex.Chars.Punctuation()` |
| `[[:upper:]]` | `rex.Chars.Upper()` |
| `[[:xdigit:]]` / `[0-9A-Fa-f]` | `rex.Chars.HexDigits()` |
| `\p{Greek}` | `rex.Chars.Unicode(unicode.Greek)` or `rex.Chars.UnicodeByName("Greek")` |

Class tokens compose: any `ClassToken` may be passed into
`rex.Common.Class(...)` / `rex.Common.NotClass(...)`.

### Anchors and boundaries

| Regex | rex |
|---|---|
| `^` | `rex.Chars.Begin()` |
| `$` | `rex.Chars.End()` |
| `\A` | `rex.Chars.BeginOfText()` |
| `\z` | `rex.Chars.EndOfText()` |
| `\b` | `rex.Chars.ASCIIWordBoundary()` |
| `\B` | `rex.Chars.NotASCIIWordBoundary()` |

### Quantifiers

Call `.Repeat()` on a class token or a group token, then chain:

| Regex | rex |
|---|---|
| `x+` | `x.Repeat().OneOrMore()` |
| `x+?` | `x.Repeat().OneOrMorePreferFewer()` |
| `x*` | `x.Repeat().ZeroOrMore()` |
| `x*?` | `x.Repeat().ZeroOrMorePreferFewer()` |
| `x?` | `x.Repeat().ZeroOrOne()` |
| `x??` | `x.Repeat().ZeroOrOnePreferZero()` |
| `x{n}` | `x.Repeat().Exactly(n)` |
| `x{n,}` | `x.Repeat().EqualOrMoreThan(n)` |
| `x{n,}?` | `x.Repeat().EqualOrMoreThanPreferFewer(n)` |
| `x{n,m}` | `x.Repeat().Between(n, m)` |
| `x{n,m}?` | `x.Repeat().BetweenPreferFewer(n, m)` |

Possessive quantifiers (`x*+`, `x++`) do not exist in Go RE2 at all — reject
them, do not approximate silently.

### Groups and alternation

| Regex | rex |
|---|---|
| `(x)` | `rex.Group.Define(x...)` |
| `(?:x)` | `rex.Group.NonCaptured(x...)` |
| `(?P<name>x)` | `rex.Group.Define(x...).WithName("name")` |
| `a\|b` (alternation) | `rex.Group.Composite(a, b)` — produces captured `(a\|b)`; chain `.NonCaptured()` for `(?:a\|b)` |

Note: rex has no way to express a bare top-level alternation without a group.
`rex.Group.Composite(...).NonCaptured()` yields `(?:a|b)` — semantically
equivalent, not byte-identical.

### Escape hatch

| Regex | rex |
|---|---|
| anything unmappable | `rex.Common.Raw("...")` |
| same, with `#` comments and ignored whitespace | `rex.Common.RawVerbose("...")` |

Also consider `rex.Helper.*` (Email, IP, IPv4, IPv6, Phone*, Hostname*,
MD5Hex, SHA1Hex, SHA256Hex, NumberRange) when the input pattern is clearly
one of those — but only if the user wants the rex-canonical pattern, since
the produced regex will differ from their original.

## Known unsupported constructs

State these explicitly instead of guessing:

- **Inline flags** `(?i)`, `(?m)`, `(?s)`, `(?U)`, `(?i:...)`: rex has no
  flags API yet (tracked in hedhyw/rex#31). Pass them through with
  `rex.Common.Raw`.
- **Lookahead/lookbehind** (`(?=...)`, `(?<=...)`, negatives) and
  **backreferences** (`\1`): unsupported by Go's regexp engine itself —
  the pattern cannot be translated at all; tell the user.
- **Possessive quantifiers**: not in RE2, see above.

## Verify

Never deliver unverified output. Write a throwaway program or test that:

1. Builds the rex expression and calls `.MustCompile()` — it must compile.
2. Prints `rex.New(...).String()` and compares it to the original pattern.
   Byte-identical output is not required or expected: rex may emit `\d` where
   the input had `[0-9]`, `[[:xdigit:]]` for `[0-9A-Fa-f]`, `(?:...)` around
   alternations, etc. If the strings differ, compile BOTH with
   `regexp.MustCompile` and compare `MatchString`/`FindStringSubmatch`
   results on a handful of matching and non-matching inputs.
3. Report equivalence explicitly: "byte-identical" or "not byte-identical
   but verified equivalent on these inputs: ...".
