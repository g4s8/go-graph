package search

// DepthFirstSearch visits all vertices using DFS algorithm
// starting from v vertex.
func DepthFirstSearch(g Graph, v int, visitor Visitor) {
	var s intQstack
	s.push(v)
	for !s.empty() {
		n := s.pop()
		if !visitor.Visited(n) {
			visitor.Visit(n)
			nb := g.Neighbors(n)
			s.pushAll(nb)
		}
	}
}

// DepthFirstSearchRecursive visits all vertices using DFS algorithm
// starting from v vertex.
func DepthFirstSearchRecursive(g Graph, v int, visitor Visitor) {
	next := visitor.Visit(v)
	if next == nil {
		return
	}
	for _, n := range g.Neighbors(v) {
		if next.Visited(n) {
			continue
		}
		DepthFirstSearchRecursive(g, n, next)
	}
}
