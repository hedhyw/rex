package rex_test

import (
	"testing"

	"github.com/hedhyw/rex/pkg/dialect/base"
	"github.com/hedhyw/rex/pkg/rex"
)

func TestBase(t *testing.T) {
	t.Parallel()

	t.Run("Chars", func(t *testing.T) {
		t.Parallel()

		if rex.Chars != base.Chars {
			t.Fatalf("Actual: %q, Expected: %q", rex.Chars, base.Chars)
		}
	})

	t.Run("Common", func(t *testing.T) {
		t.Parallel()

		if rex.Common != base.Common {
			t.Fatalf("Actual: %q, Expected: %q", rex.Common, base.Common)
		}
	})
}
