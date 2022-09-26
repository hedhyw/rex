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

// Alphanumeric specifies digits and alphabetic characters.
// It is an alias to [0-9A-Za-z]. ASCII.
//
// Regex: `[[:alnum:]]`.
func (CharsBaseDialect) Alphanumeric() ClassToken {
	return newClassToken(helper.StringToken(`[:alnum:]`))
}

// Alphabetic specifies alphabetic lowercased and uppercased characters.
// It is an alias to [A-Za-z]. ASCII.
//
// Regex: `[[:alpha:]]`.
func (CharsBaseDialect) Alphabetic() ClassToken {
	return newClassToken(helper.StringToken(`[:alpha:]`))
}

// ASCII only characters. It is an alias to [\x00-\x7F].
//
// Regex: `[[:ascii:]]`.
func (CharsBaseDialect) ASCII() ClassToken {
	return newClassToken(helper.StringToken(`[:ascii:]`))
}

// Whitespace specfies blank characters.
// It is an alias to [\t\n\f\r ]. ASCII.
//
// Regex: `\s`.
func (CharsBaseDialect) Whitespace() ClassToken {
	return newClassToken(helper.StringToken(`\s`)).withoutBrackets()
}

// WordCharacter is an alias to [0-9A-Za-z_]. ASCII.
//
// Regex: `\w`.
func (CharsBaseDialect) WordCharacter() ClassToken {
	return newClassToken(helper.StringToken(`\w`)).withoutBrackets()
}

// Blank ASCII characters. It is an alias to [\t ].
//
// Regex: `[[:blank:]]`.
func (CharsBaseDialect) Blank() ClassToken {
	return newClassToken(helper.StringToken(`[:blank:]`))
}

// Control characters. It is an alias to [\x00-\x1F\x7F]. ASCII.
//
// Regex: `[[:cntrl:]]`.
func (CharsBaseDialect) Control() ClassToken {
	return newClassToken(helper.StringToken(`[:cntrl:]`))
}

// Graphical characters. ASCII.
// It is an alias to [A-Za-z0-9!"#$%&'()*+,\-./:;<=>?@[\\\]^_`{|}~].
//
// Regex: `[[:graph:]]`.
func (CharsBaseDialect) Graphical() ClassToken {
	return newClassToken(helper.StringToken(`[:graph:]`))
}

// Lower cased ASCII characters. It is an alias to [a-z].
//
// Regex: `[[:lower:]]`.
func (CharsBaseDialect) Lower() ClassToken {
	return newClassToken(helper.StringToken(`[:lower:]`))
}

// Printable ASCII characters. It is an alias to [ [:graph:]].
//
// Regex: `[[:print:]]`.
func (CharsBaseDialect) Printable() ClassToken {
	return newClassToken(helper.StringToken(`[:print:]`))
}

// Punctuation ASCII characters. It is an alias to [!-/:-@[-`{-~].
//
// Regex: `[[:punct:]]`.
func (CharsBaseDialect) Punctuation() ClassToken {
	return newClassToken(helper.StringToken(`[:punct:]`))
}

// Upper case ASCII characters. It is an alias to [A-Z].
//
// Regex: `[[:upper:]]`.
func (CharsBaseDialect) Upper() ClassToken {
	return newClassToken(helper.StringToken(`[:upper:]`))
}

// HexDigits ASCII characters. It is an alias to  [0-9A-Fa-f].
//
// Regex: `[[:xdigit:]]`.
func (CharsBaseDialect) HexDigits() ClassToken {
	return newClassToken(helper.StringToken(`[:xdigit:]`))
}

// Begin of text by default or line if the flag EnableMultiline is set.
//
// Regex: `^`.
func (CharsBaseDialect) Begin() ClassToken {
	return newClassToken(helper.ByteToken('^')).withoutBrackets()
}

// Begin of text (even if the flag EnableMultiline is set)
//
// Regex: `\A`.
func (CharsBaseDialect) BeginOfText() ClassToken {
	return newClassToken(helper.StringToken(`\A`)).withoutBrackets()
}

// End of text or line if the flag EnableMultiline is set.
//
// Regex: `$`.
func (CharsBaseDialect) End() ClassToken {
	return newClassToken(helper.ByteToken('$')).withoutBrackets()
}

// A word boundary for ACII words. Following positions count as word boundaries:
//   - Beginning of string: If the first character is an ASCII word character.
//   - End of string: If the last character is an ASCII word character.
//   - Between a word and a non-word character.
//
// Regex: `\b`.
func (CharsBaseDialect) ASCIIWordBoundary() ClassToken {
	return newClassToken(helper.StringToken(`\b`)).withoutBrackets()
}

// Any character, possibly including newline if the flag AnyIncludeNewLine() is set.
//
// Regex: `.`.
func (CharsBaseDialect) Any() ClassToken {
	return newClassToken(helper.ByteToken('.')).withoutBrackets()
}

// Runes create a class that contains defined runes.
// It is safe to pass unicode characters.
//
// Example usage:
//
//	Runes("a") // == Chars.Single('a')
//	Runes("ab") // == Common.Class(Chars.Single('a'), Chars.Single('b'))
//
// Regex: `[abc]`.
func (CharsBaseDialect) Runes(val string) ClassToken {
	// It is not accurate capacity, but enough.
	tokens := make([]dialect.Token, 0, len(val))
	for _, r := range val {
		tokens = append(tokens, Chars.Single(r))
	}

	classToken := newClassToken(tokens...)
	if len(tokens) <= 1 {
		classToken = classToken.withoutBrackets()
	}

	return classToken
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
	// Minus can be a special case in classes.
	if r < unicode.MaxASCII && unicode.IsPrint(r) && r != '-' && r != '%' {
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

// Unicode class. It supports *unicode.RangeTable that is defined
// in `unicode.Categories` or `unicode.Scripts`.
// The input is not validated.
//
// Example usage:
//
//	Chars.Unicode(unicode.Greek)
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
// unicode.Categories and unicode.Scripts.
// The input is not validated.
//
// Example usage:
//
//	Chars.UnicodeByName("Greek")
//
// Regex: `\p{Greek}`.
func (CharsBaseDialect) UnicodeByName(name string) ClassToken {
	return newClassToken(helper.StringToken(`\p{%s}`, name)).withoutBrackets()
}
