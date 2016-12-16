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

<<<<<<< 27a95d4fccd6ffe397684df8a8d40f7e699b320b
=======
type QStack struct {
	primary, secondary []int
}

func newQStack(capacity int) *QStack {
	return &QStack{
		primary:   make([]int, 0, capacity),
		secondary: make([]int, 0, capacity),
	}
}

// pop from primary queue
func (s *QStack) pop() int {
	if len(s.primary) == 0 {
		fmt.Println("Stack empty")
		return -1
	}
	val := s.primary[0]
	s.primary = s.primary[1:]
	return val
}

// push to secondary queue
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

>>>>>>> string program
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
