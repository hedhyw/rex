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

type Token interface {
	WriteTo(w StringByteWriter) (n int, err error)
}
