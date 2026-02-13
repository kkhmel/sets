// Copyright (c) 2026 kkhmel
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

// Package sets provides a high-performance, idiomatic Set library for Go that follows
// the same design principles as the standard library's maps and slices packages.
//
// The library provides ~30 functions covering:
//   - Set theory operations (Union, Intersection, Difference, etc.)
//   - Functional programming (Map, Filter, Reduce)
//   - Predicates and comparisons (Equal, Subset, Contains, etc.)
//   - Iterator support
//
// Zero dependencies. Pure Go standard library only.
package sets

import (
	"fmt"
	"maps"
	"sort"
	"strings"
)

// Set is an implementation of a mathematical set -- a well-defined, unordered collection of distinct objects.
// Since Set[E] is just a type definition for map[E]struct{}, you can use native Go map operations to work with sets directly.
// Additionally, all functions from the standard maps package are compatible with sets.
type Set[E comparable] map[E]struct{}

// New creates a new Set with the specified initial capacity.
//
// Time complexity: O(1). Space complexity: O(n). n is the passed capacity.
func New[E comparable](capacity int) Set[E] {
	if capacity < 0 {
		panic("cannot be negative")
	}
	return make(Set[E], capacity)
}

// From creates a new Set containing the provided vals.
// See also FromSlice.
//
// Time complexity: O(len(vals)). Space complexity: O(len(vals)).
func From[E comparable](vals ...E) Set[E] {
	return FromSlice(vals)
}

// FromSlice creates a new Set from an existing slice.
//
// Time complexity: O(len(slice)). Space complexity: O(len(slice)).
func FromSlice[E comparable](slice []E) Set[E] {
	s := New[E](len(slice))
	Insert(s, slice...)
	return s
}

// FromSliceFunc creates a new Set by applying f to each element of the slice.
//
// Time complexity: O(len(slice)). Space complexity: O(len(slice)).
func FromSliceFunc[A any, E comparable](slice []A, f func(A) E) Set[E] {
	s := New[E](len(slice))
	for _, a := range slice {
		Insert(s, f(a))
	}
	return s
}

// String returns a string representation of the set in the format "{elem1, elem2, ...}".
// Elements are sorted by their string representation for consistent output.
//
// Time complexity: O(len(s)). Space complexity: O(len(s)).
func (s Set[E]) String() string {
	if len(s) == 0 {
		return "{}"
	}
	elements := make([]string, 0, len(s))
	for e := range s {
		elements = append(elements, fmt.Sprintf("%v", e))
	}
	sort.Strings(elements)
	return "{" + strings.Join(elements, ", ") + "}"
}

// Clone returns a shallow copy of s. The new set is allocated with
// a capacity hint equal to len(s), which may help reclaim excess memory
// held by the original map (e.g. after many deletions).
//
// Time complexity: O(len(s)). Space complexity: O(len(s)).
func Clone[S ~map[E]struct{}, E comparable](s S) S {
	if s == nil {
		return nil
	}
	clone := make(S, len(s))
	Copy(clone, s)
	return clone
}

// Copy copies all elements in src, adding them to dst.
//
// Time complexity: O(len(src)). Space complexity: O(len(src)).
func Copy[S1 ~map[E]struct{}, S2 ~map[E]struct{}, E comparable](dst S1, src S2) {
	for e := range src {
		dst[e] = struct{}{}
	}
}

// Grow returns a new set with capacity increased to guarantee space for n more elements without another allocation.
// After Grow(n), at least n elements can be appended to the set without another allocation.
// If n is negative or too large to allocate the memory, Grow panics.
//
// Time complexity: O(len(s)). Space complexity: O(len(s) + n).
func Grow[S ~map[E]struct{}, E comparable](s S, n int) S {
	if n < 0 {
		panic("cannot be negative")
	}
	newSet := make(S, len(s)+n)
	maps.Copy(newSet, s)
	return newSet
}

// ToSlice returns all elements of the set as a slice.
// The order of elements is non-deterministic due to map iteration.
//
// Time complexity: O(len(s)). Space complexity: O(len(s)).
func ToSlice[S ~map[E]struct{}, E comparable](s S) []E {
	if s == nil {
		return nil
	}
	r := make([]E, 0, len(s))
	for e := range s {
		r = append(r, e)
	}
	return r
}

// ToSliceFunc returns a slice created by applying f to each element of s.
// The order of elements is non-deterministic due to map iteration.
//
// Time complexity: O(len(s)). Space complexity: O(len(s)).
func ToSliceFunc[S ~map[E]struct{}, E comparable, A any](s S, f func(E) A) []A {
	if s == nil {
		return nil
	}
	r := make([]A, 0, len(s))
	for e := range s {
		r = append(r, f(e))
	}
	return r
}
