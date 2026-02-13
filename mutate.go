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

import "maps"

// Insert inserts the given elements into the set.
// Elements already present in the set are ignored.
//
// Time complexity: O(len(v)). Space complexity: O(1).
func Insert[S ~map[E]struct{}, E comparable](s S, v ...E) {
	for _, e := range v {
		s[e] = struct{}{}
	}
}

// Delete deletes the specified elements from the set.
// If s is nil or there is no such element, delete is a no-op.
//
// Time complexity: O(len(v)). Space complexity: O(1).
func Delete[S ~map[E]struct{}, E comparable](s S, v ...E) {
	for _, e := range v {
		delete(s, e)
	}
}

// DeleteFunc deletes any elements from s for which del returns true.
//
// Time complexity: O(len(s)). Space complexity: O(1).
func DeleteFunc[S ~map[E]struct{}, E comparable](s S, del func(E) bool) {
	maps.DeleteFunc(s, func(e E, _ struct{}) bool { return del(e) })
}

// Replace replaces old with new in the set. If old is not present, Replace is a no-op.
//
// Time complexity: O(1). Space complexity: O(1).
func Replace[S ~map[E]struct{}, E comparable](s S, old, new E) { //nolint:revive // 'new' follows stdlib pattern (see strings.Replace)
	if _, ok := s[old]; !ok {
		return
	}
	delete(s, old)
	s[new] = struct{}{}
}

// ReplaceFunc replaces each element e in s with f(e).
// Since f may map multiple elements to the same value, the resulting set may be smaller.
//
// Time complexity: O(len(s)). Space complexity: O(1).
func ReplaceFunc[S ~map[E]struct{}, E comparable](s S, f func(E) E) {
	keys := make([]E, 0, len(s))
	for e := range s {
		keys = append(keys, e)
	}
	clear(s)
	for _, e := range keys {
		s[f(e)] = struct{}{}
	}
}
