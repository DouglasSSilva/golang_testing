package services_test

import (
	. "golang_testing/src/api/services"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortTenThousandOrLess(t *testing.T) {
	// Init
	n := 1000
	elements := getElements(n)
	// Execution
	Sort(elements)
	// Validation
	assert.NotNil(t, elements, "Elements should not bil nil")
	assert.Equal(t, elements[0], 1, "First element should be equal to 1")
	assert.Equal(t, elements[len(elements)-1], n, "Last element should be equal to n")
	assert.EqualValues(t, len(elements), n, "the element should have size = n ")
}

func TestSortMoreThanTenThousand(t *testing.T) {
	// Init
	n := 100000
	elements := getElements(n)
	// Execution
	Sort(elements)
	// Validation
	assert.NotNil(t, elements, "Elements should not bil nil")
	assert.Equal(t, elements[0], 1, "First element should be equal to 1")
	assert.Equal(t, elements[len(elements)-1], n, "Last element should be equal to n")
	assert.EqualValues(t, len(elements), n, "the element should have size = n ")
}

func getElements(n int) []int {
	result := make([]int, n)
	for i := n; i > 0; i-- {
		result[i-1] = i
	}

	return result
}
