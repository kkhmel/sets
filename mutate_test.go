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

func TestInsert(t *testing.T) {
	tests := []struct {
		name string
		set  Set[int]
		v    []int
		want Set[int]
	}{
		{
			name: "empty set",
			set:  New[int](0),
			v:    []int{1, 2, 3},
			want: From(1, 2, 3),
		},
		{
			name: "no elements",
			set:  From(1, 2, 3),
			v:    nil,
			want: From(1, 2, 3),
		},
		{
			name: "single member",
			set:  From(1, 2, 3),
			v:    []int{2},
			want: From(1, 2, 3),
		},
		{
			name: "single non-member",
			set:  From(1, 2, 3),
			v:    []int{4},
			want: From(1, 2, 3, 4),
		},
		{
			name: "multiple members",
			set:  From(1, 2, 3),
			v:    []int{1, 2, 3},
			want: From(1, 2, 3),
		},
		{
			name: "multiple elements some members",
			set:  From(1, 2, 3),
			v:    []int{2, 3, 4},
			want: From(1, 2, 3, 4),
		},
		{
			name: "multiple non-members",
			set:  From(1, 2, 3),
			v:    []int{4, 5, 6},
			want: From(1, 2, 3, 4, 5, 6),
		},
		{
			name: "duplicate elements",
			set:  From(1, 2, 3),
			v:    []int{3, 3, 4, 4, 5, 5},
			want: From(1, 2, 3, 4, 5),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Insert(tt.set, tt.v...)
			if !Equal(tt.set, tt.want) {
				t.Errorf("\nwant: %v\ngot : %v", tt.want, tt.set)
			}
		})
	}
}

func TestDelete(t *testing.T) {
	tests := []struct {
		name     string
		set      Set[int]
		elements []int
		want     Set[int]
	}{
		{
			name:     "nil set",
			set:      New[int](0),
			elements: []int{1, 2, 3},
			want:     New[int](0),
		},
		{
			name:     "empty set",
			set:      New[int](0),
			elements: []int{1, 2, 3},
			want:     New[int](0),
		},
		{
			name:     "no elements",
			set:      From(1, 2, 3),
			elements: nil,
			want:     From(1, 2, 3),
		},
		{
			name:     "single non-member",
			set:      From(1, 2, 3),
			elements: []int{42},
			want:     From(1, 2, 3),
		},
		{
			name:     "single member",
			set:      From(1, 2, 3),
			elements: []int{2},
			want:     From(1, 3),
		},
		{
			name:     "multiple non-members",
			set:      From(1, 2, 3),
			elements: []int{4, 5, 6},
			want:     From(1, 2, 3),
		},
		{
			name:     "multiple elements some members",
			set:      From(1, 2, 3),
			elements: []int{2, 4, 5},
			want:     From(1, 3),
		},
		{
			name:     "multiple members",
			set:      From(1, 2, 3, 4, 5),
			elements: []int{1, 3, 5},
			want:     From(2, 4),
		},
		{
			name:     "delete all elements",
			set:      From(1, 2, 3),
			elements: []int{1, 2, 3},
			want:     New[int](0),
		},
		{
			name:     "duplicate elements",
			set:      From(1, 2, 3, 4),
			elements: []int{2, 2, 3, 3},
			want:     From(1, 4),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Delete(tt.set, tt.elements...)
			if !Equal(tt.set, tt.want) {
				t.Errorf("\nwant: %v\ngot : %v", tt.want, tt.set)
			}
		})
	}
}

func TestDeleteFunc(t *testing.T) {
	tests := []struct {
		name string
		s    Set[int]
		f    func(int) bool
		want Set[int]
	}{
		{
			name: "nil set",
			s:    nil,
			f:    func(n int) bool { return n < 0 },
			want: nil,
		},
		{
			name: "empty set",
			s:    New[int](0),
			f:    func(n int) bool { return n < 0 },
			want: New[int](0),
		},
		{
			name: "no elements",
			s:    From(1, 2, 3),
			f:    func(n int) bool { return n < 0 },
			want: From(1, 2, 3),
		},
		{
			name: "single element",
			s:    From(1, -2, 3),
			f:    func(n int) bool { return n < 0 },
			want: From(1, 3),
		},
		{
			name: "multiple elements",
			s:    From(1, -2, 3, -4, 5),
			f:    func(n int) bool { return n < 0 },
			want: From(1, 3, 5),
		},
		{
			name: "delete all elements",
			s:    From(-1, -2, -3),
			f:    func(n int) bool { return n < 0 },
			want: New[int](0),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DeleteFunc(tt.s, tt.f)
			if !Equal(tt.s, tt.want) {
				t.Errorf("\nwant: %v\ngot : %v", tt.want, tt.s)
			}
		})
	}
}

func TestReplace(t *testing.T) {
	tests := []struct {
		name string
		s    Set[int]
		old  int
		new  int
		want Set[int]
	}{
		{
			name: "empty set",
			s:    New[int](0),
			old:  1,
			new:  2,
			want: New[int](0),
		},
		{
			name: "single element set",
			s:    From(1),
			old:  1,
			new:  2,
			want: From(2),
		},
		{
			name: "member",
			s:    From(1, 2, 3),
			old:  2,
			new:  4,
			want: From(1, 3, 4),
		},
		{
			name: "non-member",
			s:    From(1, 2, 3),
			old:  4,
			new:  5,
			want: From(1, 2, 3),
		},
		{
			name: "replace with same value",
			s:    From(1, 2, 3),
			old:  2,
			new:  2,
			want: From(1, 2, 3),
		},
		{
			name: "replace with member",
			s:    From(1, 2, 3),
			old:  2,
			new:  3,
			want: From(1, 3),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Replace(tt.s, tt.old, tt.new)
			if !Equal(tt.s, tt.want) {
				t.Errorf("\nwant: %v\ngot : %v", tt.want, tt.s)
			}
		})
	}
}

func TestReplaceFunc(t *testing.T) {
	tests := []struct {
		name string
		s    Set[int]
		f    func(int) int
		want Set[int]
	}{
		{
			name: "empty set",
			s:    New[int](0),
			f:    func(n int) int { return n * 2 },
			want: New[int](0),
		},
		{
			name: "single element",
			s:    From(1),
			f:    func(n int) int { return n * 2 },
			want: From(2),
		},
		{
			name: "identity function",
			s:    From(1, 2, 3),
			f:    func(n int) int { return n },
			want: From(1, 2, 3),
		},
		{
			name: "double all elements",
			s:    From(1, 2, 3),
			f:    func(n int) int { return n * 2 },
			want: From(2, 4, 6),
		},
		{
			name: "partial collision",
			s:    From(1, 2, 3, 4),
			f:    func(n int) int { return n / 2 },
			want: From(0, 1, 2),
		},
		{
			name: "constant function collapses set",
			s:    From(1, 2, 3),
			f:    func(int) int { return 0 },
			want: From(0),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ReplaceFunc(tt.s, tt.f)
			if !Equal(tt.s, tt.want) {
				t.Errorf("\nwant: %v\ngot : %v", tt.want, tt.s)
			}
		})
	}
}
