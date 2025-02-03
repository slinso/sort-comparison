package sortcomparison

// AdaptiveSort sorts a slice of integers using an adaptive natural merge sort.
// It scans the input for natural runs (monotonic sequences). Descending runs are reversed
// to become ascending. Then, the discovered runs are merged iteratively in a stable manner.
// This implementation is efficient, GC friendly (by preallocating a temporary buffer for merging),
// and idiomatic Go code.
//
// Time Complexity:
//   - Best-case: O(n) for nearly-sorted data (few long runs).
//   - Worst-case: O(n log n)
//
// Space Complexity:
//   - O(n) auxiliary space for merging (buffer reused across merges).
func AdaptiveSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	type run struct {
		start, end int
	}

	// Identify natural runs.
	var runs []run
	n := len(arr)
	i := 0
	for i < n {
		start := i
		i++
		// Determine run direction.
		if i < n {
			if arr[i-1] <= arr[i] {
				// Ascending run.
				for i < n && arr[i-1] <= arr[i] {
					i++
				}
			} else {
				// Descending run.
				for i < n && arr[i-1] > arr[i] {
					i++
				}
				// Reverse the descending run to make it ascending.
				reverseAdaptive(arr, start, i)
			}
		}
		runs = append(runs, run{start: start, end: i})
	}

	// Preallocate a temporary buffer for merging.
	temp := make([]int, len(arr))

	// Iteratively merge runs.
	for len(runs) > 1 {
		var newRuns []run
		for j := 0; j < len(runs); j += 2 {
			if j+1 < len(runs) {
				r1 := runs[j]
				r2 := runs[j+1]
				mergeAdaptive(arr, r1.start, r1.end, r2.end, temp)
				newRuns = append(newRuns, run{start: r1.start, end: r2.end})
			} else {
				newRuns = append(newRuns, runs[j])
			}
		}
		runs = newRuns
	}

	return arr
}

// reverse reverses the elements in arr from index start (inclusive) to end (exclusive).
func reverseAdaptive(arr []int, start, end int) {
	for i, j := start, end-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

// merge merges two consecutive sorted subarrays arr[low:mid] and arr[mid:high].
// The sorted elements are stored back into arr using the preallocated buffer temp.
func mergeAdaptive(arr []int, low, mid, high int, temp []int) {
	// Copy the left run to temp.
	copy(temp[low:mid], arr[low:mid])

	i, j, k := low, mid, low
	for i < mid && j < high {
		if temp[i] <= arr[j] {
			arr[k] = temp[i]
			i++
		} else {
			arr[k] = arr[j]
			j++
		}
		k++
	}
	// Copy any remaining elements from temp (left run) to arr.
	for i < mid {
		arr[k] = temp[i]
		i++
		k++
	}
}
