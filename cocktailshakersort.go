package sortcomparison

/*
CocktailShaker Sort Implementation (Bidirectional Bubble Sort)

Time Complexity:
  - Average: O(n²)
  - Worst:   O(n²) - when array is reverse sorted
  - Best:    O(n) - when array is already sorted

Space Complexity:
  - O(1) - only requires a single variable for swapping
  - In-place sorting algorithm

Implementation Notes:
  - Bidirectional variant of bubble sort
  - Stable sort - maintains relative order of equal elements
  - More efficient than bubble sort on partially sorted arrays
  - Reduces "turtle" problem (small values near end of array)
  - Cache-friendly due to sequential access pattern
*/
func CocktailShakerSort(arr []int) {
	n := len(arr)
	if n <= 1 {
		return
	}

	left := 0
	right := n - 1
	swapped := true

	for swapped {
		swapped = false

		// Forward pass (left to right)
		for i := left; i < right; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				swapped = true
			}
		}
		if !swapped {
			break
		}
		right--

		// Backward pass (right to left)
		swapped = false
		for i := right; i > left; i-- {
			if arr[i-1] > arr[i] {
				arr[i-1], arr[i] = arr[i], arr[i-1]
				swapped = true
			}
		}
		left++
	}
}
