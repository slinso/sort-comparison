package sortcomparison

/*
StrandSort Implementation (Optimized for Memory Allocations)

Time Complexity:
  - Average: O(n²)
  - Worst:   O(n²) - when few ordered subsequences exist
  - Best:    O(n) - when input is already sorted

Space Complexity:
  - O(n) - requires additional storage for strands and merging buffers

Implementation Notes:
  - Extracts increasing subsequences (strands) and recursively merges them
  - Preallocates slice capacities to reduce reallocations
  - Stable sort – maintains relative order of equal elements
*/

func StrandSort(arr []int) {
	if len(arr) <= 1 {
		return
	}
	result := strandSortHelper(arr)
	copy(arr, result)
}

func strandSortHelper(input []int) []int {
	n := len(input)
	if n <= 1 {
		return input
	}

	// Preallocate with capacity equal to input length to avoid reallocations
	strand := make([]int, 0, n)
	remaining := make([]int, 0, n)

	// Add first element to strand
	strand = append(strand, input[0])

	// Build strand with increasing subsequence
	for i := 1; i < n; i++ {
		if input[i] >= strand[len(strand)-1] {
			strand = append(strand, input[i])
		} else {
			remaining = append(remaining, input[i])
		}
	}

	// If no remaining elements, the strand is fully sorted
	if len(remaining) == 0 {
		return strand
	}

	// Recursively sort remaining elements and merge strands
	return mergeStrand(strand, strandSortHelper(remaining))
}

func mergeStrand(left, right []int) []int {
	totalLen := len(left) + len(right)
	result := make([]int, 0, totalLen)
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	// Append any remaining elements
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}
