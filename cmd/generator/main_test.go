package main

import (
	"testing"
)

type testCase struct {
	name   string
	regex  string
	result string
}

func Test_generateCode(t *testing.T) {
	var actual, expected, givenRegex string

	tests := getTestData()
	for _, tt := range tests {
		expected = tt.result
		givenRegex = tt.regex
		t.Run(tt.name, func(t *testing.T) {
			actual = generateCode(givenRegex)
			if actual != expected {
				t.Errorf("Expected:\n %s, \nGot:\n %s", expected, actual)
			}
		})
	}
}

func getTestData() [3]testCase {
	return [3]testCase{
		{
			name: "One letter regex", regex: "a",
			result: "rex.New(\n" +
				"	rex.Common.Raw(`a`),\n" +
				")",
		},
		{
			name: "simple regex", regex: "a((\\d+)([a-z]+))",
			result: "rex.New(\n" +
				"	rex.Common.Raw(`a`),\n" +
				"	rex.Group.Define(\n" +
				"		rex.Group.Define(\n" +
				"			rex.Common.Raw(`\\d+`),\n" +
				"		),\n" +
				"		rex.Group.Define(\n" +
				"			rex.Common.Raw(`[a-z]+`),\n" +
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
