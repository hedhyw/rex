package base

import (
	"github.com/hedhyw/rex/pkg/dialect"
)

// Phone contains composite of different phone patterns: E.164, E.123.
//
// Examples:
//   +15555555
//   (607) 123 4567
//   +22 607 123 4567
func (h HelperDialect) Phone() dialect.Token {
	return Group.Composite(
		h.PhoneE164(),
		h.PhoneE123(),
	)
}

// PhoneE164 is a patter for E.164, that is the international telephone
// numbering plan that ensures each device on the PSTN has globally
// unique number. This number allows phone calls and text messages can
// be correctly routed to individual phones in different countries.
// E.164 numbers are formatted [+] [country code] [subscriber number
// including area code] and can have a maximum of fifteen digits.
//
// Example: +15555555.
func (HelperDialect) PhoneE164() dialect.Token {
	return Group.NonCaptured(
		// Country code.
		Chars.Single('+'),
		// There is no country code that starts with zero.
		Chars.Range('1', '9'),
		// 7...14 digits.
		Chars.Digits().Repeat().Between(7, 14),
	)
}

// PhoneE123 is a patter for E.123, it is an international standard by
// the standardization union (ITU-T), entitled Notation for national
// and international telephone numbers, e-mail addresses and Web addresses.
//
// It combines international and national formats.
//
// Examples:
//   (607) 123 4567
//   +22 607 123 4567
func (h HelperDialect) PhoneE123() dialect.Token {
	return Group.Composite(
		h.PhoneNationalE123(),
		h.PhoneInternationalE123(),
	)
}

// PhoneNationalE123 is a patter for telephone number of E.123
// international notation.
//
// It matches the standard national US 3-3-4 number format. A 3 digit
// area code in brackets, followed by a space, then 3 more digits,
// followed by another space then 4 digits.
//
// Example: (607) 123 4567.
func (HelperDialect) PhoneNationalE123() dialect.Token {
	return Group.Define(
		// Area code.
		Chars.Single('('),
		Chars.Digits().Repeat().Exactly(3),
		Chars.Single(')'),

		Chars.Whitespace(),
		Chars.Digits().Repeat().Exactly(3),

		Chars.Whitespace(),
		Chars.Digits().Repeat().Exactly(4),
	).NonCaptured()
}

// PhoneInternationalE123 is a patter for telephone number of E.123
// international notation.
//
// It matches a 1 to 3 digit country code followed by a space,
// then 2 to 3 digits, another space, another 2 to 3 digits a final
// space then 4 digits.
//
// Example: +22 607 123 4567.
func (HelperDialect) PhoneInternationalE123() dialect.Token {
	return Group.NonCaptured(
		// Country code.
		Chars.Single('+'),
		// There is no country code that starts with zero.
		Chars.Range('1', '9'),
		Chars.Digits().Repeat().Between(0, 2),

		Chars.Whitespace(),
		Chars.Digits().Repeat().Between(2, 3),

		Chars.Whitespace(),
		Chars.Digits().Repeat().Between(2, 3),

		Chars.Whitespace(),
		Chars.Digits().Repeat().Exactly(4),
	)
}
