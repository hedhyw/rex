package base

import "github.com/hedhyw/rex/pkg/dialect"

// MD5Hex is a pattern for a cryptographic hash function MD5 in hex representation.
//
// Example: d41d8cd98f00b204e9800998ecf8427e.
func (h HelperDialect) MD5Hex() dialect.Token {
	return h.hex(32)
}

// SHA1Hex is a pattern for a cryptographic hash function SHA1 in hex representation.
//
// Example: da39a3ee5e6b4b0d3255bfef95601890afd80709.
func (h HelperDialect) SHA1Hex() dialect.Token {
	return h.hex(40)
}

// MD5 is a pattern for a cryptographic hash function SHA256 in hex representation.
//
// Example: e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855.
func (h HelperDialect) SHA256Hex() dialect.Token {
	return h.hex(64)
}

func (HelperDialect) hex(length int) dialect.Token {
	return Group.NonCaptured(
		Chars.HexDigits().Repeat().Exactly(length),
	)
}
