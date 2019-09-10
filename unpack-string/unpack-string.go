/*
 * HomeWork-2: Unpack String
 * Created on 07.09.19 12:04
 * Copyright (c) 2019 - Eugene Klimov
 */

// Package unpackstring implements unpacking string.
package unpackstring

import (
	"strings"
	"unicode"
)

// UnpackString returns string unpacked.
func UnpackString(input string) string {

	digits := ""
	char := ""
	esc := false

	result := strings.Builder{}

	for i, r := range input {

		if r == 0x5c && !esc { // 0x5c = `\`
			esc = true
			continue
		}

		if unicode.IsDigit(r) && !esc {
			digits += string(r)
			if i != len(input)-1 { // last char is number
				continue
			}
		}

		digit := 0
		for _, r := range digits { // fast convert string to int
			digit = digit*10 + int(r-'0')
		}

		if digit == 0 && len(digits) > 0 { //delete last char if zero repeated
			s := strings.TrimRight(result.String(), char)
			result.Reset()
			result.WriteString(s)
		}

		if digit > 0 { // flush series
			result.WriteString(strings.Repeat(char, digit-1)) // one char already appended
			digits = ""
			char = ""
		}

		if unicode.IsLetter(r) || unicode.IsSpace(r) || esc {
			char = string(r)
			result.WriteRune(r)
			esc = false
		}
	}

	return result.String()
}
