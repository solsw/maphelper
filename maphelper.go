package maphelper

import (
	"github.com/solsw/generichelper"
)

// Keys returns a slice containing the keys from 'm'.
// If 'm' is nil, nil is returned.
func Keys[K comparable, E any](m map[K]E) []K {
	if m == nil {
		return nil
	}
	kk := make([]K, 0, len(m))
	for k := range m {
		kk = append(kk, k)
	}
	return kk
}

// Elements returns a slice containing the elements from 'm'.
// If 'm' is nil, nil is returned.
func Elements[K comparable, E any](m map[K]E) []E {
	if m == nil {
		return nil
	}
	ee := make([]E, 0, len(m))
	for _, e := range m {
		ee = append(ee, e)
	}
	return ee
}

// Tuples returns a slice containing the key/element pairs from 'm'.
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

// NewFromTuples creates a map from the slice of key/element pairs.
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
