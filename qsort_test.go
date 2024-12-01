package main

import (
	"fmt"
	"math/rand"
	"quicksort/parallel"
	"quicksort/sequential"
	"testing"

	"github.com/stretchr/testify/require"
)

const N = 100000000

func BenchmarkParQuickSorter(b *testing.B) {
	b.Run(fmt.Sprintf("Par"), func(b *testing.B) {
		sorter := parallel.ParallelQuickSorter{}
		b.StopTimer()

		for range 5 {
			a := rand.Perm(N)

			b.StartTimer()
			sorter.Sort(a)
			b.StopTimer()

			require.True(b, validate(a))
		}
	})
}

func BenchmarkSeqQuickSorter(b *testing.B) {
	b.Run(fmt.Sprintf("Seq"), func(b *testing.B) {
		sorter := sequential.SequentialQuickSorter{}
		b.StopTimer()

		for range 5 {
			a := rand.Perm(N)

			b.StartTimer()
			sorter.Sort(a)
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
