package sortcomparison

/*
ShellSort Implementation

Time Complexity:
  - Average: O(n log² n)
  - Worst:   O(n²) - depends on gap sequence
  - Best:    O(n log n) - when array is nearly sorted

Space Complexity:
  - O(1) - only requires variables for gap and temp storage
  - In-place sorting algorithm

Implementation Notes:
  - Uses Ciura's gap sequence for optimal performance
  - Adaptive - efficient for partially sorted arrays
  - Not stable - does not preserve order of equal elements
  - Cache-friendly compared to standard insertion sort
  - Good for medium-sized arrays (1000-50000 elements)
*/
func ShellSort(arr []int) {
	n := len(arr)
	if n <= 1 {
		return
	}

	// Ciura's gap sequence
	gaps := []int{701, 301, 132, 57, 23, 10, 4, 1}

	// Start with the largest gap and work down to gap of 1
	for _, gap := range gaps {
		// Skip gaps larger than array
		if gap >= n {
			continue
		}

		// Do a gapped insertion sort
		for i := gap; i < n; i++ {
			temp := arr[i]
			j := i

			// Shift elements until correct location found
			for j >= gap && arr[j-gap] > temp {
				arr[j] = arr[j-gap]
				j -= gap
			}

			arr[j] = temp
		}
	}
}
