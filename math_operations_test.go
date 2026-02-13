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
	"testing"
)

func TestUnion(t *testing.T) {
	tests := []struct {
		name string
		sets []Set[int]
		want Set[int]
	}{
		{
			name: "no sets",
			want: New[int](0),
		},
		{
			name: "single nil set",
			sets: []Set[int]{nil},
			want: New[int](0),
		},
		{
			name: "single empty set",
			sets: []Set[int]{New[int](0)},
			want: New[int](0),
		},
		{
			name: "single set",
			sets: []Set[int]{From(1, 2, 3)},
			want: From(1, 2, 3),
		},
		{
			name: "with nil sets",
			sets: []Set[int]{nil, From(1, 2, 3), nil},
			want: From(1, 2, 3),
		},
		{
			name: "with empty sets",
			sets: []Set[int]{New[int](0), From(1, 2, 3), New[int](0)},
			want: From(1, 2, 3),
		},
		{
			name: "two sets",
			sets: []Set[int]{From(1, 2, 3), From(2, 3, 4)},
			want: From(1, 2, 3, 4),
		},
		{
			name: "multiple sets",
			sets: []Set[int]{From(1, 2, 3, 4), From(2, 3, 4, 5), From(3, 4, 5, 6)},
			want: From(1, 2, 3, 4, 5, 6),
		},
		{
			name: "multiple equal sets",
			sets: []Set[int]{From(1, 2, 3), From(1, 2, 3), From(1, 2, 3)},
			want: From(1, 2, 3),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Union(tt.sets...)
			if !Equal(got, tt.want) {
				t.Errorf("\nwant: %v\ngot : %v", tt.want, got)
			}
		})
	}
}

func TestIntersection(t *testing.T) {
	tests := []struct {
		name string
		sets []Set[int]
		want Set[int]
	}{
		{
			name: "no sets",
			want: New[int](0),
		},
		{
			name: "single nil set",
			sets: []Set[int]{nil},
			want: New[int](0),
		},
		{
			name: "single empty set",
			sets: []Set[int]{New[int](0)},
			want: New[int](0),
		},
		{
			name: "single set",
			sets: []Set[int]{From(1, 2, 3)},
			want: From(1, 2, 3),
		},
		{
			name: "with nil sets",
			sets: []Set[int]{nil, From(1, 2, 3), nil},
			want: New[int](0),
		},
		{
			name: "with empty sets",
			sets: []Set[int]{New[int](0), From(1, 2, 3), New[int](0)},
			want: New[int](0),
		},
		{
			name: "two sets",
			sets: []Set[int]{From(1, 2, 3), From(2, 3, 4)},
			want: From(2, 3),
		},
		{
			name: "multiple sets",
			sets: []Set[int]{From(1, 2, 3, 4), From(2, 3, 4), From(3, 4, 5)},
			want: From(3, 4),
		},
		{
			name: "multiple disjoint sets",
			sets: []Set[int]{From(1, 2, 3), From(4, 5, 6), From(7, 8, 9)},
			want: New[int](0),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Intersection(tt.sets...)
			if !Equal(got, tt.want) {
				t.Errorf("\nwant: %v\ngot : %v", tt.want, got)
			}
		})
	}
}

func TestDifference(t *testing.T) {
	tests := []struct {
		name        string
		minuend     Set[int]
		subtrahends []Set[int]
		want        Set[int]
	}{
		{
			name:        "nil minuend",
			minuend:     nil,
			subtrahends: []Set[int]{From(1, 2, 3)},
			want:        New[int](0),
		},
		{
			name:    "no subtrahends",
			minuend: From(1, 2, 3),
			want:    From(1, 2, 3),
		},
		{
			name:        "nil subtrahend",
			minuend:     From(1, 2, 3),
			subtrahends: []Set[int]{nil},
			want:        From(1, 2, 3),
		},
		{
			name:        "empty minuend",
			minuend:     New[int](0),
			subtrahends: []Set[int]{From(1, 2, 3)},
			want:        New[int](0),
		},
		{
			name:        "empty subtrahend",
			minuend:     From(1, 2, 3),
			subtrahends: []Set[int]{New[int](0)},
			want:        From(1, 2, 3),
		},
		{
			name:        "single disjoint subtrahend",
			minuend:     From(1, 2, 3),
			subtrahends: []Set[int]{From(4, 5, 6)},
			want:        From(1, 2, 3),
		},
		{
			name:        "single subtrahend partial overlap",
			minuend:     From(1, 2, 3),
			subtrahends: []Set[int]{From(2, 3, 4)},
			want:        From(1),
		},
		{
			name:        "single subtrahend equals minuend",
			minuend:     From(1, 2, 3),
			subtrahends: []Set[int]{From(1, 2, 3)},
			want:        New[int](0),
		},
		{
			name:        "multiple subtrahends",
			minuend:     From(1, 2, 3, 4, 5, 6, 7),
			subtrahends: []Set[int]{From(1, 3), From(5, 7)},
			want:        From(2, 4, 6),
		},
		{
			name:        "multiple subtrahends remove all",
			minuend:     From(1, 2, 3, 4),
			subtrahends: []Set[int]{From(1, 2), From(2, 3), From(3, 4)},
			want:        New[int](0),
		},
		{
			name:        "complex",
			minuend:     From(1, 2, 3, 4, 5),
			subtrahends: []Set[int]{From(1, 2), nil, From(1, 2), nil, From(4), From(6, 7)},
			want:        From(3, 5),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Difference(tt.minuend, tt.subtrahends...)
			if !Equal(got, tt.want) {
				t.Errorf("\nwant: %v\ngot : %v", tt.want, got)
			}
		})
	}
}

func TestSymmetricDifference(t *testing.T) {
	tests := []struct {
		name string
		sets []Set[int]
		want Set[int]
	}{
		{
			name: "no sets",
			want: New[int](0),
		},

		{
			name: "single nil set",
			sets: []Set[int]{nil},
			want: New[int](0),
		},
		{
			name: "single empty set",
			sets: []Set[int]{New[int](0)},
			want: New[int](0),
		},
		{
			name: "single set",
			sets: []Set[int]{From(1, 2, 3)},
			want: From(1, 2, 3),
		},
		{
			name: "with nil sets",
			sets: []Set[int]{nil, From(1, 2, 3), nil},
			want: From(1, 2, 3),
		},
		{
			name: "with empty sets",
			sets: []Set[int]{New[int](0), From(1, 2, 3), New[int](0)},
			want: From(1, 2, 3),
		},
		{
			name: "disjoint",
			sets: []Set[int]{From(1, 2), From(3, 4), From(5, 6)},
			want: From(1, 2, 3, 4, 5, 6),
		},
		{
			name: "partial overlap",
			sets: []Set[int]{From(1, 2, 3), From(3, 4, 5), From(5, 6, 7)},
			want: From(1, 2, 4, 6, 7),
		},
		{
			name: "even number of equal sets",
			sets: []Set[int]{From(1, 2, 3), From(1, 2, 3)},
			want: New[int](0),
		},
		{
			name: "odd number of equal sets",
			sets: []Set[int]{From(1, 2, 3), From(1, 2, 3), From(1, 2, 3)},
			want: From(1, 2, 3),
		},
		{
			name: "complex",
			sets: []Set[int]{From(1, 2, 3), nil, From(2, 3, 4), nil, From(3, 4, 5), From(3), From(6, 7)},
			want: From(1, 5, 6, 7),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SymmetricDifference(tt.sets...)
			if !Equal(got, tt.want) {
				t.Errorf("\nwant: %v\ngot : %v", tt.want, got)
			}
		})
	}
}

func TestCartesianProduct(t *testing.T) {
	tests := []struct {
		name string
		set1 Set[int]
		set2 Set[string]
		want Set[Pair[int, string]]
	}{
		{
			name: "both sets nil",
			set1: nil,
			set2: nil,
			want: New[Pair[int, string]](0),
		},
		{
			name: "first set nil",
			set1: nil,
			set2: From("a", "b"),
			want: New[Pair[int, string]](0),
		},
		{
			name: "second set nil",
			set1: From(1, 2),
			set2: nil,
			want: New[Pair[int, string]](0),
		},
		{
			name: "both sets empty",
			set1: New[int](0),
			set2: New[string](0),
			want: New[Pair[int, string]](0),
		},
		{
			name: "first set empty",
			set1: New[int](0),
			set2: From("a", "b"),
			want: New[Pair[int, string]](0),
		},
		{
			name: "second set empty",
			set1: From(1, 2),
			set2: New[string](0),
			want: New[Pair[int, string]](0),
		},
		{
			name: "single element sets",
			set1: From(42),
			set2: From("a"),
			want: From(Pair[int, string]{First: 42, Second: "a"}),
		},
		{
			name: "different size sets",
			set1: From(1, 2, 3),
			set2: From("a", "b"),
			want: From(
				Pair[int, string]{First: 1, Second: "a"},
				Pair[int, string]{First: 1, Second: "b"},
				Pair[int, string]{First: 2, Second: "a"},
				Pair[int, string]{First: 2, Second: "b"},
				Pair[int, string]{First: 3, Second: "a"},
				Pair[int, string]{First: 3, Second: "b"},
			),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CartesianProduct(tt.set1, tt.set2)
			if !Equal(got, tt.want) {
				t.Errorf("\nwant: %v\ngot : %v", tt.want, got)
			}
		})
	}
}
