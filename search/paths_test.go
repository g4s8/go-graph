package search

import (
	"fmt"
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
		{
			size:     5,
			directed: true,
			connections: [][2]int{
				{0, 1}, {1, 2}, {0, 2}, {1, 3},
				{2, 3}, {2, 4}, {4, 3}, {3, 4},
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
				{0, 1}, {0, 2}, {0, 3},
				{1, 4}, {2, 4}, {3, 4},
				{0, 4},
			},
			path: [2]int{0, 4},
			expect: [][]int{
				{0, 1, 4}, {0, 2, 4}, {0, 3, 4},
				{0, 4},
			},
		},
		{
			size:     4,
			directed: true,
			connections: [][2]int{
				{0, 1}, {0, 2},
				{1, 3}, {2, 3},
				{0, 3},
				{1, 2}, {2, 1},
			},
			path: [2]int{0, 3},
			expect: [][]int{
				{0, 3},
				{0, 1, 3}, {0, 2, 3},
				{0, 1, 2, 3}, {0, 2, 1, 3},
			},
		},
		{
			size:     5,
			directed: true,
			connections: [][2]int{
				{0, 4},
				{0, 1}, {1, 4},
				{0, 2}, {2, 4},
				{0, 3}, {3, 4},
				{1, 2}, {2, 1}, {2, 3}, {3, 2}, {1, 3}, {3, 1},
			},
			path: [2]int{0, 4},
			expect: [][]int{
				{0, 4},
				{0, 1, 4}, {0, 2, 4}, {0, 3, 4},
				{0, 1, 2, 4}, {0, 1, 3, 4},
				{0, 2, 1, 4}, {0, 2, 3, 4},
				{0, 3, 1, 4}, {0, 3, 2, 4},
				{0, 1, 2, 3, 4}, {0, 1, 3, 2, 4},
				{0, 2, 1, 3, 4}, {0, 2, 3, 1, 4},
				{0, 3, 1, 2, 4}, {0, 3, 2, 1, 4},
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
