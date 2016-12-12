package main

import (
	"fmt"
)

type nextTree struct{
	node
	next *nextTree
}

func newNextTree() *nextTree{
	return &nextTree{
		new
	}
}