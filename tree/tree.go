package main

import (
	"fmt"
	"math"
	"strconv"
)

type node struct {
	value int
	left  *node
	right *node
}

func newTreeNode(value int, left, right *node) *node {
	return &node{value, left, right}
}

func insert(val int, root *node) *node {
	if root == nil {
		return newTreeNode(val, nil, nil)
	}
	if val <= root.value {
		root.left = insert(val, root.left)
		return root
	}
	root.right = insert(val, root.right)
	return root
}

func printLevel(root *node) string {
	if root == nil {
		return ""
	}
	queue := make([]*node, 0)
	level := make([]int, 0)
	queue = append(queue, root)
	prevLevel := 0
	level = append(level, 0)
	tree := "\n"
	for len(queue) != 0 {
		current := queue[0]
		curLevel := level[0]

		if curLevel != prevLevel {
			tree += "\n"
			prevLevel = curLevel
		}
		tree += strconv.Itoa(current.value) + "\t"
		if current.left != nil {
			queue = append(queue, current.left)
			level = append(level, curLevel+1)
		}
		if current.right != nil {
			queue = append(queue, current.right)
			level = append(level, curLevel+1)
		}
		queue = queue[1:]
		level = level[1:]
	}
	return tree
}

func TestPrintLevel() {
	tree := []int{4, 2, 6, 1, 3, 5, 7} // level order
	var root *node
	for i := 0; i < len(tree); i++ {
		root = insert(tree[i], root)
	}
	fmt.Println("Level order of: expected: ", tree, "Actual:", printLevel(root))

}

//Given a tree, write a function to find the length of the longest branch of nodes in increasing consecutive order.

// 			11
//    2				6
// 1		3		5		7
// 							8

func longestBranch(root *node) int {
	if root == nil {
		return 0
	}
	if root.left == nil && root.right == nil {
		return 1
	}
	var leftSum, rightSum int
	if root.left != nil {
		leftSum = longestBranch(root.left)
		if root.value < root.left.value {
			leftSum = leftSum + 1
		}
	}
	if root.right != nil {
		rightSum = longestBranch(root.right)
		if root.value < root.right.value {
			rightSum = rightSum + 1
		}
	}

	return max(leftSum, rightSum)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func TestLongestBranch() {
	tree := []int{4, 2, 6, 1, 3, 5, 7} // level order
	var root *node
	for i := 0; i < len(tree); i++ {
		root = insert(tree[i], root)
	}
	root = newTreeNode(11, root, newTreeNode(33, nil, newTreeNode(44, newTreeNode(55, nil, nil), nil)))
	fmt.Println(" longest branch expected: 11->33->44->55 = 4 Actual:", longestBranch(root))
}

//Given a binary tree, write a function to determine whether the tree is balanced.
// a binary tree is balanced if difference between its left and right tree's height is 0 or 1
func isBalanced(root *node) bool {
	if balancedHeght(root) > -1 {
		return true
	}
	return false
}

func balancedHeght(root *node) int {
	if root == nil {
		return 0
	}
	l := balancedHeght(root.left)
	r := balancedHeght(root.right)

	if l == -1 || r == -1 {
		return -1
	}
	if math.Abs(float64(l-r)) > 1 { // tree is unbalanced
		return -1
	}

	// tree's height is max(left,right) + 1
	return max(l, r) + 1
}

func TestBalanced() {
	tree := []int{4, 2, 6, 1, 3, 5, 7} // level order
	var root *node
	for i := 0; i < len(tree); i++ {
		root = insert(tree[i], root)
	}
	fmt.Println("balanced: expected true", isBalanced(root))
	root = newTreeNode(11, root, newTreeNode(33, nil, newTreeNode(44, newTreeNode(55, nil, nil), nil)))
	fmt.Println("balanced: expect false", isBalanced(root))
}

func main() {
	//TestPrintLevel()
	//TestLongestBranch()
	TestBalanced()
}
