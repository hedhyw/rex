// nolint: lll,nosnakecase // Generated regular expression can be long.
package rex_test

import (
	"fmt"
	"log"
	"unicode"

	"github.com/hedhyw/rex/pkg/rex"
)

func Example_basicMethods() {
	rexRe := rex.New(rex.Chars.Any().Repeat().ZeroOrMore())

	// Use Compile if you specify dynamic arguments.
	re, err := rexRe.Compile()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(`re.MatchString("a"):`, re.MatchString("a"))

	// Use MustCompile if you don't specify dynamic arguments.
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
		rex.Chars.Lower().Repeat().OneOrMore(), // `[a-z]+`
		// ID should contain number inside brackets [#].
		rex.Group.NonCaptured(
			rex.Chars.Single('['),                   // `[`
			rex.Chars.Digits().Repeat().OneOrMore(), // `[0-9]+`
			rex.Chars.Single(']'),                   // `]`
		),
		rex.Chars.End(), // `$`
	).MustCompile()

	fmt.Println("re.String():", re.String())
	fmt.Println("MatchString(\"abc[0]\"):", re.MatchString("abc[0]"))
	fmt.Println("MatchString(\"abc0\"):", re.MatchString("abc0"))
	// Output:
	// re.String(): ^[[:lower:]]+(?:\[\d+\])$
	// MatchString("abc[0]"): true
	// MatchString("abc0"): false
}

func Example_generalEmailCheck() {
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

	fmt.Println("regular expression:", re.String())
	fmt.Println("rex-lib@example.com.uk:", re.MatchString("rex-lib@example.com.uk"))
	fmt.Println("rexexample.com:", re.MatchString("rexexample.com"))
	// Output:
	// regular expression: ^[a-z[:upper:]\x2D\d]+@[\.a-z[:upper:]\x2D\d]+\.[[:alnum:]]{2,3}$
	// rex-lib@example.com.uk: true
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

func Example_numberRangeRawVerbose() {
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

	fmt.Println("regular expression:", re.String())
	fmt.Println("10.99:", re.MatchString("10.99"))
	fmt.Println("1111.99:", re.MatchString("1111.99"))
	fmt.Println("-111.99:", re.MatchString("-111.99"))
	fmt.Println("-112.00:", re.MatchString("-112.00"))
	fmt.Println("10.999:", re.MatchString("10.999"))
	fmt.Println("111:", re.MatchString("111"))
	// Output:
	// regular expression: ^((?:\x2D(?:0|(?:[1-9])|(?:[1-9][0-9])|(?:10[0-9])|(?:11[0-1])))|(?:0|(?:[1-9])|(?:[1-9][0-9])|(?:[1-9][0-9][0-9])|(?:10[0-9][0-9])|(?:110[0-9])|(?:111[0-1])))\.[0-9]{2}$
	// 10.99: true
	// 1111.99: true
	// -111.99: true
	// -112.00: false
	// 10.999: false
	// 111: false
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
		rex.Group.NonCaptured(
			rex.Chars.Digits().Repeat().OneOrMore(),
			rex.Chars.Single('a'),
		),
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
	fmt.Println(`"world" matched: `, re.MatchString("world"))
	fmt.Println(`"rex" matched: `, re.MatchString("rex"))
	// Output:
	// regular expression: (?:hello|world)
	// "hello" matched:  true
	// "world" matched:  true
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

func Example_phoneMatch() {
	re := rex.New(
		rex.Chars.Begin(),
		rex.Helper.Phone(),
		rex.Chars.End(),
	).MustCompile()

	fmt.Println("regular expression:", re.String())
	fmt.Println("+15555555:", re.MatchString("+15555555"))
	fmt.Println("(607) 123 4567:", re.MatchString("(607) 123 4567"))
	fmt.Println("+22 607 123 4567:", re.MatchString("+22 607 123 4567"))
	fmt.Println("invalid:", re.MatchString("invalid"))

	// Output:
	// regular expression: ^((?:\+[1-9]\d{7,14})|((?:\(\d{3}\)\s\d{3}\s\d{4})|(?:\+[1-9]\d{0,2}\s\d{2,3}\s\d{2,3}\s\d{4})))$
	// +15555555: true
	// (607) 123 4567: true
	// +22 607 123 4567: true
	// invalid: false
}

func Example_phoneFind() {
	re := rex.New(
		rex.Group.Define(
			rex.Helper.Phone(),
		).WithName("phone"),
	).MustCompile()

	const text = `
	E.164:      +15555555
	E.123.Intl: (607) 123 4567
	E.123.Natl: +22 607 123 4567
	`

	fmt.Println("regular expression:", re.String())
	submatches := re.FindAllStringSubmatch(text, -1)

	for i, sub := range submatches {
		fmt.Printf("submatches[%d]: %s\n", i, sub[0])
	}

	// Output:
	// regular expression: (?P<phone>((?:\+[1-9]\d{7,14})|((?:\(\d{3}\)\s\d{3}\s\d{4})|(?:\+[1-9]\d{0,2}\s\d{2,3}\s\d{2,3}\s\d{4}))))
	// submatches[0]: +15555555
	// submatches[1]: (607) 123 4567
	// submatches[2]: +22 607 123 4567
}

func Example_compositeInReadme() {
	re := rex.New(
		rex.Chars.Begin(),
		rex.Group.Composite(
			// Text matches exact text, symbols will be escaped.
			rex.Common.Text("hello."),
			// OR numbers.
			rex.Chars.Digits().Repeat().OneOrMore(),
		),
		rex.Chars.End(),
	).MustCompile()

	fmt.Println("regular expression:", re.String())
	fmt.Println("hello.:", re.MatchString("hello."))
	fmt.Println("hello:", re.MatchString("hello"))
	fmt.Println("123:", re.MatchString("123"))
	fmt.Println("hello.123:", re.MatchString("hello.123"))

	// Output:
	// regular expression: ^(hello\.|\d+)$
	// hello.: true
	// hello: false
	// 123: true
	// hello.123: false
}

func Example_emailMatch() {
	re := rex.New(
		rex.Chars.Begin(),
		rex.Helper.Email(),
		rex.Chars.End(),
	).MustCompile()

	fmt.Println("regular expression:", re.String())
	fmt.Println("example@example.com:", re.MatchString("example@example.com"))
	fmt.Println("@example.com:", re.MatchString("@example.com"))

	// Output:
	// regular expression: ^(?:(?:[[:alnum:]!#\$\x25&'\*\+\x2D/=\?\^_`\{\|\}~](?:[[:alnum:]!#\$\x25&'\*\+\x2D/=\?\^_`\{\|\}~\.]?[[:alnum:]!#\$\x25&'\*\+\x2D/=\?\^_`\{\|\}~]){0,31})@(?:[[:alnum:]][[:alnum:]\x2D]{0,62}(?:\.*[[:alnum:]][[:alnum:]\x2D]{0,62})*[[:alnum:]]))$
	// example@example.com: true
	// @example.com: false
}

func Example_emailFind() {
	re := rex.New(
		rex.Group.Define(
			rex.Helper.Email(),
		).WithName("email"),
	).MustCompile()

	const text = `
	Duyen: duyen@example.com
	Rex:   rex@rex.example.com
	`

	fmt.Println("regular expression:", re.String())
	submatches := re.FindAllStringSubmatch(text, -1)

	for i, sub := range submatches {
		fmt.Printf("submatches[%d]: %s\n", i, sub[0])
	}

	// Output:
	// regular expression: (?P<email>(?:(?:[[:alnum:]!#\$\x25&'\*\+\x2D/=\?\^_`\{\|\}~](?:[[:alnum:]!#\$\x25&'\*\+\x2D/=\?\^_`\{\|\}~\.]?[[:alnum:]!#\$\x25&'\*\+\x2D/=\?\^_`\{\|\}~]){0,31})@(?:[[:alnum:]][[:alnum:]\x2D]{0,62}(?:\.*[[:alnum:]][[:alnum:]\x2D]{0,62})*[[:alnum:]])))
	// submatches[0]: duyen@example.com
	// submatches[1]: rex@rex.example.com
}

func Example_ipv4Match() {
	re := rex.New(
		rex.Chars.Begin(),
		rex.Helper.IPv4(),
		rex.Chars.End(),
	).MustCompile()

	fmt.Println("regular expression:", re.String())
	fmt.Println("127.0.0.1:", re.MatchString("127.0.0.1"))
	fmt.Println("172.217.16.14:", re.MatchString("172.217.16.14"))
	fmt.Println("github.com:", re.MatchString("github.com"))

	// Output:
	// regular expression: ^(?:(?:(?:0|(?:[1-9])|(?:[1-9][0-9])|(?:1[0-9][0-9])|(?:2[0-4][0-9])|(?:25[0-5]))\.){3}(?:0|(?:[1-9])|(?:[1-9][0-9])|(?:1[0-9][0-9])|(?:2[0-4][0-9])|(?:25[0-5])))$
	// 127.0.0.1: true
	// 172.217.16.14: true
	// github.com: false
}

func Example_numberRange() {
	re := rex.New(
		rex.Chars.Begin(),
		rex.Helper.NumberRange(-1, 123),
		rex.Chars.End(),
	).MustCompile()

	fmt.Println("regular expression:", re.String())
	fmt.Println("9:", re.MatchString("9"))
	fmt.Println("100:", re.MatchString("100"))
	fmt.Println("-1:", re.MatchString("-1"))
	fmt.Println("-2:", re.MatchString("-2"))
	fmt.Println("124:", re.MatchString("124"))

	// Output:
	// regular expression: ^((?:\x2D(?:0|1))|(?:0|(?:[1-9])|(?:[1-9][0-9])|(?:1[0-1][0-9])|(?:12[0-3])))$
	// 9: true
	// 100: true
	// -1: true
	// -2: false
	// 124: false
}
