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
	return ClassToken{
		classTokens: tokens,
		repetition:  "",
		brackets:    true,
	}
}

// Single specifies the class of a single character.
func (commonBaseDialect) Single(r rune) ClassToken {
	return ClassToken{
		classTokens: []dialect.Token{Chars.Single(r)},
		repetition:  "",
		brackets:    false,
	}
}
