This go library provides packages with functions for
common graph operations. It doesn't provide graph data structures.
It's designed to be structure agnostic,
it depends only on required API for each operation declared as interface
(similar to standard Go `sort` package).

[![CI](https://github.com/g4s8/go-graph/actions/workflows/go.yml/badge.svg)](https://github.com/g4s8/go-graph/actions/workflows/go.yml)
[![GoDoc](https://godoc.org/github.com/g4s8/go-graph?status.svg)](https://godoc.org/github.com/g4s8/go-graph)


## Packages

### Search

Package search provides graph search operations:
 - `DepthFirstSearch(Graph g, v int, visitor Visitor)` - implements DFS stack-based
 algorithm for graph `g` from vertex `v` using `visitor`.
 - `BreathFirstSearch(g Graph, v int, visitor Visitor)` - implements BFS queue-based
 algorithm for graph `g` from vertex `v` using `visitor`.
 - `HasCicle(g Graph)` - detects if graph `g` has any cicle.
 - `SimplePaths(g Graph, from, to int)` - finds all simple paths (without cicles)
 in graph `g` from vertex `from` to `to`.

Graph type for `search` package should be able to say its size of vertices and edges,
and iterate all neighors for each vertex:
```go
type Graph interface {
	// Size of vertices and edges
	Size() (vertices, edges int)

	// Neighbors of vertex
	Neighbors(v int) []int
}
```

The `search` package provides standard visitors for DFS and BFS like:
 - `NewFullVisitor(Graph)` which allows to visit all vertices
 - `NewShortPathVisitor(g Graph, vertex int)` which stops when destination `vertex`
 is visited.

Users are free to implement custom visitor as `Visitor` interface.
