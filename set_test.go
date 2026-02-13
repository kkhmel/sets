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
	"fmt"
	"slices"
	"strconv"
	"testing"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name string
		n    int
	}{
		{
			name: "zero",
			n:    0,
		},
		{
			name: "positive",
			n:    10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := New[int](tt.n)
			if got == nil {
				t.Error("New() returned nil")
			}
			if len(got) != 0 {
				t.Errorf("New() returned non-empty set: %v", got)
			}
		})
	}
}

func TestNewPanic(t *testing.T) {
	capacity := -5
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("New() should panic when capacity is negative, but did not panic")
		}
	}()

	New[int](capacity)
}

func TestFrom(t *testing.T) {
	tests := []struct {
		name string
		vals []int
		want Set[int]
	}{
		{
			name: "no vals",
			vals: nil,
			want: map[int]struct{}{},
		},
		{
			name: "one element",
			vals: []int{1},
			want: map[int]struct{}{1: {}},
		},
		{
			name: "multiple elements",
			vals: []int{1, 2, 3},
			want: map[int]struct{}{1: {}, 2: {}, 3: {}},
		},
		{
			name: "duplicate elements",
			vals: []int{1, 2, 2, 3, 3, 3},
			want: map[int]struct{}{1: {}, 2: {}, 3: {}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := From(tt.vals...)
			if !Equal(got, tt.want) {
				t.Errorf("\nwant: %v\ngot : %v", tt.want, got)
			}
		})
	}
}

func TestFromSlice(t *testing.T) {
	tests := []struct {
		name  string
		slice []int
		want  Set[int]
	}{
		{
			name:  "no elements",
			slice: nil,
			want:  map[int]struct{}{},
		},
		{
			name:  "one element",
			slice: []int{1},
			want:  map[int]struct{}{1: {}},
		},
		{
			name:  "multiple elements",
			slice: []int{1, 2, 3},
			want:  map[int]struct{}{1: {}, 2: {}, 3: {}},
		},
		{
			name:  "duplicate elements",
			slice: []int{1, 2, 2, 3, 3, 3},
			want:  map[int]struct{}{1: {}, 2: {}, 3: {}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FromSlice(tt.slice)
			if !Equal(got, tt.want) {
				t.Errorf("\nwant: %v\ngot : %v", tt.want, got)
			}
		})
	}
}

func TestFromSliceFunc(t *testing.T) {
	tests := []struct {
		name  string
		slice []int
		f     func(e int) string
		want  Set[string]
	}{
		{
			name:  "nil slice",
			slice: nil,
			f:     func(s int) string { return strconv.Itoa(s) },
			want:  From[string](),
		},
		{
			name:  "empty slice",
			slice: []int{},
			f:     func(s int) string { return strconv.Itoa(s) },
			want:  From[string](),
		},
		{
			name:  "single element",
			slice: []int{42},
			f:     func(s int) string { return strconv.Itoa(s) },
			want:  From("42"),
		},
		{
			name:  "multiple elements",
			slice: []int{1, 2, 3},
			f:     func(s int) string { return strconv.Itoa(s) },
			want:  From("1", "2", "3"),
		},
		{
			name:  "duplicate results after transform",
			slice: []int{1, 2, 2, 3, 3, 3},
			f:     func(s int) string { return strconv.Itoa(s) },
			want:  From("1", "2", "3"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FromSliceFunc(tt.slice, tt.f)
			if !Equal(got, tt.want) {
				t.Errorf("\nwant: %v\ngot : %v", tt.want, got)
			}
		})
	}
}

func TestSet_String(t *testing.T) {
	tests := []struct {
		name string
		set  Set[int]
		want string
	}{
		{
			name: "nil set",
			set:  New[int](0),
			want: "{}",
		},
		{
			name: "empty set",
			set:  New[int](0),
			want: "{}",
		},
		{
			name: "single element",
			set:  From(1),
			want: "{1}",
		},
		{
			name: "multiple elements",
			set:  From(1, 2, 3),
			want: "{1, 2, 3}",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.set.String()
			if got != tt.want {
				t.Errorf("\nwant: %v\ngot : %v", tt.want, got)
			}
		})
	}
}

func TestClone(t *testing.T) {
	tests := []struct {
		name string
		s    Set[int]
		want Set[int]
	}{
		{
			name: "nil",
			s:    nil,
			want: nil,
		},
		{
			name: "empty",
			s:    New[int](0),
			want: New[int](0),
		},
		{
			name: "single element",
			s:    From(42),
			want: From(42),
		},
		{
			name: "multiple elements",
			s:    From(1, 2, 3, 4, 5),
			want: From(1, 2, 3, 4, 5),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Clone(tt.s)
			if !Equal(got, tt.s) {
				t.Errorf("\nwant: %v\ngot : %v", tt.s, got)
			}
			if tt.want != nil {
				// Check if a new set was created
				element := -999
				if Insert(got, element); Contains(tt.s, element) {
					t.Errorf("Grow() should have created a new set but returned the original")
				}
			}
		})
	}
}

func TestGrow(t *testing.T) {
	tests := []struct {
		name string
		s    Set[int]
		n    int
	}{
		{
			name: "nil set",
			s:    New[int](0),
			n:    10,
		},
		{
			name: "empty set",
			s:    New[int](0),
			n:    10,
		},
		{
			name: "n is zero",
			s:    From(1, 2, 3),
			n:    0,
		},
		{
			name: "non-empty set positive n",
			s:    From(1, 2, 3),
			n:    5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Grow(tt.s, tt.n)
			if !Equal(got, tt.s) {
				t.Errorf("\nwant: %v\ngot : %v", tt.s, got)
			}

			element := -999
			if Insert(got, element); Contains(tt.s, element) {
				t.Errorf("Grow() should have created a new subset but returned the original")
			}
		})
	}
}

func TestGrowPanic(t *testing.T) {
	s := From(1, 2, 3)
	n := -5
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Grow() should panic when n is negative, but did not panic")
		}
	}()

	Grow(s, n)
}

func TestCopy(t *testing.T) {
	tests := []struct {
		name string
		dst  Set[int]
		src  Set[int]
		want Set[int]
	}{
		{
			name: "empty dst",
			dst:  New[int](0),
			src:  From(1, 2, 3),
			want: From(1, 2, 3),
		},
		{
			name: "empty src",
			dst:  From(1, 2, 3),
			src:  New[int](0),
			want: From(1, 2, 3),
		},
		{
			name: "overlapping",
			dst:  From(1, 2, 3),
			src:  From(2, 3, 4),
			want: From(1, 2, 3, 4),
		},
		{
			name: "disjoint",
			dst:  From(1, 2),
			src:  From(3, 4),
			want: From(1, 2, 3, 4),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Copy(tt.dst, tt.src)
			if !Equal(tt.dst, tt.want) {
				t.Errorf("\nwant: %v\ngot : %v", tt.want, tt.dst)
			}
		})
	}
}

func TestToSlice(t *testing.T) {
	tests := []struct {
		name string
		s    Set[int]
		want []int
	}{
		{
			name: "nil set",
			s:    nil,
			want: nil,
		},
		{
			name: "empty set",
			s:    New[int](0),
			want: []int{},
		},
		{
			name: "single element",
			s:    From(42),
			want: []int{42},
		},
		{
			name: "multiple elements",
			s:    From(1, 2, 3),
			want: []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ToSlice(tt.s)
			slices.Sort(got)
			if !slices.Equal(got, tt.want) {
				t.Errorf("\nwant: %v\ngot : %v", tt.want, got)
			}
		})
	}
}

func TestToSliceFunc(t *testing.T) {
	tests := []struct {
		name string
		s    Set[int]
		f    func(int) string
		want []string
	}{
		{
			name: "nil set",
			s:    nil,
			f:    func(i int) string { return fmt.Sprintf("%d", i) },
			want: nil,
		},
		{
			name: "empty set",
			s:    New[int](0),
			f:    func(i int) string { return fmt.Sprintf("%d", i) },
			want: []string{},
		},
		{
			name: "single element",
			s:    From(42),
			f:    func(i int) string { return fmt.Sprintf("v%d", i) },
			want: []string{"v42"},
		},
		{
			name: "multiple elements",
			s:    From(1, 2, 3),
			f:    func(i int) string { return fmt.Sprintf("v%d", i) },
			want: []string{"v1", "v2", "v3"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ToSliceFunc(tt.s, tt.f)
			slices.Sort(got)
			slices.Sort(tt.want)
			if !slices.Equal(got, tt.want) {
				t.Errorf("\nwant: %v\ngot : %v", tt.want, got)
			}
		})
	}
}
