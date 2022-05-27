package search

// Path in a graph
type Path []int

// SimplePaths finds all simple (without cicles) paths in graph.
func SimplePaths(g Graph, from, to int) []Path {
	v, _ := g.Size()
	search := &simplePathSearch{
		visited: make([]bool, v),
		paths:   make([]Path, 0),
	}
	search.dfs(g, from, to)
	return search.paths
}

type simplePathSearch struct {
	visited []bool
	current intStack
	paths   []Path
}

func (sps *simplePathSearch) dfs(g Graph, from, to int) {
	if sps.visited[from] {
		return
	}
	sps.current.push(from)
	sps.visited[from] = true
	if from == to {
		p := sps.current.sliceReverted()
		sps.paths = append(sps.paths, Path(p))
		sps.current.pop()
		sps.visited[from] = false
		return
	}
	for _, n := range g.Neighbors(from) {
		sps.dfs(g, n, to)
	}
	sps.current.pop()
	sps.visited[from] = false
}
