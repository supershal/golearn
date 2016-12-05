package main

import (
	"bytes"
	"fmt"
	"strconv"
)

type cnode struct {
	value  int
	next   *cnode
	random *cnode
}

func newCNode(v int, n, r *cnode) *cnode {
	return &cnode{
		v,
		n,
		r,
	}
}

func cprint(head *cnode) string {
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

func copyUsingMap(head *cnode) *cnode {
	if head == nil {
		return nil
	}
	cache := make(map[int]*cnode, 0)
	current := head
	copy := newCNode(current.value, nil, nil)
	cache[current.value] = copy
	prev := copy
	for current != nil {
		current = current.next
		if current == nil {
			break
		}
		newN := newCNode(current.value, nil, nil)
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

func copyList(head *cnode) *cnode {
	if head == nil {
		return nil
	}

	current := head

	for current != nil {
		newn := newCNode(current.value, nil, nil)
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
	one := newCNode(1, nil, nil)
	two := newCNode(2, nil, nil)
	three := newCNode(3, nil, nil)
	four := newCNode(4, nil, nil)

	one.next, one.random = two, two
	two.next, two.random = three, four
	three.next, three.random = four, one
	four.next, four.random = nil, one

	cases := []struct {
		in, out *cnode
	}{
		{nil, nil},
		{one, one},
	}

	fmt.Println(">>>copy linked list using map<<<")
	for _, c := range cases {
		r := copyUsingMap(c.in)
		fmt.Println("case:", cprint(c.in), "expceted:", cprint(c.out), "result:", cprint(r))
	}

	fmt.Println(">>>copy linked list using pointer<<<")
	for _, c := range cases {
		fmt.Println("original in:", cprint(c.in), " original expected:", cprint(c.out))
		r := copyList(c.in)
		fmt.Println("case:", cprint(c.in), "expected:", cprint(c.out), "result:", cprint(r))
	}

}
