package search

// DepthFirstSearch visits all vertices using DFS algorithm
// starting from v vertex.
func DepthFirstSearch(g Graph, v int, visitor Visitor) {
	var s intStack
	s.push(v)
	for !s.empty() && visitor.Continue() {
		n := s.pop()
		if !visitor.Visited(n) {
			visitor.Visit(n)
			nb := g.Neighbors(n)
			s.pushAll(nb)
		}
	}
}
