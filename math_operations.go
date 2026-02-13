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

// Pair represents an ordered pair of elements, used as the result type for CartesianProduct.
type Pair[T, U comparable] struct {
	First  T
	Second U
}

// Union returns a new set containing all unique elements from all provided sets.
// If no sets are provided, returns an empty set.
//
// Time complexity: O(N). Space complexity: O(N). N is the sum of all set sizes.
func Union[S ~map[E]struct{}, E comparable](sets ...S) Set[E] {
	maxSize := 0
	for _, s := range sets {
		maxSize += len(s)
	}

	r := New[E](maxSize)
	for _, s := range sets {
		for e := range s {
			r[e] = struct{}{}
		}
	}

	return r
}

// Intersection returns a new set containing only elements that are present in all sets.
// If no sets are provided, returns an empty set.
//
// Time complexity: O(N). Space complexity: O(min). N is the sum of all set sizes. min is the size of the smallest set.
func Intersection[S ~map[E]struct{}, E comparable](sets ...S) Set[E] {
	if len(sets) == 0 {
		return New[E](0)
	}
	if len(sets) == 1 {
		r := New[E](len(sets[0]))
		for e := range sets[0] {
			r[e] = struct{}{}
		}
		return r
	}

	smallest := 0
	for i := 1; i < len(sets); i++ {
		if len(sets[i]) < len(sets[smallest]) {
			smallest = i
		}
	}

	r := New[E](len(sets[smallest]))
elementsLoop:
	for e := range sets[smallest] {
		for i, s := range sets {
			if i == smallest {
				continue
			}
			if _, ok := s[e]; !ok {
				continue elementsLoop
			}
		}
		r[e] = struct{}{}
	}

	return r
}

// Difference returns a new set containing elements that are in the minuend but not in any of the subtrahends.
//
// Time complexity: O(len(minuend) + S). Space complexity: O(len(minuend)). S is the sum of subtrahend sizes.
func Difference[S ~map[E]struct{}, E comparable](minuend S, subtrahends ...S) Set[E] {
	if len(minuend) == 0 {
		return New[E](0)
	}
	if len(subtrahends) == 0 {
		r := New[E](len(minuend))
		for e := range minuend {
			r[e] = struct{}{}
		}
		return r
	}
	if len(subtrahends) == 1 {
		r := New[E](len(minuend))
		for e := range minuend {
			if _, ok := subtrahends[0][e]; !ok {
				r[e] = struct{}{}
			}
		}
		return r
	}

	combinedMaxSize := 0
	for _, subtrahend := range subtrahends {
		combinedMaxSize += len(subtrahend)
	}

	combined := New[E](combinedMaxSize)
	for _, s := range subtrahends {
		for e := range s {
			combined[e] = struct{}{}
		}
	}

	r := New[E](len(minuend))
	for e := range minuend {
		if _, ok := combined[e]; !ok {
			r[e] = struct{}{}
		}
	}

	return r
}

// SymmetricDifference returns a new set containing elements that belong to
// an odd number of the provided sets (i.e., the n-ary XOR of the sets).
// If no sets are provided, returns an empty set.
//
// Time complexity: O(N). Space complexity: O(N). N is the sum of all set sizes.
func SymmetricDifference[S ~map[E]struct{}, E comparable](sets ...S) Set[E] {
	if len(sets) == 0 {
		return New[E](0)
	}
	if len(sets) == 1 {
		r := New[E](len(sets[0]))
		for e := range sets[0] {
			r[e] = struct{}{}
		}
		return r
	}

	maxSize := 0
	for _, s := range sets {
		maxSize += len(s)
	}

	r := New[E](maxSize)
	for _, s := range sets {
		for e := range s {
			if _, ok := r[e]; ok {
				delete(r, e)
			} else {
				r[e] = struct{}{}
			}
		}
	}

	return r
}

// CartesianProduct returns a new set containing all ordered pairs (e1, e2) where e1 is from set1 and e2 is from set2.
//
// Time complexity: O(len(set1) * len(set2)). Space complexity: O(len(set1) * len(set2)).
func CartesianProduct[S1 ~map[E1]struct{}, S2 ~map[E2]struct{}, E1, E2 comparable](set1 S1, set2 S2) Set[Pair[E1, E2]] {
	if len(set1) == 0 || len(set2) == 0 {
		return New[Pair[E1, E2]](0)
	}
	r := New[Pair[E1, E2]](len(set1) * len(set2))
	for e1 := range set1 {
		for e2 := range set2 {
			r[Pair[E1, E2]{First: e1, Second: e2}] = struct{}{}
		}
	}
	return r
}
