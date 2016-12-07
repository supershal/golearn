package main

import (
	"bytes"
	"fmt"
	"math"
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

func sumOfLeftLeaves(root *node, left bool) int {
	if root == nil {
		return 0
	}
	if root.left == nil && root.right == nil && left {
		return root.value
	}
	var sumLeft = 0
	sumLeft += sumOfLeftLeaves(root.left, true) + sumOfLeftLeaves(root.right, false)
	return sumLeft
}

func TestSumOfLeftLeaves() {
	tree := []int{4, 2, 6, 1, 3, 5, 7} // level order
	var root *node
	for i := 0; i < len(tree); i++ {
		root = insert(tree[i], root)
	}
	fmt.Println("SumOfLeftLeaves: expected: 6, got:", sumOfLeftLeaves(root, false))
}

//Given a binary tree, return all root-to-leaf paths.
func rootToLeafPath(root *node, currPath string, paths *[]string) {
	if root == nil {
		return
	}
	if root.left == nil || root.right == nil {
		*paths = append(*paths, currPath+strconv.Itoa(root.value))
	}
	currPath += strconv.Itoa(root.value) + "->"
	rootToLeafPath(root.left, currPath, paths)
	rootToLeafPath(root.right, currPath, paths)
}

func TestRootToLeafPath() {
	tree := []int{4, 2, 6, 1, 3, 5, 7} // level order
	var root *node
	for i := 0; i < len(tree); i++ {
		root = insert(tree[i], root)
	}
	paths := make([]string, 0)
	rootToLeafPath(root, "", &paths)
	fmt.Println("rootToLeafPath of:", printLevel(root), "paths:", paths)

}

//Given a binary tree and a sum, determine if the tree has a root-to-leaf path such that adding up all the values along the path equals the given sum.
func rootToLeafSum(root *node, sum int) bool {
	if root.left == nil && root.right == nil {
		if (root.value - sum) == 0 {
			return true
		} else {
			return false
		}
	}
	return rootToLeafSum(root.left, sum-root.value) || rootToLeafSum(root.right, sum-root.value)
}

func TestRootToLeafSum() {
	tree := []int{4, 2, 6, 1, 3, 5, 7} // level order
	var root *node
	for i := 0; i < len(tree); i++ {
		root = insert(tree[i], root)
	}
	fmt.Println("rootToLeafPath of:", printLevel(root), ", sum = 15, expected: true, got:", rootToLeafSum(root, 15))
	fmt.Println("rootToLeafPath of:", printLevel(root), ", sum = 25, expected: false, got:", rootToLeafSum(root, 25))

}

//Given a binary tree, find its minimum depth.
//The minimum depth is the number of nodes along the shortest path from the root node down to the nearest leaf node.
func minDepth(root *node, height int) int {
	if root == nil {
		return -1
	}
	if root.left == nil && root.right == nil {
		return height
	}

	return min(minDepth(root.left, height+1), minDepth(root.right, height+1))
}

func minDepth2(root *node) int {
	if root == nil {
		return 0
	}
	l := minDepth2(root.left)
	r := minDepth2(root.right)
	if l == 0 || r == 0 { // tree with only one child.
		return l + r + 1
	}
	return min(l, r) + 1
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Bottom up level order traversal.
// try to solve with three elements. then generalize it.
func bottomUpLeftRightLevel(root *node) (stack []*node) {
	if root == nil {
		return stack
	}
	queue := make([]*node, 0)
	queue = append(queue, root)
	for len(queue) != 0 {
		top := queue[0]
		queue = queue[1:]
		stack = append(stack, top)
		if top.right != nil {
			queue = append(queue, top.right)
		}
		if top.left != nil {
			queue = append(queue, top.left)
		}
	}
	return stack
}

func TestBottomUpLeftRightLevel() {
	tree := []int{4, 2, 6, 1, 3, 5, 7} // level order
	var root *node
	for i := 0; i < len(tree); i++ {
		root = insert(tree[i], root)
	}
	stack := bottomUpLeftRightLevel(root)
	fmt.Println("Bottom up level order traversal of ", printLevel(root))
	for i := len(stack) - 1; i >= 0; i-- {
		fmt.Printf("%v->", stack[i].value)
	}
	fmt.Println("")
}

// Given a binary tree, check whether it is a mirror of itself (ie, symmetric around its center).

// For example, this binary tree [1,2,2,3,4,4,3] is symmetric:

//     1
//    / \
//   2   2
//  / \ / \
// 3  4 4  3

func isSymetric(root *node) bool {
	if root == nil {
		return true
	}
	isSym(root.left, root.right)
}
func isSym(left, right *node) bool {
	if left == nil || right == nil {
		return left == right
	}
	if left.value != right.value {
		return false
	}
	return isSym(left.left, right.right) && isSym(left.right, right.left)
}

func main() {
	TestLevelOrder()
	//TestPrintLevel()
	//TestLongestBranch()
	TestBalanced()
	TestSumOfLeftLeaves()
	TestRootToLeafPath()
	TestRootToLeafSum()
	TestBottomUpLeftRightLevel()
}
