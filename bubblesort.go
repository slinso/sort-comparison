package sortcomparison

// BubbleSort implements the bubble sort algorithm which:
// 1. Repeatedly steps through the list
// 2. Compares adjacent elements and swaps them if they are in wrong order
// 3. Continues until no swaps are needed (list is sorted)
// Time Complexity: O(n²)
// Space Complexity: O(1)
// Stable: Yes
// In-place: Yes
// Best case: O(n) when array is already sorted

func BubbleSort[T any](arr []T, less func(a, b T) bool) {
	n := len(arr)
	if n <= 1 {
		return
	}

	for i := 0; i < n-1; i++ {
		swapped := false

		// Last i elements are already in place
		for j := 0; j < n-i-1; j++ {
			// Compare adjacent elements
			if less(arr[j+1], arr[j]) {
				// Swap if they are in wrong order
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
			}
		}

		// If no swapping occurred, array is sorted
		if !swapped {
			break
		}
	}
}
