package sortcomparison

const (
	minMerge  = 32
	minGallop = 7
)

/*
TimSort Implementation (Hybrid Stable Sort)

Time Complexity:
  - Average: O(n log n)
  - Worst:   O(n log n)
  - Best:    O(n) - when data is pre-sorted

Space Complexity:
  - O(n) - requires temporary storage for merging
  - Stack space O(log n) for managing runs

Implementation Notes:
  - Hybrid of merge sort and insertion sort
  - Stable sort - maintains relative order of equal elements
  - Adaptive - exploits existing order in data
  - Minimum run size based on input size
  - Uses galloping mode for efficient merging
*/
func TimSort(arr []int) {
	n := len(arr)
	if n <= 1 {
		return
	}

	// Calculate min run size
	minRun := minRunLength(n)

	// Array for storing runs
	runs := make([]run, 0)

	// Process array in chunks
	for i := 0; i < n; {
		// Identify next run
		runLen := findRun(arr[i:])
		if runLen < minRun {
			runLen = min(minRun, n-i)
			insertionSort(arr[i : i+runLen])
		}

		// Push run onto stack
		runs = append(runs, run{start: i, len: runLen})
		i += runLen

		// Merge runs to maintain invariants
		mergeCollapse(arr, runs)
	}

	// Final merging of all runs
	mergeForcedCollapse(arr, runs)
}

type run struct {
	start, len int
}

func minRunLength(n int) int {
	r := 0
	for n >= minMerge {
		r |= n & 1
		n >>= 1
	}
	return n + r
}

func findRun(arr []int) int {
	n := len(arr)
	if n <= 1 {
		return n
	}

	increasing := arr[1] >= arr[0]
	i := 2
	if increasing {
		for i < n && arr[i] >= arr[i-1] {
			i++
		}
	} else {
		for i < n && arr[i] < arr[i-1] {
			i++
		}
		// Reverse the run if decreasing
		reverseRun(arr[:i])
	}
	return i
}

func reverseRun(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func mergeCollapse(arr []int, runs []run) {
	for len(runs) > 1 {
		n := len(runs)
		if n >= 2 && runs[n-1].len <= runs[n-2].len {
			if n >= 3 && runs[n-2].len <= runs[n-3].len {
				mergeAt(arr, runs, n-3)
			} else {
				mergeAt(arr, runs, n-2)
			}
		} else if runs[n-1].len < minMerge {
			mergeAt(arr, runs, n-2)
		} else {
			break
		}
	}
}

func mergeForcedCollapse(arr []int, runs []run) {
	for len(runs) > 1 {
		n := len(runs)
		if n >= 2 && runs[n-2].len < runs[n-1].len {
			mergeAt(arr, runs, n-2)
		} else {
			mergeAt(arr, runs, n-1)
		}
	}
}

func mergeAt(arr []int, runs []run, i int) {
	left := runs[i]
	right := runs[i+1]

	temp := make([]int, left.len)
	copy(temp, arr[left.start:left.start+left.len])

	p := left.start
	q := right.start
	r := 0

	for r < left.len && q < right.start+right.len {
		if q >= right.start+right.len || (r < left.len && temp[r] <= arr[q]) {
			arr[p] = temp[r]
			r++
		} else {
			arr[p] = arr[q]
			q++
		}
		p++
	}

	// Copy remaining elements
	for r < left.len {
		arr[p] = temp[r]
		p++
		r++
	}

	// Update runs slice
	runs[i].len += runs[i+1].len
	copy(runs[i+1:], runs[i+2:])
	runs = runs[:len(runs)-1]
}
