package rex_test

import (
	"fmt"
	"log"

	"github.com/hedhyw/rex/pkg/rex"
)

func Example_basicMethods() {
	rexRe := rex.New(rex.Chars.Any().ZeroOrMore())

	// Use Compile if you speicfy dynamic arguments.
	re, err := rexRe.Compile()
	if err != nil {
		log.Fatal(err)
	}

	_ = re

	// Use MustCompile if you don't speicfy dynamic arguments.
	re = rexRe.MustCompile()
	_ = re

	// We can get constructed regular expression.
	fmt.Println(rexRe.String())
	// Output:
	// .*
}

func Example_basicUsage() {
	re := rex.New(
		rex.Chars.Begin(), // `^`
		// ID should begin with lowercased character.
		rex.Chars.Range('a', 'z').OneOrMore(), // `[a-z]+`
		// ID should contain number inside brackets [#].
		rex.Chars.Single('['),          // `[`
		rex.Chars.Digits().OneOrMore(), // `[0-9]+`
		rex.Chars.Single(']'),          // `]`
		rex.Chars.End(),                // `$`
	).MustCompile()

	fmt.Println("re.String():", re.String())
	fmt.Println("MatchString(\"abc[0]\"):", re.MatchString("abc[0]"))
	fmt.Println("MatchString(\"abc0\"):", re.MatchString("abc0"))
	// Output:
	// re.String(): ^[a-z]+\[\d+\]$
	// MatchString("abc[0]"): true
	// MatchString("abc0"): false
}

func Example_emailRange() {
	re := rex.New(
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

	fmt.Println("regular expression:", re.String())
	fmt.Println("rex@example.com:", re.MatchString("rex@example.com"))
	fmt.Println("rexexample.com:", re.MatchString("rexexample.com"))
	// Output:
	// regular expression: ^[a-zA-Z\d]+@[a-zA-Z\d]+\.[a-zA-Z\d]{2,3}$
	// rex@example.com: true
	// rexexample.com: false
}

func Example_emailRawFirst() {
	re := rex.New(
		rex.Chars.Begin(),              // `^`
		rex.Common.Raw("[a-zA-Z0-9]+"), // `[a-zA-Z0-9]+`
		rex.Chars.Single('@'),          // `@`
		rex.Common.Raw("[a-zA-Z0-9]+"), // `[a-zA-Z0-9]+`
		rex.Chars.Single('.'),          // \.
		rex.Common.Raw("[a-zA-Z0-9]+"), // `[a-zA-Z0-9]+`
		rex.Chars.End(),                // `$`
	).MustCompile()

	fmt.Println("regular expression:", re.String())
	fmt.Println("rex@example.com:", re.MatchString("rex@example.com"))
	fmt.Println("rexexample.com:", re.MatchString("rexexample.com"))
	// Output:
	// regular expression: ^[a-zA-Z0-9]+@[a-zA-Z0-9]+\.[a-zA-Z0-9]+$
	// rex@example.com: true
	// rexexample.com: false
}

func Example_emailRawSecond() {
	re := rex.New(
		rex.Common.Raw(`^[a-zA-Z\d]+@[a-zA-Z\d]+\.[a-zA-Z\d]{2,3}$`),
	).MustCompile()

	fmt.Println("regular expression:", re.String())
	fmt.Println("rex@example.com:", re.MatchString("rex@example.com"))
	fmt.Println("rexexample.com:", re.MatchString("rexexample.com"))
	// Output:
	// regular expression: ^[a-zA-Z\d]+@[a-zA-Z\d]+\.[a-zA-Z\d]{2,3}$
	// rex@example.com: true
	// rexexample.com: false
}
