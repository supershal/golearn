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

type QStack struct {
	primary, secondary []int
}

func newQStack(capacity int) *QStack {
	return &QStack{
		primary:   make([]int, 0, capacity),
		secondary: make([]int, 0, capacity),
	}
}

func (s *QStack) pop() int {
	if len(s.primary) == 0 {
		fmt.Println("Stack empty")
		return -1
	}
	val := s.primary[0]
	s.primary = s.primary[1:]
	return val
}

func (s *QStack) push(x int) {
	if len(s.secondary) == cap(s.secondary) {
		fmt.Println("stack overflow")
		return
	}
	s.secondary = append(s.secondary, x)
	for len(s.primary) != 0 {
		s.secondary = append(s.secondary, s.primary[0])
		s.primary = s.primary[1:]
	}
	s.primary, s.secondary = s.secondary, s.primary
}

func TestQstack() {
	stack := newQStack(5)
	stack.push(1)
	stack.push(2)
	fmt.Println("expected pop: 2, got:", stack.pop())
	fmt.Println("expected pop: 1, got:", stack.pop())
	fmt.Println("expected pop: -1, got:", stack.pop())

}

func main() {
	TestQstack()
	//TestSortedStack()
}
