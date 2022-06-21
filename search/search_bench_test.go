package search

import (
	"fmt"
	"math/rand"
	"testing"
)

func BenchmarkSearch(b *testing.B) {
	b.Run("graph-lines", func(b *testing.B) {
		for i := 10; i < 1_000_001; i *= 10 {
			g := newTestGraph(i, true)
			for j := 0; j < i-1; j++ {
				g.connect(j, j+1)
			}
			b.Run(fmt.Sprintf("v-%d", i), func(b *testing.B) {
				b.Run("DFS", searchBenchRunner(g, DepthFirstSearch))
				b.Run("BFS", searchBenchRunner(g, BreathFirstSearch))
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
				b.Run("DFS", searchBenchRunner(g, DepthFirstSearch))
				b.Run("BFS", searchBenchRunner(g, BreathFirstSearch))
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
					b.Run("DFS", searchBenchRunner(g, DepthFirstSearch))
					b.Run("BFS", searchBenchRunner(g, BreathFirstSearch))
				})
			}
		}
	})
}

func searchBenchRunner(g Graph, search func(Graph, int, Visitor)) func(b *testing.B) {
	return func(b *testing.B) {
		visitor := NewFullVisitor(g)
		b.ResetTimer()
		b.ReportAllocs()
		search(g, 0, visitor)
	}
}
