package sortcomparison

/*
CountingSort Implementation (Non-Comparative Integer Sort)

Time Complexity:
  - Average: O(n + k) where k is the range of input
  - Worst:   O(n + k)
  - Best:    O(n + k)

Space Complexity:
  - O(k) where k is the range of input (max - min + 1)
  - Requires count array of size k

Implementation Notes:
  - Stable sort - maintains relative order of equal elements
  - Efficient for small ranges of integers
  - Non-comparison based algorithm
  - Linear time complexity when k = O(n)
  - Cache-friendly sequential access pattern
*/
func CountingSort(arr []int) {
	n := len(arr)
	if n <= 1 {
		return
	}

	// Find min and max in single pass
	min, max := arr[0], arr[0]
	for i := 1; i < n; i++ {
		if arr[i] < min {
			min = arr[i]
		}
		if arr[i] > max {
			max = arr[i]
		}
	}

	range_ := max - min + 1

	// Use original array if range is too large
	if range_ > 10*n {
		return
	}

	// Create count array
	count := make([]int, range_)

	// Count occurrences
	for i := 0; i < n; i++ {
		count[arr[i]-min]++
	}

	// Calculate cumulative count
	for i := 1; i < range_; i++ {
		count[i] += count[i-1]
	}

	// Build output array
	output := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		output[count[arr[i]-min]-1] = arr[i]
		count[arr[i]-min]--
	}

	// Copy back to original array
	copy(arr, output)
}
