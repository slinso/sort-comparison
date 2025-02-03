package sortcomparison

// BitonicSort implements an optimized bitonic sorting network.
// Best for parallel hardware; performs O(n log²n) comparisons.
func BitonicSort(arr []int) {
	if len(arr) < 2 {
		return
	}

	// Find next power of 2 length
	n := len(arr)

	bitonicSort(arr, 0, n, true)
}

// bitonicMerge recursively sorts a bitonic sequence in ascending or descending order
func bitonicMerge(arr []int, low, cnt int, ascending bool) {
	if cnt > 1 {
		k := cnt / 2
		for i := low; i < low+k; i++ {
			compareAndSwap(arr, i, i+k, ascending)
		}
		bitonicMerge(arr, low, k, ascending)
		bitonicMerge(arr, low+k, k, ascending)
	}
}

// bitonicSort recursively sorts the array in ascending or descending order
func bitonicSort(arr []int, low, cnt int, ascending bool) {
	if cnt > 1 {
		k := cnt / 2
		bitonicSort(arr, low, k, true)
		bitonicSort(arr, low+k, k, false)
		bitonicMerge(arr, low, cnt, ascending)
	}
}
