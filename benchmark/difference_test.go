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

package benchmark

import (
	"testing"

	"github.com/kkhmel/sets"
)

func BenchmarkDifference_equal(b *testing.B) {
	set1 := NewRandSet(Size1K)
	set2 := sets.Clone(set1)
	for b.Loop() {
		sets.Difference(set1, set2)
	}
}

func BenchmarkDifference_halfOverlap(b *testing.B) {
	set1 := NewRandSet(Size1K)
	set2 := sets.New[int](Size1K)
	cnt := 0
	for e := range set1 {
		if cnt < Size1K/2 {
			sets.Insert(set2, e)
		} else {
			sets.Insert(set2, -e)
		}
		cnt++
	}
	for b.Loop() {
		sets.Difference(set1, set2)
	}
}

func BenchmarkDifference_disjoint(b *testing.B) {
	set1 := NewRandSet(Size1K)
	set2 := sets.Map(set1, func(e int) int { return -e })
	for b.Loop() {
		sets.Difference(set1, set2)
	}
}
