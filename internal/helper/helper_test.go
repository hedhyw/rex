package helper_test

import (
	"errors"
	"strings"
	"testing"

	"github.com/hedhyw/rex/internal/helper"
	"github.com/hedhyw/rex/pkg/dialect"
)

func TestTokenFunc(t *testing.T) {
	t.Parallel()

	const expected = "abc"

	sb := new(strings.Builder)

	n, err := helper.TokenFunc(func(w dialect.StringByteWriter) (int, error) {
		return 1, w.WriteByte('a')
	}).WriteTo(sb)

	switch {
	case err != nil:
		t.Fatal(err)
	case n != 1:
		t.Fatalf("Actual: %d, Expected: %d", n, 1)
	}

	n, err = helper.TokenFunc(func(w dialect.StringByteWriter) (int, error) {
		return w.WriteString("bc")
	}).WriteTo(sb)

	switch {
	case err != nil:
		t.Fatal(err)
	case n != 2:
		t.Fatalf("Actual: %d, Expected: %d", n, 2)
	}

	if actual := sb.String(); expected != actual {
		t.Fatalf("Actual: %q, Expected: %q", actual, expected)
	}
}

func TestStringToken(t *testing.T) {
	t.Parallel()

	const expected = "abc"

	sb := new(strings.Builder)

	n, err := helper.StringToken(expected).WriteTo(sb)

	switch {
	case err != nil:
		t.Fatal(err)
	case n != len(expected):
		t.Fatalf("Actual: %d, Expected: %d", n, len(expected))
	}

	if actual := sb.String(); expected != actual {
		t.Fatalf("Actual: %q, Expected: %q", actual, expected)
	}
}

func TestByteToken(t *testing.T) {
	t.Parallel()

	const expected = 'a'

	sb := new(strings.Builder)

	n, err := helper.ByteToken(expected).WriteTo(sb)

	switch {
	case err != nil:
		t.Fatal(err)
	case n != 1:
		t.Fatalf("Actual: %d, Expected: %d", n, 1)
	}

	if actual := sb.String(); string(expected) != actual {
		t.Fatalf("Actual: %q, Expected: %q", actual, expected)
	}
}

func TestProcessTokens(t *testing.T) {
	t.Parallel()

	const expected = "abc"

	sb := new(strings.Builder)

	n, err := helper.ProcessTokens(sb, []dialect.Token{
		helper.StringToken("ab"),
		helper.ByteToken('c'),
	})

	switch {
	case err != nil:
		t.Fatal(err)
	case n != len(expected):
		t.Fatalf("Actual: %d, Expected: %d", n, len(expected))
	}

	if actual := sb.String(); string(expected) != actual {
		t.Fatalf("Actual: %q, Expected: %q", actual, expected)
	}
}

func TestProcessTokensFailed(t *testing.T) {
	t.Parallel()

	sb := new(strings.Builder)

	_, err := helper.ProcessTokens(sb, []dialect.Token{
		helper.TokenFunc(func(dialect.StringByteWriter) (int, error) {
			// nolint: err113 // Test.
			return 0, errors.New("failed")
		}),
	})
	if err == nil {
		t.Fatal("Expected error")
	}
}
