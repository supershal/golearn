package main

import (
	"fmt"
	"os"
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

func (st *stack) print() {
	fmt.Println(st.s)
}
func main() {
	tests := map[string]func(){
		"mergeTwoStack":    TestTwoStacks,
		"queueUsingStacks": TestQueueUsingStacks,
		"nstack":           TestNStack,
		"sortStack":        TestSortedStack,
		"reverseStack":     TestReverseStack,
	}
	//TestSortedStack()
	//TestTwoStacks()
	filename := os.Args[1]
	testFunc, _ := tests[filename]
	testFunc()
}
