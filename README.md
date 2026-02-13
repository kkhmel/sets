# sets

[![Go Reference](https://pkg.go.dev/badge/github.com/kkhmel/sets.svg)](https://pkg.go.dev/github.com/kkhmel/sets)
[![Github release](https://img.shields.io/github/release/kkhmel/sets.svg)](https://github.com/kkhmel/sets/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/kkhmel/sets)](https://goreportcard.com/report/github.com/kkhmel/sets)
[![Lint Status](https://img.shields.io/github/actions/workflow/status/kkhmel/sets/golangci-lint.yml?branch=main&label=lint)](https://github.com/kkhmel/sets/actions)
[![Coverage Status](https://codecov.io/gh/kkhmel/sets/branch/main/graph/badge.svg)](https://codecov.io/gh/kkhmel/sets/branch/main)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

A high-performance, idiomatic Go library for set operations that follows the same design principles as the standard library's [`maps`](https://pkg.go.dev/maps) and [`slices`](https://pkg.go.dev/slices) packages.

## Design Philosophy: Idiomatic Go

This library combines **idiomatic Go design** with **practical functionality** that developers need in real-world projects.

**Native map at the core.** `Set[T]` is a **type alias** for `map[T]struct{}`, not a wrapper type. All native Go map operations work out of the box: `len()`, `clear()`, `delete()`, `for...range`.

**Follows standard library conventions.** The API mirrors the design of [`maps`](https://pkg.go.dev/maps) and [`slices`](https://pkg.go.dev/slices) — composable functions that each do one thing well, consistent naming, variadic parameters where they make sense. No method chaining or fluent APIs.

**Comprehensive.** ~30 functions covering set theory, functional programming, and iterators.

**Robust.** Read-only functions gracefully handle nil sets, consistent with how the standard library's `maps` and `slices` packages treat nil inputs. Pre-allocated capacity for optimal performance.

**Zero dependencies.** Pure Go standard library only — no external dependencies to manage or security vulnerabilities to track.

---

## Installation

```bash
go get github.com/kkhmel/sets
```

Requires Go 1.23+.

---

## Usage

### Quick Start

```go
package main

import (
    "fmt"
    "github.com/kkhmel/sets"
)

func main() {
    // Create sets of user permissions
    admin := sets.From("read", "write", "delete", "admin")
    editor := sets.From("read", "write")
    viewer := sets.From("read")

    // Check permissions
    if sets.ContainsAll(admin, "delete", "admin") {
        fmt.Println("Admin has delete and admin privileges") // Output: Admin has delete and admin privileges
    }

    // Combine all unique permissions
    allPerms := sets.Union(admin, editor, viewer)
    fmt.Println("All permissions:", allPerms) // Output: All permissions: {admin, delete, read, write}

    // Find admin-exclusive permissions
    exclusive := sets.Difference(admin, editor, viewer)
    fmt.Println("Admin-only:", exclusive) // Output: Admin-only: {admin, delete}

    // Check subset relationships
    if sets.Subset(viewer, editor) {
        fmt.Println("Viewers are a subset of editors") // Output: Viewers are a subset of editors
    }
}
```

### Guideline

Use builtin `len()` for size, `clear()` for clearing, and `range` for iteration. For all other operations, use `sets` package functions for consistency and cleaner code.

```go
s := sets.From("a", "b", "c")

// Use native operations for these common cases
count := len(s)    // Get size
clear(s)           // Clear all elements
for e := range s { // Iterate over elements
    ...
}

// Use sets package functions for everything else -- cleaner and more consistent
items := []string{"d", "e"}
sets.Insert(s, items...)                 // Instead of: for _, e := range items { s[e] = struct{}{} }
sets.Delete(s, items...)                 // Instead of: for _, e := range items { delete(s, e) }
if sets.Contains(s, "x") { ... }         // Instead of: if _, ok := s["x"]; ok { ... }
if sets.ContainsAny(s, items...) { ... } // Instead of: manual loop with checks
```

But since `Set[T]` is a type alias for `map[T]struct{}`, you have full access to native Go map operations and the [`maps`](https://pkg.go.dev/maps) package.

---

## Performance

The library is designed for maximum performance through several key principles:

**Core Performance Principles:**

- **Zero-cost abstraction:** Leverages Go's highly-optimized map implementation directly
- **Zero-byte values:** Uses `struct{}` as map value (0 bytes per element) for minimal memory footprint
- **Smart pre-allocation:** Set operations pre-allocate capacity based on input sizes to minimize reallocation
- **Inlining-friendly:** Simple functions are designed to be inlined by the compiler

**Complexity Information:**
Detailed time and space complexity for each function is documented in the [API documentation](https://pkg.go.dev/github.com/kkhmel/sets).

**Note:** Call `sets.Clone()` on results to reclaim excess memory if needed.

---

## Thread Safety

- **Safe for concurrent reads** — Multiple goroutines can safely read from a set simultaneously
- **Requires synchronization for writes** — Use `sync.RWMutex` or `sync.Mutex` when modifying sets concurrently

---

## Contributing

Contributions are welcome! Please see [CONTRIBUTING.md](CONTRIBUTING.md) for details on how to submit pull requests, report issues, and contribute to the code.
