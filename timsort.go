package sortcomparison

// TimSort sorts the input slice using the Timsort algorithm.
//
// Time Complexity:
//   - Average: O(n log n)
//   - Worst:   O(n log n)
//   - Best:    O(n)
//
// Space Complexity:
//   - O(n) - requires auxiliary space for merging runs
//   - In-place sorting algorithm
//
// Timsort is a hybrid sorting algorithm derived from merge sort and insertion sort.
// It is designed to perform well on many kinds of real-world data. Timsort is the
// default sorting algorithm in Python and Java.
func TimSort(arr []int) {
	n := len(arr)
	if n < 2 {
		return
	}

	// Compute minimum run length (similar to Python's Timsort)
	minRun := minRunLength(n)

	// Identify runs and sort short runs via insertion sort.
	// Each run is represented as a pair: [start, length].
	var runs [][2]int
	for i := 0; i < n; {
		runStart := i
		i++

		// Identify run: detect ascending or descending order.
		if i < n {
			if arr[i] < arr[i-1] {
				// Descending run.
				for i < n && arr[i] < arr[i-1] {
					i++
				}
				reverseTim(arr, runStart, i-1)
			} else {
				// Ascending run.
				for i < n && arr[i] >= arr[i-1] {
					i++
				}
			}
		}

		runLen := i - runStart
		// Extend run to at least minRun.
		if runLen < minRun {
			end := runStart + minRun
			if end > n {
				end = n
			}
			insertionSortTim(arr, runStart, end)
			runLen = end - runStart
			i = end
		}
		runs = append(runs, [2]int{runStart, runLen})
	}

	// Merge runs until one run remains.
	for len(runs) > 1 {
		var mergedRuns [][2]int
		for i := 0; i < len(runs); i += 2 {
			if i+1 < len(runs) {
				left := runs[i][0]
				mid := runs[i+1][0]
				right := mid + runs[i+1][1]
				mergeTim(arr, left, mid, right)
				mergedRuns = append(mergedRuns, [2]int{left, runs[i][1] + runs[i+1][1]})
			} else {
				mergedRuns = append(mergedRuns, runs[i])
			}
		}
		runs = mergedRuns
	}
}

// minRunLength computes the minimum run length for Timsort.
func minRunLength(n int) int {
	r := 0
	for n >= 64 {
		r |= n & 1
		n >>= 1
	}
	return n + r
}

// insertionSort sorts the subarray arr[left:right] using insertion sort.
func insertionSortTim(arr []int, left, right int) {
	for i := left + 1; i < right; i++ {
		key := arr[i]
		j := i - 1
		for j >= left && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

// reverse reverses the subarray arr[left:right] in place.
func reverseTim(arr []int, left, right int) {
	for left < right {
		arr[left], arr[right] = arr[right], arr[left]
		left++
		right--
	}
}

// merge merges two sorted subarrays: arr[left:mid] and arr[mid:right] into one sorted segment.
func mergeTim(arr []int, left, mid, right int) {
	n1, n2 := mid-left, right-mid
	L := make([]int, n1)
	R := make([]int, n2)
	copy(L, arr[left:mid])
	copy(R, arr[mid:right])
	i, j, k := 0, 0, left
	for i < n1 && j < n2 {
		if L[i] <= R[j] {
			arr[k] = L[i]
			i++
		} else {
			arr[k] = R[j]
			j++
		}
		k++
	}
	for i < n1 {
		arr[k] = L[i]
		i++
		k++
	}
	for j < n2 {
		arr[k] = R[j]
		j++
		k++
	}
}
