package search

// Graph API for search operations
type Graph interface {
	// Size of vertices and edges
	Size() (vertices, edges int)

	// Neighbors of vertex
	Neighbors(v int) []int
}
