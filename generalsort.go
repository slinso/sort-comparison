package sortcomparison

// GeneralSort implements a robust, stable, general-purpose sorting algorithm inspired by Timsort.
// It detects natural runs in the input and then merges them using a stable merge procedure.
// For small runs, it applies insertion sort to ensure run length reaches a minimum threshold.
// This implementation is efficient, GC friendly (by minimizing allocations), and idiomatic Go code.
//
// Time Complexity:
//   - Best-case: O(n) for nearly sorted data.
//   - Worst-case: O(n log n)
//
// Space Complexity:
//   - O(n) auxiliary space for merging; the temporary buffer is allocated once.
func GeneralSort(arr []int) []int {
	n := len(arr)
	if n < 2 {
		return arr
	}

	const minRun = 32

	// Detect and process natural runs.
	var runs []run
	i := 0
	for i < n {
		runStart := i
		i++

		// Identify whether current run is ascending or descending.
		if i < n {
			if arr[i-1] <= arr[i] {
				// Ascending run.
				for i < n && arr[i-1] <= arr[i] {
					i++
				}
			} else {
				// Descending run; reverse to make ascending.
				for i < n && arr[i-1] > arr[i] {
					i++
				}
				reverseGeneral(arr, runStart, i)
			}
		}

		// If run length is less than minRun, extend run using insertion sort.
		runEnd := i
		if runEnd-runStart < minRun {
			limit := runStart + minRun
			if limit > n {
				limit = n
			}
			insertionSortRangeGeneral(arr, runStart, limit)
			runEnd = limit
			i = runEnd
		}

		runs = append(runs, run{start: runStart, end: runEnd})
	}

	// Allocate a temporary buffer once for merging.
	temp := make([]int, n)

	// Iteratively merge runs.
	for len(runs) > 1 {
		var mergedRuns []run
		// Merge pairs of runs sequentially.
		for j := 0; j < len(runs); j += 2 {
			if j+1 < len(runs) {
				r1 := runs[j]
				r2 := runs[j+1]
				mergeGeneral(arr, temp, r1.start, r1.end, r2.end)
				mergedRuns = append(mergedRuns, run{start: r1.start, end: r2.end})
			} else {
				// Odd run out remains.
				mergedRuns = append(mergedRuns, runs[j])
			}
		}
		runs = mergedRuns
	}

	return arr
}

// run defines a contiguous sorted segment of the array.
type run struct {
	start, end int
}

// insertionSortRangeGeneral performs an in-place insertion sort on arr[from:to].
func insertionSortRangeGeneral(arr []int, from, to int) {
	for i := from + 1; i < to; i++ {
		key := arr[i]
		j := i - 1
		for j >= from && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

// mergeGeneral merges two consecutive sorted subarrays of arr:
// arr[left:mid] and arr[mid:right]. The merged result is written back into arr,
// using temp as a temporary buffer to minimize allocations.
func mergeGeneral(arr, temp []int, left, mid, right int) {
	i, j, k := left, mid, left
	// Copy left half into temp.
	copy(temp[left:mid], arr[left:mid])
	for i < mid && j < right {
		if temp[i] <= arr[j] {
			arr[k] = temp[i]
			i++
		} else {
			arr[k] = arr[j]
			j++
		}
		k++
	}
	// Copy any remaining elements from temp (left half).
	for i < mid {
		arr[k] = temp[i]
		i++
		k++
	}
}

// reverseGeneral reverses the elements in arr[from:to] in-place.
func reverseGeneral(arr []int, from, to int) {
	to-- // Adjust to be inclusive.
	for from < to {
		arr[from], arr[to] = arr[to], arr[from]
		from++
		to--
	}
}
