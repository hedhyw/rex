// nolint: funlen // Unit tests.
package base_test

import (
	"fmt"
	"net"
	"strings"
	"testing"

	"github.com/hedhyw/rex/internal/test"
	"github.com/hedhyw/rex/pkg/dialect/base"
	"github.com/hedhyw/rex/pkg/rex"
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

func getIPv4ValidTestCases() test.MatchTestCaseSlice {
	return test.MatchTestCaseSlice{{
		Name:  "ipv4_local",
		Value: "127.0.0.1",
	}, {
		Name:  "ipv4_zero",
		Value: "0.0.0.0",
	}, {
		Name:  "ipv4_zero",
		Value: "0.0.0.0",
	}, {
		Name:  "ipv4_max",
		Value: "255.255.255.255",
	}, {
		Name:  "ipv4_max",
		Value: "199.199.199.199",
	}, {
		Name:  "ipv4_google",
		Value: "172.217.16.46",
	}, {
		Name:  "ipv4_private",
		Value: "192.168.1.1",
	}}
}

func getIPv4InvalidTestCases() test.MatchTestCaseSlice {
	return test.MatchTestCaseSlice{{
		Name:  "ipv4_max",
		Value: "256.255.255.255",
	}, {
		Name:  "ipv4_three",
		Value: "255.255.255",
	}, {
		Name:  "ipv4_extra_dot",
		Value: "0.0.0.0.",
	}, {
		Name:  "ipv4_300",
		Value: "255.300.255.255",
	}, {
		Name:  "ipv4_leading_zero",
		Value: "009.009.009.009",
	}}
}

func TestIPv4(t *testing.T) {
	test.MatchTestCaseGroupSlice{
		getIPv4ValidTestCases().WithMatched(true),
		getIPv6ValidTestCases().WithMatched(false),
		getIPv4InvalidTestCases().WithMatched(false),
	}.Run(t, base.Helper.IPv4())
}

func FuzzIPv4(f *testing.F) {
	f.Add(255, 255, 255, 255)
	f.Add(0, 0, 0, 0)
	f.Add(127, 0, 0, 1)
	f.Add(256, 0, 0, 0)

	re := rex.New(base.Group.NonCaptured(
		base.Chars.Begin(),
		base.Group.NonCaptured(base.Helper.IPv4()),
		base.Chars.End(),
	)).MustCompile()

	f.Fuzz(func(t *testing.T, a int, b int, c int, d int) {
		fuzzIP := fmt.Sprintf("%d.%d.%d.%d", a, b, c, d)

		expected := net.ParseIP(fuzzIP) != nil
		actual := re.MatchString(fuzzIP)
		if expected != actual {
			t.Errorf("Actual: %v, Expected: %v", actual, expected)
		}
	})
}

func getIPv6ValidTestCases() test.MatchTestCaseSlice {
	return test.MatchTestCaseSlice{{
		Name:  "ipv6_wiki",
		Value: "2001:0db8:85a3:0000:0000:8a2e:0370:7334",
	}, {
		Name:  "ipv6_pattern_1",
		Value: "1:2:3:4:5:6:7:8",
	}, {
		Name:  "ipv6_pattern_2",
		Value: "1::",
	}, {
		Name:  "ipv6_pattern_3",
		Value: "1:2:3:4:5:6:7::",
	}, {
		Name:  "ipv6_pattern_4",
		Value: "1::8",
	}, {
		Name:  "ipv6_pattern_5",
		Value: "1:2:3:4:5:6::8",
	}, {
		Name:  "ipv6_pattern_6",
		Value: "1::7:8",
	}, {
		Name:  "ipv6_pattern_7",
		Value: "1:2:3:4:5::7:8",
	}, {
		Name:  "ipv6_pattern_8",
		Value: "1:2:3:4:5::8",
	}, {
		Name:  "ipv6_pattern_9",
		Value: "1::6:7:8",
	}, {
		Name:  "ipv6_pattern_10",
		Value: "1:2:3:4::6:7:8",
	}, {
		Name:  "ipv6_pattern_11",
		Value: "1:2:3:4::8",
	}, {
		Name:  "ipv6_pattern_12",
		Value: "1::7:8",
	}, {
		Name:  "ipv6_pattern_13",
		Value: "1:2:3:4:5::7:8",
	}, {
		Name:  "ipv6_pattern_14",
		Value: "1:2:3:4:5::8",
	}, {
		Name:  "ipv6_pattern_15",
		Value: "1::6:7:8",
	}, {
		Name:  "ipv6_pattern_16",
		Value: "1:2:3:4::6:7:8",
	}, {
		Name:  "ipv6_pattern_17",
		Value: "1:2:3:4::8",
	}, {
		Name:  "ipv6_pattern_18",
		Value: "1::5:6:7:8",
	}, {
		Name:  "ipv6_pattern_19",
		Value: "1:2:3::5:6:7:8",
	}, {
		Name:  "ipv6_pattern_20",
		Value: "1:2:3::8",
	}, {
		Name:  "ipv6_pattern_21",
		Value: "1::4:5:6:7:8",
	}, {
		Name:  "ipv6_pattern_22",
		Value: "1:2::4:5:6:7:8",
	}, {
		Name:  "ipv6_pattern_23",
		Value: "1:2::8",
	}, {
		Name:  "ipv6_pattern_24",
		Value: "1::3:4:5:6:7:8",
	}, {
		Name:  "ipv6_pattern_25",
		Value: "1::8",
	}, {
		Name:  "ipv6_pattern_26",
		Value: "::2:3:4:5:6:7:8",
	}, {
		Name:  "ipv6_pattern_27",
		Value: "::8",
	}, {
		Name:  "ipv6_pattern_28",
		Value: "::",
	}, {
		Name:  "ipv6_pattern_29",
		Value: "fe80::7:8%eth0",
	}, {
		Name:  "ipv6_pattern_30",
		Value: "fe80::7:8%1",
	}, {
		Name:  "ipv6_pattern_31",
		Value: "::255.255.255.255",
	}, {
		Name:  "ipv6_pattern_32",
		Value: "::ffff:255.255.255.255",
	}, {
		Name:  "ipv6_pattern_33",
		Value: "::ffff:0:255.255.255.255",
	}, {
		Name:  "ipv6_pattern_34",
		Value: "2001:db8:3:4::192.0.2.33",
	}, {
		Name:  "ipv6_pattern_35",
		Value: "64:ff9b::192.0.2.33",
	}}
}

func getIPv6InvalidTestCases() test.MatchTestCaseSlice {
	return test.MatchTestCaseSlice{{
		Name:  "ipv6_not_hex",
		Value: "G001:0db8:85a3:0000:0000:8a2e:0370:7334",
	}, {
		Name:  "ipv6_invalid_ip",
		Value: "::ffff:0:256.255.255.255",
	}, {
		Name:  "ipv6_ipv4",
		Value: "225.1.4.2",
	}, {
		Name:  "ipv6_count_tokens",
		Value: "fe80:2030:31:24",
	}}
}

func TestIPv6(t *testing.T) {
	test.MatchTestCaseGroupSlice{
		getIPv6ValidTestCases().WithMatched(true),
		getIPv4ValidTestCases().WithMatched(false),
		getIPv6InvalidTestCases().WithMatched(false),
	}.Run(t, base.Helper.IPv6())
}

func TestIP(t *testing.T) {
	test.MatchTestCaseGroupSlice{
		getIPv6ValidTestCases().WithMatched(true),
		getIPv4ValidTestCases().WithMatched(true),
	}.Run(t, base.Helper.IP())
}
