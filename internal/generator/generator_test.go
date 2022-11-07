package generator_test

import (
	"log"
	"testing"

	"github.com/hedhyw/rex/internal/generator"
)

type generatorTestCase struct {
	name   string
	regex  string
	result string
}

func TestGenerateCode(t *testing.T) {
	t.Parallel()

	var actual, expected, givenRegex string

	var err error

	tests := getGroupTestCases()
	for _, testCaseNotInParallel := range tests {
		testCase := testCaseNotInParallel
		expected = testCase.result
		givenRegex = testCase.regex
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			actual, err = generator.GenerateCode(givenRegex)
			if err != nil {
				log.Fatal(err)
			}
			if actual != expected {
				t.Errorf("Expected:\n %s, \nGot:\n %s", expected, actual)
			}
		})
	}
}

func getGroupTestCases() []generatorTestCase {
	return []generatorTestCase{
		{
			name: "one letter regex", regex: "a",
			result: "rex.New(\n" +
				"	rex.Common.Raw(`a`),\n" +
				")",
		},
		{
			name: "simple regex", regex: "a((\\d+)([a-z]+\\()))",
			result: "rex.New(\n" +
				"	rex.Common.Raw(`a`),\n" +
				"	rex.Group.Define(\n" +
				"		rex.Group.Define(\n" +
				"			rex.Common.Raw(`\\d+`),\n" +
				"		),\n" +
				"		rex.Group.Define(\n" +
				"			rex.Common.Raw(`[a-z]+\\(`),\n" +
				"		),\n" +
				"	),\n" +
				")",
		},
		{
			name: "long regex",
			regex: "((\\d+)([a-z]+))a((\\d+)([a-z]+))a" +
				"((\\d+)([a-z]+))a",
			result: "rex.New(\n" +
				"	rex.Group.Define(\n" +
				"		rex.Group.Define(\n" +
				"			rex.Common.Raw(`\\d+`),\n" +
				"		),\n" +
				"		rex.Group.Define(\n" +
				"			rex.Common.Raw(`[a-z]+`),\n" +
				"		),\n" +
				"	),\n" +
				"	rex.Common.Raw(`a`),\n" +
				"	rex.Group.Define(\n" +
				"		rex.Group.Define(\n" +
				"			rex.Common.Raw(`\\d+`),\n" +
				"		),\n" +
				"		rex.Group.Define(\n" +
				"			rex.Common.Raw(`[a-z]+`),\n" +
				"		),\n" +
				"	),\n" +
				"	rex.Common.Raw(`a`),\n" +
				"	rex.Group.Define(\n" +
				"		rex.Group.Define(\n" +
				"			rex.Common.Raw(`\\d+`),\n" +
				"		),\n" +
				"		rex.Group.Define(\n" +
				"			rex.Common.Raw(`[a-z]+`),\n" +
				"		),\n" +
				"	),\n" +
				"	rex.Common.Raw(`a`),\n" +
				")",
		},
	}
}
