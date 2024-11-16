package search

// BreathFirstSearch visits all vertices in graph using BFS
// algorithm starting from v vertex.
func BreathFirstSearch(g Graph, v int, visitor Visitor) {
	var q intQstack
	q.enq(v)
	for !q.empty() {
		next := q.deq()
		if v := visitor.Visit(next); v == nil {
			return
		}

		for _, n := range g.Neighbors(next) {
			if !visitor.Visited(n) {
				q.enq(n)
			}
		}
	}
}
