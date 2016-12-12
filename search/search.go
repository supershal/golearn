package main

import (
	"fmt"
	"math"
)

// Given two arrays, write a function to compute their intersection.

// Example:
// Given nums1 = [1, 2, 2, 1], nums2 = [2, 2], return [2, 2].

// Note:
// Each element in the result should appear as many times as it shows in both arrays.
// The result can be in any order.

//unordered list. so use map to count numbers.
func inersection(a []int, b []int) (c []int) {
	counts := make(map[int]int)
	for i := 0; i < len(a); i++ {
		if counts[a[i]] != 0 {
			counts[a[i]] = counts[a[i]] + 1
		} else {
			counts[a[i]] = 1
		}
	}

	for i := 0; i < len(b); i++ {
		if counts[b[i]] != 0 {
			c = append(c, b[i])
			counts[b[i]] = counts[b[i]] - 1
		}
	}
	return c
}

// sqrt

func sqrt(n int) int {
	start, end := 0, n
	mid := (start + end) / 2
	for math.Abs(float64((mid*mid)-n)) > 0.00001 {
		if mid*mid == n {
			return mid
		}
		if mid*mid < n {
			start = mid + 1
		}
		end = mid
		mid = (start + end) / 2
	}
	return start
}

func searchBinary(a []int, x int) bool {
	if bst(a, 0, len(a)-1, x) == -1 {
		return false
	}
	return true
}

func bst(a []int, low, high, target int) int {
	if low > high {
		return -1
	}
	mid := (low + high) / 2
	if a[mid] == target {
		return mid
	}

	if a[mid] < target {
		return bst(a, mid+1, high, target)
	}
	return bst(a, low, mid-1, target)
}

func TestBST() {
	cases := []struct {
		in     []int
		target int
		out    bool
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7}, 7, true},
		{[]int{1, 3, 4}, 7, false},
	}

	for _, c := range cases {
		fmt.Println("BST search of ", c.in, " search:", c.target, " expected:", c.out, " got:", searchBinary(c.in, c.target))
	}
}

// search lowest rotated.
// cases: not rotated. if low < midval < high, return low
// cases : one element: return low
// cases: else
//    1. lower can be in first half or second half.
//    if midval < lowval, its in first half
//    if midval > highval, its in second half
func searchLowestRotated(a []int) int {
	if len(a) == 0 {
		return -1
	}
	return a[lowestRotated(a, 0, len(a)-1)]
}

func lowestRotated(a []int, low, high int) int {
	if low == high {
		return low
	}
	mid := (low + high) / 2
	// important. <= condition. for array of size 2
	// array is not rotated.
	if a[low] <= a[mid] && a[mid] < a[high] {
		return low
	}
	// array is rotated.
	// important. check base condition before going left or right.
	if a[mid] < a[mid-1] {
		return mid
	}
	if a[mid] > a[mid+1] {
		return mid + 1
	}

	if a[mid] > a[high] {
		return lowestRotated(a, mid+1, high)
	}
	return lowestRotated(a, low, mid-1)
}

// since array is rotated, lower element is always going to be between an rotated array where a[low] > a[high]
// half the array until a[low] is not greater than a[high] . when this happens low would be the minimum element.
// if midval > highval that means minimum is in right half.
// else element is in left half.
func lowestRotatedOptimized(a []int) int {
	low, high := 0, len(a)-1
	for a[low] > a[high] {
		mid := (low + high) / 2
		if a[mid] > a[high] {
			low = mid + 1
		} else {
			// important. it should not be mid-1. check case where 1 is the middle element.{4,5,1,2,3}
			high = mid
		}
	}
	return a[low]
}

func TestSearchLowestRotated() {
	cases := []struct {
		in  []int
		out int
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7}, 1},
		{[]int{7, 8, 9, 10, 11, 1, 3, 4}, 1},
		{[]int{10, 11, 1, 3, 4, 7, 8, 9}, 1},
		{[]int{3, 4, 7, 8, 9, 10, 11, 1}, 1},
	}

	for _, c := range cases {
		fmt.Println("Lowest of ", c.in, " expected:", c.out, " got:", searchLowestRotated(c.in))
	}

	for _, c := range cases {
		fmt.Println("Optimized: Lowest of ", c.in, " expected:", c.out, " got:", lowestRotatedOptimized(c.in))
	}
}

// x can be middle
// find pivot element that is the lowest of heighest element in the array.
// divide array into two sorted halves. search element in both half one by one.
// below method is using pivot. does two BST. it can be done with one BST.
func searchRotated(a []int, x int) int {
	low, high := 0, len(a)-1
	lowest := lowestRotated(a, low, high)
	if x >= a[lowest] && x <= a[high] {
		indexX := bst(a, lowest, high, x)
		return a[indexX]
	}
	indexX := bst(a, low, lowest-1, x)
	return a[indexX]
}

// find middle element.
// if midval > lowval, left is sorted.
// else right is sorted.

// if element is found in sorted array then reduce lower or upper bound.
// else go left or right where the array is not sorted.
// search until middle is the element being searched.

func searchRotatedOptimized(a []int, x int) int {
	low, high := 0, len(a)-1

	// important: forgot to put <=, used < instead. cases failed {7, 8, 9, 10, 11, 1, 3, 4} and serach 9
	for low <= high {
		mid := (low + high) / 2
		if a[mid] == x {
			return a[mid]
		}

		if a[mid] > a[low] { // left is sorted.
			if x >= a[low] && x < a[mid] {
				high = mid - 1
			} else {
				low = mid + 1
			}
		} else { // right is sorted
			if x > a[mid] && x <= a[high] {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}

	}
	return -1
}

func TestSearchRotated() {
	cases := []struct {
		in     []int
		target int
		out    int
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7}, 6, 6},
		{[]int{7, 8, 9, 10, 11, 1, 3, 4}, 9, 9},
		{[]int{10, 11, 1, 3, 4, 7, 8, 9}, 8, 8},
		{[]int{3, 4, 7, 8, 9, 10, 11, 1}, 9, 9},
	}

	for _, c := range cases {
		fmt.Println("search in rotated of ", c.in, " expected:", c.out, " got:", searchRotated(c.in, c.target))
	}

	for _, c := range cases {
		fmt.Println("Optimized: search in rotated of ", c.in, " expected:", c.out, " got:", searchRotatedOptimized(c.in, c.target))
	}
}

func searchRanges(a []int, x int) (int, int) {
	rl, rh := (1<<31)-1, -1
	findRange(a, 0, len(a)-1, x, &rl, &rh)
	return rl, rh
}
func findRange(a []int, low, high, x int, rl, rh *int) {
	if low > high {
		return
	}
	mid := (low + high) / 2
	if a[mid] == x {
		if mid < *rl {
			*rl = mid
			findRange(a, low, mid-1, x, rl, rh)
		}
		if mid > *rh {
			*rh = mid
			findRange(a, mid+1, high, x, rl, rh)
		}

	} else if a[mid] < x {
		findRange(a, mid+1, high, x, rl, rh)
	}
	findRange(a, low, mid-1, x, rl, rh)
}

func TestSearchRanges() {
	cases := []struct {
		in     []int
		target int
		lower  int
		higher int
	}{
		{[]int{1, 2, 5, 5, 5, 6, 7}, 5, 2, 4},
		{[]int{1, 2, 3, 4, 5, 6, 7}, 2, 1, 1},
		{[]int{1, 1, 1, 1, 4, 7, 7, 9}, 1, 0, 3},
		{[]int{3, 4, 7, 8, 9, 10, 11, 11}, 11, 6, 7},
	}

	for _, c := range cases {
		l, h := searchRanges(c.in, c.target)
		fmt.Println("search in rotated of ", c.in, " expected:", c.lower, ",", c.higher, " got:", l, ",", h)
	}

}

// func searchRotatedDuplicate(a []int, x int) int{}
// func searchRanges(a []int) (int, int){}

func main() {
	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
	TestBST()
	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
	TestSearchLowestRotated()
	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
	TestSearchRotated()
	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
	TestSearchRanges()
}
