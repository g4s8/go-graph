package search

import (
	"fmt"
	"math/rand"
	"testing"
)

func BenchmarkSearch(b *testing.B) {
	for _, t := range []struct {
		name   string
		search func(Graph, int, Visitor)
	}{
		{"DFS", DepthFirstSearch},
		{"BFS", BreathFirstSearch},
	} {
		b.Run(t.name, func(b *testing.B) {
			b.Run("graph-lines", func(b *testing.B) {
				for i := 10; i <= 1_000_000; i *= 10 {
					g := newTestGraph(i, true)
					for j := 0; j < i-1; j++ {
						g.connect(j, j+1)
					}
					b.Run(fmt.Sprintf("v-%d", i), func(b *testing.B) {
						runSearchBench(b, g, t.search)
					})
				}
			})
			b.Run("star-graphs", func(b *testing.B) {
				for i := 10; i < 1_000_001; i *= 10 {
					g := newTestGraph(i, true)
					for j := 1; j < i; j++ {
						g.connect(0, j)
					}
					b.Run(fmt.Sprintf("v-%d", i), func(b *testing.B) {
						runSearchBench(b, g, t.search)
					})
				}
			})
			b.Run("connected-graphs", func(b *testing.B) {
				for i := 10; i < 100_000; i *= 10 {
					g := newTestGraph(i, true)
					// connect each vertex to x pseudo-random vertices
					for x := 2; x < 10; x++ {
						rng := rand.New(rand.NewSource(0))
						for j := 0; j < i; j++ {
							for k := 0; k < x; k++ {
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
						b.Run(fmt.Sprintf("v-%d x-%d", i, x), func(b *testing.B) {
							runSearchBench(b, g, t.search)
						})
					}
				}
			})
		})
	}
}

func runSearchBench(b *testing.B, g *testGraph, search func(Graph, int, Visitor)) {
	b.Helper()

	visitor := NewFullVisitor(g).(visitor)
	for i := 0; i < b.N; i++ {
		visitor.reset()
		search(g, 0, visitor)
	}
}
