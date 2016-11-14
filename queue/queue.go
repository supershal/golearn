package main

import (
	"fmt"
	"log"
)

type node struct {
	value int
	next  *node
}

func newNode(val int, n *node) *node {
	return &node{val, n}
}

func print(root *node) {
	for root != nil {
		fmt.Printf("%d->", root.value)
		root = root.next
	}
	fmt.Println("")
}

type queue struct {
	first, last *node
}

func (q *queue) enqueue(val int) {
	n := newNode(val, nil)
	if q.isEmpty() {
		q.first, q.last = n, n
		return
	}
	q.last.next = n
	q.last = n
}

func (q *queue) dequeue() int {
	if q.isEmpty() {
		log.Fatalln("cant dequeue from empty list")
		return -1
	}
	n := q.first
	if q.first == q.last {
		q.first, q.last = nil, nil
	} else {
		q.first = q.first.next
	}
	return n.value
}

func (q *queue) isEmpty() bool {
	return q.first == nil && q.last == nil
}

func main() {
	q := &queue{}
	q.enqueue(1)
	q.enqueue(2)
	q.enqueue(3)
	q.dequeue()
	print(q.first)
}
