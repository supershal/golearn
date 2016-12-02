package main

import "fmt"

//Given a list of bytes a, each representing one byte of a larger integer (ie. {0x12, 0x34, 0x56, 0x78} represents the integer 0x12345678), and an integer b, find a % b.
//http://www.byte-by-byte.com/bigintmod/

func mod(n []byte, a int) int {
	m := 0
	for i := 0; i < len(n); i++ {
		m = m << 8
		m = m + int(n[i]&0xFF)
		m = m % a
	}
	return m
}

func TestModulo() {
	b := []byte{0x03, 0xED}
	a := 10
	fmt.Println(b)
	fmt.Println("modulo:", mod(b, a))
}

func fizzbuzz(n int) {
	for i := 1; i <= n; i++ {
		if i%15 == 0 {
			fmt.Println("fizzbuzz")
		} else if i%3 == 0 {
			fmt.Println("fizz")
		} else if i%5 == 0 {
			fmt.Println("buzz")
		} else {
			fmt.Println(i)
		}
	}

}

func TestFizzbuzz() {
	fizzbuzz(16)
}

//Given an input amount of change x, write a function to determine the minimum number of coins required to make that amount of change.
// assume american coins. [1,5,10,25]
// change > coin[i], change < coin[i], change = coin[i], change = 0
// 27 = 25 * 1 + 1 *2 = 3
// 25 = 25 *1
// 16 = 10 *1 + 5 *1 + 1 *1 = 3
// 0 = 0
// sort coins in descending order
// find smallest element m > change
// count = change / m
// change = change % m
// reapeat

// coints = {25, 5, 10, 1}
func smallestChange(n int) int {
	coins := []int{25, 10, 5, 1}
	var count int
	for i := 0; i < len(coins) && n > 0; i++ {
		if coins[i] > n {
			continue
		}
		count = count + (n / coins[i])
		n = n % coins[i]
	}
	return count
}

func TestSmallestChange() {
	cases := []struct {
		in, out int
	}{
		{0, 0},
		{27, 3},
		{25, 1},
		{16, 3},
	}
	for _, c := range cases {
		fmt.Println("In:", c.in, "Expected:", c.out, "result:", smallestChange(c.in))
	}
}

//Given an integer, write a function to compute the number of ones in the binary representation of the number.
func sumOnes(n int) int {
	var sum int
	for n > 0 {
		sum = sum + n&1
		n = n >> 1
	}
	return sum
}

func TestSumOnes() {
	fmt.Println("sumones(7)=", sumOnes(7))
}

//Given 2 sorted arrays, A and B, where A is long enough to hold the contents of A and B, write a function to copy the contents of B into A without using any buffer or additional memory.
//http://www.byte-by-byte.com/mergearrays/
func MergeTwoSortedArray(A, B []int) []int {
	if len(A) == 0 {
		return A
	}
	var aMax int
	for i := 0; i < len(A); i++ {
		if A[i] == 0 {
			aMax = i - 1
			break
		}
	}
	aLast, bMax := len(A)-1, len(B)-1
	fmt.Println(aMax, aLast, bMax)
	for ; aMax >= 0 && bMax >= 0; aLast-- {
		if A[aMax] > B[bMax] {
			A[aLast] = A[aMax]
			aMax--
		} else if B[bMax] > A[aMax] {
			A[aLast] = B[bMax]
			bMax--
		} else {
			A[aLast] = B[aMax]
			aLast--
			A[aLast] = A[aMax]
			aMax--
			bMax--
		}
	}
	if aMax == 0 {
		for bMax > 0 {
			A[aLast] = B[bMax]
			aLast--
			bMax--
		}
	}

	return A
}

func TestMergeTwoSortedArray() {
	A := []int{1, 3, 0, 0, 0}
	B := []int{1, 4, 6}
	fmt.Println("A=", A, "B=", B)
	C := MergeTwoSortedArray(A, B)
	fmt.Println("Result=", C)
}

// Write function to find permutation of string
// permutation: n! .
// if dupliactes then permutations are n!/(a! * b! *...) where a , b are number of duplicates
// {a} => {a}
// {ab} => {ab,ba}
//{abc} =>{ abc, acb, bac, bca, cab, cba}
// logic: taken one char and append permutation of rest of chars to it. and then iterate through it.
// abc
//  /
// a, bc
//   /   \
// ab, c  ac, b
//     /      \
//    abc, nil acb, nil
func stringPermutation(s string) []string {
	result := make([]string, 0)
	permutation("", s, &result)
	return result
}

func permutation(prefix, suffix string, results *[]string) {
	if len(suffix) == 0 {
		*results = append(*results, prefix)
	}
	for i := 0; i < len(suffix); i++ {
		permutation(prefix+string(suffix[i]), suffix[:i]+suffix[i+1:], results)
	}
}

func TestPermuatation() {
	fmt.Println("permuatation of abc", stringPermutation("abc"))
}

func fibo(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fibo(n-1) + fibo(n-2)
}

func TestFibo() {
	fmt.Println("6th fibo. expected: 8, got:", fibo(6))
	fmt.Println("10th fibo. expected: 55, got:", fibo(10))
}

// Nested weighted sum
// Given a nested list of integers, return the sum of all integers in the list weighted by their depth.
// Each element is either an integer, or a list -- whose elements may also be integers or other lists.
// Example 1:
// Given the list [[1,1],2,[1,1]], return 10. (four 1's at depth 2, one 2 at depth 1)

// its a slice of integers or list. so elements can be either int or []int or slice of interface
// recursively look through elements. if an element is integer its depth is current.
// first element's depth is 1.
// if it is an list then its depth is depth+1, 2 if the the list has a list then its depth is currdepth +1
func nestedWeightedSum(list []interface{}, depth int) int {
	sum := 0
	for _, l := range list {
		if data, ok := l.(int); ok {
			sum += data * depth
		} else if dataList, ok := l.([]interface{}); ok {
			sum += nestedWeightedSum(dataList, depth+1)
		} else {
			fmt.Println("invalid element")
			return -1
		}
	}
	return sum
}

func TestnestedWeightedSum() {
	cases := []struct {
		in  []interface{}
		out int
	}{
		{[]interface{}{}, 0},
		{[]interface{}{1, 2, 3}, 6},
		{
			[]interface{}{[]interface{}{1, 1}, 2, []interface{}{1, 1}},
			10},
	}
	for _, c := range cases {
		fmt.Println("Nested Weighted Sum of ", c.in, "expected:", c.out, "got:", nestedWeightedSum(c.in, 1))
	}
}

// find middle element in rotated array
func findMinRotated(arr []int, low, high int) int {
	if len(arr) == 0 {
		return -1
	}
	if low == high {
		return arr[low]
	}
	if arr[low] < arr[high] {
		return arr[low]
	}
	mid := (low + high) / 2
	return min(findMinRotated(arr, low, mid), findMinRotated(arr, mid+1, high))
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func TestfindMinRotated() {
	cases := []struct {
		in  []int
		out int
	}{
		{[]int{}, -1},
		{[]int{1}, 1},
		{[]int{4, 5, 6, 7, 8, 1, 2, 3}, 1},
		{[]int{4, 5, 6, 7, 8, 1, 2, 3, 4}, 1},
	}

	for _, c := range cases {
		fmt.Println("Min of rotated array:", c.in, "expected:", c.out, "got:", findMinRotated(c.in, 0, len(c.in)-1))
	}
}

// Implement an algorithm to determine if a string has all unique characters. What if you
// can not use additional data structures?
func uniqueString(s string) bool {
	if len(s) <= 1 {
		return true
	}
	cache := make(map[byte]bool)
	for _, b := range []byte(s) {
		if cache[b] {
			return false
		}
		cache[b] = true
	}
	return true
}

// without using any data structure.
// think of it as a bloom filter
// crate integer storage.
// set the ascii equivalent number of bit.
// do logical AND. so if the char is already there, AND will be greater than 0
func uniqueString2(s string) bool {
	if len(s) <= 1 {
		return true
	}
	hash := 0
	for _, b := range []byte(s) {
		charIndex := b - 'a'
		if (hash & (1 << charIndex)) > 0 {
			return false
		}
		hash = hash | (1 << charIndex)
	}

	return true

}

func TestUniqueStringWithMap() {
	cases := []struct {
		in  string
		out bool
	}{
		{"", true},
		{"a", true},
		{"abc", true},
		{"abca", false},
	}

	for _, c := range cases {
		fmt.Println("Unique string:", c.in, "expected:", c.out, "got:", uniqueString(c.in))
	}
	for _, c := range cases {
		fmt.Println("Unique string  without map:", c.in, "expected:", c.out, "got:", uniqueString2(c.in))
	}
}

func reverseString(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func TestReverseString() {
	cases := []struct {
		in  string
		out string
	}{
		{"", ""},
		{"a", "a"},
		{"abc", "cba"},
		{"abca", "acba"},
	}

	for _, c := range cases {
		fmt.Println("reverse string:", c.in, "expected:", c.out, "got:", reverseString(c.in))
	}
}

// Design an algorithm and write code to remove the duplicate characters in a string
// without using any additional buffer. NOTE: One or two additional variables are fine.
// An extra copy of the array is not.

func removeDupliateChars(s string) string {
	if len(s) <= 1 {
		return s
	}
	chars, result := []byte(s), ""
	checksum := 0
	for _, c := range chars {
		charIndex := c - 'a'
		charInt := (1 << charIndex)
		if (checksum & charInt) != charInt {
			result += string(c)
		}
		checksum |= charInt
	}
	return result

}

func TestRemoveDuplicateChars() {
	cases := []struct {
		in  string
		out string
	}{
		{"", ""},
		{"a", "a"},
		{"aa", "a"},
		{"abc", "abc"},
		{"abca", "abc"},
	}

	for _, c := range cases {
		fmt.Println("remove duplicate chars string:", c.in, "expected:", c.out, "got:", removeDupliateChars(c.in))
	}
}

//Write a method to replace all spaces in a string with ‘%20’.
func replaceSpaces(s []byte) []byte {
	if len(s) == 0 {
		return s
	}
	pop := s[len(s)-1]
	s = s[:len(s)-1]
	s = replaceSpaces(s)
	if pop == ' ' {
		s = append(s, '%', '2', '0')
	} else {
		s = append(s, pop)
	}
	rev := string(s)
	return []byte(rev)
}

func TestReplaceSpaces() {
	cases := []struct {
		in  string
		out string
	}{
		{"", ""},
		{"a ", "a%20"},
		{"a bc", "a%20bc"},
	}

	for _, c := range cases {
		chars := []byte(c.in)
		result := string(replaceSpaces(chars))
		fmt.Println("replace spaces:", c.in, "expected:", c.out, "got:", result)
	}
}

func main() {
	//TestModulo()
	//TestFizzbuzz()
	//TestSmallestChange()
	//TestSumOnes()

	//TestMergeTwoSortedArray()
	//TestNStack()
	TestPermuatation()
	TestFibo()
	TestMergeTwoSortedArray()
	TestnestedWeightedSum()
	TestfindMinRotated()
	TestUniqueStringWithMap()
	TestReverseString()
	TestRemoveDuplicateChars()
	TestReplaceSpaces()

}
