package sortcomparison

// MergeSort sorts the slice using the merge sort algorithm
func MergeSort[T any](arr []T, less func(a, b T) bool) {
	// Create temporary array for merging
	temp := make([]T, len(arr))
	mergeSort(arr, temp, 0, len(arr)-1, less)
}

// mergeSort recursively splits and merges the array
func mergeSort[T any](arr []T, temp []T, left, right int, less func(a, b T) bool) {
	if left < right {
		mid := (left + right) / 2

		// Sort first and second halves
		mergeSort(arr, temp, left, mid, less)
		mergeSort(arr, temp, mid+1, right, less)

		// Merge the sorted halves
		merge(arr, temp, left, mid, right, less)
	}
}

// merge combines two sorted subarrays into one sorted array
func merge[T any](arr []T, temp []T, left, mid, right int, less func(a, b T) bool) {
	// Copy both parts to temp array
	for i := left; i <= right; i++ {
		temp[i] = arr[i]
	}

	i := left    // First subarray index
	j := mid + 1 // Second subarray index
	k := left    // Merged array index

	// Merge temp arrays back into arr[left..right]
	for i <= mid && j <= right {
		if less(temp[i], temp[j]) {
			arr[k] = temp[i]
			i++
		} else {
			arr[k] = temp[j]
			j++
		}
		k++
	}

	// Copy remaining elements of left half if any
	for i <= mid {
		arr[k] = temp[i]
		k++
		i++
	}
	// Note: remaining elements of right half are already in correct position
}
