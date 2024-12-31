package goassist_test

import (
	"strings"
	"testing"

	goassist "github.com/fobus1289/go_assist"
)

func TestMap(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}
	doubled := goassist.Map(numbers, func(x int) int {
		return x * 2
	})
	expected := []int{2, 4, 6, 8, 10}
	for i, v := range doubled {
		if v != expected[i] {
			t.Errorf("Map failed: expected %d, got %d", expected[i], v)
		}
	}
}

func TestFilter(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}
	evens := goassist.Filter(numbers, func(x int) bool {
		return x%2 == 0
	})
	expected := []int{2, 4}
	for i, v := range evens {
		if v != expected[i] {
			t.Errorf("Filter failed: expected %d, got %d", expected[i], v)
		}
	}
}

func TestReduce(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}
	sum := goassist.Reduce(numbers, func(acc, x int) int {
		return acc + x
	}, 0)
	expected := 15
	if sum != expected {
		t.Errorf("Reduce failed: expected %d, got %d", expected, sum)
	}
}

func TestFind(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}
	first, found := goassist.Find(numbers, func(x int) bool {
		return x > 3
	})
	if !found || first != 4 {
		t.Errorf("Find failed: expected 4, got %d, found %v", first, found)
	}
}

func TestSome(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}
	hasEven := goassist.Some(numbers, func(x int) bool {
		return x%2 == 0
	})
	if !hasEven {
		t.Error("Some failed: expected true, got false")
	}
}

func TestEvery(t *testing.T) {
	numbers := []int{2, 4, 6, 8}
	allEven := goassist.Every(numbers, func(x int) bool {
		return x%2 == 0
	})
	if !allEven {
		t.Error("Every failed: expected true, got false")
	}
}

func TestFlatten(t *testing.T) {
	nested := [][]int{{1, 2}, {3, 4}, {5, 6}}
	flat := goassist.Flatten(nested)
	expected := []int{1, 2, 3, 4, 5, 6}
	for i, v := range flat {
		if v != expected[i] {
			t.Errorf("Flatten failed: expected %d, got %d", expected[i], v)
		}
	}
}

func TestZip(t *testing.T) {
	numbers := []int{1, 2, 3}
	letters := []string{"a", "b", "c"}
	pairs := goassist.Zip(numbers, letters)
	expected := [][]interface{}{{1, "a"}, {2, "b"}, {3, "c"}}
	for i, v := range pairs {
		if v[0] != expected[i][0] || v[1] != expected[i][1] {
			t.Errorf("Zip failed: expected %v, got %v", expected[i], v)
		}
	}
}

func TestUnzip(t *testing.T) {
	pairs := [][]interface{}{{1, 10}, {2, 20}, {3, 30}}
	first, second := goassist.Unzip[int, int](pairs)
	expectedFirst := []int{1, 2, 3}
	expectedSecond := []int{10, 20, 30}
	for i, v := range first {
		if v != expectedFirst[i] {
			t.Errorf("Unzip failed: expected %d, got %d", expectedFirst[i], v)
		}
	}
	for i, v := range second {
		if v != expectedSecond[i] {
			t.Errorf("Unzip failed: expected %d, got %d", expectedSecond[i], v)
		}
	}
}

func TestBinarySearch(t *testing.T) {
	numbers := []int{1, 3, 5, 7, 9}
	index, found := goassist.BinarySearch(numbers, 5)
	if !found || index != 2 {
		t.Errorf("BinarySearch failed: expected index 2, found %d, found %v", index, found)
	}
	index, found = goassist.BinarySearch(numbers, 6)
	if found || index != 3 {
		t.Errorf("BinarySearch failed: expected index 3, found %d, found %v", index, found)
	}
}

func TestBinarySearchFunc(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}
	people := []Person{
		{"Alice", 25},
		{"Bob", 30},
		{"Charlie", 35},
	}
	index, found := goassist.BinarySearchFunc(people, 30, func(p Person, age int) int {
		return p.Age - age
	})
	if !found || index != 1 {
		t.Errorf("BinarySearchFunc failed: expected index 1, found %d, found %v", index, found)
	}
	index, found = goassist.BinarySearchFunc(people, 40, func(p Person, age int) int {
		return p.Age - age
	})
	if found || index != 3 {
		t.Errorf("BinarySearchFunc failed: expected index 3, found %d, found %v", index, found)
	}
}

func TestClip(t *testing.T) {
	s := make([]int, 3, 10)
	s = goassist.Clip(s)
	if cap(s) != 3 {
		t.Errorf("Clip failed: expected capacity 3, got %d", cap(s))
	}
}

func TestClone(t *testing.T) {
	original := []int{1, 2, 3}
	copy := goassist.Clone(original)
	copy[0] = 99
	if original[0] == 99 {
		t.Error("Clone failed: original slice should not be modified")
	}
}

func TestCompact(t *testing.T) {
	numbers := []int{1, 1, 2, 3, 3, 3, 4, 4, 5}
	unique := goassist.Compact(numbers)
	expected := []int{1, 2, 3, 4, 5}
	for i, v := range unique {
		if v != expected[i] {
			t.Errorf("Compact failed: expected %d, got %d", expected[i], v)
		}
	}
}

func TestCompactFunc(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}
	people := []Person{
		{"Alice", 25},
		{"Alice", 30},
		{"Bob", 30},
	}
	unique := goassist.CompactFunc(people, func(a, b Person) bool {
		return a.Name == b.Name
	})
	if len(unique) != 2 {
		t.Errorf("CompactFunc failed: expected 2 unique people, got %d", len(unique))
	}
}

func TestContains(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}
	exists := goassist.Contains(numbers, 3)
	if !exists {
		t.Error("Contains failed: expected true, got false")
	}
	exists = goassist.Contains(numbers, 6)
	if exists {
		t.Error("Contains failed: expected false, got true")
	}
}

func TestContainsFunc(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}
	hasEven := goassist.ContainsFunc(numbers, func(x int) bool {
		return x%2 == 0
	})
	if !hasEven {
		t.Error("ContainsFunc failed: expected true, got false")
	}
}

func TestDelete(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}
	numbers = goassist.Delete(numbers, 1, 3)
	expected := []int{1, 4, 5}
	for i, v := range numbers {
		if v != expected[i] {
			t.Errorf("Delete failed: expected %d, got %d", expected[i], v)
		}
	}
}

func TestDeleteFunc(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}
	numbers = goassist.DeleteFunc(numbers, func(x int) bool {
		return x%2 == 0
	})
	expected := []int{1, 3, 5}
	for i, v := range numbers {
		if v != expected[i] {
			t.Errorf("DeleteFunc failed: expected %d, got %d", expected[i], v)
		}
	}
}

func TestEqual(t *testing.T) {
	a := []int{1, 2, 3}
	b := []int{1, 2, 3}
	equal := goassist.Equal(a, b)
	if !equal {
		t.Error("Equal failed: expected true, got false")
	}
}

func TestEqualFunc(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}
	a := []Person{{"Alice", 25}}
	b := []Person{{"alice", 25}}
	equal := goassist.EqualFunc(a, b, func(p1, p2 Person) bool {
		return p1.Name == p2.Name && p1.Age == p2.Age
	})
	if equal {
		t.Error("EqualFunc failed: expected false, got true")
	}
}

func TestGrow(t *testing.T) {
	s := make([]int, 0, 5)
	s = goassist.Grow(s, 10)
	if cap(s) < 10 {
		t.Errorf("Grow failed: expected capacity at least 10, got %d", cap(s))
	}
}

func TestIndex(t *testing.T) {
	numbers := []int{1, 2, 3, 2, 4}
	index := goassist.Index(numbers, 2)
	if index != 1 {
		t.Errorf("Index failed: expected 1, got %d", index)
	}
}

func TestIndexFunc(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}
	index := goassist.IndexFunc(numbers, func(x int) bool {
		return x%2 == 0
	})
	if index != 1 {
		t.Errorf("IndexFunc failed: expected 1, got %d", index)
	}
}

func TestInsert(t *testing.T) {
	numbers := []int{1, 2, 5}
	numbers = goassist.Insert(numbers, 2, 3, 4)
	expected := []int{1, 2, 3, 4, 5}
	for i, v := range numbers {
		if v != expected[i] {
			t.Errorf("Insert failed: expected %d, got %d", expected[i], v)
		}
	}
}

func TestIsSorted(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}
	sorted := goassist.IsSorted(numbers)
	if !sorted {
		t.Error("IsSorted failed: expected true, got false")
	}
}

func TestIsSortedFunc(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}
	people := []Person{
		{"Alice", 25},
		{"Bob", 30},
		{"Charlie", 35},
	}
	sorted := goassist.IsSortedFunc(people, func(a, b Person) int {
		return a.Age - b.Age
	})
	if !sorted {
		t.Error("IsSortedFunc failed: expected true, got false")
	}
}

func TestMax(t *testing.T) {
	numbers := []int{1, 4, 2, 5, 3}
	max := goassist.Max(numbers)
	if max != 5 {
		t.Errorf("Max failed: expected 5, got %d", max)
	}
}

func TestMaxFunc(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}
	people := []Person{
		{"Alice", 25},
		{"Bob", 30},
		{"Charlie", 35},
	}
	oldest := goassist.MaxFunc(people, func(a, b Person) int {
		return a.Age - b.Age
	})
	if oldest.Age != 35 {
		t.Errorf("MaxFunc failed: expected age 35, got %d", oldest.Age)
	}
}

func TestMin(t *testing.T) {
	numbers := []int{4, 2, 5, 1, 3}
	min := goassist.Min(numbers)
	if min != 1 {
		t.Errorf("Min failed: expected 1, got %d", min)
	}
}

func TestMinFunc(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}
	people := []Person{
		{"Alice", 25},
		{"Bob", 30},
		{"Charlie", 35},
	}
	youngest := goassist.MinFunc(people, func(a, b Person) int {
		return a.Age - b.Age
	})
	if youngest.Age != 25 {
		t.Errorf("MinFunc failed: expected age 25, got %d", youngest.Age)
	}
}

func TestReplace(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}
	numbers = goassist.Replace(numbers, 1, 4, 6, 7)
	expected := []int{1, 6, 7, 5}
	for i, v := range numbers {
		if v != expected[i] {
			t.Errorf("Replace failed: expected %d, got %d", expected[i], v)
		}
	}
}

func TestReverse(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}
	goassist.Reverse(numbers)
	expected := []int{5, 4, 3, 2, 1}
	for i, v := range numbers {
		if v != expected[i] {
			t.Errorf("Reverse failed: expected %d, got %d", expected[i], v)
		}
	}
}

func TestSort(t *testing.T) {
	numbers := []int{3, 1, 4, 1, 5, 9}
	goassist.Sort(numbers)
	expected := []int{1, 1, 3, 4, 5, 9}
	for i, v := range numbers {
		if v != expected[i] {
			t.Errorf("Sort failed: expected %d, got %d", expected[i], v)
		}
	}
}

func TestSortFunc(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}
	people := []Person{
		{"Charlie", 35},
		{"Alice", 25},
		{"Bob", 30},
	}
	goassist.SortFunc(people, func(a, b Person) int {
		return strings.Compare(a.Name, b.Name)
	})
	expected := []Person{
		{"Alice", 25},
		{"Bob", 30},
		{"Charlie", 35},
	}
	for i, v := range people {
		if v != expected[i] {
			t.Errorf("SortFunc failed: expected %v, got %v", expected[i], v)
		}
	}
}

func TestSortStableFunc(t *testing.T) {
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
	goassist.SortStableFunc(people, func(a, b Person) int {
		return a.Age - b.Age
	})
	expected := []Person{
		{"Alice", 25, 1},
		{"Bob", 25, 2},
		{"Charlie", 35, 3},
	}
	for i, v := range people {
		if v != expected[i] {
			t.Errorf("SortStableFunc failed: expected %v, got %v", expected[i], v)
		}
	}
}
