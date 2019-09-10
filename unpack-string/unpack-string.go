/*
 * HomeWork-2: Unpack String
 * Created on 07.09.19 12:04
 * Copyright (c) 2019 - Eugene Klimov
 */

// Package unpackstring implements unpacking string.
package unpackstring

import (
	"strconv"
	"strings"
	"unicode"
)

// UnpackString returns string unpacked.
func UnpackString(input string) string {

	result := ""
	digits := ""
	char := ""
	esc := false

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

		digit, _ := strconv.Atoi(digits) // no need to check error - IsDigit "checks" it

		if digit == 0 && len(digits) > 0 { //delete char if zero repeated
			result = result[:len(result)-1]
			char = ""
		}

		if digit > 0 { // flush series
			result += strings.Repeat(char, digit-1) // one char already appended
			digits = ""
			char = ""
		}

		if unicode.IsLetter(r) || unicode.IsSpace(r) || esc {
			char = string(r)
			result += char
			esc = false
		}
	}

	return result
}
