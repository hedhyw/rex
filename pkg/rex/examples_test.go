package rex_test

import (
	"fmt"
	"log"
	"unicode"

	"github.com/hedhyw/rex/pkg/rex"
)

func Example_basicMethods() {
	rexRe := rex.New(rex.Chars.Any().ZeroOrMore())

	// Use Compile if you spcify dynamic arguments.
	re, err := rexRe.Compile()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(`re.MatchString("a"):`, re.MatchString("a"))

	// Use MustCompile if you don't speicfy dynamic arguments.
	re = rexRe.MustCompile()
	fmt.Println(`re.MatchString("a"):`, re.MatchString("a"))

	// We can get constructed regular expression.
	fmt.Println(`rexRe.String():`, rexRe.String())
	// Output:
	// re.MatchString("a"): true
	// re.MatchString("a"): true
	// rexRe.String(): .*
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
	alphaNum := rex.Common.Class(
		rex.Chars.Range('a', 'z'),
		rex.Chars.Range('A', 'Z'),
		rex.Chars.Digits(),
	).OneOrMore() // `[a-zA-Z0-9]`

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

func Example_unicode() {
	re := rex.New(
		rex.Chars.Begin(),
		rex.Chars.Unicode(unicode.Hiragana),
		rex.Chars.End(),
	).MustCompile()

	fmt.Println("ひ", re.MatchString("ひ"))
	fmt.Println("a", re.MatchString("a"))
	// Output:
	// ひ true
	// a false
}

func Example_unicodeByName() {
	re := rex.New(
		rex.Chars.Begin(),
		rex.Common.Class(
			rex.Chars.UnicodeByName("Hiragana"),
			rex.Chars.UnicodeByName("Katakana"),
		),
		rex.Chars.End(),
	).MustCompile()

	fmt.Println("Hiragana: ひ", re.MatchString("ひ"))
	fmt.Println("Katakana: ヒ", re.MatchString("ひ"))
	fmt.Println("Kanji: 家", re.MatchString("家"))
	fmt.Println("Other: a", re.MatchString("a"))
	// Output:
	// Hiragana: ひ true
	// Katakana: ヒ true
	// Kanji: 家 false
	// Other: a false
}
