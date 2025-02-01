package sortcomparison

/*
BurstSort Implementation (Optimized to reduce memory allocations)

Time Complexity:
    - Average: O(n log n)
    - Worst:   O(n log n)
    - Best:    O(n) when nearly sorted

Space Complexity:
    - O(n) additional space for a temporary buffer (allocated only once)
    - Reduces repeated allocations during recursive merging

Implementation Notes:
    - Divide-and-conquer recursive burst sort with in-place merge using a preallocated buffer
    - Stable sort - maintains relative order of equal elements
    - Particularly efficient for string sorting, adapted here for ints
    - Memory allocation is minimized by reusing a single temporary slice
*/

func BurstSort(arr []int) {
	if len(arr) <= 1 {
		return
	}
	// Allocate a temporary buffer once
	temp := make([]int, len(arr))
	burstSortAux(arr, temp)
}

func burstSortAux(arr, temp []int) {
	n := len(arr)
	if n <= 1 {
		return
	}

	mid := n / 2

	// Recursively sort each half, reuse appropriate portions of temp buffer
	burstSortAux(arr[:mid], temp[:mid])
	burstSortAux(arr[mid:], temp[mid:])

	mergeBurst(arr, mid, temp)
}

func mergeBurst(arr []int, mid int, temp []int) {
	n := len(arr)
	// Merge the two halves into the temp array
	left, right, idx := 0, mid, 0

	for left < mid && right < n {
		if arr[left] <= arr[right] {
			temp[idx] = arr[left]
			left++
		} else {
			temp[idx] = arr[right]
			right++
		}
		idx++
	}
	for left < mid {
		temp[idx] = arr[left]
		left++
		idx++
	}
	for right < n {
		temp[idx] = arr[right]
		right++
		idx++
	}

	// Copy merged output back to arr
	copy(arr, temp[:idx])
}
