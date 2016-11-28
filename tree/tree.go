package main

import (
	"bytes"
	"fmt"
	"strconv"
)

type node struct {
	value int
	left  *node
	right *node
}

func newTree(val int, left, right *node) *node {
	return &node{val, left, right}
}

func printTree(root *node) string {
	if root == nil {
		return ""
	}
	return ""
}

func levelOrder(root *node) string {
	if root == nil {
		return ""
	}
	var sb bytes.Buffer
	queue := make([]*node, 0)
	queue = append(queue, root)
	for len(queue) != 0 {
		current := queue[0]
		sb.WriteString(strconv.Itoa(current.value))
		sb.WriteString("->")
		if current.left != nil {
			queue = append(queue, current.left)
		}
		if current.right != nil {
			queue = append(queue, current.right)
		}
		queue = queue[1:]
	}
	return sb.String()
}

func TestLevelOrder() {
	two := newTree(2, nil, nil)
	three := newTree(3, nil, nil)
	root := newTree(1, two, three)

	fmt.Println("level order: ", levelOrder(root))
}

func main() {
	TestLevelOrder()
}
