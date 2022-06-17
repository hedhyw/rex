package base

import (
	"github.com/hedhyw/rex/internal/helper"
	"github.com/hedhyw/rex/pkg/dialect"
)

// CompositToken defines multiple possible expression patterns.
// It can be described as T1 or T2 or T3, ...
type CompositToken struct {
	tokens []dialect.Token
}

// WriteTo implements dialect.Token interface.
func (ct CompositToken) WriteTo(w dialect.StringByteWriter) (n int, err error) {
	tokens := make([]dialect.Token, 0, len(ct.tokens)*2-1)

	for i, tok := range ct.tokens {
		if i > 0 {
			tokens = append(tokens, helper.ByteToken('|'))
		}

		tokens = append(tokens, tok)
	}

	return helper.ProcessTokens(w, tokens)
}
