package sortcomparison

import "sort"

// PostmanSort sorts a slice of non-negative integers using a stable counting sort approach,
// often referred to as PostmanSort because it "delivers" elements into their proper buckets.
//
// How it works:
//   - It scans the array to determine the minimum and maximum values.
//   - It creates a count array of size (max - min + 1) to count occurrences.
//   - It computes cumulative counts to identify the final positions of each element.
//   - It iterates backwards through the input array to maintain stability when placing elements
//     in the output slice.
//   - The sorted output is copied back into the original array.
//
// Time Complexity:
//   - O(n + r), where n is the number of elements and r is the range of the input values.
//
// Space Complexity:
//   - O(n + r), due to the count array and auxiliary output slice.
func PostmanSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	// Find minimum and maximum values
	min, max := arr[0], arr[0]
	for _, v := range arr {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}

	rangeSize := max - min + 1

	// TODO: if the range is too large, use a different sorting algorithm
	if rangeSize > 1000000 {
		sort.Ints(arr)
		return arr
	}

	// Create and populate the count array
	count := make([]int, rangeSize)
	for _, v := range arr {
		count[v-min]++
	}

	// Compute cumulative count for positions
	for i := 1; i < rangeSize; i++ {
		count[i] += count[i-1]
	}

	// Build the output slice
	output := make([]int, len(arr))
	// Iterate backwards for stability
	for i := len(arr) - 1; i >= 0; i-- {
		val := arr[i]
		count[val-min]--
		output[count[val-min]] = val
	}

	// Copy the sorted output back to the original array
	copy(arr, output)
	return arr
}
