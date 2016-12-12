package main

import (
	"bytes"
	"fmt"
)

// Given a non-empty string check if it can be constructed by taking a substring of it and appending multiple copies of the substring together. You may assume the given string consists of lowercase English letters only and its length will not exceed 10000.
// Example 1:
// Input: "abab"
// Output: True
// Explanation: It's the substring "ab" twice.
// Example 2:
// Input: "aba"
// Output: False
// Example 3:
// Input: "abcabcabcabc"
// Output: True
// Explanation: It's the substring "abc" four times. (And the substring "abcabc" twice.)

// find substring first: loop until first char is diff then next char.
// then check length. if it is multiple of the substring
// then compare take next len(substring) chars and compare with the substring.
// however this doest not work for when there are repeated chars in the string[aaabaaabaaab]

func repeatedSubString(s string) bool {
	if len(s) == 0 {
		return false
	}
	prefix := findPrefix(s)
	if (len(s) % len(prefix)) != 0 {
		return false
	}
	for p := len(prefix); p < len(s); p += len(prefix) {
		// important. made a mistake here. the higher bound is p+len(prefix) not len(prefix)
		next := s[p : p+len(prefix)]
		if prefix != next {
			return false
		}
	}
	return true
}

func findPrefix(s string) string {
	var prefix string = string(s[0])
	first := s[0]
	for i := 1; i < len(s); i++ {
		if s[i] == first {
			break
		}
		prefix = prefix + string(s[i])
	}

	return prefix
}

//correct version.
// the len(string) should be multiple of len(substring)
// start from half the string and start adding substring to itself upto the devisor.
// break when main string and substring are equal.

func repeatedSubStringOptimized(s string) bool {
	if len(s) == 0 {
		return false
	}
	l := len(s)
	// i should not reach to 0 to ensure repeation.
	for i := l / 2; i >= 1; i-- {
		if l%i == 0 {
			multiple := l / i
			sub := s[:i]
			buf := bytes.NewBufferString(sub)
			for j := 1; j < multiple; j++ {
				buf.WriteString(sub)
			}
			if buf.String() == s {
				return true
			}
		}
	}
	return false
}

// TODO: compare using kmp

func TestRepeatedSubString() {
	cases := []struct {
		in  string
		out bool
	}{
		{"", false},
		{"a", true},
		{"ab", false},
		{"abc", false},
		{"abab", true},
		{"aaabaaabaaab", true}, // important case.
	}

	for _, c := range cases {
		fmt.Println("RepeatedSubSting of ", c.in, "expected:", c.out, "got:", repeatedSubString(c.in))
	}

	for _, c := range cases {
		fmt.Println("Optimized RepeatedSubSting of ", c.in, "expected:", c.out, "got:", repeatedSubStringOptimized(c.in))
	}
}

func main() {
	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
	TestRepeatedSubString()
}
