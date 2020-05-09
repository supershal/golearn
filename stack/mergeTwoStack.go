package main

// implement two stacks in single array.
// start from both ends and overflow when top collides.

import "fmt"

type twoStacks struct {
	stack      []int
	top1, top2 int
}

func NewTwoStacks(len int) *twoStacks {
	return &twoStacks{
		stack: make([]int, len),
		top1:  -1,
		top2:  len,
	}
}

func (t *twoStacks) Push(stackNo int, val int) error {
	// validate stackNo
	if stackNo < 1 || stackNo > 2 {
		return fmt.Errorf("invalid params")
	}
	if stackNo == 1 {
		return t.push1(val)
	}
	return t.push2(val)
}

func (t *twoStacks) push1(val int) error {
	if t.top1+1 == t.top2 {
		return fmt.Errorf("stack overflow")
	}
	t.stack[t.top1+1] = val
	t.top1++
	return nil
}

func (t *twoStacks) push2(val int) error {
	if t.top2-1 == t.top1 {
		return fmt.Errorf("stack overflow")
	}
	t.stack[t.top2-1] = val
	t.top2--
	return nil
}

func (t *twoStacks) Pop(stackNo int) (int, error) {
	if stackNo == 1 {
		return t.pop1()
	}
	return t.pop2()
}

func (t *twoStacks) pop1() (int, error) {
	if t.top1-1 < 0 {
		return 0, fmt.Errorf("stack empty")
	}
	val := t.stack[t.top1]
	t.top1--
	return val, nil
}

func (t *twoStacks) pop2() (int, error) {
	if t.top2+1 >= len(t.stack) {
		return 0, fmt.Errorf("stack empty")
	}
	val := t.stack[t.top2]
	t.top2++
	return val, nil
}

func (t *twoStacks) Print() {
	fmt.Printf("+%v\n", t)
}

// TestTwoStacks foo
func TestTwoStacks() {
	ts := NewTwoStacks(5)
	ts.Push(1, 5)
	ts.Push(2, 10)
	ts.Push(2, 15)
	ts.Push(1, 11)
	ts.Push(2, 7)

	ts.Print()
	p1, err1 := ts.Pop(1)
	p2, err2 := ts.Pop(2)
	fmt.Println("pop1 = ", p1, err1)
	fmt.Println("pop2 =", p2, err2)
}
