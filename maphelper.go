package maphelper

import (
	"github.com/solsw/generichelper"
)

// Tuples returns a slice containing the key/element pairs
// (in the form of [generichelper.Tuple2]) from 'm'.
// If 'm' is nil, nil is returned.
func Tuples[K comparable, E any](m map[K]E) []generichelper.Tuple2[K, E] {
	if m == nil {
		return nil
	}
	tt := make([]generichelper.Tuple2[K, E], 0, len(m))
	for k, e := range m {
		tt = append(tt, generichelper.Tuple2[K, E]{Item1: k, Item2: e})
	}
	return tt
}

// NewFromTuples creates a map from the slice of key/element pairs
// (in the form of [generichelper.Tuple2]).
// If 'tt' contains tuples with duplicate keys, the last one wins.
// If 'tt' is nil, nil is returned.
func NewFromTuples[K comparable, E any](tt []generichelper.Tuple2[K, E]) map[K]E {
	if tt == nil {
		return nil
	}
	m := make(map[K]E, len(tt))
	for _, t := range tt {
		m[t.Item1] = t.Item2
	}
	return m
}
