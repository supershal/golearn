package main

import "fmt"

// https://www.byte-by-byte.com/reversestack/
// Given a stack, reverse the items without creating any additional data structures
/*
 can not use map or array.
 can use internal stack by using recursion.

 [1,2,3] -> [3,2,1]
 1) pop elements from the stack and store it in the funciton stack.
 2) but if you push in the same order of pop then it will restore to itself.
 3) since we want to push last popped element in the bottom again we reverse it.
*/

func reverse(s *stack) {
	if s.isEmpty() {
		return
	}
	top, _ := s.pop()
	reverse(s)
	insertAtBottom(s, top)
}

// 1, 2
func insertAtBottom(s *stack, x int) {
	if s.isEmpty() {
		s.push(x)
		return
	}
	top, _ := s.pop() // 2 -> top ==1; [2, 1]
	insertAtBottom(s, x)
	s.push(top)
}

func TestReverseStack() {
	s := NewStack(3)
	s.push(1)
	s.push(2)
	s.push(3)
	fmt.Println("before revese:")
	s.print()
	reverse(s)
	fmt.Println("after revese:")
	s.print()
}
