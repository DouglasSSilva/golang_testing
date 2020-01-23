package services

import "golang_testing/src/api/utils/sort"

// Sort function to make integration test
func Sort(elements []int) {
	if len(elements) > 10000 {
		sort.Sort(elements)
		return
	}
	sort.BubbleSort(elements)
}
