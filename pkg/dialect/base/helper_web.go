package base

import (
	"github.com/hedhyw/rex/pkg/dialect"
)

// HostnameRFC952 is a pattern for a text string drawn from the
// alphabet (A-Z), digits (0-9), minus sign (-), and period (.).
// Periods are only allowed when they serve to delimit components
// of "domain style names". No blank or space characters are permitted
// as part of a name. No distinction is made between upper and lower case.
// The first character must be an alpha character. The last character
// must not be a minus sign or period.
func (HelperDialect) HostnameRFC952() dialect.Token {
	return Group.Define(
		// Cannot start with a number or '-'.
		Chars.Alphabetic(),
		Group.Define(
			Common.Class(
				Chars.Alphanumeric(),
				Chars.Single('-'),
			).Repeat().OneOrMore(),
			Chars.Single('.').Repeat().ZeroOrMore(),
		).NonCaptured().Repeat().ZeroOrMore(),
		// Cannot end with a '-' or '.'.
		Chars.Alphanumeric().Repeat().OneOrMore(),
	).NonCaptured()
}

// HostnameRFC1123 is a pattern like HostnameRFC952, but the restriction
// on the first character is relaxed to allow either a letter or a digit.
// Host software must handle host names of up to 63 characters.
func (HelperDialect) HostnameRFC1123() dialect.Token {
	alphanumericWithMinus := Common.Class(
		Chars.Alphanumeric(),
		Chars.Single('-'),
	)

	return Group.Define(
		Chars.Alphanumeric(),
		alphanumericWithMinus.Repeat().Between(0, 62),
		Group.Define(
			Chars.Single('.').Repeat().ZeroOrMore(),
			Chars.Alphanumeric(),
			alphanumericWithMinus.Repeat().Between(0, 62),
		).NonCaptured().Repeat().ZeroOrMore(),
		Chars.Alphanumeric(),
	).NonCaptured()
}

// Email is a pattern, that checks <local_part>@<host_name>.
//
// Hostname is validated considering RFC-1123.
//
// Localpart is unquoted, and may use any of these ASCII characters:
// - uppercase and lowercase Latin letters A to Z and a to z, digits 0 to 9
// - printable characters !#$%&'*+-/=?^_`{|}~
// - dot ., provided that it is not the first or last character and provided
//   also that it does not appear consecutively (e.g., John..Doe@example.com is not allowed).
func (h HelperDialect) Email() dialect.Token {
	localCharsWithoutDot := Common.Class(
		Chars.Alphanumeric(),
		Chars.Runes("!#$%&'*+-/=?^_`{|}~"),
	)

	unquotedLocalPart := Group.Define(
		// Email must not start with a dot.
		localCharsWithoutDot,
		Group.Define(
			// A dot must not appear consecutively, for this wrap non-dot
			// tokens around.
			Common.Class(
				localCharsWithoutDot,
				Chars.Single('.'),
			).Repeat().ZeroOrOne(),
			localCharsWithoutDot,
		).NonCaptured().Repeat().Between(0, 31),
	).NonCaptured()

	return Group.Define(
		unquotedLocalPart,
		Chars.Single('@'),
		h.HostnameRFC1123(),
	).NonCaptured()
}
