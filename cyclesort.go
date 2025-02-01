package sortcomparison

/*
CycleSort Implementation (Optimal Memory Writes)

Time Complexity:
  - Average: O(n²)
  - Worst:   O(n²)
  - Best:    O(n²)

Space Complexity:
  - O(1) - only requires a single variable for swapping
  - Minimum memory writes: exactly n writes to sort n elements

Implementation Notes:
  - In-place sorting algorithm
  - Not stable - does not preserve order of equal elements
  - Optimal for minimizing memory writes
  - Each element is moved exactly once to its final position
  - Good for situations where memory write is costly
*/
func CycleSort(arr []int) {
	n := len(arr)
	if n <= 1 {
		return
	}

	for cycleStart := 0; cycleStart < n-1; cycleStart++ {
		item := arr[cycleStart]
		pos := cycleStart

		// Find final position for item
		for i := cycleStart + 1; i < n; i++ {
			if arr[i] < item {
				pos++
			}
		}

		// Skip if item is already in correct position
		if pos == cycleStart {
			continue
		}

		// Handle duplicates by finding correct position
		for item == arr[pos] {
			pos++
		}

		// Put item in its final position and get displaced item
		if pos != cycleStart {
			arr[pos], item = item, arr[pos]
		}

		// Rotate remainder of cycle
		for pos != cycleStart {
			pos = cycleStart

			// Find final position for current item
			for i := cycleStart + 1; i < n; i++ {
				if arr[i] < item {
					pos++
				}
			}

			// Handle duplicates
			for item == arr[pos] {
				pos++
			}

			// Put item in its final position and get next item
			if item != arr[pos] {
				arr[pos], item = item, arr[pos]
			}
		}
	}
}
