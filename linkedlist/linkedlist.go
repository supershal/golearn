package main

import (
	"bytes"
	"fmt"
	"strconv"
)

type node struct {
	value int
	next  *node
}

func newNode(val int, n *node) *node {
	return &node{val, n}
}

func reverse(head *node) *node {
	if head == nil {
		return head
	}
	if head.next == nil {
		return head
	}

	var prev, rest *node
	current := head

	for current != nil {
		rest = current.next
		current.next = prev
		prev = current
		current = rest
	}
	return prev
}

func print(root *node) string {
	var r bytes.Buffer
	for root != nil {
		r.WriteString(strconv.Itoa(root.value))
		r.WriteString("->")
		root = root.next
	}
	r.WriteString("nil\n")
	return r.String()
}

func middle(root *node) *node {
	fast, slow := root, root
	for fast != nil && fast.next != nil {
		fast = fast.next.next
		if fast == nil {
			break // important for list with length 2
		}
		slow = slow.next
	}
	return slow
}

func sumlist(num1 *node, num2 *node) *node {
	if num1 == nil {
		return num2
	}
	if num2 == nil {
		return num1
	}
	num1 = reverse(num1)
	num2 = reverse(num2)

	result := newNode(0, nil)
	retSum := result
	for num1 != nil || num2 != nil {
		var n1, n2 int
		if num1 == nil {
			n1 = 0
		} else {
			n1 = num1.value
		}
		if num2 == nil {
			n2 = 0
		} else {
			n2 = num2.value
		}
		sum := result.value + n1 + n2
		rem := sum % 10
		carry := sum / 10
		result.value = rem
		result.next = newNode(carry, nil)
		result = result.next

		if num1 != nil {
			num1 = num1.next
		}
		if num2 != nil {
			num2 = num2.next
		}
	}

	return reverse(retSum)

}

func deleteAlt(head *node) {
	if head == nil {
		return
	}
	if head.next == nil {
		return
	}
	rest := head.next.next
	deleteAlt(rest)
	head.next = rest
}

func mergeSort(head *node) *node {
	if head == nil {
		return head
	}
	if head.next == nil {
		return head
	}
	fast, slow := head.next, head
	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
	}
	rest := slow.next
	slow.next = nil
	head = mergeSort(head)
	rest = mergeSort(rest)
	return sortedMerge(head, rest)
}
func sortedMerge(head *node, rest *node) *node {
	if head == nil {
		return rest
	}
	if rest == nil {
		return head
	}

	var merged *node
	if head.value < rest.value {
		merged = head
		merged.next = sortedMerge(head.next, rest)
	} else if head.value > rest.value {
		merged = rest
		merged.next = sortedMerge(head, rest.next)
	} else {
		merged = head
		merged.next = rest
		merged = merged.next
		merged.next = sortedMerge(head.next, rest.next)
	}
	return merged
}

func pairwiseSwap(head *node) *node {
	if head == nil {
		return head
	}
	if head.next == nil {
		return head
	}
	current := head.next
	rest := current.next
	current.next = head
	head.next = pairwiseSwap(rest)
	return current
}

func isPalindrome(head *node) bool {
	if head == nil {
		return true
	}
	stack := make([]int, 0)
	curr, fast := head, head
	for fast != nil && fast.next != nil {
		stack = append(stack, curr.value)
		curr = curr.next
		fast = fast.next.next
	}
	// 1, 2, 1 -> stack (1), slow = 2, fast = 1
	// 1,2,2,1 -> stack 1,2   slow=2' fast=nil
	// so in case 1: if we want to compare 1 with 1 then when fast != nil then advance current pointer
	if fast != nil {
		curr = curr.next
	}

	for curr != nil {
		pop := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if pop != curr.value {
			return false
		}
		curr = curr.next
	}
	return true
}

func TestPalindrome() {
	cases := []struct {
		in  *node
		out bool
	}{
		{nil, true},
		{newNode(1, nil), true},
		{newNode(1, newNode(2, newNode(1, nil))), true},
		{newNode(1, newNode(2, newNode(2, newNode(1, nil)))), true},
		{newNode(1, newNode(2, newNode(2, newNode(3, nil)))), false},
		{newNode(1, newNode(2, newNode(3, newNode(3, nil)))), false},
	}

	for _, c := range cases {
		fmt.Println("palindrom of :", print(c.in), "Expected:", c.out, "Result:", isPalindrome(c.in))
	}
}

func nthNode(head *node, n int) *node {
	if head == nil {
		return nil
	}
	current := head
	for current != nil && n > 0 {
		current = current.next
		n = n - 1
	}

	rest := head
	for current != nil {
		rest = rest.next
		current = current.next
	}
	return rest
}

func recNthNode(head *node, n int) int {
	if head == nil {
		return 0
	}
	i := recNthNode(head.next, n) + 1
	if i == n {
		fmt.Printf("nth node recursive: %v\n", head)
	}
	return i
}

//first solve using Visited Map method.
// then optimize it if asked.
func detectLoop(head *node) bool {
	if head == nil {
		return false
	}
	if head.next == head {
		return true
	}
	slow, fast := head, head
	for fast != nil && fast.next != nil {
		fast = fast.next.next
		if slow == fast {
			return true
		}
		slow = slow.next

	}
	return false
}

// very difficult
// 1->2->4->5->6
//       |     |
//       9<-8<-7
// output should be 4
//https://www.youtube.com/watch?v=apIw0Opq5nk
// lets say both meet at m distance. loop size is l
// fast and slow. fast = 2*slow
// fast and slow meet at k in loop
// 2(m+s*l+k) = m+f*l+k
// m+k = (f - 2s)*l
// m = (f-2s)*l - k
// m =  (1 - 2(0)) *6 -k
// m =  6 - k
// problem reduce to find kth element from end of thelist.
// since we already have kth element, if we move n-k, we will find end of the list.
// m = 6-4 = 2
// move fast to two points to find end of the list. then break.
func detectLoopStart(head *node) *node {
	if head == nil {
		return head
	}
	if head.next == head {
		return head
	}
	slow, fast := head, head
	for fast != nil && fast.next != nil {
		fast = fast.next.next
		slow = slow.next
		if fast == nil {
			return nil // no loop
		}
		if slow == fast {
			break
		}
	}

	// Important.
	// now fast is at k distance.
	// move len - k disance is equivalent of m distance from head.
	slow = head
	for fast != slow {
		fast = fast.next
		slow = slow.next
	}

	return slow
}

func detectLoopAndBreak(head *node) *node {
	if head == nil {
		return head
	}
	if head.next == head {
		return head
	}
	slow, fast := head, head
	for fast != nil && fast.next != nil {
		fast = fast.next.next
		slow = slow.next
		if fast == nil {
			return nil // no loop
		}
		if slow == fast {
			break
		}
	}

	slow = head
	var prev *node
	for fast != slow {
		prev = fast
		fast = fast.next
		slow = slow.next
	}
	prev.next = nil
	return head
}

// insert element in already sorted list
// if the root is nil, assign head = newnode
// if root > element, head -> new, new.next = head
// continue to next
// test case: (nil, 1)=1, (1->nil, 2)=1->2, (1->2->3->5,4 )=1->2->3->4->5
func sortedInsert(head *node, n int) {
	if head == nil {
		return
	}
	if head.value > n {
		return
	}

	sortedInsert(head.next, n)
	// important condition. usecase (1->nil, 2)
	if head.value < n && (head.next == nil || n < head.next.value) {
		rest := head.next
		head.next = newNode(n, rest)
	}

}

// find intersection of two linkedl list.
// usecase: (nil, nil) = nil
//  (nil, 1) = nil
// (1->2->3->4, 10->11->2->3->4) = 2
//1) diff = len(a)-len(b), traverse first list till diff. then traverse both until find common element.
//2) create loop from one list. find loop in second. the looped element is common element.

func intersectionPoint(head1, head2 *node) *node {
	if head1 == nil || head2 == nil {
		return nil
	}
	curr1 := head1
	for curr1 != nil && curr1.next != nil {
		curr1 = curr1.next
	}
	curr1.next = head1
	loop := detectLoopStart(head2)
	return loop
}

func intersectionOfSorted(head1 *node, head2 *node) *node {
	if head1 == nil || head2 == nil {
		return nil // no intersection
	}
	var inter *node
	if head1.value == head2.value {
		inter = newNode(head1.value, nil)
		inter.next = intersectionOfSorted(head1.next, head2.next)
	} else if head1.value < head2.value {
		inter = intersectionOfSorted(head1.next, head2)
	} else {
		inter = intersectionOfSorted(head1, head2.next)
	}
	return inter
}

func TestIntersectionOfSorted() {
	fmt.Println("intersection of sorted")
	cases := []struct {
		root1, root2, result *node
	}{
		{nil, nil, nil},
		{nil, newNode(1, nil), nil},
		{newNode(1, newNode(2, newNode(4, newNode(7, nil)))), newNode(4, newNode(6, newNode(7, nil))), newNode(4, newNode(7, nil))},
	}

	for _, c := range cases {
		r := intersectionOfSorted(c.root1, c.root2)
		fmt.Println("root1=", print(c.root1), "root2=", print(c.root2), "expected result=", print(c.result), "actual result=", print(r))
	}

}

// revrese linked liset in group of k elements.
func reverseGroupofK(head *node, k int) *node {
	if head == nil || k == 0 {
		return head
	}

	current := head
	for c := 0; c < k-1; c++ {
		current = current.next
		if current == nil {
			return head
		}
	}

	rest := current.next
	current.next = nil
	rev := reverse(head)
	head.next = reverseGroupofK(rest, k)
	return rev
}

//cases: nil, 0 = nil
// 1->nil, 1 = nil
// 1->2->3, 1 = 1->2->3
// 1->2->3->4,  2 = 2->1->4->3
func TestReverseGroupofK() {
	cases := []struct {
		head   *node
		k      int
		result *node
	}{
		{nil, 0, nil},
		{newNode(1, nil), 1, newNode(1, nil)},
		{newNode(1, newNode(2, newNode(3, nil))), 1, newNode(1, newNode(2, newNode(3, nil)))},
		{newNode(1, newNode(2, newNode(3, newNode(4, nil)))), 2, newNode(2, newNode(1, newNode(4, newNode(3, nil))))},
		{newNode(1, newNode(2, newNode(3, newNode(4, nil)))), 3, newNode(3, newNode(2, newNode(1, newNode(4, nil))))},
	}

	for _, c := range cases {
		fmt.Printf("root=%v", print(c.head))
		r := reverseGroupofK(c.head, c.k)
		fmt.Println("k=", c.k, "\nexpected result=", print(c.result), "actual result=", print(r))
	}
}

// delete node from linked list with greater value on right side.
// iterate recursively through list
// when return from list, keep track of max element
// if current element is less than last max, delete it.
// if current element is greater than last max, continue
var maxRight int = 0

func deleteNodeGreaterRight(head *node) {
	if head == nil {
		return
	}
	rest := head.next

	deleteNodeGreaterRight(rest)
	maxRight = max(maxRight, head.value)

	if head.next == nil {
		return
	}
	if head.value >= maxRight {
		return
	}
	head.value = rest.value
	head.next = rest.next
	rest = nil
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func TestDeleteNodeGreaterRight() {
	cases := []struct {
		head   *node
		result *node
	}{
		{nil, nil},
		{newNode(1, nil), newNode(1, nil)},
		{newNode(1, newNode(2, newNode(3, nil))), newNode(3, nil)},
		{newNode(4, newNode(3, newNode(2, newNode(1, nil)))), newNode(4, newNode(3, newNode(2, newNode(1, nil))))},
		{newNode(1, newNode(8, newNode(5, newNode(6, nil)))), newNode(8, newNode(6, nil))},
	}

	for _, c := range cases {
		fmt.Printf("root=%v", print(c.head))
		maxRight = 0
		deleteNodeGreaterRight(c.head)
		fmt.Println("expected result=", print(c.result), "actual result=", print(c.head))
	}
}

func main() {
	//TestIntersectionOfSorted()
	//TestReverseGroupofK()
	TestDeleteNodeGreaterRight()
	TestCopy()

	// fmt.Println("detect start of loop")
	// 1->2->4->5->6
	//       |     |
	//       9<-8<-7
	// startNode := newNode(4, nil)
	// rest := newNode(5, newNode(6,newNode(7, newNode(8, newNode(9, startNode)))))
	// startNode.next = rest
	// root := newNode(1, newNode(2, startNode))

	// fmt.Println("loop start: (1->2->4->5->6->7->8->9->4", detectLoopStart(root))

	// fmt.Println("detect and break loop")
	// startNode := newNode(4, nil)
	// rest := newNode(5, newNode(6, newNode(7, newNode(8, newNode(9, startNode)))))
	// startNode.next = rest
	// root := newNode(1, newNode(2, startNode))

	// root = detectLoopAndBreak(root)
	// fmt.Printf("loop break: (1->2->4->5->6->7->8->9->4) : ")
	// print(root)

	// fmt.Println("Find intersection point")
	// var root1, root2 *node = nil, nil
	// fmt.Println("(nil, nil) =", intersectionPoint(root1, root2))

	// root1, root2 = newNode(1, nil), nil
	// fmt.Println("(1->nil, nil) =", intersectionPoint(root1, root2))

	// intersect := newNode(2, newNode(3, nil))
	// root1, root2 = newNode(1, intersect), newNode(11, newNode(22, newNode(33, intersect)))
	// fmt.Println("(1->2->3, 11->12->33->2->3) =", intersectionPoint(root1, root2))

	//one := newNode(1, nil)
	//root := newNode(2, newNode(3, one))
	//one.next = root
	//fmt.Println(detectLoop2(root))

	// fmt.Println("SortedInsert ....")
	// root = newNode(1, nil)
	// sortedInsert(root, 1)
	// fmt.Printf("sortedinsert (1->nil,1) :")
	// print(root)

	// root = nil
	// sortedInsert(root, 1)
	// fmt.Printf("sortedinsert (nil,1) :")
	// print(root)

	// root = newNode(1, nil)
	// sortedInsert(root, 2)
	// fmt.Printf("sortedinsert (1,2):")
	// print(root)

	// root = newNode(1, newNode(2, newNode(3, newNode(5, nil))))
	// sortedInsert(root, 4)
	// fmt.Printf("sortedinsert (1->2->3->5,4) :")
	// print(root)

	// fmt.Println("detect loop")
	// root := newNode(1, nil)
	// fmt.Println("loop - 1", detectLoop(root))

	// one := newNode(1, nil)
	// one.next = one
	// fmt.Println("loop - 1->1", detectLoop(one))

	// one = newNode(1, nil)
	// one.next = newNode(2, newNode(3, one))
	// fmt.Println("loop - 1->2->3->1", detectLoop(one))

	// fmt.Println("nth element")
	// root := newNode(1, nil)
	// fmt.Println("nth element:", nthNode(root, 2)) // incorrect

	// root = newNode(1, newNode(2, newNode(3, newNode(4, newNode(5, nil)))))
	// fmt.Println("nth element:", nthNode(root, 2))

	// root = newNode(1, nil)
	// recNthNode(root, 1) //correct

	// root = newNode(1, newNode(2, newNode(3, newNode(4, newNode(5, nil)))))
	// recNthNode(root, 2)

	//root := newNode(1, newNode(2, newNode(3, newNode(4, newNode(5, nil)))))
	// fmt.Println("reverse")
	// print(reverse(root))
	// root = newNode(1, newNode(2, newNode(3, newNode(4, newNode(5, newNode(6, nil))))))
	// fmt.Println("middle")
	// fmt.Println(middle(root))

	// fmt.Println("sum")
	// num1 := newNode(11, newNode(22, newNode(32, nil)))
	// num2 := newNode(12, newNode(22, nil))
	// print(sumlist(num1, num2))

	// fmt.Println("Delete Alternative")
	// root = newNode(1, newNode(2, newNode(3, newNode(4, newNode(5, newNode(6, nil))))))
	// deleteAlt(root)
	// print(root)

	// fmt.Println("mergeSort-1")
	// root = newNode(1, newNode(2, newNode(3, newNode(4, newNode(5, newNode(6, nil))))))
	// print(mergeSort(root))

	// fmt.Println("mergeSort-2")
	// root = newNode(6, newNode(5, newNode(4, newNode(3, newNode(2, newNode(1, nil))))))
	//print(mergeSort(root))

	//fmt.Println("pairwise swap")
	///root = newNode(6, newNode(5, newNode(4, newNode(3, newNode(2, newNode(1, nil))))))
	//print(pairwiseSwap(root))
}
