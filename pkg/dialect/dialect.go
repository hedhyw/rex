package dialect

import (
	"io"
)

// Dialect specifies group of dialect tokens. It saves the name.
type Dialect string

// StringByteWriter is a union of io.ByteWriter and io.StringWriter.
type StringByteWriter interface {
	io.ByteWriter
	io.StringWriter
}

// Token provides a simple write operation for regular expressions.
type Token interface {
	// WriteTo current position.
	WriteTo(w StringByteWriter) (n int, err error)
}

// ClassToken represents a character or set of characters.
type ClassToken interface {
	Token

	// Unwrap removes brackets that wrap ClassToken.
	// If ClassToken is already unwrapped, then this operation does nothing.
	// Example: [a-z] -> a-z.
	Unwrap() ClassToken
}
