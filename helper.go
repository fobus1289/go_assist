package goassist

import (
	"cmp"
	"slices"
)

// Map applies a function to each element of the input slice and returns a new slice with the results.
//
// Example:
//
//	numbers := []int{1, 2, 3, 4, 5}
//	doubled := Map(numbers, func(x int) int {
//		return x * 2
//	})
//	// doubled is []int{2, 4, 6, 8, 10}
//
//	names := []string{"alice", "bob", "charlie"}
//	lengths := Map(names, func(s string) int {
//		return len(s)
//	})
//	// lengths is []int{5, 3, 7}
func Map[T any, R any](arr []T, fn func(T) R) []R {
	result := make([]R, len(arr))
	for i, v := range arr {
		result[i] = fn(v)
	}
	return result
}

// Filter returns a new slice containing only the elements that satisfy the provided predicate function.
//
// Example:
//
//	numbers := []int{1, 2, 3, 4, 5}
//	evens := Filter(numbers, func(x int) bool {
//		return x%2 == 0
//	})
//	// evens is []int{2, 4}
//
//	names := []string{"alice", "bob", "charlie"}
//	longNames := Filter(names, func(s string) bool {
//		return len(s) > 4
//	})
//	// longNames is []string{"alice", "charlie"}
func Filter[T any](arr []T, fn func(T) bool) []T {
	result := make([]T, 0)
	for _, v := range arr {
		if fn(v) {
			result = append(result, v)
		}
	}
	return result
}

// Reduce applies a function cumulatively to the elements of the slice, reducing it to a single value.
//
// Example:
//
//	numbers := []int{1, 2, 3, 4, 5}
//	sum := Reduce(numbers, func(acc, x int) int {
//		return acc + x
//	}, 0)
//	// sum is 15
//
//	words := []string{"hello", " ", "world"}
//	sentence := Reduce(words, func(acc, s string) string {
//		return acc + s
//	}, "")
//	// sentence is "hello world"
func Reduce[T any, R any](arr []T, fn func(R, T) R, initial R) R {
	result := initial
	for _, v := range arr {
		result = fn(result, v)
	}
	return result
}

// Find returns the first element that satisfies the predicate function and a boolean indicating success.
//
// Example:
//
//	numbers := []int{1, 2, 3, 4, 5}
//	first, found := Find(numbers, func(x int) bool {
//		return x > 3
//	})
//	// first is 4, found is true
//
//	noMatch, found := Find(numbers, func(x int) bool {
//		return x > 10
//	})
//	// noMatch is 0, found is false
func Find[T any](arr []T, fn func(T) bool) (T, bool) {
	for _, v := range arr {
		if fn(v) {
			return v, true
		}
	}

	var zero T

	return zero, false
}

// Some returns true if at least one element satisfies the predicate function.
//
// Example:
//
//	numbers := []int{1, 2, 3, 4, 5}
//	hasEven := Some(numbers, func(x int) bool {
//		return x%2 == 0
//	})
//	// hasEven is true
//
//	hasNegative := Some(numbers, func(x int) bool {
//		return x < 0
//	})
//	// hasNegative is false
func Some[T any](arr []T, fn func(T) bool) bool {
	for _, v := range arr {
		if fn(v) {
			return true
		}
	}
	return false
}

// Every returns true if all elements satisfy the predicate function.
//
// Example:
//
//	numbers := []int{2, 4, 6, 8}
//	allEven := Every(numbers, func(x int) bool {
//		return x%2 == 0
//	})
//	// allEven is true
//
//	mixed := []int{2, 4, 5, 8}
//	allEven = Every(mixed, func(x int) bool {
//		return x%2 == 0
//	})
//	// allEven is false
func Every[T any](arr []T, fn func(T) bool) bool {
	for _, v := range arr {
		if !fn(v) {
			return false
		}
	}
	return true
}

// Flatten flattens a slice of slices into a single slice.
//
// Example:
//
//	nested := [][]int{{1, 2}, {3, 4}, {5, 6}}
//	flat := Flatten(nested)
//	// flat is []int{1, 2, 3, 4, 5, 6}
//
//	matrix := [][]string{{"a", "b"}, {"c", "d"}}
//	chars := Flatten(matrix)
//	// chars is []string{"a", "b", "c", "d"}
func Flatten[T any](arr [][]T) []T {
	result := make([]T, 0)
	for _, v := range arr {
		result = append(result, v...)
	}
	return result
}

// Zip combines two slices into a slice of pairs. If the input slices have different lengths,
// the result will have the length of the shorter slice.
//
// Example:
//
//	numbers := []int{1, 2, 3}
//	letters := []int{'a', 'b', 'c'}
//	pairs := Zip(numbers, letters)
//	// pairs is [][]int{{1, 'a'}, {2, 'b'}, {3, 'c'}}
func Zip[T any, R any](arr []T, arr2 []R) [][]any {
	result := make([][]any, 0)
	for i := range arr {
		result = append(result, []any{arr[i], arr2[i]})
	}
	return result
}

// Unzip splits a slice of pairs into two separate slices.
//
// Example:
//
//	pairs := [][]int{{1, 10}, {2, 20}, {3, 30}}
//	first, second := Unzip[int, int](pairs)
//	// first is []int{1, 2, 3}
//	// second is []int{10, 20, 30}
func Unzip[T any, R any](arr [][]any) ([]T, []R) {

	result := make([]T, 0)
	result2 := make([]R, 0)
	for _, v := range arr {
		result = append(result, v[0].(T))
		result2 = append(result2, v[1].(R))
	}
	return result, result2
}

// BinarySearch performs a binary search on a sorted slice and returns the index and a boolean indicating if the target was found.
// The slice must be sorted in increasing order.
//
// Example:
//
//	numbers := []int{1, 3, 5, 7, 9}
//	index, found := BinarySearch(numbers, 5)
//	// index is 2, found is true
//	index, found = BinarySearch(numbers, 6)
//	// index is 3, found is false
func BinarySearch[S ~[]E, E cmp.Ordered](x S, target E) (int, bool) {
	return slices.BinarySearch(x, target)
}

// BinarySearchFunc performs a binary search using a custom comparison function.
// The slice must be sorted in increasing order according to the comparison function.
// The cmp function should return 0 if the values are equal, a negative number if a < b,
// or a positive number if a > b.
//
// Example:
//
//	type Person struct {
//		Name string
//		Age  int
//	}
//	people := []Person{
//		{"Alice", 25},
//		{"Bob", 30},
//		{"Charlie", 35},
//	}
//	index, found := BinarySearchFunc(people, 30, func(p Person, age int) int {
//		return p.Age - age
//	})
//	// index is 1, found is true
func BinarySearchFunc[S ~[]E, E, T any](x S, target T, cmp func(E, T) int) (int, bool) {
	return slices.BinarySearchFunc(x, target, cmp)
}

// Clip returns a slice with the capacity adjusted to the length of the slice.
// This can be used to release memory if a slice is known to not need additional growth.
//
// Example:
//
//	s := make([]int, 3, 10)
//	s = Clip(s)
//	// len(s) == 3, cap(s) == 3
func Clip[S ~[]E, E any](s S) S {
	return slices.Clip(s)
}

// Clone creates a copy of the slice.
// The elements are copied using assignment, so this is a shallow clone.
//
// Example:
//
//	original := []int{1, 2, 3}
//	copy := Clone(original)
//	copy[0] = 99
//	// original is still []int{1, 2, 3}
func Clone[S ~[]E, E any](s S) S {
	return slices.Clone(s)
}

// Compact removes adjacent duplicate elements from the slice.
//
// Example:
//
//	numbers := []int{1, 1, 2, 3, 3, 3, 4, 4, 5}
//	unique := Compact(numbers)
//	// unique is []int{1, 2, 3, 4, 5}
func Compact[S ~[]E, E comparable](s S) S {
	return slices.Compact(s)
}

// CompactFunc removes adjacent duplicate elements based on a custom equality function.
//
// Example:
//
//	type Person struct {
//		Name string
//		Age  int
//	}
//	people := []Person{
//		{"Alice", 25},
//		{"Alice", 30},
//		{"Bob", 30},
//	}
//	unique := CompactFunc(people, func(a, b Person) bool {
//		return a.Name == b.Name
//	})
//	// unique has 2 elements: Alice and Bob
func CompactFunc[S ~[]E, E any](s S, eq func(E, E) bool) S {
	return slices.CompactFunc(s, eq)
}

// Compare compares two slices lexicographically and returns:
//   - a negative number if s1 < s2
//   - zero if s1 == s2
//   - a positive number if s1 > s2
//
// Example:
//
//	a := []int{1, 2, 3}
//	b := []int{1, 2, 4}
//	result := Compare(a, b)
//	// result is negative because a < b
func Compare[S ~[]E, E cmp.Ordered](s1, s2 S) int {
	return slices.Compare(s1, s2)
}

// CompareFunc compares two slices using a custom comparison function.
// The cmp function should return:
//   - a negative number if a < b
//   - zero if a == b
//   - a positive number if a > b
//
// Example:
//
//	type Person struct {
//		Name string
//		Age  int
//	}
//	a := []Person{{"Alice", 25}}
//	b := []Person{{"Bob", 30}}
//	result := CompareFunc(a, b, func(p1, p2 Person) int {
//		return strings.Compare(p1.Name, p2.Name)
//	})
//	// result is negative because "Alice" < "Bob"
func CompareFunc[S1 ~[]E1, S2 ~[]E2, E1, E2 any](s1 S1, s2 S2, cmp func(E1, E2) int) int {
	return slices.CompareFunc(s1, s2, cmp)
}

// Contains checks if a value exists in the slice.
//
// Example:
//
//	numbers := []int{1, 2, 3, 4, 5}
//	exists := Contains(numbers, 3)
//	// exists is true
//	exists = Contains(numbers, 6)
//	// exists is false
func Contains[S ~[]E, E comparable](s S, v E) bool {
	return slices.Contains(s, v)
}

// ContainsFunc checks if any element satisfies a predicate function.
//
// Example:
//
//	numbers := []int{1, 2, 3, 4, 5}
//	hasEven := ContainsFunc(numbers, func(x int) bool {
//		return x%2 == 0
//	})
//	// hasEven is true
func ContainsFunc[S ~[]E, E any](s S, f func(E) bool) bool {
	return slices.ContainsFunc(s, f)
}

// Delete removes elements from the slice between indices i and j.
//
// Example:
//
//	numbers := []int{1, 2, 3, 4, 5}
//	numbers = Delete(numbers, 1, 3)
//	// numbers is []int{1, 4, 5}
func Delete[S ~[]E, E any](s S, i, j int) S {
	return slices.Delete(s, i, j)
}

// DeleteFunc removes elements that satisfy a predicate function.
//
// Example:
//
//	numbers := []int{1, 2, 3, 4, 5}
//	numbers = DeleteFunc(numbers, func(x int) bool {
//		return x%2 == 0
//	})
//	// numbers is []int{1, 3, 5}
func DeleteFunc[S ~[]E, E any](s S, del func(E) bool) S {
	return slices.DeleteFunc(s, del)
}

// Equal checks if two slices are equal.
//
// Example:
//
//	a := []int{1, 2, 3}
//	b := []int{1, 2, 3}
//	equal := Equal(a, b)
//	// equal is true
func Equal[S ~[]E, E comparable](s1, s2 S) bool {
	return slices.Equal(s1, s2)
}

// EqualFunc checks if two slices are equal using a custom equality function.
//
// Example:
//
//	type Person struct {
//		Name string
//		Age  int
//	}
//	a := []Person{{"Alice", 25}}
//	b := []Person{{"alice", 25}}
//	equal := EqualFunc(a, b, func(p1, p2 Person) bool {
//		return strings.EqualFold(p1.Name, p2.Name) && p1.Age == p2.Age
//	})
//	// equal is true
func EqualFunc[S1 ~[]E1, S2 ~[]E2, E1, E2 any](s1 S1, s2 S2, eq func(E1, E2) bool) bool {
	return slices.EqualFunc(s1, s2, eq)
}

// Grow increases the capacity of the slice by n.
//
// Example:
//
//	s := make([]int, 0, 5)
//	s = Grow(s, 10)
//	// cap(s) is at least 15
func Grow[S ~[]E, E any](s S, n int) S {
	return slices.Grow(s, n)
}

// Index returns the index of the first occurrence of v in s, or -1 if not present.
//
// Example:
//
//	numbers := []int{1, 2, 3, 2, 4}
//	index := Index(numbers, 2)
//	// index is 1
func Index[S ~[]E, E comparable](s S, v E) int {
	return slices.Index(s, v)
}

// IndexFunc returns the index of the first element satisfying f, or -1 if none do.
//
// Example:
//
//	numbers := []int{1, 2, 3, 4, 5}
//	index := IndexFunc(numbers, func(x int) bool {
//		return x%2 == 0
//	})
//	// index is 1 (first even number)
func IndexFunc[S ~[]E, E any](s S, f func(E) bool) int {
	return slices.IndexFunc(s, f)
}

// Insert inserts the values v... into s at index i.
//
// Example:
//
//	numbers := []int{1, 2, 5}
//	numbers = Insert(numbers, 2, 3, 4)
//	// numbers is []int{1, 2, 3, 4, 5}
func Insert[S ~[]E, E any](s S, i int, v ...E) S {
	return slices.Insert(s, i, v...)
}

// IsSorted reports whether x is sorted in ascending order.
//
// Example:
//
//	numbers := []int{1, 2, 3, 4, 5}
//	sorted := IsSorted(numbers)
//	// sorted is true
func IsSorted[S ~[]E, E cmp.Ordered](x S) bool {
	return slices.IsSorted(x)
}

// IsSortedFunc reports whether x is sorted in ascending order, as determined by cmp.
//
// Example:
//
//	type Person struct {
//		Name string
//		Age  int
//	}
//	people := []Person{
//		{"Alice", 25},
//		{"Bob", 30},
//		{"Charlie", 35},
//	}
//	sorted := IsSortedFunc(people, func(a, b Person) int {
//		return a.Age - b.Age
//	})
//	// sorted is true
func IsSortedFunc[S ~[]E, E any](x S, cmp func(a, b E) int) bool {
	return slices.IsSortedFunc(x, cmp)
}

// Max returns the maximum element in x. It panics if x is empty.
//
// Example:
//
//	numbers := []int{1, 4, 2, 5, 3}
//	max := Max(numbers)
//	// max is 5
func Max[S ~[]E, E cmp.Ordered](x S) E {
	return slices.Max(x)
}

// MaxFunc returns the maximum element in x using cmp. It panics if x is empty.
//
// Example:
//
//	type Person struct {
//		Name string
//		Age  int
//	}
//	people := []Person{
//		{"Alice", 25},
//		{"Bob", 30},
//		{"Charlie", 35},
//	}
//	oldest := MaxFunc(people, func(a, b Person) int {
//		return a.Age - b.Age
//	})
//	// oldest is Person{"Charlie", 35}
func MaxFunc[S ~[]E, E any](x S, cmp func(a, b E) int) E {
	return slices.MaxFunc(x, cmp)
}

// Min returns the minimum element in x. It panics if x is empty.
//
// Example:
//
//	numbers := []int{4, 2, 5, 1, 3}
//	min := Min(numbers)
//	// min is 1
func Min[S ~[]E, E cmp.Ordered](x S) E {
	return slices.Min(x)
}

// MinFunc returns the minimum element in x using cmp. It panics if x is empty.
//
// Example:
//
//	type Person struct {
//		Name string
//		Age  int
//	}
//	people := []Person{
//		{"Alice", 25},
//		{"Bob", 30},
//		{"Charlie", 35},
//	}
//	youngest := MinFunc(people, func(a, b Person) int {
//		return a.Age - b.Age
//	})
//	// youngest is Person{"Alice", 25}
func MinFunc[S ~[]E, E any](x S, cmp func(a, b E) int) E {
	return slices.MinFunc(x, cmp)
}

// Replace replaces the elements from index i to j with v.
//
// Example:
//
//	numbers := []int{1, 2, 3, 4, 5}
//	numbers = Replace(numbers, 1, 4, 6, 7)
//	// numbers is []int{1, 6, 7, 5}
func Replace[S ~[]E, E any](s S, i, j int, v ...E) S {
	return slices.Replace(s, i, j, v...)
}

// Reverse reverses the elements of the slice in place.
//
// Example:
//
//	numbers := []int{1, 2, 3, 4, 5}
//	Reverse(numbers)
//	// numbers is []int{5, 4, 3, 2, 1}
func Reverse[S ~[]E, E any](s S) {
	slices.Reverse(s)
}

// Sort sorts the slice in ascending order.
//
// Example:
//
//	numbers := []int{3, 1, 4, 1, 5, 9}
//	Sort(numbers)
//	// numbers is []int{1, 1, 3, 4, 5, 9}
func Sort[S ~[]E, E cmp.Ordered](x S) {
	slices.Sort(x)
}

// SortFunc sorts the slice using a custom comparison function.
//
// Example:
//
//	type Person struct {
//		Name string
//		Age  int
//	}
//	people := []Person{
//		{"Charlie", 35},
//		{"Alice", 25},
//		{"Bob", 30},
//	}
//	SortFunc(people, func(a, b Person) int {
//		return strings.Compare(a.Name, b.Name)
//	})
//	// people is sorted by name
func SortFunc[S ~[]E, E any](x S, cmp func(a, b E) int) {
	slices.SortFunc(x, cmp)
}

// SortStableFunc sorts the slice using a custom comparison function while keeping the
// original order of equal elements (stable sort).
//
// Example:
//
//	type Person struct {
//		Name    string
//		Age     int
//		OrderID int
//	}
//	people := []Person{
//		{"Alice", 25, 1},
//		{"Bob", 25, 2},
//		{"Charlie", 35, 3},
//	}
//	SortStableFunc(people, func(a, b Person) int {
//		return a.Age - b.Age
//	})
//	// people is sorted by age, with original order preserved for same ages
func SortStableFunc[S ~[]E, E any](x S, cmp func(a, b E) int) {
	slices.SortStableFunc(x, cmp)
}
