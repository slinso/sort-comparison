package sortcomparison

// CascadeSort finds naturally ordered runs in the slice and then
// merges them in cascading passes until the entire slice is sorted.
// It is efficient for partially sorted data, GC friendly (single aux allocation)
// and idiomatic Go code.
func CascadeSort(arr []int) []int {
	nLen := len(arr)
	if nLen < 2 {
		return arr
	}

	// Identify natural runs.
	// Each run is represented as a two-element array: [start, end)
	var runs [][2]int
	start := 0
	for start < nLen {
		end := start + 1
		// Extend run as long as next element is in order.
		for end < nLen && arr[end-1] <= arr[end] {
			end++
		}
		runs = append(runs, [2]int{start, end})
		start = end
	}

	// If only one run exists, array is sorted.
	if len(runs) == 1 {
		return arr
	}

	// Allocate temporary buffer once for merging.
	temp := make([]int, nLen)

	// Cascade merge: merge adjacent runs until one run remains.
	for len(runs) > 1 {
		var newRuns [][2]int
		for i := 0; i < len(runs); i += 2 {
			// If a pair is available, merge them.
			if i+1 < len(runs) {
				leftRun := runs[i]
				rightRun := runs[i+1]
				merge(arr, temp, leftRun[0], leftRun[1], rightRun[1])
				newRuns = append(newRuns, [2]int{leftRun[0], rightRun[1]})
			} else {
				// Otherwise, carry over the last run.
				newRuns = append(newRuns, runs[i])
			}
		}
		runs = newRuns
	}

	return arr
}

// merge merges two consecutive sorted subarrays:
// arr[low:mid] and arr[mid:high] using temp as auxiliary storage.
func merge(arr, temp []int, low, mid, high int) {
	// Copy the segment to be merged into temp.
	copy(temp[low:high], arr[low:high])
	i, j := low, mid
	for k := low; k < high; k++ {
		if i < mid && (j >= high || temp[i] <= temp[j]) {
			arr[k] = temp[i]
			i++
		} else {
			arr[k] = temp[j]
			j++
		}
	}
}
