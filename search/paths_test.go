package search

import (
	"fmt"
	"math/rand"
	"testing"

	m "github.com/g4s8/go-matchers"
)

func TestSearchPaths(t *testing.T) {
	for i, tc := range []struct {
		size        int
		directed    bool
		connections [][2]int
		path        [2]int
		expect      [][]int
	}{
		// FIXME: this test is failing
		// {
		// 	size:        4,
		// 	directed:    true,
		// 	connections: [][2]int{},
		// },

		/*
			                          +-----------------+
			                          |                 |
					 +-----------------+        |
			                 |        |        ↓        ↓
				0 -----> 1 -----> 2 -----> 3 -----> 4
				|                 ↑        ↑        ↑
				+-----------------+        ↓        ↓
							   +--------+
		*/
		{
			size:     5,
			directed: true,
			connections: [][2]int{
				{0, 1},
				{1, 2},
				{0, 2},
				{1, 3},
				{2, 3},
				{2, 4},
				{4, 3},
				{3, 4},
			},
			path: [2]int{0, 3},
			expect: [][]int{
				{0, 1, 2, 3},
				{0, 1, 2, 4, 3},
				{0, 1, 3},
				{0, 2, 3},
				{0, 2, 4, 3},
			},
		},
		{
			size:     5,
			directed: true,
			connections: [][2]int{
				{0, 1},
				{0, 2},
				{0, 3},
				{1, 4},
				{2, 4},
				{3, 4},
				{0, 4},
			},
			path: [2]int{0, 4},
			expect: [][]int{
				{0, 1, 4},
				{0, 2, 4},
				{0, 3, 4},
				{0, 4},
			},
		},
		{
			size:     4,
			directed: true,
			connections: [][2]int{
				{0, 1},
				{0, 2},
				{1, 3},
				{2, 3},
				{0, 3},
				{1, 2},
				{2, 1},
			},
			path: [2]int{0, 3},
			expect: [][]int{
				{0, 3},
				{0, 1, 3},
				{0, 2, 3},
				{0, 1, 2, 3},
				{0, 2, 1, 3},
			},
		},
		{
			size:     5,
			directed: true,
			connections: [][2]int{
				{0, 4},
				{0, 1},
				{1, 4},
				{0, 2},
				{2, 4},
				{0, 3},
				{3, 4},
				{1, 2},
				{2, 1},
				{2, 3},
				{3, 2},
				{1, 3},
				{3, 1},
			},
			path: [2]int{0, 4},
			expect: [][]int{
				{0, 4},
				{0, 1, 4},
				{0, 2, 4},
				{0, 3, 4},
				{0, 1, 2, 4},
				{0, 1, 3, 4},
				{0, 2, 1, 4},
				{0, 2, 3, 4},
				{0, 3, 1, 4},
				{0, 3, 2, 4},
				{0, 1, 2, 3, 4},
				{0, 1, 3, 2, 4},
				{0, 2, 1, 3, 4},
				{0, 2, 3, 1, 4},
				{0, 3, 1, 2, 4},
				{0, 3, 2, 1, 4},
			},
		},
	} {
		g := newTestGraph(tc.size, tc.directed)
		for _, con := range tc.connections {
			g.connect(con[0], con[1])
		}
		t.Run(fmt.Sprintf("case-%d", i), func(t *testing.T) {
			paths := SimplePaths(g, tc.path[0], tc.path[1])
			assert := m.Assert(t)
			assert.That("size", paths, m.LenIs(len(tc.expect)))
			for _, p := range tc.expect {
				assert.That(fmt.Sprintf("expect-%v", p),
					paths, m.HasItemEq(Path(p)))
			}
		})
	}
}

func BenchmarkSearchPaths(b *testing.B) {
	benchRunner := func(g Graph, from, to int, f func(g Graph, from, to int) []Path) func(b *testing.B) {
		return func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				_ = f(g, from, to)
			}
		}
	}

	for _, tc := range []struct {
		name string
		f    func(g Graph, from, to int) []Path
	}{
		{"SimplePaths", SimplePaths},
	} {
		b.Run(tc.name, func(b *testing.B) {
			b.Run("graph-lines", func(b *testing.B) {
				for i := 10; i <= 10_000; i *= 10 {
					g := newTestGraph(i, true)
					for j := 0; j < i-1; j++ {
						g.connect(j, j+1)
					}
					b.Run(fmt.Sprintf("v-%d", i), benchRunner(g, 0, i-1, tc.f))
				}
			})
			b.Run("star-graphs", func(b *testing.B) {
				for i := 10; i <= 10_000; i *= 10 {
					g := newTestGraph(i, true)
					for j := 1; j < i; j++ {
						g.connect(0, j)
					}
					b.Run(fmt.Sprintf("v-%d", i), benchRunner(g, 0, i-1, tc.f))
				}
			})
			b.Run("connected-graphs", func(b *testing.B) {
				for i := 10; i <= 80; i *= 2 {
					g := newTestGraph(i, true)
					// connect each vertex to 2 pseudo-random vertices
					rng := rand.New(rand.NewSource(0))
					for j := 0; j < i; j++ {
						for k := 0; k < 2; k++ {
							var other int
							for {
								other = rng.Intn(i)
								if other != j {
									break
								}
							}
							g.connect(j, other)
						}
					}
					b.Run(fmt.Sprintf("v-%d", i), benchRunner(g, 0, i-1, tc.f))
				}
			})
		})
	}
}
