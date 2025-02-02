package sortcomparison

// WeaveMergeSort sorts the given slice of integers using a merge sort variant
// that "weaves" elements from two sorted halves in an unrolled merge loop.
// This technique improves cache locality and reduces branch mispredictions.
// It is efficient, GC friendly (performs a single allocation for a temporary buffer)
// and is idiomatic Go.
//
// Time Complexity:
//   - Average and Worst: O(n log n)
//
// Best suited when benefit from cache-friendly merging is desired.
func WeaveMergeSort(arr []int) []int {
	n := len(arr)
	if n < 2 {
		return arr
	}
	temp := make([]int, n)
	weaveMergeSortHelper(arr, temp, 0, n)
	return arr
}

// weaveMergeSortHelper recursively sorts the subarray arr[low:high].
func weaveMergeSortHelper(arr, temp []int, low, high int) {
	if high-low < 16 { // small slice threshold for insertion sort
		insertionSort(arr, low, high)
		return
	}
	mid := low + (high-low)/2
	weaveMergeSortHelper(arr, temp, low, mid)
	weaveMergeSortHelper(arr, temp, mid, high)
	weaveMerge(arr, temp, low, mid, high)
}

// weaveMerge merges two sorted subarrays [low, mid) and [mid, high)
// from temp into arr. The merge loop is unrolled in blocks to "weave" elements
// and improve CPU pipeline usage.
func weaveMerge(arr, temp []int, low, mid, high int) {
	// Copy the range to be merged into temp.
	copy(temp[low:high], arr[low:high])
	i, j, k := low, mid, low
	// Unrolled loop: process blocks of 4 elements at a time.
	for k < high {
		for block := 0; block < 4 && k < high; block++ {
			if i < mid && (j >= high || temp[i] <= temp[j]) {
				arr[k] = temp[i]
				i++
			} else {
				arr[k] = temp[j]
				j++
			}
			k++
		}
	}
}

// insertionSort performs an in-place insertion sort on arr[low:high].
func insertionSort(arr []int, low, high int) {
	for i := low + 1; i < high; i++ {
		key := arr[i]
		j := i - 1
		for j >= low && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}
