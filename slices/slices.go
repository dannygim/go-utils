// Package slices defines additional various functions useful
// not included in https://pkg.go.dev/slices.
package slices

import (
	std_slices "slices"
)

// Filter returns a new slice containing only the elements of the slice that satisfy the predicate.
func Filter[S ~[]E, E any](slice S, predicate func(E) bool) []E {
	var filtered S

	for _, v := range slice {
		if predicate(v) {
			filtered = append(filtered, v)
		}
	}

	return filtered
}

// FindFunc returns the first element in the slice that satisfies the predicate.
func FindFunc[S ~[]E, E any](slice S, predicate func(E) bool) (*E, bool) {
	if index := std_slices.IndexFunc(slice, predicate); index == -1 {
		return nil, false
	} else {
		return &slice[index], true
	}
}
