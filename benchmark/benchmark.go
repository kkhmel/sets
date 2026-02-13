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
	"math/rand/v2"

	"github.com/kkhmel/sets"
)

const (
	Size10   = 10
	Size100  = 100
	Size1K   = 1000
	Size10K  = 10_000
	Size100K = 100_000
	Size1M   = 1_000_000
)

var Sizes = [...]int{
	Size10,
	Size100,
	Size1K,
	Size10K,
	Size100K,
	Size1M,
}

func NewRandSet(size int) sets.Set[int] {
	return sets.FromSlice(NewRandSlice(size))
}

func NewRandSlice(size int) []int {
	result := make([]int, size)
	for i := range result {
		result[i] = rand.Int()
	}
	return result
}
