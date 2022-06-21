package base_test

import (
	"testing"

	"github.com/hedhyw/rex/internal/test"
	"github.com/hedhyw/rex/pkg/dialect"
	"github.com/hedhyw/rex/pkg/dialect/base"
)

// nolint: funlen // Unit test.
func TestRexGroup(t *testing.T) {
	test.RexTestCasesSlice{{
		Name: "GroupSingle",
		Chain: []dialect.Token{
			base.Group.Define(base.Chars.Any()),
		},
		Expected: `(.)`,
	}, {
		Name: "GroupMultiple",
		Chain: []dialect.Token{
			base.Group.Define(
				base.Chars.Single('a'),
				base.Chars.Single('b'),
				base.Chars.Single('c'),
			),
		},
		Expected: `(abc)`,
	}, {
		Name: "GroupEmpty",
		Chain: []dialect.Token{
			base.Group.Define(),
		},
		Expected: ``,
	}, {
		Name: "GroupEmptyRepeat",
		Chain: []dialect.Token{
			base.Group.Define().Repeat().OneOrMore(),
		},
		Expected: ``,
	}, {
		Name: "GroupNonCaptured",
		Chain: []dialect.Token{
			base.Group.NonCaptured(base.Chars.Single('a')),
		},
		Expected: `(?:a)`,
	}, {
		Name: "GroupNonCapturedMark",
		Chain: []dialect.Token{
			base.Group.Define(base.Chars.Single('a')).NonCaptured(),
		},
		Expected: `(?:a)`,
	}, {
		Name: "GroupNamed",
		Chain: []dialect.Token{
			base.Group.Define(base.Chars.Single('a')).WithName("my_name"),
		},
		Expected: `(?P<my_name>a)`,
	}, {
		Name: "Repeat_OneOrMore",
		Chain: []dialect.Token{
			base.Group.Define(base.Chars.Single('a')).Repeat().OneOrMore(),
		},
		Expected: `(a)+`,
	}, {
		Name: "Repeat_NonCaptured_ZeroOrMore",
		Chain: []dialect.Token{
			base.Group.Define(
				base.Chars.Single('a'),
			).NonCaptured().Repeat().ZeroOrMore(),
		},
		Expected: `(?:a)*`,
	}}.Run(t)
}
