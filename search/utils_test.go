package search

import (
	"fmt"
	"testing"
)

func TestIntQueue(t *testing.T) {
	t.Run("enq-then-deq", func(t *testing.T) {
		const l = 10
		var q intQstack
		assertErrorf(t, q.empty(), "queue should be empty")
		for i := 0; i < l; i++ {
			q.enq(i)
		}
		assertErrorf(t, !q.empty(), "queue should not be empty")
		assertErrorf(t, q.len() == l, "expected %d, got %d", l, q.len())
		for i := 0; i < l; i++ {
			d := q.deq()
			assertErrorf(t, d == i, "[%d] expected %d, got %d", i, i, d)
		}
		assertErrorf(t, q.empty(), "queue should be empty")
	})
	t.Run("enq-deq-enq-deq", func(t *testing.T) {
		var q intQstack
		assertErrorf(t, q.empty(), "queue should be empty")
		for i := 0; i < 10; i++ {
			q.enq(i)
			d := q.deq()
			assertErrorf(t, d == i, "[%d] expected %d, got %d", i, i, d)
		}
	})
	t.Run("enq3-deq3", func(t *testing.T) {
		var q intQstack
		assertErrorf(t, q.empty(), "queue should be empty")
		expect := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
		assertErrorf(t, q.empty(), "queue should be empty")
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				q.enq(expect[i*3+j])
			}
			for j := 0; j < 3; j++ {
				d := q.deq()
				assertErrorf(t, d == expect[i*3+j], "[%d] expected %d, got %d", i*3+j, expect[i*3+j], d)
			}
		}
	})
	t.Run("enq-grow-deq", func(t *testing.T) {
		var q intQstack
		assertErrorf(t, q.empty(), "queue should be empty")
		expect := make([]int, 256)
		for i := 0; i < 256; i++ {
			expect[i] = i
		}
		for _, v := range expect {
			q.enq(v)
		}
		for i, v := range expect {
			d := q.deq()
			assertErrorf(t, d == v, "[%d] expected %d, got %d", i, v, d)
		}
		assertErrorf(t, q.empty(), "queue should be empty")
	})
	t.Run("enc-compact-deq", func(t *testing.T) {
		var q intQstack
		assertErrorf(t, q.empty(), "queue should be empty")
		for i := 0; i < 256; i++ {
			q.enq(i)
		}
		for i := 0; i < 128; i++ {
			d := q.deq()
			assertErrorf(t, d == i, "[%d] expected %d, got %d", i, i, d)
		}
		for i := 0; i < 256; i++ {
			q.enq(i + 256)
		}
		for i := 128; i < 512; i++ {
			d := q.deq()
			assertErrorf(t, d == i, "[%d] expected %d, got %d", i, i, d)
		}

		assertErrorf(t, q.empty(), "queue should be empty; len=%d", q.len())
	})
}

func TestIntStack(t *testing.T) {
	t.Run("push-then-pop", func(t *testing.T) {
		const l = 10
		var s intQstack
		for i := 0; i < l; i++ {
			s.push(i)
		}
		for i := l - 1; i >= 0; i-- {
			v := s.pop()
			assertErrorf(t, v == i, "[%d] expected %d, got %d", i, i, v)
		}
		assertErrorf(t, s.empty(), "stack should be empty")
	})
	t.Run("push-pop-push-pop", func(t *testing.T) {
		var s intQstack
		for i := 0; i < 10; i++ {
			s.push(i)
			v := s.pop()
			assertErrorf(t, v == i, "[%d] expected %d, got %d", i, i, v)
		}
		assertErrorf(t, s.empty(), "stack should be empty")
	})
	t.Run("push3-pop3", func(t *testing.T) {
		var s intQstack
		expect := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				s.push(expect[i*3+j])
			}
			for j := 2; j >= 0; j-- {
				v := s.pop()
				assertErrorf(t, v == expect[i*3+j], "[%d] expected %d, got %d", i*3+j, expect[i*3+j], v)
			}
		}
		assertErrorf(t, s.empty(), "stack should be empty")
	})
	t.Run("push-grow-pop", func(t *testing.T) {
		var s intQstack
		for i := 0; i < 256; i++ {
			s.push(i)
		}
		for i := 255; i >= 0; i-- {
			v := s.pop()
			assertErrorf(t, v == i, "[%d] expected %d, got %d", i, i, v)
		}
		assertErrorf(t, s.empty(), "stack should be empty")
	})
	t.Run("pushall", func(t *testing.T) {
		var s intQstack
		s.pushAll([]int{1, 2, 3, 4, 5})
		for i := 5; i > 0; i-- {
			v := s.pop()
			assertErrorf(t, v == i, "expected %d, got %d", i, v)
		}
	})
	t.Run("slise", func(t *testing.T) {
		var s intQstack
		for i := 0; i < 256; i++ {
			s.push(i)
		}
		sl := s.slice()
		assertErrorf(t, len(sl) == 256, "expected 256, got %d", len(sl))
		for i := 255; i >= 0; i-- {
			assertErrorf(t, sl[i] == i, "[%d] expected %d, got %d", i, i, sl[i])
		}
	})
}

func BenchmarkIntQueue(b *testing.B) {
	const count = 1000
	b.Run(fmt.Sprintf("enqN-deqM-%d", count), func(b *testing.B) {
		for n := 10; n <= 100; n += 20 {
			for m := 10; m <= n; m += 20 {
				b.Run(fmt.Sprintf("n%dm%d", n, m), func(b *testing.B) {
					for i := 0; i < b.N; i++ {
						var q intQstack
						var i int
						for i < count {
							for j := 0; j < n; j++ {
								q.enq(i)
								i++
							}
							for j := 0; j < m; j++ {
								q.deq()
							}
						}
					}
				})
			}
		}
	})
}

func BenchmarkIntStack(b *testing.B) {
	const count = 1000
	b.Run(fmt.Sprintf("pushN-popM-%d", count), func(b *testing.B) {
		for n := 10; n <= 100; n += 20 {
			for m := 10; m <= n; m += 20 {
				b.Run(fmt.Sprintf("n%dm%d", n, m), func(b *testing.B) {
					for i := 0; i < b.N; i++ {
						var s intQstack
						var i int
						for i < count {
							for j := 0; j < n; j++ {
								s.push(i)
								i++
							}
							for j := 0; j < m; j++ {
								s.pop()
							}
						}
					}
				})
			}
		}
	})
	b.Run("pushall", func(b *testing.B) {
		const (
			count     = 1000
			batchSize = 100
		)
		batch := make([]int, batchSize)
		for i := 0; i < batchSize; i++ {
			batch[i] = i
		}
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			var s intQstack
			for i := 0; i < count; i += batchSize {
				s.pushAll(batch)
			}
		}
	})
}
