package sequential

func NewSequentialQuickSorter() *SequentialQuickSorter {
	return &SequentialQuickSorter{}
}

type SequentialQuickSorter struct {
}

func (sorter *SequentialQuickSorter) Sort(array []int) []int {
	sorter.sort(array)
	return array
}

func (sorter *SequentialQuickSorter) sort(a []int) {
	if len(a) < 2 {
		return
	}

	l, r := 0, len(a)-1
	for i := range a {
		if a[i] < a[r] {
			a[i], a[l] = a[l], a[i]
			l++
		}
	}
	a[l], a[r] = a[r], a[l]

	sorter.sort(a[:l])
	sorter.sort(a[l+1:])
}
