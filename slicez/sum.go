package slicez

import "golang.org/x/exp/constraints"

// Sum all entries in a slice together.
func Sum[N constraints.Integer | constraints.Float | constraints.Complex](s []N) N {
	var sum N
	for _, n := range s {
		sum += n
	}
	return sum
}
