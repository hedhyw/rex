package base

import (
	"regexp"
	"strconv"
	"strings"
	"unicode"

	"github.com/hedhyw/rex/internal/helper"
	"github.com/hedhyw/rex/pkg/dialect"
)

type charsBaseDialect dialect.Dialect

// Chars contains character class elements.
const Chars charsBaseDialect = "charsBaseDialect"

// Digits is an alias to [0-9].
func (charsBaseDialect) Digits() dialect.Token {
	return helper.StringToken(`\d`)
}

// Begin of text by default or line if the flag EnableMultiline is set.
func (charsBaseDialect) Begin() dialect.Token {
	return helper.ByteToken('^')
}

// End of text or line if the flag EnableMultiline is set.
func (charsBaseDialect) End() dialect.Token {
	return helper.ByteToken('$')
}

// Any character, possibly including newline if the flag AnyIncludeNewLine() is set.
func (charsBaseDialect) Any() dialect.Token {
	return helper.ByteToken('.')
}

// Single character. It supports not ascii characters.
func (charsBaseDialect) Single(r rune) dialect.Token {
	if r < unicode.MaxASCII {
		return helper.StringToken(regexp.QuoteMeta(string([]rune{r})))
	}

	hexValue := strings.ToUpper(strconv.FormatInt(int64(r), 16))

	if len(hexValue) == 2 {
		return helper.StringToken("\\x" + hexValue)
	}

	return helper.StringToken("\\x{%s}", hexValue)
}
