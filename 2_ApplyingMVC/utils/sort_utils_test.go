package utils

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func TestBubbleSortWorstCase(t *testing.T) {
	// Initialization
	elements := []int{9,8,7,6,5}

	// Execution
	BubbleSort(elements)

	// Validation
	assert.NotNil(t, elements)
	assert.EqualValues(t, 5, len(elements))
	assert.EqualValues(t, 5, elements[0])
	assert.EqualValues(t, 6, elements[1])
	assert.EqualValues(t, 7, elements[2])
	assert.EqualValues(t, 8, elements[3])
	assert.EqualValues(t, 9, elements[4])
}

func TestBubbleSortBestCase(t *testing.T) {
	// Initialization
	elements := []int{5,6,7,8,9}

	// Execution
	BubbleSort(elements)

	// Validation
	assert.NotNil(t, elements)
	assert.EqualValues(t, 5, len(elements))
	assert.EqualValues(t, 5, elements[0])
	assert.EqualValues(t, 6, elements[1])
	assert.EqualValues(t, 7, elements[2])
	assert.EqualValues(t, 8, elements[3])
	assert.EqualValues(t, 9, elements[4])
}

func TestSortWorstCase(t *testing.T) {
	// Initialization
	elements := []int{9,8,7,6,5}

	// Execution
	Sort(elements)

	// Validation
	assert.NotNil(t, elements)
	assert.EqualValues(t, 5, len(elements))
	assert.EqualValues(t, 5, elements[0])
	assert.EqualValues(t, 6, elements[1])
	assert.EqualValues(t, 7, elements[2])
	assert.EqualValues(t, 8, elements[3])
	assert.EqualValues(t, 9, elements[4])
}

func TestSortBestCase(t *testing.T) {
	// Initialization
	elements := []int{5,6,7,8,9}

	// Execution
	Sort(elements)

	// Validation
	assert.NotNil(t, elements)
	assert.EqualValues(t, 5, len(elements))
	assert.EqualValues(t, 5, elements[0])
	assert.EqualValues(t, 6, elements[1])
	assert.EqualValues(t, 7, elements[2])
	assert.EqualValues(t, 8, elements[3])
	assert.EqualValues(t, 9, elements[4])
}

func getElements(n int) []int {
	result := make([]int, n)
	i := 0
	for j := n-1; j>=0; j-- {
		result[i] = j
		i++
	}
	return result
}

func TestGetElements(t *testing.T) {
	elements := getElements(5)

	assert.NotNil(t, elements)
	assert.EqualValues(t, 5, len(elements))
	assert.EqualValues(t, 4, elements[0])
	assert.EqualValues(t, 3, elements[1])
	assert.EqualValues(t, 2, elements[2])
	assert.EqualValues(t, 1, elements[3])
	assert.EqualValues(t, 0, elements[4])
}

func benchmarkBubbleSortNElements(n int, b *testing.B) {
	elements := getElements(n)
	for i := 0; i < b.N; i++ {
		BubbleSort(elements)
	}
}

func BenchmarkBubbleSort10Elements(b *testing.B) {
	benchmarkBubbleSortNElements(10, b)
}

func BenchmarkBubbleSort1000Elements(b *testing.B) {
	benchmarkBubbleSortNElements(1000, b)
}

func BenchmarkBubbleSort100000Elements(b *testing.B) {
	benchmarkBubbleSortNElements(100000, b)
}

func benchmarkGoSortNativeLibrary(n int, b *testing.B){
	elements := getElements(n)
	for i := 0; i < b.N; i++ {
		sort.Ints(elements)
	}
}

func BenchmarkGoSort10Elements(b *testing.B) {
	benchmarkGoSortNativeLibrary(10, b)
}

func BenchmarkGoSort1000Elements(b *testing.B) {
	benchmarkGoSortNativeLibrary(1000, b)
}

func BenchmarkGoSort100000Elements(b *testing.B) {
	benchmarkGoSortNativeLibrary(100000, b)
}

func benchmarkSort(n int, b *testing.B){
	elements := getElements(n)
	for i := 0; i < b.N; i++ {
		Sort(elements)
	}
}

func BenchmarkSort10Elements(b *testing.B) {
	benchmarkSort(10, b)
}

func BenchmarkSort1000Elements(b *testing.B) {
	benchmarkSort(1000, b)
}

func BenchmarkSort100000Elements(b *testing.B) {
	benchmarkSort(100000, b)
}




