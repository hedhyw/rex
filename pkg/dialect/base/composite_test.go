package base_test

import (
	"testing"

	"github.com/hedhyw/rex/internal/test"
	"github.com/hedhyw/rex/pkg/dialect"
	"github.com/hedhyw/rex/pkg/dialect/base"
)

func TestRexComposite(t *testing.T) {
	test.RexTestCasesSlice{{
		Name: "CompositeSingle",
		Chain: []dialect.Token{
			base.Group.Composite(base.Chars.Any()),
		},
		Expected: `(.)`,
	}, {
		Name: "CompositeMultiple",
		Chain: []dialect.Token{
			base.Group.Composite(
				base.Chars.Single('a'),
				base.Chars.Single('b'),
				base.Chars.Single('c'),
			),
		},
		Expected: `(a|b|c)`,
	}, {
		Name: "CompositeEmpty",
		Chain: []dialect.Token{
			base.Group.Composite(),
		},
		Expected: ``,
	}}.Run(t)
}
