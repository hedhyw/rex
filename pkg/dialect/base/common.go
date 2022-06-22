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
func (CommonBaseDialect) Raw(raw string) RawToken {
	return RawToken{value: raw, verbose: true}
}

// RawVerbose appends regular expression and removes comments, spaces and new lines.
//
// When a line contains a # that is not in a character class and is not
// preceded by an unescaped backslash, all characters from the leftmost
// such # through the end of the line are ignored.
//
// Leading and trailling whitespaces are ignored.
//
// It also ignores all new lines.
func (CommonBaseDialect) RawVerbose(raw string) RawToken {
	return RawToken{value: raw, verbose: true}
}

// Text appends the text, and escapes all regular expression metacharacters.
func (CommonBaseDialect) Text(text string) dialect.Token {
	return helper.StringToken(regexp.QuoteMeta(text))
}

// Class specifies the class of characters.
func (CommonBaseDialect) Class(tokens ...dialect.ClassToken) ClassToken {
	return newClassToken(unwrapClassTokens(tokens)...)
}

// NotClass specifies the class of characters that should be excluded.
func (CommonBaseDialect) NotClass(tokens ...dialect.ClassToken) ClassToken {
	return newClassToken(unwrapClassTokens(tokens)...).withExclude()
}

func unwrapClassTokens(classTokens []dialect.ClassToken) []dialect.Token {
	tokens := make([]dialect.Token, 0, len(classTokens))

	for _, ct := range classTokens {
		tokens = append(tokens, ct.Unwrap())
	}

	return tokens
}
