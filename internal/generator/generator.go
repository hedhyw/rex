package generator

import (
	"fmt"
	"io"
	"regexp/syntax"
	"strings"
)

// GenerateCode returns rex code for a given regex.
func GenerateCode(regex string) (generatedCode string, err error) {
	regExpr, err := syntax.Parse(regex, syntax.Perl)
	if err != nil {
		return "", fmt.Errorf("failed to parse regexp: %w", err)
	}

	var strBuilder strings.Builder

	strBuilder.Grow(len(regex))
	_, _ = strBuilder.WriteString("rex.New(\n")
	writeRegexp(&strBuilder, regExpr, 1)
	_, _ = strBuilder.WriteString(")")

	return strBuilder.String(), nil
}

func writeRegexp(w io.StringWriter, regExpr *syntax.Regexp, indent int) {
	//nolint: exhaustive // All cases captured in default.
	switch regExpr.Op {
	case syntax.OpConcat:
		writeConcat(w, regExpr, indent)
	case syntax.OpCapture:
		writeCapture(w, regExpr, indent)
	default:
		writeRaw(w, []*syntax.Regexp{regExpr}, indent)
	}
}

func writeCapture(w io.StringWriter, regExpr *syntax.Regexp, indent int) {
	strIndent := strings.Repeat("\t", indent)

	_, _ = w.WriteString(strIndent + "rex.Group.Define(\n")

	if regExpr.Sub[0].Op != syntax.OpEmptyMatch {
		writeRegexp(w, regExpr.Sub[0], indent+1)
	}

	if regExpr.Name != "" {
		_, _ = w.WriteString(fmt.Sprintf("%s).WithName(%q),\n", strIndent, regExpr.Name))
	} else {
		_, _ = w.WriteString(strIndent + "),\n")
	}
}

func writeConcat(w io.StringWriter, regExpr *syntax.Regexp, indent int) {
	rawExprs := make([]*syntax.Regexp, 0, len(regExpr.Sub))

	for _, sub := range regExpr.Sub {
		//nolint: exhaustive // All cases captured in default.
		switch sub.Op {
		case syntax.OpCapture:
			writeRaw(w, rawExprs, indent)

			rawExprs = rawExprs[:0]

			writeRegexp(w, sub, indent)
		default:
			rawExprs = append(rawExprs, sub)
		}
	}

	writeRaw(w, rawExprs, indent)
}

func writeRaw(w io.StringWriter, regExprs []*syntax.Regexp, indent int) {
	if len(regExprs) == 0 {
		return
	}

	strIndent := strings.Repeat("\t", indent)

	_, _ = w.WriteString(strIndent + "rex.Common.Raw(`")

	for _, re := range regExprs {
		_, _ = w.WriteString(re.String())
	}

	_, _ = w.WriteString("`),\n")
}
