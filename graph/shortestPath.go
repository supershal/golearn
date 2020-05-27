package main

import "fmt"

// A simple graph can be presented with map[int][]int instead of complex structure. it would be easier in the interview.

type sGraph map[int][]int

// brute force
/*
1. undestand problem.
  path can be between a to b or b to a in graph.
  what if there is loop.
2. bruteforuce
  DFS from a to its children until b is found.
   for each iteration incrase path variable.
	b is found then return path. or return 0
	if loop is found then return 0
  same for b to a
  min of both path
3. optimize
   a better way to find path in the graph is to use BFS because if a node is closer to the root node, it can be found faster.
   If we use DFS we may go through longest path first.
4. code
5. test
*/

func (g sGraph) shortestPathBFS(a, b int) []int {
	// a and b are same
	if a == b {
		return []int{}
	}
	// a or b not found
	if _, ok := g[a]; !ok {
		return []int{}
	}

	// mistake here. a path does not have to be originated from b.
	// if _, ok := g[b]; !ok {
	// 	return []int{}
	// }

	queue := []int{a}
	visited := map[int]bool{}
	parents := map[int]int{
		a: -1,
	}

	for len(queue) != 0 {
		curr := queue[0]
		visited[curr] = true
		queue = queue[1:]

		if curr == b {
			break
		}
		for _, c := range g[curr] {
			if _, ok := visited[c]; !ok {
				queue = append(queue, c)
				parents[c] = curr
			}
		}
	}

	// path not found
	if _, ok := parents[b]; !ok {
		return []int{}
	}
	path := []int{b}
	curr := b
	for {
		parent := parents[curr]
		if parent == -1 {
			break
		}
		path = append(path, parent)
		curr = parent
	}
	return path
}

func (g sGraph) shortestPathDFS(a, b int) []int {
	return findTargetNode(g, a, b, map[int]bool{}, []int{}, []int{})
}

func findTargetNode(g sGraph, curr int, target int, visited map[int]bool, currPath, maxPath []int) []int {
	// check loop
	if _, ok := visited[curr]; ok {
		return []int{}
	}
	// mark node as visited
	visited[curr] = true
	currPath = append(currPath, curr)
	// visit each child node
	for _, c := range g[curr] {
		if c == target {
			return append(currPath, c)
		}
		// DFS each child node
		path := findTargetNode(g, c, target, visited, currPath, maxPath)
		if len(path) > len(maxPath) {
			maxPath = path
		}
	}
	// path not found
	return maxPath
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func TestShortestPath() {
	g := sGraph{}
	g[2] = []int{5}
	g[5] = []int{4}
	g[4] = []int{3, 1}
	g[1] = []int{2, 3}
	fmt.Println("BFS shortestpath:", g.shortestPathBFS(2, 3))
	fmt.Println("DFS shortestpath:", g.shortestPathDFS(2, 3))
}
