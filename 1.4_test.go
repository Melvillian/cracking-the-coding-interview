package main

import (
	"github.com/melvillian/cracking-the-coding-interview/lib"
	"testing"
)

func TestPalindrome(t *testing.T) {
	// test positives
	positiveTestCases := [...]string{
		"",
		"a",
		"abccba",
		"1221",
		"abcbc",
		"bcbc",
	};

	for _, test := range positiveTestCases {
		if (!lib.Palindrome(test)) {
			t.Error(test);
		}
	}

	// test negatives
	negativeTestCases := [...]string{
		"blah",
		"12",
		"abc",
		"abbc",
		"aabccc",
		"laskdfj",
	};

	for _, test := range negativeTestCases {
		if (lib.Palindrome(test)) {
			t.Error(test);
		}
	}
}
