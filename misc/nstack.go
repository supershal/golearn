package main

import (
	"fmt"
)

// // n stack array.
// lets say array is of size 5
// N = 1

// // start from 0 to 4
//  A := {5, 10, 7, 9, 5 }
//        0, 1,  2, 3, 4
// if N = 2 then start from both end
//    A := {5, 10, 7, 9, 5 }
//          0, 1,  2, 3, 4
//          ^            ^

// If N = 3 and have to have stack operations push and pop in O(1) then
// we have to think of it as a doubly linked list where you know next and prev pointer
// and add and remove node in middle of the list in o(1) given where to insert the node.

// We will keep two arrays.
// top[N] keep track of doubly linked list head
// next[L] keep track of address of prev element. which will be next element in linked list.

// A := {5, 10, 7, 9, 5 }
//        0, 1,  2, 3, 4

// Stack := {5, 0, 0, 0, 0}
// top[0] = 0
// top[1] = -1
// top[2] = -1

// Next := {-1, 2, 3, 4, 5, -1}
//         { 0, 1, 2, 3, 4, 5}
// push(0,5)
// avail = 0, 1

// pop(0)
// current = top[0] = 0
// top[0] = next[current] = -1
// next[current] = avail = next[0] = 1
// avail = current = 0

type NStack struct {
	top      []int
	next     []int
	stack    []int
	capacity int
}

var nextFree = 0

func newNstack(noofStack, capacity int) *NStack {
	nStack := &NStack{
		top:      make([]int, noofStack),
		next:     make([]int, capacity),
		stack:    make([]int, capacity),
		capacity: capacity,
	}

	for i := 0; i < noofStack; i++ {
		nStack.top[i] = -1
	}
	for i := 0; i < capacity; i++ {
		nStack.next[i] = i + 1
	}

	nStack.next[capacity-1] = -1
	return nStack
}
func (n *NStack) push(stack, element int) {
	if stack < 0 || nextFree == -1 {
		fmt.Println("buffer overflow")
		return
	}
	currentIndex := nextFree
	nextFree = n.next[currentIndex]
	n.next[currentIndex] = n.top[stack]
	n.top[stack] = currentIndex
	n.stack[currentIndex] = element
}

func (n *NStack) pop(stack int) int {
	if n.top[stack] == -1 {
		fmt.Println("no element left to pop from stack: ", stack)
		return -1
	}

	currentIndex := n.top[stack]
	n.top[stack] = n.next[currentIndex]
	n.next[currentIndex] = nextFree
	nextFree = currentIndex
	return n.stack[currentIndex]
}

func TestNStack() {
	nStack := newNstack(3, 5)
	// {5, 6, 7, 8}
	nStack.push(0, 5)
	nStack.push(1, 6)
	nStack.push(2, 7)

	fmt.Println("pop(0) should be 5. result: ", nStack.pop(0)) //5
	nStack.push(2, 8)
	fmt.Println("pop(1) should be 6. result: ", nStack.pop(1)) //5
	nStack.push(2, 9)
	nStack.push(0, 10)
	fmt.Println("pop(2) should be 9. result: ", nStack.pop(2)) //5

}

func main() {
	TestNStack()
}
