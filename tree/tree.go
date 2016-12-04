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

func maxDepth(root *node) int {
	if root == nil {
		return 0
	}
	lsum, rsum := 0, 0
	if root.left != nil {
		lsum += (1 + maxDepth(root.left))
	}
	if root.right != nil {
		rsum += (1 + maxDepth(root.right))
	}
	return max(lsum, rsum)
}

func maxDepth2(root *node) int {
	if root == nil {
		return 0
	}
	return 1 + max(maxDepth2(root.left), maxDepth2(root.right))
}
func TestMaxDepth() {
	ele := newTreeNode(11, nil, nil)
	twe := newTreeNode(12, ele, nil)
	nine := newTreeNode(9, nil, nil)
	ten := newTreeNode(10, nine, twe)
	fmt.Println("Depth of ", printLevel(ten), "expected: 2", "got:", maxDepth(ten))
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
func min(a, b int) int {
	if a < b {
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
	if balancedHeight(root) > -1 {
		return true
	}
	return false
}

func balancedHeight(root *node) int {
	if root == nil {
		return 0
	}
	l := balancedHeight(root.left)
	r := balancedHeight(root.right)

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

// Invert tree.
// below problem can also be done using BSF.
// when pop the node, change the pointers and add those nodes to the list.
func invertTree(root *node) {
	if root == nil {
		return
	}
	invertTree(root.left)
	invertTree(root.right)
	root.left, root.right = root.right, root.left
}

//     4
//    /   \
//   2     7
//  / \   / \
// 1   3 6   9
func TestInvertTree() {
	tree := []int{4, 2, 7, 1, 3, 6, 9} // level order
	var root *node
	for i := 0; i < len(tree); i++ {
		root = insert(tree[i], root)
	}

	fmt.Println("Inver tree:", printLevel(root))
	invertTree(root)
	fmt.Println("Got:", printLevel(root))
}

func sameTree(a, b *node) bool {
	if a == nil || b == nil {
		return a == b
	}
	return (a.value == b.value) && sameTree(a.left, b.left) && sameTree(a.right, b.right)
}

// LCA of BST
//    _______6______
//     /              \
//  ___2__          ___8__
// /      \        /      \
// 0      _4       7       9
//       /  \
//       3   5

func lcaBst(root *node, val1, val2 int) *node {
	if root == nil {
		return nil
	}
	if root.value < val1 && root.value < val2 {
		return lcaBst(root.right, val1, val2)
	}

	if root.value > val1 && root.value > val2 {
		return lcaBst(root.left, val1, val2)
	}
	return root
}

func TestLcaBst() {
	tree := []int{4, 2, 7, 1, 3, 6, 9} // level order
	var root *node
	for i := 0; i < len(tree); i++ {
		root = insert(tree[i], root)
	}
	cases := []struct {
		in   *node
		val1 int
		val2 int
		out  int
	}{
		{root, 7, 1, 4},
		{root, 1, 3, 2},
		{root, 1, 2, 2},
	}
	fmt.Println("LCABst of ", printLevel(root))
	for _, c := range cases {
		result := lcaBst(c.in, c.val1, c.val2)
		fmt.Println("val1=", c.val1, "val2=", c.val2, "expected=", c.out, "lca=", result.value)
	}
}

func closestBST(root *node, val int) *node {
	if root == nil {
		return nil
	}
	if root.value == val {
		return root
	}

	if root.value > val {
		return smallest(root, closestBST(root.left, val), val)
	}

	return smallest(root, closestBST(root.right, val), val)

}

func smallest(node1, node2 *node, val int) *node {
	if node1 == nil {
		return node2
	}
	if node2 == nil {
		return node1
	}
	node1Diff := abs(node1.value - val)
	node2Diff := abs(node2.value - val)

	if node1Diff < node2Diff {
		return node1
	}
	return node2
}
func abs(a int) int {
	if a < 0 {
		return a * -1
	}
	return a
}

func TestClosestBST() {
	tree := []int{4, 2, 7, 1, 3, 6, 9} // level order
	var root *node
	for i := 0; i < len(tree); i++ {
		root = insert(tree[i], root)
	}
	cases := []struct {
		in   *node
		val1 int
		out  int
	}{
		{root, 5, 6},
		{root, 7, 7},
		{root, 100, 9},
		{root, -10, 1},
	}
	fmt.Println("Closest of ", printLevel(root))
	for _, c := range cases {
		result := closestBST(root, c.val1)
		fmt.Println("val1=", c.val1, "expected=", c.out, "closest=", result.value)
	}
}
func main() {
	//TestPrintLevel()
	//TestLongestBranch()
	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
	TestBalanced()
	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
	TestMaxDepth()
	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
	TestInvertTree()
	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
	TestLcaBst()
	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
	TestClosestBST()
}
