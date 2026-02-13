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
	"slices"
	"testing"
)

func TestAll(t *testing.T) {
	tests := []struct {
		name string
		s    Set[int]
		want []int
	}{
		{
			name: "nil",
			s:    nil,
			want: []int{},
		},
		{
			name: "empty",
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
			got := make([]int, 0, len(tt.s))
			for e := range All(tt.s) {
				got = append(got, e)
			}
			slices.Sort(got)
			if !slices.Equal(got, tt.want) {
				t.Errorf("\nwant: %v\ngot : %v", tt.want, got)
			}
		})
	}
}

func TestAllEarlyTermination(t *testing.T) {
	s := From(1, 2, 3, 4, 5)
	count := 0
	for range All(s) {
		count++
		if count == 2 {
			break
		}
	}
	if count != 2 {
		t.Errorf("expected to iterate 2 times, got %d", count)
	}
}

func TestInsertSeq(t *testing.T) {
	tests := []struct {
		name string
		s    Set[int]
		seq  iter.Seq[int]
		want Set[int]
	}{
		{
			name: "empty set",
			s:    New[int](0),
			seq:  slices.Values([]int{1, 2, 3}),
			want: From(1, 2, 3),
		},
		{
			name: "empty seq",
			s:    From(1, 2),
			seq:  slices.Values([]int{}),
			want: From(1, 2),
		},
		{
			name: "single element",
			s:    From(1, 2),
			seq:  slices.Values([]int{3}),
			want: From(1, 2, 3),
		},
		{
			name: "multiple elements",
			s:    From(1, 2),
			seq:  slices.Values([]int{3, 4, 5}),
			want: From(1, 2, 3, 4, 5),
		},
		{
			name: "duplicate elements",
			s:    From(1, 2),
			seq:  slices.Values([]int{2, 3, 3}),
			want: From(1, 2, 3),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			InsertSeq(tt.s, tt.seq)
			if !Equal(tt.s, tt.want) {
				t.Errorf("\nwant: %v\ngot : %v", tt.want, tt.s)
			}
		})
	}
}

func TestCollect(t *testing.T) {
	tests := []struct {
		name string
		seq  iter.Seq[int]
		want Set[int]
	}{
		{
			name: "empty seq",
			seq:  slices.Values([]int{}),
			want: New[int](0),
		},
		{
			name: "single element",
			seq:  slices.Values([]int{42}),
			want: From(42),
		},
		{
			name: "multiple elements",
			seq:  slices.Values([]int{1, 2, 3}),
			want: From(1, 2, 3),
		},
		{
			name: "duplicate elements",
			seq:  slices.Values([]int{1, 2, 2, 3, 3, 3}),
			want: From(1, 2, 3),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Collect(tt.seq)
			if !Equal(got, tt.want) {
				t.Errorf("\nwant: %v\ngot : %v", tt.want, got)
			}
		})
	}
}

func TestChunk(t *testing.T) {
	tests := []struct {
		name       string
		s          Set[int]
		n          int
		wantChunks int
		wantTotal  int
	}{
		{
			name:       "nil set",
			s:          nil,
			n:          2,
			wantChunks: 0,
			wantTotal:  0,
		},
		{
			name:       "empty set",
			s:          New[int](0),
			n:          2,
			wantChunks: 0,
			wantTotal:  0,
		},
		{
			name:       "chunk size equals set size",
			s:          From(1, 2, 3),
			n:          3,
			wantChunks: 1,
			wantTotal:  3,
		},
		{
			name:       "chunk size larger than set",
			s:          From(1, 2),
			n:          5,
			wantChunks: 1,
			wantTotal:  2,
		},
		{
			name:       "even division",
			s:          From(1, 2, 3, 4),
			n:          2,
			wantChunks: 2,
			wantTotal:  4,
		},
		{
			name:       "uneven division",
			s:          From(1, 2, 3, 4, 5),
			n:          2,
			wantChunks: 3,
			wantTotal:  5,
		},
		{
			name:       "chunk size of 1",
			s:          From(1, 2, 3),
			n:          1,
			wantChunks: 3,
			wantTotal:  3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			chunks := 0
			total := 0
			for chunk := range Chunk(tt.s, tt.n) {
				chunks++
				total += len(chunk)
				if len(chunk) > tt.n {
					t.Errorf("chunk size %d exceeds limit %d", len(chunk), tt.n)
				}
			}
			if chunks != tt.wantChunks {
				t.Errorf("\nwant chunks: %d\ngot  chunks : %d", tt.wantChunks, chunks)
			}
			if total != tt.wantTotal {
				t.Errorf("\nwant total: %d\ngot  total : %d", tt.wantTotal, total)
			}
		})
	}
}

func TestChunkEarlyTermination(t *testing.T) {
	s := From(1, 2, 3, 4, 5, 6, 7, 8)
	chunkCount := 0
	for range Chunk(s, 2) {
		chunkCount++
		if chunkCount == 2 {
			break
		}
	}
	if chunkCount != 2 {
		t.Errorf("expected to iterate 2 chunks, got %d", chunkCount)
	}
}

func TestChunkPanic(t *testing.T) {
	tests := []struct {
		name string
		n    int
	}{
		{
			name: "zero",
			n:    0,
		},
		{
			name: "negative",
			n:    -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r == nil {
					t.Errorf("Chunk() should panic when n is %d, but did not panic", tt.n)
				}
			}()

			s := From(1, 2, 3)
			for chunk := range Chunk(s, tt.n) {
				_ = chunk
			}
		})
	}
}

func TestUnionSeq(t *testing.T) {
	tests := []struct {
		name string
		seq  iter.Seq[Set[int]]
		want Set[int]
	}{
		{
			name: "empty seq",
			seq:  slices.Values([]Set[int]{}),
			want: New[int](0),
		},
		{
			name: "single nil set",
			seq:  slices.Values([]Set[int]{nil}),
			want: New[int](0),
		},
		{
			name: "single empty set",
			seq:  slices.Values([]Set[int]{New[int](0)}),
			want: New[int](0),
		},
		{
			name: "single set",
			seq:  slices.Values([]Set[int]{From(1, 2, 3)}),
			want: From(1, 2, 3),
		},
		{
			name: "disjoint sets",
			seq:  slices.Values([]Set[int]{From(1, 2), From(3, 4), From(5, 6)}),
			want: From(1, 2, 3, 4, 5, 6),
		},
		{
			name: "overlapping sets",
			seq:  slices.Values([]Set[int]{From(1, 2, 3), From(2, 3, 4), From(3, 4, 5)}),
			want: From(1, 2, 3, 4, 5),
		},
		{
			name: "complex",
			seq:  slices.Values([]Set[int]{New[int](0), From(1, 2), From(1, 2), nil, From(3, 4), From(1, 2)}),
			want: From(1, 2, 3, 4),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Collect(UnionSeq(tt.seq))
			if !Equal(got, tt.want) {
				t.Errorf("\nwant: %v\ngot : %v", tt.want, got)
			}
		})
	}
}

func TestUnionSeqEarlyTermination(t *testing.T) {
	seq := slices.Values([]Set[int]{From(1, 2), From(3, 4), From(5, 6)})
	count := 0
	for range UnionSeq(seq) {
		count++
		if count == 3 {
			break
		}
	}
	if count != 3 {
		t.Errorf("expected to iterate 3 elements, got %d", count)
	}
}
