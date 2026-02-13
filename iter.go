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

import (
	"iter"
)

// All returns an iterator over elements from s.
// The iteration order is not specified and is not guaranteed to be the same from one call to the next.
//
// Creation: O(1) time, O(1) space.
// Iteration: O(len(s)) time, O(1) space.
func All[S ~map[E]struct{}, E comparable](s S) iter.Seq[E] {
	return func(yield func(E) bool) {
		for e := range s {
			if !yield(e) {
				return
			}
		}
	}
}

// InsertSeq inserts the elements from seq to s.
// Duplicate elements are ignored.
//
// Time complexity: O(n). Space complexity: O(n). n is the number of seq elements.
func InsertSeq[S ~map[E]struct{}, E comparable](s S, seq iter.Seq[E]) {
	for e := range seq {
		s[e] = struct{}{}
	}
}

// Collect collects elements from seq into a new set and returns it.
//
// Time complexity: O(n). Space complexity: O(n). n is the number of seq elements.
func Collect[E comparable](seq iter.Seq[E]) Set[E] {
	s := New[E](0)
	InsertSeq(s, seq)
	return s
}

// Chunk returns an iterator over consecutive subsets of up to n elements of s.
// All but the last subset will have size n.
// If s is empty, the sequence is empty: there is no empty set in the sequence.
// Chunk panics if n is less than 1.
//
// Creation: O(1) time, O(1) space.
// Iteration: O(len(s)) time, O(n) space.
func Chunk[S ~map[E]struct{}, E comparable](s S, n int) iter.Seq[S] {
	if n < 1 {
		panic("cannot be less than 1")
	}

	return func(yield func(S) bool) {
		chunk := make(S, n)
		cnt := 0

		for e := range s {
			chunk[e] = struct{}{}
			cnt++

			if cnt == n {
				if !yield(chunk) {
					return
				}
				chunk = make(S, n)
				cnt = 0
			}
		}

		// Yield remaining elements if any
		if cnt > 0 {
			yield(chunk)
		}
	}
}

// UnionSeq returns an iterator over all unique elements from all sets in the sequence.
// It eliminates duplicates, so each element is yielded only once even if it appears in multiple sets.
//
// Creation: O(1) time, O(1) space.
// Iteration: O(N) time, O(U) space. N is the total number of elements across all sets.
// U is the number of unique elements across all sets.
func UnionSeq[S ~map[E]struct{}, E comparable](seq iter.Seq[S]) iter.Seq[E] {
	return func(yield func(E) bool) {
		seen := make(map[E]struct{})
		for s := range seq {
			for e := range s {
				if _, ok := seen[e]; !ok {
					seen[e] = struct{}{}
					if !yield(e) {
						return
					}
				}
			}
		}
	}
}
