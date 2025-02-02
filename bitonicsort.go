package sortcomparison

// BitonicSort implements an optimized bitonic sorting network.
// Best for parallel hardware; performs O(n log²n) comparisons.
func BitonicSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	// Find next power of 2 length
	n := len(arr)

	bitonicSort(arr, 0, n, true)
	return arr
}

/*
The parameter dir indicates the sorting direction, ASCENDING

	or DESCENDING; if (a[i] > a[j]) agrees with the direction,
	then a[i] and a[j] are interchanged.
*/
func compAndSwap(a []int, i int, j int, ascending bool) {
	if (ascending && a[i] > a[j]) || (!ascending && a[i] < a[j]) {
		a[i], a[j] = a[j], a[i]
	}
}

/*
It recursively sorts a bitonic sequence in ascending order,

	if dir = 1, and in descending order otherwise (means dir=0).
	The sequence to be sorted starts at index position low,
	the parameter cnt is the number of elements to be sorted.
*/
func bitonicMerge(a []int, low int, cnt int, ascending bool) {
	if cnt > 1 {
		k := cnt / 2
		for i := low; i < low+k; i++ {
			compAndSwap(a, i, i+k, ascending)
		}
		bitonicMerge(a, low, k, ascending)
		bitonicMerge(a, low+k, k, ascending)
	}
}

/*
This function first produces a bitonic sequence by recursively

	sorting its two halves in opposite sorting orders, and then
	calls bitonicMerge to make them in the same order
*/
func bitonicSort(a []int, low int, cnt int, ascending bool) {
	if cnt > 1 {
		k := cnt / 2

		// sort in ascending order
		bitonicSort(a, low, k, true)

		// sort in descending order
		bitonicSort(a, low+k, k, false)

		// Will merge whole sequence in ascending order
		bitonicMerge(a, low, cnt, ascending)
	}
}
