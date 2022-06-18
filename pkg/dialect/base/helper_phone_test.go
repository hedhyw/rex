package base_test

import (
	"testing"

	"github.com/hedhyw/rex/internal/test"
	"github.com/hedhyw/rex/pkg/dialect/base"
)

func getPhoneE164ValidTestCases() test.MatchTestCaseSlice {
	return test.MatchTestCaseSlice{{
		Name:  "e164_ok_1",
		Value: "+14155552671",
	}, {
		Name:  "e164_ok_2",
		Value: "+442071838750",
	}, {
		Name:  "e164_ok_3",
		Value: "+551155256325",
	}, {
		Name:  "e164_ok_4",
		Value: "+2125551212",
	}, {
		Name:  "e164_ok_example_docs",
		Value: "+15555555",
	}}
}

func getPhoneE164InvalidTestCases() test.MatchTestCaseSlice {
	return test.MatchTestCaseSlice{{
		Name:  "e164_space_and_brackets",
		Value: "+1 (212) 555-1212",
	}, {
		Name:  "e164_minuses",
		Value: "+1-212-555-1212",
	}, {
		Name:  "e164_spaces",
		Value: "+1 212 555 1212",
	}, {
		Name:  "e164_no_plus",
		Value: "2125551212",
	}, {
		Name:  "e164_zero_country_code",
		Value: "+0125551212",
	}, {
		Name:  "e164_short",
		Value: "+1234567",
	}}
}

func TestPhoneE164(t *testing.T) {
	test.MatchTestCaseGroupSlice{
		getPhoneE164ValidTestCases().WithMatched(true),
		getPhoneE164InvalidTestCases().WithMatched(false),
		getPhoneNationalE123ValidTestCases().WithMatched(false),
		getPhoneInternationalE123ValidTestCases().WithMatched(false),
	}.Run(t, base.Helper.PhoneE164())
}

func getPhoneInternationalE123ValidTestCases() test.MatchTestCaseSlice {
	return test.MatchTestCaseSlice{{
		Name:  "e123_international_ok_1",
		Value: "+1 212 555 1212",
	}, {
		Name:  "e123_international_ok_2",
		Value: "+31 42 123 4567",
	}, {
		Name:  "e123_international_ok_example",
		Value: "+22 607 123 4567",
	}}
}

func getPhoneInternationalE123InvalidTestCases() test.MatchTestCaseSlice {
	return test.MatchTestCaseSlice{{
		Name:  "e123_international_zero_country_code",
		Value: "+0 212 555 1212",
	}, {
		Name:  "e123_international_e164",
		Value: "+14155552671",
	}, {
		Name:  "e123_international_brackets",
		Value: "+1 (212) 555-1212",
	}, {
		Name:  "e123_international_minuses",
		Value: "+1-212-555-1212",
	}, {
		Name:  "e123_international_pluses",
		Value: "1 123 456 7890",
	}}
}

func TestPhoneInternationalE123(t *testing.T) {
	test.MatchTestCaseGroupSlice{
		getPhoneInternationalE123ValidTestCases().WithMatched(true),
		getPhoneInternationalE123InvalidTestCases().WithMatched(false),
		getPhoneNationalE123ValidTestCases().WithMatched(false),
		getPhoneE164ValidTestCases().WithMatched(false),
	}.Run(t, base.Helper.PhoneInternationalE123())
}

func getPhoneNationalE123ValidTestCases() test.MatchTestCaseSlice {
	return test.MatchTestCaseSlice{{
		Name:  "e123_national_ok_1",
		Value: "(112) 123 4567",
	}, {
		Name:  "e123_national_ok_example",
		Value: "(607) 123 4567",
	}}
}

func getPhoneNationalE123InvalidTestCases() test.MatchTestCaseSlice {
	return test.MatchTestCaseSlice{{
		Name:  "e123_national_international",
		Value: "+1 212 555 1212",
	}, {
		Name:  "e123_national_e164",
		Value: "+14155552671",
	}, {
		Name:  "e123_national_minus",
		Value: "(112) 123-4567",
	}, {
		Name:  "e123_national_no_brackets",
		Value: "112 123 4567",
	}, {
		Name:  "e123_national_no_spaces",
		Value: "(112)1234567",
	}}
}

func TestPhoneNationalE123(t *testing.T) {
	test.MatchTestCaseGroupSlice{
		getPhoneNationalE123ValidTestCases().WithMatched(true),
		getPhoneNationalE123InvalidTestCases().WithMatched(false),
		getPhoneInternationalE123ValidTestCases().WithMatched(false),
		getPhoneE164ValidTestCases().WithMatched(false),
	}.Run(t, base.Helper.PhoneNationalE123())
}

func TestPhoneE123(t *testing.T) {
	test.MatchTestCaseGroupSlice{
		getPhoneNationalE123ValidTestCases().WithMatched(true),
		getPhoneInternationalE123ValidTestCases().WithMatched(true),
		getPhoneE164ValidTestCases().WithMatched(false),
	}.Run(t, base.Helper.PhoneE123())
}

func TestPhone(t *testing.T) {
	test.MatchTestCaseGroupSlice{
		getPhoneNationalE123ValidTestCases().WithMatched(true),
		getPhoneInternationalE123ValidTestCases().WithMatched(true),
		getPhoneE164ValidTestCases().WithMatched(true),
		test.MatchTestCaseSlice{{
			Name:  "string",
			Value: "invalid",
		}}.WithMatched(false),
	}.Run(t, base.Helper.Phone())
}
