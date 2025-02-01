package sortcomparison

/*
CombSort Implementation

Time Complexity:
  - Average: O(n log n)
  - Worst:   O(n²)
  - Best:    O(n log n) - when array is nearly sorted

Space Complexity:
  - O(1) - only requires variables for gap and swapping
  - In-place sorting algorithm

Implementation Notes:
  - Improvement over bubble sort using gap sequence
  - Gap shrinks by factor of 1.3 (found empirically)
  - Eliminates small values near end of array ("turtle" values)
  - Not stable - does not preserve order of equal elements
  - Cache-friendly for initial passes with large gaps
*/
func CombSort(arr []int) {
	n := len(arr)
	if n <= 1 {
		return
	}

	// Initialize gap
	gap := n
	shrink := 1.3
	sorted := false

	for !sorted {
		// Update gap
		gap = int(float64(gap) / shrink)
		if gap <= 1 {
			gap = 1
			sorted = true
		}

		// Compare elements with gap
		for i := 0; i < n-gap; i++ {
			if arr[i] > arr[i+gap] {
				arr[i], arr[i+gap] = arr[i+gap], arr[i]
				sorted = false
			}
		}
	}
}
