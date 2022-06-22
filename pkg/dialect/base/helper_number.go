package base

import (
	"sort"
	"strconv"

	"github.com/hedhyw/rex/pkg/dialect"
)

// 20 is a maximum count of digts in an int64 number.
const numberRangeTokensCapacity = 20

// NumberRange helper.
type NumberRange struct {
	initialFrom int64
	initialTo   int64
}

// NumberRange helps to define a pattern that matches number ranges.
// The arguments from and to can have any order or even be equal.
// Negative numbers are supported.
//
// It doesn't match leading zeros. If you want to match them, use:
//
//   Group.NonCaptured(
//     Chars.Single('0').Repeat().ZeroOrMore(),
//     Helper.NumberRange(0, 99),
//   )
func (h HelperDialect) NumberRange(from int32, to int32) NumberRange {
	return NumberRange{
		initialFrom: int64(from),
		initialTo:   int64(to),
	}
}

func (nr NumberRange) WriteTo(w dialect.StringByteWriter) (n int, err error) {
	return nr.processRange(nr.initialFrom, nr.initialTo).WriteTo(w)
}

func (nr NumberRange) processRange(from, to int64) dialect.Token {
	if from > to {
		to, from = from, to
	}

	switch {
	case to < 0:
		return Group.NonCaptured(
			Chars.Single('-'),
			nr.processRange(-from, -to),
		)
	case from < 0:
		return Group.Composite(
			Group.NonCaptured(
				Chars.Single('-'),
				nr.processRange(0, -from),
			),
			nr.processRange(0, to),
		)
	case from == to:
		return Common.Text(strconv.Itoa(int(from)))
	}

	steps := prepareSteps(from, to)
	pattern := numberRangePatterMaker{
		cachedTokens: make([]dialect.Token, 0, numberRangeTokensCapacity),
	}

	tokens := make([]dialect.Token, 0, numberRangeTokensCapacity)
	for i := 1; i < len(steps); i += 2 {
		tokens = append(tokens, pattern.Make(steps[i-1], steps[i]))
	}

	return Group.Composite(tokens...).NonCaptured()
}

// prepareSteps weakens the numbers boundaries.
func prepareSteps(from, to int64) (steps []int64) {
	steps = make([]int64, 0, numberRangeTokensCapacity)
	numberRange := numberRangeWeakeaner{
		cachedDigits: make([]int64, 0, numberRangeTokensCapacity),
	}

	fromBound := from
	steps = append(steps, fromBound)

	for fromBound < to {
		fromBound = numberRange.Next(fromBound)

		if fromBound < to {
			steps = append(steps, fromBound)
		}

		fromBound++
		if fromBound <= to {
			steps = append(steps, fromBound)
		}
	}

	// Lower bound.
	fromBound = steps[len(steps)-1]
	toBound := to
	steps = append(steps, toBound)

	for toBound > fromBound {
		toBound = numberRange.Prev(toBound)
		if toBound > fromBound {
			steps = append(steps, toBound)
		}

		toBound--
		if toBound >= fromBound {
			steps = append(steps, toBound)
		}
	}

	sort.Slice(steps, func(i, j int) bool {
		return steps[j] > steps[i]
	})

	return steps
}

type numberRangePatterMaker struct {
	// Save tokens between calls.
	cachedTokens []dialect.Token
}

// Make creates a number range pattern.
// Arguments from and to should have the same number of digits.
//
// Example:
// 123-129 -> 12[3-9].
// 100-199 -> 1[0-9][0-9].
func (m numberRangePatterMaker) Make(from, to int64) dialect.Token {
	if from == to {
		return Common.Text(strconv.FormatInt(from, 10))
	}

	tokens := m.cachedTokens[:0]

	for to != 0 {
		digitTo := '0' + rune(to%10)
		digitFrom := '0' + rune(from%10)

		to /= 10
		from /= 10

		if digitTo == digitFrom {
			tokens = append(
				[]dialect.Token{Chars.Single(digitTo)},
				tokens...,
			)

			continue
		}

		tokens = append(
			[]dialect.Token{
				Chars.Range(digitFrom, digitTo),
			},
			tokens...,
		)
	}

	return Group.NonCaptured(tokens...)
}

type numberRangeWeakeaner struct {
	// Save slice between calls.
	cachedDigits []int64
}

// Next finds the next number of the range.
//
// Examples:
// 150 -> 199.
// 199 -> 999.
func (w numberRangeWeakeaner) Next(val int64) int64 {
	return w.weakenNumber(val, 0, 9)
}

// Prev finds the previous number of the range.
//
// Examples:
// 162 -> 160.
// 150 -> 100.
// 149 -> 100.
// 590 -> 589.
func (w numberRangeWeakeaner) Prev(val int64) int64 {
	return w.weakenNumber(val, 9, 0)
}

// weakenNumber replaces first occurrences of fromDigit to toDigit and
// the next digit after it.
//
// Example: val = 150, fromDigit = 0, toDigit = 9
// Then it will return 199.
func (w numberRangeWeakeaner) weakenNumber(val, fromDigit, toDigit int64) (res int64) {
	digitsStack := w.cachedDigits[:0]

	var end bool

	for val > 0 {
		digit := val % 10
		val /= 10

		if !end {
			end = digit != fromDigit
			digit = toDigit
		}

		digitsStack = append(digitsStack, digit)
	}

	for i := len(digitsStack) - 1; i >= 0; i-- {
		res *= 10
		res += digitsStack[i]
	}

	return res
}
