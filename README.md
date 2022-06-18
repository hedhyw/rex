# Rex [work in progress]

![Version](https://img.shields.io/github/v/tag/hedhyw/rex)
[![Go Report Card](https://goreportcard.com/badge/github.com/hedhyw/rex)](https://goreportcard.com/report/github.com/hedhyw/rex)
[![Coverage Status](https://coveralls.io/repos/github/hedhyw/rex/badge.svg?branch=main)](https://coveralls.io/github/hedhyw/rex?branch=main)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/hedhyw/rex)](https://pkg.go.dev/github.com/hedhyw/rex?tab=doc)

![rex-gopher](_docs/gopher.png)

This is a regular expressions builder for gophers!

## Why?

It improves readability and helps to construct regular expressions using human-friendly constructions. Also, it allows commenting and reusing blocks, which improves the quality of the code.

Let's see an example:
```golang
// Using regular expression string.
regexp.MustCompile(`^[a-z]+\[[0-9]+\]$`)

// Using this builder.
rex.New(
    rex.Chars.Begin(), // `^`
    // ID should begin with lowercased character.
    rex.Chars.Range('a', 'z').Repeat().OneOrMore(), // `[a-z]+`
    // ID should contain number inside brackets [#].
    rex.Chars.Single('['), // `[`
    rex.Chars.Digits().Repeat().OneOrMore(), // `[0-9]+`
    rex.Chars.Single(']'), // `]`
    rex.Chars.End(), // `$`
).MustCompile()
```

Yes, it requires more code, but it has its advantages.
> More, but simpler code, fewer bugs.

## Meme

<img alt="Drake Hotline Bling meme" width=300px src="_docs/meme.png" />

_The picture contains two frame fragments from [video](https://www.youtube.com/watch?v=uxpDa-c-4Mc)._

## Documentation

```golang
import "github.com/hedhyw/rex/pkg/rex"

func main() {
    rex.New(/* tokens */).MustCompile() // The same as `regexp.MustCompile`.
    rex.New(/* tokens */).Compile() // The same as `regexp.Compile`.
    rex.New(/* tokens */).String() // Get constructed regular expression as a string.
}
```

### Common

Common operators for core operations.

```golang
rex.Common.Raw(raw string) // Raw regular expression.
rex.Common.Text(text string) // Escaped text.
rex.Common.Class(tokens ...dialect.ClassToken) // Include specified characters.
rex.Common.NotClass(tokens ...dialect.ClassToken) // Exclude specified characters.
```

### Character classes

Single characters and classes, that can be used as-is, as well as childs to `rex.CommonClass` or `rex.CommonNotClass`.

```golang
rex.Chars.Digits()               // `[0-9]`
rex.Chars.Begin()                // `^`
rex.Chars.End()                  // `$`
rex.Chars.Any()                  // `.`
rex.Chars.Range('a', 'z')        // `[a-z]`
rex.Chars.Single('r')            // `r`
rex.Chars.Runes("abc")           // `[abc]`
rex.Chars.Unicode(unicode.Greek) // `\p{Greek}`
rex.Chars.UnicodeByName("Greek") // `\p{Greek}`
rex.Chars.Alphanumeric()         // `[0-9A-Za-z]`
rex.Chars.Alphabetic()           // `[A-Za-z]`
rex.Chars.ASCII()                // `[\x00-\x7F]`
rex.Chars.Whitespace()           // `[\t\n\v\f\r ]`
rex.Chars.WordCharacter()        // `[0-9A-Za-z_]`
rex.Chars.Blank()                // `[\t ]`
rex.Chars.Control()              // `[\x00-\x1F\x7F]`
rex.Chars.Graphical()            // `[[:graph:]]`
rex.Chars.Lower()                // `[a-z]`
rex.Chars.Printable()            // `[ [:graph:]]`
rex.Chars.Punctuation()          // `[!-/:-@[-`{-~]`
rex.Chars.Upper()                // `[A-Z]`
rex.Chars.HexDigits()            // `[0-9A-Fa-f]`
```

If you want to combine mutiple character classes, use `rex.Common.Class`:
```golang
// Only specific characters:
rex.Common.Class(rex.Chars.Digits(), rex.Chars.Single('a'))
// It will produce `[0-9a]`.

// All characters except:
rex.Common.NotClass(rex.Chars.Digits(), rex.Chars.Single('a'))
// It will produce `[^0-9a]`.
```

### Groups

Helpers for grouping expressions.

```golang
// Define a captured group. That can help to select part of the text.
rex.Group.Define(rex.Chars.Single('a'), rex.Chars.Single('b')) // (ab)
// A group that defines "OR" condition for given expressions.
// Example: "a" or "rex", ...
rex.Group.Composite(rex.Chars.Single('a'), rex.Common.Text("rex")) // (?:a|rex)

// Define non-captured group. The result will not be captured.
rex.Group.Define(rex.Chars.Single('a')).NonCaptured() // (?:a)
// Define a group with a name.
rex.Group.Define(rex.Chars.Single('a')).WithName("my_name") // (?P<my_name>a)
```

### Repetitions

Helpers that specify how to repeat characters. They can be called on character class tokens.

```golang
RepetableClassToken.Repeat().OneOrMore() // `+`
RepetableClassToken.ZeroOrMore() // `*`
RepetableClassToken.ZeroOrOne() // `?`
RepetableClassToken.EqualOrMoreThan(n int) // `{n,}`
RepetableClassToken.Between(n, m int) // `{n,m}`

// Example:
rex.Chars.Digits().Repeat().OneOrMore() // [0-9]+
rex.Group.Define(rex.Chars.Single('a')).Repeat().OneOrMore() // (a)+
```

## Examples

### Simple email validator

If we describe an email as `(alphanum)@(alphanum).(2-3 characters)`, then we can define our code:

1. using ASCII classes:

    Issue: [#9](https://github.com/hedhyw/rex/issues/9)
    ```golang
    // TODO
    ```

2. using character ranges:

    ```golang
    alphaNum := rex.Common.Class(
        rex.Chars.Range('a', 'z'),
        rex.Chars.Range('A', 'Z'),
        rex.Chars.Digits(),
    ).Repeat().OneOrMore() // `[a-zA-Z0-9]`

    re := rex.New(
        rex.Chars.Begin(), // `^`

        alphaNum, 
        // Email delimeter.
        rex.Chars.Single('@'), // `@`

        // Domain part.
        alphaNum,

        // Should contain at least one dot.
        rex.Chars.Single('.'), // `\`
        alphaNum.Between(2, 3),

        rex.Chars.End(), // `$`
    ).MustCompile()
    ```

3. using predefined helper:

    Issue: [#10](https://github.com/hedhyw/rex/issues/10)
    ```golang
    // TODO
    ```

4. using raw regular expression:

    ```golang
    rex.New(
        rex.Chars.Begin(), // `^`
        rex.Common.Raw("[a-zA-Z0-9]+"), // `[a-zA-Z0-9]+`
        rex.Chars.Single('@'), // `@`
        rex.Common.Raw("[a-zA-Z0-9]+"), // `[a-zA-Z0-9]+`
        rex.Chars.End(), // `$`
    ).MustCompile()

    // Or even!

    rex.New(
        rex.Common.Raw(`^[a-zA-Z\d]+@[a-zA-Z\d]+\.[a-zA-Z\d]{2,3}$`),
    ).MustCompile()
    ```

#### Sample text matcher

```golang
rex.New(
    // It is safe to use any text in a regular expression, because it will
    // be escaped.
    rex.Common.Text(`hello worldr.`), // `hello worldr\.`
    // It will match exactly the same text.
).MustCompile()
```

#### More examples

More examples can be found here: [pkg/rex/examples_test.go](pkg/rex/examples_test.go).
