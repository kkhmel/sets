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
	"strconv"
	"testing"
)

func TestMap(t *testing.T) {
	tests := []struct {
		name string
		s    Set[int]
		f    func(int) string
		want Set[string]
	}{
		{
			name: "nil set",
			s:    nil,
			f:    strconv.Itoa,
			want: New[string](0),
		},
		{
			name: "empty set",
			s:    New[int](0),
			f:    strconv.Itoa,
			want: New[string](0),
		},
		{
			name: "single element",
			s:    From(42),
			f:    strconv.Itoa,
			want: From("42"),
		},
		{
			name: "multiple elements",
			s:    From(1, 2, 3),
			f:    strconv.Itoa,
			want: From("1", "2", "3"),
		},
		{
			name: "duplicates in result",
			s:    From(1, 2, 3, 4),
			f:    func(n int) string { return strconv.Itoa(n & 1) },
			want: From("0", "1"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Map(tt.s, tt.f)
			if !Equal(got, tt.want) {
				t.Errorf("\nwant: %v\ngot : %v", tt.want, got)
			}
		})
	}
}

func TestFilter(t *testing.T) {
	tests := []struct {
		name string
		s    Set[int]
		f    func(int) bool
		want Set[int]
	}{
		{
			name: "nil set",
			s:    nil,
			f:    func(n int) bool { return n > 0 },
			want: New[int](0),
		},
		{
			name: "empty set",
			s:    New[int](0),
			f:    func(n int) bool { return n > 0 },
			want: New[int](0),
		},
		{
			name: "single element fails",
			s:    From(-42),
			f:    func(n int) bool { return n > 0 },
			want: New[int](0),
		},
		{
			name: "single element passes",
			s:    From(42),
			f:    func(n int) bool { return n > 0 },
			want: From(42),
		},
		{
			name: "no elements pass",
			s:    From(1, 2, 3),
			f:    func(int) bool { return false },
			want: New[int](0),
		},
		{
			name: "part of elements passes",
			s:    From(1, -2, 3, -4),
			f:    func(n int) bool { return n > 0 },
			want: From(1, 3),
		},
		{
			name: "all elements pass",
			s:    From(1, 2, 3),
			f:    func(int) bool { return true },
			want: From(1, 2, 3),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Filter(tt.s, tt.f)
			if !Equal(got, tt.want) {
				t.Errorf("\nwant: %v\ngot : %v", tt.want, got)
			}
		})
	}
}

func TestReduce(t *testing.T) {
	tests := []struct {
		name string
		s    Set[string]
		acc  int
		f    func(int, string) int
		want int
	}{
		{
			name: "nil set",
			s:    nil,
			acc:  10,
			want: 10,
		},
		{
			name: "empty set",
			s:    New[string](0),
			acc:  10,
			want: 10,
		},
		{
			name: "zero acc",
			s:    From("1", "2", "3"),
			acc:  0,
			f: func(i int, s string) int {
				n, _ := strconv.Atoi(s)
				return i + n
			},
			want: 6,
		},
		{
			name: "non-zero acc",
			s:    From("1", "2", "3"),
			acc:  10,
			f: func(i int, s string) int {
				n, _ := strconv.Atoi(s)
				return i + n
			},
			want: 16,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Reduce(tt.s, tt.acc, tt.f)
			if got != tt.want {
				t.Errorf("\nwant: %v\ngot : %v", tt.want, got)
			}
		})
	}
}
