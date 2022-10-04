# Rex

![Version](https://img.shields.io/github/v/tag/hedhyw/rex)
[![Go Report Card](https://goreportcard.com/badge/github.com/hedhyw/rex)](https://goreportcard.com/report/github.com/hedhyw/rex)
[![Coverage Status](https://coveralls.io/repos/github/hedhyw/rex/badge.svg?branch=main)](https://coveralls.io/github/hedhyw/rex?branch=main)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/hedhyw/rex)](https://pkg.go.dev/github.com/hedhyw/rex?tab=doc)

![rex-gopher](_docs/gopher.png)

**This is a regular expressions builder for gophers!**

- **[Why?](#why)**
- **[FAQ](#faq)**
- **[Documentation](_docs/library.md)**
- **[Examples](pkg/rex/examples_test.go)**
- **[License](#license)**

## Why?

It makes readability better and helps to construct regular expressions using human-friendly constructions. Also, it allows commenting and reusing blocks, which improves the quality of code. It provides a convenient way to use parameterized patterns. It is easy to implement custom patterns or use a combination of others.

It is just a builder, so it returns standart [`*regexp.Regexp`](https://pkg.go.dev/regexp#Regexp).

The library supports [groups](_docs/library.md#groups), [composits](_docs/library.md#groups), [classes](_docs/library.md#character-classes), [flags](_docs/library.md#flags), [repetitions](_docs/library.md#repetitions) and if you want you can even use `raw regular expressions` in any place. Also it contains a set of [predefined helpers](_docs/library.md#helper) with patterns for number ranges, phones, emails, etc...

Let's see an example of validating or matching `someid[#]` using a verbose pattern:
```golang
re := rex.New(
    rex.Chars.Begin(), // `^`
    // ID should begin with lowercased character.
    rex.Chars.Lower().Repeat().OneOrMore(), // `[a-z]+`
    // ID should contain number inside brackets [#].
    rex.Group.NonCaptured( // (?:)
        rex.Chars.Single('['),                   // `[`
        rex.Chars.Digits().Repeat().OneOrMore(), // `[0-9]+`
        rex.Chars.Single(']'),                   // `]`
    ),
    rex.Chars.End(), // `$`
).MustCompile()
```

Yes, it requires more code, but it has its advantages.
> More, but simpler code, fewer bugs.

You can still use original regular expressions as is in any place. Example of
matching numbers between `-111.99` and `1111.99` using a combination of
patterns and raw regular expression:

```golang
re := rex.New(
    rex.Common.Raw(`^`),
    rex.Helper.NumberRange(-111, 1111),
    rex.Common.RawVerbose(`
        # RawVerbose is a synonym to Raw,
        # but ignores comments, spaces and new lines.
        \.        # Decimal delimter.  
        [0-9]{2}  # Only two digits.
        $         # The end.
    `),
).MustCompile()

// Produces:
// ^((?:\x2D(?:0|(?:[1-9])|(?:[1-9][0-9])|(?:10[0-9])|(?:11[0-1])))|(?:0|(?:[1-9])|(?:[1-9][0-9])|(?:[1-9][0-9][0-9])|(?:10[0-9][0-9])|(?:110[0-9])|(?:111[0-1])))\.[0-9]{2}$
```

> The style you prefer is up to you.

## Meme

<img alt="Drake Hotline Bling meme" width=350px src="_docs/meme.png" />

## FAQ

1. **It is too verbose. Too much code.**

    More, but simpler code, fewer bugs.
    Anyway, you can still use the raw regular expressions syntax in combination with helpers.
    ```golang
    rex.New(
        rex.Chars.Begin(),
        rex.Group.Define(
            // `Raw` can be placed anywhere in blocks.
            rex.Common.Raw(`[a-z]+\d+[A-Z]*`),
        ),
        rex.Chars.End(),
    )
    ```
    Or just raw regular expression with comments:
    ```golang
    rex.Common.RawVerbose(`
        ^                # Start of the line.
        [a-zA-Z0-9]+     # Local part.
        @                # delimeter.
        [a-zA-Z0-9\.]+   # Domain part.
        $                # End of the line.
    `)
    ```

2. **Should I know regular expressions?**

   It is better to know them in order to use this library most effectively.
   But in any case, it is not strictly necessary.

3. **Is it language-dependent? Is it transferable to other languages?**

   We can use this library only in Go. If you want to use any parts
   in other places, then just call `rex.New(...).String()` and copy-paste
   generated regular expression.

4. **What about my favourite `DSL`?**

   Every IDE has convenient auto-completion for languages. So all helpers
   of this library are easy to use out of the box. Also, it is easier
   to create custom parameterized helpers.

5. **Is it stable?**

   It is `0.X.Y` version, but there are some backward compatibility guarantees:
   - `rex.Chars` helpers can change output to an alternative synonym.
   - `rex.Common` helpers can be deprecated, but not removed.
   - `rex.Group` some methods can be deprecated.
   - `rex.Helper` can be changed with breaking changes due to specification complexities.
   - The test coverage should be `~100%` without covering [test helpers](internal/test/test.go).
   - Any breaking change will be prevented as much as possible.

   _All of the above may not be respected when upgrading the major version._

6. **I have another question. I found an issue. I have a feature request. I want to contribute.**

   Please, [create an issue](https://github.com/hedhyw/rex/issues/new?labels=question&title=I+have+a+question).

## License

- The library is under [MIT Lecense](LICENSE)
- [The gopher](_docs/gopher.png) is under [Creative Commons Attribution 3.0](https://creativecommons.org/licenses/by/3.0/) license. It was originally created by [Renée French](https://en.wikipedia.org/wiki/Ren%C3%A9e_French) and redrawed by me.
- [The meme](_docs/meme.png) contains two frame fragments from [the video](https://www.youtube.com/watch?v=uxpDa-c-4Mc).
