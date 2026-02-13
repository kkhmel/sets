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

package sets

// Contains reports whether v is present in s.
//
// Time complexity: O(1). Space complexity: O(1).
func Contains[S ~map[E]struct{}, E comparable](s S, v E) bool {
	_, ok := s[v]
	return ok
}

// ContainsAny reports whether at least one of the specified elements is present in s.
//
// Time complexity: O(len(v)). Space complexity: O(1).
func ContainsAny[S ~map[E]struct{}, E comparable](s S, v ...E) bool {
	if len(s) == 0 {
		return false
	}
	for _, e := range v {
		if Contains(s, e) {
			return true
		}
	}
	return false
}

// ContainsAll reports whether all specified elements are present in s.
//
// Time complexity: O(len(v)). Space complexity: O(1).
func ContainsAll[S ~map[E]struct{}, E comparable](s S, v ...E) bool {
	if len(v) == 0 {
		return true
	}
	if len(s) == 0 {
		return false
	}
	for _, e := range v {
		if !Contains(s, e) {
			return false
		}
	}
	return true
}

// Some reports whether at least one element e of s satisfies f(e).
//
// Time complexity: O(len(s)). Space complexity: O(1).
func Some[S ~map[E]struct{}, E comparable](s S, f func(E) bool) bool {
	for e := range s {
		if f(e) {
			return true
		}
	}
	return false
}

// Every reports whether all elements e of s satisfy f(e).
//
// Time complexity: O(len(s)). Space complexity: O(1).
func Every[S ~map[E]struct{}, E comparable](s S, f func(E) bool) bool {
	for e := range s {
		if !f(e) {
			return false
		}
	}
	return true
}

// Equal reports whether two sets contain the same elements.
// Elements are compared using ==.
//
// Time complexity: O(len(s1)). Space complexity: O(1).
func Equal[S1, S2 ~map[E]struct{}, E comparable](s1 S1, s2 S2) bool {
	if len(s1) != len(s2) {
		return false
	}
	for e := range s1 {
		if _, ok := s2[e]; !ok {
			return false
		}
	}
	return true
}

// Overlaps reports whether s1 and s2 have any element in common.
//
// Time complexity: O(min(len(s1), len(s2))). Space complexity: O(1).
func Overlaps[S ~map[E]struct{}, E comparable](s1 S, s2 S) bool {
	if len(s1) == 0 || len(s2) == 0 {
		return false
	}
	if len(s1) > len(s2) {
		s1, s2 = s2, s1
	}
	for e := range s1 {
		if _, ok := s2[e]; ok {
			return true
		}
	}
	return false
}

// Subset reports whether all elements of subset are also in superset.
//
// Time complexity: O(len(subset)). Space complexity: O(1).
func Subset[S1, S2 ~map[E]struct{}, E comparable](subset S1, superset S2) bool {
	if len(subset) == 0 {
		return true
	}
	if len(superset) == 0 {
		return false
	}
	for e := range subset {
		if _, ok := superset[e]; !ok {
			return false
		}
	}
	return true
}

// ProperSubset reports whether subset is a proper subset of superset,
// i.e. all elements of subset are in superset and the sets are not equal.
//
// Time complexity: O(len(subset)). Space complexity: O(1).
func ProperSubset[S1, S2 ~map[E]struct{}, E comparable](subset S1, superset S2) bool {
	if len(subset) >= len(superset) {
		return false
	}
	return Subset(subset, superset)
}
