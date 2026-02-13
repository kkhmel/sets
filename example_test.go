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

package sets_test

import (
	"fmt"

	"github.com/kkhmel/sets"
)

func Example() {
	// Create sets of user permissions
	admin := sets.From("read", "write", "delete", "admin")
	editor := sets.From("read", "write")
	viewer := sets.From("read")

	// Check permissions
	if sets.ContainsAll(admin, "delete", "admin") {
		fmt.Println("Admin has delete and admin privileges")
	}

	// Combine all unique permissions
	allPerms := sets.Union(admin, editor, viewer)
	fmt.Println("All permissions:", allPerms)

	// Find admin-exclusive permissions
	exclusive := sets.Difference(admin, editor, viewer)
	fmt.Println("Admin-only:", exclusive)

	// Output:
	// Admin has delete and admin privileges
	// All permissions: {admin, delete, read, write}
	// Admin-only: {admin, delete}
}

func ExampleFromSlice() {
	numbers := []int{1, 2, 3, 3, 2, 1}
	s := sets.FromSlice(numbers)
	fmt.Println(s)

	// Output:
	// {1, 2, 3}
}

func ExampleFromSliceFunc() {
	numbers := []int{1, 2, 3, 3, 2, 1}
	s := sets.FromSliceFunc(numbers, func(e int) int { return -e })
	fmt.Println(s)

	// Output:
	// {-1, -2, -3}
}

func ExampleUnion() {
	set1 := sets.From(1, 2, 3)
	set2 := sets.From(3, 4, 5)
	set3 := sets.From(5, 6, 7)
	union := sets.Union(set1, set2, set3)
	fmt.Println(union)

	// Output:
	// {1, 2, 3, 4, 5, 6, 7}
}

func ExampleIntersection() {
	set1 := sets.From(1, 2, 3, 4)
	set2 := sets.From(2, 3, 4, 5)
	set3 := sets.From(3, 4, 5, 6)
	intersection := sets.Intersection(set1, set2, set3)
	fmt.Println(intersection)

	// Output:
	// {3, 4}
}

func ExampleFilter() {
	numbers := sets.From(1, 2, 3, 4, 5)
	odds := sets.Filter(numbers, func(e int) bool { return e%2 == 1 })
	fmt.Println(odds)

	// Output:
	// {1, 3, 5}
}

func ExampleDeleteFunc() {
	numbers := sets.From(1, 2, 3, 4, 5)
	sets.DeleteFunc(numbers, func(e int) bool { return e%2 == 0 })
	fmt.Println(numbers)

	// Output:
	// {1, 3, 5}
}

func ExampleMap() {
	numbers := sets.From(1, 2, 3)
	doubled := sets.Map(numbers, func(e int) int { return -e })
	fmt.Println(doubled)

	// Output:
	// {-1, -2, -3}
}

func ExampleReplaceFunc() {
	numbers := sets.From(1, 2, 3)
	sets.ReplaceFunc(numbers, func(e int) int { return -e })
	fmt.Println(numbers)

	// Output:
	// {-1, -2, -3}
}

func ExampleEvery() {
	mixed := sets.From(1, 2, 3, 4)
	allPositive := sets.Every(mixed, func(e int) bool { return e > 0 })
	fmt.Println("All positive:", allPositive)

	// Output:
	// All positive: true
}

func ExampleChunk() {
	numbers := sets.From(1, 2, 3, 4, 5, 6, 7)
	chunkCount := 0
	for chunk := range sets.Chunk(numbers, 3) {
		chunkCount++
		fmt.Printf("Chunk %d size: %d\n", chunkCount, len(chunk))
	}

	// Output:
	// Chunk 1 size: 3
	// Chunk 2 size: 3
	// Chunk 3 size: 1
}
