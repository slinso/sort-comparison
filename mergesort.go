package sortcomparison

/*
MergeSort Implementation (Divide and Conquer Sort)

Time Complexity:
  - Average: O(n log n)
  - Worst:   O(n log n)
  - Best:    O(n log n)

Space Complexity:
  - O(n) - requires temporary arrays for merging
  - Additional recursive call stack space O(log n)

Implementation Notes:
  - Stable sort - maintains relative order of equal elements
  - Divide-and-conquer strategy
  - Predictable performance regardless of input
  - Cache-inefficient due to non-local memory access
  - Well-suited for sorting linked lists
*/
func MergeSort(arr []int) {
	if len(arr) <= 1 {
		return
	}

	mid := len(arr) / 2

	// Create temp arrays
	left := make([]int, mid)
	right := make([]int, len(arr)-mid)

	// Copy data to temp arrays
	copy(left, arr[:mid])
	copy(right, arr[mid:])

	// Sort temp arrays
	MergeSort(left)
	MergeSort(right)

	// Merge sorted arrays
	mergeSort(arr, left, right)
}

func mergeSort(arr, left, right []int) {
	i, j, k := 0, 0, 0

	// Compare and merge
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			arr[k] = left[i]
			i++
		} else {
			arr[k] = right[j]
			j++
		}
		k++
	}

	// Copy remaining elements
	for i < len(left) {
		arr[k] = left[i]
		i++
		k++
	}

	for j < len(right) {
		arr[k] = right[j]
		j++
		k++
	}
}
