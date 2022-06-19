// nolint: gosec // It is a test.
package base_test

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
	"testing"

	"github.com/hedhyw/rex/internal/test"
	"github.com/hedhyw/rex/pkg/dialect/base"
)

func getMD5ValidTestCases() test.MatchTestCaseSlice {
	return test.MatchTestCaseSlice{{
		Name:  "md5_ok_example",
		Value: "d41d8cd98f00b204e9800998ecf8427e",
	}, {
		Name:  "md5_ok_upper",
		Value: "D41D8CD98F00B204E9800998ECF8427E",
	}, {
		Name:  "md5_ok_3",
		Value: fmt.Sprintf("%x", md5.Sum([]byte("3"))),
	}, {
		Name:  "md5_ok_4",
		Value: fmt.Sprintf("%x", md5.Sum([]byte("4"))),
	}, {
		Name:  "md5_ok_5",
		Value: fmt.Sprintf("%x", md5.Sum([]byte("5"))),
	}, {
		Name:  "md5_ok_6",
		Value: fmt.Sprintf("%x", md5.Sum([]byte("6"))),
	}}
}

func getMD5InvalidTestCases() test.MatchTestCaseSlice {
	return test.MatchTestCaseSlice{{
		Name:  "md5_non_hex_p",
		Value: "p41d8cd98f00b204e9800998ecf8427e",
	}, {
		Name:  "md5_short",
		Value: fmt.Sprintf("%x", md5.Sum([]byte("short")))[0:31],
	}, {
		Name:  "md5_long",
		Value: fmt.Sprintf("%x", md5.Sum([]byte("long"))) + "0",
	}, {
		Name:  "md5_empty",
		Value: "",
	}}
}

func TestMD5Hex(t *testing.T) {
	test.MatchTestCaseGroupSlice{
		getMD5ValidTestCases().WithMatched(true),
		getMD5InvalidTestCases().WithMatched(false),
		getSHA1ValidTestCases().WithMatched(false),
		getSHA256ValidTestCases().WithMatched(false),
	}.Run(t, base.Helper.MD5Hex())
}

func getSHA1ValidTestCases() test.MatchTestCaseSlice {
	return test.MatchTestCaseSlice{{
		Name:  "sha1_ok_example",
		Value: "da39a3ee5e6b4b0d3255bfef95601890afd80709",
	}, {
		Name:  "sha1_ok_upper",
		Value: "DA39A3EE5E6B4B0D3255BFEF95601890AFD80709",
	}, {
		Name:  "sha1_ok_3",
		Value: fmt.Sprintf("%x", sha1.Sum([]byte("3"))),
	}, {
		Name:  "sha1_ok_4",
		Value: fmt.Sprintf("%x", sha1.Sum([]byte("4"))),
	}, {
		Name:  "sha1_ok_5",
		Value: fmt.Sprintf("%x", sha1.Sum([]byte("5"))),
	}, {
		Name:  "sha1_ok_6",
		Value: fmt.Sprintf("%x", sha1.Sum([]byte("6"))),
	}}
}

func getSHA1InvalidTestCases() test.MatchTestCaseSlice {
	return test.MatchTestCaseSlice{{
		Name:  "sha1_non_hex_p",
		Value: "pa39a3ee5e6b4b0d3255bfef95601890afd80709",
	}, {
		Name:  "sha1_short",
		Value: fmt.Sprintf("%x", sha1.Sum([]byte("short")))[0:39],
	}, {
		Name:  "sha1_long",
		Value: fmt.Sprintf("%x", sha1.Sum([]byte("long"))) + "0",
	}, {
		Name:  "sha1_empty",
		Value: "",
	}}
}

func TestSHA1Hex(t *testing.T) {
	test.MatchTestCaseGroupSlice{
		getSHA1ValidTestCases().WithMatched(true),
		getSHA1InvalidTestCases().WithMatched(false),
		getMD5ValidTestCases().WithMatched(false),
		getSHA256ValidTestCases().WithMatched(false),
	}.Run(t, base.Helper.SHA1Hex())
}

func getSHA256ValidTestCases() test.MatchTestCaseSlice {
	return test.MatchTestCaseSlice{{
		Name:  "sha256_ok_example",
		Value: "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855",
	}, {
		Name:  "sha256_ok_upper",
		Value: "E3B0C44298FC1C149AFBF4C8996FB92427AE41E4649B934CA495991B7852B855",
	}, {
		Name:  "sha256_ok_3",
		Value: fmt.Sprintf("%x", sha256.Sum256([]byte("3"))),
	}, {
		Name:  "sha256_ok_4",
		Value: fmt.Sprintf("%x", sha256.Sum256([]byte("4"))),
	}, {
		Name:  "sha256_ok_5",
		Value: fmt.Sprintf("%x", sha256.Sum256([]byte("5"))),
	}, {
		Name:  "sha256_ok_6",
		Value: fmt.Sprintf("%x", sha256.Sum256([]byte("6"))),
	}}
}

func getSHA256InvalidTestCases() test.MatchTestCaseSlice {
	return test.MatchTestCaseSlice{{
		Name:  "sha256_non_hex_p",
		Value: "p3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855",
	}, {
		Name:  "sha256_short",
		Value: fmt.Sprintf("%x", sha256.Sum256([]byte("short")))[0:63],
	}, {
		Name:  "sha256_long",
		Value: fmt.Sprintf("%x", sha256.Sum256([]byte("long"))) + "0",
	}, {
		Name:  "sha256_empty",
		Value: "",
	}}
}

func TestSHA256Hex(t *testing.T) {
	test.MatchTestCaseGroupSlice{
		getSHA256ValidTestCases().WithMatched(true),
		getSHA256InvalidTestCases().WithMatched(false),
		getSHA1ValidTestCases().WithMatched(false),
		getMD5ValidTestCases().WithMatched(false),
	}.Run(t, base.Helper.SHA256Hex())
}
