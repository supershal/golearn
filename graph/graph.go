package main

type children []int

type graph struct {
	vertices []int
	adj      map[int]*children
	visited  map[int]bool
}

func newGraph() *graph {
	return &graph{
		vertices: []int{},
		adj:      map[int]*children{},
		visited:  make(map[int]bool),
	}
}

func (g *graph) addVertex(x int) {
	g.vertices = append(g.vertices, x)
	g.visited[x] = false
}

func (g *graph) addEdge(x, y int) {
	if _, ok := g.adj[x]; !ok {
		g.adj[x] = &children{}
	}
	children := g.adj[x]
	*children = append(*children, y)
	g.adj[x] = children
}

func (g *graph) dfs(x int) {
	g.visited[x] = true
	for _, child := range *g.adj[x] {
		if !g.visited[child] {
			g.dfs(child)
		}
	}
}

func (g *graph) bfs(x int) {
	queue := make([]int, 0)
	queue = append(queue, x)
	for len(queue) != 0 {
		curr := queue[0]
		queue := queue[1:]
		g.visited[curr] = true
		for _, child := range *g.adj[curr] {
			if !g.visited[child] {
				queue = append(queue, child)
			}
		}

	}
}

func main() {

}
