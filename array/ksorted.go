package main

import "container/heap"

// merg k sorted array.
type sortedArr []int

func ksorted(klist []sortedArr) []int {
	pq := pqueue{}
	// add first elements
	var row int
	for _, s := range klist {
		pq = append(pq, &node{s[0], row, 0})
		row++
	}
	heap.Init(pq)
	var sorted []int
	for len(pq) != 0 {
		pop := heap.Pop(pq).(node)
		sorted = append(sorted, pop.element)
		next := klist[pop.rowIndex][pop.nextIndex]
		x := &node{next, pop.rowIndex, pop.nextIndex + 1}
		heap.Push(pq, x)
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

func (p pqueue) Less(i, j int) bool {
	return p[i].element < p[j].element
}

func (p pqueue) Pop() interface{} {
	old := *p
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
