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

// ZeroOrMore repeats zero or more, prefer more chars.
//
// Regex: `*`.
func (r Repetition) ZeroOrMore() dialect.Token {
	return r.withSuffix("*")
}

// ZeroOrMore repeats zero or one x, prefer one.
//
// Regex: `?`.
func (r Repetition) ZeroOrOne() dialect.Token {
	return r.withSuffix("?")
}

// EqualOrMoreThan repeats i or i+1 or ... or n, prefer more.
// It doesn't validate an input.
//
// Regex: `{n,}`.
func (r Repetition) EqualOrMoreThan(n int) dialect.Token {
	return r.withSuffix(fmt.Sprintf("{%d,}", n))
}

// Between repeats i=from or i+1 or ... or to, prefer more.
// It doesn't validate an input.
//
// Regex: `{from,to}`.
func (r Repetition) Between(from, to int) dialect.Token {
	return r.withSuffix(fmt.Sprintf("{%d,%d}", from, to))
}
