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

// Map returns a new set containing the results of applying a transformation function to each element in the set.
// The resulting set may have fewer elements than the input set if the function is not injective.
//
// Time complexity: O(len(s)). Space complexity: O(len(s)).
func Map[S ~map[From]struct{}, From, To comparable](s S, f func(From) To) Set[To] {
	r := New[To](len(s))
	for e := range s {
		r[f(e)] = struct{}{}
	}
	return r
}

// Filter returns a new set containing only the elements from the input set that satisfy the predicate function f.
// The resulting set is pre-allocated with the capacity of the input set for efficiency.
// Use Clone() on the result if memory optimization is needed.
//
// Time complexity: O(len(s)). Space complexity: O(len(s)).
func Filter[S ~map[E]struct{}, E comparable](s S, f func(E) bool) Set[E] {
	r := New[E](len(s))
	for e := range s {
		if f(e) {
			r[e] = struct{}{}
		}
	}
	return r
}

// Reduce returns the result of applying a reduction function f to each element in the input set,
// accumulating a result starting from the initial value acc.
// Note: The order of reduction is non-deterministic due to map iteration.
//
// Time complexity: O(len(s)). Space complexity: O(1).
func Reduce[S ~map[E]struct{}, E comparable, A any](s S, acc A, f func(A, E) A) A {
	for e := range s {
		acc = f(acc, e)
	}
	return acc
}
