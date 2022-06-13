package rex_test

import (
	"testing"

	"github.com/hedhyw/rex/internal/test"
	"github.com/hedhyw/rex/pkg/dialect"
	"github.com/hedhyw/rex/pkg/rex"
)

func TestRexGeneral(t *testing.T) {
	test.RexTestCasesSlice{{
		Name: "hello world",
		Chain: []dialect.Token{
			rex.Common.Text("hello world"),
		},
		Expected: `hello world`,
	}}.Run(t)
}

func TestRexString(t *testing.T) {
	t.Parallel()

	const expected = "test"

	actual := rex.New(rex.Common.Raw(expected)).String()
	if actual != expected {
		t.Fatalf("Actual: %q, Expected: %q", actual, expected)
	}
}

func TestRexCompile(t *testing.T) {
	t.Parallel()

	t.Run("ok", func(t *testing.T) {
		t.Parallel()

		const expected = "test"

		re, err := rex.New(rex.Common.Raw(expected)).Compile()
		if err != nil {
			t.Fatal(err)
		}

		actual := re.String()
		if actual != expected {
			t.Fatalf("Actual: %q, Expected: %q", actual, expected)
		}
	})

	t.Run("failed", func(t *testing.T) {
		t.Parallel()

		_, err := rex.New(rex.Common.Raw(`[a-`)).Compile()
		if err == nil {
			t.Fatal(err)
		}
	})
}

func TestRexMustCompile(t *testing.T) {
	t.Parallel()

	t.Run("ok", func(t *testing.T) {
		t.Parallel()

		const expected = "test"

		re := rex.New(rex.Common.Raw(expected)).MustCompile()

		actual := re.String()
		if actual != expected {
			t.Fatalf("Actual: %q, Expected: %q", actual, expected)
		}
	})

	t.Run("failed", func(t *testing.T) {
		t.Parallel()

		var recovered interface{}

		func() {
			defer func() { recovered = recover() }()

			_ = rex.New(rex.Common.Raw(`[a-`)).MustCompile()
		}()

		if recovered == nil {
			t.Fatal("Expected panic")
		}
	})
}
