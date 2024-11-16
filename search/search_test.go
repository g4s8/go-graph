package search

import (
	"fmt"
	"testing"

	m "github.com/g4s8/go-matchers"
)

func TestBreathFirstSearch(t *testing.T) {
	t.Run("BFS", func(t *testing.T) {
		testSearch(t, BreathFirstSearch)
	})
}

func TestDepthFirstSearch(t *testing.T) {
	t.Run("DFS", func(t *testing.T) {
		testSearch(t, DepthFirstSearch)
	})
}

func TestDepthFirstSearchRecursive(t *testing.T) {
	t.Run("DFS-recursive", func(t *testing.T) {
		testSearch(t, DepthFirstSearchRecursive)
	})
}

func testSearch(t *testing.T, search func(Graph, int, Visitor)) {
	g := newTestGraph(6, false)
	g.connect(0, 1)
	g.connect(1, 2)
	g.connect(2, 3)
	g.connect(0, 2)
	g.connect(1, 3)
	g.connect(4, 5)

	visitor := NewFullVisitor(g)
	search(g, 1, visitor)

	assert := m.Assert(t)
	for v := 0; v < 4; v++ {
		assert.That(fmt.Sprintf("visited-%d", v),
			visitor.Visited(v), m.Is(true))
	}
	for v := 4; v < 6; v++ {
		assert.That(fmt.Sprintf("not-visited-%d", v),
			visitor.Visited(v), m.Is(false))
	}
}
