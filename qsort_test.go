package main

import (
	"fmt"
	"math/rand"
	goodparallel "quicksort/good_parallel"
	"quicksort/parallel"
	"quicksort/sequential"
	"testing"

	"github.com/stretchr/testify/require"
)

const N = 100000000

func BenchmarkQuickSorter(b *testing.B) {
	b.Run(fmt.Sprintf("Par(filter)"), func(b *testing.B) {
		sorter := parallel.NewParallelQuickSorter(4)
		b.StopTimer()

		for range 5 {
			a := rand.Perm(N)

			b.StartTimer()
			a = sorter.Sort(a)
			b.StopTimer()

			require.True(b, validate(a))
		}
	})

	b.Run(fmt.Sprintf("Par(good)"), func(b *testing.B) {
		sorter := goodparallel.NewParallelQuickSorter()
		b.StopTimer()

		for range 5 {
			a := rand.Perm(N)

			b.StartTimer()
			a = sorter.Sort(a)
			b.StopTimer()

			require.True(b, validate(a))
		}
	})

	b.Run(fmt.Sprintf("Seq"), func(b *testing.B) {
		sorter := sequential.NewSequentialQuickSorter()
		b.StopTimer()

		for range 5 {
			a := rand.Perm(N)

			b.StartTimer()
			a = sorter.Sort(a)
			b.StopTimer()

			require.True(b, validate(a))
		}
	})
}

func validate(a []int) bool {
	for i, v := range a {
		if i != v {
			return false
		}
	}
	return true
}
