package parallel

import (
	// "quicksort/sequential"

	"quicksort/sequential"
	"sync"
)

const block = 500000

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
}

func (sorter *ParallelQuickSorter) Sort(array []int) []int {
	for i := range sorter.numWorkers {
		sorter.workersLeft[i] = make([]int, 0, len(array))
		sorter.workersRight[i] = make([]int, 0, len(array))
	}

	return sorter.sort(array)
}

func (sorter *ParallelQuickSorter) sort(a []int) []int {
	if len(a) < 2 {
		return a
	}

	if len(a) < block {
		sequential.NewSequentialQuickSorter().Sort(a)
		return a
	}

	pivot := a[len(a)-1]

	left, right := sorter.parallelFilter(a, pivot)

	left = sorter.sort(left)
	right = sorter.sort(right)

	middleSize := len(a) - len(left) - len(right)
	for range middleSize {
		left = append(left, pivot)
	}

	return append(left, right...)
}

// parallel filter optimized for this task
func (sorter *ParallelQuickSorter) parallelFilter(nums []int, pivot int) ([]int, []int) {
	n := len(nums)
	chunkSize := n / sorter.numWorkers

	resultLeft := make([]int, 0, n)
	resultRight := make([]int, 0, n)

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
			resultLeft = append(resultLeft, left...)
			resultRight = append(resultRight, right...)
			sorter.mu.Unlock()

		}(start, end, i)
	}

	sorter.wg.Wait()
	return resultLeft, resultRight
}
