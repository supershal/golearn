package main

import "fmt"

type queue struct {
	pushStack, popStack *stack
}

func NewQueue(len int) *queue {
	return &queue{
		NewStack(len),
		NewStack(len),
	}
}

func (q *queue) insert(x int) error {
	return q.pushStack.push(x)
}

func (q *queue) dequeue() (int, error) {
	if !q.popStack.isEmpty() {
		return q.popStack.pop()
	}
	for !q.pushStack.isEmpty() {
		val, err := q.pushStack.pop()
		if err != nil {
			return -1, err
		}
		if err := q.popStack.push(val); err != nil {
			return -1, err
		}
	}
	return q.popStack.pop()
}

func TestQueueUsingStacks() {
	q := NewQueue(4)
	q.insert(1)
	q.insert(2)
	val, _ := q.dequeue()
	fmt.Println(val)
	q.insert(3)
	val, _ = q.dequeue()
	fmt.Println(val)
}
