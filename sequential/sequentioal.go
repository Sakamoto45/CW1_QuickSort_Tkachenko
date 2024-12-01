package sequential

type SequentialQuickSorter struct {
}

func (sorter *SequentialQuickSorter) Sort(array []int) {
	sorter.sort(array)
}

func (sorter *SequentialQuickSorter) sort(a []int) {
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

	sorter.sort(a[:left])
	sorter.sort(a[left+1:])
}
