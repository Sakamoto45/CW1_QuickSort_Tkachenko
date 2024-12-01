package parallel

import (
	"runtime"
	"sync"
)

type ParallelQuickSorter struct {
	wg sync.WaitGroup
}

func (sorter *ParallelQuickSorter) Sort(array []int) {
	sorter.sort(array, runtime.GOMAXPROCS(0)*8)
	sorter.wg.Wait()
}

func (sorter *ParallelQuickSorter) sort(a []int, depth int) {
	if len(a) < 2 {
		return
	}

	left, right := 0, len(a)-1
	for i := range a {
		if a[i] < a[right] {
			a[i], a[left] = a[left], a[i]
			left++
		}
	}
	a[left], a[right] = a[right], a[left]

	if depth > 0 {
		sorter.wg.Add(2)
		go func() {
			sorter.sort(a[:left], depth/2)
			sorter.wg.Done()
		}()
		go func() {
			sorter.sort(a[left+1:], depth/2)
			sorter.wg.Done()
		}()
	} else {
		sorter.sort(a[:left], depth/2)
		sorter.sort(a[left+1:], depth/2)
	}
}
