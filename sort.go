package genericz

import (
	"sort"

	"golang.org/x/exp/constraints"
)

// SortWithData sorts two slices together. The slices will be sorted in order of comparables. The data slice is only reordered.
// SortWithData may fail to sort correctly when sorting slices of floating-point numbers containing Not-a-number (NaN) values.
func SortWithData[C constraints.Ordered, D any](comparables []C, data []D) {
	sort.Sort(comparablesAndData[C, D]{comparables, data})
}

type comparablesAndData[C constraints.Ordered, D any] struct {
	comparables []C
	data        []D
}

func (s comparablesAndData[C, D]) Len() int {
	return len(s.comparables)
}

func (s comparablesAndData[C, D]) Less(i, j int) bool {
	return s.comparables[i] < s.comparables[j]
}

func (s comparablesAndData[C, D]) Swap(i, j int) {
	s.comparables[i], s.comparables[j] = s.comparables[j], s.comparables[i]
	s.data[i], s.data[j] = s.data[j], s.data[i]
}
