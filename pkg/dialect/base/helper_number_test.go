package base_test

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"testing"

	"github.com/hedhyw/rex/pkg/dialect/base"
	"github.com/hedhyw/rex/pkg/rex"
)

func TestNumberRange_broot(t *testing.T) {
	t.Parallel()

	for i := -100; i < 100; i++ {
		for j := i; j < 100; j++ {
			// nolint: gosec // It's a test, and range is defined in the loop.
			newNumberRangeTestCase(int32(i), int32(j)).Run(t, 100)
		}
	}
}

func TestNumberRange_base(t *testing.T) {
	t.Parallel()

	newNumberRangeTestCase(50, 1230).Run(t, 100)
	newNumberRangeTestCase(250, 250).Run(t, 100)
	newNumberRangeTestCase(250, 255).Run(t, 100)
	newNumberRangeTestCase(999, 1000).Run(t, 100)
	newNumberRangeTestCase(1001, 1768).Run(t, 100)
	newNumberRangeTestCase(-123, 456).Run(t, 100)
	newNumberRangeTestCase(-456, 123).Run(t, 100)
	newNumberRangeTestCase(-456, -100).Run(t, 100)
	newNumberRangeTestCase(math.MaxInt32-20, math.MaxInt32-10).Run(t, 1000)
}

func TestNumberRange_leadingZeros(t *testing.T) {
	t.Parallel()

	reNoLeading := rex.New(
		base.Chars.Begin(),
		base.Helper.NumberRange(1, 100),
		base.Chars.End(),
	).MustCompile()

	reLeading := rex.New(
		base.Chars.Begin(),
		base.Group.NonCaptured(
			base.Chars.Single('0').Repeat().ZeroOrMore(),
			base.Helper.NumberRange(1, 100),
		),
		base.Chars.End(),
	).MustCompile()

	switch {
	case !reNoLeading.MatchString("1"):
		t.Fatal("no leading: 1")
	case reNoLeading.MatchString("001"):
		t.Fatal("no leading: 001")
	case reNoLeading.MatchString("101"):
		t.Fatal("no leading: 101")
	case !reLeading.MatchString("099"):
		t.Fatal("leading 099")
	case !reLeading.MatchString("000099"):
		t.Fatal("leading 000099")
	case !reLeading.MatchString("99"):
		t.Fatal("leading: 99")
	case reLeading.MatchString("101"):
		t.Fatal("leading: 101")
	}
}

type numberRangeTestCase struct {
	from int32
	to   int32
	re   *regexp.Regexp
}

func newNumberRangeTestCase(from, to int32) *numberRangeTestCase {
	if from > to {
		to, from = from, to
	}

	return &numberRangeTestCase{
		from: from,
		to:   to,
		re: rex.New(
			base.Chars.Begin(),
			base.Helper.NumberRange(from, to),
			base.Chars.End(),
		).MustCompile(),
	}
}

func (tc numberRangeTestCase) Run(t *testing.T, threshold int64) {
	t.Run(fmt.Sprintf("from_%d_to_%d", tc.from, tc.to), func(t *testing.T) {
		t.Parallel()

		t.Log("re is ", tc.re.String())

		for i := int64(tc.from) - threshold; i <= int64(tc.to)+threshold; i++ {
			tc.assert(t, i)
		}
	})
}

func (tc numberRangeTestCase) assert(tb testing.TB, n int64) {
	tb.Helper()

	expected := n >= int64(tc.from) && n <= int64(tc.to)
	actual := tc.re.MatchString(strconv.FormatInt(n, 10))

	if expected != actual {
		tb.Fatalf(
			"Actual: %t, Expected: %t (%d in %d - %d)",
			actual,
			expected,
			n,
			tc.from,
			tc.to,
		)
	}
}

func FuzzRangeNumber(f *testing.F) {
	f.Add(int32(0), int32(100), int64(50))
	f.Add(int32(0), int32(100), int64(-1))
	f.Add(int32(0), int32(100), int64(101))

	f.Fuzz(func(t *testing.T, from int32, to int32, num int64) {
		newNumberRangeTestCase(from, to).assert(t, num)
	})
}
