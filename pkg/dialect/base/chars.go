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
func (CharsBaseDialect) Digits() RepetableClassToken {
	return newRepetableClassToken(
		newClassToken(helper.StringToken(`\d`)).withoutBrackets(),
	)
}

// Begin of text by default or line if the flag EnableMultiline is set.
//
// Regex: `^`.
func (CharsBaseDialect) Begin() RepetableClassToken {
	return newRepetableClassToken(
		newClassToken(helper.ByteToken('^')).withoutBrackets(),
	)
}

// End of text or line if the flag EnableMultiline is set.
//
// Regex: `$`.
func (CharsBaseDialect) End() RepetableClassToken {
	return newRepetableClassToken(
		newClassToken(helper.ByteToken('$')).withoutBrackets(),
	)
}

// Any character, possibly including newline if the flag AnyIncludeNewLine() is set.
//
// Regex: `.`.
func (CharsBaseDialect) Any() RepetableClassToken {
	return newRepetableClassToken(
		newClassToken(helper.ByteToken('.')).withoutBrackets(),
	)
}

// Runes create a class that contains defined runes.
// It is safe to pass unicode characters.
//
// Example usage:
//   Runes("a") // == Chars.Single('a')
//   Runes("ab") // == Common.Class(Chars.Single('a'), Chars.Single('b'))
//
// Regex: `[abc]`.
func (CharsBaseDialect) Runes(val string) RepetableClassToken {
	// It is not accurate capacity, but enough.
	tokens := make([]dialect.Token, 0, len(val))
	for _, r := range val {
		tokens = append(tokens, Chars.Single(r))
	}

	classToken := newClassToken(tokens...)
	if len(tokens) <= 1 {
		classToken = classToken.withoutBrackets()
	}

	return newRepetableClassToken(classToken)
}

// Range of characters.
// The input is not validated.
//
// Regex: `[a-z]`.
func (CharsBaseDialect) Range(from rune, to rune) RepetableClassToken {
	return newRepetableClassToken(
		newClassToken(helper.StringToken("%c-%c", from, to)),
	)
}

// Single character. It supports not ascii characters.
// The input is not validated.
//
// Regex: `r`, `\\xHEX_CODE`, or  `\\x{HEX_CODE}`.
func (CharsBaseDialect) Single(r rune) RepetableClassToken {
	if r < unicode.MaxASCII {
		return newRepetableClassToken(newClassToken(
			helper.StringToken(regexp.QuoteMeta(string(r))),
		).withoutBrackets())
	}

	hexValue := strings.ToUpper(strconv.FormatInt(int64(r), 16))

	if len(hexValue) == 2 {
		return newRepetableClassToken(newClassToken(
			helper.StringToken("\\x" + hexValue),
		).withoutBrackets())
	}

	return newRepetableClassToken(newClassToken(
		helper.StringToken("\\x{%s}", hexValue),
	).withoutBrackets())
}

// Unicode class. It supports *unicode.RangeTable that is defined
// in `unicode.Categories` or `unicode.Scripts`.
// The input is not validated.
//
// Example usage:
//
//   Chars.Unicode(unicode.Greek)
//
// Regex: `\p{Greek}`.
func (d CharsBaseDialect) Unicode(table *unicode.RangeTable) RepetableClassToken {
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

	return newRepetableClassToken(
		newClassToken(helper.StringToken("")).withoutBrackets(),
	)
}

// UnicodeByName class. It is alternative to Chars.Unicode, but accepts
// name of the RangeTable. Unicode character classes are those in
// unicode.Categories and unicode.Scripts.
// The input is not validated.
//
// Example usage:
//
//   Chars.UnicodeByName("Greek")
//
// Regex: `\p{Greek}`.
func (CharsBaseDialect) UnicodeByName(name string) RepetableClassToken {
	return newRepetableClassToken(
		newClassToken(helper.StringToken(`\p{%s}`, name)).withoutBrackets(),
	)
}
