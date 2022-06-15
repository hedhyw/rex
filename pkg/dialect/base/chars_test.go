package base_test

import (
	"testing"
	"unicode"

	"github.com/hedhyw/rex/internal/test"
	"github.com/hedhyw/rex/pkg/dialect"
	"github.com/hedhyw/rex/pkg/dialect/base"
)

func TestRexChars(t *testing.T) {
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
	}, {
		Name:     "unicode_greek",
		Chain:    []dialect.Token{base.Chars.Unicode(unicode.Greek)},
		Expected: `\p{Greek}`,
	}, {
		Name:     "unicode_control",
		Chain:    []dialect.Token{base.Chars.Unicode(unicode.Cc)},
		Expected: `\p{Cc}`,
	}, {
		Name:     "unicode_invalid",
		Chain:    []dialect.Token{base.Chars.Unicode(&unicode.RangeTable{})},
		Expected: ``,
	}}.Run(t)
}
