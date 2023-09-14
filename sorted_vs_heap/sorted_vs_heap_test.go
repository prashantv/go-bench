package main

import (
	"cmp"
	"container/heap"
	"fmt"
	"math/rand"
	"slices"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

type sortedSlice []int

func newSortedSlice() *sortedSlice        { return &sortedSlice{} }
func (ss sortedSlice) Len() int           { return len(ss) }
func (ss sortedSlice) Less(i, j int) bool { return ss[i] < ss[j] }
func (ss sortedSlice) Swap(i, j int)      { ss[i], ss[j] = ss[j], ss[i] }

func (ss *sortedSlice) PushV(v int) {
	if cap(*ss) == 0 {
		*ss = make(sortedSlice, 0, 10)
	}
	*ss = append(*ss, v)
	sort.Sort(ss)
}

func (ss *sortedSlice) PopV() int {
	v := (*ss)[0]
	*ss = (*ss)[1:]
	return v
}

type sortedSliceGen []int

func newSortedSliceGen() *sortedSliceGen { return &sortedSliceGen{} }
func (ss sortedSliceGen) Len() int       { return len(ss) }
func (ss *sortedSliceGen) PushV(v int) {
	if cap(*ss) == 0 {
		*ss = make(sortedSliceGen, 0, 10)
	}
	*ss = append(*ss, v)
	slices.SortFunc(*ss, func(a, b int) int {
		return cmp.Compare(a, b)
	})
}

func (ss *sortedSliceGen) PopV() int {
	v := (*ss)[0]
	*ss = (*ss)[1:]
	return v
}

type heapSlice struct {
	values []int
	arg    int
}

func newHeapSlice() *heapSlice         { return &heapSlice{} }
func (h heapSlice) Len() int           { return len(h.values) }
func (h heapSlice) Less(i, j int) bool { return h.values[i] < h.values[j] }
func (h heapSlice) Swap(i, j int)      { h.values[i], h.values[j] = h.values[j], h.values[i] }

func (h *heapSlice) Push(v any) {
	h.values = append(h.values, h.arg)
}

func (h *heapSlice) Pop() any {
	n := len(h.values)
	h.arg = h.values[n-1]
	h.values = h.values[0 : n-1]
	return 0
}

func (h *heapSlice) PushV(v int) {
	h.arg = v
	heap.Push(h, 0)
}

func (h *heapSlice) PopV() int {
	heap.Pop(h)
	return h.arg
}

type Stack interface {
	PushV(v int)
	PopV() int
	Len() int
}

func BenchmarkSorted(b *testing.B) {
	runBenchmarks[*sortedSlice](b, newSortedSlice)
}

func BenchmarkSortedGen(b *testing.B) {
	runBenchmarks[*sortedSliceGen](b, newSortedSliceGen)
}

func TestSorted(t *testing.T) {
	runTests[*sortedSlice](t, newSortedSlice)
}

func TestSortedGenerics(t *testing.T) {
	runTests[*sortedSliceGen](t, newSortedSliceGen)
}

func BenchmarkHeap(b *testing.B) {
	runBenchmarks[*heapSlice](b, newHeapSlice)
}

func TestHeap(t *testing.T) {
	runTests[*heapSlice](t, newHeapSlice)
}

func runBenchmarks[S Stack](b *testing.B, newFn func() S) {
	for _, numElems := range []int{1, 10, 100, 1000} {
		elems := randElems(numElems)
		b.Run(fmt.Sprintf("%v shuffled elems", numElems), func(b *testing.B) {
			s := newFn()
			for i := 0; i < b.N; i++ {
				for _, v := range elems {
					s.PushV(v)
				}
				for range elems {
					s.PopV()
				}
			}
		})

		b.Run(fmt.Sprintf("%v sorted elems", numElems), func(b *testing.B) {
			s := newFn()
			for i := 0; i < b.N; i++ {
				for j := 0; j < numElems; j++ {
					s.PushV(j)
				}
				for j := 0; j < numElems; j++ {
					s.PopV()
				}
			}
		})
	}
}

func runTests[S Stack](t *testing.T, newFn func() S) {
	t.Run("single element", func(t *testing.T) {
		s := newFn()
		for i := 0; i < 100; i++ {
			s.PushV(i)
			assert.Equal(t, 1, s.Len())
			assert.Equal(t, i, s.PopV())
			assert.Equal(t, 0, s.Len())
		}
	})

	t.Run("multiple push/pop", func(t *testing.T) {
		push := []int{10, 5, 6, 3, 2, 8, 4, 9, 7, 1, 0}
		s := newFn()
		for _, v := range push {
			s.PushV(v)
		}

		// Pop some elements
		for i := 0; i < 5; i++ {
			assert.Equal(t, i, s.PopV())
		}

		// Then push them back
		for i := 0; i < 5; i++ {
			s.PushV(4 - i)
		}

		// Push all
		for i := 0; i < 10; i++ {
			assert.Equal(t, i, s.PopV())
		}
	})
}

func randElems(n int) []int {
	// Deterministic element ordering
	r := rand.New(rand.NewSource(1))

	var elems []int
	for i := 0; i < n; i++ {
		elems = append(elems, i)
	}
	r.Shuffle(len(elems), func(i, j int) {
		elems[i], elems[j] = elems[j], elems[i]
	})
	return elems
}
