package test

import (
	"testing"

	"github.com/hedhyw/rex/pkg/dialect"
	"github.com/hedhyw/rex/pkg/rex"
)

// RexTestCase is a general test case.
type RexTestCase struct {
	Name     string
	Chain    []dialect.Token
	Expected string
}

// RexTestCasesSlice helps to process slice of test cases.
type RexTestCasesSlice []RexTestCase

// Run runs in parallel.
func (testCases RexTestCasesSlice) Run(t *testing.T) {
	t.Parallel()

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()

			b := rex.New(tc.Chain...)
			if b.String() != tc.Expected {
				t.Fatalf("Actual: %#q, Expected: %#q", b.String(), tc.Expected)
			}
		})
	}
}
