package main

import (
	"fmt"
)

type node struct {
	value      int
	prev, next *node
}

func newNode(v int) *node {
	return &node{v, nil, nil}
}

func print(root *node) {
	if root == nil {
		return
	}
	current := root

	for {
		switch {
		case current.next != root:
			fmt.Printf("%v<->", current.value)
			current = current.next
		default:
			fmt.Printf("%v<->", current.value)
			return

		}
	}

}

func concate(left, right *node) *node {
	if left == nil {
		return right
	}
	if right == nil {
		return left
	}
	leftLast := left.prev
	rightLast := right.prev

	left.prev = rightLast

	right.prev = leftLast
	leftLast.next = right

	rightLast.next = left

	return left

}

func btod(root *node) *node {
	if root == nil {
		return nil
	}
	left := btod(root.prev)
	right := btod(root.next)
	root.prev = root
	root.next = root
	return concate(concate(left, root), right)
}

func main() {
	fmt.Println("Tree to Doubly linked list")
	root := newNode(10)
	root.prev = newNode(12)
	root.next = newNode(15)
	root.prev.prev = newNode(25)
	root.prev.next = newNode(30)
	root.next.prev = newNode(36)
	print(btod(root))
}
