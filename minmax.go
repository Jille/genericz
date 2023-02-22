package genericz

import "golang.org/x/exp/constraints"

func Min[T constraints.Ordered](a ...T) T {
	ret := a[0]
	for _, e := range a[1:] {
		if e < ret {
			ret = e
		}
	}
	return ret
}

func Max[T constraints.Ordered](a ...T) T {
	ret := a[0]
	for _, e := range a[1:] {
		if e > ret {
			ret = e
		}
	}
	return ret
}
