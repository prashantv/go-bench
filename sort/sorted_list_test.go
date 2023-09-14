package sort

import (
	"fmt"
	"math/rand"
	"slices"
	"sort"
	"testing"
)

func sortInterface(ints []int) {
	sort.Ints(ints)
}

func sortInterfaceReverse(ints []int) {
	sort.Sort(sort.Reverse(sort.IntSlice(ints)))
}

func slicesSort(ints []int) {
	slices.SortFunc(ints, func(a, b int) int {
		return a - b
	})
}

func slicesSortReverse(ints []int) {
	slices.SortFunc(ints, func(a, b int) int {
		return b - a
	})
}

func BenchmarkSortReverse(b *testing.B) {
	for _, size := range []int{1, 10, 100, 1000} {
		elements := make([]int, size)
		for i := range elements {
			elements[i] = i
		}

		b.Run(fmt.Sprintf("slices/size=%d", size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				if i%2 == 0 {
					slicesSortReverse(elements)
				} else {
					slicesSort(elements)
				}
			}
		})

		b.Run(fmt.Sprintf("sortIface/size=%d", size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				if i%2 == 0 {
					sortInterfaceReverse(elements)
				} else {
					sortInterface(elements)
				}
			}
		})
	}
}

func BenchmarkSortRandom(b *testing.B) {
	for _, size := range []int{1, 10, 100, 1000} {
		orig := rand.Perm(size)
		elements := make([]int, size)
		b.Run(fmt.Sprintf("slices/size=%d", size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				copy(elements, orig)
				slicesSort(elements)
			}
		})

		b.Run(fmt.Sprintf("sortIface/size=%d", size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				copy(elements, orig)
				sortInterface(elements)
			}
		})
	}
}

func BenchmarkSortedList(b *testing.B) {
	for _, size := range []int{1, 10, 100, 1000} {
		elements := make([]int, size)
		for i := range elements {
			elements[i] = i
		}
		b.Run(fmt.Sprintf("slices/size=%d", size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				slicesSort(elements)
			}
		})

		b.Run(fmt.Sprintf("sortIface/size=%d", size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				sortInterface(elements)
			}
		})
	}
}
