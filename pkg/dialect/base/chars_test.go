package base_test

import (
	"testing"
	"unicode"

	"github.com/hedhyw/rex/internal/test"
	"github.com/hedhyw/rex/pkg/dialect"
	"github.com/hedhyw/rex/pkg/dialect/base"
)

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
		Name:     "end",
		Chain:    []dialect.Token{base.Chars.End()},
		Expected: `$`,
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

func TestRexChars_Runes(t *testing.T) {
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
		Name: "wrapped_once",
		Chain: []dialect.Token{
			base.Common.Class(
				base.Chars.Runes("abc"),
			),
		},
		Expected: `[abc]`,
	}}.Run(t)
}
