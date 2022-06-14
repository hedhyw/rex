package base

import (
	"github.com/hedhyw/rex/internal/helper"
	"github.com/hedhyw/rex/pkg/dialect"
)

// ClassToken helps to specify class tokens.
type ClassToken struct {
	brackets    bool
	exclude     bool
	classTokens []dialect.Token
	repetition  string
}

func newClassToken(classTokens ...dialect.Token) ClassToken {
	return ClassToken{
		exclude:     false,
		brackets:    true,
		classTokens: classTokens,
		repetition:  "",
	}
}

func (ct ClassToken) withoutBrackets() ClassToken {
	ct.brackets = false

	return ct
}

func (ct ClassToken) withExclude() ClassToken {
	ct.exclude = true

	return ct
}

// WriteTo implements dialect.Token interface.
func (ct ClassToken) WriteTo(w dialect.StringByteWriter) (n int, err error) {
	tokens := make([]dialect.Token, 0, 3+len(ct.classTokens))

	if ct.brackets {
		tokens = append(tokens, helper.ByteToken('['))

		if ct.exclude {
			tokens = append(tokens, helper.ByteToken('^'))
		}
	}

	for _, tok := range ct.classTokens {
		if classTok, ok := tok.(ClassToken); ok {
			tokens = append(tokens, classTok.withoutBrackets())

			continue
		}

		tokens = append(tokens, tok)
	}

	if ct.brackets {
		tokens = append(tokens, helper.ByteToken(']'))
	}

	if ct.repetition != "" {
		tokens = append(tokens, helper.StringToken(ct.repetition))
	}

	return helper.ProcessTokens(w, tokens)
}