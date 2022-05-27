package search

import "testing"

func TestFindCicle(t *testing.T) {
	g1 := newTestGraph(4, true)
	g1.connect(0, 1)
	g1.connect(1, 2)
	g1.connect(2, 3)
	g1.connect(3, 1)
	t.Run("simple circle", testCicle(t, g1, true))

	g2 := newTestGraph(4, true)
	g2.connect(0, 3)
	g2.connect(1, 2)
	g2.connect(0, 1)
	g2.connect(2, 0)
	g2.connect(3, 1)
	t.Run("multipath circle", testCicle(t, g2, true))

	g3 := newTestGraph(4, true)
	g3.connect(0, 1)
	g3.connect(1, 2)
	g3.connect(2, 0)
	g3.connect(0, 3)
	t.Run("choose correct path 1", testCicle(t, g3, true))

	g4 := newTestGraph(4, true)
	g4.connect(0, 1)
	g4.connect(1, 2)
	g4.connect(0, 2)
	g4.connect(1, 3)
	t.Run("choose correct path 2", testCicle(t, g4, true))
}

func TestNoCicle(t *testing.T) {
	g1 := newTestGraph(4, true)
	g1.connect(0, 1)
	g1.connect(2, 3)
	g1.connect(1, 3)
	t.Run("simple", testCicle(t, g1, false))

	g2 := newTestGraph(6, true)
	g2.connect(0, 3)
	g2.connect(1, 4)
	g2.connect(2, 5)
	g2.connect(0, 1)
	g2.connect(4, 5)
	t.Run("snake", testCicle(t, g1, false))
}

func testCicle(t *testing.T, g Graph, expect bool) func(t *testing.T) {
	t.Helper()
	return func(t *testing.T) {
		if HasCicle(g) != expect {
			t.Fail()
		}
	}
}
