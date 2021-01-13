package utils

import "sort"

// Input: {9,8,7,6,5}
// Expected output: {5,6,7,8,9}
func BubbleSort(elements []int ) {
	keepRunning := true

	for keepRunning {
		keepRunning = false

		for i := 0; i< len(elements) -1; i++ {
			if elements[i] > elements[i+1] {
				elements[i], elements[i+1] = elements[i+1], elements[i]
				keepRunning = true
			}
		}
	}
}

func Sort(elements []int) {
	if len(elements) < 1000 {
		BubbleSort(elements)
	}
	sort.Ints(elements)
}