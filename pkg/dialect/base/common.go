package base

import (
	"regexp"

	"github.com/hedhyw/rex/internal/helper"
	"github.com/hedhyw/rex/pkg/dialect"
)

// CommonBaseDialect is a namespace that contains common operations.
//
// Use the alias `rex.Common`.
type CommonBaseDialect dialect.Dialect

// Common contains base regular expression helpers.
const Common CommonBaseDialect = "CommonBaseDialect"

// Raw appends regular expression as is.
func (CommonBaseDialect) Raw(raw string) dialect.Token {
	return helper.StringToken(raw)
}

// Text appends the text, and escapes all regular expression metacharacters.
func (CommonBaseDialect) Text(text string) dialect.Token {
	return helper.StringToken(regexp.QuoteMeta(text))
}

// Class specifies the class of characters.
func (CommonBaseDialect) Class(tokens ...dialect.Token) ClassToken {
	return newClassToken(tokens...)
}

// NotClass specifies the class of characters that should be excluded.
func (CommonBaseDialect) NotClass(tokens ...dialect.Token) ClassToken {
	return newClassToken(tokens...).withExclude()
}

// Single specifies the class of a single character.
// It is a synonym to `Chars.Single``.
func (CommonBaseDialect) Single(r rune) ClassToken {
	return Chars.Single(r)
}
