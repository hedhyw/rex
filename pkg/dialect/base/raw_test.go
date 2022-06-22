package base_test

import (
	"testing"

	"github.com/hedhyw/rex/internal/test"
	"github.com/hedhyw/rex/pkg/dialect"
	"github.com/hedhyw/rex/pkg/dialect/base"
)

func TestRexRaw(t *testing.T) {
	test.RexTestCasesSlice{{
		Name:     "RawToken",
		Chain:    []dialect.Token{base.Common.Raw(`.+`)},
		Expected: `.+`,
	}, {
		Name:     "RawClass",
		Chain:    []dialect.Token{base.Common.Class(base.Common.Raw("A-Z"))},
		Expected: `[A-Z]`,
	}}.Run(t)
}

func TestRexRawVerbose(t *testing.T) {
	test.RexTestCasesSlice{{
		Name: "RawVerboseRegular",
		Chain: []dialect.Token{
			base.Common.RawVerbose(`.+`),
		},
		Expected: `.+`,
	}, {
		Name: "RawVerboseComment",
		Chain: []dialect.Token{
			base.Common.RawVerbose(`.+ # any character`),
		},
		Expected: `.+`,
	}, {
		Name: "RawVerboseMultiline",
		Chain: []dialect.Token{
			base.Common.RawVerbose(`
			.+ # any character
			\d+ # Digits
			`),
		},
		Expected: `.+\d+`,
	}, {
		Name: "RawVerboseEscapedHashSign",
		Chain: []dialect.Token{
			base.Common.RawVerbose(`\#\d+ # Digits`),
		},
		Expected: `#\d+`,
	}, {
		Name: "RawVerboseEscapedHashInClass",
		Chain: []dialect.Token{
			base.Common.RawVerbose(`[#]\d+ # Digits`),
		},
		Expected: `[#]\d+`,
	}, {
		Name: "RawVerboseEscapedRegularExpressionInComment",
		Chain: []dialect.Token{
			base.Common.RawVerbose(`[#]\d+ # [Hi].+`),
		},
		Expected: `[#]\d+`,
	}, {
		Name: "RawVerboseEscapedNoCommentMultiline",
		Chain: []dialect.Token{
			base.Common.RawVerbose(`
			.+
			\d+
			`),
		},
		Expected: `.+\d+`,
	}}.Run(t)
}
