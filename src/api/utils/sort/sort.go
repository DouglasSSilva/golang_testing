package sort

import (
	"sort"
)

// BubbleSort to be tested
func BubbleSort(elements []int) {
	keepWorking := true
	elen := len(elements)
	for keepWorking {
		keepWorking = false

		for i := 0; i < elen-1; i++ {
			if elements[i] > elements[i+1] {
				keepWorking = true
				elements[i], elements[i+1] = elements[i+1], elements[i]
			}
		}
	}
}

// Sort based on quicksort
func Sort(elements []int) {
	sort.Ints(elements)
}
