package base_test

import (
	"testing"
	"unicode"

	"github.com/hedhyw/rex/internal/test"
	"github.com/hedhyw/rex/pkg/dialect"
	"github.com/hedhyw/rex/pkg/dialect/base"
)

func TestRexCharsASCII(t *testing.T) {
	test.RexTestCasesSlice{{
		Name:     "Alphanumeric",
		Chain:    []dialect.Token{base.Chars.Alphanumeric()},
		Expected: `[[:alnum:]]`,
	}, {
		Name:     "Alphabetic",
		Chain:    []dialect.Token{base.Chars.Alphabetic()},
		Expected: `[[:alpha:]]`,
	}, {
		Name:     "ASCII",
		Chain:    []dialect.Token{base.Chars.ASCII()},
		Expected: `[[:ascii:]]`,
	}, {
		Name:     "Whitespace",
		Chain:    []dialect.Token{base.Chars.Whitespace()},
		Expected: `\s`,
	}, {
		Name:     "WordCharacter",
		Chain:    []dialect.Token{base.Chars.WordCharacter()},
		Expected: `\w`,
	}, {
		Name:     "Blank",
		Chain:    []dialect.Token{base.Chars.Blank()},
		Expected: `[[:blank:]]`,
	}, {
		Name:     "Control",
		Chain:    []dialect.Token{base.Chars.Control()},
		Expected: `[[:cntrl:]]`,
	}, {
		Name:     "Graphical",
		Chain:    []dialect.Token{base.Chars.Graphical()},
		Expected: `[[:graph:]]`,
	}, {
		Name:     "Lower",
		Chain:    []dialect.Token{base.Chars.Lower()},
		Expected: `[[:lower:]]`,
	}, {
		Name:     "Printable",
		Chain:    []dialect.Token{base.Chars.Printable()},
		Expected: `[[:print:]]`,
	}, {
		Name:     "Punctuation",
		Chain:    []dialect.Token{base.Chars.Punctuation()},
		Expected: `[[:punct:]]`,
	}, {
		Name:     "Upper",
		Chain:    []dialect.Token{base.Chars.Upper()},
		Expected: `[[:upper:]]`,
	}, {
		Name:     "HexDigits",
		Chain:    []dialect.Token{base.Chars.HexDigits()},
		Expected: `[[:xdigit:]]`,
	}}.Run(t)
}

func TestRexChars_base(t *testing.T) {
	test.RexTestCasesSlice{{
		Name:     "any",
		Chain:    []dialect.Token{base.Chars.Any()},
		Expected: `.`,
	}, {
		Name:     "digits",
		Chain:    []dialect.Token{base.Chars.Digits()},
		Expected: `\d`,
	}, {
		Name:     "begin",
		Chain:    []dialect.Token{base.Chars.Begin()},
		Expected: `^`,
	}, {
		Name:     "beginOfText",
		Chain:    []dialect.Token{base.Chars.BeginOfText()},
		Expected: `\A`,
	}, {
		Name:     "end",
		Chain:    []dialect.Token{base.Chars.End()},
		Expected: `$`,
	}, {
		Name:     "endOfText",
		Chain:    []dialect.Token{base.Chars.EndOfText()},
		Expected: `\z`,
	}, {
		Name:     "ASCIIWordBoundary",
		Chain:    []dialect.Token{base.Chars.ASCIIWordBoundary()},
		Expected: `\b`,
	}, {
		Name:     "notASCIIWordBoundary",
		Chain:    []dialect.Token{base.Chars.NotASCIIWordBoundary()},
		Expected: `\B`,
	}, {
		Name:     "single",
		Chain:    []dialect.Token{base.Chars.Single('a')},
		Expected: `a`,
	}, {
		Name:     "single_escaped",
		Chain:    []dialect.Token{base.Chars.Single('.')},
		Expected: `\.`,
	}, {
		Name:     "single_hex_large",
		Chain:    []dialect.Token{base.Chars.Single('ở')},
		Expected: `\x{1EDF}`,
	}, {
		Name:     "single_hex_small",
		Chain:    []dialect.Token{base.Chars.Single(unicode.MaxASCII + 1)},
		Expected: `\x80`,
	}, {
		Name:     "range_upper",
		Chain:    []dialect.Token{base.Chars.Range('A', 'Z')},
		Expected: `[A-Z]`,
	}, {
		Name:     "range_digits",
		Chain:    []dialect.Token{base.Chars.Range('0', '9')},
		Expected: `[0-9]`,
	}}.Run(t)
}

func TestRexChars_unicode(t *testing.T) {
	test.RexTestCasesSlice{{
		Name:     "unicode_greek",
		Chain:    []dialect.Token{base.Chars.Unicode(unicode.Greek)},
		Expected: `\p{Greek}`,
	}, {
		Name:     "unicode_control",
		Chain:    []dialect.Token{base.Chars.Unicode(unicode.Cc)},
		Expected: `\p{Cc}`,
	}, {
		Name: "unicode_invalid",
		Chain: []dialect.Token{base.Chars.Unicode(&unicode.RangeTable{
			R16:         nil,
			R32:         nil,
			LatinOffset: 0,
		})},
		Expected: ``,
	}, {
		Name:     "unicode_by_name_greek",
		Chain:    []dialect.Token{base.Chars.UnicodeByName("Greek")},
		Expected: `\p{Greek}`,
	}, {
		Name:     "unicode_by_name_control",
		Chain:    []dialect.Token{base.Chars.UnicodeByName("Cc")},
		Expected: `\p{Cc}`,
	}}.Run(t)
}

func TestRexChars_runes(t *testing.T) {
	test.RexTestCasesSlice{{
		Name: "single",
		Chain: []dialect.Token{
			base.Chars.Runes("a"),
		},
		Expected: `a`,
	}, {
		Name: "abc",
		Chain: []dialect.Token{
			base.Chars.Runes("abc"),
		},
		Expected: `[abc]`,
	}, {
		Name: "escaped",
		Chain: []dialect.Token{
			base.Chars.Runes(".+"),
		},
		Expected: `[\.\+]`,
	}, {
		Name: "unicode",
		Chain: []dialect.Token{
			// nolint: gosmopolitan // A test.
			base.Chars.Runes("ひヒ家"),
		},
		Expected: `[\x{3072}\x{30D2}\x{5BB6}]`,
	}, {
		Name: "single_unicode",
		Chain: []dialect.Token{
			base.Chars.Runes("ひ"),
		},
		Expected: `\x{3072}`,
	}, {
		Name: "empty",
		Chain: []dialect.Token{
			base.Chars.Runes(""),
		},
		Expected: ``,
	}, {
		Name: "WrappedOnce",
		Chain: []dialect.Token{
			base.Common.Class(
				base.Chars.Runes("abc"),
			),
		},
		Expected: `[abc]`,
	}, {
		Name: "CanBeRepetable",
		Chain: []dialect.Token{
			base.Chars.Runes("abc").Repeat().OneOrMore(),
		},
		Expected: `[abc]+`,
	}, {
		Name: "Punctuation",
		Chain: []dialect.Token{
			base.Chars.Runes("!#$%&'*+-/=?^_`{|}~"),
		},
		Expected: "[!#\\$\\x25&'\\*\\+\\x2D/=\\?\\^_`\\{\\|\\}~]",
	}}.Run(t)
}
