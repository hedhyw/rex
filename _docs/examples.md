
## Examples

### Simple email validator

Let's describe a simple email regular expression in order to show the basic functionality (there is a more advanced helper `rex.Helper.Email()`):

```golang
// We can define a set of characters and reuse the block.
customCharacters := rex.Common.Class(
    rex.Chars.Range('a', 'z'), // `[a-z]`
    rex.Chars.Upper(),         // `[A-Z]`
    rex.Chars.Single('-'),     // `\x2D`
    rex.Chars.Digits(),        // `[0-9]`
) // `[a-zA-Z-0-9]`

re := rex.New(
    rex.Chars.Begin(), // `^`
    customCharacters.Repeat().OneOrMore(),

    // Email delimeter.
    rex.Chars.Single('@'), // `@`

    // Allow dot after delimter.
    rex.Common.Class(
        rex.Chars.Single('.'), // \.
        customCharacters,
    ).Repeat().OneOrMore(),

    // Email should contain at least one dot.
    rex.Chars.Single('.'), // `\.`
    rex.Chars.Alphanumeric().Repeat().Between(2, 3),

    rex.Chars.End(), // `$`
).MustCompile()
```

#### Simple composite

```golang
re := rex.New(
    rex.Chars.Begin(),
    rex.Group.Composite(
        // Text matches exact text (symbols will be escaped).
        rex.Common.Text("hello."),
        // OR one or more numbers.
        rex.Chars.Digits().Repeat().OneOrMore(),
    ),
    rex.Chars.End(),
).MustCompile()

re.MatchString("hello.")    // true
re.MatchString("hello")     // false
re.MatchString("123")       // true
re.MatchString("hello.123") // false
```

## Example match usage

```golang
re := rex.New(
    // Define a named group.
    rex.Group.Define(
        rex.Helper.Phone(),
    ).WithName("phone"),
).MustCompile()

const text = `
E.164:      +15555555
E.123.Intl: (607) 123 4567
E.123.Natl: +22 607 123 4567
`

submatches := re.FindAllStringSubmatch(text, -1)
// submatches[0]: +15555555
// submatches[1]: (607) 123 4567
// submatches[2]: +22 607 123 4567
```

#### More examples

More examples can be found here: [examples_test.go](examples_test.go).
