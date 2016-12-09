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

func TestPlusOne() {
	if r := plusOne([]int{1, 2, 3}); r != 124 {
		fmt.Println("Sum one failed for:{1,2,3}, expected: 124, got:", r)
	}
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

func main() {
	TestRemoveDuplicatedSorted()
	TestPlusOne()
	TestMergeSortedToOne()
}
