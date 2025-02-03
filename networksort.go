package sortcomparison

// NetworkSort sorts a slice of non-negative integers using a comparator-based sorting network.
// This implementation uses a recursive Batcher odd-even mergesort network. It requires the input
// length to be a power of two, so the algorithm pads the input with a sentinel value (maxInt)
// if needed, and then removes the padding after sorting.
//
// How it works:
//   - If the input length is not a power of two, it is padded with a large integer value.
//   - The padded array is sorted using a recursive sorting network (BatcherSort), which first
//     sorts two halves in opposite orders and then performs an odd-even merge.
//   - Finally, the padding is removed.
//
// Time Complexity:
//   - Worst-case: O(n (log n)^2)
//
// Space Complexity:
//   - O(n) due to padding, with in-place network comparisons minimizing additional allocations.
func NetworkSort(arr []int) []int {
	n := len(arr)
	if n < 2 {
		return arr
	}

	// Compute next power of 2
	pow2 := nextPowerOfTwo(n)
	maxInt := int(^uint(0) >> 1)

	// Create padded slice if necessary.
	padded := make([]int, pow2)
	copy(padded, arr)
	for i := n; i < pow2; i++ {
		padded[i] = maxInt
	}

	// Sort the padded array using the Batcher sorting network.
	batcherSort(padded, 0, pow2, true)

	// Copy the sorted part (exclude the padded sentinels).
	sorted := padded[:n]
	// Optionally, copy back into original slice.
	copy(arr, sorted)
	return arr
}

// batcherSort recursively sorts the slice a[lo:lo+n] using the Batcher odd-even mergesort network.
// dir = true sorts in ascending order, false in descending order.
func batcherSort(a []int, lo, n int, dir bool) {
	if n > 1 {
		m := n / 2
		// Sort first half in ascending order.
		batcherSort(a, lo, m, true)
		// Sort second half in descending order.
		batcherSort(a, lo+m, n-m, false)
		// Merge the two halves using odd-even merge.
		oddevenMerge(a, lo, n, 1, dir)
	}
}

// oddevenMerge performs an odd-even merge on the slice a[lo:lo+n] with a stride r.
// It merges the two sorted sequences into one sorted sequence in the order specified by dir.
func oddevenMerge(a []int, lo, n, r int, dir bool) {
	m := r * 2
	if m < n {
		oddevenMerge(a, lo, n, m, dir)
		oddevenMerge(a, lo+r, n, m, dir)
		for i := lo + r; i+m/2 <= lo+n-r; i += m {
			compareAndSwap(a, i, i+r, dir)
		}
	} else {
		compareAndSwap(a, lo, lo+r, dir)
	}
}
