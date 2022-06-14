package base

import (
	"regexp"

	"github.com/hedhyw/rex/internal/helper"
	"github.com/hedhyw/rex/pkg/dialect"
)

type commonBaseDialect dialect.Dialect

// Common contains base regular expression helpers.
const Common commonBaseDialect = "commonBaseDialect"

// Raw appends regular expression as is.
func (commonBaseDialect) Raw(raw string) dialect.Token {
	return helper.StringToken(raw)
}

// Text appends the text, and escapes all regular expression metacharacters.
func (commonBaseDialect) Text(text string) dialect.Token {
	return helper.StringToken(regexp.QuoteMeta(text))
}

// Class specifies the class of characters.
func (commonBaseDialect) Class(tokens ...dialect.Token) ClassToken {
	return newClassToken(tokens...)
}

// NotClass specifies the class of characters that should be excluded.
func (commonBaseDialect) NotClass(tokens ...dialect.Token) ClassToken {
	return newClassToken(tokens...).withExclude()
}

// Single specifies the class of a single character.
// It is a synonym to `Chars.Single``.
func (commonBaseDialect) Single(r rune) ClassToken {
	return Chars.Single(r)
}
