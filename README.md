## Installation

To use this package, you can install it using:

```bash
go get github.com/fobus1289/go_assist
```

To add examples for the functions in your `helper.go` file to the README, you can create a section that describes each function along with code snippets demonstrating how to use them. Below is an example of how you might structure the README with examples for each function.

````markdown
# Go Assist

This package provides various utility functions for working with slices in Go.

## Functions

### Map

`func Map[T any, R any](arr []T, fn func(T) R) []R`

Applies a function to each element of the input slice and returns a new slice with the results.

**Example:**

```go
numbers := []int{1, 2, 3, 4, 5}
doubled := Map(numbers, func(x int) int {
    return x * 2
})
// doubled is []int{2, 4, 6, 8, 10}
```
````

### Filter

`func Filter[T any](arr []T, fn func(T) bool) []T`

Returns a new slice containing only the elements that satisfy the provided predicate function.

**Example:**

```go
numbers := []int{1, 2, 3, 4, 5}
evens := Filter(numbers, func(x int) bool {
    return x%2 == 0
})
// evens is []int{2, 4}
```

### Reduce

`func Reduce[T any, R any](arr []T, fn func(R, T) R, initial R) R`

Applies a function cumulatively to the elements of the slice, reducing it to a single value.

**Example:**

```go
numbers := []int{1, 2, 3, 4, 5}
sum := Reduce(numbers, func(acc, x int) int {
    return acc + x
}, 0)
// sum is 15
```

### Find

`func Find[T any](arr []T, fn func(T) bool) (T, bool)`

Returns the first element that satisfies the predicate function and a boolean indicating success.

**Example:**

```go
numbers := []int{1, 2, 3, 4, 5}
first, found := Find(numbers, func(x int) bool {
    return x > 3
})
// first is 4, found is true
```

### Some

`func Some[T any](arr []T, fn func(T) bool) bool`

Returns true if at least one element satisfies the predicate function.

**Example:**

```go
numbers := []int{1, 2, 3, 4, 5}
hasEven := Some(numbers, func(x int) bool {
    return x%2 == 0
})
// hasEven is true
```

### Every

`func Every[T any](arr []T, fn func(T) bool) bool`

Returns true if all elements satisfy the predicate function.

**Example:**

```go
numbers := []int{2, 4, 6, 8}
allEven := Every(numbers, func(x int) bool {
    return x%2 == 0
})
// allEven is true
```

### Flatten

`func Flatten[T any](arr [][]T) []T`

Flattens a slice of slices into a single slice.

**Example:**

```go
nested := [][]int{{1, 2}, {3, 4}, {5, 6}}
flat := Flatten(nested)
// flat is []int{1, 2, 3, 4, 5, 6}
```

### Zip

`func Zip[T any, R any](arr []T, arr2 []R) [][]any`

Combines two slices into a slice of pairs. If the input slices have different lengths, the result will have the length of the shorter slice.

**Example:**

```go
numbers := []int{1, 2, 3}
letters := []string{"a", "b", "c"}
pairs := Zip(numbers, letters)
// pairs is [][]any{{1, "a"}, {2, "b"}, {3, "c"}}
```

### Unzip

`func Unzip[T any, R any](arr [][]any) ([]T, []R)`

Splits a slice of pairs into two separate slices.

**Example:**

```go
pairs := [][]any{{1, 10}, {2, 20}, {3, 30}}
first, second := Unzip(pairs)
// first is []int{1, 2, 3}
// second is []int{10, 20, 30}
```

### BinarySearch

`func BinarySearch[S ~[]E, E cmp.Ordered](x S, target E) (int, bool)`

Performs a binary search on a sorted slice and returns the index and a boolean indicating if the target was found.

**Example:**

```go
numbers := []int{1, 3, 5, 7, 9}
index, found := BinarySearch(numbers, 5)
// index is 2, found is true
```

### BinarySearchFunc

`func BinarySearchFunc[S ~[]E, E, T any](x S, target T, cmp func(E, T) int) (int, bool)`

Performs a binary search using a custom comparison function.

**Example:**

```go
type Person struct {
    Name string
    Age  int
}
people := []Person{
    {"Alice", 25},
    {"Bob", 30},
    {"Charlie", 35},
}
index, found := BinarySearchFunc(people, 30, func(p Person, age int) int {
    return p.Age - age
})
// index is 1, found is true
```

### Clip

`func Clip[S ~[]E, E any](s S) S`

Returns a slice with the capacity adjusted to the length of the slice.

**Example:**

```go
s := make([]int, 3, 10)
s = Clip(s)
// len(s) == 3, cap(s) == 3
```

### Clone

`func Clone[S ~[]E, E any](s S) S`

Creates a copy of the slice.

**Example:**

```go
original := []int{1, 2, 3}
copy := Clone(original)
copy[0] = 99
// original is still []int{1, 2, 3}
```

### Compact

`func Compact[S ~[]E, E comparable](s S) S`

Removes adjacent duplicate elements from the slice.

**Example:**

```go
numbers := []int{1, 1, 2, 3, 3, 3, 4, 4, 5}
unique := Compact(numbers)
// unique is []int{1, 2, 3, 4, 5}
```

### CompactFunc

`func CompactFunc[S ~[]E, E any](s S, eq func(E, E) bool) S`

Removes adjacent duplicate elements based on a custom equality function.

**Example:**

```go
type Person struct {
    Name string
    Age  int
}
people := []Person{
    {"Alice", 25},
    {"Alice", 30},
    {"Bob", 30},
}
unique := CompactFunc(people, func(a, b Person) bool {
    return a.Name == b.Name
})
// unique has 2 elements: Alice and Bob
```

### Contains

`func Contains[S ~[]E, E comparable](s S, v E) bool`

Checks if a value exists in the slice.

**Example:**

```go
numbers := []int{1, 2, 3, 4, 5}
exists := Contains(numbers, 3)
// exists is true
```

### ContainsFunc

`func ContainsFunc[S ~[]E, E any](s S, f func(E) bool) bool`

Checks if any element satisfies a predicate function.

**Example:**

```go
numbers := []int{1, 2, 3, 4, 5}
hasEven := ContainsFunc(numbers, func(x int) bool {
    return x%2 == 0
})
// hasEven is true
```

### Delete

`func Delete[S ~[]E, E any](s S, i, j int) S`

Removes elements from the slice between indices i and j.

**Example:**

```go
numbers := []int{1, 2, 3, 4, 5}
numbers = Delete(numbers, 1, 3)
// numbers is []int{1, 4, 5}
```

### DeleteFunc

`func DeleteFunc[S ~[]E, E any](s S, del func(E) bool) S`

Removes elements that satisfy a predicate function.

**Example:**

```go
numbers := []int{1, 2, 3, 4, 5}
numbers = DeleteFunc(numbers, func(x int) bool {
    return x%2 == 0
})
// numbers is []int{1, 3, 5}
```

### Equal

`func Equal[S ~[]E, E comparable](s1, s2 S) bool`

Checks if two slices are equal.

**Example:**

```go
a := []int{1, 2, 3}
b := []int{1, 2, 3}
equal := Equal(a, b)
// equal is true
```

### EqualFunc

`func EqualFunc[S1 ~[]E1, S2 ~[]E2, E1, E2 any](s1 S1, s2 S2, eq func(E1, E2) bool) bool`

Checks if two slices are equal using a custom equality function.

**Example:**

```go
type Person struct {
    Name string
    Age  int
}
a := []Person{{"Alice", 25}}
b := []Person{{"alice", 25}}
equal := EqualFunc(a, b, func(p1, p2 Person) bool {
    return p1.Name == p2.Name && p1.Age == p2.Age
})
// equal is false
```

### Grow

`func Grow[S ~[]E, E any](s S, n int) S`

Increases the capacity of the slice by n.

**Example:**

```go
s := make([]int, 0, 5)
s = Grow(s, 10)
// cap(s) is at least 15
```

### Index

`func Index[S ~[]E, E comparable](s S, v E) int`

Returns the index of the first occurrence of v in s, or -1 if not present.

**Example:**

```go
numbers := []int{1, 2, 3, 2, 4}
index := Index(numbers, 2)
// index is 1
```

### IndexFunc

`func IndexFunc[S ~[]E, E any](s S, f func(E) bool) int`

Returns the index of the first element satisfying f, or -1 if none do.

**Example:**

```go
numbers := []int{1, 2, 3, 4, 5}
index := IndexFunc(numbers, func(x int) bool {
    return x%2 == 0
})
// index is 1 (first even number)
```

### Insert

`func Insert[S ~[]E, E any](s S, i int, v ...E) S`

Inserts the values v... into s at index i.

**Example:**

```go
numbers := []int{1, 2, 5}
numbers = Insert(numbers, 2, 3, 4)
// numbers is []int{1, 2, 3, 4, 5}
```

### IsSorted

`func IsSorted[S ~[]E, E cmp.Ordered](x S) bool`

Reports whether x is sorted in ascending order.

**Example:**

```go
numbers := []int{1, 2, 3, 4, 5}
sorted := IsSorted(numbers)
// sorted is true
```

### IsSortedFunc

`func IsSortedFunc[S ~[]E, E any](x S, cmp func(a, b E) int) bool`

Reports whether x is sorted in ascending order, as determined by cmp.

**Example:**

```go
type Person struct {
    Name string
    Age  int
}
people := []Person{
    {"Alice", 25},
    {"Bob", 30},
    {"Charlie", 35},
}
sorted := IsSortedFunc(people, func(a, b Person) int {
    return a.Age - b.Age
})
// sorted is true
```

### Max

`func Max[S ~[]E, E cmp.Ordered](x S) E`

Returns the maximum element in x. It panics if x is empty.

**Example:**

```go
numbers := []int{1, 4, 2, 5, 3}
max := Max(numbers)
// max is 5
```

### MaxFunc

`func MaxFunc[S ~[]E, E any](x S, cmp func(a, b E) int) E`

Returns the maximum element in x using cmp. It panics if x is empty.

**Example:**

```go
type Person struct {
    Name string
    Age  int
}
people := []Person{
    {"Alice", 25},
    {"Bob", 30},
    {"Charlie", 35},
}
oldest := MaxFunc(people, func(a, b Person) int {
    return a.Age - b.Age
})
// oldest is Person{"Charlie", 35}
```

### Min

`func Min[S ~[]E, E cmp.Ordered](x S) E`

Returns the minimum element in x. It panics if x is empty.

**Example:**

```go
numbers := []int{4, 2, 5, 1, 3}
min := Min(numbers)
// min is 1
```

### MinFunc

`func MinFunc[S ~[]E, E any](x S, cmp func(a, b E) int) E`

Returns the minimum element in x using cmp. It panics if x is empty.

**Example:**

```go
type Person struct {
    Name string
    Age  int
}
people := []Person{
    {"Alice", 25},
    {"Bob", 30},
    {"Charlie", 35},
}
youngest := MinFunc(people, func(a, b Person) int {
    return a.Age - b.Age
})
// youngest is Person{"Alice", 25}
```

### Replace

`func Replace[S ~[]E, E any](s S, i, j int, v ...E) S`

Replaces the elements from index i to j with v.

**Example:**

```go
numbers := []int{1, 2, 3, 4, 5}
numbers = Replace(numbers, 1, 4, 6, 7)
// numbers is []int{1, 6, 7, 5}
```

### Reverse

`func Reverse[S ~[]E, E any](s S)`

Reverses the elements of the slice in place.

**Example:**

```go
numbers := []int{1, 2, 3, 4, 5}
Reverse(numbers)
// numbers is []int{5, 4, 3, 2, 1}
```

### Sort

`func Sort[S ~[]E, E cmp.Ordered](x S)`

Sorts the slice in ascending order.

**Example:**

```go
numbers := []int{3, 1, 4, 1, 5, 9}
Sort(numbers)
// numbers is []int{1, 1, 3, 4, 5, 9}
```

### SortFunc

`func SortFunc[S ~[]E, E any](x S, cmp func(a, b E) int)`

Sorts the slice using a custom comparison function.

**Example:**

```go
type Person struct {
    Name string
    Age  int
}
people := []Person{
    {"Charlie", 35},
    {"Alice", 25},
    {"Bob", 30},
}
SortFunc(people, func(a, b Person) int {
    return strings.Compare(a.Name, b.Name)
})
// people is sorted by name
```

### SortStableFunc

`func SortStableFunc[S ~[]E, E any](x S, cmp func(a, b E) int)`

Sorts the slice using a custom comparison function while keeping the original order of equal elements (stable sort).

**Example:**

```go
type Person struct {
    Name    string
    Age     int
    OrderID int
}
people := []Person{
    {"Alice", 25, 1},
    {"Bob", 25, 2},
    {"Charlie", 35, 3},
}
SortStableFunc(people, func(a, b Person) int {
    return a.Age - b.Age
})
// people is sorted by age, with original order preserved for same ages
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

```

### Explanation of the README:
- Each function is documented with its signature, a brief description, and an example of how to use it.
- The examples demonstrate typical use cases for each function, making it easier for users to understand how to implement them in their own code.

Feel free to modify the README further to suit your project's needs! If you have any other requests or need additional examples, let me know!
```
