package base_test

import (
	"testing"

	"github.com/hedhyw/rex/internal/test"
	"github.com/hedhyw/rex/pkg/dialect"
	"github.com/hedhyw/rex/pkg/dialect/base"
)

func TestRexClassRepetitions(t *testing.T) {
	getABClass := func() base.Repetition {
		return base.Common.Class(
			base.Chars.Single('a'),
			base.Chars.Single('b'),
		).Repeat()
	}

	test.RexTestCasesSlice{{
		Name:     "Between",
		Chain:    []dialect.Token{getABClass().Between(0, 1)},
		Expected: `[ab]{0,1}`,
	}, {
		Name:     "EqualOrMoreThan",
		Chain:    []dialect.Token{getABClass().EqualOrMoreThan(5)},
		Expected: `[ab]{5,}`,
	}, {
		Name:     "OneOrMore",
		Chain:    []dialect.Token{getABClass().OneOrMore()},
		Expected: `[ab]+`,
	}, {
		Name:     "OneOrMore",
		Chain:    []dialect.Token{getABClass().ZeroOrMore()},
		Expected: `[ab]*`,
	}, {
		Name:     "ZeroOrOne",
		Chain:    []dialect.Token{getABClass().ZeroOrOne()},
		Expected: `[ab]?`,
	}, {
		Name:     "AnyOneOrMore",
		Chain:    []dialect.Token{base.Chars.Any().Repeat().OneOrMore()},
		Expected: `.+`,
	}, {
		Name:     "RangeOneOrMore",
		Chain:    []dialect.Token{base.Chars.Range('0', '9').Repeat().OneOrMore()},
		Expected: `[0-9]+`,
	}, {
		Name:     "Exactly",
		Chain:    []dialect.Token{base.Chars.Range('0', '9').Repeat().Exactly(2)},
		Expected: `[0-9]{2}`,
	}}.Run(t)
}
