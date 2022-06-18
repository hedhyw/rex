package test

import (
	"testing"

	"github.com/hedhyw/rex/pkg/dialect"
	"github.com/hedhyw/rex/pkg/dialect/base"
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

// MatchTestCaseSlice helps to process slice of test cases.
type MatchTestCaseSlice []MatchTestCase

// WithMatched defines matched for all test cases.
func (tcs MatchTestCaseSlice) WithMatched(matched bool) MatchTestCaseSlice {
	for i := range tcs {
		tcs[i].matched = matched
	}

	return tcs
}

// MatchTestCaseGroupSlice helps to process groups of test cases.
type MatchTestCaseGroupSlice [][]MatchTestCase

// // Run runs in parallel.
func (tcs MatchTestCaseGroupSlice) Run(t *testing.T, tokens ...dialect.Token) {
	t.Parallel()

	re := rex.New(base.Group.Define(
		base.Chars.Begin(),
		base.Group.Define(tokens...).NonCaptured(),
		base.Chars.End(),
	).NonCaptured()).MustCompile()

	for _, g := range tcs {
		for _, tc := range g {
			tc := tc

			t.Run(tc.Name, func(t *testing.T) {
				t.Parallel()

				t.Logf("\nmatching: %#q\nby %#q", tc.Value, re.String())

				actual := re.MatchString(tc.Value)
				if actual != tc.matched {
					t.Fatalf("Actual: %v, Expected: %v", actual, tc.matched)
				}
			})
		}
	}
}

// MatchTestCase used for testing matching.
type MatchTestCase struct {
	Name    string
	Value   string
	matched bool
}
