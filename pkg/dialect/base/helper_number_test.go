package base_test

import (
	"testing"

	"github.com/hedhyw/rex/pkg/dialect/base"
	"github.com/hedhyw/rex/pkg/rex"
)

func TestNumberRange(t *testing.T) {
	t.Parallel()

	val := rex.New(
		base.Helper.NumberRange(9, 590),
	).String()
	t.Fatal(val)
}
