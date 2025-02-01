package sortcomparison

/*
OddEvenSort Implementation (Brick Sort)

Time Complexity:
  - Average: O(n²)
  - Worst:   O(n²)
  - Best:    O(n) - when array is already sorted

Space Complexity:
  - O(1) - only requires a single boolean flag
  - In-place sorting algorithm

Implementation Notes:
  - Parallel-friendly sorting algorithm
  - Stable sort - maintains relative order of equal elements
  - Similar to bubble sort but with odd-even indexing pattern
  - Good for parallel processing architectures
  - Cache-friendly due to sequential access pattern
*/
func OddEvenSort(arr []int) {
	n := len(arr)
	if n <= 1 {
		return
	}

	sorted := false
	for !sorted {
		sorted = true

		// Even indices (0, 2, 4, ...)
		for i := 0; i < n-1; i += 2 {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				sorted = false
			}
		}

		// Odd indices (1, 3, 5, ...)
		for i := 1; i < n-1; i += 2 {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				sorted = false
			}
		}
	}
}
