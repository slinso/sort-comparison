package sortcomparison

/*
ExchangeSort Implementation

Time Complexity:
  - Average: O(n²)
  - Worst:   O(n²) - when array is reverse sorted
  - Best:    O(n²) - even with sorted input

Space Complexity:
  - O(1) - uses only a single temp variable for swapping
  - In-place sorting algorithm

Implementation Notes:
  - Simple comparison-based sorting algorithm
  - Similar to selection sort but with immediate swaps
  - Not stable - does not preserve order of equal elements
  - Cache-friendly due to sequential access pattern
  - Good for small arrays or educational purposes
*/
func ExchangeSort(arr []int) {
	n := len(arr)
	if n <= 1 {
		return
	}

	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
}
