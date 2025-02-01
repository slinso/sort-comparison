package sortcomparison

// SpreadSort sorts a slice of non-negative integers using an adaptive recursion
// that partitions the array based on the binary representation of the numbers.
// It is a hybrid approach similar in spirit to SpreadSort where the input distribution
// is exploited to achieve high performance.
//
// How it works:
//   - It computes the most significant bit (MSB) needed to represent the maximum value.
//   - It recursively partitions the slice into two groups based on the presence of the bit
//     at the current position (0-group and 1-group).
//   - Once the partition length falls below a predefined threshold, insertion sort is used.
//   - This in-place, recursive approach minimizes memory allocations and is idiomatic Go.
//
// Time Complexity:
//   - Average-case: O(n) when the data is uniformly distributed.
//   - Worst-case: O(n log m), where m is the range of input values, if many bits are processed.
//
// Space Complexity:
//   - O(1) auxiliary space (in-place recursion; recursion depth is O(log m)).
//
// Note:
// - The algorithm assumes non-negative integers.
func SpreadSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	// Compute maximum value to determine the starting bit for partitioning.
	max := arr[0]
	for _, v := range arr {
		if v > max {
			max = v
		}
	}

	// Find the most significant bit (MSB) in 'max'
	bit := 0
	for max > 0 {
		bit++
		max >>= 1
	}
	// Adjust bit to be zero-indexed.
	bit--

	spreadSortHelper(arr, bit)
	return arr
}

// threshold below which insertion sort is used for efficiency.
const insertionThreshold = 16

// spreadSortHelper recursively partitions the array based on the 'bit' position.
func spreadSortHelper(arr []int, bit int) {
	if len(arr) < insertionThreshold {
		insertionSort(arr)
		return
	}
	if bit < 0 {
		// No further bits to partition on.
		return
	}

	mask := 1 << bit
	// Partition the array in-place into two groups:
	// - left part: numbers where the 'bit' is 0.
	// - right part: numbers where the 'bit' is 1.
	i, j := 0, len(arr)-1
	for i <= j {
		if (arr[i] & mask) == 0 {
			i++
		} else {
			arr[i], arr[j] = arr[j], arr[i]
			j--
		}
	}

	// Recursively sort the partitions based on the next lower bit.
	spreadSortHelper(arr[:i], bit-1)
	spreadSortHelper(arr[i:], bit-1)
}
