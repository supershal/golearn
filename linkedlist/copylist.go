package main

import (
	"bytes"
	"fmt"
	"strconv"
)

type node struct {
	value  int
	next   *node
	random *node
}

func newNode(v int, n, r *node) *node {
	return &node{
		v,
		n,
		r,
	}
}

func print(head *node) string {
	var r bytes.Buffer
	for head != nil {
		r.WriteString(strconv.Itoa(head.value))
		r.WriteString(",")
		if head.random != nil {
			r.WriteString(strconv.Itoa(head.random.value))
		}
		r.WriteString("->")
		head = head.next
	}
	r.WriteString("\n")
	return r.String()
}

func copyUsingMap(head *node) *node {
	if head == nil {
		return nil
	}
	cache := make(map[int]*node, 0)
	current := head
	copy := newNode(current.value, nil, nil)
	cache[current.value] = copy
	prev := copy
	for current != nil {
		current = current.next
		if current == nil {
			break
		}
		newN := newNode(current.value, nil, nil)
		cache[current.value] = newN
		prev.next = newN
		prev = newN
	}

	current = head
	cc := copy
	for current != nil && cc != nil {
		cc.random = cache[current.random.value]
		current = current.next
		cc = cc.next
	}

	return copy

}

func copyList(head *node) *node {
	if head == nil {
		return nil
	}

	current := head

	for current != nil {
		newn := newNode(current.value, nil, nil)
		rest := current.next
		current.next = newn
		newn.next = rest
		current = rest
	}
	current = head
	for current != nil {
		current.next.random = current.random.next
		current = current.next.next
	}

	current = head
	copy := current.next

	for current != nil {
		rest := current.next
		if rest != nil {
			current.next = rest.next
		}
		current = rest
	}

	return copy

}

func TestCopy() {
	one := newNode(1, nil, nil)
	two := newNode(2, nil, nil)
	three := newNode(3, nil, nil)
	four := newNode(4, nil, nil)

	one.next, one.random = two, two
	two.next, two.random = three, four
	three.next, three.random = four, one
	four.next, four.random = nil, one

	cases := []struct {
		in, out *node
	}{
		{nil, nil},
		{one, one},
	}

	fmt.Println(">>>copy linked list using map<<<")
	for _, c := range cases {
		r := copyUsingMap(c.in)
		fmt.Println("case:", print(c.in), "expceted:", print(c.out), "result:", print(r))
	}

	fmt.Println(">>>copy linked list using pointer<<<")
	for _, c := range cases {
		fmt.Println("original in:", print(c.in), " original expected:", print(c.out))
		r := copyList(c.in)
		fmt.Println("case:", print(c.in), "expected:", print(c.out), "result:", print(r))
	}

}

func main() {
	TestCopy()
}
