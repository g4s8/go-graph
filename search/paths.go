package search

// Path in a graph
type Path []int

func SimplePaths(g Graph, from, to int) []Path {
	vs, _ := g.Size()
	visited := make([]bool, vs)
	var current intQstack
	var paths []Path

	var dfs func(g Graph, from, to int)
	dfs = func(g Graph, from, to int) {
		if visited[from] {
			return
		}
		current.push(from)
		visited[from] = true
		if from == to {
			p := current.slice()
			paths = append(paths, Path(p))
			current.pop()
			visited[from] = false
			return
		}
		for _, n := range g.Neighbors(from) {
			dfs(g, n, to)
		}
		current.pop()
		visited[from] = false
	}
	dfs(g, from, to)
	return paths
}
