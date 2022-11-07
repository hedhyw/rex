package generator

import (
	"errors"
	"strings"
)

var errUnmachedBraces = errors.New("braces are unmatched")

// GenerateCode returns rex code for a given regex.
func GenerateCode(regex string) (generatedCode string, err error) {
	var (
		beforeCurrentBrace string
		previousRune       rune
		result             strings.Builder
		afterLastBrace     = regex
		indentations       int
		currentOpenBraceI  int
		bracesCounter      int
	)

	result.WriteString("rex.New(\n")

	for i, runeValue := range regex {
		if runeValue == '(' && previousRune != '\\' {
			bracesCounter++

			afterLastBrace = ""
			beforeCurrentBrace = regex[currentOpenBraceI:i]

			addRawExpressionIfNeeded(beforeCurrentBrace, &result, indentations)

			currentOpenBraceI = i + 1
			indentations++
			result.WriteString(strings.Repeat("\t", indentations) + "rex.Group.Define(\n")
		}

		if runeValue == ')' && previousRune != '\\' {
			bracesCounter--

			beforeCurrentBrace = regex[currentOpenBraceI:i]
			currentOpenBraceI = i + 1

			addRawExpressionIfNeeded(beforeCurrentBrace, &result, indentations)

			if i < len(regex)-1 && bracesCounter == 0 {
				afterLastBrace = regex[(i + 1):]
			}

			result.WriteString(strings.Repeat("\t", indentations) + "),\n")
			indentations--
		}

		previousRune = runeValue
	}

	if bracesCounter != 0 {
		err = errUnmachedBraces
	}

	if len(afterLastBrace) != 0 {
		result.WriteString("\trex.Common.Raw(`" + afterLastBrace + "`),\n")
	}

	result.WriteRune(')')

	return result.String(), err
}

func addRawExpressionIfNeeded(beforeCurrentBrace string, result *(strings.Builder), indentations int) {
	if len(beforeCurrentBrace) != 0 {
		(*result).WriteString(strings.Repeat("\t", indentations+1) + "rex.Common.Raw(`" + beforeCurrentBrace + "`),\n")
	}
}
