package helper

import (
	"fmt"

	"github.com/hedhyw/rex/pkg/dialect"
)

// TokenFunc implements Token.
type TokenFunc func(w dialect.StringByteWriter) (int, error)

// WriteTo implements dialect.Token interface.
func (tf TokenFunc) WriteTo(w dialect.StringByteWriter) (int, error) {
	return tf(w)
}

// StringToken creates a string token with formatting.
func StringToken(f string, args ...interface{}) dialect.Token {
	return TokenFunc(func(w dialect.StringByteWriter) (int, error) {
		return w.WriteString(fmt.Sprintf(f, args...))
	})
}

// ByteToken creates a byte.
func ByteToken(val byte) dialect.Token {
	return TokenFunc(func(w dialect.StringByteWriter) (int, error) {
		return 1, w.WriteByte(val)
	})
}

// ProcessTokens goes thought all tokens and call WriteTo method.
func ProcessTokens(w dialect.StringByteWriter, tokens []dialect.Token) (int, error) {
	var totalWritten int

	for _, t := range tokens {
		n, err := t.WriteTo(w)
		if err != nil {
			return n, err
		}

		totalWritten += n
	}

	return totalWritten, nil
}
