package slices

import (
	"fmt"
	"reflect"
	"testing"
)

func ExampleFilter() {
	// Filter even numbers from a slice of integers.
	slice := []int{1, 2, 3, 4, 5, 6}
	even := Filter(slice, func(n int) bool {
		return n%2 == 0
	})

	fmt.Println(even)
	// Output: [2 4 6]
}

func ExampleFilter_notEmpty() {
	// Filter non-empty strings from a slice of strings.
	slice := []string{"", "a", "", "b", "c", ""}
	nonEmpty := Filter(slice, func(s string) bool {
		return s != ""
	})

	fmt.Println(nonEmpty)
	// Output: [a b c]
}

func TestFilter(t *testing.T) {
	// Test filtering even numbers from a slice of integers.
	slice := []int{1, 2, 3, 4, 5, 6}
	even := Filter(slice, func(n int) bool {
		return n%2 == 0
	})
	if !reflect.DeepEqual(even, []int{2, 4, 6}) {
		t.Errorf("got %v, want [2 4 6]", even)
	}

	// Test filtering non-empty strings from a slice of strings.
	sliceStr := []string{"", "a", "", "b", "c", ""}
	nonEmpty := Filter(sliceStr, func(s string) bool {
		return s != ""
	})
	if !reflect.DeepEqual(nonEmpty, []string{"a", "b", "c"}) {
		t.Errorf("got %v, want [a b c]", nonEmpty)
	}
}

func ExampleFindFunc() {
	// Find the first even number in a slice of integers.
	slice := []int{1, 2, 3, 4, 5, 6}
	even, found := FindFunc(slice, func(n int) bool {
		return n%2 == 0
	})

	fmt.Println(*even, found)
	// Output: 2 true
}

func ExampleFindFunc_notFound() {
	// Find the first even number in a slice of integers.
	slice := []int{1, 3, 5}
	even, found := FindFunc(slice, func(n int) bool {
		return n%2 == 0
	})

	fmt.Println(even, found)
	// Output: <nil> false
}
