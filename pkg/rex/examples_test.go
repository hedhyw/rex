package rex_test

import (
	"fmt"
	"log"
	"unicode"

	"github.com/hedhyw/rex/pkg/rex"
)

func Example_basicMethods() {
	rexRe := rex.New(rex.Chars.Any().Repeat().ZeroOrMore())

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
		rex.Chars.Range('a', 'z').Repeat().OneOrMore(), // `[a-z]+`
		// ID should contain number inside brackets [#].
		rex.Chars.Single('['),                   // `[`
		rex.Chars.Digits().Repeat().OneOrMore(), // `[0-9]+`
		rex.Chars.Single(']'),                   // `]`
		rex.Chars.End(),                         // `$`
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
	) // `[a-zA-Z0-9]`

	re := rex.New(
		rex.Chars.Begin(), // `^`

		alphaNum.Repeat().OneOrMore(),
		// Email delimeter.
		rex.Chars.Single('@'), // `@`

		// Domain part.
		alphaNum.Repeat().OneOrMore(),

		// Should contain at least one dot.
		rex.Chars.Single('.'), // `\`
		alphaNum.Repeat().Between(2, 3),

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

func Example_group() {
	re := rex.New(
		rex.Common.NotClass(rex.Chars.Digits()).Repeat().OneOrMore(),
		rex.Group.Define(
			rex.Chars.Digits().Repeat().OneOrMore(),
		),
		rex.Common.NotClass(rex.Chars.Digits()).Repeat().OneOrMore(),
	).MustCompile()

	fmt.Println("regular expression:", re.String())

	submatches := re.FindStringSubmatch("abc123def")
	fmt.Println("first submatch is:", submatches[0])
	fmt.Println("second submatch is (group):", submatches[1])
	// Output:
	// regular expression: [^\d]+(\d+)[^\d]+
	// first submatch is: abc123def
	// second submatch is (group): 123
}

func Example_groupNonCapturing() {
	re := rex.New(
		rex.Chars.Any().Repeat().ZeroOrMore(),
		rex.Group.Define(
			rex.Chars.Digits().Repeat().OneOrMore(),
			rex.Chars.Single('a'),
		).NonCaptured(),
	).MustCompile()

	fmt.Println("regular expression:", re.String())

	submatches := re.FindStringSubmatch("abc123a")
	fmt.Println("first submatch is:", submatches[0])
	fmt.Println("len (submatches):", len(submatches))
	// Output:
	// regular expression: .*(?:\d+a)
	// first submatch is: abc123a
	// len (submatches): 1
}

func Example_composite() {
	re := rex.New(
		rex.Group.Composite(
			rex.Common.Text("hello"),
			// OR.
			rex.Common.Text("world"),
		).NonCaptured(),
	).MustCompile()

	fmt.Println("regular expression:", re.String())
	fmt.Println(`"hello" matched: `, re.MatchString("hello"))
	fmt.Println(`"worldr" matched: `, re.MatchString("worldr"))
	fmt.Println(`"rex" matched: `, re.MatchString("rex"))
	// Output:
	// regular expression: (?:hello|world)
	// "hello" matched:  true
	// "worldr" matched:  true
	// "rex" matched:  false
}

func Example_groupNamed() {
	re := rex.New(
		rex.Chars.Any().Repeat().OneOrMore(),
		rex.Group.Define(
			rex.Chars.Digits().Repeat().OneOrMore(),
		).WithName("digits"),
		rex.Chars.Any().Repeat().OneOrMore(),
	).MustCompile()

	fmt.Println("regular expression:", re.String())
	fmt.Println("name[1]:", re.SubexpNames()[1])
	fmt.Println("submatch[1]:", re.FindStringSubmatch("abc123abc")[1])

	// Output:
	// regular expression: .+(?P<digits>\d+).+
	// name[1]: digits
	// submatch[1]: 3
}

func Example_groupRepeat() {
	re := rex.New(
		rex.Chars.Begin(),
		rex.Group.Define(
			rex.Common.Text("hello"),
		).Repeat().Between(2, 3),
		rex.Chars.End(),
	).MustCompile()

	fmt.Println("regular expression:", re.String())
	fmt.Println("hello:", re.MatchString("hello"))
	fmt.Println("hellohello:", re.MatchString("hellohello"))
	fmt.Println("hellohellohello:", re.MatchString("hellohellohello"))
	fmt.Println("hellohellohellohello:", re.MatchString("hellohellohellohello"))
	// Output:
	// regular expression: ^(hello){2,3}$
	// hello: false
	// hellohello: true
	// hellohellohello: true
	// hellohellohellohello: false
}
