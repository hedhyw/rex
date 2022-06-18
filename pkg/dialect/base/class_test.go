package base_test

import (
	"testing"

	"github.com/hedhyw/rex/internal/test"
	"github.com/hedhyw/rex/pkg/dialect"
	"github.com/hedhyw/rex/pkg/dialect/base"
)

// nolint: funlen // Unit test.
func TestRexClass(t *testing.T) {
	test.RexTestCasesSlice{{
		Name: "ClassRangeAndSingle",
		Chain: []dialect.Token{
			base.Common.Class(
				base.Chars.Range('A', 'Z'),
				base.Chars.Single('0'),
			),
		},
		Expected: `[A-Z0]`,
	}, {
		Name: "ClassInClass",
		Chain: []dialect.Token{
			base.Common.Class(
				base.Common.Class(
					base.Chars.Range('A', 'Z'),
				),
				base.Chars.Single('0'),
			),
		},
		Expected: `[A-Z0]`,
	}, {
		Name: "RegularToken",
		Chain: []dialect.Token{
			base.Common.Class(
				base.Common.Raw(`A-Z`),
				base.Chars.Single('0'),
			),
		},
		Expected: `[A-Z0]`,
	}, {
		Name: "ClassInClassInClass",
		Chain: []dialect.Token{
			base.Common.Class(
				base.Common.Class(
					base.Common.Class(
						base.Chars.Digits(),
					),
				),
			),
		},
		Expected: `[\d]`,
	}, {
		Name: "HexDigitsWrapOnce",
		Chain: []dialect.Token{
			base.Common.Class(
				base.Chars.HexDigits(),
			),
		},
		Expected: `[[:xdigit:]]`,
	}, {
		Name: "ClassEmpty",
		Chain: []dialect.Token{
			base.Common.Class(),
		},
		Expected: ``,
	}, {
		Name: "ClassEmptyRepetition",
		Chain: []dialect.Token{
			base.Common.Class().Repeat().OneOrMore(),
		},
		Expected: ``,
	}}.Run(t)
}
