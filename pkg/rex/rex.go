package rex

import (
	"regexp"
	"strings"

	"github.com/hedhyw/rex/internal/helper"
	"github.com/hedhyw/rex/pkg/dialect"
)

// RegExp helps to build regular expressions.
// Use rex.New() for creating.
type RegExp struct {
	expr *strings.Builder
}

// New creates a new RegExp from tokens.
func New(tokens ...dialect.Token) *RegExp {
	b := &RegExp{
		expr: &strings.Builder{},
	}

	_, _ = helper.ProcessTokens(b.expr, tokens)

	return b
}

// String returns a text of the regular expression.
// It can be called multiple times.
// It implements fmt.Stringer interface.
func (r RegExp) String() string {
	return r.expr.String()
}

// Compile parses a regular expression and returns, if successful,
// a Regexp object that can be used to match against text.
func (r RegExp) Compile() (*regexp.Regexp, error) {
	re, err := regexp.Compile(r.String())
	if err != nil {
		return nil, err
	}

	return re, nil
}

// MustCompile is like Compile but panics if the expression cannot be parsed.
// It simplifies safe initialization of global variables holding compiled regular
// expressions.
func (r RegExp) MustCompile() *regexp.Regexp {
	return regexp.MustCompile(r.String())
}
