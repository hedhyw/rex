package base_test

import (
	"strings"
	"testing"

	"github.com/hedhyw/rex/internal/test"
	"github.com/hedhyw/rex/pkg/dialect/base"
)

func getHostnameRFC952ValidTestCases() test.MatchTestCaseSlice {
	return test.MatchTestCaseSlice{{
		Name:  "rfc952_ok_google",
		Value: "google.com",
	}, {
		Name:  "rfc952_ok_third_level",
		Value: "git.github.com",
	}, {
		Name:  "rfc952_ok_com",
		Value: "com",
	}}
}

func getHostnameRFC952InvalidTestCases() test.MatchTestCaseSlice {
	return test.MatchTestCaseSlice{{
		Name:  "rfc952_ip",
		Value: "127.0.0.1",
	}, {
		Name:  "start_with_zero",
		Value: "0.example.com",
	}, {
		Name:  "end_with_a_dot",
		Value: "example.com.",
	}, {
		Name:  "end_with_a_sign",
		Value: "-example.com.",
	}, {
		Name:  "start_with_a_dot",
		Value: ".example.com",
	}}
}

func TestHostnameRFC952(t *testing.T) {
	test.MatchTestCaseGroupSlice{
		getHostnameRFC952ValidTestCases().WithMatched(true),
		getHostnameRFC952InvalidTestCases().WithMatched(false),
	}.Run(t, base.Helper.HostnameRFC952())
}

func getHostnameRFC1123ValidTestCases() test.MatchTestCaseSlice {
	return test.MatchTestCaseSlice{{
		Name:  "rfc1123_ok_google",
		Value: "google.com",
	}, {
		Name:  "rfc1123_ok_third_level",
		Value: "git.github.com",
	}, {
		Name:  "rfc1123_ok_com",
		Value: "com",
	}, {
		Name:  "rfc1123_starts_with_zero",
		Value: "0.example.com",
	}}
}

func getHostnameRFC1123InvalidTestCases() test.MatchTestCaseSlice {
	return test.MatchTestCaseSlice{{
		Name:  "rfc1123_ip",
		Value: "127.0.0.1",
	}, {
		Name:  "rfc1123_end_with_a_dot",
		Value: "example.com.",
	}, {
		Name:  "rfc1123_ends_with_a_sign",
		Value: "example.com-",
	}, {
		Name:  "rfc1123_end_with_a_sign",
		Value: "-example.com.",
	}, {
		Name:  "rfc1123_start_with_a_dot",
		Value: ".example.com",
	}}
}

func TestHostnameRFC1123(t *testing.T) {
	test.MatchTestCaseGroupSlice{
		getHostnameRFC952ValidTestCases().WithMatched(true),
		getHostnameRFC1123ValidTestCases().WithMatched(true),
		getHostnameRFC1123InvalidTestCases().WithMatched(false),
	}.Run(t, base.Helper.HostnameRFC1123())
}

func getEmailValidTestCases() test.MatchTestCaseSlice {
	return test.MatchTestCaseSlice{{
		Name:  "email_ok_simple",
		Value: "test@example.com",
	}, {
		Name:  "email_ok_dash",
		Value: "test-email@example.com",
	}, {
		Name:  "email_ok_dot",
		Value: "test.email@example.com",
	}, {
		Name:  "email_ok_underscore",
		Value: "test_email@example.com",
	}, {
		Name:  "subdomain",
		Value: "email@sub.example.com",
	}, {
		Name:  "number",
		Value: "123@example.com",
	}, {
		Name:  "many_dots",
		Value: "1.2.3.4.5.6.7.8.9.10@example.com",
	}, {
		Name:  "email_ok_long",
		Value: strings.Repeat("1", 63) + "@example.com",
	}, {
		Name:  "email_ok_single_local",
		Value: "0@example.com",
	}}
}

func getEmailInvalidTestCases() test.MatchTestCaseSlice {
	return test.MatchTestCaseSlice{{
		Name:  "email_starts_with_a_dot",
		Value: ".test@example.com",
	}, {
		Name:  "email_end_with_a_dot",
		Value: "test.@example.com",
	}, {
		Name:  "email_two_dots_consecutively",
		Value: "te..st@example.com",
	}, {
		Name:  "email_host_invalid",
		Value: "test@127.0.0.1",
	}, {
		Name:  "email_two_at",
		Value: "two@parts@example.com",
	}, {
		Name:  "email_long",
		Value: strings.Repeat("1", 65) + "@example.com",
	}, {
		Name:  "email_no_local",
		Value: "@example.com",
	}}
}

func TestEmail(t *testing.T) {
	test.MatchTestCaseGroupSlice{
		getEmailValidTestCases().WithMatched(true),
		getEmailInvalidTestCases().WithMatched(false),
	}.Run(t, base.Helper.Email())
}
