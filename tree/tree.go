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

func makeTestTree() *node {
	tree := []int{4, 2, 7, 1, 3, 6, 9} // level order
	var root *node
	for i := 0; i < len(tree); i++ {
		root = insert(tree[i], root)
	}
	return root
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
	fmt.Println(printLevel(root))
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
	fmt.Println("Invert tree:", printLevel(root))
	invertTree(root)
	fmt.Println("Got:", printLevel(root))
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

//Given a binary tree, return all root-to-leaf paths.
func rootToLeafPath(root *node, currPath string, paths *[]string) {
	if root == nil {
		return
	}
	if root.left == nil && root.right == nil {
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
	return isSym(root.left, root.right)
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

// LCA of binary tree
// if root == nil then return nil
// if root is one of the value return value
// search value in left. if found left will be non nil
//search vaue in right. if found, right will be non nil.
// if both non nil that means root is lca
// if one of them is nil then thats the lca, given other node was under that node and exists in the tree.

func lcaBinary(root *node, p, q int) *node {
	if root == nil {
		return nil
	}
	if root.value == p || root.value == q {
		return root
	}

	left := lcaBinary(root.left, p, q)
	right := lcaBinary(root.right, p, q)

	if left != nil && right != nil {
		return root
	}

	if left == nil {
		return right
	}
	return left
}

func TestLcaBinary() {
	root := makeTestTree()

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
	fmt.Println("LCABinary of ", printLevel(root))
	for _, c := range cases {
		result := lcaBinary(c.in, c.val1, c.val2)
		fmt.Println("val1=", c.val1, "val2=", c.val2, "expected=", c.out, "lca=", result.value)
	}
}

//Given a binary tree, return the zigzag level order traversal of its nodes' values. (ie, from left to right, then right to left for the next level and alternate between).

func zigZag(root *node) {
	if root == nil {
		return
	}

	stack1, stack2 := []*node{root}, []*node{}

	for len(stack1) != 0 || len(stack2) != 0 {
		for len(stack1) != 0 {
			top := stack1[len(stack1)-1]
			fmt.Printf("%v->", top.value)
			stack1 = stack1[:len(stack1)-1]
			if top.right != nil {
				stack2 = append(stack2, top.right)
			}
			if top.left != nil {
				stack2 = append(stack2, top.left)
			}
		}

		for len(stack2) != 0 {
			top := stack2[len(stack2)-1]
			fmt.Printf("%v->", top.value)
			stack2 = stack2[:len(stack2)-1]
			if top.left != nil {
				stack1 = append(stack1, top.left)
			}
			if top.right != nil {
				stack1 = append(stack1, top.right)
			}
		}
	}
}

func zigZag1(root *node) {
	h := height(root)
	var direction bool
	for i := 1; i <= h; i++ {
		printAtLevel(root, i, direction)
		direction = !direction
	}
}

func printAtLevel(root *node, level int, direction bool) {
	if root == nil {
		return
	}

	if level == 1 {
		fmt.Printf("%v->", root.value)
	} else {
		if direction {
			printAtLevel(root.left, level-1, direction)
			printAtLevel(root.right, level-1, direction)
		} else {
			printAtLevel(root.right, level-1, direction)
			printAtLevel(root.left, level-1, direction)
		}
	}
}

func height(root *node) int {
	if root == nil {
		return 0
	}
	return 1 + max(height(root.left), height(root.left))
}

func TestZigzag() {
	root := makeTestTree()
	fmt.Println("zigzag of ", printLevel(root))
	zigZag(root)
	fmt.Println()

	fmt.Println("zigzag2(using height)")
	zigZag1(root)
	fmt.Println()
}

// keep k outside
var kth int

func kthSmallestBST(root *node) *node {
	if root == nil || kth < 0 {
		return nil
	}

	left := kthSmallestBST(root.left)
	kth--
	if kth == 0 {
		return root
	}
	right := kthSmallestBST(root.right)
	if left == nil && right == nil {
		return nil
	}
	if left != nil {
		return left
	}
	return right
}

func TestKthSmallestBst() {
	root := makeTestTree()
	cases := []struct {
		in  *node
		k   int
		out int
	}{
		{root, 3, 3},
		{root, 5, 6},
	}

	fmt.Println("kth element of : ", printLevel(root))
	for _, c := range cases {
		kth = c.k
		fmt.Println("k=", c.k, "expected:", c.out, "got:", kthSmallestBST(c.in).value)
	}

}

// count nodes in complete binary tree
// class Solution {
//     int height(TreeNode root) {
//         return root == null ? -1 : 1 + height(root.left);
//     }
//     public int countNodes(TreeNode root) {
//         int h = height(root);
//         return h < 0 ? 0 :
//                height(root.right) == h-1 ? (1 << h) + countNodes(root.right)
//                                          : (1 << h-1) + countNodes(root.left);
//     }
// }

// Given a binary tree, count the number of uni-value subtrees.
// A Uni-value subtree means all nodes of the subtree have the same value.
// For example:
// Given binary tree,
//               5
//              / \
//             1   5
//            / \   \
//           5   5   5

// may be incorrect solution
var uni = 0

func uniValueSubtree(root *node) *node {
	if root == nil {
		return root
	}

	left := uniValueSubtree(root.left)
	right := uniValueSubtree(root.right)

	if left == nil && right == nil {
		uni++
		return root
	}
	if left == nil {
		if root.value == right.value {
			uni++
			return root
		}
		return nil
	}
	if right == nil {
		if root.value == left.value {
			uni++
			return root
		}
		return nil
	}

	if root.value == left.value && root.value == right.value {
		uni++
		return root
	}
	return nil
}

// right view and left view are the same way. when traverse left->right or right->left.
func rightView(root *node) []int {
	view := []int{}
	if root == nil {
		return view
	}

	queue, level := []*node{root}, []int{1}
	prevLvl := 0

	for len(queue) != 0 {
		top := queue[0]
		queue = queue[1:]
		currLvl := level[0]
		level = level[1:]
		if currLvl != prevLvl {
			view = append(view, top.value)
			prevLvl = currLvl
		}
		if top.right != nil {
			queue = append(queue, top.right)
			level = append(level, currLvl+1)
		}
		if top.left != nil {
			queue = append(queue, top.left)
			level = append(level, currLvl+1)
		}
	}
	return view
}

func rightView1(root *node) []int {
	view := []int{}
	rightView1Helper(root, &view, 0)
	return view
}
func rightView1Helper(root *node, result *[]int, currLevel int) {
	if root == nil {
		return
	}
	if len(*result) == currLevel {
		*result = append(*result, root.value)
	}
	rightView1Helper(root.right, result, currLevel+1)
	rightView1Helper(root.left, result, currLevel+1)
}

func TestRightView() {
	root := makeTestTree()
	fmt.Println("rightside view of ", printLevel(root), "got:", rightView(root))
	fmt.Println("rightside view 2:", rightView1(root))
}

// Implement an iterator over a binary search tree (BST). Your iterator will be initialized with the root node of a BST.
// Calling next() will return the next smallest number in the BST.
// Note: next() and hasNext() should run in average O(1) time and uses O(h) memory, where h is the height of the tree.

type treeIterator struct {
	stack []*node
}

func newTreeIterator(root *node) *treeIterator {
	if root == nil {
		return &treeIterator{}
	}
	itr := &treeIterator{}
	itr.stack = make([]*node, 0)
	for root != nil {
		itr.stack = append(itr.stack, root)
		root = root.left
	}

	return itr
}

func (itr *treeIterator) hasNext() bool {
	return len(itr.stack) != 0
}

func (itr *treeIterator) next() *node {
	if !itr.hasNext() {
		return nil
	}
	pop := itr.stack[len(itr.stack)-1]
	itr.stack = itr.stack[:len(itr.stack)-1]
	if pop.right != nil {
		succ := pop.right
		for succ != nil {
			itr.stack = append(itr.stack, succ)
			succ = succ.left
		}
	}
	return pop
}

func TestTreeIterator() {
	root := makeTestTree()
	fmt.Println("Tree iterator of BST:", printLevel(root))
	itr := newTreeIterator(root)
	for itr.hasNext() {
		fmt.Println(itr.next().value, "->")
	}
}

// Given a binary tree where all the right nodes are either leaf nodes with a sibling (a left node that shares the same parent node) or empty, flip it upside down and turn it into a tree where the original right nodes turned into left leaf nodes. Return the new root.

// For example:
// Given a binary tree {1,2,3,4,5},

//      1
//    / \
//   2   3
//  / \
// 4   5
// return the root of the binary tree [4,5,2,#,#,3,1].

//     4
//   / \
//  5   2
//     / \
//    3   1
//http://qa.geeksforgeeks.org/5011/turn-the-binary-tree-upside-down
// func flipTree(root *node) *node{
// 	// TODO
// }
var deepest = 0
var deepesNode *node

func deepestLeftLeafNode(root *node, left bool, level int) {
	if root == nil {
		return
	}
	if left && root.left == nil && root.right == nil {
		if level > deepest {
			deepesNode = root
			deepest = level
		}
	}
	deepestLeftLeafNode(root.left, true, level+1)
	deepestLeftLeafNode(root.right, false, level+1)
}

func TestDeepestLeftLeafNode() {
	root := makeTestTree()
	deepestLeftLeafNode(root, false, 0)
	fmt.Println("deepest left leaf node of :", printLevel(root), "\ngot:", deepesNode.value)
}

// Given a binary tree containing digits from 0-9 only, each root-to-leaf path could represent a number.

// An example is the root-to-leaf path 1->2->3 which represents the number 123.

// Find the total sum of all root-to-leaf numbers.

// For example,

//     1
//    / \
//   2   3
// The root-to-leaf path 1->2 represents the number 12.
// The root-to-leaf path 1->3 represents the number 13.

// Return the sum = 12 + 13 = 25.
func sumOfRoottoLeafPaths(root *node, sum int) int {
	if root == nil {
		return 0
	}
	if root.left == nil && root.right == nil {
		return sum*10 + root.value
	}
	return sumOfRoottoLeafPaths(root.left, (sum*10)+root.value) + sumOfRoottoLeafPaths(root.right, (sum*10)+root.value)

}
func TestSumOfRoottoLeafPaths() {
	root := makeTestTree()
	fmt.Println("sum of root to leaf paths of ", printLevel(root), "\n expected: 1799, got:", sumOfRoottoLeafPaths(root, 0))
}

// delete node in bst.
// in the process we have to set its parent's pointer to nil. that means we have to return referece to the iterate node and assign as left or right.
// so function should return value.
// if leaf node, return nil.
// if left is nil then right is the answer
// if right is nil then left is answer
// if both are nil then find in order succ and copy its value to root.
// then delete the succ
func deleteNodeBST(root *node, key int) *node {
	if root == nil {
		return nil
	}
	if root.value < key {
		root.right = deleteNodeBST(root.right, key)
	} else if root.value > key {
		root.left = deleteNodeBST(root.left, key)
	} else if root.value == key {
		if root.left == nil {
			return root.right
		}
		if root.right == nil {
			return root.left
		}
		succ := root.right
		for succ.left != nil {
			succ = succ.left
		}
		root.value = succ.value
		root.right = deleteNodeBST(root.right, root.value) // delete successor node.
	}
	return root
}

func TestDeleteBST() {
	root := makeTestTree()
	fmt.Println("Delete node from bst:", printLevel(root), "Delete 7")
	deleteNodeBST(root, 7)
	fmt.Println("result:\n", printLevel(root))
	fmt.Println("delete 6")
	deleteNodeBST(root, 6)
	fmt.Println("result:\n", printLevel(root))

}

//   5
//  / \
//  3  7
// / \ /
// 2 4 6

// 3 and 4
// 5->3->4
// 5->3

//4 and 6
// 5->3->4
// 5->7->6

// 7 and 6
// 5->7->6
// 5->7

// 2 and 4
// 5->3->2
// 5->3->4

// find path to the node
// find lca by comparing element by element
// skip both list until the lca is found
// then join both lists.

func pathBetweenNodesBST(root *node, x, y int) []int {
	path1, path2 := []int{}, []int{}
	nodePathBST(root, x, &path1)
	nodePathBST(root, y, &path2)
	lca := intLca(path1, path2)
	// skip until lca found
	var i, j int
	for i = 0; path1[i] != lca; i++ {
	}
	for j = 0; path2[j] != lca; j++ {
	}
	path1 = path1[i:]
	path2 = path2[j:]

	reverse(&path1)
	for i := 1; i < len(path2); i++ {
		path1 = append(path1, path2[i])
	}
	return path1
}

func reverse(path *[]int) {
	for start, end := 0, len(*path)-1; start < end; start, end = start+1, end-1 {
		(*path)[start], (*path)[end] = (*path)[end], (*path)[start]
	}

}
func intLca(p, q []int) int {
	lca := -1
	for i, j := 0, 0; i < len(p) && j < len(q); i, j = i+1, j+1 {
		if p[i] != q[j] {
			break
		}
		lca = p[i]
	}
	return lca
}

func nodePathBST(root *node, x int, path *[]int) {
	if root == nil {
		return
	}
	*path = append(*path, root.value)
	if root.value == x {
		return
	}
	if root.value > x {
		nodePathBST(root.left, x, path)
	} else {
		nodePathBST(root.right, x, path)
	}
}

//     4
//    /   \
//   2     7
//  / \   / \
// 1   3 6   9

func TestPathBetweenNodesBST() {
	root := makeTestTree()
	cases := []struct {
		in   *node
		x, y int
		out  []int
	}{
		{root, 1, 3, []int{1, 2, 3}},
		{root, 2, 3, []int{2, 3}},
		{root, 1, 9, []int{1, 2, 4, 7, 9}},
		{root, 7, 6, []int{7, 6}},
	}
	fmt.Println("path between two nodes BST:", printLevel(root))
	for _, c := range cases {
		fmt.Println("x:", c.x, "y:", c.y, "expected:", c.out, "got:", pathBetweenNodesBST(c.in, c.x, c.y))
	}

}

//find distance for each level using tree traversal. O(n)
// since root to lca path will be accounted twice in that distance
// remove 2*lcadistance from total distance.
func distanceTwoNodeBinary(root *node, x, y int) int {
	xDist := distance(root, 0, x)
	yDist := distance(root, 0, y)

	lca := lcaBinary(root, x, y)
	lcaDist := distance(root, 0, lca.value)

	return (xDist + yDist) - (2 * lcaDist)
}

func distance(root *node, level, x int) int {
	if root == nil {
		return -1
	}

	if root.value == x {
		return level
	}
	if leftDist := distance(root.left, level+1, x); leftDist != -1 {
		return leftDist
	}
	return distance(root.right, level+1, x)
}

//     4
//    /   \
//   2     7
//  / \   / \
// 1   3 6   9

func TestDistanceTwoNodeBinary() {
	root := makeTestTree()
	cases := []struct {
		in        *node
		x, y, out int
	}{
		{root, 1, 3, 2},
		{root, 2, 3, 1},
		{root, 1, 9, 4},
		{root, 7, 6, 1},
	}
	fmt.Println("Distance between two nodes in binary tree:", printLevel(root))

	for _, c := range cases {
		fmt.Println("x:", c.x, "y:", c.y, "expected:", c.out, "got:", distanceTwoNodeBinary(c.in, c.x, c.y))
	}

}

// find maximum root to leaf path sum
func maxRootToLeafPathSum(root *node) int {
	if root == nil {
		return 0
	}
	maxSum := 0
	maxRToLSum(root, 0, &maxSum)
	return maxSum
}

func maxRToLSum(root *node, curr int, maxSum *int) {
	if root.left == nil && root.right == nil {
		*maxSum = max(*maxSum, curr+root.value)
		return
	}

	curr += root.value
	maxRToLSum(root.left, curr, maxSum)
	maxRToLSum(root.right, curr, maxSum)

}

//     4
//    /   \
//   2     7
//  / \   / \
// 1   3 6   9

func TestMaxRootToLeafPathSum() {
	root := makeTestTree()
	fmt.Println("Max sum root to leaf:", printLevel(root), "expected:20 got:", maxRootToLeafPathSum(root))
}

// maximum sum between two leaf nodes in binary tree
// find lsum, rsum, add root.value. compare it with max, if greater then update max
// return maxof lsum or rsum + root.value for uppper level node.
func maxSumAnyTwoLeafsBinary(root *node) int {
	if root == nil {
		return 0
	}
	maxSum := 0
	maxLeafSum(root, &maxSum)
	return maxSum
}

func maxLeafSum(root *node, maxSum *int) int {
	if root == nil {
		return 0
	}
	leftSum := maxLeafSum(root.left, maxSum)
	rightSum := maxLeafSum(root.right, maxSum)

	if leftSum == 0 {
		return root.value + rightSum
	}

	if rightSum == 0 {
		return root.value + leftSum
	}

	// check if current path going through root has max sum.
	*maxSum = max(*maxSum, root.value+leftSum+rightSum)
	// return max leaf path sum from root to either left or right
	return max(leftSum, rightSum) + root.value
}

//     4
//    /   \
//   2     7
//  / \   / \
// 1   3 6   9

//it should be  3 + 2 + 4+ 7 +9=25

func TestMaxSumAnyTwoLeafsBinary() {
	root := makeTestTree()
	fmt.Println("Max sum two leaves binary:", printLevel(root), "expected:25 got:", maxSumAnyTwoLeafsBinary(root))
}

// if right child is not null then its loweset value in right subtree
// if right child is null then it would be one of its parent
// inorder traversal visits left node before root node then right node. root node being parent.
// so in inorder travrsal if we go left then mark current node as root and proceed until node is found.
func inOrderSucc(root, x, succ *node) *node {
	if root == nil {
		return succ
	}
	// lowest node in right subtree
	if x.right != nil {
		succ := x.right
		for succ.left != nil {
			succ = succ.left
		}
		return succ
	}
	// right subtree is nil so one of its parent is succ.
	// if the node is in right subtree then
	if root.value > x.value {
		succ = inOrderSucc(root.left, x, root)
	} else if root.value < x.value {
		succ = inOrderSucc(root.right, x, succ)
	}
	// node found. send last assigned succ
	return succ
}

func inOrderSuccItr(root, x *node) *node {
	if root == nil {
		return nil
	}
	// lowest node in right subtree
	if x.right != nil {
		succ := x.right
		for succ.left != nil {
			succ = succ.left
		}
		return succ
	}
	var succ *node
	for root != nil {
		if root.value > x.value {
			succ = root
			root = root.left
		} else if root.value < x.value {
			root = root.right
		} else {
			break
		}
	}
	return succ
}

//     4
//    /   \
//   2     7
//  / \   / \
// 1   3 6   9
func TestInorderSuccRecurr() {
	root := makeTestTree()
	cases := []struct {
		in  *node
		x   *node
		out interface{}
	}{
		{root, root, 6},
		{root, root.left.right, 4},
		{root, root.right.right, nil},
		{root, root.left, 3},
		{root, root.right, 9},
	}
	fmt.Println("Inorder successor of ", printLevel(root))
	for _, c := range cases {
		fmt.Println("(Recursive) Successor of ", c.x, "expected:", c.out, "got:", inOrderSucc(c.in, c.x, nil))
	}

	for _, c := range cases {
		fmt.Println("(Iterative) Successor of ", c.x, "expected:", c.out, "got:", inOrderSuccItr(c.in, c.x))
	}
}

type pnode struct {
	value               int
	left, right, parent *pnode
}

//inorder successor with parent pointer.
// if right node is not null then lowest node in right subtree.
// Inorder traversal does left,root, right.
// that means when a node is visited, is in left subtree of root then its succ is its parent.
// if right subtree then its parent will be last unvisited parent.
// so last unvisited parent would be the node whose left child was last visited in in order traversal.
// that means travel upword until you find a a node whose left child is current node
func inOrdersuccParent(root, x *pnode) *pnode {
	if root == nil {
		return nil
	}
	// lowest node in right subtree
	if x.right != nil {
		succ := x.right
		for succ.left != nil {
			succ = succ.left
		}
		return succ
	}
	parent := x.parent
	for parent != nil && parent.left != x {
		x = parent
		parent = x.parent
	}
	return parent
}

// validate if a tree is BST or not.
func validateBST(root *node) bool {
	if root == nil {
		return true
	}
	return validateBSTUtil(root, -1<<31, 1>>31-1)
}

func validateBSTUtil(root *node, min, max int) bool {
	if root == nil {
		return true
	}
	// fucked up here. I had && here. root should be either more than min or less then max
	if root.value < min || root.value > max {
		return false
	}
	left := validateBSTUtil(root.left, min, root.value-1)
	right := validateBSTUtil(root.right, root.value+1, max)
	return left && right
}

var prev *node

func validateBSTInOrder(root *node) bool {
	if root == nil {
		return true
	}
	if !validateBSTInOrder(root.left) {
		return false
	}
	if prev != nil && prev.value > root.value {
		return false
	}
	prev = root
	return validateBSTInOrder(root.right)
}

//     4
//    /   \
//   2     7
//  / \   / \
// 1   3 6   9

//     4
//    /   \
//   2     5
//  / \   / \
// 1   3 6   9

func TestValidateBST() {
	root := makeTestTree()
	fmt.Println("validate BST of: ", printLevel(root), "expected: true got:", validateBST(root))
	root.right.value = 5
	fmt.Println("validate BST of: ", printLevel(root), "expected: false got:", validateBST(root))

	root = makeTestTree()
	fmt.Println("validate BST (inorder) of: ", printLevel(root), "expected: true got:", validateBSTInOrder(root))
	root.right.value = 5
	fmt.Println("validate BST(inorder) of: ", printLevel(root), "expected: false got:", validateBSTInOrder(root))
}

func main() {
	TestLevelOrder()
	//TestPrintLevel()
	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
	TestLongestBranch()
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
	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
	TestSumOfLeftLeaves()
	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
	TestRootToLeafPath()
	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
	TestRootToLeafSum()
	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
	TestBottomUpLeftRightLevel()
	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
	TestLcaBinary()
	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
	TestZigzag()
	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
	TestKthSmallestBst()
	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
	TestRightView()
	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
	TestTreeIterator()
	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
	TestDeepestLeftLeafNode()
	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
	TestSumOfRoottoLeafPaths()
	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
	TestDeleteBST()
	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
	TestPathBetweenNodesBST()
	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
	TestDistanceTwoNodeBinary()
	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
	TestMaxRootToLeafPathSum()
	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
	TestMaxSumAnyTwoLeafsBinary()
	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
	TestInorderSuccRecurr()
	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
	TestValidateBST()
}
