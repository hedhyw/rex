package base_test

import (
	"testing"

	"github.com/hedhyw/rex/internal/test"
	"github.com/hedhyw/rex/pkg/dialect"
	"github.com/hedhyw/rex/pkg/dialect/base"
)

func TestRexCommon(t *testing.T) {
	test.RexTestCasesSlice{{
		Name:     "Class_Any",
		Chain:    []dialect.Token{base.Common.Class(base.Chars.Any())},
		Expected: `[.]`,
	}, {
		Name:     "NotClass_Any",
		Chain:    []dialect.Token{base.Common.NotClass(base.Chars.Any())},
		Expected: `[^.]`,
	}, {
		Name:     "Single_Any",
		Chain:    []dialect.Token{base.Common.Single('.')},
		Expected: `\.`,
	}, {
		Name:     "Raw",
		Chain:    []dialect.Token{base.Common.Raw(`^[A-Z]+$`)},
		Expected: `^[A-Z]+$`,
	}, {
		Name:     "Text",
		Chain:    []dialect.Token{base.Common.Text(`hello world`)},
		Expected: `hello world`,
	}, {
		Name:     "Text_Escaped",
		Chain:    []dialect.Token{base.Common.Text(`^[A-Z]+$`)},
		Expected: `\^\[A-Z\]\+\$`,
	}}.Run(t)
}
