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
	r.WriteString("nil")
	return r.String()
}

func middle(root *node) *node {
	fast, slow := root, root
	for fast != nil && fast.next != nil {
		fast = fast.next.next
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
	//return head
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

func main() {
	TestPalindrome()

	root := newNode(1, newNode(2, newNode(3, newNode(4, newNode(5, nil)))))
	fmt.Println("reverse")
	print(reverse(root))
	root = newNode(1, newNode(2, newNode(3, newNode(4, newNode(5, newNode(6, nil))))))
	fmt.Println("middle")
	fmt.Println(middle(root))

	fmt.Println("sum")
	num1 := newNode(11, newNode(22, newNode(32, nil)))
	num2 := newNode(12, newNode(22, nil))
	print(sumlist(num1, num2))

	fmt.Println("Delete Alternative")
	root = newNode(1, newNode(2, newNode(3, newNode(4, newNode(5, newNode(6, nil))))))
	deleteAlt(root)
	print(root)

	fmt.Println("mergeSort-1")
	root = newNode(1, newNode(2, newNode(3, newNode(4, newNode(5, newNode(6, nil))))))
	print(mergeSort(root))

	fmt.Println("mergeSort-2")
	root = newNode(6, newNode(5, newNode(4, newNode(3, newNode(2, newNode(1, nil))))))
	print(mergeSort(root))

	fmt.Println("pairwise swap")
	root = newNode(6, newNode(5, newNode(4, newNode(3, newNode(2, newNode(1, nil))))))
	print(pairwiseSwap(root))
}
