package genericz

import "golang.org/x/exp/constraints"

// Min returns the lowest value given.
func Min[T constraints.Ordered](a ...T) T {
	ret := a[0]
	for _, e := range a[1:] {
		if e < ret {
			ret = e
		}
	}
	return ret
}

// Max returns the highest value given.
func Max[T constraints.Ordered](a ...T) T {
	ret := a[0]
	for _, e := range a[1:] {
		if e > ret {
			ret = e
		}
	}
	return ret
}
