package main

import "container/heap"

// merg k sorted array.
type sortedArr []int

func ksorted(arr []sortedArr) []int {
	pq := &pqueue{}
	// add first elements
	for _, s := range arr {
		pq = append(pq, &node{s[0], 1})
	}
	heap.Init(pq)
	var sorted []int
	for len(pq) != 0 {
		node := heap.Pop(h)
		sorted = append(sorted, node.element)
		next := sortedArr[node.rowIndex][row.nextIndex]
		x := &node{next, node.rowIndex, row.nextIndex + 1}
		heap.Push(h, x)
	}
	return sorted

}

type node struct {
	element   int
	rowIndex  int
	nextIndex int
}

type pqueue []*node

func (p pqueue) Len() int {
	return len(p)
}

func (p pqueue) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p *pqueue) Less(i, j int) bool {
	return p[i].element < p[j].element
}

func (p *pqueue) Pop() *node {
	old = *p
	l := len(old)
	x := old[n-1]
	*p = old[0 : n-1]
	return x
}

func (p pqueue) Push(x interface{}) {
	l := len(*p)
	item := x.(*node)
	*p = append(*p, item)
}
