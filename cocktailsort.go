package sortcomparison

// CocktailSort sorts a slice of integers in ascending order using the cocktail
// shaker sort algorithm. This algorithm traverses the array bidirectionally,
// bubbling the largest element to the end on the forward pass and the smallest
// element to the beginning on the backward pass, until no swaps occur.
//
// How it works:
//   - In the forward pass, it compares adjacent elements and swaps them if they
//     are in the wrong order.
//   - In the backward pass, it moves in reverse order and does the same.
//   - The process repeats until a complete forward and backward pass yields no swaps,
//     which means the array is sorted.
//
// Time Complexity:
//   - Worst-case: O(n²)
//   - Best-case: O(n) when the array is nearly sorted
//
// Space Complexity:
//   - O(1) auxiliary space (in-place)
func CocktailSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	swapped := true
	start := 0
	end := len(arr) - 1

	for swapped {
		swapped = false

		// Forward pass: bubble up the largest element.
		for i := start; i < end; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				swapped = true
			}
		}

		// If no swaps occurred, array is sorted.
		if !swapped {
			break
		}

		// Otherwise, prepare for backward pass.
		swapped = false
		end--

		// Backward pass: bubble down the smallest element.
		for i := end; i > start; i-- {
			if arr[i-1] > arr[i] {
				arr[i-1], arr[i] = arr[i], arr[i-1]
				swapped = true
			}
		}
		start++
	}

	return arr
}
