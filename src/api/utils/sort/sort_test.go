package sort_test

import (
	. "golang_testing/src/api/utils/sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBubbleSortOrderASC(t *testing.T) {
	// Init
	elements := []int{9, 7, 5, 3, 1, 2, 4, 6, 8, 0}
	// Execution
	BubbleSort(elements)
	// Validation
	assert.Equal(t, elements[0], 0)
	assert.Equal(t, elements[len(elements)-1], 9)

}

func TestSortOrderASC(t *testing.T) {
	// Init
	elements := []int{9, 7, 5, 3, 1, 2, 4, 6, 8, 0}
	// Execution
	Sort(elements)
	// Validation
	assert.Equal(t, elements[0], 0)
	assert.Equal(t, elements[len(elements)-1], 9)

}

func BenchmarkBubbleSort(b *testing.B) {
	// Init

	elements := getElements(10000)
	// Execution
	for i := 0; i < b.N; i++ {
		BubbleSort(elements)
	}
}

func BenchmarkSort(b *testing.B) {
	// Init

	elements := getElements(10000)
	// Execution
	for i := 0; i < b.N; i++ {
		Sort(elements)
	}
}

func getElements(n int) []int {
	result := make([]int, n)
	for i := n; i > 0; i-- {
		result[i-1] = i
	}

	return result
}
