package base

import "github.com/hedhyw/rex/pkg/dialect"

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
		).Repeat().ZeroOrMore(),
		Chars.Alphanumeric(),
	).NonCaptured()
}
