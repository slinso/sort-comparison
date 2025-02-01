package sortcomparison

/*
SelectionSort Implementation

Time Complexity:
  - Average: O(n²)
  - Worst:   O(n²)
  - Best:    O(n²)

Space Complexity:
  - O(1) - only requires a single variable for swapping
  - In-place sorting algorithm

Implementation Notes:
  - Simple comparison-based sorting algorithm
  - Not stable - does not preserve order of equal elements
  - Memory writes are O(n) - optimal for write-sensitive memory
  - Cache-unfriendly due to non-local memory access
  - Suitable for small arrays or when memory writes are expensive
*/
func SelectionSort(arr []int) {
	n := len(arr)
	if n <= 1 {
		return
	}

	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}
		// Swap only if we found a smaller element
		if minIdx != i {
			arr[i], arr[minIdx] = arr[minIdx], arr[i]
		}
	}
}
