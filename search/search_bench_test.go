package search

import (
	"fmt"
	"math"
	"testing"
)

func BenchmarkBreathFirstSearch(b *testing.B) {
	for _, spec := range []struct{ epv, depth int }{
		{2, 2}, {2, 3}, {2, 4}, {2, 5},
		{5, 2}, {5, 3}, {5, 4}, {5, 5},
		{10, 2}, {10, 3}, {10, 4}, {10, 5},
		{15, 2}, {15, 3}, {15, 4}, {15, 5},
	} {
		b.Run(fmt.Sprintf("DFS-%d-%d", spec.epv, spec.depth), func(b *testing.B) {
			benchmarkBranching(b, spec.epv, spec.depth, DepthFirstSearch)
		})
		b.Run(fmt.Sprintf("BFS-%d-%d", spec.epv, spec.depth), func(b *testing.B) {
			benchmarkBranching(b, spec.epv, spec.depth, BreathFirstSearch)
		})
	}
}

func benchmarkBranching(b *testing.B, epv int, depth int, search func(Graph, int, Visitor)) {
	var size int
	for d := depth; d >= 0; d-- {
		size += int(math.Pow(float64(epv), float64(d)))
	}

	g := newTestGraph(size, false)
	var pos int
	fillBranchingGraph(g, 0, &pos, epv, depth)

	visitor := NewFullVisitor(g)
	b.ResetTimer()
	search(g, 0, visitor)
}

func fillBranchingGraph(g *testGraph, v int, pos *int, epv int, depth int) {
	for i := 0; i < epv; i++ {
		x := *pos + i
		g.connect(v, x)
	}
	*pos += epv
	if depth < 1 {
		return
	}
	for i := 0; i < epv; i++ {
		x := *pos + i
		fillBranchingGraph(g, x, pos, epv, depth-1)
	}
}
