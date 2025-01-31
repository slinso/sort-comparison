package sortcomparison

/*
BubbleSort Implementation

Time Complexity:
  - Average: O(n²)
  - Worst:   O(n²) - when array is reverse sorted
  - Best:    O(n) - when array is already sorted

Space Complexity:
  - O(1) - only requires a single additional variable for swapping

Implementation Notes:
  - Simple in-place sorting algorithm
  - Stable sort - maintains relative order of equal elements
  - Cache-friendly due to sequential access pattern
  - Practical for very small arrays (n < 50)
  - Adaptive - stops early when array becomes sorted
*/
func BubbleSort(arr []int) {
	n := len(arr)
	if n <= 1 {
		return
	}

	for i := 0; i < n-1; i++ {
		swapped := false
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
}
