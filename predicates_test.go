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

import "testing"

func TestContains(t *testing.T) {
	tests := []struct {
		name string
		s    Set[int]
		v    int
		want bool
	}{
		{
			name: "nil set",
			s:    nil,
			v:    1,
			want: false,
		},
		{
			name: "empty set",
			s:    New[int](0),
			v:    1,
			want: false,
		},
		{
			name: "member",
			s:    From(1, 2, 3),
			v:    2,
			want: true,
		},
		{
			name: "non-member",
			s:    From(1, 2, 3),
			v:    4,
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Contains(tt.s, tt.v)
			if got != tt.want {
				t.Errorf("\nwant: %v\ngot : %v", tt.want, got)
			}
		})
	}
}

func TestContainsAny(t *testing.T) {
	tests := []struct {
		name string
		s    Set[int]
		v    []int
		want bool
	}{
		{
			name: "nil set no elements",
			s:    nil,
			v:    nil,
			want: false,
		},
		{
			name: "nil set with elements",
			s:    nil,
			v:    []int{1, 2, 3},
			want: false,
		},
		{
			name: "empty set no elements",
			s:    New[int](0),
			v:    nil,
			want: false,
		},
		{
			name: "empty set with elements",
			s:    New[int](0),
			v:    []int{1, 2, 3},
			want: false,
		},
		{
			name: "no elements",
			s:    From(1, 2, 3),
			v:    nil,
			want: false,
		},
		{
			name: "single member",
			s:    From(1, 2, 3),
			v:    []int{2},
			want: true,
		},
		{
			name: "single non-member",
			s:    From(1, 2, 3),
			v:    []int{4},
			want: false,
		},
		{
			name: "all members",
			s:    From(1, 2, 3),
			v:    []int{1, 2},
			want: true,
		},
		{
			name: "some members",
			s:    From(1, 2, 3),
			v:    []int{3, 4, 5},
			want: true,
		},
		{
			name: "no members",
			s:    From(1, 2, 3),
			v:    []int{4, 5, 6},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ContainsAny(tt.s, tt.v...)
			if got != tt.want {
				t.Errorf("\nwant: %v\ngot : %v", tt.want, got)
			}
		})
	}
}

func TestContainsAll(t *testing.T) {
	tests := []struct {
		name string
		s    Set[int]
		v    []int
		want bool
	}{
		{
			name: "nil set no elements",
			s:    nil,
			v:    nil,
			want: true,
		},
		{
			name: "nil set with elements",
			s:    nil,
			v:    []int{1, 2, 3},
			want: false,
		},
		{
			name: "no elements",
			s:    From(1, 2, 3),
			v:    nil,
			want: true,
		},
		{
			name: "empty set no elements",
			s:    New[int](0),
			v:    nil,
			want: true,
		},
		{
			name: "empty set with elements",
			s:    New[int](0),
			v:    []int{1, 2, 3},
			want: false,
		},
		{
			name: "single member",
			s:    From(1, 2, 3),
			v:    []int{2},
			want: true,
		},
		{
			name: "single non-member",
			s:    From(1, 2, 3),
			v:    []int{4},
			want: false,
		},
		{
			name: "all members",
			s:    From(1, 2, 3),
			v:    []int{1, 2},
			want: true,
		},
		{
			name: "some non-members",
			s:    From(1, 2, 3),
			v:    []int{1, 4},
			want: false,
		},
		{
			name: "no members",
			s:    From(1, 2, 3),
			v:    []int{4, 5, 6},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ContainsAll(tt.s, tt.v...)
			if got != tt.want {
				t.Errorf("\nwant: %v\ngot : %v", tt.want, got)
			}
		})
	}
}

func TestSome(t *testing.T) {
	tests := []struct {
		name string
		s    Set[int]
		f    func(int) bool
		want bool
	}{
		{
			name: "nil set",
			s:    nil,
			f:    func(n int) bool { return n > 0 },
			want: false,
		},
		{
			name: "empty set",
			s:    New[int](0),
			f:    func(n int) bool { return n > 0 },
			want: false,
		},
		{
			name: "single element passes",
			s:    From(42),
			f:    func(n int) bool { return n > 0 },
			want: true,
		},
		{
			name: "single element fails",
			s:    From(-42),
			f:    func(n int) bool { return n > 0 },
			want: false,
		},
		{
			name: "all elements pass",
			s:    From(1, 2, 3),
			f:    func(n int) bool { return n > 0 },
			want: true,
		},
		{
			name: "some elements pass",
			s:    From(-1, 2, -3),
			f:    func(n int) bool { return n > 0 },
			want: true,
		},
		{
			name: "all elements fail",
			s:    From(-1, -2, -3),
			f:    func(n int) bool { return n > 0 },
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Some(tt.s, tt.f)
			if got != tt.want {
				t.Errorf("\nwant: %v\ngot : %v", tt.want, got)
			}
		})
	}
}

func TestEvery(t *testing.T) {
	tests := []struct {
		name string
		s    Set[int]
		f    func(int) bool
		want bool
	}{
		{
			name: "nil set",
			s:    nil,
			f:    func(n int) bool { return n > 0 },
			want: true,
		},
		{
			name: "empty set",
			s:    New[int](0),
			f:    func(n int) bool { return n > 0 },
			want: true,
		},
		{
			name: "single element passes",
			s:    From(42),
			f:    func(n int) bool { return n > 0 },
			want: true,
		},
		{
			name: "single element fails",
			s:    From(-42),
			f:    func(n int) bool { return n > 0 },
			want: false,
		},
		{
			name: "all elements pass",
			s:    From(1, 2, 3),
			f:    func(n int) bool { return n > 0 },
			want: true,
		},
		{
			name: "some elements fail",
			s:    From(1, -2, 3),
			f:    func(n int) bool { return n > 0 },
			want: false,
		},
		{
			name: "all elements fail",
			s:    From(-1, -2, -3),
			f:    func(n int) bool { return n > 0 },
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Every(tt.s, tt.f)
			if got != tt.want {
				t.Errorf("\nwant: %v\ngot : %v", tt.want, got)
			}
		})
	}
}

func TestEqual(t *testing.T) {
	tests := []struct {
		name string
		s1   Set[int]
		s2   Set[int]
		want bool
	}{
		{
			name: "nil sets",
			s1:   nil,
			s2:   nil,
			want: true,
		},
		{
			name: "first set nil",
			s1:   nil,
			s2:   From(1, 2, 3),
			want: false,
		},
		{
			name: "second set nil",
			s1:   From(1, 2, 3),
			s2:   nil,
			want: false,
		},
		{
			name: "empty sets",
			s1:   New[int](0),
			s2:   New[int](0),
			want: true,
		},
		{
			name: "first set empty",
			s1:   New[int](0),
			s2:   From(1, 2, 3),
			want: false,
		},
		{
			name: "second set empty",
			s1:   From(1, 2, 3),
			s2:   New[int](0),
			want: false,
		},
		{
			name: "equal single element sets",
			s1:   From(42),
			s2:   From(42),
			want: true,
		},
		{
			name: "different single element sets",
			s1:   From(42),
			s2:   From(24),
			want: false,
		},
		{
			name: "equal",
			s1:   From(1, 2, 3),
			s2:   From(1, 2, 3),
			want: true,
		},
		{
			name: "partial overlap",
			s1:   From(1, 2, 3),
			s2:   From(2, 3, 4),
			want: false,
		},
		{
			name: "disjoint",
			s1:   From(1, 2, 3),
			s2:   From(4, 5, 6),
			want: false,
		},
		{
			name: "different sizes",
			s1:   From(1, 2, 3),
			s2:   From(1, 2),
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Equal(tt.s1, tt.s2)
			if got != tt.want {
				t.Errorf("\nwant: %v\ngot : %v", tt.want, got)
			}
		})
	}
}

func TestOverlaps(t *testing.T) {
	tests := []struct {
		name string
		s1   Set[int]
		s2   Set[int]
		want bool
	}{
		{
			name: "nil sets",
			s1:   nil,
			s2:   nil,
			want: false,
		},
		{
			name: "first set nil",
			s1:   nil,
			s2:   From(1, 2, 3),
			want: false,
		},
		{
			name: "second set nil",
			s1:   From(1, 2, 3),
			s2:   nil,
			want: false,
		},
		{
			name: "empty sets",
			s1:   New[int](0),
			s2:   New[int](0),
			want: false,
		},
		{
			name: "first set empty",
			s1:   New[int](0),
			s2:   From(1, 2, 3),
			want: false,
		},
		{
			name: "second set empty",
			s1:   From(1, 2, 3),
			s2:   New[int](0),
			want: false,
		},
		{
			name: "first set bigger",
			s1:   From(1, 2, 3, 4),
			s2:   From(1, 2, 3),
			want: true,
		},
		{
			name: "second set bigger",
			s1:   From(1, 2, 3),
			s2:   From(1, 2, 3, 4),
			want: true,
		},
		{
			name: "single common element",
			s1:   From(1, 2, 3),
			s2:   From(3, 4, 5),
			want: true,
		},
		{
			name: "disjoint",
			s1:   From(1, 2, 3),
			s2:   From(4, 5, 6),
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Overlaps(tt.s1, tt.s2)
			if got != tt.want {
				t.Errorf("\nwant: %v\ngot : %v", tt.want, got)
			}
		})
	}
}

func TestSubset(t *testing.T) {
	tests := []struct {
		name     string
		subset   Set[int]
		superset Set[int]
		want     bool
	}{
		{
			name:     "nil sets",
			subset:   nil,
			superset: nil,
			want:     true,
		},
		{
			name:     "nil subset",
			subset:   nil,
			superset: From(1, 2, 3),
			want:     true,
		},
		{
			name:     "nil superset",
			subset:   From(1, 2, 3),
			superset: nil,
			want:     false,
		},
		{
			name:     "empty sets",
			subset:   New[int](0),
			superset: New[int](0),
			want:     true,
		},
		{
			name:     "empty subset",
			subset:   New[int](0),
			superset: From(1, 2, 3),
			want:     true,
		},
		{
			name:     "empty superset",
			subset:   From(1, 2, 3),
			superset: New[int](0),
			want:     false,
		},
		{
			name:     "equal sets",
			subset:   From(1, 2, 3),
			superset: From(1, 2, 3),
			want:     true,
		},
		{
			name:     "proper subset",
			subset:   From(1, 2),
			superset: From(1, 2, 3),
			want:     true,
		},
		{
			name:     "partial overlap",
			subset:   From(1, 2, 3),
			superset: From(2, 3, 4),
			want:     false,
		},
		{
			name:     "disjoint",
			subset:   From(1, 2, 3),
			superset: From(4, 5, 6),
			want:     false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Subset(tt.subset, tt.superset)
			if got != tt.want {
				t.Errorf("\nwant: %v\ngot : %v", tt.want, got)
			}
		})
	}
}

func TestProperSubset(t *testing.T) {
	tests := []struct {
		name     string
		subset   Set[int]
		superset Set[int]
		want     bool
	}{
		{
			name:     "nil sets",
			subset:   nil,
			superset: nil,
			want:     false,
		},
		{
			name:     "nil subset",
			subset:   nil,
			superset: From(1, 2, 3),
			want:     true,
		},
		{
			name:     "nil superset",
			subset:   From(1, 2, 3),
			superset: nil,
			want:     false,
		},
		{
			name:     "empty sets",
			subset:   New[int](0),
			superset: New[int](0),
			want:     false,
		},
		{
			name:     "empty subset",
			subset:   New[int](0),
			superset: From(1, 2, 3),
			want:     true,
		},
		{
			name:     "empty superset",
			subset:   From(1, 2, 3),
			superset: New[int](0),
			want:     false,
		},
		{
			name:     "equal sets",
			subset:   From(1, 2, 3),
			superset: From(1, 2, 3),
			want:     false,
		},
		{
			name:     "proper subset",
			subset:   From(1, 2),
			superset: From(1, 2, 3),
			want:     true,
		},
		{
			name:     "partial overlap",
			subset:   From(1, 2, 3),
			superset: From(2, 3, 4),
			want:     false,
		},
		{
			name:     "disjoint",
			subset:   From(1, 2, 3),
			superset: From(4, 5, 6),
			want:     false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ProperSubset(tt.subset, tt.superset)
			if got != tt.want {
				t.Errorf("\nwant: %v\ngot : %v", tt.want, got)
			}
		})
	}
}
