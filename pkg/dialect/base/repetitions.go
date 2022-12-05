package base

import (
	"fmt"

	"github.com/hedhyw/rex/internal/helper"
	"github.com/hedhyw/rex/pkg/dialect"
)

// Repetition helps to add repetition suffix.
type Repetition struct {
	token  dialect.Token
	suffix string
}

func newRepetition(token dialect.Token) Repetition {
	return Repetition{
		token:  token,
		suffix: "",
	}
}

// WriteTo implements dialect.Token interface.
func (r Repetition) WriteTo(w dialect.StringByteWriter) (n int, err error) {
	tokens := make([]dialect.Token, 0, 2)

	if r.token != nil {
		tokens = append(tokens, r.token)

		if r.suffix != "" {
			tokens = append(tokens, helper.StringToken(r.suffix))
		}
	}

	return helper.ProcessTokens(w, tokens)
}

func (r Repetition) withSuffix(suffix string) dialect.Token {
	r.suffix = suffix

	return r
}

// OneOrMore repeats one or more, prefer more chars.
//
// Regex: `+`.
func (r Repetition) OneOrMore() dialect.Token {
	return r.withSuffix("+")
}

// OneOrMorePreferFewer repeats one or more, prefer fewer chars.
//
// Regex: `+?`.
func (r Repetition) OneOrMorePreferFewer() dialect.Token {
	return r.withSuffix("+?")
}

// ZeroOrMore repeats zero or more, prefer more chars.
//
// Regex: `*`.
func (r Repetition) ZeroOrMore() dialect.Token {
	return r.withSuffix("*")
}

// ZeroOrMorePreferFewer repeats zero or more, prefer fewer chars.
//
// Regex: `*?`.
func (r Repetition) ZeroOrMorePreferFewer() dialect.Token {
	return r.withSuffix("*?")
}

// ZeroOrOne repeats zero or one x, prefer one.
//
// Regex: `?`.
func (r Repetition) ZeroOrOne() dialect.Token {
	return r.withSuffix("?")
}

// ZeroOrOnePreferZero repeats zero or one x, prefer zero.
//
// Regex: `??`.
func (r Repetition) ZeroOrOnePreferZero() dialect.Token {
	return r.withSuffix("??")
}

// Exactly n times. It doesn't validate an input.
//
// Regex: `{n}`.
func (r Repetition) Exactly(n int) dialect.Token {
	return r.withSuffix(fmt.Sprintf("{%d}", n))
}

// EqualOrMoreThan repeats i or i+1 or ... or n, prefer more.
// It doesn't validate an input.
//
// Regex: `{n,}`.
func (r Repetition) EqualOrMoreThan(n int) dialect.Token {
	return r.withSuffix(fmt.Sprintf("{%d,}", n))
}

// EqualOrMoreThanPreferFewer repeats i or i+1 or ... or n, prefer fewer.
// It doesn't validate an input.
//
// Regex: `{n,}?`.
func (r Repetition) EqualOrMoreThanPreferFewer(n int) dialect.Token {
	return r.withSuffix(fmt.Sprintf("{%d,}?", n))
}

// Between repeats i=from or i+1 or ... or to, prefer more.
// It doesn't validate an input.
//
// Regex: `{from,to}`.
func (r Repetition) Between(from, to int) dialect.Token {
	return r.withSuffix(fmt.Sprintf("{%d,%d}", from, to))
}

// BetweenPreferFewer repeats i=from or i+1 or ... or to, prefer fewer.
// It doesn't validate an input.
//
// Regex: `{from,to}?`.
func (r Repetition) BetweenPreferFewer(from, to int) dialect.Token {
	return r.withSuffix(fmt.Sprintf("{%d,%d}?", from, to))
}
