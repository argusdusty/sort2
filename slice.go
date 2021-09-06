// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sort2

type genericSlice[T any] struct {
    slice []T
    less func(i, j int) bool
}

func (s genericSlice[T]) Less(i, j int) bool { return s.less(i, j) }
func (s genericSlice[T]) Swap(i, j int) { s.slice[i], s.slice[j] = s.slice[j], s.slice[i] }

// Slice sorts the provided slice given the provided less function.
//
// The sort is not guaranteed to be stable. For a stable sort, use
// SliceStable.
func Slice[T any](slice []T, less func(i, j int) bool) {
	length := len(slice)
	quickSort(genericSlice[T]{slice, less}, 0, length, maxDepth(length))
}

// SliceStable sorts the provided slice given the provided less
// function while keeping the original order of equal elements.
func SliceStable[T any](slice []T, less func(i, j int) bool) {
	stable(genericSlice[T]{slice, less}, len(slice))
}

// SliceIsSorted tests whether a slice is sorted.
func SliceIsSorted[T any](slice []T, less func(i, j int) bool) bool {
	n := len(slice)
	for i := n - 1; i > 0; i-- {
		if less(i, i-1) {
			return false
		}
	}
	return true
}

type orderedSlice[T Ordered] []T

func (s orderedSlice[T]) Less(i, j int) bool { return s[i] < s[j] }
func (s orderedSlice[T]) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

// SliceSortable sorts the provided slice with Sortable elements given the
// provided less function.
//
// The sort is not guaranteed to be stable. For a stable sort, use
// SliceStableSortable.
func SliceOrdered[T Ordered](slice []T) {
	length := len(slice)
	quickSort(orderedSlice[T](slice), 0, length, maxDepth(length))
}

// SliceStableSortable sorts the provided slice with Sortable elements
// given the provided less function while keeping the original order of
// equal elements.
func SliceStableOrdered[T Ordered](slice []T) {
	stable(orderedSlice[T](slice), len(slice))
}

// SliceIsSortedSortable tests whether a slice with Sortable elements is
// sorted.
func SliceIsSortedOrdered[T Ordered](slice []T) bool {
	n := len(slice)
	for i := n - 1; i > 0; i-- {
		if slice[i] < slice[i-1] {
			return false
		}
	}
	return true
}