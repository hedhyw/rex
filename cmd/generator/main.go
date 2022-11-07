package main

import (
	"os"
)

func main() {
	argsWithProg := os.Args

	if len(argsWithProg) != 2 {
		panic("Wrong amount of arguments!")
	}

	if len(argsWithProg[1]) == 0 {
		panic("Given regex is empty!")
	}

	regex := argsWithProg[1]

	_, err := os.Stdout.WriteString((generateCode(regex) + "\n"))
	if err != nil {
		panic("Error writing to stdout")
	}
}

func generateCode(regex string) string {
	afterLastBrace, beforeCurrentBrace, indentations, result := regex, "", "", "rex.New(\n"

	currentOpenBraceIndex, bracesCounter := 0, 0

	for index, runeValue := range regex {
		if runeValue == '(' {
			bracesCounter++

			afterLastBrace = ""
			beforeCurrentBrace = regex[currentOpenBraceIndex:index]

			if len(beforeCurrentBrace) != 0 {
				indentations += "\t"
				result = result + indentations + "rex.Common.Raw(`" + beforeCurrentBrace + "`),\n"
				indentations = indentations[:(len(indentations) - 1)]
			}

			indentations += "\t"
			result += indentations + "rex.Group.Define(\n"
			currentOpenBraceIndex = index + 1
		}

		if runeValue == ')' {
			bracesCounter--

			beforeCurrentBrace = regex[currentOpenBraceIndex:index]
			if len(beforeCurrentBrace) != 0 {
				indentations += "\t"
				result = result + indentations + "rex.Common.Raw(`" + beforeCurrentBrace + "`),\n"
				indentations = indentations[:(len(indentations) - 1)]
			}

			if index < len(regex)-1 && bracesCounter == 0 {
				afterLastBrace = regex[(index + 1):]
			}

			result += indentations + "),\n"
			indentations = indentations[:(len(indentations) - 1)]
			currentOpenBraceIndex = index + 1
		}
	}

	if bracesCounter != 0 {
		panic("Braces aren't mached!")
	}

	if len(afterLastBrace) != 0 {
		result += "	rex.Common.Raw(`" + afterLastBrace + "`),\n"
	}

	return result + ")"
}
