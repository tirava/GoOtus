/*
 * HomeWork-2: Unpack String tests
 * Created on 07.09.19 13:12
 * Copyright (c) 2019 - Eugene Klimov
 */

package unpackstring

import "testing"

var decodeTests = []struct {
	input       string
	expected    string
	description string
}{
	{"", "", "empty string"},
	{"a4bc2d5e", "aaaabccddddde", "simple coded"},
	{"abcd", "abcd", "single characters lower"},
	{"45", "", "fail string"},
	{"XYZ", "XYZ", "single characters upper"},
	{"A2B3C4", "AABBBCCCC", "no single characters upper"},
	{"W12BW12B3W24B", "WWWWWWWWWWWWBWWWWWWWWWWWWBBBWWWWWWWWWWWWWWWWWWWWWWWWB", "many characters with repeat"},
	{" 2hsq2 qw2 2", "  hsqq qww  ", "whitespace mixed in string"},
	{"a2b3c4", "aabbbcccc", "no single characters lower"},
	{"a0b2", "bb", "with zero count"},
	{"z1y1x1", "zyx", "only one count per char"},
	{`\,1\$2\.3\*4`, ",$$...****", "esc punctuation chars"},
	{`qwe\4\5`, `qwe45`, "string with 2 esc numbers"},
	{`qwe\45`, `qwe44444`, "string with 1 esc char"},
	{`qwe\\5`, `qwe\\\\\`, "string with same esc character"},
	{`\`, "", "fail esc string"},
	{"А1Б2Ц3Я0", "АББЦЦЦ", "cyrillic string"},
	{`a4bc2d5eabcdXYZA2B3C4W12BW12B3W24B 2hsq2 qw2 2a2b3c4a0b2z1y1x1\,1\$2\.3\*4qwe\4\5qwe\45qwe\\5А1Б2Ц3`,
		`aaaabccdddddeabcdXYZAABBBCCCCWWWWWWWWWWWWBWWWWWWWWWWWWBBBWWWWWWWWWWWWWWWWWWWWWWWWB  hsqq qww  aabbbccccbbzyx,$$...****qwe45qwe44444qwe\\\\\АББЦЦЦ`,
		"mixed all test strings"},
}

func TestUnpackString(t *testing.T) {
	for _, test := range decodeTests {
		if actual := UnpackString(test.input); actual != test.expected {
			t.Errorf("FAIL %s - UnpackString(%s) = '%s', expected '%s'.",
				test.description, test.input, actual, test.expected)
			continue
		}
		t.Logf("PASS UnpackString - %s", test.description)
	}
}

func BenchmarkUnpackString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UnpackString(`a4bc2d5eabcdXYZA2B3C4W12BW12B3W24B 2hsq2 qw2 2a2b3c4a0b2a0000b2z1y1x1\,1\$2\.3\*4qwe\4\5qwe\45qwe\\5`)
	}
}
