package main

import "fmt"

//remove duplicate in constant space O(1). its ok to leave elements leave behind after removal
// return length of the new deduped array
func removeDuplicateSorted(a []int) int {
	if len(a) == 0 {
		return 0
	}
	curr := 1
	for i := 1; i < len(a); i++ {
		if a[i] != a[i-1] {
			a[curr] = a[i]
			curr++
		}

	}
	return curr
}

func TestRemoveDuplicatedSorted() {
	cases := []struct {
		in  []int
		out int
	}{
		{[]int{1}, 1},
		{[]int{1, 1}, 1},
		{[]int{1, 1, 2, 2, 2, 3}, 3},
	}

	for _, c := range cases {
		got := removeDuplicateSorted(c.in)
		if got != c.out {
			fmt.Printf("test failed. %v, expected %v, got %v\n", c.in, c.out, got)
		}
	}
}

// Given a non-negative number represented as an array of digits, plus one to the number.

// The digits are stored such that the most significant digit is at the head of the list.

func plusOne(a []int) int {
	sum := 0
	for i := 0; i < len(a); i++ {
		sum = (sum * 10) + a[i]
	}
	return sum + 1
}

func plusOne1(a []int) {
	for i := len(a) - 1; i >= 0; i-- {
		if a[i] < 9 {
			a[i] = a[i] + 1
			return
		}
		a[i] = 0

	}
	if a[0] == 0 {
		a[0] = 1
	}
}

func TestPlusOne() {
	if r := plusOne([]int{1, 2, 3}); r != 124 {
		fmt.Println("Sum one failed for:{1,2,3}, expected: 124, got:", r)
	}
}

func TestPlusOne1() {
	a := []int{1, 2, 3}
	fmt.Println("Sum one 1 for:", a, ", expected: 124")
	plusOne1(a)
	fmt.Println(" got:", a)
}

//Given two sorted integer arrays nums1 and nums2, merge nums2 into nums1 as one sorted array.
func mergeSortedToOne(a, b []int) {
	aLast := len(a) - 1
	for a[aLast] == 0 {
		aLast--
	}
	k := len(a) - 1
	bLast := len(b) - 1

	// made mistake here. I had || here. it should be &&
	for bLast >= 0 && aLast >= 0 {
		if a[aLast] > b[bLast] {
			a[k] = a[aLast]
			k--
			aLast--
		} else {
			a[k] = b[bLast]
			k--
			bLast--
		}
	}
	if bLast >= 0 {
		for bLast >= 0 {
			a[k] = b[bLast]
			k--
			bLast--
		}
	}
}

func TestMergeSortedToOne() {
	a := []int{3, 5, 6, 0, 0, 0, 0}
	b := []int{2, 4, 7, 9}
	fmt.Println("sort two arrays A:", a, "B:", b)
	mergeSortedToOne(a, b)
	fmt.Println("result: ", a)

}

// Say you have an array for which the ith element is the price of a given stock on day i.

// If you were only permitted to complete at most one transaction (ie, buy one and sell one share of the stock), design an algorithm to find the maximum profit.

// Example 1:
// Input: [7, 1, 5, 3, 6, 4]
// Output: 5

// max. difference = 6-1 = 5 (not 7-1 = 6, as selling price needs to be larger than buying price)
// Example 2:
// Input: [7, 6, 4, 3, 1]
// Output: 0

// In this case, no transaction is done, i.e. max profit = 0.

func maxProfit(stock []int) int {
	if len(stock) == 0 {
		return 0
	}
	buy, profit := stock[0], 0
	for i := 1; i < len(stock); i++ {
		buy = min(buy, stock[i])
		profit = max(profit, stock[i]-buy)
	}
	return profit
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func TestMaxProfit() {
	cases := []struct {
		in  []int
		out int
	}{
		{[]int{7, 1, 5, 3, 6, 4}, 5},
		{[]int{7, 6, 4, 3, 1}, 0},
	}

	for _, c := range cases {
		fmt.Println("Maxprofit of ", c.in, " expected:", c.out, " got:", maxProfit(c.in))
	}
}

// Given an array of size n, find the majority element. The majority element is the element that appears more than ⌊ n/2 ⌋ times.
// You may assume that the array is non-empty and the majority element always exist in the array.
// Moor's voting method. Each different element negates the majority element occurance
// increase counter when same number found.
// decrease counter when diff number found.
// if count sets to 0 , reset the element.
func majorityNum(list []int) int {
	if len(list) == 1 { // majority element does exists in array
		return list[0]
	}
	major, count := list[0], 1
	for i := 1; i < len(list); i++ {
		if count == 0 {
			count++
			major = list[i]
		} else if major == list[i] {
			count++
		} else {
			count--
		}
	}
	return major
}

func TestMajor() {
	cases := []struct {
		in  []int
		out int
	}{
		{[]int{7, 1, 8, 7, 7, 7}, 7}, //should apprear more than n/2 ~ > 3
		{[]int{4, 6, 4, 4, 1}, 4},    // should appear more than n/2 ~ > 2
	}

	for _, c := range cases {
		fmt.Println("Majority of ", c.in, " expected:", c.out, " got:", majorityNum(c.in))
	}
}

// Major element where element appear more than n/3
func majorityNum3(list []int) []int {
	if len(list) == 1 { // majority element does exists in array
		return list
	}
	major1, count1, major2, count2 := list[0], 1, list[0], 1
	for i := 1; i < len(list); i++ {
		if count1 == 0 {
			count1++
			major1 = list[i]
		} else if count2 == 0 {
			count2++
			major2 = list[i]
		} else if major1 == list[i] {
			count1++
		} else if major2 == list[i] {
			count2++
		} else {
			count1--
			count2--
		}
	}
	c1, c2 := 0, 0
	for i := 0; i < len(list); i++ {
		if list[i] == major1 {
			c1++
		} else if list[i] == major2 {
			c2++
		}
	}

	var majors []int = make([]int, 0)
	if c1 > len(list)/3 {
		majors = append(majors, major1)
	}
	majors = append(majors, major2)
	return majors
}

// Rotate an array of n elements to the right by k steps.
// For example, with n = 7 and k = 3, the array [1,2,3,4,5,6,7] is rotated to [5,6,7,1,2,3,4].

func rotateArray(a []int, k int) {
	temp := make([]int, 0, k)
	nk := len(a) - k // move n-k element to right.
	for k := 0; k < len(a) && k < nk; k++ {
		temp = append(temp, a[k])
	}

	var i int
	for ; nk < len(a); i++ {
		a[i] = a[nk]
		nk++
	}

	for t := 0; i < len(a); t++ {
		a[i] = temp[t]
		i++
	}
}

// order O(1) space.
func rotateArray2(a []int, k int) {
	nk := len(a) - k
	reverse(a, 0, nk-1)
	reverse(a, nk, len(a)-1)
	reverse(a, 0, len(a)-1)
}

func reverse(a []int, l, r int) {
	for l < r {
		a[l], a[r] = a[r], a[l]
		l++
		r--
	}
}

func TestRotateArray() {
	cases := []struct {
		in     []int
		rotate int
		out    []int
	}{
		{[]int{7, 1, 5, 3, 6, 4}, 3, []int{3, 6, 4, 7, 1, 5}},
		{[]int{7, 6, 4, 3, 1}, 4, []int{6, 4, 3, 1, 7}},
	}

	for _, c := range cases {

		fmt.Println("rotate of ", c.in, " by:", c.rotate, " expected:", c.out)
		rotateArray(c.in, c.rotate)
		fmt.Println("got:", c.in)
	}
}

func TestRotateArray2() {
	cases := []struct {
		in     []int
		rotate int
		out    []int
	}{
		{[]int{7, 1, 5, 3, 6, 4}, 3, []int{3, 6, 4, 7, 1, 5}},
		{[]int{7, 6, 4, 3, 1}, 4, []int{6, 4, 3, 1, 7}},
	}

	for _, c := range cases {

		fmt.Println("rotate O(1) space of  ", c.in, " by:", c.rotate, " expected:", c.out)
		rotateArray2(c.in, c.rotate)
		fmt.Println("got:", c.in)
	}
}
func main() {
	TestRemoveDuplicatedSorted()
	TestPlusOne()
	TestPlusOne1()
	TestMergeSortedToOne()
	TestMaxProfit()
	TestMajor()
	TestRotateArray()
	TestRotateArray2()
}
