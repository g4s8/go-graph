package search

// BreathFirstSearch visits all vertices in graph using BFS
// algorithm starting from v vertex.
func BreathFirstSearch(g Graph, v int, visitor Visitor) {
	var q intQstack
	q.enq(v)
	for !q.empty() && visitor.Continue() {
		next := q.deq()
		visitor.Visit(next)
		for _, n := range g.Neighbors(next) {
			if !visitor.Continue() {
				return
			}
			if !visitor.Visited(n) {
				q.enq(n)
			}
		}
	}
}
