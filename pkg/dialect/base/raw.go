package base

import "github.com/hedhyw/rex/pkg/dialect"

// RawToken holds raw regular expression.
type RawToken struct {
	value string
}

// Unwrap implements dialect.ClassToken.
func (rt RawToken) Unwrap() dialect.ClassToken {
	return rt
}

// WriteTo implements dialect.Token interface.
func (rt RawToken) WriteTo(w dialect.StringByteWriter) (n int, err error) {
	return w.WriteString(rt.value)
}
