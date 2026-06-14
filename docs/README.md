# maphelper
[![Go Reference](https://pkg.go.dev/badge/github.com/solsw/maphelper.svg)](https://pkg.go.dev/github.com/solsw/maphelper)
[![GitHub](https://img.shields.io/badge/github--green?logo=github)](https://github.com/solsw/maphelper)

Package `maphelper` contains [map](https://go.dev/ref/spec#Map_types)-related
helpers for converting between a map and a slice of key/element pairs.

For obtaining a map's keys or values as slices, use the standard library
[`maps`](https://pkg.go.dev/maps) and [`slices`](https://pkg.go.dev/slices)
packages directly (e.g. `slices.Collect(maps.Keys(m))`); `maphelper` provides
only the tuple-conversion helpers that have no standard-library equivalent.

## Installation

```sh
go get github.com/solsw/maphelper
```

## Functions

### `Tuples`

```go
func Tuples[K comparable, E any](m map[K]E) []generichelper.Tuple2[K, E]
```

Returns a slice containing the key/element pairs (as
[`generichelper.Tuple2`](https://pkg.go.dev/github.com/solsw/generichelper#Tuple2))
from `m`. The order of the pairs is unspecified, mirroring map iteration order.
If `m` is `nil`, `nil` is returned.

### `NewFromTuples`

```go
func NewFromTuples[K comparable, E any](tt []generichelper.Tuple2[K, E]) map[K]E
```

Creates a map from the slice of key/element pairs. If `tt` contains tuples with
duplicate keys, the last one wins. If `tt` is `nil`, `nil` is returned.

## Example

```go
package main

import (
	"fmt"

	"github.com/solsw/generichelper"
	"github.com/solsw/maphelper"
)

func main() {
	m := map[int]string{1: "one", 2: "two", 3: "three"}

	// map -> slice of key/element pairs
	tt := maphelper.Tuples(m)

	// slice of key/element pairs -> map
	m2 := maphelper.NewFromTuples(tt)

	fmt.Println(len(tt), len(m2)) // 3 3

	_ = generichelper.Tuple2[int, string]{Item1: 1, Item2: "one"}
}
```

## License

[MIT](../LICENSE)
