[![](https://godoc.org/github.com/Jille/genericz?status.svg)](https://pkg.go.dev/github.com/Jille/genericz)

# genericz - convenience made possible since generics

The main package only contains `Min(a ...T) T`, `Max(a ...T) T` and `Ternary(cond bool, a, b T) T`.

The packages below are intended to contain additions to [golang.org/x/exp/slices](https://pkg.go.dev/golang.org/x/exp/slices) and [golang.org/x/exp/maps](https://pkg.go.dev/golang.org/x/exp/maps) and thus won't be the full set that you need. (Over time there will be overlap as we won't remove methods as it would break backwards compatibility.)

## mapz

[![](https://godoc.org/github.com/Jille/genericz/mapz?status.svg)](https://pkg.go.dev/github.com/Jille/genericz/mapz)

The mapz package contains:

* `KeysSorted(m map[K]V) []K` and `ValuesSorted(m map[K]V) []V`
* `MinKey(m map[K]V) K` and `MaxKey(m map[K]V) K`
* `DeleteWithLock(l sync.Locker, m map[K]V, key K)` and `StoreWithLock(l sync.Locker, m map[K]V, key K, value V)`

The `SyncMap` is a type-safe `sync.Map`.

The `MutexMap` is a variant of `SyncMap` that uses a mutex and a regular map.

The `AppendMap` is an append-only map perfect for caching values that never change. It slightly cheaper than a `sync.Map` because values can't change.

## slicez

[![](https://godoc.org/github.com/Jille/genericz/slicez?status.svg)](https://pkg.go.dev/github.com/Jille/genericz/slicez)

The slicez packages contains Diff, Filter, Map, Unique, Concat and Sum.

## orderedobject

[![](https://godoc.org/github.com/Jille/genericz/orderedobject?status.svg)](https://pkg.go.dev/github.com/Jille/genericz/orderedobject)

The orderedobject allows you to decode a JSON dict while preserving order. Can be used with most json encoders/decoders.
