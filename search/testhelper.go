package search

type testGraph struct {
	size     struct{ vertices, edges int }
	edges    map[int][]int
	directed bool
}

func newTestGraph(v int, directed bool) *testGraph {
	var g testGraph
	g.size.vertices = v
	g.edges = make(map[int][]int)
	g.directed = directed
	for i := 0; i < v; i++ {
		g.edges[i] = make([]int, 0)
	}
	return &g
}

func (g *testGraph) Size() (int, int) {
	return g.size.vertices, g.size.edges
}

func (g *testGraph) Neighbors(v int) []int {
	return g.edges[v]
}

func (g *testGraph) connect(a, b int) {
	g.edges[a] = append(g.edges[a], b)
	if !g.directed {
		g.edges[b] = append(g.edges[b], a)
	}
	g.size.edges++
}
