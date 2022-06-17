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
