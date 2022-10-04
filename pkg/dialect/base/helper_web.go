// nolint: funlen // Complex patterns.
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
	return Group.NonCaptured(
		// Cannot start with a number or '-'.
		Chars.Alphabetic(),
		Group.NonCaptured(
			Common.Class(
				Chars.Alphanumeric(),
				Chars.Single('-'),
			).Repeat().OneOrMore(),
			Chars.Single('.').Repeat().ZeroOrMore(),
		).Repeat().ZeroOrMore(),
		// Cannot end with a '-' or '.'.
		Chars.Alphanumeric().Repeat().OneOrMore(),
	)
}

// HostnameRFC1123 is a pattern like HostnameRFC952, but the restriction
// on the first character is relaxed to allow either a letter or a digit.
// Host software must handle host names of up to 63 characters.
func (HelperDialect) HostnameRFC1123() dialect.Token {
	alphanumericWithMinus := Common.Class(
		Chars.Alphanumeric(),
		Chars.Single('-'),
	)

	return Group.NonCaptured(
		Chars.Alphanumeric(),
		alphanumericWithMinus.Repeat().Between(0, 62),
		Group.NonCaptured(
			Chars.Single('.').Repeat().ZeroOrMore(),
			Chars.Alphanumeric(),
			alphanumericWithMinus.Repeat().Between(0, 62),
		).Repeat().ZeroOrMore(),
		Chars.Alphanumeric(),
	)
}

// Email is a pattern, that checks <local_part>@<host_name>.
//
// Hostname is validated considering RFC-1123.
//
// Localpart is unquoted, and may use any of these ASCII characters:
//   - uppercase and lowercase Latin letters A to Z and a to z, digits 0 to 9
//   - printable characters !#$%&'*+-/=?^_`{|}~
//   - dot ., provided that it is not the first or last character and provided
//     also that it does not appear consecutively (e.g., John..Doe@example.com is not allowed).
func (h HelperDialect) Email() dialect.Token {
	localCharsWithoutDot := Common.Class(
		Chars.Alphanumeric(),
		Chars.Runes("!#$%&'*+-/=?^_`{|}~"),
	)

	unquotedLocalPart := Group.NonCaptured(
		// Email must not start with a dot.
		localCharsWithoutDot,
		Group.NonCaptured(
			// A dot must not appear consecutively, for this wrap non-dot
			// tokens around.
			Common.Class(
				localCharsWithoutDot,
				Chars.Single('.'),
			).Repeat().ZeroOrOne(),
			localCharsWithoutDot,
		).Repeat().Between(0, 31),
	)

	return Group.NonCaptured(
		unquotedLocalPart,
		Chars.Single('@'),
		h.HostnameRFC1123(),
	)
}

// IP is a pattern for IPv4 or IPv6.
func (h HelperDialect) IP() dialect.Token {
	return Group.Composite(h.IPv4(), h.IPv6()).NonCaptured()
}

// IPv4 is a pattern for an IPv4 address that has the following format:
// x.x.x.x where x is called an octet and must be a decimal value between
// 0 and 255.
//
// Example: 127.0.0.1.
func (HelperDialect) IPv4() dialect.Token {
	ipv4Octet := Helper.NumberRange(0, 255)

	return Group.NonCaptured(
		Group.NonCaptured(
			ipv4Octet,
			// Numbers are divided by a dot.
			Chars.Single('.'),
		).Repeat().Exactly(3),
		ipv4Octet,
	)
}

// IPv6 is a pattern for IPv6 (Normal) address that has the following format:
// y:y:y:y:y:y:y:y where y is called a segment and can be any hexadecimal
// value between 0 and FFFF. The segments are separated by colons - not periods.
// Zero segments can be skipped. It can also match an IPv6 (Dual) address,
// that combines an IPv6 and an IPv4.
func (h HelperDialect) IPv6() dialect.Token {
	ipv6Segment := Chars.HexDigits().Repeat().Between(1, 4)
	delimeter := Chars.Single(':')
	ipv6SegmentDelimeter := Group.NonCaptured(
		ipv6Segment,
		delimeter,
	)
	delimeterIPv6Segment := Group.NonCaptured(
		delimeter,
		ipv6Segment,
	)

	return Group.Composite(
		Group.NonCaptured(
			// 1:2:3:4:5:6:7:8
			ipv6SegmentDelimeter.Repeat().Exactly(7),
			ipv6Segment,
		),
		Group.NonCaptured(
			// 1::
			// 1:2:3:4:5:6:7::
			ipv6SegmentDelimeter.Repeat().Between(1, 7),
			delimeter,
		),
		Group.NonCaptured(
			// 1::8
			// 1:2:3:4:5:6::8
			ipv6SegmentDelimeter.Repeat().Between(1, 6),
			delimeterIPv6Segment,
		),
		Group.Define(
			// 1::7:8
			// 1:2:3:4:5::7:8
			// 1:2:3:4:5::8
			ipv6SegmentDelimeter.Repeat().Between(1, 5),
			delimeterIPv6Segment.Repeat().Between(1, 2),
		),
		Group.NonCaptured(
			// 1::6:7:8
			// 1:2:3:4::6:7:8
			// 1:2:3:4::8
			ipv6SegmentDelimeter.Repeat().Between(1, 4),
			delimeterIPv6Segment.Repeat().Between(1, 3),
		),
		Group.NonCaptured(
			// 1::5:6:7:8
			// 1:2:3::5:6:7:8
			// 1:2:3::8
			ipv6SegmentDelimeter.Repeat().Between(1, 3),
			delimeterIPv6Segment.Repeat().Between(1, 4),
		),
		Group.NonCaptured(
			// 1::4:5:6:7:8
			// 1:2::4:5:6:7:8
			// 1:2::8
			ipv6SegmentDelimeter.Repeat().Between(1, 2),
			delimeterIPv6Segment.Repeat().Between(1, 5),
		),
		Group.NonCaptured(
			// 1::3:4:5:6:7:8
			// 1::8
			ipv6SegmentDelimeter,
			delimeterIPv6Segment.Repeat().Between(1, 6),
		),
		Group.NonCaptured(
			// ::2:3:4:5:6:7:8
			// ::8
			// ::
			delimeter,
			Group.Composite(
				delimeterIPv6Segment.Repeat().Between(1, 7),
				// Or.
				delimeter,
			).NonCaptured(),
		),
		Group.NonCaptured(
			// fe80::7:8%eth0
			// fe80::7:8%1
			// (link-local IPv6 addresses with zone index)
			Group.Composite(
				Common.Text("fe"),
				Common.Text("FE"),
			).NonCaptured(),
			Common.Text("80"),
			delimeter,
			delimeterIPv6Segment.Repeat().Between(0, 4),
			Chars.Single('%'),
			Chars.Alphanumeric().Repeat().OneOrMore(),
		),
		Group.NonCaptured(
			// ::255.255.255.255
			// ::ffff:255.255.255.255
			// ::ffff:0:255.255.255.255
			// (IPv4-mapped IPv6 addresses and IPv4-translated addresses).
			delimeter.Repeat().Exactly(2),
			Group.NonCaptured(
				Chars.Runes("fF").Repeat().Exactly(4),
				Group.NonCaptured(
					delimeter,
					Chars.Single('0').Repeat().Between(1, 4),
				).Repeat().ZeroOrOne(),
				delimeter,
			).Repeat().ZeroOrOne(),
			h.IPv4(),
		),
		Group.NonCaptured(
			// 2001:db8:3:4::192.0.2.33
			// 64:ff9b::192.0.2.33
			// (IPv4-Embedded IPv6 Address).
			ipv6SegmentDelimeter.Repeat().Between(1, 4),
			delimeter,
			h.IPv4(),
		),
	)
}
