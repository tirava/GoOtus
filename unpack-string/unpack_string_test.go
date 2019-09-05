/*
 * HomeWork-2: Unpack String
 * Created on 05.09.19 23:12
 * Copyright (c) 2019 - Eugene Klimov
 */

package unpackstring

import "testing"

// run-length decode a string
var decodeTests = []struct {
	input       string
	expected    string
	description string
}{
	{"", "", "empty string"},
	{"XYZ", "XYZ", "single characters only"},
	{"2A3B4C", "AABBBCCCC", "string with no single characters"},
	{"12WB12W3B24WB", "WWWWWWWWWWWWBWWWWWWWWWWWWBBBWWWWWWWWWWWWWWWWWWWWWWWWB", "single characters with repeated characters"},
	{"2 hs2q q2w2 ", "  hsqq qww  ", "multiple whitespace mixed in string"},
	{"2a3b4c", "aabbbcccc", "lower case string"},
}

func TestRunLengthDecode(t *testing.T) {
	for _, test := range decodeTests {
		if actual := RunLengthDecode(test.input); actual != test.expected {
			t.Errorf("FAIL %s - RunLengthDecode(%s) = %q, expected %q.",
				test.description, test.input, actual, test.expected)
		}
		t.Logf("PASS RunLengthDecode - %s", test.description)
	}
}

func BenchmarkNth(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RunLengthDecode("2a3b4c")
	}
}
