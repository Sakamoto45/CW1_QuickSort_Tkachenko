package parallel

import (
	// "quicksort/sequential"

	"quicksort/sequential"
	"sync"
)

const block = 500

func NewParallelQuickSorter(numWorkers int) *ParallelQuickSorter {
	sorter := &ParallelQuickSorter{
		numWorkers: numWorkers,
	}

	sorter.workersLeft = make([][]int, numWorkers)
	sorter.workersRight = make([][]int, numWorkers)

	return sorter
}

type ParallelQuickSorter struct {
	wg         sync.WaitGroup
	mu         sync.Mutex
	numWorkers int

	workersLeft  [][]int
	workersRight [][]int
	resultLeft   []int
	resultRight  []int
}

func (sorter *ParallelQuickSorter) Sort(array []int) []int {
	for i := range sorter.numWorkers {
		sorter.workersLeft[i] = make([]int, 0, len(array)/sorter.numWorkers)
		sorter.workersRight[i] = make([]int, 0, len(array)/sorter.numWorkers)
	}
	sorter.resultLeft = make([]int, 0, len(array))
	sorter.resultRight = make([]int, 0, len(array))

	sorter.sort(array)
	return array
}

func (sorter *ParallelQuickSorter) sort(a []int) {
	if len(a) < 2 {
		return
	}

	if len(a) < block {
		sequential.NewSequentialQuickSorter().Sort(a)
		return
	}

	pivot := a[len(a)-1]

	left, right := sorter.parallelFilter(a, pivot)

	l := len(left)
	r := len(right)
	middleSize := len(a) - l - r

	a = a[:0]
	a = append(a, left...)
	for range middleSize {
		a = append(a, pivot)
	}
	a = append(a, right...)

	sorter.sort(a[:l])
	sorter.sort(a[l+middleSize:])
}

// parallel filter optimized for this task
func (sorter *ParallelQuickSorter) parallelFilter(nums []int, pivot int) ([]int, []int) {
	n := len(nums)
	chunkSize := n / sorter.numWorkers

	sorter.resultLeft = sorter.resultLeft[:0]
	sorter.resultRight = sorter.resultRight[:0]

	for i := 0; i < sorter.numWorkers; i++ {
		start := i * chunkSize
		end := start + chunkSize
		if i == sorter.numWorkers-1 {
			end = n
		}
		sorter.wg.Add(1)
		go func(start, end, id int) {
			defer sorter.wg.Done()
			left := sorter.workersLeft[id]
			left = left[:0]
			right := sorter.workersRight[id]
			right = right[:0]

			for j := start; j < end; j++ {
				if nums[j] < pivot {
					left = append(left, nums[j])
				} else if nums[j] > pivot {
					right = append(right, nums[j])
				}
			}

			sorter.mu.Lock()
			sorter.resultLeft = append(sorter.resultLeft, left...)
			sorter.resultRight = append(sorter.resultRight, right...)
			sorter.mu.Unlock()
		}(start, end, i)
	}

	sorter.wg.Wait()
	return sorter.resultLeft, sorter.resultRight
}
