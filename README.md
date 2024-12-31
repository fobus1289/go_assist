# go_assist

go_assist is a collection of helper functions for Go.

## Installation

```bash
go get github.com/fobus1289/go_assist
```

## Usage

```go
// Import the package
import "github.com/fobus1289/go_assist"

// Map example - double all numbers
numbers := []int{1, 2, 3, 4, 5}
doubled := goassist.Map(numbers, func(x int) int {
    return x * 2
})
// doubled is []int{2, 4, 6, 8, 10}

// Filter example - get even numbers
evens := goassist.Filter(numbers, func(x int) bool {
    return x%2 == 0
})
// evens is []int{2, 4}

// Reduce example - sum all numbers
sum := goassist.Reduce(numbers, func(acc, x int) int {
    return acc + x
}, 0)
// sum is 15

// Find example - get first number greater than 3
first, found := goassist.Find(numbers, func(x int) bool {
    return x > 3
})
// first is 4, found is true

// String examples
names := []string{"alice", "bob", "charlie"}

// Map strings to their lengths
lengths := goassist.Map(names, func(s string) int {
    return len(s)
})
// lengths is []int{5, 3, 7}

// Filter long names
longNames := goassist.Filter(names, func(s string) bool {
    return len(s) > 4
})
// longNames is []string{"alice", "charlie"}

// Concatenate strings
words := []string{"hello", " ", "world"}
sentence := goassist.Reduce(words, func(acc, s string) string {
    return acc + s
}, "")
// sentence is "hello world"


```
