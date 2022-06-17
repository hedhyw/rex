package base

import (
	"github.com/hedhyw/rex/internal/helper"
	"github.com/hedhyw/rex/pkg/dialect"
)

// RepetableClassToken handles repetitions of the class.
type RepetableClassToken struct {
	classToken dialect.ClassToken
	Repetable
}

// Unwrap implements dialect.ClassToken.
func (rct RepetableClassToken) Unwrap() dialect.ClassToken {
	return rct.classToken.Unwrap()
}

func newRepetableClassToken(classToken dialect.ClassToken) RepetableClassToken {
	return RepetableClassToken{
		classToken: classToken,
		Repetable:  newRepetable(classToken),
	}
}

// ClassToken helps to specify class tokens.
type ClassToken struct {
	brackets    bool
	exclude     bool
	classTokens []dialect.Token
}

func newClassToken(classTokens ...dialect.Token) ClassToken {
	return ClassToken{
		exclude:     false,
		brackets:    true,
		classTokens: classTokens,
	}
}

// Unwrap implements dialect.ClassToken.
func (ct ClassToken) Unwrap() dialect.ClassToken {
	return ct.withoutBrackets()
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

	tokens = append(tokens, ct.classTokens...)

	if ct.brackets {
		tokens = append(tokens, helper.ByteToken(']'))
	}

	return helper.ProcessTokens(w, tokens)
}
