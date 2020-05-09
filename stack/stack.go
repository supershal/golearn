package main

import (
	"fmt"
	"os"
	"strconv"
)

type stack struct {
	s   []int
	top int
}

func NewStack(len int) *stack {
	return &stack{make([]int, len), -1}
}

func (st *stack) push(x int) error {
	if st.top == len(st.s)-1 {
		return fmt.Errorf("overflow")
	}
	st.top++
	st.s[st.top] = x
	return nil
}

func (st *stack) pop() (int, error) {
	if st.top < 0 {
		return -1, fmt.Errorf("underflow")
	}
	val := st.s[st.top]
	st.top--
	return val, nil
}

func (st *stack) isEmpty() bool {
	return st.top == -1
}

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

func main() {
	tests := map[string]func(){
		"mergeTwoStack":    TestTwoStacks,
		"queueUsingStacks": TestQueueUsingStacks,
		"nstack":           TestNStack,
	}
	//TestSortedStack()
	//TestTwoStacks()
	filename := os.Args[1]
	testFunc, _ := tests[filename]
	testFunc()
}
