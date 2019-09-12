/*
 * HomeWork-3: Frequency Analysis
 * Created on 13.09.19 19:04
 * Copyright (c) 2019 - Eugene Klimov
 */

// Package frequency_analysis implements counting most popular words.
package frequency_analysis

import (
	"regexp"
	"sort"
	"strings"
)

// Frequency is the base type for counting words.
type Frequency map[string]int

// pairs need for sort map
type pairs struct {
	key   string
	value int
}

// reg is global - for best benchmark
var reg = regexp.MustCompile("[a-z0-9а-яё]+('[a-z0-9а-яё])*")

// WordCount returns the frequencies of words in a string for most popular 'num' words.
func WordCount(s string, num int) Frequency {

	result := Frequency{}
	s = strings.ToLower(s)

	// regexp instead strings.Split - for clean words
	words := reg.FindAllString(s, -1)

	// count words
	for _, c := range words {
		result[c]++
	}

	// no need sort for words <= num
	if len(result) <= num {
		return result
	}

	// sort map if result > num
	sortSlice := make([]pairs, 0, len(result))
	for k, v := range result {
		sortSlice = append(sortSlice, pairs{k, v})
	}

	// sort by count then by word
	sort.Slice(sortSlice, func(i, j int) bool {
		if sortSlice[i].value == sortSlice[j].value {
			return sortSlice[i].key < sortSlice[j].key
		}
		return sortSlice[i].value > sortSlice[j].value
	})

	// delete words > num
	for i := num; i < len(sortSlice); i++ {
		delete(result, sortSlice[i].key)
	}

	return result
}
