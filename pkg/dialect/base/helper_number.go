package base

import (
	"fmt"
	"strconv"

	"github.com/hedhyw/rex/pkg/dialect"
)

// NumberRange helps to define a pattern that matches number ranges.
// from and to can have any order or even be equal.
//
// Negative numbers are supported.
func (HelperDialect) NumberRange(from int, to int) dialect.Token {
	if from > to {
		to, from = from, to
	}

	if from == to {
		return Common.Text(strconv.Itoa(from))
	}

	tokens := make([]dialect.Token, 0, 20)

	// For 0-593, it is
	// 0-9
	// 10-99
	// 100-499
	// 500-589
	// 590-593
	for from < to {
		upperBound, ok := nextRangeNumberUpperBound(from, to)
		fmt.Println(from, upperBound)
		tokens = append(tokens, numberRangePattern(from, upperBound))
		from = upperBound + 1

		if !ok {
			break
		}
	}

	// Handle upper special case.

	return Group.Composite(tokens...).NonCaptured()
}

// nextRangeNumberUpperBound returns the next upper bound.
//
// For 0-590, it is
// first call:   0-9
// second call:  10-99
// third call:   100-499
func nextRangeNumberUpperBound(from, to int) (nextNumber int, ok bool) {
	originalFrom := from

	if from >= to {
		return to, false
	}

	for from > 0 {
		from /= 10
		nextNumber *= 10
		nextNumber += 9
	}

	if nextNumber == 0 {
		nextNumber = 9
	}

	if nextNumber > to {
		if originalFrom == 0 {
			return to, false
		}

		// Floor to lower.
		// 590 -> 499.
		return (to/originalFrom)*originalFrom - 1, false
	}

	return nextNumber, true
}

// numberRangePattern create a number range pattern.
// from and to should have the same number of digits.
func numberRangePattern(from, to int) dialect.Token {
	if from == to {
		return Common.Text(strconv.Itoa(from))
	}

	// 20 is a maximum count of digts in int64 number.
	tokens := make([]dialect.Token, 0, 20)

	for to != 0 {
		digitTo := '0' + rune(to%10)
		digitFrom := '0' + rune(from%10)

		tokens = append(
			[]dialect.Token{
				Chars.Range(digitFrom, digitTo),
			},
			tokens...,
		)

		to /= 10
		from /= 10
	}

	return Group.Define(tokens...).NonCaptured()
}
