package search

import (
	m "github.com/g4s8/go-matchers"
	"testing"
)

func TestSearchPaths(t *testing.T) {
	g := newTestGraph(5, true)
	g.connect(0, 1)
	g.connect(1, 2)
	g.connect(0, 2)
	g.connect(1, 3)
	g.connect(2, 3)
	g.connect(2, 4)
	g.connect(4, 3)
	g.connect(3, 4)

	paths := SimplePaths(g, 0, 3)

	assert := m.Assert(t)
	assert.That("find 5 paths", paths, m.LenIs(5))
	assert.That("find 0-1-2-3 path", paths,
		m.HasItemEq(Path{0, 1, 2, 3}))
	assert.That("find 0-1-2-4-3 path", paths,
		m.HasItemEq(Path{0, 1, 2, 4, 3}))
	assert.That("find 0-1-3 path", paths,
		m.HasItemEq(Path{0, 1, 3}))
	assert.That("find 0-2-3 path", paths,
		m.HasItemEq(Path{0, 2, 3}))
	assert.That("find 0-2-4-3 path", paths,
		m.HasItemEq(Path{0, 2, 4, 3}))
}
