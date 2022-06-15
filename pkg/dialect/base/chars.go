package base

import (
	"regexp"
	"strconv"
	"strings"
	"unicode"

	"github.com/hedhyw/rex/internal/helper"
	"github.com/hedhyw/rex/pkg/dialect"
)

// CharsBaseDialect is a namespace that contains common character
// classes tokens.
//
// Use the alias `rex.Chars`.
type CharsBaseDialect dialect.Dialect

// Chars contains character class elements.
const Chars CharsBaseDialect = "CharsBaseDialect"

// Digits is an alias to [0-9]. ASCII.
//
// Regex: `\d`.
func (CharsBaseDialect) Digits() ClassToken {
	return newClassToken(helper.StringToken(`\d`)).withoutBrackets()
}

// Begin of text by default or line if the flag EnableMultiline is set.
//
// Regex: `^`.
func (CharsBaseDialect) Begin() ClassToken {
	return newClassToken(helper.ByteToken('^')).withoutBrackets()
}

// End of text or line if the flag EnableMultiline is set.
//
// Regex: `$`.
func (CharsBaseDialect) End() ClassToken {
	return newClassToken(helper.ByteToken('$')).withoutBrackets()
}

// Any character, possibly including newline if the flag AnyIncludeNewLine() is set.
//
// Regex: `.`.
func (CharsBaseDialect) Any() ClassToken {
	return newClassToken(helper.ByteToken('.')).withoutBrackets()
}

// Range of characters.
// The input is not validated.
//
// Regex: `[a-z]`.
func (CharsBaseDialect) Range(from rune, to rune) ClassToken {
	return newClassToken(helper.StringToken("%c-%c", from, to))
}

// Single character. It supports not ascii characters.
// The input is not validated.
//
// Regex: `r`, `\\xHEX_CODE`, or  `\\x{HEX_CODE}`.
func (CharsBaseDialect) Single(r rune) ClassToken {
	if r < unicode.MaxASCII {
		return newClassToken(
			helper.StringToken(regexp.QuoteMeta(string(r))),
		).withoutBrackets()
	}

	hexValue := strings.ToUpper(strconv.FormatInt(int64(r), 16))

	if len(hexValue) == 2 {
		return newClassToken(
			helper.StringToken("\\x" + hexValue),
		).withoutBrackets()
	}

	return newClassToken(
		helper.StringToken("\\x{%s}", hexValue),
	).withoutBrackets()
}

// Unicode class. It supports only unicode.* range tables that are defined
// in `unicode.Categories` or `unicode.Scripts`. The input is not validated.
//
// Example usage:
//
//   Chars.Unicode(unicode.Greek)
//
// Regex: `\p{Greek}`.
func (d CharsBaseDialect) Unicode(table *unicode.RangeTable) ClassToken {
	for name, t := range unicode.Categories {
		if table == t {
			return d.UnicodeByName(name)
		}
	}

	for name, t := range unicode.Scripts {
		if table == t {
			return d.UnicodeByName(name)
		}
	}

	return newClassToken(helper.StringToken("")).withoutBrackets()
}

// UnicodeByName class. It is alternative to Chars.Unicode, but accepts
// name of the RangeTable. Unicode character classes are those in
// unicode.Categories and unicode.Scripts. The input is not validated.
//
// Example usage:
//
//   Chars.UnicodeByName("Greek")
//
// Regex: `\p{Greek}`.
func (CharsBaseDialect) UnicodeByName(name string) ClassToken {
	return newClassToken(helper.StringToken(`\p{%s}`, name)).withoutBrackets()
}
