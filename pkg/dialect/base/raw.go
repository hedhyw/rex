package base

import (
	"strings"

	"github.com/hedhyw/rex/pkg/dialect"
)

// RawToken holds raw regular expression.
type RawToken struct {
	value   string
	verbose bool
}

// Unwrap implements dialect.ClassToken.
func (rt RawToken) Unwrap() dialect.ClassToken {
	return rt
}

// WriteTo implements dialect.Token interface.
func (rt RawToken) WriteTo(w dialect.StringByteWriter) (n int, err error) {
	value := rt.value

	if rt.verbose {
		lines := strings.Split(value, "\n")
		for i, l := range lines {
			l = removeComment(l)
			l = strings.ReplaceAll(l, "\\#", "#")
			l = strings.TrimSpace(l)
			lines[i] = l
		}

		value = strings.Join(lines, "")
	}

	return w.WriteString(value)
}

// removeComment removes everythong after '#' if it is not escaped by
// backslash '\#' or it is not in the character class '[#]'.
//
// This input:
//
//	.+\#[#] # comment
//
// Will be converted to:
//
//	.+\#[#]
func removeComment(val string) string {
	var (
		backslash      bool
		characterClass bool
	)

	for i, ch := range val {
		switch ch {
		case '\\':
			backslash = !backslash

			continue
		case '#':
			if !backslash && !characterClass {
				return string([]rune(val)[:i])
			}
		case '[':
			if !backslash {
				characterClass = true
			}
		case ']':
			if !backslash {
				characterClass = false
			}
		}

		backslash = false
	}

	return val
}
