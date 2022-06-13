package base

import (
	"fmt"

	"github.com/hedhyw/rex/internal/helper"
	"github.com/hedhyw/rex/pkg/dialect"
)

// ClassToken helps to specify class tokens.
type ClassToken struct {
	brackets    bool
	classTokens []dialect.Token
	repetition  string
}

// WriteTo implements dialect.Token interface.
func (ct ClassToken) WriteTo(w dialect.StringByteWriter) (n int, err error) {
	tokens := make([]dialect.Token, 0, 3+len(ct.classTokens))

	if ct.brackets {
		tokens = append(tokens, helper.ByteToken('['))
	}

	tokens = append(tokens, ct.classTokens...)

	if ct.brackets {
		tokens = append(tokens, helper.ByteToken(']'))
	}

	if ct.repetition != "" {
		tokens = append(tokens, helper.StringToken(ct.repetition))
	}

	return helper.ProcessTokens(w, tokens)
}

// OneOrMore repeats one or more, prefer more chars.
//
// Regex: `+`.
func (ct ClassToken) OneOrMore() ClassToken {
	ct.repetition = "+"

	return ct
}

// ZeroOrMore repeats zero or more, prefer more chars.
//
// Regex: `*`.
func (ct ClassToken) ZeroOrMore() ClassToken {
	ct.repetition = "*"

	return ct
}

// ZeroOrMore repeats zero or one x, prefer one.
//
// Regex: `?`.
func (ct ClassToken) ZeroOrOne() ClassToken {
	ct.repetition = "?"

	return ct
}

// EqualOrMoreThan repeats i or i+1 or ... or n, prefer more.
// It doesn't validate an input.
//
// Regex: `{n,}`.
func (ct ClassToken) EqualOrMoreThan(n int) ClassToken {
	ct.repetition = fmt.Sprintf("{%d,}", n)

	return ct
}

// Between repeats i=from or i+1 or ... or to, prefer more.
// It doesn't validate an input.
//
// Regex: `{from,to}`.
func (ct ClassToken) Between(from, to int) ClassToken {
	ct.repetition = fmt.Sprintf("{%d,%d}", from, to)

	return ct
}
