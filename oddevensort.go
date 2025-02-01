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
func OddEvenSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	sorted := false
	for !sorted {
		sorted = true

		// Odd phase
		for i := 1; i < len(arr)-1; i += 2 {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				sorted = false
			}
		}

		// Even phase
		for i := 0; i < len(arr)-1; i += 2 {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				sorted = false
			}
		}
	}

	return arr
}
