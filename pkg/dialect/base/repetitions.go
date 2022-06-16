package base

import (
	"fmt"

	"github.com/hedhyw/rex/internal/helper"
	"github.com/hedhyw/rex/pkg/dialect"
)

// Repetable helps to add repetition suffix.
type Repetable struct {
	token  dialect.Token
	suffix string
}

func newRepetable(token dialect.Token) Repetable {
	return Repetable{
		token:  token,
		suffix: "",
	}
}

// WriteTo implements dialect.Token interface.
func (r Repetable) WriteTo(w dialect.StringByteWriter) (n int, err error) {
	tokens := make([]dialect.Token, 0, 2)
	tokens = append(tokens, r.token)

	if r.suffix != "" {
		tokens = append(tokens, helper.StringToken(r.suffix))
	}

	return helper.ProcessTokens(w, tokens)
}

func (r Repetable) withSuffix(suffix string) Repetable {
	r.suffix = suffix

	return r
}

// OneOrMore repeats one or more, prefer more chars.
//
// Regex: `+`.
func (r Repetable) OneOrMore() Repetable {
	return r.withSuffix("+")
}

// ZeroOrMore repeats zero or more, prefer more chars.
//
// Regex: `*`.
func (r Repetable) ZeroOrMore() Repetable {
	return r.withSuffix("*")
}

// ZeroOrMore repeats zero or one x, prefer one.
//
// Regex: `?`.
func (r Repetable) ZeroOrOne() Repetable {
	return r.withSuffix("?")
}

// EqualOrMoreThan repeats i or i+1 or ... or n, prefer more.
// It doesn't validate an input.
//
// Regex: `{n,}`.
func (r Repetable) EqualOrMoreThan(n int) Repetable {
	return r.withSuffix(fmt.Sprintf("{%d,}", n))
}

// Between repeats i=from or i+1 or ... or to, prefer more.
// It doesn't validate an input.
//
// Regex: `{from,to}`.
func (r Repetable) Between(from, to int) Repetable {
	return r.withSuffix(fmt.Sprintf("{%d,%d}", from, to))
}
