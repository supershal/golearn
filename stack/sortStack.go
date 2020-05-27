package main

import (
	"fmt"
	"strconv"
)

func sortStack(orig []int) []int {
	var sorted = make([]int, 0)

	for len(orig) != 0 {
		lastIndex := len(orig) - 1
		origTop := orig[lastIndex]
		orig = orig[:lastIndex]
		for len(sorted) != 0 {
			sortedTop := sorted[len(sorted)-1]
			if origTop <= sortedTop {
				break
			}
			orig = append(orig, sortedTop)
			sorted = sorted[0 : len(sorted)-1]
		}
		sorted = append(sorted, origTop)
	}
	return sorted
}

func printStack(stack []int) string {
	s := ""
	for i := len(stack) - 1; i >= 0; i-- {
		s += strconv.Itoa(stack[i]) + ","
	}
	return s
}

func TestSortedStack() {
	orig := []int{7, 1, 6, 3, 4, 5}
	fmt.Println("original:", orig)
	result := sortStack(orig)
	fmt.Println("sorted:", printStack(result))
}
