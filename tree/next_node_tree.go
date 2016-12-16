package main

type nextTree struct {
	value             int
	left, right, next *nextTree
}

func newNextTree(val int, left, right, next *nextTree) *nextTree {
	return &nextTree{
		val,
		left,
		right,
		next,
	}
}

// connect every node to their next node in same level.

//     4 -> nil
//    /   \
//   2  -> 7 -> nil
//  / \   / \
// 1->3->6-> 9->nil

// method 1.
// level order traversal. keep pointer to prev node.
// keep level when adding the node
// if level is same as prev
//  connect prev. next to current
// move prev to current
// if not  then
//   just change prev to current
func nextNodeLevelOrder(root *nextTree) {
	if root == nil {
		return
	}
	var prev *nextTree = nil
	prevLevel := 0
	queue := []*nextTree{root}
	level := []int{1}

	for len(queue) != 0 {
		current := queue[0]
		queue = queue[1:]
		currLevel := level[0]
		level = level[1:]

		if currLevel == prevLevel {
			prev.next = current
		}
		prev = current
	}
}

//     4 -> nil
//    /   \
//   2  -> 7 -> nil
//  / \   / \
// 1->3->6-> 9->nil
// method2: outer loop: start from root and go left.
//  inner loop: go to next.

func nextNode(root *nextTree) {
	if root == nil {
		return
	}
	curr := root
	for curr != nil {
		currLevel := curr
		for currLevel != nil {
			if currLevel.left != nil && currLevel.right != nil {
				currLevel.left.next = currLevel.right
				if currLevel.next != nil {
					currLevel.right.next = currLevel.next.left
				}
				currLevel = currLevel.next
			}
		}
		curr = curr.left
	}
}
