package generator_test

import (
	"testing"

	"github.com/hedhyw/rex/internal/generator"
)

type generatorTestCase struct {
	name   string
	regex  string
	result string
}

func TestGenerateCodeOK(t *testing.T) {
	t.Parallel()

	testCases := getSuccessGroupTestCases()

	for _, testCaseNotInParallel := range testCases {
		testCase := testCaseNotInParallel

		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			actual, err := generator.GenerateCode(testCase.regex)
			if err != nil {
				t.Fatal(err)
			}

			if actual != testCase.result {
				t.Errorf("Expected:\n%s\nGot:\n%s", testCase.result, actual)
			}
		})
	}
}

func TestGenerateCodeInvalidRegexpr(t *testing.T) {
	t.Parallel()

	_, err := generator.GenerateCode("(")
	if err == nil {
		t.Fatal(err)
	}
}

// nolint: funlen // test cases.
func getSuccessGroupTestCases() []generatorTestCase {
	return []generatorTestCase{{
		name:  "one_letter_regex",
		regex: `a`,
		result: "rex.New(\n" +
			"	rex.Common.Raw(`a`),\n" +
			")",
	}, {
		name:  "uncaptured",
		regex: `(?P<name>1234)`,
		result: "rex.New(\n" +
			"	rex.Group.Define(\n" +
			"		rex.Common.Raw(`1234`),\n" +
			"	).WithName(\"name\"),\n" +
			")",
	}, {
		name:  "concat",
		regex: `(1|12|123)`,
		result: "rex.New(\n" +
			"	rex.Group.Define(\n" +
			"		rex.Common.Raw(`1(?:)|2(?:(?:)|3)`),\n" +
			"	),\n" +
			")",
	}, {
		name:  "simple regex",
		regex: `a((\d+)([a-z]+\())`,
		result: "rex.New(\n" +
			"	rex.Common.Raw(`a`),\n" +
			"	rex.Group.Define(\n" +
			"		rex.Group.Define(\n" +
			"			rex.Common.Raw(`[0-9]+`),\n" +
			"		),\n" +
			"		rex.Group.Define(\n" +
			"			rex.Common.Raw(`[a-z]+\\(`),\n" +
			"		),\n" +
			"	),\n" +
			")",
	}, {
		name: "long_regex",
		regex: "(([0-9]+)([a-z]+))a(([0-9]+)([a-z]+))a" +
			"(([0-9]+)([a-z]+))a",
		result: "rex.New(\n" +
			"	rex.Group.Define(\n" +
			"		rex.Group.Define(\n" +
			"			rex.Common.Raw(`[0-9]+`),\n" +
			"		),\n" +
			"		rex.Group.Define(\n" +
			"			rex.Common.Raw(`[a-z]+`),\n" +
			"		),\n" +
			"	),\n" +
			"	rex.Common.Raw(`a`),\n" +
			"	rex.Group.Define(\n" +
			"		rex.Group.Define(\n" +
			"			rex.Common.Raw(`[0-9]+`),\n" +
			"		),\n" +
			"		rex.Group.Define(\n" +
			"			rex.Common.Raw(`[a-z]+`),\n" +
			"		),\n" +
			"	),\n" +
			"	rex.Common.Raw(`a`),\n" +
			"	rex.Group.Define(\n" +
			"		rex.Group.Define(\n" +
			"			rex.Common.Raw(`[0-9]+`),\n" +
			"		),\n" +
			"		rex.Group.Define(\n" +
			"			rex.Common.Raw(`[a-z]+`),\n" +
			"		),\n" +
			"	),\n" +
			"	rex.Common.Raw(`a`),\n" +
			")",
	}}
}
