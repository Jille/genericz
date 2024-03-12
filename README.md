[![](https://godoc.org/github.com/Jille/genericz?status.svg)](https://pkg.go.dev/github.com/Jille/genericz)

# genericz - convenience made possible since generics

The main package only contains `Min(a ...T) T` and `Max(a ...T) T`.

The packages below are intended to contain additions to [golang.org/x/exp/slices](https://pkg.go.dev/golang.org/x/exp/slices) and [golang.org/x/exp/maps](https://pkg.go.dev/golang.org/x/exp/maps) and thus won't be the full set that you need. (Over time there will be overlap as we won't remove methods as it would break backwards compatibility.)

## mapz

[![](https://godoc.org/github.com/Jille/genericz/mapz?status.svg)](https://pkg.go.dev/github.com/Jille/genericz/mapz)

The mapz package contains:

* `KeysSorted(m map[K]V) []K` and `ValuesSorted(m map[K]V) []V`
* `MinKey(m map[K]V) K` and `MaxKey(m map[K]V) K`
* `DeleteWithLock(l sync.Locker, m map[K]V, key K)` and `StoreWithLock(l sync.Locker, m map[K]V, key K, value V)`

The `SyncMap` is a type-safe `sync.Map`.

The `MutexMap` is a variant of `SyncMap` that uses a mutex and a regular map.

## slicez

[![](https://godoc.org/github.com/Jille/genericz/slicez?status.svg)](https://pkg.go.dev/github.com/Jille/genericz/slicez)

The slicez packages contains Diff, Filter, Unique, Concat and Sum.
