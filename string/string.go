package main

import (
	"bytes"
	"fmt"
	"strings"
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

func ransomNote(ransom, dict []byte) bool {
	charMap := make(map[byte]int)
	for _, b := range dict {
		charMap[b] = charMap[b] + 1
	}

	for _, r := range ransom {
		charMap[r] = charMap[r] - 1
		if charMap[r] < 0 {
			return false
		}
	}
	return true
}

// longest prefix in the array of strings.
func longestPrefix(s []string) string {
	if len(s) == 0 {
		return ""
	}
	pre := s[0]
	for i := 1; i < len(s); i++ {
		pre = prefix(pre, s[i])
		if pre == "" {
			return ""
		}
	}
	return pre
}

func prefix(first, second string) string {
	sb := bytes.Buffer{}
	for f, s := 0, 0; f < len(first) && s < len(second); f, s = f+1, s+1 {
		if first[f] != second[s] {
			break
		}
		sb.WriteString(string(first[f]))
	}
	return sb.String()
}

// case: {"foo", "foobar", "f"}, we have to scan all the string to find "f"
// think of it as a two dimention array.
// get first char of all the string and compare with first char of base string.
func longestPrefixOptimized(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	min := (1 << 31) - 1
	for _, s := range strs {
		if l := len(s); l < min {
			min = l
		}
	}
	sb := bytes.Buffer{}
	for i := 0; i < min; i++ {
		curr := strs[0][i]
		for j := 1; j < len(strs); j++ {
			if strs[j][i] != curr {
				return sb.String()
			}
		}
		sb.WriteByte(curr)
	}
	return sb.String()
}

//TODO
// func longestPrefixTrie(strs []string) string{

// }

func TestLongestPrefix() {
	cases := []struct {
		in  []string
		out string
	}{
		{[]string{"foo", "bar"}, ""},
		{[]string{"foo", "foobar", "f"}, "f"},
		{[]string{"foo", "foobar", "fooz"}, "foo"},
	}
	for _, c := range cases {
		fmt.Println("Longest Prefix of ", c.in, "expected:", c.out, "got:", longestPrefix(c.in))
	}

	for _, c := range cases {
		fmt.Println("Optimized Longest Prefix of ", c.in, "expected:", c.out, "got:", longestPrefixOptimized(c.in))
	}
}

func validateIP(s string) bool {
	if len(s) == 0 {
		return false
	}
	tokens := strings.Split(s, ".")
	if len(tokens) != 4 {
		return false
	}

	for _, t := range tokens {
		if !isValidSubnet(t) {
			return false
		}
	}
	return true
}

func isValidSubnet(s string) bool {
	if len(s) == 0 || len(s) > 3 {
		return false
	}
	if s[0] == '0' && len(s) != 1 {
		return false
	}
	num := 0
	for _, r := range s {
		if r >= '0' && r <= '9' {
			num = (num * 10) + (int(r) - '0')
			continue
		}
		return false
	}
	if num > 255 {
		return false
	}
	return true
}

func TestValidIP() {
	cases := []struct {
		in  string
		out bool
	}{
		{"0.1.2.3", false},
		{"192.168.0.1", true},
		{"255.255.255.255", true},
		{"256.256.256.256", false},
		{"2.3", false},
	}

	for _, c := range cases {
		fmt.Println("Validate IP:", c.in, "expected:", c.out, "got:", validateIP(c.in))
	}
}

func main() {
	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
	TestRepeatedSubString()
	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
	TestLongestPrefix()
	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
	TestValidIP()
}
