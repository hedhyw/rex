package rex

import "github.com/hedhyw/rex/pkg/dialect/base"

const (
	// Chars is a namespace that contains character class elements.
	// It is an alias to github.com/hedhyw/rex/pkg/dialect/base.Chars.
	//
	// Example usage:
	//
	//   rex.New(rex.Chars.Digits()) // `[0-9]`
	//
	//   rex.New(rex.Common.Class(
	//     rex.Chars.Digits(),
	//     rex.Chars.Single('a'),
	//     rex.Chars.Runes("bc"),
	//   ) // `[0-9abc]`
	Chars = base.Chars
	// Common is a namespace that contains base regular expression helpers.
	// It is an alias to github.com/hedhyw/rex/pkg/dialect/base.Common.
	//
	// Example usage:
	//
	//   // Raw regular expression.
	//   rex.New(rex.Common.Raw(`[a-z]+`)) // `[a-z]+`
	//
	//   // Escaped text. It helps to match the same text.
	//   rex.New(rex.Common.Text(`escaped text.`)) // `escaped text\.`
	//
	//   // Combine characters
	//   rex.New(rex.Common.Class(
	//     rex.Chars.Digits(),
	//     rex.Chars.Single('a'),
	//   ) // `[0-9a]`
	//
	//   // Exclude characters.
	//   rex.New(rex.Common.Class(rex.Chars.Digits())) // `[^0-9]`
	Common = base.Common
	// Group is a namespace that contains helpers for grouping expressions.
	//
	// Example usage:
	//
	//   base.Group.Define(
	//     base.Chars.Single('a'),
	//     base.Chars.Composite('a'),
	//   ).Repeat().OneOrMore() // (a)+
	//
	Group = base.Group
)
