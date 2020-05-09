package main

import (
	"fmt"
	"log"
)

// https://www.byte-by-byte.com/nstacks
type nstack struct {
	s         []int // stack
	next      []int // next available index
	top       []int // top of the stacks
	nextIndex int
}

func NewNStack(len, stacks int) *nstack {
	n := &nstack{
		s:         make([]int, len),
		next:      make([]int, len),
		top:       make([]int, stacks),
		nextIndex: 0,
	}
	for ; stacks > 0; stacks-- {
		n.top[stacks-1] = -1
	}
	for i := 0; i < len; i++ {
		n.next[i] = i + 1
	}
	n.next[len-1] = -1
	return n
}

func (n *nstack) push(stackno, x int) {
	if n.nextIndex < 0 {
		log.Fatal("overflow")
	}
	curr := n.nextIndex
	n.s[curr] = x
	n.nextIndex = n.next[curr]
	n.next[curr] = n.top[stackno-1]
	n.top[stackno-1] = curr
}

func (n *nstack) pop(stackno int) int {
	curr := n.top[stackno-1]
	if curr < 0 {
		log.Fatal("underflow")
	}
	val := n.s[curr]
	n.top[stackno-1] = n.next[curr]
	n.next[curr] = n.nextIndex
	n.nextIndex = curr
	return val
}

func (n *nstack) print() {
	fmt.Println("stack: ", n.s)
	fmt.Println("top: ", n.top)
	fmt.Println("next: ", n.next)
	fmt.Println("next available: ", n.nextIndex)
}
func TestNStack() {
	n := NewNStack(4, 3)
	n.push(1, 1)
	n.push(2, 2)
	n.push(3, 3)
	n.push(1, 2)
	//n.push(2, 3) // overflow
	//n.print()
	fmt.Println(n.pop(1)) // 2
	fmt.Println(n.pop(1)) // 1
	fmt.Println(n.pop(3)) // 3
	fmt.Println(n.pop(2)) // 2
	fmt.Println(n.pop(1)) // underflow

}
