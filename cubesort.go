package sortcomparison

import (
	"math"
)

// CubeSort sorts the input slice using a multi–way merge strategy.
// The algorithm divides the array into blocks of size d (where
// d is roughly the cube root of n) and uses insertion sort to sort
// each block. Then it repeatedly merges d sorted blocks at a time
// using a d–way merge until the entire array is sorted.
//
// Time Complexity:
//   - Average: O(n log n), with a logarithm base d = ∛n.
//   - Worst:   O(n log n)
//   - Best:    O(n) for nearly sorted inputs (due to insertion sort).
//
// Space Complexity:
//   - O(n) extra space for a temporary buffer (allocated once).
//
// Implementation Notes:
//   - Uses a “ping–pong” temporary buffer to avoid repeated allocations.
//   - Performs an initial block sort (using insertion sort) to reduce work.
//   - Merges d blocks at a time using an in–place d–way merge.
func CubeSort(arr []int) {
	n := len(arr)
	if n <= 32 {
		insertionSort(arr)
		return
	}

	// Compute d as ceil(cuberoot(n)). Ensure d >= 2.
	d := int(math.Cbrt(float64(n)))
	if d < 2 || d*d*d < n {
		d++
	}
	blockSize := d // initial block size

	// Initial block sort using insertion sort.
	for i := 0; i < n; i += blockSize {
		end := i + blockSize
		if end > n {
			end = n
		}
		insertionSort(arr[i:end])
	}

	// Allocate a temporary buffer once.
	temp := make([]int, n)

	// Merge in passes. In each pass, we merge groups of d blocks,
	// each block being of current size.
	currentSize := blockSize
	// isSrc indicates where the current sorted data is (true means in arr,
	// false means in temp). After each merge pass we swap.
	isSrc := true
	for currentSize < n {
		if isSrc {
			// Merge from arr into temp.
			for start := 0; start < n; start += currentSize * d {
				multiWayMerge(arr, temp, start, currentSize, d, n)
			}
		} else {
			// Merge from temp into arr.
			for start := 0; start < n; start += currentSize * d {
				multiWayMerge(temp, arr, start, currentSize, d, n)
			}
		}
		isSrc = !isSrc
		currentSize *= d
	}
	// If final sorted data is not in arr, copy it from temp.
	if !isSrc {
		copy(arr, temp)
	}
}

// multiWayMerge merges up to d sorted blocks from src into dst.
// Each block has maximum length "size". The merge is performed
// on the interval starting at index "start" and ending at min(n, start+d*size).
func multiWayMerge(src, dst []int, start, size, d, n int) {
	end := start + d*size
	if end > n {
		end = n
	}
	// For each block j (0 <= j < d), set up indices and end positions.
	indices := make([]int, d)
	seqStart := make([]int, d)
	seqEnd := make([]int, d)
	for j := 0; j < d; j++ {
		s := start + j*size
		if s >= n {
			seqStart[j] = n
			seqEnd[j] = n
			indices[j] = n
		} else {
			seqStart[j] = s
			e := s + size
			if e > n {
				e = n
			}
			seqEnd[j] = e
			indices[j] = s
		}
	}
	// Perform the d–way merge.
	for pos := start; pos < end; pos++ {
		minFound := false
		var minVal int
		minSeq := -1
		for j := 0; j < d; j++ {
			if indices[j] < seqEnd[j] {
				val := src[indices[j]]
				if !minFound || val < minVal {
					minVal = val
					minSeq = j
					minFound = true
				}
			}
		}
		// Place the smallest value into dst.
		dst[pos] = minVal
		indices[minSeq]++
	}
}
