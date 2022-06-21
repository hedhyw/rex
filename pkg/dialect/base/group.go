package base

import (
	"regexp"

	"github.com/hedhyw/rex/internal/helper"
	"github.com/hedhyw/rex/pkg/dialect"
)

// GroupBaseDialect is a namespace that contains group helpers.
//
// Use the alias `rex.Group`.
type GroupBaseDialect dialect.Dialect

// Define a group with ranges of expressions.
func (GroupBaseDialect) Define(tokens ...dialect.Token) GroupToken {
	return GroupToken{
		prefix: "",
		tokens: tokens,
	}
}

// NonCaptured defines a group as a non-captured. It is a synonym to
// Group.Define(...).NonCaptured().
func (g GroupBaseDialect) NonCaptured(tokens ...dialect.Token) GroupToken {
	return g.Define(tokens...).NonCaptured()
}

// Composite defines logical "OR" between tokens. It can be used for
// matching one of given expression. It creates non-captured group.
func (GroupBaseDialect) Composite(tokens ...dialect.Token) GroupToken {
	if len(tokens) <= 1 {
		return Group.Define(tokens...)
	}

	return GroupToken{
		prefix: "",
		tokens: []dialect.Token{
			CompositToken{tokens: tokens},
		},
	}
}

// Group helps to define groups.
const Group GroupBaseDialect = "GroupBaseDialect"

// GroupToken defines a token that wraps a range of tokens with a `(...)`.
type GroupToken struct {
	prefix string
	tokens []dialect.Token
}

// WriteTo implements dialect.Token interface.
func (gt GroupToken) WriteTo(w dialect.StringByteWriter) (n int, err error) {
	if len(gt.tokens) == 0 {
		return 0, nil
	}

	tokens := make([]dialect.Token, 0, 3+len(gt.tokens))

	tokens = append(tokens, helper.ByteToken('('))
	if gt.prefix != "" {
		tokens = append(tokens, helper.StringToken(gt.prefix))
	}

	tokens = append(tokens, gt.tokens...)
	tokens = append(tokens, helper.ByteToken(')'))

	return helper.ProcessTokens(w, tokens)
}

// WithName add a name to captured group.
//
// It overrides non-captured if set.
func (gt GroupToken) WithName(name string) GroupToken {
	gt.prefix = "?P<" + regexp.QuoteMeta(name) + ">"

	return gt
}

// NonCaptured marks group as non-captured. It means that this group
// will not be included in group submatch results. Use this for optimizing
// regular expressions if you don't need captured result.
//
// It overrides name if set.
func (gt GroupToken) NonCaptured() GroupToken {
	gt.prefix = "?:"

	return gt
}

// Repeat group.
func (gt GroupToken) Repeat() Repetition {
	if len(gt.tokens) == 0 {
		return newRepetition(nil)
	}

	return newRepetition(gt)
}
