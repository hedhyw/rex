# Rex [work in progress]

![Version](https://img.shields.io/github/v/tag/hedhyw/rex)
[![Go Report Card](https://goreportcard.com/badge/github.com/hedhyw/rex)](https://goreportcard.com/report/github.com/hedhyw/rex)
[![Coverage Status](https://coveralls.io/repos/github/hedhyw/rex/badge.svg?branch=main)](https://coveralls.io/github/hedhyw/rex?branch=main)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/hedhyw/rex)](https://pkg.go.dev/github.com/hedhyw/rex?tab=doc)

This is a regular expressions builder for humans!

## Why?

It improves readability and helps to construct regular expressions using human-friendly constructions. Also, it allows commenting blocks, which improves the quality of the code.

Let's see an example:
```golang
// Using reg expression string.
regexp.MustCompile(`^[a-z]+\[[0-9]+\]$`)

// Using this builder.
rex.New(
    rex.Chars.Begin(),
    // ID should begin with lowercased character.
    rex.Chars.Range('a', 'z').OneOrMore(),
    // ID should contain number inside brackets [#].
    rex.Chars.Single('['),
    rex.Chars.Digits().OneOrMore(),
    rex.Chars.Single(']'),
    rex.Chars.End(),
).MustCompile()
```

Yes, it requires more code, but it has its advantages.
> More, but simpler code, fewer bugs.

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

Common operators.

```golang
Raw(raw string) // Raw regular expression.
Text(text string) // Escaped text.
Class(tokens ...dialect.Token) // Include specified characters.
NotClass(tokens ...dialect.Token) // Exclude specified characters.
Single(r rune) // Single character.
```

### Character classes

Single characters and classes. They can be used as-is, as well as a child to `Class` or `NotClass`.

```golang
rex.Chars.Digits() // [0-9]
rex.Chars.Begin() // ^
rex.Chars.End() // $
rex.Chars.Any() // .
rex.Chars.Range(from rune, to rune)  // [a-z]
rex.Chars.Single(r rune) // r
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

### Character classes

```golang
rex.Chars.Digits() // [0-9]
rex.Chars.Begin() // ^
rex.Chars.End() // $
rex.Chars.Any() // .
rex.Chars.Range(from rune, to rune)  // [a-z]
rex.Chars.Single(r rune) // r
```

### Repetitions

```golang
ClassToken.OneOrMore() // +
ClassToken.ZeroOrMore() // *
ClassToken.ZeroOrOne() // ?
ClassToken.EqualOrMoreThan(n int) // {n,}
ClassToken.Between(n, m int) // {n,m}
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
    rex.New(
        rex.Chars.Begin(), // `^`

        rex.Common.Class( // `[a-zA-Z0-9]`
            rex.Chars.Range('a', 'z'),
            rex.Chars.Range('A', 'Z'),
            rex.Chars.Digits(),
        ).OneOrMore(),

        // Email delimeter.
        rex.Chars.Single('@'), // `@`

        // Domain part.
        rex.Common.Class(
            rex.Chars.Range('a', 'z'),
            rex.Chars.Range('A', 'Z'),
            rex.Chars.Digits(),
        ).OneOrMore(),

        // Should contain at least one dot.
        rex.Chars.Single('.'), // `\`

        rex.Common.Class(
            rex.Chars.Range('a', 'z'),
            rex.Chars.Range('A', 'Z'),
            rex.Chars.Digits(),
        ).Between(2, 3),

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
