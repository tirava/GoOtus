/*
 * HomeWork-6: FileCopy utility like dd
 * Created on 07.10.2019 17:26
 * Copyright (c) 2019 - Eugene Klimov
 */

package main

import (
	"io"
	"io/ioutil"
	"os"
	"testing"
)

var testCases = []struct {
	writer           io.Writer
	fromFile, toFile string
	offset, limit    int64
	expError         bool
	description      string
	checkStr         string
}{
	{
		ioutil.Discard,
		"in_tests.txt",
		"out_tests.txt",
		0, 0,
		false,
		"simple copy all file with no limit and offset",
		"",
	},
	{
		ioutil.Discard,
		"fake.fak",
		"out_tests.txt",
		0, 0,
		true,
		"error open source file",
		"",
	},
	{
		ioutil.Discard,
		"in_tests.txt",
		"",
		0, 0,
		true,
		"error create destination file",
		"",
	},
	{
		ioutil.Discard,
		"in_tests.txt",
		"out_tests.txt",
		100, 200,
		false,
		"copy 200 bytes from offset 100",
		"",
	},
	{
		ioutil.Discard,
		"in_tests.txt",
		"out_tests.txt",
		23, 25,
		false,
		"copy 25 bytes from offset 23",
		"There are seven days of t",
	},
	{
		ioutil.Discard,
		"in_tests.txt",
		"out_tests.txt",
		2340, 250,
		false,
		"copy 250 bytes from offset 2340 but end of file and get less bytes",
		"",
	},
}

func TestCopyFileSeekLimit(t *testing.T) {
	for _, test := range testCases {
		actualBytes, err := CopyFileSeekLimit(test.writer, test.toFile, test.fromFile, test.offset, test.limit)

		// check if returned errors but must not be
		if err != nil {
			if !test.expError {
				t.Errorf("FAIL %s - CopyFileSeekLimit() returns error = '%s', expected = 'no error'.",
					test.description, err)
			} else {
				t.Logf("PASS CopyFileSeekLimit - %s", test.description)
			}
			continue
		}

		// check if no returned errors but must be
		if test.expError {
			t.Errorf("FAIL %s - CopyFileSeekLimit() returns = 'no error', expected error = '%s'.",
				test.description, err)
			continue
		}

		// check test file exists
		from, err := os.Open(test.fromFile)
		if err != nil {
			t.Fatalf("can't open test file: %s", err)
		}
		stat, err := from.Stat()
		if err != nil {
			t.Fatalf("can't get file stat: %s", err)
		}

		// check if length is equal for expected and output
		var expectedBytes int64
		if test.limit == 0 {
			expectedBytes = stat.Size()
		} else {
			if test.offset+test.limit <= stat.Size() {
				expectedBytes = test.limit
			} else {
				expectedBytes = stat.Size() - test.offset
			}
		}
		if expectedBytes != actualBytes {
			t.Errorf("FAIL %s - CopyFileSeekLimit() returns bytes = '%d', expected bytes = '%d'.",
				test.description, actualBytes, expectedBytes)
			continue
		}

		// check output test string and example is equal
		if test.checkStr != "" {
			to, err := os.Open(test.toFile)
			if err != nil {
				t.Fatalf("can't open test file: %s", err)
			}
			b, err := ioutil.ReadAll(to)
			if err != nil {
				t.Fatalf("can't read test file: %s", err)
			}
			s := string(b)
			if s != test.checkStr {
				t.Errorf("FAIL %s - CopyFileSeekLimit() copied data = '%s', expected data = '%s'.",
					test.description, s, test.checkStr)
				continue
			}
		}

		t.Logf("PASS CopyFileSeekLimit - %s", test.description)
	}
}
